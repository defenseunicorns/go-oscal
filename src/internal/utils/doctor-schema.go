package utils

import (
	"fmt"
	"strings"
)

// DoctorJSON recursively replaces all $ref keys with the referenced object
// when an object has both $id and $ref keys defined to support JSON Schema draft 7
func DoctorJSON(obj map[string]interface{}, parentObj map[string]interface{}, parentKey string) (err error) {
	for key, value := range obj {
		switch typedValue := value.(type) {
		case map[string]interface{}:
			if _, idExists := typedValue["$id"]; idExists {
				if ref, refExists := typedValue["$ref"]; refExists {
					refParts := strings.Split(ref.(string), "/")[1:]
					// Ensure the $ref is to definitions
					// TODO: expand to support other $ref types if needed.
					if refParts[0] == "definitions" {
						delete(typedValue, "$ref")
						replaceValue := parentObj
						for _, refPart := range refParts {
							if replaceValue[refPart] == nil {
								return fmt.Errorf("$ref %s not found", ref)
							}
							replaceValue = replaceValue[refPart].(map[string]interface{})
						}
						obj[key] = MergeMaps(replaceValue, typedValue)
					} else {
						return fmt.Errorf("$ref %s not supported. Only $ref to definitions supported", ref)
					}
				}
			}
			err = DoctorJSON(typedValue, obj, key)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// MergeMaps merges two maps
func MergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}
