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
	fieldsPresentFilePath        = "../../../testdata/fields-present.json"
	fieldsMissingFilePath        = "../../../testdata/fields-missing.json"
)

// TestOscalComponentSchemaVersion tests that the OSCAL Component Definition schema version is correct.
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

// TestCheckTopLevelRequiredFieldExists tests that we can check if an OSCAL schema file has a top-level 'required' field correctly.
// This test case uses faked json data that contains a 'required' field, so the checkTopLevelRequiredField() function
// that we're testing should return true.
func TestCheckTopLevelRequiredFieldExists(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: fieldsPresentFilePath,
	}

	testMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	// We expect the checkTopLevelRequiredField() function to return true because the 'required' field is present in our fake data.
	expected := true
	actual := checkTopLevelRequiredField(testMap)

	if expected != actual {
		t.Fatal("error checkTopLevelRequiredField(): the function should have returned true if a 'required' field is present.")
	}
}

// TestCheckTopLevelRequiredFieldMissing tests that we can check if an OSCAL schema file is missing a top-level 'required' field correctly.
// This test case uses faked json data that does not have a 'required' field, so the checkTopLevelRequiredField() function
// that we're testing should return false.
func TestCheckTopLevelRequiredFieldMissing(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: fieldsMissingFilePath,
	}

	testMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	// We expect the checkTopLevelRequiredField() function to return false because the 'required' field is missing in our fake data.
	expected := false
	actual := checkTopLevelRequiredField(testMap)

	if expected != actual {
		t.Fatal("error checkTopLevelRequiredField(): the function should have returned false if a 'required' field is not present.")
	}
}

// TestGetOscalModel tests that we can get the value of the top-level 'required' field,
// which is the name of the OSCAL model, and that we can convert it to a string properly.
func TestGetOscalModel(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: fieldsPresentFilePath,
	}

	testMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	expected := "test-data"
	actual := getOscalModel(testMap)

	if expected != actual {
		t.Fatalf("error getOscalModel(): expected: %s | got: %s", expected, actual)
	}
}

// TestSetOscalModelRefComponent tests that we can set the OSCAL model $ref correctly for the OSCAL Component Definition schema.
func TestSetOscalModelRefComponent(t *testing.T) {
	expected := "#assembly_oscal-component-definition_component-definition"
	actual := setOscalModelRef("component-definition")

	if expected != actual {
		t.Fatalf("error setOscalModelRef(): expected: %s | got: %s", expected, actual)
	}
}

// TestSetOscalModelRefSSP tests that we can set the OSCAL model $ref correctly for the OSCAL SSP schema.
func TestSetOscalModelRefSSP(t *testing.T) {
	expected := "#assembly_oscal-ssp_system-security-plan"
	actual := setOscalModelRef("system-security-plan")

	if expected != actual {
		t.Fatalf("error setOscalModelRef(): expected: %s | got: %s", expected, actual)
	}
}

// TestCheckPropertiesFieldExistsComponent tests that we can check if an OSCAL Component Definition schema file has a 'properties' field correctly.
func TestCheckPropertiesFieldExistsComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	modelId := setOscalModelRef(oscalModel)
	actualMap := generateUniqueIdMap(oscalMap)

	// We expect the checkPropertiesField() function to return true because the 'properties' field is present in the OSCAL Component Definition schema.
	expected := true
	actual := checkPropertiesField(actualMap, modelId)

	if expected != actual {
		t.Fatal("error checkPropertiesField(): the function should have returned true if a 'properties' field is present.")
	}
}

// TestCheckPropertiesFieldExistsSSP tests that we can check if an OSCAL SSP schema file has a 'properties' field correctly.
func TestCheckPropertiesFieldExistsSSP(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	modelId := setOscalModelRef(oscalModel)
	actualMap := generateUniqueIdMap(oscalMap)

	// We expect the checkPropertiesField() function to return true because the 'properties' field is present in the OSCAL SSP schema.
	expected := true
	actual := checkPropertiesField(actualMap, modelId)

	if expected != actual {
		t.Fatal("error checkPropertiesField(): the function should have returned true if a 'properties' field is present.")
	}
}

