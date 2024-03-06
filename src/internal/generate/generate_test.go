package generate

import (
	"go/format"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
	"github.com/swaggest/jsonschema-go"
)

const (
	schemaPath            string = "../schemas/"
	fieldsPresentFilePath string = "../../../testdata/fields-present.json"
	fieldsMissingFilePath string = "../../../testdata/fields-missing.json"
)

var (
	schemaMutex            = sync.Mutex{}
	schemaByteMap          = map[string][]byte{}
	writeOutput            = true
	deterministicTestCount = 10
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	t.Run("CompleteSchema", func(t *testing.T) {

		for path, schemaBytes := range schemaByteMap {
			if !strings.Contains(path, "_complete_schema") {
				continue
			}

			pkgPath := strings.ReplaceAll(path, "_complete_schema", "")
			pkgPath = strings.ReplaceAll(pkgPath, ".json", "")
			pkgName := strings.ReplaceAll(pkgPath, "-", "_")
			pkgName = strings.ReplaceAll(pkgName, "oscal", "oscalTypes")

			bytes, err := Generate(schemaBytes, pkgName, []string{"json", "yaml"})
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if writeOutput {
				utils.WriteOutput(bytes, "../../types/"+pkgPath+"/types.go")
			}
		}
	})

}

func TestGenerateDeterministic(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	var previousBytes []byte

	for _, schemaBytes := range schemaByteMap {
		previousBytes = nil
		for i := 0; i < deterministicTestCount; i++ {
			bytes, err := Generate(schemaBytes, "oscalTypes", []string{"json", "yaml"})
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if previousBytes == nil {
				previousBytes = bytes
				continue
			}
			if string(previousBytes) != string(bytes) {
				t.Error("expected deterministic output")
			}
		}
	}
}

func TestBuildStructs(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	for _, schemaBytes := range schemaByteMap {
		schema, err := buildSchema(schemaBytes)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		config := NewGeneratorConfig([]string{"json", "yaml"}, "oscalTypes")

		err = config.initBuild(&schema)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

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
	schema, err := buildSchema(schemaByteMap["oscal_complete_schema-1-1-1.json"])
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

		bytes, err := format.Source([]byte(result))
		if err != nil {
			t.Fatal(err)
		}
		actual := string(bytes)

		bytes, err = os.ReadFile("../../../testdata/generation/structs/catalog.go")
		if err != nil {
			t.Fatal(err)
		}

		expected := string(bytes)

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, result)
		}
	})

	t.Run("It handles an object with no properties by aliasing it to map[string]interface{} for json representation", func(t *testing.T) {
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

		bytes, err := format.Source([]byte(result))
		if err != nil {
			t.Fatal(err)
		}
		actual := string(bytes)
		expected, err := os.ReadFile("../../../testdata/generation/structs/include-all.go")
		if err != nil {
			t.Fatal(err)
		}

		if actual != string(expected) {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})

	t.Run("It adds an alias if one exists", func(t *testing.T) {
		t.Parallel()
		config := NewGeneratorConfig([]string{"json", "yaml"}, "oscalTypes")
		config.definitions = definitions
		config.initBuild(&schema)

		result, err := config.buildStructString(config.definitions[config.refQueue.Pop()])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if !strings.Contains(result, "type OscalModels = OscalCompleteSchema") {
			t.Errorf("expected %s to contain %s", result, "type OscalModels = OscalCompleteSchema")
		}

	})
}

func TestGetTypeSuffix(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	schema, err := buildSchema(schemaByteMap["oscal_complete_schema-1-1-1.json"])
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

	schema, err := buildSchema(schemaByteMap["oscal_complete_schema-1-1-1.json"])
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
		dir, err := os.ReadDir(schemaPath)
		if err != nil {
			panic(err)
		}
		for _, file := range dir {
			path := file.Name()
			if !strings.Contains(path, ".json") {
				continue
			}
			schemaBytes, err := os.ReadFile(schemaPath + path)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			schemaByteMap[path] = schemaBytes
		}
	}
}
