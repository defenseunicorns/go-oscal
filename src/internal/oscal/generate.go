package oscal

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// BaseFlags represents command-line flags for the base go-oscal command.
type BaseFlags struct {
	InputFile  string // -f / --input-file
	OutputFile string // -o / --output-file
	Pkg        string // -p / --pkg
	Tags       string // -t / --tags
}

const headerComment string = `/*
	This file was auto-generated with go-oscal.

	To regenerate:
	
	go-oscal \
		--input-file <path_to_oscal_json_schema_file> \
		--output-file <name_of_go_types_file> // the path to this file must already exist \
		--tags json,yaml // the tags to add to the Go structs \
		--pkg <name_of_your_go_package> // defaults to "main"

	For more information on how to use go-oscal: go-oscal --help

	Source: https://github.com/defenseunicorns/go-oscal
	*/`

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"NTP":   true,
	"DB":    true,
}

var intToWordMap = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

// Generate a struct definition given a JSON string representation of an object.
func Generate(oscalSchema map[string]interface{}, pkgName string, tags []string) ([]byte, error) {
	oscalModel := getOscalModel(oscalSchema)

	modelId := setOscalModelRef(oscalModel)

	// Generate a map with unique Id as key and existing interface as value
	idMap := generateUniqueIdMap(oscalSchema)

	// Instantiate variable for storing the data
	modelTypeMap := make(map[string][]string)

	generateModelTypes(idMap, modelId, strings.Split(modelId, "_")[2], tags, modelTypeMap)

	// Construct header comment and package name.
	structString := fmt.Sprintf("%s\n\npackage %s\n", headerComment, pkgName)

	// Construct top-level OscalModel struct.
	structString += generateOscalModelStruct(oscalSchema, pkgName, tags)

	// Construct structs for oscal models.
	structString += generateStruct(modelTypeMap)

	formattedStruct, err := format.Source([]byte(structString))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, structString)
		return nil, err
	}

	return formattedStruct, nil
}

/*
ParseJson reads a user-provided oscal json schema file as input,
stores it to a map[string]interface{} pointer,
and returns the map[string]interface{}.
*/
func ParseJson(opts *BaseFlags) (map[string]interface{}, error) {
	oscalMap := map[string]interface{}{}

	bytes, err := os.ReadFile(opts.InputFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &oscalMap); err != nil {
		return nil, err
	}

	return oscalMap, nil
}

// setOscalModelRef determines which OSCAL model $ref to use based on the model name.
func setOscalModelRef(oscalModel string) string {
	// Check which OSCAL model we're working with, and set the $ref accordingly.
	switch oscalModel {
	case "system-security-plan":
		oscalModelRef := "#assembly_oscal-ssp_system-security-plan"
		return oscalModelRef
	case "component-definition":
		oscalModelRef := "#assembly_oscal-component-definition_component-definition"
		return oscalModelRef
	default:
		fmt.Println("Unsupported OSCAL model. Currently supported OSCAL models are Component Definition and System Security Plan.")
		os.Exit(1)
	}

	return ""
}

