package generate

import (
	"testing"

	"github.com/swaggest/jsonschema-go"
)

func TestBuildTagString(t *testing.T) {
	t.Parallel()

	t.Run("It returns a tag string given a list of tags and a field name", func(t *testing.T) {
		t.Parallel()
		expected := "`json:\"test,omitempty\" yaml:\"test,omitempty\"`"
		actual := buildTagString([]string{"json", "yaml"}, "test", false)
		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})

	t.Run("It leaves out omitempty if the field is required", func(t *testing.T) {
		t.Parallel()
		expected := "`json:\"test\" yaml:\"test\"`"
		result := buildTagString([]string{"json", "yaml"}, "test", true)
		if result != expected {
			t.Errorf("expected %s, got %s", expected, result)
		}
	})
}

func TestGetRef(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	t.Run("It returns the ref if it exists", func(t *testing.T) {
		t.Parallel()
		schema, err := buildSchema(schemaByteMap[oscal111FilePath])
		if err != nil {
			t.Fatalf("error building schema: %s", err)
		}
		schemaWithRef := schema.OneOf[0].TypeObject.Properties["catalog"].TypeObject

		expected := "#assembly_oscal-catalog_catalog"

		actual, err := getRef(*schemaWithRef)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})

	t.Run("It returns the id if it exists", func(t *testing.T) {
		t.Parallel()
		schema, err := buildSchema(schemaByteMap[oscal111FilePath])
		if err != nil {
			t.Fatalf("error building schema: %s", err)
		}
		schemaWitId := schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject
		expected := "#assembly_oscal-catalog_catalog"

		actual, err := getRef(*schemaWitId)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})

	t.Run("It builds a ref with the title if no id and no ref exist", func(t *testing.T) {
		t.Parallel()
		schema, err := buildSchema(schemaByteMap[oscal111FilePath])
		if err != nil {
			t.Fatalf("error building schema: %s", err)
		}
		schemaWithNoIdOrRef := schema.Definitions["oscal-complete-oscal-control-common:parameter-constraint"].TypeObject.Properties["tests"].TypeObject.Items.SchemaOrBool.TypeObject
		expected := "#/definitions/ConstraintTest"

		actual, err := getRef(*schemaWithNoIdOrRef)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})

	t.Run("It returns an error if no ref, id, or title exist", func(t *testing.T) {
		t.Parallel()
		schema, err := buildSchema(schemaByteMap[oscal111FilePath])
		if err != nil {
			t.Fatalf("error building schema: %s", err)
		}
		schemaWithNoIdOrRef := schema.Definitions["oscal-complete-oscal-control-common:parameter-constraint"].TypeObject.Properties["tests"].TypeObject.Items.SchemaOrBool.TypeObject
		schemaWithNoIdOrRef.Title = nil
		schemaWithNoIdOrRef.ID = nil

		_, err = getRef(*schemaWithNoIdOrRef)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}

func TestGetJsonType(t *testing.T) {
	t.Parallel()
	getSchemaByteMap(t)

	schema, err := buildSchema(schemaByteMap[oscal111FilePath])
	if err != nil {
		t.Fatalf("error building schema: %s", err)
	}

	t.Run("It returns the json type if it exists", func(t *testing.T) {
		t.Parallel()
		schemaWithSimpleType := schema.Definitions["oscal-complete-oscal-catalog:catalog"].TypeObject

		expected := "object"

		actual := getJsonType(*schemaWithSimpleType)

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})

	t.Run("It returns an empty string if there is no json type", func(t *testing.T) {
		t.Parallel()
		schemaWithNoSimpleType := schema.OneOf[0].TypeObject // This schema has no type
		expected := ""

		actual := getJsonType(*schemaWithNoSimpleType)

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}

func TestIsPrimitiveJsonType(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out bool
	}

	testCases := []TestCase{
		{in: "string", out: true},
		{in: "boolean", out: true},
		{in: "number", out: true},
		{in: "integer", out: true},
		{in: "object", out: false},
		{in: "array", out: false},
		{in: "date-time", out: true},
	}

	for _, testCase := range testCases {
		actual := isPrimitiveJsonType(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Errorf("error isPrimitiveJsonType(): expected: %t | got: %t", expected, actual)
		}
	}
}

func TestGetGoType(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "string", out: "string"},
		{in: "boolean", out: "bool"},
		{in: "number", out: "float64"},
		{in: "integer", out: "int"},
		{in: "array", out: "[]"},
		{in: "object", out: ""},
		{in: "date-time", out: "time.Time"},
	}

	for _, testCase := range testCases {
		actual := getGoType(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Errorf("error getGoType(): expected: %s | got: %s", expected, actual)
		}
	}
}

func TestGetNameFromRef(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "#/definitions/ConstraintTest", out: "ConstraintTest"},
		{in: "#assembly_oscal-catalog_catalog", out: "Catalog"},
		{in: "#assembly_oscal-control-common_parameter-guideline", out: "ParameterGuideline"},
		{in: "#http://csrc.nist.gov/ns/oscal/1.0/1.1.1/oscal-complete-schema.json", out: "OscalCompleteSchema"},
	}

	for _, testCase := range testCases {
		actual := getNameFromRef(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Fatalf("error getNameFromRef(): expected: %s | got: %s", expected, actual)
		}
	}
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

// TestFmtFieldName tests that we handle formatting a json string to a go struct correctly.
func TestFmtFieldName(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "foo_id", out: "FooID"},
		{in: "fooId", out: "FooID"},
		{in: "foo_url", out: "FooURL"},
		{in: "foobar", out: "Foobar"},
		{in: "url_sample", out: "URLSample"},
		{in: "_id", out: "ID"},
		{in: "__id", out: "ID"},
	}

	for _, testCase := range testCases {
		actual := FmtFieldName(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Fatalf("error FmtFieldName(): expected: %s | got: %s", expected, actual)
		}
	}
}

func TestGetImportKey(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "uri", out: ""},
		{in: "date-time", out: "date-time"},
	}

	for _, testCase := range testCases {
		schema := jsonschema.Schema{Format: &testCase.in}
		actual := getCustomTypeKey(schema)
		expected := testCase.out
		if expected != actual {
			t.Fatalf("error getImportKey(): expected: %s | got: %s", expected, actual)
		}
	}
}

func TestGetImportType(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		{in: "uri", out: ""},
		{in: "date-time", out: "time.Time"},
	}

	for _, testCase := range testCases {
		actual := getCustomType(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Fatalf("error getImportType(): expected: %s | got: %s", expected, actual)
		}
	}
}

func TestHasImportKey(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		in  string
		out bool
	}

	testCases := []TestCase{
		{in: "uri", out: false},
		{in: "date-time", out: true},
	}

	for _, testCase := range testCases {
		actual := hasCustomType(testCase.in)
		expected := testCase.out
		if expected != actual {
			t.Fatalf("error hasImportKey(): expected: %t | got: %t", expected, actual)
		}
	}
}
