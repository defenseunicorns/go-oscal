package oscal

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/swaggest/jsonschema-go"
)

const (
	oscalComponentSchemaFilePath    string = "../../../schema/component/oscal_component_schema-1-1-1.json"
	oscalSSPSchemaFilePath          string = "../../../schema/ssp/oscal_ssp_schema-1-1-1.json"
	fieldsPresentFilePath           string = "../../../testdata/fields-present.json"
	fieldsMissingFilePath           string = "../../../testdata/fields-missing.json"
	expected104FilePath             string = "../../types/oscal-1-0-4/types.go"
	oscal104FilePath                string = "../../../schema/complete/oscal_complete_schema-1-0-4.json"
	componentExpectedPropertiesFile string = "../../../testdata/generation/component/expected-properties.txt"
	sspExpectedPropertiesFile       string = "../../../testdata/generation/ssp/expected-properties.txt"
	componentExpectedStructDataFile string = "../../../testdata/generation/component/expected-struct-data.txt"
	sspExpectedStructDataFile       string = "../../../testdata/generation/ssp/expected-struct-data.txt"
	expectedComponentOutputFile     string = "../../../testdata/generation/component/expected-oscal-model-struct.txt"
	sspExpectedOutputFile           string = "../../../testdata/generation/ssp/expected-oscal-model-struct.txt"
)

func TestTypesOutput(t *testing.T) {
	expected, err := os.ReadFile(expected104FilePath)
	if err != nil {
		t.Fatal(err)
	}

	schema, err := os.ReadFile(oscal104FilePath)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := Generate(schema, "oscalTypes", []string{"json", "yaml"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("expected output to match")
	}
}

// TestGetOscalModel tests that we can get the value of the top-level 'required' field,
// which is the name of the OSCAL model, and that we can convert it to a string properly.
func TestGetOscalModel(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: fieldsPresentFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"test-data"}
	actual, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("error getOscalModel(): expected: %s | got: %s", expected, actual)
	}
}

// TestSetOscalModelRefComponent tests that we can set the OSCAL model $ref correctly for the OSCAL Component Definition schema.
func TestSetOscalModelRefComponent(t *testing.T) {
	expected := "#assembly_oscal-component-definition_component-definition"
	actual, err := setOscalModelRef("component-definition")
	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Fatalf("error setOscalModelRef(): expected: %s | got: %s", expected, actual)
	}
}

