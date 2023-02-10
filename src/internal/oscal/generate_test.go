package oscal

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

// TestGenerateUniqueIdMapProperties tests that the generateUniqueIdMap function
// returns the correct 'properties'.
func TestGenerateUniqueIdMap(t *testing.T) {
	var actualProperties string
	expectedPropertiesFile := "../../../testdata/expected-properties.txt"

	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	actualMap := generateUniqueIdMap(oscalMap)

	// Check if there's a properties field.
	// If there is, loop over to collect the properties and assert we have the correct properties.
	// If there isn't, fail the test and print error message.
	if properties := actualMap[id].(map[string]interface{})["properties"]; properties != nil {
		properties := actualMap[id].(map[string]interface{})["properties"]

		// Store the properties fields to a string slice and sort it.
		propertiesSlice := make([]string, 0)
		for property := range properties.(map[string]interface{}) {
			propertiesSlice = append(propertiesSlice, property)
		}
		sort.Strings(propertiesSlice)

		// Convert the sorted string slice to a string so that we can assert the output.
		for _, property := range propertiesSlice {
			actualProperties += fmt.Sprintf("%s\n", property)
		}
	} else {
		t.Error("The 'properties' field was not found. Please verify the OSCAL JSON schema file is valid.")
	}

	expectedPropertiesBytes, err := os.ReadFile(expectedPropertiesFile)
	if err != nil {
		t.Error(err)
	}

	expected := string(expectedPropertiesBytes)

	actual := strings.TrimSpace(actualProperties)

	if expected != actual {
		t.Errorf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestFormatStructTags tests that we can format Go struct tags correctly.
func TestFormatStructTags(t *testing.T) {
	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	actualTagList := formatStructTags(idMap, id, "uuid", []string{"json", "yaml"})

	actualTagString := strings.Join(actualTagList, " ")

	expectedTagString := "json:\"uuid\" yaml:\"uuid\""

	if expectedTagString != actualTagString {
		t.Errorf("error formatStructTags(): expected: %s | got: %s", expectedTagString, actualTagString)
	}
}

// TestBuildStructData tests that we can construct Go struct data correctly.
func TestBuildStructData(t *testing.T) {
	var actualStructDataString string
	expectedStructDataFile := "../../../testdata/expected-struct-data.txt"

	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	modelTypeMap := make(map[string][]string)

	// Check if there's a properties field.
	// If there is, call the buildStructData function and store the result as a string for assertion.
	// If there isn't, fail the test and print error message.
	if properties := idMap[id].(map[string]interface{})["properties"]; properties != nil {
		actualStructData := buildStructData(properties, idMap, id, []string{"json", "yaml"}, []string{""}, modelTypeMap)
		// Sort the data in increasing order so that we always have consistent, predictable output.
		sort.Strings(actualStructData)

		for _, data := range actualStructData {
			actualStructDataString += fmt.Sprintf("%s\n", data)
		}
	} else {
		t.Error("The 'properties' field was not found. Please verify the OSCAL JSON schema file is valid.")
	}

	// Trim leading and trailing white space from struct data string output.
	actual := strings.TrimSpace(actualStructDataString)

	expectedStructDataBytes, err := os.ReadFile(expectedStructDataFile)
	if err != nil {
		t.Error(err)
	}

	expected := string(expectedStructDataBytes)

	if expected != actual {
		t.Errorf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateModelTypes tests that we can generate the top-level "ComponentDefinition" struct name correctly.
func TestGenerateModelTypes(t *testing.T) {
	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	modelTypeMap := make(map[string][]string)

	expected := "ComponentDefinition"

	actual := generateModelTypes(idMap, id, strings.Split(id, "_")[2], []string{""}, modelTypeMap)

	if expected != actual {
		t.Errorf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateOscalComponentDocumentStruct tests that we can generate the 'OscalComponentDocument' struct correctly.
func TestGenerateOscalComponentDocumentStruct(t *testing.T) {
	expectedOutputFile := "../../../testdata/expected-oscal-component-document-struct.txt"

	oscalMap, err := parseOscalFileToMap()
	if err != nil {
		t.Error(err)
	}

	expectedBytes, err := os.ReadFile(expectedOutputFile)
	if err != nil {
		t.Error(err)
	}

	expected := string(expectedBytes)

	actual := generateOscalComponentDocumentStruct(oscalMap, "", []string{"json", "yaml"})

	// TODO: Need to figure out a more accurate/reliable way to compare strings read from a file.
	if expected != actual {
		t.Errorf("error generateOscalComponentDocumentStruct():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// parseOscalFileToMap reads and unmarshals the OSCAL JSON schema file
// into a map[string]interface{} structure for further processing/testing.
func parseOscalFileToMap() (map[string]interface{}, error) {
	oscalJSONSchemaFile := "../../../testdata/schema/oscal_component_schema.json"

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
