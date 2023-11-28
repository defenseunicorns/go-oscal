package validation

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

//go:embed schema/*.json
var schemas embed.FS

const (
	SCHEMA_PREFIX         = "oscal_complete_schema-"
	DEFAULT_OSCAL_VERSION = "1.0.4"
)

var (
	versionRegexp    = regexp.MustCompile(`^\d+([-\.]\d+){2}$`)
	supportedVersion = map[string]bool{
		"1.0.4": true,
		"1.0.5": true,
		"1.0.6": true,
		"1.1.0": true,
		"1.1.1": true,
	}
)

// InterfaceOrBytes is an interface{} or []byte for generic functions that can support either type
type InterfaceOrBytes interface {
	interface{} | []byte
}

// IsValidOscal takes an interface{} or []byte and returns error if the yaml/json is not valid.
func IsValidOscal[T InterfaceOrBytes](docBytes T) (modelJson map[string]interface{}, err error) {
	modelJson, err = CoerceToJSONForTypeSafety(docBytes)
	if err != nil {
		return modelJson, err
	}
	version, err := GetOscalVersionFromMap(modelJson)
	if err != nil {
		return modelJson, err
	}

	return ValidateOscalAgainstVersion(version, modelJson)
}

// ValidateOscalAgainstVersion takes a version string and a []byte or interface{} and returns true if the yaml/json is valid for the specified oscal-version
func ValidateOscalAgainstVersion[T InterfaceOrBytes](oscalVersion string, docBytes T) (modelJson map[string]interface{}, err error) {
	modelJson, err = CoerceToJSONForTypeSafety(docBytes)
	if err != nil {
		return modelJson, err
	}

	if err = IsValidOscalVersion(oscalVersion); err != nil {
		return modelJson, err
	}

	// Build the schema file-path
	schemaPath := SCHEMA_PREFIX + strings.ReplaceAll(oscalVersion, ".", "-") + ".json"

	schemaBytes, err := schemas.ReadFile("schema/" + schemaPath)
	if err != nil {
		return modelJson, err
	}

	sch, err := jsonschema.CompileString(oscalVersion, string(schemaBytes))
	if err != nil {
		return modelJson, err
	}

	err = sch.Validate(modelJson)
	if err != nil {
		// If the error is not a validation error, return the error
		validationErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			return modelJson, err
		}

		// TODO: More succinct error message.
		// Return the detailed output of the validation error
		formattedError, _ := json.MarshalIndent(validationErr.DetailedOutput(), "", "  ")
		return modelJson, errors.New(string(formattedError))
	}

	// Successful validation
	return modelJson, nil
}

// CoerceToJSONForTypeSafety takes a yaml byte array and coerces it to a json interface{}
// This is necessary because the jsonschema library only accepts valid json data types that may not match yaml.
// Example: yaml allows for DateTimes to be time.Time, but json requires them to be strings
// This also allows for structs to be passed in, and they will be converted to map[string]interface{}
func CoerceToJSONForTypeSafety[T InterfaceOrBytes](ymlData T) (model map[string]interface{}, err error) {
	model, err = convertInterfaceOrBytesToMap(ymlData)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

// GetOscalVersionFromMap returns formatted OSCAL version if valid version is passed, returns error if not.
func GetOscalVersionFromMap(model map[string]interface{}) (version string, err error) {
	var submodel map[string]interface{}
	for _, value := range model {
		submodel = value.(map[string]interface{})
		break
	}

	metadata, ok := submodel["metadata"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("required field: metadata not found")
	}

	if _, ok := metadata["oscal-version"]; !ok {
		return "", fmt.Errorf("required field: oscal-version not found")
	}

	version, ok = reflect.ValueOf(metadata["oscal-version"]).Interface().(string)
	if !ok {
		return "", fmt.Errorf("required field: oscal-version not found")
	}

	return version, nil
}

// IsValidOscalVersion returns true if the version is supported, false if not.
func IsValidOscalVersion(version string) error {
	version = FormatOscalVersion(version)

	if !versionRegexp.MatchString(version) {
		return fmt.Errorf("version %s is not a valid version", version)
	}

	if !supportedVersion[version] {
		return fmt.Errorf("version %s is not supported", version)
	}
	return nil
}

// FormatOscalVersion takes a version string and returns a formatted version string
func FormatOscalVersion(version string) (formattedVersion string) {
	formattedVersion = strings.ToLower(version)
	formattedVersion = strings.ReplaceAll(formattedVersion, "-", ".")
	formattedVersion = strings.Trim(formattedVersion, "v")
	return formattedVersion
}

// convertInterfaceOrBytesToMap takes an InterfaceOrByte and returns a map[string]interface{}
func convertInterfaceOrBytesToMap[T InterfaceOrBytes](incomingModel T) (model map[string]interface{}, err error) {
	// Check if interface{} and can be coerced to map[string]interface{}
	model, ok := reflect.ValueOf(incomingModel).Interface().(map[string]interface{})
	if !ok {
		// Check if []byte
		ymlBytes, ok := reflect.ValueOf(incomingModel).Interface().([]byte)
		// If not []byte or map[string]interface{}, marshal to []byte
		if !ok {
			ymlBytes, err = yaml.Marshal(incomingModel)
			if err != nil {
				return nil, err
			}
		}
		// Unmarshal to map[string]interface{}
		err = yaml.Unmarshal(ymlBytes, &model)
		if err != nil {
			return nil, err
		}
	}
	return model, nil
}
