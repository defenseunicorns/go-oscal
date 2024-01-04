package validation

import (
	"embed"
	"encoding/json"
	"errors"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validationError"
	"github.com/defenseunicorns/go-oscal/src/pkg/validationResult"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schema/*.json
var schemas embed.FS

const (
	SCHEMA_PREFIX = "oscal_complete_schema-"
)

type Validator struct {
	jsonMap          map[string]interface{}
	schemaVersion    string
	modelType        string
	validationResult validationResult.ValidationResult
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
		jsonMap:          model,
		schemaVersion:    version,
		modelType:        modelType,
		validationResult: validationResult.ValidationResult{},
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
		jsonMap:          model,
		modelType:        modelType,
		schemaVersion:    formattedVersion,
		validationResult: validationResult.ValidationResult{},
	}, nil
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

// GetValidationResult returns the result of the validation.
func (v *Validator) GetValidationResult() validationResult.ValidationResult {
	return v.validationResult
}

// Validate validates the model against the schema.
func (v *Validator) Validate() error {
	// Build the schema file-path
	schemaPath := SCHEMA_PREFIX + strings.ReplaceAll(v.schemaVersion, ".", "-") + ".json"

	schemaBytes, err := schemas.ReadFile("schema/" + schemaPath)
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
		basicErrors := validationError.ExtractErrors(v.jsonMap, validationErr.BasicOutput())
		v.validationResult = validationResult.CreateValidationResult(basicErrors)
		formattedErrors, _ := json.MarshalIndent(basicErrors, "", "  ")
		return errors.New(string(formattedErrors))
	}

	v.validationResult = validationResult.CreateValidationResult([]validationError.ValidatorError{})
	return nil
}