// getOscalModel determines which OSCAL model we're working with.
func getOscalModel(oscalSchema map[string]interface{}) string {
	if checkTopLevelRequiredField(oscalSchema) {
		requiredFieldValue := getRequiredFieldValue(oscalSchema)
		oscalModel := convertRequiredFieldInterfaceToString(requiredFieldValue)
		return oscalModel
	} else {
		fmt.Println("Top-level 'required' field not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
		os.Exit(1)
	}

	return ""
}

func generateUniqueIdMap(obj map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	if definitions, ok := obj["definitions"].(map[string]interface{}); ok && definitions != nil {
		for _, val := range definitions {
			result[val.(map[string]interface{})["$id"].(string)] = val
		}

	}

	return result
}

// getRequiredFields collects the required fields in each OSCAL model definition.
func getRequiredFields(obj map[string]interface{}, structId string) map[string]bool {
	requiredFields := make(map[string]bool)

	if required, ok := obj[structId].(map[string]interface{})["required"]; ok && required != nil {
		for _, v := range required.([]interface{}) {
			requiredFields[v.(string)] = true
		}
	}

	return requiredFields
}

// formatStructTags formats Go struct tags.
func formatStructTags(obj map[string]interface{}, structId string, key string, tags []string) []string {
	requiredFields := getRequiredFields(obj, structId)

	tagList := []string{}
	for _, tag := range tags {
		// If this field is not required, then add omitempty to tag.
		if _, ok := requiredFields[key]; !ok {
			tagList = append(tagList, fmt.Sprintf("%s:\"%s,%s\"", tag, key, "omitempty"))
		} else {
			tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", tag, key))
		}
	}

	return tagList
}

// buildStructData loops through "properties" fields
// and constructs data for Go structs.
func buildStructData(prop interface{}, obj map[string]interface{}, structId string, tags []string, structData []string, modelMap map[string][]string) []string {
	for k, v := range prop.(map[string]interface{}) {
		valueName := FmtFieldName(k)
		tagList := formatStructTags(obj, structId, k, tags)

		if valueType, ok := v.(map[string]interface{})["type"]; ok && valueType != nil {
			switch value := valueType.(string); value {
			case "string":
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "bool":
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "integer":
				value = "int"
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "array":
				if ref, ok := v.(map[string]interface{})["items"].(map[string]interface{})["$ref"]; ok && ref != nil {
					objectType := generateModelTypes(obj, ref.(string), strings.Split(ref.(string), "_")[2], tags, modelMap)
					structData = append(structData, fmt.Sprintf("%s []%s `%s`", valueName, objectType, strings.Join(tagList, " ")))
				} else {
					// If there is no $ref - then we must be attempting to determine type without a ref
					obj[valueName] = v.(map[string]interface{})["items"]
					objectType := generateModelTypes(obj, valueName, valueName, tags, modelMap)
					structData = append(structData, fmt.Sprintf("%s []%s `%s`", valueName, objectType, strings.Join(tagList, " ")))
				}
			case "object":
				obj[valueName] = v
				objectType := generateModelTypes(obj, valueName, valueName, tags, modelMap)
				structData = append(structData, fmt.Sprintf("%s []%s `%s`", valueName, objectType, strings.Join(tagList, " ")))
			default:
				fmt.Printf("type not defined: %v", value)
			}
		} else if ref := v.(map[string]interface{})["$ref"]; ref != nil {
			objectType := generateModelTypes(obj, ref.(string), strings.Split(ref.(string), "_")[2], tags, modelMap)
			structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, objectType, strings.Join(tagList, " ")))
		} else {
			fmt.Printf("no type or ref for: %v", v)
		}
	}

	return structData
}

func generateModelTypes(obj map[string]interface{}, structId string, structName string, tags []string, modelMap map[string][]string) string {
	// use structId to search for existing data in modelMap
	if mapId, ok := modelMap[structId]; ok && mapId != nil {
		return modelMap[structId][0]
	}

	// If our data has a properties field - evaluate
	if checkPropertiesField(obj, structId) {
		properties := getPropertiesFieldValue(obj, structId)
		formattedStructName := []string{FmtFieldName(structName)}
		structData := buildStructData(properties, obj, structId, tags, formattedStructName, modelMap)
		modelMap[structId] = structData
		// else it may be the property itself and we need to get the type
	} else if objType, ok := obj[structId].(map[string]interface{})["type"]; ok && objType != nil {
		switch value := objType.(string); value {
		case "string":
			return "string"
		case "bool":
			return "bool"
		case "array":
			return "array"
		default:
			fmt.Printf("type not defined: %v", value)
		}
	} else {
		fmt.Println("did not find properties or type")
	}

	formattedStructName := FmtFieldName(structName)

	return formattedStructName
}

