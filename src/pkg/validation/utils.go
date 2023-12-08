package validation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

// Extension of the jsonschema.BasicError struct to include the failed value
// if the failed value is a map or slice, it will be omitted
type ValidatorError struct {
	jsonschema.BasicError
	FailedValue interface{} `json:"failedValue,omitempty"`
}

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

// Creates a []ValidatorError from a jsonschema.Basic
// The jsonschema.Basic contains the errors from the validation
func extractErrors(originalObject map[string]interface{}, validationError jsonschema.Basic) (validationErrors []ValidatorError) {
	for _, error := range validationError.Errors {
		if error.InstanceLocation == "" || error.Error == "" || strings.HasPrefix(error.Error, "doesn't validate with") {
			continue
		}
		if len(validationErrors) > 0 && validationErrors[len(validationErrors)-1].InstanceLocation == error.InstanceLocation {
			validationErrors[len(validationErrors)-1].Error += ", " + error.Error
		} else {
			failedValue := findValue(originalObject, strings.Split(error.InstanceLocation, "/")[1:])
			_, mapOk := failedValue.(map[string]interface{})
			_, sliceOk := failedValue.([]interface{})
			if mapOk || sliceOk {
				failedValue = nil
			}
			validationErrors = append(validationErrors, ValidatorError{BasicError: error, FailedValue: failedValue})
		}
	}
	return validationErrors

}

// Finds the value of a key in a model in a map[string]interface{} given a slice of keys
// Returns nil if the key is not found or the model type is not supported
// Internal Recursive function so that the model can keep type safety
func findValue(model map[string]interface{}, keys []string) interface{} {

	// Define recursive function
	var find func(model interface{}, keys []string) interface{}
	find = func(root interface{}, keys []string) interface{} {
		// If there are no more keys to find, return the parent
		if len(keys) == 0 {
			return root
		}

		childKey := keys[0]
		descendentKeys := keys[1:]

		// If the model is a map find the next value
		if rootAsMap, ok := root.(map[string]interface{}); ok {
			if child, found := rootAsMap[childKey]; found {
				return find(child, descendentKeys)
			}
		}

		// If the model is a rootAsSlice find the next value
		if rootAsSlice, ok := root.([]interface{}); ok {
			// Convert the key to an int (should always be an int/index)
			index, err := strconv.Atoi(childKey)
			if err != nil {
				return nil
			}
			child := rootAsSlice[index]
			return find(child, descendentKeys)
		}

		// If the key is not found or model type is not supported, return nil
		return nil
	}
	return find(model, keys)
}

// getModelType returns the type of the model if the model is valid
// returns error if more than one model is found or no models are found (consistent with OSCAL spec)
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
	var ok bool
	for _, value := range model {
		submodel, ok = value.(map[string]interface{})
		if !ok {
			continue
		}
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
