package validation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
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

func getModelType(model map[string]interface{}) (modelType string, err error) {
	if len(model) != 1 {
		return "", fmt.Errorf("expected model to have 1 key, got %d", len(model))
	}
	for key := range model {
		modelType = key
	}
	return modelType, nil
}

// getOscalVersionFromMap returns formatted OSCAL version if valid version is passed, returns error if not.
func getOscalVersionFromMap(model map[string]interface{}) (version string, err error) {
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

	if err = isValidOscalVersion(version); err != nil {
		return "", err
	}

	return version, nil
}

// isValidOscalVersion returns true if the version is supported, false if not.
func isValidOscalVersion(version string) error {
	if !versionRegexp.MatchString(version) {
		return fmt.Errorf("version %s is not a valid version", version)
	}

	if !supportedVersion[version] {
		return fmt.Errorf("version %s is not supported", version)
	}
	return nil
}

// formatOscalVersion takes a version string and returns a formatted version string
func formatOscalVersion(version string) (formattedVersion string) {
	formattedVersion = strings.ToLower(version)
	formattedVersion = strings.ReplaceAll(formattedVersion, "-", ".")
	formattedVersion = strings.Trim(formattedVersion, "v")
	return formattedVersion
}

// coerceToJsonMap takes a yaml byte array and coerces it to a json interface{}
// This is necessary because the jsonschema library only accepts valid json data types that may not match yaml.
// Example: yaml allows for DateTimes to be time.Time, but json requires them to be strings
// This also allows for structs to be passed in, and they will be converted to map[string]interface{}
func coerceToJsonMap(ymlData InterfaceOrBytes) (model map[string]interface{}, err error) {
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

// convertInterfaceOrBytesToMap takes an InterfaceOrByte and returns a map[string]interface{}
func convertInterfaceOrBytesToMap(incomingModel InterfaceOrBytes) (model map[string]interface{}, err error) {
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
