package oscal

import (
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/swaggest/jsonschema-go"
)

const (
	oscalComponentSchemaFilePath    string = "../../../schema/component/oscal_component_schema-1-1-1.json"
	oscalSSPSchemaFilePath          string = "../../../schema/ssp/oscal_ssp_schema-1-1-1.json"
	fieldsPresentFilePath           string = "../../../testdata/fields-present.json"
	fieldsMissingFilePath           string = "../../../testdata/fields-missing.json"
	expected104FilePath             string = "../../types/oscal-1-0-4/types.go"
	expected105FilePath             string = "../../types/oscal-1-0-5/types.go"
	expected106FilePath             string = "../../types/oscal-1-0-6/types.go"
	expected110FilePath             string = "../../types/oscal-1-1-0/types.go"
	expected111FilePath             string = "../../types/oscal-1-1-1/types.go"
	oscal104FilePath                string = "../../../schema/complete/oscal_complete_schema-1-0-4.json"
	oscal105FilePath                string = "../../../schema/complete/oscal_complete_schema-1-0-5.json"
	oscal106FilePath                string = "../../../schema/complete/oscal_complete_schema-1-0-6.json"
	oscal110FilePath                string = "../../../schema/complete/oscal_complete_schema-1-1-0.json"
	oscal111FilePath                string = "../../../schema/complete/oscal_complete_schema-1-1-1.json"
	componentExpectedPropertiesFile string = "../../../testdata/generation/component/expected-properties.txt"
	sspExpectedPropertiesFile       string = "../../../testdata/generation/ssp/expected-properties.txt"
	componentExpectedStructDataFile string = "../../../testdata/generation/component/expected-struct-data.txt"
	sspExpectedStructDataFile       string = "../../../testdata/generation/ssp/expected-struct-data.txt"
	expectedComponentOutputFile     string = "../../../testdata/generation/component/expected-oscal-model-struct.txt"
	sspExpectedOutputFile           string = "../../../testdata/generation/ssp/expected-oscal-model-struct.txt"
)

var (
	schemaPaths   = []string{oscal104FilePath, oscal105FilePath, oscal106FilePath, oscal110FilePath, oscal111FilePath, oscalComponentSchemaFilePath, oscalSSPSchemaFilePath}
	schemaMutex   = sync.Mutex{}
	schemaByteMap = map[string][]byte{}
)

// Assumes that current types are valid.
func TestGenerateRegression(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)
	t.Run("Test104", func(t *testing.T) {
		t.Parallel()
		expected, err := os.ReadFile(expected104FilePath)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := Generate(schemaByteMap[oscal104FilePath], "oscalTypes", []string{"json", "yaml"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Fatal("expected output to match")
		}
	})

	t.Run("Test105", func(t *testing.T) {
		t.Parallel()
		expected, err := os.ReadFile(expected105FilePath)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := Generate(schemaByteMap[oscal105FilePath], "oscalTypes", []string{"json", "yaml"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Fatal("expected output to match")
		}
	})

	t.Run("Test106", func(t *testing.T) {
		t.Parallel()
		expected, err := os.ReadFile(expected106FilePath)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := Generate(schemaByteMap[oscal106FilePath], "oscalTypes", []string{"json", "yaml"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Fatal("expected output to match")
		}
	})

	t.Run("Test110", func(t *testing.T) {
		t.Parallel()
		expected, err := os.ReadFile(expected110FilePath)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := Generate(schemaByteMap[oscal110FilePath], "oscalTypes", []string{"json", "yaml"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Fatal("expected output to match")
		}
	})

	t.Run("Test111", func(t *testing.T) {
		t.Parallel()
		expected, err := os.ReadFile(expected111FilePath)
		if err != nil {
			t.Fatal(err)
		}

		actual, err := Generate(schemaByteMap[oscal111FilePath], "oscalTypes", []string{"json", "yaml"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Fatal("expected output to match")
		}
	})

}

func TestGenerate(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)
	t.Run("TestDefinitionMap", func(t *testing.T) {
		t.Parallel()
		schema, err := buildSchema(schemaByteMap[oscal111FilePath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		definitionMap, err := generateUniqueIdMap(schema)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		t.Run("It maps $id's to their respective definitions", func(t *testing.T) {
			t.Parallel()
			result, ok := definitionMap["#assembly_oscal-catalog_catalog"]
			if !ok {
				t.Errorf("expected a result, got nil")
			}
			if *result.TypeObject.ID != "#assembly_oscal-catalog_catalog" {
				t.Errorf("expected %s, got %s", "#assembly_oscal-catalog_catalog", *result.TypeObject.ID)
			}
		})

		t.Run("it maps definitions to a $ref pattern if no $id is present", func(t *testing.T) {
			t.Parallel()
			result, ok := definitionMap["#/definitions/TokenDatatype"]
			if !ok {
				t.Errorf("expected a result, got nil")
			}

			if *result.TypeObject.Type.SimpleTypes != "string" {
				t.Errorf("expected %s, got %v", "string", result.TypeObject.Type.SimpleTypes)
			}
		})
	})
}

func buildSchema(schemaBytes []byte) (jsonschema.Schema, error) {
	schema := jsonschema.Schema{}
	err := schema.UnmarshalJSON(schemaBytes)
	if err != nil {
		return jsonschema.Schema{}, err
	}
	return schema, nil
}

func getSchemaByteMap(t *testing.T) {
	schemaMutex.Lock()
	defer schemaMutex.Unlock()
	if len(schemaByteMap) == 0 {
		for _, path := range schemaPaths {
			bytes, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			schemaByteMap[path] = bytes
		}
	}
}