// TestGenerateUniqueIdMapComponent tests that the generateUniqueIdMap function
// returns the correct 'properties' from the OSCAL Component Definition schema.
func TestGenerateUniqueIdMapComponent(t *testing.T) {
	expectedPropertiesFile := "../../../testdata/schema/component/expected-properties.txt"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	modelId := setOscalModelRef(oscalModel)
	actualMap := generateUniqueIdMap(oscalMap)

	// Check if there's a properties field and whether it is populated.
	// If there is, assert we have the correct properties.
	// If there isn't, fail the test and print error message.
	if checkPropertiesField(actualMap, modelId) {
		properties := getPropertiesFieldValue(actualMap, modelId)

		expected := readTestFile(expectedPropertiesFile)
		actual := preparePropertiesForAssertion(properties)

		if expected != actual {
			t.Fatalf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
		}
	} else {
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}
}

// TestGenerateUniqueIdMapSSP tests that the generateUniqueIdMap function
// returns the correct 'properties' from the OSCAL SSP schema.
func TestGenerateUniqueIdMapSSP(t *testing.T) {
	expectedPropertiesFile := "../../../testdata/schema/ssp/expected-properties.txt"

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	modelId := setOscalModelRef(oscalModel)
	actualMap := generateUniqueIdMap(oscalMap)

	// Check if there's a properties field and whether it is populated.
	// If there is, assert we have the correct properties.
	// If there isn't, fail the test and print error message.
	if checkPropertiesField(actualMap, modelId) {
		properties := getPropertiesFieldValue(actualMap, modelId)

		expected := readTestFile(expectedPropertiesFile)
		actual := preparePropertiesForAssertion(properties)

		if expected != actual {
			t.Fatalf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
		}
	} else {
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}
}

