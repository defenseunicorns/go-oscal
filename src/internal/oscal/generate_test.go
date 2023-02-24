package oscal

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

const (
	oscalComponentSchemaFilePath = "../../../testdata/schema/component/oscal_component_schema.json"
	oscalSSPSchemaFilePath       = "../../../testdata/schema/ssp/oscal_ssp_schema.json"
)

// TestOscalComponentSchemaVersion tests that the OSCAL component schema version is correct.
func TestOscalComponentSchemaVersion(t *testing.T) {
	oscalSchemaVersion := "1.0.4"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	// Check whether the top-level "$id" field exists and whether it is empty or not.
	// If it doesn't exist or is empty, fail the test and print error message.
	// If it exists and is populated, validate the string contains the proper version number.
	if schemaVersionString, ok := oscalMap["$id"].(string); ok && schemaVersionString != "" {
		schemaVersionString = oscalMap["$id"].(string)

		expected := true
		actual := strings.Contains(schemaVersionString, oscalSchemaVersion)

		if expected != actual {
			t.Fatalf("OSCAL JSON schema version %s was not found. The version may have been updated.", oscalSchemaVersion)
		}
	} else {
		t.Fatal("The top-level '$id' field was not found or is not populated. Please verify that the OSCAL JSON schema file is valid.")
	}
}

// TestOscalSSPSchemaVersion tests that the OSCAL SSP schema version is correct.
func TestOscalSSPSchemaVersion(t *testing.T) {
	oscalSchemaVersion := "1.0.4"

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	// Check whether the top-level "$id" field exists and whether it is empty or not.
	// If it doesn't exist or is empty, fail the test and print error message.
	// If it exists and is populated, validate the string contains the proper version number.
	if schemaVersionString, ok := oscalMap["$id"].(string); ok && schemaVersionString != "" {
		schemaVersionString = oscalMap["$id"].(string)

		expected := true
		actual := strings.Contains(schemaVersionString, oscalSchemaVersion)

		if expected != actual {
			t.Fatalf("OSCAL JSON schema version %s was not found. The version may have been updated.", oscalSchemaVersion)
		}
	} else {
		t.Fatal("The top-level '$id' field was not found or is not populated. Please verify that the OSCAL JSON schema file is valid.")
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
			t.Fatalf("error FmtFieldName(): expected: %s | got: %s", expected, actual)
		}
	}
}

