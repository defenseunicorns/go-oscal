package utils_test

import (
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	oscalTypes_1_0_4 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	"gopkg.in/yaml.v3"
)

var yamlString = `component-definition:
  metadata:
    oscal-version: 1.0.4`

var jsonString = `{"component-definition": {"metadata": {"oscal-version": "1.0.4"}}}`

var (
	unsupportedVersion    = "-2.0.7"
	validVersion          = "1.0.4"
	tooManyVersionNumbers = "1.0.4.1"
	upgradeVersion        = "1.0.6"
)

func TestModelUtils(t *testing.T) {
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

	t.Run("findValue", func(t *testing.T) {
		t.Parallel()

		var catalog map[string]interface{}
		originalDocBytes, err := os.ReadFile(gooscaltest.InvalidCatalogPath)
		if err != nil {
			t.Fatal(err)
		}

		// Unmarshal into associated structs
		if err := yaml.Unmarshal(originalDocBytes, &catalog); err != nil {
			t.Fatal(err)
		}

		t.Run("returns the value given a model and a slice of keys", func(t *testing.T) {
			t.Parallel()
			value := utils.FindValue(catalog, []string{"catalog", "metadata", "oscal-version"})
			if value != "1.0.4" {
				t.Errorf("expected %s, got %s", "1.0.4", value)
			}
		})

		t.Run("converts key to int if the previous key is a slice", func(t *testing.T) {
			t.Parallel()
			value := utils.FindValue(catalog, []string{"catalog", "metadata", "parties", "0", "uuid"})
			if value != "invalid-uuid" {
				t.Errorf("expected %s, got %s", "invalid-uuid", value)
			}
		})

		t.Run("returns nil given a model and a slice of keys that do not exist", func(t *testing.T) {
			t.Parallel()
			value := utils.FindValue(catalog, []string{"catalog", "metadata", "oscal-version", "invalid-key"})
			if value != nil {
				t.Errorf("expected nil, got %v", value)
			}
		})

	})

	t.Run("getModelType", func(t *testing.T) {
		t.Parallel()
		t.Run("returns the model type given a model.", func(t *testing.T) {
			t.Parallel()
			modelType, err := utils.GetModelType(validModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if modelType != "component-definition" {
				t.Errorf("expected %s, got %s", "component-definition", modelType)
			}
		})

		t.Run("throws error when model has more than one key", func(t *testing.T) {
			t.Parallel()
			_, err := utils.GetModelType(invalidOscalModel)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("throws error for empty model", func(t *testing.T) {
			t.Parallel()
			_, err := utils.GetModelType(map[string]interface{}{})
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("throws error for nil model", func(t *testing.T) {
			t.Parallel()
			_, err := utils.GetModelType(nil)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

	})

	t.Run("coerceToJsonMap", func(t *testing.T) {
		t.Parallel()
		var (
			componentDefinition = oscalTypes_1_0_4.OscalModels{}
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
			jsonMap, err := utils.CoerceToJsonMap(yamlBytes)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("returns a json map when a valid yaml interface{} is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := utils.CoerceToJsonMap(yamlModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("returns a json map when a valid json byte array is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := utils.CoerceToJsonMap(bytes)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("returns a json map when a valid struct is passed", func(t *testing.T) {
			t.Parallel()
			jsonMap, err := utils.CoerceToJsonMap(componentDefinition)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if jsonMap["component-definition"] == nil {
				t.Errorf("expected json map, got nil")
			}
		})

	})

}