// TestFormatStructTagsComponent tests that we can format Go struct tags from the OSCAL Component Definition schema correctly.
func TestFormatStructTagsComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	oscalModelId := setOscalModelRef(oscalModel)
	idMap := generateUniqueIdMap(oscalMap)
	actualTagList := formatStructTags(idMap, oscalModelId, "uuid", []string{"json", "yaml"})

	expected := "json:\"uuid\" yaml:\"uuid\""
	// Convert the string slice to a string for assertion.
	actual := strings.Join(actualTagList, " ")

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

	oscalModel := getOscalModel(oscalMap)
	oscalModelId := setOscalModelRef(oscalModel)
	idMap := generateUniqueIdMap(oscalMap)
	actualTagList := formatStructTags(idMap, oscalModelId, "uuid", []string{"json", "yaml"})

	expected := "json:\"uuid\" yaml:\"uuid\""
	// Convert the string slice to a string for assertion.
	actual := strings.Join(actualTagList, " ")

	if expected != actual {
		t.Fatalf("error formatStructTags(): expected: %s | got: %s", expected, actual)
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

// TestBuildStructDataComponent tests that we can construct Go struct data for the OSCAL Component Definition schema correctly.
func TestBuildStructDataComponent(t *testing.T) {
	expectedStructDataFile := "../../../testdata/schema/component/expected-struct-data.txt"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	modelId := setOscalModelRef(oscalModel)
	idMap := generateUniqueIdMap(oscalMap)

	modelTypeMap := make(map[string][]string)

	// Check if there's a properties field and whether it is populated.
	// If there is, call the buildStructData function and store the result as a string for assertion.
	// If there isn't, fail the test and print error message.
	if checkPropertiesField(idMap, modelId) {
		properties := getPropertiesFieldValue(idMap, modelId)
		actualStructData := buildStructData(properties, idMap, modelId, []string{"json", "yaml"}, []string{""}, modelTypeMap)

		expected := readTestFile(expectedStructDataFile)
		actual := sortStringSliceAndConvertToString(actualStructData)

		if expected != actual {
			t.Fatalf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
		}
	} else {
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}
}

// TestBuildStructDataSSP tests that we can construct Go struct data for the OSCAL SSP schema correctly.
func TestBuildStructDataSSP(t *testing.T) {
	expectedStructDataFile := "../../../testdata/schema/ssp/expected-struct-data.txt"

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	modelId := setOscalModelRef(oscalModel)
	idMap := generateUniqueIdMap(oscalMap)

	modelTypeMap := make(map[string][]string)

	// Check if there's a properties field and whether it is populated.
	// If there is, call the buildStructData function and store the result as a string for assertion.
	// If there isn't, fail the test and print error message.
	if checkPropertiesField(idMap, modelId) {
		properties := getPropertiesFieldValue(idMap, modelId)
		actualStructData := buildStructData(properties, idMap, modelId, []string{"json", "yaml"}, []string{""}, modelTypeMap)

		expected := readTestFile(expectedStructDataFile)
		actual := sortStringSliceAndConvertToString(actualStructData)

		if expected != actual {
			t.Fatalf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
		}
	} else {
		t.Fatal("The 'properties' field was not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
	}
}

// TestGenerateModelTypesComponent tests that we can generate the top-level "ComponentDefinition" struct name correctly for the OSCAL Component Definition schema.
func TestGenerateModelTypesComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel := getOscalModel(oscalMap)
	oscalModelId := setOscalModelRef(oscalModel)
	idMap := generateUniqueIdMap(oscalMap)

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

	oscalModel := getOscalModel(oscalMap)
	oscalModelId := setOscalModelRef(oscalModel)
	idMap := generateUniqueIdMap(oscalMap)

	modelTypeMap := make(map[string][]string)

	expected := "SystemSecurityPlan"
	actual := generateModelTypes(idMap, oscalModelId, strings.Split(oscalModelId, "_")[2], []string{""}, modelTypeMap)

	if expected != actual {
		t.Fatalf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateOscalComponentModelStruct tests that we can generate the 'OscalModel' struct correctly for the OSCAL Component Definition schema.
func TestGenerateOscalComponentModelStruct(t *testing.T) {
	expectedOutputFile := "../../../testdata/schema/component/expected-oscal-model-struct.txt"

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	oscalMap, err := ParseJson(testdata)
	if err != nil {
		t.Fatal(err)
	}

	expected := readTestFile(expectedOutputFile)
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

	expected := readTestFile(expectedOutputFile)
	actualString := generateOscalModelStruct(oscalMap, "", []string{"json", "yaml"})
	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualString)

	if expected != actual {
		t.Fatalf("error generateOscalModelStruct():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// sortPropertiesStringsForAssertion converts 'properties' values in an OSCAL schema file
// to strings and sorts them in alphabetical order for asserting the output against test data.
func preparePropertiesForAssertion(properties interface{}) string {
	// Store the properties fields to a string slice and sort it.
	propertiesSlice := make([]string, 0)
	for property := range properties.(map[string]interface{}) {
		propertiesSlice = append(propertiesSlice, property)
	}

	propertiesString := sortStringSliceAndConvertToString(propertiesSlice)

	return propertiesString
}

// sortStringSliceAndConvertToString sorts a string slice in increasing order,
// converts the string slice to a string, and trims away any white space for assertion.
func sortStringSliceAndConvertToString(dataSlice []string) string {
	var dataString string

	sort.Strings(dataSlice)

	// Convert the sorted string slice to a string so that we can assert the output.
	for _, data := range dataSlice {
		dataString += fmt.Sprintf("%s\n", data)
	}

	// Trim leading and trailing white space from string.
	testDataString := strings.TrimSpace(dataString)

	return testDataString
}

// readTestFile reads data from a file and returns it as a string.
func readTestFile(testFile string) string {
	dataBytes, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Printf("Failed to read data from file at %s", testFile)
		os.Exit(1)
	}

	testFileString := string(dataBytes)

	return testFileString
}
