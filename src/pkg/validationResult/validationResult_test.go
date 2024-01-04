package validationResult_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/validationError"
	"github.com/defenseunicorns/go-oscal/src/pkg/validationResult"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

func TestValidationResult(t *testing.T) {
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

	validationErrors := validationError.ExtractErrors(catalog, basic)

	t.Run("CreateValidationResult", func(t *testing.T) {
		t.Parallel()
		t.Run("sets success to true if no errors are provided", func(t *testing.T) {
			t.Parallel()
			result := validationResult.CreateValidationResult([]validationError.ValidatorError{})
			if result.Success != true {
				t.Errorf("expected true, got %v", result.Success)
			}
		})

		t.Run("sets success to false if errors are provided", func(t *testing.T) {
			t.Parallel()
			result := validationResult.CreateValidationResult(validationErrors)
			if result.Success != false {
				t.Errorf("expected false, got %v", result.Success)
			}
		})
	})
}
