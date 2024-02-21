package generate_test

import (
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

var (
	rev4YamlPath  = "../../../testdata/generation/e2e/rev4/yaml/"
	rev5YamlPath  = "../../../testdata/generation/e2e/rev5/yaml/"
	oscal104Types = "../../types/oscal-1-0-4/types.go"
	oscal111Types = "../../types/oscal-1-1-1/types.go"
)

func TestFedrampBaselineYamlFieldsInTypes(t *testing.T) {

	t.Run("Rev4", func(t *testing.T) {
		typeBytes, err := os.ReadFile(oscal104Types)
		if err != nil {
			t.Fatal(err)
		}

		dir, err := os.ReadDir(rev4YamlPath)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range dir {

			bytes, err := os.ReadFile(rev4YamlPath + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			oscalDoc := map[string]interface{}{}
			err = yaml.Unmarshal(bytes, &oscalDoc)
			if err != nil {
				t.Fatal(err)
			}
			ValidateKeys(oscalDoc, string(typeBytes), t)
		}
	})

	t.Run("Rev5", func(t *testing.T) {
		typeBytes, err := os.ReadFile(oscal111Types)
		if err != nil {
			t.Fatal(err)
		}

		dir, err := os.ReadDir(rev5YamlPath)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range dir {

			bytes, err := os.ReadFile(rev5YamlPath + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			oscalDoc := map[string]interface{}{}
			err = yaml.Unmarshal(bytes, &oscalDoc)
			if err != nil {
				t.Fatal(err)
			}
			ValidateKeys(oscalDoc, string(typeBytes), t)
		}

	})
}

func ValidateKeys(model map[string]interface{}, typeString string, t *testing.T) {
	for key, value := range model {
		if !strings.Contains(typeString, "yaml:\""+key) {
			t.Errorf("expected key %s not found", key)
		}

		// If the model is a map find the next value
		if rootAsMap, ok := value.(map[string]interface{}); ok {
			ValidateKeys(rootAsMap, typeString, t)
		}

		// If the model is a rootAsSlice find the next value
		if rootAsSlice, ok := value.([]interface{}); ok {
			if len(rootAsSlice) > 0 {
				if rootAsMap, ok := rootAsSlice[0].(map[string]interface{}); ok {
					ValidateKeys(rootAsMap, typeString, t)
				}
			}
		}
	}
}