// generateOscalModelStruct generates the top-level struct for OSCAL data models.
func generateOscalModelStruct(oscalSchema map[string]interface{}, pkgName string, tags []string) string {
	oscalModelName := getOscalModel(oscalSchema)

	// Example: Format the string from 'oscal-model' to 'OscalModel'.
	formattedOscalModelName := FmtFieldName(oscalModelName)

	// Format struct tags.
	tagList := []string{}
	for _, tag := range tags {
		tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", tag, oscalModelName))
	}

	// Construct the struct string.
	structString := fmt.Sprintf("type Oscal%sModel struct {\n\t%s %s `%s`\n}\n", formattedOscalModelName, formattedOscalModelName, formattedOscalModelName, strings.Join(tagList, " "))

	// Format the Go struct.
	formattedStruct, err := format.Source([]byte(structString))
	if err != nil {
		fmt.Printf("error formatting:\n%s", structString)
		os.Exit(1)
	}

	return string(formattedStruct)
}

// handleDuplicateStructNames ensures we do not generate structs with duplicate struct names.
func handleDuplicateStructNames(existingValueMap map[string]bool, key, value string) string {
	var typesString string
	var i int = 0

	// If the struct name does not already exist, mark it as existing and generate a struct for it.
	if _, ok := existingValueMap[value]; !ok {
		existingValueMap[value] = true
		typesString += fmt.Sprintf("\ntype %s struct {", value)
	} else if strings.Contains(key, "#assembly") {
		// If the key for this value contains "#assembly" in the string,
		// let's use the key itself as the struct name.
		structName := strings.Trim(key, "#")
		formattedStructName := FmtFieldName(structName)
		typesString += fmt.Sprintf("\ntype %s struct {", formattedStructName)
	} else {
		// In this case, the key is identical to the value,
		// which means we need to add a unique identifier to the struct name.
		i++
		typesString += fmt.Sprintf("\ntype %s%v struct {", key, i)
	}

	return typesString
}

func generateStruct(structMap map[string][]string) string {
	var typesString string
	existingValueMap := map[string]bool{}

	// Begin generation of struct
	for k, v := range structMap {
		for index, value := range v {
			if index == 0 {
				typesString += handleDuplicateStructNames(existingValueMap, k, value)
			} else {
				typesString += fmt.Sprintf("\n\t%s", value)
			}
		}
		typesString += "\n}\n"
	}

	return typesString
}

// checkTopLevelRequiredField checks if an OSCAL schema file has a valid top-level 'required' key-value pair.
// If the key-value pair is valid, it returns true.
// If the key-value pair is invalid, it returns false.
func checkTopLevelRequiredField(oscalSchema map[string]interface{}) bool {
	if oscalModelField, ok := oscalSchema["required"]; ok && oscalModelField != nil {
		return true
	}

	if oscalModelField, ok := oscalSchema["required"]; !ok || oscalModelField == nil {
		return false
	}

	return false
}

// getRequiredFieldValue returns the value of the top-level 'required' field of an OSCAL schema file.
func getRequiredFieldValue(oscalSchema map[string]interface{}) interface{} {
	requiredFieldValue := oscalSchema["required"]
	return requiredFieldValue
}

// convertRequiredFieldInterfaceToString converts the value of the top-level 'required' field of an OSCAL schema file
// from type interface{} to type string and trims away the surrounding [] braces.
func convertRequiredFieldInterfaceToString(requiredFieldValue interface{}) string {
	oscalModelString := fmt.Sprintf("%v", requiredFieldValue)
	trimmedOscalModelString := strings.Trim(oscalModelString, "[]")

	return trimmedOscalModelString
}

// checkPropertiesField checks if an OSCAL schema file has a valid 'properties' key-value pair.
// If the key-value pair is valid, it returns true.
// If the key-value pair is invalid, it returns false.
func checkPropertiesField(oscalSchema map[string]interface{}, modelId string) bool {
	if properties, ok := oscalSchema[modelId].(map[string]interface{})["properties"]; ok && properties != nil {
		return true
	}

	if properties, ok := oscalSchema[modelId].(map[string]interface{})["properties"]; !ok || properties == nil {
		return false
	}

	return false
}

