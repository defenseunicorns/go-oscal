package oscal

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

var oscalJSONSchemaFile string = "../../test/oscal_component_schema.json"

// Test that the OSCAL schema version is correct
func TestOscalSchemaVersion(t *testing.T) {
	oscalMap := make(map[string]interface{})

	oscalBytes, err := os.ReadFile(oscalJSONSchemaFile)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(oscalBytes, &oscalMap)
	if err != nil {
		t.Error(err)
	}

	// Check whether the top-level "$id" field is empty or not
	// if it is, fail the test and print error message
	// if it isn't, validate the string contains the proper version number
	if schemaVersion := oscalMap["$id"].(string); schemaVersion != "" {
		schemaVersion = oscalMap["$id"].(string)

		expected := true
		actual := strings.Contains(schemaVersion, "1.0.4")

		if expected != actual {
			t.Error("OSCAL JSON schema version 1.0.4 was not found. The version may have been updated.")
		}
	} else {
		t.Error("The top-level '$id' field was not found or is not populated. Please verify that the OSCAL JSON schema file is valid.")
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
