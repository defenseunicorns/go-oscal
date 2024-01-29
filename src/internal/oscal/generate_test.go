package oscal

import (
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/swaggest/jsonschema-go"
)

const (
	oscalComponentSchemaFilePath    string = "../../../schema/component/oscal_component_schema-1-1-1.json"
	oscalSSPSchemaFilePath          string = "../../../schema/ssp/oscal_ssp_schema-1-1-1.json"
	fieldsPresentFilePath           string = "../../../testdata/fields-present.json"
	fieldsMissingFilePath           string = "../../../testdata/fields-missing.json"
	expected104FilePath             string = "../../../testdata/generation/oscal-1-0-4/types.go"
	expected105FilePath             string = "../../types/oscal-1-0-5/types.go"
	expected106FilePath             string = "../../types/oscal-1-0-6/types.go"
	expected110FilePath             string = "../../types/oscal-1-1-0/types.go"
	expected111FilePath             string = "../../types/oscal-1-1-1/types.go"
	oscal104FilePath                string = "../../../schema/complete/oscal_complete_schema-1-0-5.json"
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
	writeOutput   = false
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	t.Run("It generates the types for a given complete schema", func(t *testing.T) {
		t.Parallel()
		bytes, err := Generate(schemaByteMap[oscal110FilePath], "oscalTypes", []string{"json", "yaml"})

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if writeOutput {
			utils.WriteOutput(bytes, "../../../test_out/oscal-1-1-0/types.go")
		}
	})

	t.Run("It generates the types for an individual schema", func(t *testing.T) {
		t.Parallel()
		bytes, err := Generate(schemaByteMap[oscalComponentSchemaFilePath], "oscalTypes", []string{"json", "yaml"})

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if writeOutput {
			utils.WriteOutput(bytes, "../../../test_out/component/types.go")
		}
	})
}

func TestGenerateDifferences(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	t.Skip("Skipping test. This test is not meant to be run in CI. It is meant to be run locally to generate the expected output files.")
	expectedBytes, err := os.ReadFile(expected111FilePath)
	if err != nil {
		t.Fatal(err)
	}
	expected := string(expectedBytes)

	schema, err := buildSchema(schemaByteMap[oscal111FilePath])
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	definitions, err := getDefinitionMap(schema)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	t.Run("It builds a struct map given a schema object", func(t *testing.T) {
		refQueue := NewRefQueue()
		for _, oneOf := range schema.OneOf {
			for _, pSOB := range oneOf.TypeObject.Properties {
				prop := pSOB.TypeObject
				if prop.Ref != nil && *prop.Ref != "#json-schema-directive" {
					refQueue.Add(*prop.Ref)
				}
			}

		}
		config := NewGeneratorConfig([]string{"json", "yaml"}, "oscalTypes")
		config.refQueue = refQueue
		config.definitions = definitions

		err := config.buildStructs()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		diffs := []string{}
		all := []string{}
		for _, v := range config.structMap {
			all = append(all, v)
			if !strings.Contains(expected, v) {
				diffs = append(diffs, v)
			}
		}

		if writeOutput {
			allTypes := "package allTypes\n\n" + strings.Join(all, "\n\n")
			os.WriteFile("../../../test_out/all/allTypes.go", []byte(allTypes), 0644)

			diffTypes := "package diffTypes\n\n" + strings.Join(diffs, "\n\n")
			os.WriteFile("../../../test_out/diff/diffTypes.go", []byte(diffTypes), 0644)
		}
	})
}

func TestBuildStructs(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	schemas := []string{oscal104FilePath, oscal105FilePath, oscal106FilePath, oscal110FilePath, oscal111FilePath}

	for _, path := range schemas {
		schema, err := buildSchema(schemaByteMap[path])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		config := NewGeneratorConfig([]string{"json", "yaml"}, "oscalTypes")

		err = config.initBuild(&schema)

		err = config.buildStructs()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		var structMap = map[string]string{}
		duplicates := []string{}
		for k, value := range config.structMap {
			firstLine := strings.ReplaceAll(strings.Split(value, "\n")[0], " ", "")
			mapValue := structMap[firstLine]
			if mapValue == "" {
				structMap[firstLine] = k
				continue
			}

			if mapValue != k {
				duplicates = append(duplicates, "Duplicate struct name found: ", k, " and ", mapValue)
			}
		}

		if len(duplicates) > 0 {
			t.Errorf("expected no duplicates, got %v", duplicates)
		}
	}

}

