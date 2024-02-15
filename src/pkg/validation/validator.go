package validation

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Validator struct {
	documentPath     string
	jsonMap          map[string]interface{}
	schemaVersion    string
	modelType        string
	validationResult ValidationResult
}

// NewValidator returns a validator with the models version of the schema.
func NewValidator(oscalDoc utils.InterfaceOrBytes) (validator Validator, err error) {
	model, err := utils.CoerceToJsonMap(oscalDoc)
	if err != nil {
		return validator, err
	}

	modelType, err := utils.GetModelType(model)
	if err != nil {
		return validator, err
	}

	version, err := utils.GetOscalVersionFromMap(model)
	if err != nil {
		return validator, err
	}
	utils.VersionWarning(version)

	return Validator{
		jsonMap:       model,
		schemaVersion: version,
		modelType:     modelType,
	}, nil
}

// NewValidatorDesiredVersion returns a validator with the desired version of the schema.
func NewValidatorDesiredVersion(oscalDoc utils.InterfaceOrBytes, desiredVersion string) (validator Validator, err error) {
	model, err := utils.CoerceToJsonMap(oscalDoc)
	if err != nil {
		return validator, err
	}

	modelType, err := utils.GetModelType(model)
	if err != nil {
		return validator, err
	}

	formattedVersion := utils.FormatOscalVersion(desiredVersion)
	utils.VersionWarning(formattedVersion)

	if err = utils.IsValidOscalVersion(formattedVersion); err != nil {
		return validator, err
	}

	return Validator{
		jsonMap:       model,
		modelType:     modelType,
		schemaVersion: formattedVersion,
	}, nil
}

// SetDocumentPath sets the path of the document being validated.
func (v *Validator) SetDocumentPath(documentPath string) {
	v.documentPath = documentPath
}

// GetDocumentPath returns the path of the document being validated.
func (v *Validator) GetDocumentPath() string {
	return v.documentPath
}

// GetSchemaVersion returns the version of the schema used to validate the model.
func (v *Validator) GetSchemaVersion() string {
	return v.schemaVersion
}

// GetJsonModel returns the model being validated.
func (v *Validator) GetJsonModel() map[string]interface{} {
	return v.jsonMap
}

// GetModelType returns the type of the model being validated.
func (v *Validator) GetModelType() string {
	return v.modelType
}

// GetValidationResult returns a ValidationResult.
// If the validation has not been run, an error is returned.
func (v *Validator) GetValidationResult() (ValidationResult, error) {
	if v.validationResult.TimeStamp == (ValidationResult{}).TimeStamp {
		return v.validationResult, errors.New("validation has not been run")
	}
	return v.validationResult, nil
}

// Validate validates the model against the schema.
func (v *Validator) Validate() error {
	// Build the schema file-path
	schemaPath := schemas.SCHEMA_PREFIX + strings.ReplaceAll(v.schemaVersion, ".", "-") + ".json"

	schemaBytes, err := schemas.Schemas.ReadFile(schemaPath)
	if err != nil {
		return err
	}

	sch, err := jsonschema.CompileString(v.schemaVersion, string(schemaBytes))
	if err != nil {
		return err
	}

	err = sch.Validate(v.jsonMap)
	if err != nil {
		// If the error is not a validation error, return the error
		validationErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			return err
		}

		// Extract the specific errors from the schema error
		// Return the errors as a string
		basicErrors := ExtractErrors(v.jsonMap, validationErr.BasicOutput())
		v.validationResult = NewValidationResult(v, basicErrors)
		formattedErrors, _ := json.MarshalIndent(basicErrors, "", "  ")
		return errors.New(string(formattedErrors))
	}

	v.validationResult = NewValidationResult(v, []ValidatorError{})
	return nil
}
