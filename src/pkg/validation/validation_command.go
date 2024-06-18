package validation

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type ValidationResponse struct {
	Validator Validator
	// Parsed validation result
	Result ValidationResult
	// Non-failing go-oscal warnings (ie: deprecated fields, newer schema versions, etc)
	Warnings []string
	// Unparsed Failing validation errors from the jsonschema library
	JsonSchemaError *jsonschema.ValidationError
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

	// Set the document path
	validator.SetDocumentPath(inputFile)

	// Run the validation
	err = validator.Validate()
	if err != nil {
		validationError, ok := err.(*jsonschema.ValidationError)
		// If the error is not a validation error, return the error
		if !ok {
			return validationResponse, err
		}
		// Set the jsonschema error in the validation response
		validationResponse.JsonSchemaError = validationError
	}

	// Get and set version warnings if upgrade available
	version := validator.GetSchemaVersion()
	err = versioning.VersionWarning(version)
	if err != nil {
		validationResponse.Warnings = append(validationResponse.Warnings, err.Error())
	}

	// Get the validation result
	validationResponse.Result, err = validator.GetValidationResult()
	// If there is an error, return it, but there shouldn't be an error
	// Referenced in [#268](https://github.com/defenseunicorns/go-oscal/issues/268)
	if err != nil {
		return validationResponse, fmt.Errorf("shouldn't error,check the GetValidationResult method for more information: %s", err)
	}

	return validationResponse, nil
}