func TestBuildStructString(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)
	schema, err := buildSchema(schemaByteMap[oscal111FilePath])
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	definitions, err := getDefinitionMap(schema)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	t.Run("It builds a struct string given a schema object", func(t *testing.T) {
		t.Parallel()
		config := GeneratorConfig{
			tags:        []string{"json", "yaml"},
			definitions: definitions,
			pkgName:     "oscalTypes",
			refQueue:    NewRefQueue(),
			nameMap:     map[string]string{},
		}
		result, err := config.buildStructString(*schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected, err := os.ReadFile("../../../testdata/generation/structs/catalog.go")
		if err != nil {
			t.Fatal(err)
		}

		os.WriteFile("../../../testdata/generation/catalog.go", []byte(result), 0644)
		if result != string(expected) {
			t.Errorf("expected %s, got %s", expected, result)
		}
	})

	t.Run("It handles an object with no properties", func(t *testing.T) {
		t.Parallel()
		config := GeneratorConfig{
			tags:        []string{"json", "yaml"},
			definitions: definitions,
			pkgName:     "oscalTypes",
			refQueue:    NewRefQueue(),
			nameMap:     map[string]string{},
		}
		result, err := config.buildStructString(*schema.Definitions["oscal-complete-oscal-control-common:include-all"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		expected, err := os.ReadFile("../../../testdata/generation/structs/include-all.go")
		if err != nil {
			t.Fatal(err)
		}

		if result != string(expected) {
			t.Errorf("expected %s, got %s", expected, result)
		}
	})
}

func TestDefinitionMap(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	schema, err := buildSchema(schemaByteMap[oscal111FilePath])
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	definitionMap, err := getDefinitionMap(schema)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	t.Run("It maps $id's to their respective definitions", func(t *testing.T) {
		t.Parallel()
		result, ok := definitionMap["#assembly_oscal-catalog_catalog"]
		if !ok {
			t.Errorf("expected a result, got nil")
		}
		if *result.ID != "#assembly_oscal-catalog_catalog" {
			t.Errorf("expected %s, got %s", "#assembly_oscal-catalog_catalog", *result.ID)
		}
	})

	t.Run("it maps definitions to a $ref pattern if no $id is present", func(t *testing.T) {
		t.Parallel()
		result, ok := definitionMap["#/definitions/TokenDatatype"]
		if !ok {
			t.Errorf("expected a result, got nil")
		}

		if *result.Type.SimpleTypes != "string" {
			t.Errorf("expected %s, got %v", "string", result.Type.SimpleTypes)
		}
	})
}

func TestGetTypeSuffix(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	schema, err := buildSchema(schemaByteMap[oscal111FilePath])
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	definitions, err := getDefinitionMap(schema)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	config := GeneratorConfig{
		tags:        []string{"json", "yaml"},
		definitions: definitions,
		pkgName:     "oscalTypes",
		refQueue:    NewRefQueue(),
		nameMap:     map[string]string{},
	}
	t.Run("It returns the object typename given a schema object", func(t *testing.T) {
		result, err := config.findSubType(*schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "Catalog" {
			t.Errorf("expected %s, got %s", "Catalog", result)
		}
	})

	t.Run("It returns an array complex type name given a schema array", func(t *testing.T) {
		result, err := config.findSubType(*schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject.Properties["groups"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "Group" {
			t.Errorf("expected %s, got %s", "Group", result)
		}
	})

	t.Run("It uses the the title to contain the type if no $ref or $id is present and adds to queue", func(t *testing.T) {
		ref := "#/definitions/ConstraintTest"
		found := false
		result, err := config.findSubType(*schema.Definitions["oscal-complete-oscal-control-common:parameter-constraint"].TypeObject.Properties["tests"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "ConstraintTest" {
			t.Errorf("expected %s, got %s", "ConstraintTest", result)
		}
		for _, q := range config.refQueue.refs {
			if q == ref {
				found = true
			}
		}
		if !found {
			t.Errorf("expected %s to be in queue", ref)
		}
		if config.definitions[ref].Type == nil {
			t.Errorf("expected %s to be in definitions", ref)
		}
	})
}

func TestBuildTypeString(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	schema, err := buildSchema(schemaByteMap[oscal111FilePath])
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	definitions, err := getDefinitionMap(schema)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	config := GeneratorConfig{
		tags:        []string{"json", "yaml"},
		definitions: definitions,
		pkgName:     "oscalTypes",
		refQueue:    NewRefQueue(),
		nameMap:     map[string]string{},
	}

	t.Run("It returns the associated primitive gotype when given a schema of a primitive type", func(t *testing.T) {
		result, err := config.buildTypeString(*schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject.Properties["uuid"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "string" {
			t.Errorf("expected %s, got %s", "string", result)
		}

	})

	t.Run("It returns the first ref type when given a schema that implements allOf", func(t *testing.T) {
		result, err := config.buildTypeString(*schema.Definitions["oscal-complete-oscal-control-common:parameter-selection"].TypeObject.Properties["how-many"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "string" {
			t.Errorf("expected %s, got %s", "string", result)
		}

	})

	t.Run("It returns the first ref type when given a schema that implements anyOf", func(t *testing.T) {
		result, err := config.buildTypeString(*schema.Definitions["oscal-complete-oscal-metadata:link"].TypeObject.Properties["rel"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "string" {
			t.Errorf("expected %s, got %s", "string", result)
		}
	})

	t.Run("It returns an array type in the format []type when given a schema that implements array", func(t *testing.T) {
		result, err := config.buildTypeString(*schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject.Properties["groups"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "[]Group" {
			t.Errorf("expected %s, got %s", "[]Group", result)
		}
	})

	t.Run("It returns the struct name of the object when given a schema that implements object", func(t *testing.T) {
		result, err := config.buildTypeString(*schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != "Catalog" {
			t.Errorf("expected %s, got %s", "Catalog", result)
		}
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
