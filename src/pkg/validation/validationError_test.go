package validation_test

import (
	"os"
	"strings"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"gopkg.in/yaml.v3"
)

func TestExtractErrors(t *testing.T) {
	t.Parallel()
	var catalog map[string]interface{}
	// Read Catalog and detailedOutput errors.
	originalDocBytes, err := os.ReadFile(gooscaltest.InvalidCatalogPath)
	if err != nil {
		t.Fatal(err)
	}
	// Unmarshal into associated structs
	if err := yaml.Unmarshal(originalDocBytes, &catalog); err != nil {
		t.Fatal(err)
	}

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

	t.Run("returns an array of Validator errors given the original document and detailedOutput errors", func(t *testing.T) {
		t.Parallel()
		detailedOutput := validatorErr.DetailedOutput()
		validatorErrors := validation.ExtractErrors(catalog, detailedOutput)
		if len(validatorErrors) == 0 {
			t.Errorf("expected validator errors, got nil")
		}
	})

	t.Run("filters out errors with empty InstanceLocations", func(t *testing.T) {
		t.Parallel()
		detailedOutput := validatorErr.DetailedOutput()
		validatorErrors := validation.ExtractErrors(catalog, detailedOutput)
		for _, err := range validatorErrors {
			if err.InstanceLocation == "" {
				t.Error("expected errors with empty InstanceLocation to be filtered")
			}
		}
	})

	t.Run("filters out empty errors", func(t *testing.T) {
		t.Parallel()
		detailedOutput := validatorErr.DetailedOutput()
		validatorErrors := validation.ExtractErrors(catalog, detailedOutput)
		for _, err := range validatorErrors {
			if err.Error == "" {
				t.Error("expected empty errors to be filtered")
			}
		}
	})

	t.Run("filters out errors starting with 'doesn't validate with' as these are not specific to the value and therefore not useful", func(t *testing.T) {
		t.Parallel()
		detailedOutput := validatorErr.DetailedOutput()
		validatorErrors := validation.ExtractErrors(catalog, detailedOutput)
		for _, err := range validatorErrors {
			if strings.HasPrefix(err.Error, "doesn't validate with") {
				t.Error("expected errors starting with 'doesn't validate with' to be filtered")
			}
		}
	})

	t.Run("returns FailedValue if not a map or a slice", func(t *testing.T) {
		t.Parallel()
		// Create a sample invalid catalog with specific values that will fail validation
		invalidCatalog := map[string]interface{}{
			"metadata": map[string]interface{}{
				"title": "Invalid Catalog",
				"uuid":  "invalid-uuid", // This should trigger a validation error
			},
			"groups": []interface{}{
				map[string]interface{}{
					"id": "group-1",
				},
			},
		}

		// Validate the invalid catalog
		err := sch.Validate(invalidCatalog)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		validatorErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			t.Fatalf("expected validation error, got %T", err)
		}

		detailedOutput := validatorErr.DetailedOutput()
		validationErrors := validation.ExtractErrors(invalidCatalog, detailedOutput)
		for _, err := range validationErrors {
			if err.FailedValue != nil {
				_, mapOk := err.FailedValue.(map[string]interface{})
				_, sliceOk := err.FailedValue.([]interface{})
				if mapOk || sliceOk {
					t.Errorf("expected FailedValue to be nil for maps and slices, got %v", err.FailedValue)
				}
			}
		}
	})

	t.Run("returns nil if the FailedValue is a map[string]interface{}", func(t *testing.T) {
		t.Parallel()
		detailedOutput := validatorErr.DetailedOutput()
		validationErrors := validation.ExtractErrors(catalog, detailedOutput)
		for _, err := range validationErrors {
			if _, ok := err.FailedValue.(map[string]interface{}); ok {
				if err.FailedValue != nil {
					t.Errorf("expected nil, got %v", err.FailedValue)
				}
			}
		}
	})

	t.Run("returns nil if the FailedValue is a []interface{}", func(t *testing.T) {
		t.Parallel()
		detailedOutput := validatorErr.DetailedOutput()
		validationErrors := validation.ExtractErrors(catalog, detailedOutput)
		for _, err := range validationErrors {
			if _, ok := err.FailedValue.([]interface{}); ok {
				if err.FailedValue != nil {
					t.Errorf("expected nil, got %v", err.FailedValue)
				}
			}
		}
	})

	t.Run("Adds top-level errors back in if there are no errors with instance locations", func(t *testing.T) {
		t.Parallel()
		// Create a sample invalid catalog with top-level validation errors
		invalidCatalog := map[string]interface{}{
			"metadata": "Invalid Metadata", // This should trigger a top-level validation error
		}

		// Validate the invalid catalog
		err := sch.Validate(invalidCatalog)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		validatorErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			t.Fatalf("expected validation error, got %T", err)
		}

		detailedOutput := validatorErr.DetailedOutput()
		validationErrors := validation.ExtractErrors(invalidCatalog, detailedOutput)
		if len(validationErrors) == 0 {
			t.Errorf("expected top-level validation errors, got nil")
		}
		for _, err := range validationErrors {
			if err.InstanceLocation != "" {
				t.Errorf("expected top-level error with empty InstanceLocation, got %v", err.InstanceLocation)
			}
		}
	})
}
