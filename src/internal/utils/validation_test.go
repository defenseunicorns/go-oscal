package utils_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	"github.com/santhosh-tekuri/jsonschema/v5"
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

	t.Run("ValidationErrors", func(t *testing.T) {
		t.Parallel()
		var catalog map[string]interface{}
		var basic jsonschema.Basic
		// Read Catalog and basic errors.
		originalDocBytes, err := os.ReadFile(gooscaltest.InvalidCatalogPath)
		if err != nil {
			t.Fatal(err)
		}
		errorsBytes, err := os.ReadFile(gooscaltest.BasicErrorPathJson)
		if err != nil {
			t.Fatal(err)
		}
		// Unmarshal into associated structs
		if err := yaml.Unmarshal(originalDocBytes, &catalog); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(errorsBytes, &basic); err != nil {
			t.Fatal(err)
		}

		t.Run("findValue", func(t *testing.T) {
			t.Parallel()

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

		t.Run("extractErrors", func(t *testing.T) {

			t.Run("returns an array of Validator errors given the original document and basic errors", func(t *testing.T) {
				t.Parallel()
				validatorErrors := utils.ExtractErrors(catalog, basic)
				if len(validatorErrors) == 0 {
					t.Errorf("expected validator errors, got nil")
				}
			})

			t.Run("filters out errors with empty InstanceLocations", func(t *testing.T) {
				t.Parallel()
				validatorErrors := utils.ExtractErrors(catalog, basic)
				if validatorErrors[0].InstanceLocation == "" {
					t.Error("expected first Basic.Error to be filtered")
				}
			})

			t.Run("filters out empty errors", func(t *testing.T) {
				t.Parallel()
				validatorErrors := utils.ExtractErrors(catalog, basic)
				if validatorErrors[0].Error == basic.Errors[0].Error {
					t.Error("expected second Basic.Error to be filtered")
				}
			})

			t.Run("filters out errors starting with 'doesn't validate with' as these are not specific to the value and therefore not useful", func(t *testing.T) {
				t.Parallel()
				validatorErrors := utils.ExtractErrors(catalog, basic)
				if validatorErrors[0].Error == basic.Errors[0].Error {
					t.Error("expected third Basic.Error to be filtered")
				}
			})

			t.Run("returns FailedValue if not a map or a slice", func(t *testing.T) {
				t.Parallel()
				validationErrors := utils.ExtractErrors(catalog, basic)
				if validationErrors[2].FailedValue != "invalid-uuid" {
					t.Errorf("expected string, got %v", validationErrors[2].FailedValue)
				}
			})

			t.Run("returns nil if the FailedValue is a map[string]interface{}", func(t *testing.T) {
				t.Parallel()
				validationErrors := utils.ExtractErrors(catalog, basic)
				if validationErrors[0].FailedValue != nil {
					t.Errorf("expected nil, got %v", validationErrors[0].FailedValue)
				}
			})

			t.Run("returns nil if the FailedValue is a []interface{}", func(t *testing.T) {
				t.Parallel()
				validationErrors := utils.ExtractErrors(catalog, basic)
				if validationErrors[3].FailedValue != nil {
					t.Errorf("expected nil, got %v", validationErrors[3].FailedValue)
				}
			})

		})
	})
}
