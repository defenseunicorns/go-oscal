package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
)

func TestDoctorJson(t *testing.T) {
	jsonStr := `{
        "definitions": {
            "example": {
                "$id": "http://example.com/schema",
                "$ref": "#/definitions/exampleDefinition",
				"description": "Example Description"
            },
			"exampleDefinition": {
				"type": "object",
				"description": "Should not override the description of the referring object",
				"properties": {
					"prop1": {
						"type": "string"
					},
					"prop2": {
						"type": "number"
					}
				}
			}
		}
	}`

	t.Run("Should remove the $ref field from the object with both $ref and $id", func(t *testing.T) {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &obj)
		if err != nil {
			t.Fatal(err)
		}
		utils.DoctorJSON(obj, nil, "")

		if obj["definitions"].(map[string]interface{})["example"].(map[string]interface{})["$ref"] != nil {
			t.Error("Expected $ref to be resolved")
		}
	})

	t.Run("Should not override the description of the referring object", func(t *testing.T) {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &obj)
		if err != nil {
			t.Fatal(err)
		}
		utils.DoctorJSON(obj, nil, "")

		if obj["definitions"].(map[string]interface{})["example"].(map[string]interface{})["description"] != "Example Description" {
			t.Error("Expected description to be preserved")
		}
	})

	t.Run("should merge the properties of the referenced object", func(t *testing.T) {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &obj)
		if err != nil {
			t.Fatal(err)
		}
		utils.DoctorJSON(obj, nil, "")

		if obj["definitions"].(map[string]interface{})["example"].(map[string]interface{})["properties"].(map[string]interface{})["prop1"].(map[string]interface{})["type"] != "string" {
			t.Error("Expected properties to be merged")
		}
	})

	t.Run("should throw an error if the $ref is not found", func(t *testing.T) {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &obj)
		if err != nil {
			t.Fatal(err)
		}
		obj["definitions"].(map[string]interface{})["example"].(map[string]interface{})["$ref"] = "#/definitions/notFound"
		err = utils.DoctorJSON(obj, nil, "")
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("should throw an error if the $ref is not to definitions", func(t *testing.T) {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &obj)
		if err != nil {
			t.Fatal(err)
		}
		obj["definitions"].(map[string]interface{})["example"].(map[string]interface{})["$ref"] = "#/notDefinitions/notFound"
		err = utils.DoctorJSON(obj, nil, "")
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
