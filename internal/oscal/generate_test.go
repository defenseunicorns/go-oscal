package oscal

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// TestGenerateModelTypes tests that we can generate the top-level "ComponentDefinition" struct name correctly.
func TestGenerateModelTypes(t *testing.T) {
	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	// Should we be calling this currently untested function directly? Could we mock this function call?
	// or is there further decoupling we could do?
	idMap, id := generateUniqueIdMap(oscalMap)

	// Instantiate variable for storing the data
	modelTypeMap := make(map[string][]string)

	expected := "ComponentDefinition"

	actual := generateModelTypes(idMap, id, strings.Split(id, "_")[2], []string{""}, modelTypeMap)

	if expected != actual {
		t.Errorf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}

}

// TestOscalSchemaVersion tests that the OSCAL schema version is correct.
func TestOscalSchemaVersion(t *testing.T) {
	oscalSchemaVersion := "1.0.4"

	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	// Check whether the top-level "$id" field is empty or not
	// if it is, fail the test and print error message
	// if it isn't, validate the string contains the proper version number
	if schemaVersionString := oscalMap["$id"].(string); schemaVersionString != "" {
		schemaVersionString = oscalMap["$id"].(string)

		expected := true
		actual := strings.Contains(schemaVersionString, oscalSchemaVersion)

		if expected != actual {
			t.Errorf("OSCAL JSON schema version %s was not found. The version may have been updated.", oscalSchemaVersion)
		}
	} else {
		t.Error("The top-level '$id' field was not found or is not populated. Please verify that the OSCAL JSON schema file is valid.")
	}

}

// TestFmtFieldName tests that we handle formatting a json string to a go struct correctly.
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

// parseOscalFileToMap reads and unmarshals the OSCAL JSON schema file
// into a map[string]interface{} structure for further processing/testing.
func parseOscalFileToMap() (map[string]interface{}, error) {
	oscalJSONSchemaFile := "../../test/oscal_component_schema.json"

	oscalMap := make(map[string]interface{})

	oscalBytes, err := os.ReadFile(oscalJSONSchemaFile)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(oscalBytes, &oscalMap); err != nil {
		return nil, err
	}

	return oscalMap, nil
}
