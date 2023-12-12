package validation

import (
	"embed"
	"encoding/json"
	"errors"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed schema/*.json
var schemas embed.FS

const (
	SCHEMA_PREFIX = "oscal_complete_schema-"
)

type Validator struct {
	jsonMap   map[string]interface{}
	version   string
	modelType string
}

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

	return Validator{
		jsonMap:   model,
		version:   version,
		modelType: modelType,
	}, nil
}

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

	if err = utils.IsValidOscalVersion(formattedVersion); err != nil {
		return validator, err
	}

	return Validator{
		jsonMap:   model,
		modelType: modelType,
		version:   formattedVersion,
	}, nil
}

func (v *Validator) GetVersion() string {
	return v.version
}

func (v *Validator) GetJsonModel() map[string]interface{} {
	return v.jsonMap
}

func (v *Validator) GetModelType() string {
	return v.modelType
}

func (v *Validator) Validate() error {
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
		basicErrors := utils.ExtractErrors(v.jsonMap, validationErr.BasicOutput())
		formattedErrors, _ := json.MarshalIndent(basicErrors, "", "  ")
		return errors.New(string(formattedErrors))
	}

	return nil
}
