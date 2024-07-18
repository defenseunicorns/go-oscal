package validation

import (
	"errors"

	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
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

// GetValidationParams returns the validation params for the validator.
func (v *Validator) GetValidationParams() (*ValidationParams, error) {
	modelVersion, err := versioning.GetOscalVersionFromMap(v.jsonMap)
	if err != nil {
		modelVersion = ""
	}
	schemaModel, err := schemas.GetOscalMapFromVersion(v.schemaVersion)
	if err != nil {
		return nil, err
	}

	return &ValidationParams{
		ModelType:     v.modelType,
		ModelPath:     v.documentPath,
		ModelVersion:  modelVersion,
		ModelData:     v.jsonMap,
		SchemaVersion: v.schemaVersion,
		SchemaPath:    schemas.CreateSchemaPath(v.schemaVersion),
		SchemaData:    schemaModel.(map[string]interface{}),
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
	params, err := v.GetValidationParams()
	if err != nil {
		return err
	}
	validationResult, err := ValidateModelAgainstSchema(params)
	v.validationResult = *validationResult
	return err
}
