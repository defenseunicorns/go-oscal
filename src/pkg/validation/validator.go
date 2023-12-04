package validation

import (
	"embed"
	"encoding/json"
	"errors"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schema/*.json
var schemas embed.FS

const (
	SCHEMA_PREFIX = "oscal_complete_schema-"
)

type validator struct {
	jsonMap   map[string]interface{}
	version   string
	modelType string
}

// InterfaceOrBytes is an interface{} or []byte for generic functions that can support either type
type InterfaceOrBytes interface {
	interface{} | []byte
}

func NewValidator(oscalDoc InterfaceOrBytes) (*validator, error) {
	model, err := coerceToJsonMap(oscalDoc)
	if err != nil {
		return nil, err
	}

	modelType, err := getModelType(model)
	if err != nil {
		return nil, err
	}

	version, err := getOscalVersionFromMap(model)
	if err != nil {
		return nil, err
	}

	return &validator{
		jsonMap:   model,
		version:   version,
		modelType: modelType,
	}, nil
}

func NewValidatorDesiredVersion(oscalDoc InterfaceOrBytes, desiredVersion string) (*validator, error) {
	model, err := coerceToJsonMap(oscalDoc)
	if err != nil {
		return nil, err
	}

	modelType, err := getModelType(model)
	if err != nil {
		return nil, err
	}

	formattedVersion := formatOscalVersion(desiredVersion)

	if err = isValidOscalVersion(formattedVersion); err != nil {
		return nil, err
	}

	return &validator{
		jsonMap:   model,
		modelType: modelType,
		version:   formattedVersion,
	}, nil
}

func (v *validator) GetVersion() string {
	return v.version
}

func (v *validator) GetJsonModel() map[string]interface{} {
	return v.jsonMap
}

func (v *validator) GetModelType() string {
	return v.modelType
}

func (v *validator) Validate() error {
	// Build the schema file-path
	schemaPath := SCHEMA_PREFIX + strings.ReplaceAll(v.version, ".", "-") + ".json"

	schemaBytes, err := schemas.ReadFile("schema/" + schemaPath)
	if err != nil {
		return err
	}

	sch, err := jsonschema.CompileString(v.version, string(schemaBytes))
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
		basicErrors := extractErrors(v.jsonMap, validationErr.BasicOutput())
		formattedErrors, _ := json.MarshalIndent(basicErrors, "", "  ")
		return errors.New(string(formattedErrors))
	}

	return nil
}
