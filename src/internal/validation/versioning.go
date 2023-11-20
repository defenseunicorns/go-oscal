package validation

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
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

func findOscalVersion(data map[string]interface{}) (string, error) {
	// Check if the "oscal-version" field exists in the top-level map
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

	// Return an error if the field is not found
	return "", fmt.Errorf("required field: oscal-version not found")
}

// GetOscalModelVersion takes an interface{} or []byte and returns the metadata.oscal_version as a string
func GetOscalModelVersion[T interface{} | []byte](incomingModel T) (version string, err error) {
	// Check if interface{} and can be coerced to map[string]interface{}
	model, ok := reflect.ValueOf(incomingModel).Interface().(map[string]interface{})
	if !ok {
		// Check if []byte
		ymlBytes, ok := reflect.ValueOf(incomingModel).Interface().([]byte)
		// If not []byte, marshal to []byte
		if !ok {
			ymlBytes, err = yaml.Marshal(incomingModel)
			if err != nil {
				return "", err
			}
		}
		// Unmarshal to map[string]interface{}
		err = yaml.Unmarshal(ymlBytes, &model)
		if err != nil {
			return "", err
		}
	}

	// find the oscal-version field
	version, err = findOscalVersion(model)
	if err != nil {
		return "", err
	}
	// if version == "" {
	// 	return "", fmt.Errorf("required field: oscal-version not found")
	// }
	return version, nil
}

// CoerceToJSONForTypeSafety takes a yaml byte array and coerces it to a json interface{}
// This is necessary because the jsonschema library does not support yaml date types that are not declared as strings ie "2021-01-01" vs 2021-01-01
func CoerceToJSONForTypeSafety[T interface{} | []byte](version string, ymlData T) (model interface{}, err error) {
	ymlBytes, ok := reflect.ValueOf(ymlData).Interface().([]byte)
	if ok {
		err = yaml.Unmarshal(ymlBytes, &model)
		if err != nil {
			return nil, err
		}
	} else {
		model = ymlData
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

// IsValidSchemaVersion takes a version string and a []byte or interface{} and returns true if the yaml/json is valid for the specified oscal-version
func IsValidSchemaVersion[T interface{} | []byte](oscalVersion string, docBytes T) bool {
	component, err := CoerceToJSONForTypeSafety(oscalVersion, docBytes)
	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}

	compiler := jsonschema.NewCompiler()
	schemaPath := SCHEMA_PREFIX + strings.ReplaceAll(oscalVersion, ".", "-") + ".json"
	schemaBytes, err := schemas.ReadFile("schema/" + schemaPath)

	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}
	compiler.AddResource(schemaPath, strings.NewReader(string(schemaBytes)))
	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}
	sch, err := compiler.Compile(schemaPath)
	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}
	err = sch.Validate(component)
	if err != nil {
		log.Printf("%#v\n", err)
		return false
		// b, _ := json.MarshalIndent(err.(*jsonschema.ValidationError).DetailedOutput(), "", "  ")
		// fmt.Println(string(b))
		// return false
	}
	return true
}

// GetVersion returns formatted OSCAL version if valid version is passed, returns error if not.
func GetVersion(userVersion string) (string, error) {
	if userVersion == "" {
		return DEFAULT_OSCAL_VERSION, nil
	}
	builtVersion := formatUserVersion(userVersion)

	if !isVersionRegexp(builtVersion) {
		return builtVersion, fmt.Errorf("version %s is not a valid version", userVersion)
	}

	if !supportedVersion[builtVersion] {
		return builtVersion, fmt.Errorf("version %s is not supported", userVersion)
	}

	return builtVersion, nil
}

func isVersionRegexp(v string) bool {
	return versionRegexp.MatchString(v)
}

func formatUserVersion(v string) string {
	builtVersion := v
	if builtVersion[0] == 'v' {
		builtVersion = builtVersion[1:]
	}
	builtVersion = strings.ReplaceAll(builtVersion, "-", ".")
	return builtVersion
}
