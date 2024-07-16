package validation

import (
	"encoding/json"
	"errors"

	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/defenseunicorns/go-oscal/src/pkg/files"
	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
	"github.com/santhosh-tekuri/jsonschema/v6"
)

type Validator struct {
	documentPath     string
	jsonMap          map[string]interface{}
	schemaVersion    string
	modelType        string
	validationResult ValidationResult
}

// NewValidator returns a validator with the models version of the schema.
func NewValidator(oscalDoc model.InterfaceOrBytes) (validator Validator, err error) {
	jsonMap, err := model.CoerceToJsonMap(oscalDoc)
	if err != nil {
		return validator, err
	}

	modelType, err := model.GetModelType(jsonMap)
	if err != nil {
		return validator, err
	}

	version, err := versioning.GetOscalVersionFromMap(jsonMap)
	if err != nil {
		return validator, err
	}

	return Validator{
		jsonMap:       jsonMap,
		schemaVersion: version,
		modelType:     modelType,
	}, nil
}

// NewValidatorDesiredVersion returns a validator with the desired version of the schema.
func NewValidatorDesiredVersion(oscalDoc model.InterfaceOrBytes, desiredVersion string) (validator Validator, err error) {
	jsonMap, err := model.CoerceToJsonMap(oscalDoc)
	if err != nil {
		return validator, err
	}

	modelType, err := model.GetModelType(jsonMap)
	if err != nil {
		return validator, err
	}

	formattedVersion := versioning.FormatOscalVersion(desiredVersion)

	if err = versioning.IsValidOscalVersion(formattedVersion); err != nil {
		return validator, err
	}

	return Validator{
		jsonMap:       jsonMap,
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

// IsLatestOscalVersion returns true if the model is the latest version of the OSCAL schema.
func (v *Validator) IsLatestOscalVersion() (bool, error) {
	latestVersion := versioning.GetLatestSupportedVersion()

	version, err := versioning.GetOscalVersionFromMap(v.jsonMap)
	if err != nil {
		return false, err
	}

	return version == latestVersion, nil
}

// Validate validates the model against the schema.
func (v *Validator) Validate() error {
	compiler := jsonschema.NewCompiler()
	compiler.UseLoader(schemas.NewSchemaLoader())
	sch, err := compiler.Compile(v.schemaVersion)
	if err != nil {
		return err
	}

	err = sch.Validate(v.jsonMap)
	if err != nil {
		// If the error is not a `ValidationError`, return the error
		validationErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			return err
		}

		jsonBytes, err := json.Marshal(validationErr)
		files.WriteOutput(jsonBytes, "../../../testdata/validation/basic-error-v6.json")

		// Extract the specific errors from the schema error
		basicErrors := ExtractErrors(v.jsonMap, *validationErr)
		// Set the validation result
		v.validationResult = NewValidationResult(v, basicErrors)
		return err
	}

	v.validationResult = NewValidationResult(v, []ValidatorError{})
	return nil
}
