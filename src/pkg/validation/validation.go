package validation

import (
	"errors"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

const (
	NON_SCHEMA_ERROR_ABSOLUTE_KEYWORD_LOCATION = "non-schema-error"
)

type ValidationParams struct {
	// This is the name of the model to validate (used for documentation and error reporting purposes) (required)
	ModelType string `json:"model-type,omitempty" yaml:"model-type,omitempty"`
	// This is the reference of the model to validate (does not have to be a url, used for documentation and error reporting purposes) (optional)
	ModelPath string `json:"model-path,omitempty" yaml:"model-path,omitempty"`
	// This is the version of the model to validate against (for example, 1.0.0) (optional)
	ModelVersion string `json:"model-version,omitempty" yaml:"model-version,omitempty"`
	// Must be a json map (map[string]interface{}) (required)
	ModelData map[string]interface{} `json:"model-data,omitempty" yaml:"model-data,omitempty"`
	// This is the reference of the schema to validate against (does not have to be a url, used for documentation and error reporting purposes) (required)
	SchemaPath string `json:"schema-path,omitempty" yaml:"schema-path,omitempty"`
	// Must be a json map (map[string]interface{}) (required)
	SchemaData map[string]interface{} `json:"schema-data,omitempty" yaml:"schema-data,omitempty"`
	// This is the version of the schema to validate against (for example, 1.0.0) (optional)
	SchemaVersion string `json:"schema-version,omitempty" yaml:"schema-version,omitempty"`
}

func ValidateModelAgainstSchema(params *ValidationParams) (*ValidationResult, error) {
	validationResult := handleRequirements(params)
	if !validationResult.Valid {
		return validationResult, GetNonSchemaError(validationResult)
	}

	validatorErrors := []ValidatorError{}

	compiler := jsonschema.NewCompiler()

	compiler.AddResource(params.SchemaPath, params.SchemaData)

	sch, err := compiler.Compile(params.SchemaPath)
	if err != nil {
		return NewNonSchemaValidationError(err, params), err
	}

	err = sch.Validate(params.ModelData)
	if err != nil {
		// If the error is not a `ValidationError`, return the error
		validationErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			return NewNonSchemaValidationError(err, params), err
		}

		// Extract the specific errors from the schema error
		validatorErrors = ExtractErrors(params.ModelData, validationErr.DetailedOutput())
		return NewValidationResultFromParams(params, validatorErrors), err
	}

	return NewValidationResultFromParams(params, validatorErrors), nil
}

func handleRequirements(params *ValidationParams) *ValidationResult {
	if params == nil {
		return NewNonSchemaValidationError(errors.New("params are required"), params)
	}
	if params.ModelType == "" {
		return NewNonSchemaValidationError(errors.New("model-type is required"), params)
	}
	if params.ModelData == nil {
		return NewNonSchemaValidationError(errors.New("model-data is required"), params)
	}
	if params.SchemaPath == "" {
		return NewNonSchemaValidationError(errors.New("schema-url is required"), params)
	}
	if params.SchemaData == nil {
		return NewNonSchemaValidationError(errors.New("schema-data is required"), params)
	}
	return &ValidationResult{
		Valid: true,
	}
}

// NewNonSchemaValidationError creates a system validation error
func NewNonSchemaValidationError(err error, params *ValidationParams) *ValidationResult {
	return &ValidationResult{
		Valid:     false,
		TimeStamp: time.Now(),
		Errors: []ValidatorError{
			{
				Error:                   err.Error(),
				AbsoluteKeywordLocation: NON_SCHEMA_ERROR_ABSOLUTE_KEYWORD_LOCATION,
			},
		},
		Metadata: NewValidationResultMetadataFromParams(params),
	}
}

// NewValidationResultFromParams creates a new validation result from the params
func NewValidationResultFromParams(params *ValidationParams, valErrs []ValidatorError) *ValidationResult {
	if params == nil {
		return NewNonSchemaValidationError(errors.New("params are required"), nil)
	}
	valid := len(valErrs) == 0
	return &ValidationResult{
		Valid:     valid,
		TimeStamp: time.Now(),
		Errors:    valErrs,
		Metadata:  NewValidationResultMetadataFromParams(params),
	}
}

func NewValidationResultMetadataFromParams(params *ValidationParams) ValidationResultMetadata {
	if params == nil {
		return ValidationResultMetadata{}
	}
	return ValidationResultMetadata{
		DocumentType:    params.ModelType,
		DocumentPath:    params.ModelPath,
		DocumentVersion: params.ModelVersion,
		SchemaVersion:   params.SchemaVersion,
	}
}

// IsNonSchemaValidationError checks if the result is a system validation error
func IsNonSchemaValidationError(result *ValidationResult) bool {
	if result == nil {
		return false
	}
	return len(result.Errors) == 1 && result.Errors[0].AbsoluteKeywordLocation == NON_SCHEMA_ERROR_ABSOLUTE_KEYWORD_LOCATION
}

// GetNonSchemaError extracts the system validation error
// If the result is not a system validation error or if there are no errors, return nil
func GetNonSchemaError(result *ValidationResult) error {
	if result == nil {
		return nil
	}
	if !IsNonSchemaValidationError(result) {
		return nil
	}
	return errors.New(result.Errors[0].Error)
}
