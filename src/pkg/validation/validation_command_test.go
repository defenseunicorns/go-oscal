package validation_test

import (
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

func TestValidationCommand(t *testing.T) {
	t.Parallel()

	t.Run("returns an error if no input file is provided", func(t *testing.T) {
		_, err := validation.ValidationCommand("")
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error if the input file is not a json or yaml file", func(t *testing.T) {
		_, err := validation.ValidationCommand("test.txt")
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error if the input file is not a valid oscal document", func(t *testing.T) {
		_, err := validation.ValidationCommand("test.json")
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns a warning if the oscal version is not the latest", func(t *testing.T) {
		validationResponse, err := validation.ValidationCommand(gooscaltest.ValidComponentPath)
		if err != nil {
			t.Error("expected an error, got nil")
		}

		if len(validationResponse.Warnings) == 0 {
			t.Error("expected a warning, got none")
		}
	})

	t.Run("returns a validation response if the input file is valid", func(t *testing.T) {
		validationResponse, err := validation.ValidationCommand(gooscaltest.ValidComponentPath)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if validationResponse.Result.Valid != true {
			t.Error("expected a valid result, got invalid")
		}
	})

	t.Run("returns an error and validation result if the input file fails validation", func(t *testing.T) {
		validationResponse, err := validation.ValidationCommand(gooscaltest.InvalidCatalogPath)
		if err == nil {
			t.Error("expected an error, got nil")
		}

		if validationResponse.Result.Valid != false {
			t.Error("expected an invalid result, got valid")
		}
	})
}