// TestSetOscalModelRefSSP tests that we can set the OSCAL model $ref correctly for the OSCAL SSP schema.
func TestSetOscalModelRefSSP(t *testing.T) {
	expected := "#assembly_oscal-ssp_system-security-plan"
	actual, err := setOscalModelRef("system-security-plan")
	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Fatalf("error setOscalModelRef(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateUniqueIdMapComponent tests that the generateUniqueIdMap function
// returns the correct 'properties' from the OSCAL Component Definition schema.
func TestGenerateUniqueIdMapComponent(t *testing.T) {

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	actualMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}

	properties := actualMap[modelId].TypeObject.Properties

	expected, err := readTestFile(componentExpectedPropertiesFile)
	if err != nil {
		t.Fatal(err)
	}
	actual := preparePropertiesForAssertion(properties)

	if expected != actual {
		t.Fatalf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateUniqueIdMapSSP tests that the generateUniqueIdMap function
// returns the correct 'properties' from the OSCAL SSP schema.
func TestGenerateUniqueIdMapSSP(t *testing.T) {

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	actualMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}

	properties := actualMap[modelId].TypeObject.Properties

	expected, err := readTestFile(sspExpectedPropertiesFile)
	if err != nil {
		t.Fatal(err)
	}
	actual := preparePropertiesForAssertion(properties)

	if expected != actual {
		t.Fatalf("error generateUniqueIdMap():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestFormatStructTagsComponent tests that we can format Go struct tags from the OSCAL Component Definition schema correctly.
func TestFormatStructTagsComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}
	actualTagList := formatStructTags(idMap, modelId, "uuid", []string{"json", "yaml"})

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

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}
	actualTagList := formatStructTags(idMap, modelId, "uuid", []string{"json", "yaml"})

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

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	properties := idMap[modelId].TypeObject.Properties
	actualStructData, err := buildStructData(properties, idMap, modelId, []string{"json", "yaml"}, []string{""}, modelTypeMap)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := readTestFile(componentExpectedStructDataFile)
	if err != nil {
		t.Fatal(err)
	}
	actual := sortStringSliceAndConvertToString(actualStructData)

	if expected != actual {
		t.Fatalf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestBuildStructDataSSP tests that we can construct Go struct data for the OSCAL SSP schema correctly.
func TestBuildStructDataSSP(t *testing.T) {

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	properties := idMap[modelId].TypeObject.Properties
	actualStructData, err := buildStructData(properties, idMap, modelId, []string{"json", "yaml"}, []string{""}, modelTypeMap)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := readTestFile(sspExpectedStructDataFile)
	if err != nil {
		t.Fatal(err)
	}
	actual := sortStringSliceAndConvertToString(actualStructData)

	if expected != actual {
		t.Fatalf("error buildStructData():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateModelTypesComponent tests that we can generate the top-level "ComponentDefinition" struct name correctly for the OSCAL Component Definition schema.
func TestGenerateModelTypesComponent(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	expected := "ComponentDefinition"
	actual, err := generateModelTypes(idMap, modelId, strings.Split(modelId, "_")[2], []string{""}, modelTypeMap)
	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Fatalf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateModelTypesSSP tests that we can generate the top-level "SystemSecurityPlan" struct name correctly for the OSCAL SSP schema.
func TestGenerateModelTypesSSP(t *testing.T) {
	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}
	modelId, err := setOscalModelRef(oscalModel[0])
	if err != nil {
		t.Fatal(err)
	}
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		t.Fatal(err)
	}

	modelTypeMap := make(map[string][]string)

	expected := "SystemSecurityPlan"
	actual, err := generateModelTypes(idMap, modelId, strings.Split(modelId, "_")[2], []string{""}, modelTypeMap)
	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Fatalf("error generateModelTypes(): expected: %s | got: %s", expected, actual)
	}
}

// TestGenerateOscalComponentModelStruct tests that we can generate the 'OscalModel' struct correctly for the OSCAL Component Definition schema.
func TestGenerateOscalComponentModelStruct(t *testing.T) {

	testdata := &BaseFlags{
		InputFile: oscalComponentSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := readTestFile(expectedComponentOutputFile)
	if err != nil {
		t.Fatal(err)
	}
	actualString, err := generateOscalModelStruct(oscalModel, "", []string{"json", "yaml"})
	if err != nil {
		t.Fatal(err)
	}
	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualString)

	if expected != actual {
		t.Fatalf("error generateOscalModelsStruct():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestGenerateOscalSSPModelStruct tests that we can generate the 'OscalModel' struct correctly for the OSCAL SSP schema.
func TestGenerateOscalSSPModelStruct(t *testing.T) {

	testdata := &BaseFlags{
		InputFile: oscalSSPSchemaFilePath,
	}

	schema, err := readFileToSchema(testdata.InputFile)
	if err != nil {
		t.Fatal(err)
	}

	oscalModel, err := getOscalModel(schema)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := readTestFile(sspExpectedOutputFile)
	if err != nil {
		t.Fatal(err)
	}
	actualString, err := generateOscalModelStruct(oscalModel, "", []string{"json", "yaml"})
	if err != nil {
		t.Fatal(err)
	}
	// Trim leading and trailing white space from string.
	actual := strings.TrimSpace(actualString)

	if expected != actual {
		t.Fatalf("error generateOscalModelsStruct():\n\nexpected: \n%s\n\ngot: \n%s", expected, actual)
	}
}

// TestHandleDuplicateStructNamesWithAssemblyKey tests that we can handle duplicate struct names correctly when map keys begin with #assembly.
func TestHandleDuplicateStructNamesWithAssemblyKey(t *testing.T) {
	existingValueMap := map[string]bool{}

	// Populate the map with a random value to test against.
	existingValueMap["RandomValue"] = true

	// The fake data we're passing in mocks a map key that begins with #assembly,
	// and a map value (RandomValue) that already exists, so we're expecting the handleDuplicateStructNames() function
	// to format the map key and use it as the struct name.
	expected := fmt.Sprintf("\ntype %s struct {", "AssemblyTestDataRandom")
	actual := handleDuplicateStructNames(existingValueMap, "#assembly_test-data_random", "RandomValue")

	if expected != actual {
		t.Fatalf("error handleDuplicateStructNames():\nexpected:%s\n\ngot: %s", expected, actual)
	}
}

// TestHandleDuplicateStructNamesWithIdenticalKeyValue tests that we can handle duplicate struct names correctly when a map key is identical to its value.
func TestHandleDuplicateStructNamesWithIdenticalKeyValue(t *testing.T) {
	existingValueMap := map[string]bool{}

	// Populate the map with a random value to test against.
	existingValueMap["RandomValue"] = true

	// The fake data we're passing in mocks a map key that is identical to its value,
	// so the key and value is RandomValue, which already exists, so we're expecting the handleDuplicateStructNames() function
	// to append a number to the value as a unique identifier.
	expected := fmt.Sprintf("\ntype %s struct {", "RandomValue1")
	actual := handleDuplicateStructNames(existingValueMap, "RandomValue", "RandomValue")

	if expected != actual {
		t.Fatalf("error handleDuplicateStructNames():\nexpected:%s\n\ngot: %s", expected, actual)
	}
}

// preparePropertiesForAssertion converts 'properties' values in an OSCAL schema file
// to strings and sorts them in alphabetical order for asserting the output against test data.
func preparePropertiesForAssertion(properties map[string]jsonschema.SchemaOrBool) string {
	// Store the properties fields to a string slice and sort it.
	propertiesSlice := make([]string, 0)
	for property := range properties {
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
func readTestFile(testFile string) (string, error) {
	dataBytes, err := os.ReadFile(testFile)
	if err != nil {
		return "", err
	}

	testFileString := string(dataBytes)

	return testFileString, nil
}

func readFileToSchema(filepath string) (jsonschema.Schema, error) {

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return jsonschema.Schema{}, err
	}

	schema := jsonschema.Schema{}
	schema.UnmarshalJSON(bytes)

	return schema, nil

}
