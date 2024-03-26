package generate_test

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	oscalTypes_1_0_4 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	oscalTypes_1_1_1 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-1"
	"gopkg.in/yaml.v3"
)

var (
	keysToNotRecurse = map[string]bool{
		"flat":        true,
		"include-all": true,
	}
	rev4YamlPath  = "../../../testdata/generation/e2e/rev4/yaml/"
	rev4JsonPath  = "../../../testdata/generation/e2e/rev4/json/"
	rev5YamlPath  = "../../../testdata/generation/e2e/rev5/yaml/"
	rev5JsonPath  = "../../../testdata/generation/e2e/rev5/json/"
	oscal104Types = "../../types/oscal-1-0-4/types.go"
	oscal111Types = "../../types/oscal-1-1-1/types.go"
)

// TestFedrampBaselineYamlFieldsInTypes tests that the fields in the yaml files are present in the types.go files
// for the fedramp baseline schemas. While this test is not exhaustive,
// it is a good way to capture any missing fields in the types.go files should they arise.
func TestFedrampBaselineYamlFieldsInTypes(t *testing.T) {
	t.Parallel()

	t.Run("Rev4", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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

// TestFieldStability tests that the fields in the yaml files are stable when marshaled and unmarshaled
// for the fedramp baseline schemas.
// This is important to ensure that the fields are not lost when marshaled and unmarshaled.
func TestFieldStability(t *testing.T) {
	t.Parallel()

	t.Run("Rev4", func(t *testing.T) {
		t.Parallel()
		dir, err := os.ReadDir(rev4YamlPath)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range dir {

			bytes, err := os.ReadFile(rev4YamlPath + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			oscalDoc := oscalTypes_1_0_4.OscalCompleteSchema{}
			err = yaml.Unmarshal(bytes, &oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			// Marshal the document back to yaml
			marshaled, err := yaml.Marshal(oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			actual := map[string]interface{}{}
			expected := map[string]interface{}{}
			err = yaml.Unmarshal(marshaled, &actual)
			if err != nil {
				t.Fatal(err)
			}
			err = yaml.Unmarshal(bytes, &expected)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Error("expected marshaled yaml to be equal to the original yaml")
			}

		}
	})

	t.Run("Rev5", func(t *testing.T) {
		t.Parallel()
		dir, err := os.ReadDir(rev5YamlPath)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range dir {

			bytes, err := os.ReadFile(rev5YamlPath + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			oscalDoc := oscalTypes_1_1_1.OscalCompleteSchema{}
			err = yaml.Unmarshal(bytes, &oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			// Marshal the document back to yaml
			marshaled, err := yaml.Marshal(oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			actual := map[string]interface{}{}
			expected := map[string]interface{}{}
			err = yaml.Unmarshal(marshaled, &actual)
			if err != nil {
				t.Fatal(err)
			}
			err = yaml.Unmarshal(bytes, &expected)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Error("expected marshaled yaml to be equal to the original yaml")
			}

		}
	})
}

func TestNoExtraJsonFields(t *testing.T) {
	t.Parallel()

	t.Run("Rev4", func(t *testing.T) {
		t.Parallel()
		dir, err := os.ReadDir(rev4JsonPath)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range dir {
			bytes, err := os.ReadFile(rev4JsonPath + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			oscalDoc := oscalTypes_1_0_4.OscalCompleteSchema{}
			err = json.Unmarshal(bytes, &oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			// Marshal the document back to yaml
			marshaled, err := json.Marshal(oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			actual := map[string]interface{}{}
			expected := map[string]interface{}{}
			err = json.Unmarshal(marshaled, &actual)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(bytes, &expected)
			if err != nil {
				t.Fatal(err)
			}

			if !validateMapKeys(actual, expected) {
				t.Error("expected marshaled json to be equal to the original json")
			}
		}
	})
	t.Run("Rev5", func(t *testing.T) {
		t.Parallel()
		dir, err := os.ReadDir(rev5JsonPath)
		if err != nil {
			t.Fatal(err)
		}

		for _, file := range dir {
			bytes, err := os.ReadFile(rev5JsonPath + file.Name())
			if err != nil {
				t.Fatal(err)
			}
			oscalDoc := oscalTypes_1_1_1.OscalCompleteSchema{}
			err = json.Unmarshal(bytes, &oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			// Marshal the document back to yaml
			marshaled, err := json.Marshal(oscalDoc)
			if err != nil {
				t.Fatal(err)
			}

			actual := map[string]interface{}{}
			expected := map[string]interface{}{}
			err = json.Unmarshal(marshaled, &actual)
			if err != nil {
				t.Fatal(err)
			}
			err = json.Unmarshal(bytes, &expected)
			if err != nil {
				t.Fatal(err)
			}

			if !validateMapKeys(actual, expected) {
				t.Error("expected marshaled json to be equal to the original json")
			}
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
			if _, ok := keysToNotRecurse[key]; !ok {
				ValidateKeys(rootAsMap, typeString, t)
			}
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

func collectKeys(mapData map[string]interface{}, keys map[string]struct{}) {
	for key, value := range mapData {
		keys[key] = struct{}{}
		if nestedMap, ok := value.(map[string]interface{}); ok {
			collectKeys(nestedMap, keys)
		}
	}
}

func validateMapKeys(map1, map2 map[string]interface{}) bool {
	keys1 := make(map[string]struct{})
	keys2 := make(map[string]struct{})

	collectKeys(map1, keys1)
	collectKeys(map2, keys2)

	return reflect.DeepEqual(keys1, keys2)
}
