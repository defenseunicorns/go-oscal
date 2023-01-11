package oscal

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

var examplesDir string = "../../examples"
var pkgName string = "oscal"

// TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
// It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar"}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestSimpleArray tests that an array without conflicting types is handled correctly
func TestSimpleArray(t *testing.T) {
	i := strings.NewReader(`{"foo" : [{"bar": 24}, {"bar" : 42}]}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestInvalidFieldChars tests that a document with invalid field chars is handled correctly
func TestInvalidFieldChars(t *testing.T) {
	i := strings.NewReader(`{"f.o-o" : 42}`)
	if _, err := Generate(i, ParseJson, "TestStruct", pkgName, []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestInferFloatInt tests that we can correctly infer a float or an int from a
// JSON number when no command-line flag is provided.
func TestInferFloatInt(t *testing.T) {
	f, err := os.Open(filepath.Join(examplesDir, "floats.json"))
	if err != nil {
		t.Fatalf("error opening "+examplesDir+"/floats.json: %s", err)
	}
	defer f.Close()

	expected, err := os.ReadFile(filepath.Join(examplesDir, "expected-floats.go.out"))
	if err != nil {
		t.Fatalf("error reading expected-floats.go.out: %s", err)
	}

	actual, err := Generate(f, ParseJson, "Stats", pkgName, []string{"json"}, false, true)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}

}

// TestYamlNumbers tests that we handle Yaml's number system that has both floats and ints correctly
func TestYamlNumbers(t *testing.T) {
	f, err := os.Open(filepath.Join(examplesDir, "numbers.yaml"))
	if err != nil {
		t.Fatalf("error opening "+examplesDir+"/numbers.yaml: %s", err)
	}
	defer f.Close()

	expected, err := os.ReadFile(filepath.Join(examplesDir, "expected-numbers.go.out"))
	if err != nil {
		t.Fatalf("error reading expected-numbers.go.out: %s", err)
	}

	actual, err := Generate(f, ParseYaml, "Stats", pkgName, []string{"yaml"}, false, false)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}

// Test example document
func TestExample(t *testing.T) {
	i, err := os.Open(filepath.Join(examplesDir, "example.json"))
	if err != nil {
		t.Error("error opening example.json", err)
	}

	expected, err := os.ReadFile(filepath.Join(examplesDir, "expected-output-test.go.out"))
	if err != nil {
		t.Error("error reading expected-output-test.go", err)
	}

	actual, err := Generate(i, ParseJson, "User", pkgName, []string{"json"}, false, true)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}

// TestFmtFieldName tests that we handle formatting a json string to a go struct correctly
func TestFmtFieldName(t *testing.T) {
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
			t.Errorf("error FmtFieldName(): expected: %s | got: %s", expected, actual)
		}
	}
}

// Test example document
func TestExampleArray(t *testing.T) {
	i, err := os.Open(filepath.Join(examplesDir, "example-array.json"))
	if err != nil {
		t.Fatalf("error opening example.json: %s", err)
	}
	defer i.Close()

	expectedf, err := os.Open(filepath.Join(examplesDir, "example-array.go.out"))
	if err != nil {
		t.Fatalf("error opening example-array.go: %s", err)
	}
	defer expectedf.Close()

	expectedBts, err := io.ReadAll(expectedf)
	if err != nil {
		t.Fatalf("error reading example-array.go: %s", err)
	}

	actual, err := Generate(i, ParseJson, "Users", pkgName, []string{"json"}, false, true)
	if err != nil {
		t.Fatal(err)
	}
	sactual, sexpected := string(actual), string(expectedBts)
	if sactual != sexpected {
		t.Fatalf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}

// TestExampleOscalComponentDefinition tests that we can unmarshal
// an oscal component definition yaml file into the generated go data types
func TestExampleOscalComponentDefinition(t *testing.T) {
	file := filepath.Join(examplesDir, "example-component-definition.yaml")

	rawYaml, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	oscalComponentDefinition := OscalComponentDefinition{}
	componentDefinition := oscalComponentDefinition.Definitions.Oscal_component_definition_oscal_component_definition_component_definition.Properties

	err = yaml.Unmarshal(rawYaml, &componentDefinition)
	if err != nil {
		t.Fatal(err)
	}
}
