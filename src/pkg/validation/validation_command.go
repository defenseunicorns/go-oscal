package validation

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
)

type ValidationResponse struct {
	Validator Validator
	Result    ValidationResult
	Warnings  []string
}

// ValidationCommand validates an OSCAL document
// Returns a ValidationResponse and an error
func ValidationCommand(inputFile string) (validationResponse ValidationResponse, err error) {
	// Validate the input file is a json or yaml file
	if !strings.HasSuffix(inputFile, "json") && !strings.HasSuffix(inputFile, "yaml") {
		return validationResponse, errors.New("please specify a json or yaml file")
	}

	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return validationResponse, fmt.Errorf("reading input file: %s", err)
	}

	// Create and set the validator in the validation response
	validator, err := NewValidator(bytes)
	if err != nil {
		return validationResponse, fmt.Errorf("failed to create validator: %s", err)
	}
	validationResponse.Validator = validator

	// Get and set version warnings
	version := validator.GetSchemaVersion()
	err = versioning.VersionWarning(version)
	if err != nil {
		validationResponse.Warnings = append(validationResponse.Warnings, err.Error())
	}

	// Set the document path
	validator.SetDocumentPath(inputFile)

	// Run the validation
	validationError := validator.Validate()

	// Write validation result if it was specified and exists before returning ValidateCommand error
	validationResult, _ := validator.GetValidationResult()
	validationResponse.Result = validationResult

	// Handle the validation error
	if validationError != nil {
		return validationResponse, fmt.Errorf("failed to validate %s version %s: %s", validator.GetModelType(), validator.GetSchemaVersion(), err)
	}

	return validationResponse, nil
}
