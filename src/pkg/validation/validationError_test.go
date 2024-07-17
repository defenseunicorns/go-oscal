package validation_test

import (
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"gopkg.in/yaml.v3"
)

func TestValidationError(t *testing.T) {
	t.Parallel()
	var catalog map[string]interface{}
	// Read Catalog and basic errors.
	originalDocBytes, err := os.ReadFile(gooscaltest.InvalidCatalogPath)
	if err != nil {
		t.Fatal(err)
	}
	// Unmarshal into associated structs
	if err := yaml.Unmarshal(originalDocBytes, &catalog); err != nil {
		t.Fatal(err)
	}

	t.Run("extractErrors", func(t *testing.T) {
		validator, err := validation.NewValidator(catalog)
		if err != nil {
			t.Fatal(err)
		}
		sch, err := schemas.GetOscalSchema(validator.GetSchemaVersion())
		if err != nil {
			t.Fatal(err)
		}
		err = sch.Validate(catalog)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		validatorErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			t.Fatalf("expected validation error, got %T", err)
		}

		t.Run("returns an array of Validator errors given the original document and basic errors", func(t *testing.T) {
			t.Parallel()
			outputError := validatorErr.DetailedOutput()
			validatorErrors := validation.ExtractErrors(catalog, *outputError)
			if len(validatorErrors) == 0 {
				t.Errorf("expected validator errors, got nil")
			}
		})

		t.Run("returns nil if the FailedValue is a map[string]interface{}", func(t *testing.T) {
			t.Parallel()
			outputError := validatorErr.DetailedOutput()
			validationErrors := validation.ExtractErrors(catalog, *outputError)
			if validationErrors[0].FailedValue != nil {
				t.Errorf("expected nil, got %v", validationErrors[0].FailedValue)
			}
		})

		t.Run("returns nil if the FailedValue is a []interface{}", func(t *testing.T) {
			t.Parallel()
			outputError := validatorErr.DetailedOutput()
			validationErrors := validation.ExtractErrors(catalog, *outputError)
			if validationErrors[3].FailedValue != nil {
				t.Errorf("expected nil, got %v", validationErrors[3].FailedValue)
			}
		})

	})
}