// getPropertiesFieldValue returns the value of the 'properties' field of an OSCAL schema file.
func getPropertiesFieldValue(oscalSchema map[string]interface{}, modelId string) interface{} {
	propertiesFieldValue := oscalSchema[modelId].(map[string]interface{})["properties"]
	return propertiesFieldValue
}

/*
	FmtFieldName formats a string as a struct key

Example:

	FmtFieldName("foo_id")

	Output: FooID
*/
func FmtFieldName(s string) string {
	runes := []rune(s)
	for len(runes) > 0 && !unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0]) {
		runes = runes[1:]
	}
	if len(runes) == 0 {
		return "_"
	}

	s = stringifyFirstChar(string(runes))
	name := lintFieldName(s)
	runes = []rune(name)
	for i, c := range runes {
		ok := unicode.IsLetter(c) || unicode.IsDigit(c)
		if i == 0 {
			ok = unicode.IsLetter(c)
		}
		if !ok {
			runes[i] = '_'
		}
	}
	s = string(runes)
	s = strings.Trim(s, "_")
	if len(s) == 0 {
		return "_"
	}
	result := ""
	for _, v := range strings.Split(s, "_") {
		runes := []rune(v)
		runes[0] = unicode.ToUpper(runes[0])
		result += string(runes)
	}
	return result
}

func lintFieldName(name string) string {
	// Fast path for simple cases: "_" and all lowercase.
	if name == "_" {
		return name
	}

	allLower := true
	for _, r := range name {
		if !unicode.IsLower(r) {
			allLower = false
			break
		}
	}
	if allLower {
		runes := []rune(name)
		if u := strings.ToUpper(name); commonInitialisms[u] {
			copy(runes[0:], []rune(u))
		} else {
			runes[0] = unicode.ToUpper(runes[0])
		}
		return string(runes)
	}

	allUpperWithUnderscore := true
	for _, r := range name {
		if !unicode.IsUpper(r) && r != '_' {
			allUpperWithUnderscore = false
			break
		}
	}
	if allUpperWithUnderscore {
		name = strings.ToLower(name)
	}

	// Split camelCase at any lower->upper transition, and split on underscores.
	// Check each word for common initialisms.
	runes := []rune(name)
	w, i := 0, 0 // index of start of word, scan
	for i+1 <= len(runes) {
		eow := false // whether we hit the end of a word

		if i+1 == len(runes) {
			eow = true
		} else if runes[i+1] == '_' {
			// underscore; shift the remainder forward over any run of underscores
			eow = true
			n := 1
			for i+n+1 < len(runes) && runes[i+n+1] == '_' {
				n++
			}

			// Leave at most one underscore if the underscore is between two digits
			if i+n+1 < len(runes) && unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i+n+1]) {
				n--
			}

			copy(runes[i+1:], runes[i+n+1:])
			runes = runes[:len(runes)-n]
		} else if unicode.IsLower(runes[i]) && !unicode.IsLower(runes[i+1]) {
			// lower -> non-lower
			eow = true
		}
		i++
		if !eow {
			continue
		}

		// [w,i) is a word.
		word := string(runes[w:i])
		if u := strings.ToUpper(word); commonInitialisms[u] {
			// All the common initialisms are ASCII,
			// so we can replace the bytes exactly.
			copy(runes[w:], []rune(u))

		} else if strings.ToLower(word) == word {
			// already all lowercase, and not the first word, so uppercase the first character.
			runes[w] = unicode.ToUpper(runes[w])
		}
		w = i
	}
	return string(runes)
}

// convert first character ints to strings
func stringifyFirstChar(str string) string {
	first := str[:1]

	i, err := strconv.ParseInt(first, 10, 8)

	if err != nil {
		return str
	}

	return intToWordMap[i] + "_" + str[1:]
}