// TestGenerateUniqueIdMapComponent tests that the generateUniqueIdMap function
// returns the correct 'properties' from the OSCAL component schema.
func TestGenerateUniqueIdMapComponent(t *testing.T) {
	var actualProperties string
	expectedPropertiesFile := "../../../testdata/schema/component/expected-properties.txt"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	actualMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	// Check if there's a properties field and whether it is populated.
	// If there is, loop over to collect the properties and assert we have the correct properties.
	// If there isn't, fail the test and print error message.
	if properties, ok := actualMap[oscalModelId].(map[string]interface{})["properties"]; ok && properties != nil {
		properties := actualMap[oscalModelId].(map[string]interface{})["properties"]

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
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}

	expectedPropertiesBytes, err := os.ReadFile(expectedPropertiesFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedPropertiesBytes)

	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualProperties)

	if expected != actual {
		t.Fatalf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateUniqueIdMapSSP tests that the generateUniqueIdMap function
// returns the correct 'properties' from the OSCAL SSP schema.
func TestGenerateUniqueIdMapSSP(t *testing.T) {
	var actualProperties string
	expectedPropertiesFile := "../../../testdata/schema/ssp/expected-properties.txt"

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	actualMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	// Check if there's a properties field and whether it is populated.
	// If there is, loop over to collect the properties and assert we have the correct properties.
	// If there isn't, fail the test and print error message.
	if properties, ok := actualMap[oscalModelId].(map[string]interface{})["properties"]; ok && properties != nil {
		properties := actualMap[oscalModelId].(map[string]interface{})["properties"]

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
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}

	expectedPropertiesBytes, err := os.ReadFile(expectedPropertiesFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedPropertiesBytes)

	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualProperties)

	if expected != actual {
		t.Fatalf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestFormatStructTagsComponent tests that we can format Go struct tags from the OSCAL component schema correctly.
func TestFormatStructTagsComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	actualTagList := formatStructTags(idMap, oscalModelId, "uuid", []string{"json", "yaml"})

	// Convert the string slice to a string for assertion.
	actual := strings.Join(actualTagList, " ")

	expected := "json:\"uuid\" yaml:\"uuid\""

	if expected != actual {
		t.Fatalf("error formatStructTags(): expected: %s | got: %s", expected, actual)
	}
}

// TestFormatStructTagsSSP tests that we can format Go struct tags from the OSCAL SSP schema correctly.
func TestFormatStructTagsSSP(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	actualTagList := formatStructTags(idMap, oscalModelId, "uuid", []string{"json", "yaml"})

	// Convert the string slice to a string for assertion.
	actual := strings.Join(actualTagList, " ")

	expected := "json:\"uuid\" yaml:\"uuid\""

	if expected != actual {
		t.Fatalf("error formatStructTags(): expected: %s | got: %s", expected, actual)
	}
}

// TestBuildStructDataComponent tests that we can construct Go struct data for the OSCAL component schema correctly.
func TestBuildStructDataComponent(t *testing.T) {
	var actualStructDataString string
	expectedStructDataFile := "../../../testdata/schema/component/expected-struct-data.txt"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	// Check if there's a properties field and whether it is populated.
	// If there is, call the buildStructData function and store the result as a string for assertion.
	// If there isn't, fail the test and print error message.
	if properties, ok := idMap[oscalModelId].(map[string]interface{})["properties"]; ok && properties != nil {
		actualStructData := buildStructData(properties, idMap, oscalModelId, []string{"json", "yaml"}, []string{""}, modelTypeMap)
		// Sort the data in increasing order so that we always have consistent, predictable output.
		sort.Strings(actualStructData)

		// Convert the sorted string slice to a string so that we can assert the output.
		for _, data := range actualStructData {
			actualStructDataString += fmt.Sprintf("%s\n", data)
		}
	} else {
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}

	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualStructDataString)

	expectedStructDataBytes, err := os.ReadFile(expectedStructDataFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedStructDataBytes)

	if expected != actual {
		t.Fatalf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestBuildStructDataSSP tests that we can construct Go struct data for the OSCAL SSP schema correctly.
func TestBuildStructDataSSP(t *testing.T) {
	var actualStructDataString string
	expectedStructDataFile := "../../../testdata/schema/ssp/expected-struct-data.txt"

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	// Check if there's a properties field and whether it is populated.
	// If there is, call the buildStructData function and store the result as a string for assertion.
	// If there isn't, fail the test and print error message.
	if properties, ok := idMap[oscalModelId].(map[string]interface{})["properties"]; ok && properties != nil {
		actualStructData := buildStructData(properties, idMap, oscalModelId, []string{"json", "yaml"}, []string{""}, modelTypeMap)
		// Sort the data in increasing order so that we always have consistent, predictable output.
		sort.Strings(actualStructData)

		// Convert the sorted string slice to a string so that we can assert the output.
		for _, data := range actualStructData {
			actualStructDataString += fmt.Sprintf("%s\n", data)
		}
	} else {
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}

	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualStructDataString)

	expectedStructDataBytes, err := os.ReadFile(expectedStructDataFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedStructDataBytes)

	if expected != actual {
		t.Fatalf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateModelTypesComponent tests that we can generate the top-level "ComponentDefinition" struct name correctly for the OSCAL component schema.
func TestGenerateModelTypesComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	expected := "ComponentDefinition"

	actual := generateModelTypes(idMap, oscalModelId, strings.Split(oscalModelId, "_")[2], []string{""}, modelTypeMap)

	if expected != actual {
		t.Fatalf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateModelTypesSSP tests that we can generate the top-level "SystemSecurityPlan" struct name correctly for the OSCAL SSP schema.
func TestGenerateModelTypesSSP(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	idMap := generateUniqueIdMap(oscalMap)

	oscalModelId, err := setOscalModelRef(oscalMap)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	expected := "SystemSecurityPlan"

	actual := generateModelTypes(idMap, oscalModelId, strings.Split(oscalModelId, "_")[2], []string{""}, modelTypeMap)

	if expected != actual {
		t.Fatalf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateOscalComponentModelStruct tests that we can generate the 'OscalModel' struct correctly for the OSCAL component schema.
func TestGenerateOscalComponentModelStruct(t *testing.T) {
	expectedOutputFile := "../../../testdata/schema/component/expected-oscal-model-struct.txt"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	expectedBytes, err := os.ReadFile(expectedOutputFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedBytes)

	actualString := generateOscalModelStruct(oscalMap, "", []string{"json", "yaml"})

	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualString)

	if expected != actual {
		t.Fatalf("error generateOscalModelsStruct():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateOscalSSPModelStruct tests that we can generate the 'OscalModel' struct correctly for the OSCAL SSP schema.
func TestGenerateOscalSSPModelStruct(t *testing.T) {
	expectedOutputFile := "../../../testdata/schema/ssp/expected-oscal-model-struct.txt"

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	expectedBytes, err := os.ReadFile(expectedOutputFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedBytes)

	actualString := generateOscalModelStruct(oscalMap, "", []string{"json", "yaml"})

	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualString)

	if expected != actual {
		t.Fatalf("error generateOscalModelStruct():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}
