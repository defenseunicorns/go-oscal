package validation

import (
	"fmt"

	"github.com/defenseunicorns/go-oscal/src/pkg/files"
	"github.com/defenseunicorns/go-oscal/src/pkg/model"
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
	bytes, err := files.ReadOscalFile(inputFile)
	if err != nil {
		return validationResponse, err
	}
	return ValidationCommandWithModel(inputFile, bytes)
}

// ValidationCommandMulti validates an OSCAL document with multiple models ie (Catalog, Component, Profile, System)
// Ignores "oneOf" constraint for oscalTypes_X_X_X.OscalCompleteSchema
// Does not handle yaml delimited files
// Returns a ValidationResponse and an error
func ValidationCommandMulti(inputFile string) (responses []ValidationResponse, err error) {
	// Read the input file
	bytes, err := files.ReadOscalFile(inputFile)
	if err != nil {
		return responses, err
	}

	// Coerce the bytes to a json map
	jsonModel, err := model.CoerceToJsonMap(bytes)
	if err != nil {
		return responses, err
	}

	// Generate the models
	models, err := model.GenModels(jsonModel)
	if err != nil {
		return responses, err
	}

	// Validate each model
	for _, m := range models {
		// get the key (model type), at this point is the only one
		modelType, _ := model.GetModelType(m)

		// TODO: get input on path format (current: multiple_model.json#component-definition)
		pathWithType := fmt.Sprintf("%s#%s", inputFile, modelType)
		validationResponse, err := ValidationCommandWithModel(pathWithType, m)
		// TODO: should we bail on non-validation errors?
		if err != nil {
			return responses, err
		}
		responses = append(responses, validationResponse)
	}

	return responses, nil
}

// ValidationCommandWithModel validates an OSCAL document with a file path and a InterfaceOrBytes
// Returns a ValidationResponse and an error
func ValidationCommandWithModel(inputFile string, jsonModel model.InterfaceOrBytes) (validationResponse ValidationResponse, err error) {
	// Create and set the validator in the validation response
	validator, err := NewValidator(jsonModel)
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
