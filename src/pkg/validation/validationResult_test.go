package validation_test

import (
	"reflect"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

func TestValidationResult(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)

	t.Run("NewValidationResult", func(t *testing.T) {
		t.Parallel()

		validator, err := validation.NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		t.Run("returns a ValidationResult", func(t *testing.T) {
			t.Parallel()
			validationResult := validation.NewValidationResult(&validator, []validation.ValidatorError{})
			if _, ok := reflect.ValueOf(validationResult).Interface().(validation.ValidationResult); !ok {
				t.Errorf("expected ValidationResult, got nil")
			}
		})

		t.Run("sets the ValidationResult.Valid to true if there are no errors", func(t *testing.T) {
			t.Parallel()
			validationResult := validation.NewValidationResult(&validator, []validation.ValidatorError{})
			if !validationResult.Valid {
				t.Errorf("expected ValidationResult.Valid to be true, got false")
			}
		})

		t.Run("sets the ValidationMetadata.DocumentType to the type of the model", func(t *testing.T) {
			t.Parallel()
			validationResult := validation.NewValidationResult(&validator, []validation.ValidatorError{})
			if validationResult.Metadata.DocumentType != "component-definition" {
				t.Errorf("expected ValidationMetadata.DocumentType to be %s, got %s", validator.GetModelType(), validationResult.Metadata.DocumentType)
			}
		})

		t.Run("sets the ValidationMetadata.DocumentVersion to the version of the model", func(t *testing.T) {
			t.Parallel()
			validationResult := validation.NewValidationResult(&validator, []validation.ValidatorError{})
			if validationResult.Metadata.DocumentVersion != "1.0.4" {
				t.Errorf("expected ValidationMetadata.DocumentVersion to be %s, got %s", validator.GetSchemaVersion(), validationResult.Metadata.DocumentVersion)
			}
		})

		t.Run("sets the ValidationMetadata.SchemaVersion to the version of the schema", func(t *testing.T) {
			t.Parallel()
			validationResult := validation.NewValidationResult(&validator, []validation.ValidatorError{})
			if validationResult.Metadata.SchemaVersion != "1.0.4" {
				t.Errorf("expected ValidationMetadata.SchemaVersion to be %s, got %s", validator.GetSchemaVersion(), validationResult.Metadata.SchemaVersion)
			}
		})
	})
}
