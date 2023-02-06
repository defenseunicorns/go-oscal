package oscal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestOscalSchemaVersion tests that the OSCAL schema version is correct.
func TestOscalSchemaVersion(t *testing.T) {
	oscalSchemaVersion := "1.0.4"

	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	// Check whether the top-level "$id" field is empty or not.
	// If it is, fail the test and print error message.
	// If it isn't, validate the string contains the proper version number.
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

// TestGenerateUniqueIdMap tests that the generateUniqueIdMap function
// returns the correct '$id' and 'properties'.
func TestGenerateUniqueIdMap(t *testing.T) {
	var actualProperties string

	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	actualMap, actualId := generateUniqueIdMap(oscalMap)

	expectedId := "#assembly_oscal-component-definition_component-definition"

	if expectedId != actualId {
		t.Errorf("error generateUniqueIdMap(): expected: %s | got: %s", expectedId, actualId)
	}

	// Check if there's a properties field.
	// If there is, loop over to collect the properties and assert we have the correct properties.
	// If there isn't, fail the test and print error message.
	if properties := actualMap[actualId].(map[string]interface{})["properties"]; properties != nil {
		properties := actualMap[actualId].(map[string]interface{})["properties"]

		for property := range properties.(map[string]interface{}) {
			actualProperties += fmt.Sprintf("%s\n", property)
		}
	} else {
		t.Error("The 'properties' field was not found. Please verify the OSCAL JSON schema file is valid.")
	}

	if !strings.Contains(actualProperties, "uuid") {
		t.Errorf("error generateUniqueIdMap(): expected 'uuid' property to be present, but wasn't found. \n%s", actualProperties)
	}

	if !strings.Contains(actualProperties, "metadata") {
		t.Errorf("error generateUniqueIdMap(): expected 'metadata' property to be present, but wasn't found. \n%s", actualProperties)
	}

	if !strings.Contains(actualProperties, "import-component-definitions") {
		t.Errorf("error generateUniqueIdMap(): expected 'import-component-definitions' property to be present, but wasn't found. \n%s", actualProperties)
	}

	if !strings.Contains(actualProperties, "components") {
		t.Errorf("error generateUniqueIdMap(): expected 'components' property to be present, but wasn't found. \n%s", actualProperties)
	}

	if !strings.Contains(actualProperties, "capabilities") {
		t.Errorf("error generateUniqueIdMap(): expected 'capabilities' property to be present, but wasn't found. \n%s", actualProperties)
	}

	if !strings.Contains(actualProperties, "back-matter") {
		t.Errorf("error generateUniqueIdMap(): expected 'back-matter' property to be present, but wasn't found. \n%s", actualProperties)
	}

}

// TestFormatStructTags tests that we can format Go struct tags correctly.
func TestFormatStructTags(t *testing.T) {
	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	idMap, id := generateUniqueIdMap(oscalMap)

	actualTagList := formatStructTags(idMap, id, "uuid", []string{"json", "yaml"})

	actualTagString := strings.Join(actualTagList, " ")

	expectedTagString := "json:\"uuid\" yaml:\"uuid\""

	if expectedTagString != actualTagString {
		t.Errorf("error formatStructTags(): expected: %s | got: %s", expectedTagString, actualTagString)
	}
}

// TestGenerateModelTypes tests that we can generate the top-level "ComponentDefinition" struct name correctly.
func TestGenerateModelTypes(t *testing.T) {
	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	idMap, id := generateUniqueIdMap(oscalMap)

	modelTypeMap := make(map[string][]string)

	expected := "ComponentDefinition"

	actual := generateModelTypes(idMap, id, strings.Split(id, "_")[2], []string{""}, modelTypeMap)

	if expected != actual {
		t.Errorf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
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
