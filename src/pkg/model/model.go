package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	oscalTypes_1_0_4 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	oscalTypes_1_0_6 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-6"
	oscalTypes_1_1_0 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-0"
	oscalTypes_1_1_1 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-1"
	oscalTypes_1_1_2 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-2"
	"gopkg.in/yaml.v3"
)

// InterfaceOrBytes is an interface{} or []byte for generic functions that can support either type
type InterfaceOrBytes interface {
	interface{} | []byte
}

// Finds the value of a key in a model in a map[string]interface{} given a slice of keys
// Returns nil if the key is not found or the model type is not supported
// Internal Recursive function so that the model can keep type safety
func FindValue(model map[string]interface{}, keys []string) interface{} {

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

// GetModelType returns the type of the model if the model is valid
// returns error if more than one model is found or no models are found (consistent with OSCAL spec)
func GetModelType(model map[string]interface{}) (modelType string, err error) {
	if len(model) != 1 {
		return "", fmt.Errorf("expected model to have 1 key, got %d", len(model))
	}
	for key := range model {
		modelType = key
	}
	return modelType, nil
}

// CoerceToJsonMap takes a yaml byte array and coerces it to a json interface{}
// This is necessary because the jsonschema library only accepts valid json data types that may not match yaml.
// Example: yaml allows for DateTimes to be time.Time, but json requires them to be strings
// This also allows for structs to be passed in, and they will be converted to map[string]interface{}
func CoerceToJsonMap(ymlData InterfaceOrBytes) (model map[string]interface{}, err error) {
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

// MarshalByExtension takes a model and marshals it to json or yaml based on the extension of the output file
func MarshalByExtension(model interface{}, outputFile string) (bytes []byte, err error) {
	split := strings.Split(outputFile, ".")
	ext := split[len(split)-1]

	switch ext {
	case "json":
		bytes, err = json.Marshal(model)
		if err != nil {
			return nil, err
		}
	case "yaml":
		bytes, err = yaml.Marshal(model)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}

func marshallYamlToStruct(data []byte) ([]byte, error) {

	var oscalModel = oscalTypes_1_1_2.OscalModels{}

	err := yaml.Unmarshal(data, &oscalModel)
	if err != nil {
		return nil, err
	}
	bytes, err := yaml.Marshal(oscalModel)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func Marshall104(data []byte, ext string) (bytes []byte, err error) {
	var oscalModel = oscalTypes_1_0_4.OscalModels{}

	switch ext {
	case "json":
		err := json.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	case "yaml":
		err := yaml.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = yaml.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}

func Marshall106(data []byte, ext string) (bytes []byte, err error) {
	var oscalModel = oscalTypes_1_0_6.OscalModels{}

	switch ext {
	case "json":
		err := json.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	case "yaml":
		err := yaml.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = yaml.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}

func Marshall110(data []byte, ext string) (bytes []byte, err error) {
	var oscalModel = oscalTypes_1_1_0.OscalModels{}

	switch ext {
	case "json":
		err := json.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	case "yaml":
		err := yaml.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = yaml.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}

func Marshall111(data []byte, ext string) (bytes []byte, err error) {
	var oscalModel = oscalTypes_1_1_1.OscalModels{}

	switch ext {
	case "json":
		err := json.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	case "yaml":
		err := yaml.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = yaml.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}

func Marshall112(data []byte, ext string) (bytes []byte, err error) {
	var oscalModel = oscalTypes_1_1_2.OscalModels{}

	switch ext {
	case "json":
		err := json.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = json.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	case "yaml":
		err := yaml.Unmarshal(data, &oscalModel)
		if err != nil {
			return nil, err
		}
		bytes, err = yaml.Marshal(oscalModel)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid file extension must be yaml or json: %s", ext)
	}

	return bytes, nil
}
