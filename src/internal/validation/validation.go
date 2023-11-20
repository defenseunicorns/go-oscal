package validation

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	V105 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-5"
	V106 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-6"
	V110 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-0"
	V111 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-1"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

//go:embed schema/*.json
var schemas embed.FS

const SCHEMA_PREFIX = "oscal_complete_schema-"
const DEFAULT_OSCAL_VERSION = "1.0.4"

var versionRegexp = regexp.MustCompile(`^\d+([-\.]\d+){2}$`)

var supportedVersion = map[string]bool{
	"1.0.4": true,
	"1.0.5": true,
	"1.0.6": true,
	"1.1.0": true,
	"1.1.1": true,
}

// InterfaceOrBytes is an interface{} or []byte for generic functions that can support either type
type InterfaceOrBytes interface {
	interface{} | []byte
}

// GetModels returns an instance of the OSCAL model for the specified version.
// Returns an empty map[string]interface{} if the version is not supported.
func GetVersionedModel(version string) interface{} {
	switch version {
	case "1.0.4":
		return V104.OscalModels{}
	case "1.0.5":
		return V105.OscalModels{}
	case "1.0.6":
		return V106.OscalModels{}
	case "1.1.0":
		return V110.OscalModels{}
	case "1.1.1":
		return V111.OscalModels{}
	default:
		return map[string]interface{}{}
	}
}

// IsValidSchemaVersion takes a version string and a []byte or interface{} and returns true if the yaml/json is valid for the specified oscal-version
func IsValidSchemaVersion[T InterfaceOrBytes](oscalVersion string, docBytes T) (err error) {
	modelJson, err := CoerceToJSONForTypeSafety(oscalVersion, docBytes)
	if err != nil {
		return err
	}

	compiler := jsonschema.NewCompiler()
	schemaPath := SCHEMA_PREFIX + strings.ReplaceAll(oscalVersion, ".", "-") + ".json"
	schemaBytes, err := schemas.ReadFile("schema/" + schemaPath)
	if err != nil {
		return err
	}
	compiler.AddResource(schemaPath, strings.NewReader(string(schemaBytes)))
	if err != nil {
		return err
	}
	sch, err := compiler.Compile(schemaPath)
	if err != nil {
		return err
	}
	err = sch.Validate(modelJson)
	if err != nil {

		validationErr, ok := err.(*jsonschema.ValidationError)
		if !ok {
			return err
		}
		// TODO: More succinct error message.
		formattedError, _ := json.MarshalIndent(validationErr.DetailedOutput(), "", "  ")
		return errors.New(string(formattedError))
	}
	return nil
}

// CoerceToJSONForTypeSafety takes a yaml byte array and coerces it to a json interface{}
// This is necessary because the jsonschema library only accepts valid json data types that may not match yaml.
// Example: yaml allows for DateTimes to be time.Time, but json requires them to be strings
// This also allows for structs to be passed in, and they will be converted to map[string]interface{}
func CoerceToJSONForTypeSafety[T InterfaceOrBytes](version string, ymlData T) (model map[string]interface{}, err error) {
	model, err = ConvertInterfaceOrBytesToJson(ymlData)
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

// GetOscalVersionFromModel takes an interface{} or []byte and returns the metadata.oscal_version as a string
func GetOscalVersionFromModel[T InterfaceOrBytes](incomingModel T) (version string, err error) {
	// Check if interface{} and can be coerced to map[string]interface{}
	model, err := ConvertInterfaceOrBytesToJson(incomingModel)
	if err != nil {
		return "", err
	}

	// find the oscal-version field
	version, err = findOscalVersion(model)
	if err != nil {
		return "", err
	}
	return version, nil
}

// FormatCmdLineOscalVersion returns formatted OSCAL version if valid version is passed, returns error if not.
func FormatCmdLineOscalVersion(cmdLineVersion string) (formattedVersion string, err error) {
	if cmdLineVersion == "" {
		return DEFAULT_OSCAL_VERSION, nil
	}

	formattedVersion = strings.ToLower(cmdLineVersion)
	formattedVersion = strings.TrimPrefix(formattedVersion, "v")
	formattedVersion = strings.ReplaceAll(formattedVersion, "-", ".")

	if !isVersionRegexp(formattedVersion) {
		return formattedVersion, fmt.Errorf("version %s is not a valid version", cmdLineVersion)
	}

	if !supportedVersion[formattedVersion] {
		return formattedVersion, fmt.Errorf("version %s is not supported", cmdLineVersion)
	}

	return formattedVersion, nil
}

func isVersionRegexp(v string) bool {
	return versionRegexp.MatchString(v)
}

// ConvertInterfaceOrBytesToJson takes an InterfaceOrByte and returns a map[string]interface{}
func ConvertInterfaceOrBytesToJson[T InterfaceOrBytes](incomingModel T) (model map[string]interface{}, err error) {
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

// Recursively search for the oscal-version field
func deepFindOscalVersion(data map[string]interface{}) (string, error) {
	if version, ok := data["oscal-version"].(string); ok && version != "" {
		return version, nil
	}

	// If not found at the top level, recursively search nested structures
	for _, value := range data {
		if nestedMap, isMap := value.(map[string]interface{}); isMap {
			if version, err := findOscalVersion(nestedMap); err == nil {
				return version, nil
			}
		}
	}
	return "", fmt.Errorf("required field: oscal-version not found")
}

// Assumes that the data follows the OscalModel structure of any version.
// If that assumption is false, this function will do a deep recursive search for the oscal-version field
// Note: Not sure that deep search is necessary, as the oscal-version field is required at the top level of the model
func findOscalVersion(data map[string]interface{}) (version string, err error) {
	// Check for a top level field that has the Metadata substructure
	for _, value := range data {
		// Ensure the value is a map[string]interface{}
		modelType, ok := value.(map[string]interface{})
		if ok {
			// Check if the map has a metadata field
			if modelType["metadata"] != nil {
				// Ensure the metadata field is a map[string]interface{}
				metadata, ok := modelType["metadata"].(map[string]interface{})
				if ok {
					// Check if the metadata field has an oscal-version field
					if metadata["oscal-version"] != nil {
						version, ok := metadata["oscal-version"].(string)
						// Return the version if it is a string and not empty
						if ok && version != "" {
							return version, nil
						}
					}
				}
			}
		}
	}

	// Failed to find the oscal-version in the assumed OscalModel structure, so do a deep search
	// TODO: Remove this deep search if it is not necessary
	if version == "" {
		return deepFindOscalVersion(data)
	}

	// Return an error if the field is not found
	return "", fmt.Errorf("required field: \"oscal-version\" not found")
}
