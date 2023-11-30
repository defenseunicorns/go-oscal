package validation

import (
	"testing"

	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	"gopkg.in/yaml.v3"
)

var yamlString = `component-definition:
  metadata:
    oscal-version: 1.0.4`

var jsonString = `{"component-definition": {"metadata": {"oscal-version": "1.0.4"}}}`

func TestValidationUtils(t *testing.T) {
	t.Parallel()

	validModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": validVersion,
			},
		},
	}

	invalidOscalModel := map[string]interface{}{
		"component-definition": map[string]interface{}{},
		"oscal-version":        validVersion,
	}

	noOscalVersionModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": "",
			},
		},
	}

	t.Run("getModelType", func(t *testing.T) {
		t.Parallel()
		t.Run("returns the model type given a model.", func(t *testing.T) {
			t.Parallel()
			modelType, err := getModelType(validModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if modelType != "component-definition" {
				t.Errorf("expected %s, got %s", "component-definition", modelType)
			}
		})

		t.Run("throws error when model has more than one key", func(t *testing.T) {
			t.Parallel()
			_, err := getModelType(invalidOscalModel)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("throws error for empty model", func(t *testing.T) {
			t.Parallel()
			_, err := getModelType(map[string]interface{}{})
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("throws error for nil model", func(t *testing.T) {
			t.Parallel()
			_, err := getModelType(nil)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

	})

	t.Run("getOscalVersionFromMap", func(t *testing.T) {
		t.Parallel()
		t.Run("returns the oscal version given a model.", func(t *testing.T) {
			t.Parallel()
			version, err := getOscalVersionFromMap(validModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if version != validVersion {
				t.Errorf("expected %s, got %s", validVersion, "1.0.4")
			}
		})

		t.Run("throws error when version is not found", func(t *testing.T) {
			t.Parallel()
			_, err := getOscalVersionFromMap(invalidOscalModel)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("throws error for empty version", func(t *testing.T) {
			t.Parallel()
			_, err := getOscalVersionFromMap(noOscalVersionModel)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})
	})

	t.Run("isValidOscalVersion", func(t *testing.T) {
		t.Parallel()
		var (
			validVersion       = "1.0.4"
			unsupportedVersion = "1.0.7"
			invalidVersion     = "invalid-version-1-0-4"
		)

		t.Run("returns nil when a valid version is passed", func(t *testing.T) {
			t.Parallel()
			err := isValidOscalVersion(validVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("returns error when an unsupported version is passed", func(t *testing.T) {
			t.Parallel()
			err := isValidOscalVersion(unsupportedVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when an invalid version is passed", func(t *testing.T) {
			t.Parallel()
			err := isValidOscalVersion(invalidVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})
	})

	t.Run("formatOscalVersion", func(t *testing.T) {
		t.Parallel()
		var (
			validVersion  = "1.0.4"
			vPrefixed     = "v1.0.4"
			capVPrefixed  = "V1.0.4"
			dashedVersion = "1-0-4"
		)

		t.Run("returns formatted version when a valid version is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := formatOscalVersion(validVersion)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})

		t.Run("returns formatted version when a valid version prefixed with 'v' is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := formatOscalVersion(vPrefixed)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})

		t.Run("returns formatted version when a valid version prefixed with 'V' is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := formatOscalVersion(capVPrefixed)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})

		t.Run("returns formatted version when a valid version with dashes is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := formatOscalVersion(dashedVersion)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})
	})

	t.Run("coerceToJsonMap", func(t *testing.T) {
		t.Parallel()
		var (
			componentDefinition = V104.OscalModels{}
			yamlModel           = map[string]interface{}{}
		)
		bytes := []byte(jsonString)

		yamlBytes := []byte(yamlString)

		if err := yaml.Unmarshal(bytes, &componentDefinition); err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(yamlBytes, &yamlModel); err != nil {
			panic(err)
		}

		t.Run("returns a json map when a valid yaml byte array is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := coerceToJsonMap(yamlBytes)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("returns a json map when a valid yaml interface{} is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := coerceToJsonMap(yamlModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("returns a json map when a valid json byte array is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := coerceToJsonMap(bytes)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("returns a json map when a valid struct is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := coerceToJsonMap(componentDefinition)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

	})
}
