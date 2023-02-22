package oscal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

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

type Parser func(io.Reader) (interface{}, error)

// ParseJson reads encoded JSON as input, stores to an interface pointer, and returns the interface.
func ParseJson(input io.Reader) (interface{}, error) {
	var result interface{}
	if err := json.NewDecoder(input).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ParseYaml reads encoded YAML as input, stores to an interface pointer, and returns the interface.
func ParseYaml(input io.Reader) (interface{}, error) {
	var result interface{}
	b, err := readFile(input)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// readFile stores the input data as a buffer and returns the buffer's bytes.
func readFile(input io.Reader) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, input)
	if err != nil {
		return []byte{}, nil
	}
	return buf.Bytes(), nil
}

// Generate a struct definition given a JSON string representation of an object.
func Generate(input io.Reader, parser Parser, pkgName string, tags []string, convertFloats bool) ([]byte, error) {
	var result map[string]interface{}

	interfaceResult, err := parser(input)
	if err != nil {
		return nil, err
	}

	// Parse input interface and perform necessary data transformations
	switch interfaceResultType := interfaceResult.(type) {

	// Convert interface map keys to string map keys
	case map[interface{}]interface{}:
		result = convertKeysToStrings(interfaceResultType)

	// No transformations, this is what we want
	case map[string]interface{}:
		result = interfaceResultType

	default:
		return nil, fmt.Errorf("unexpected type: %T", interfaceResult)
	}

	id := setOscalModelRef(result)

	// Generate a map with unique Id as key and existing interface as value
	idMap := generateUniqueIdMap(result)

	// Instantiate variable for storing the data
	modelTypeMap := make(map[string][]string)

	generateModelTypes(idMap, id, strings.Split(id, "_")[2], tags, modelTypeMap)

	// Construct header comment and package name.
	structString := fmt.Sprintf("%s\n\npackage %s\n", headerComment, pkgName)

	// Construct top-level 'OscalModels' struct.
	structString += generateOscalModelStruct(result, pkgName, tags)

	// Construct structs for oscal models.
	structString += generateStruct(modelTypeMap, pkgName)

	formatted, err := format.Source([]byte(structString))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, structString)
	}

	return formatted, err
}

// getOscalModel determines which OSCAL model we're working with.
func getOscalModel(oscalSchema map[string]interface{}) (oscalModel string) {
	if oscalModelField := oscalSchema["required"]; oscalModelField != nil {
		oscalModelString := fmt.Sprintf("%v", oscalModelField)

		// Trim the [] braces away.
		oscalModel = strings.Trim(oscalModelString, "[]")

		return
	} else {
		fmt.Println("Top-level required field is not populated as expected, please verify the provided schema is valid.")
		os.Exit(1)
	}

	return ""
}

// setOscalModelRef determines which OSCAL model $ref to use based on the model name.
func setOscalModelRef(oscalSchema map[string]interface{}) (oscalModelRef string) {
	oscalModel := getOscalModel(oscalSchema)

	// Check which OSCAL model we're working with, and set the $ref accordingly.
	switch oscalModel {
	case "system-security-plan":
		oscalModelRef = "#assembly_oscal-ssp_system-security-plan"
		return
	case "component-definition":
		oscalModelRef = "#assembly_oscal-component-definition_component-definition"
		return
	default:
		fmt.Println("Do not recognize this model name.")
	}

	return ""
}

func generateUniqueIdMap(obj map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, val := range obj["definitions"].(map[string]interface{}) {
		result[val.(map[string]interface{})["$id"].(string)] = val
	}

	return result

}

// convertKeysToStrings converts interface{} map keys to strings.
func convertKeysToStrings(obj map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for key, value := range obj {
		res[fmt.Sprintf("%v", key)] = value
	}

	return res
}

// getRequiredFields collects the required fields in each OSCAL definition.
func getRequiredFields(obj map[string]interface{}, structId string) map[string]bool {
	requiredFields := make(map[string]bool)

	if required := obj[structId].(map[string]interface{})["required"]; required != nil {
		for _, v := range required.([]interface{}) {
			requiredFields[v.(string)] = true
		}
	}

	return requiredFields
}

// formatStructTags formats Go struct tags.
func formatStructTags(obj map[string]interface{}, structId string, key string, tags []string) []string {
	requiredFields := getRequiredFields(obj, structId)

	tagList := make([]string, 0)

	for _, tag := range tags {
		// If this field is not required, then add omitempty to tag.
		if _, required := requiredFields[key]; !required {
			omitEmpty := ",omitempty"
			tagList = append(tagList, fmt.Sprintf("%s:\"%s%s\"", tag, key, omitEmpty))
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

		if valueType := v.(map[string]interface{})["type"]; valueType != nil {
			switch value := valueType.(string); value {
			case "string":
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "bool":
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "integer":
				value = "int"
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "array":
				if ref := v.(map[string]interface{})["items"].(map[string]interface{})["$ref"]; ref != nil {
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
	if existing := modelMap[structId]; existing != nil {
		return modelMap[structId][0]
	}

	// If our data has a properties field - evaluate
	// else it may be the property itself and we need to get the type
	if prop := obj[structId].(map[string]interface{})["properties"]; prop != nil {
		structData := []string{FmtFieldName(structName)}
		structData = buildStructData(prop, obj, structId, tags, structData, modelMap)
		modelMap[structId] = structData

	} else if objType := obj[structId].(map[string]interface{})["type"]; objType != nil {
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

// TODO: Make this function extensible to handle multiple OSCAL	models/schemas. Namely, system security plan and assesment plan.
// generateOscalModelStruct generates the top-level struct for OSCAL data models.
func generateOscalModelStruct(oscalSchema map[string]interface{}, pkgName string, tags []string) string {
	// Check if there is a top-level 'required' field.
	if oscalModelField := oscalSchema["required"]; oscalModelField != nil {
		// There is, so let's convert it to a string.
		oscalModelString := fmt.Sprintf("%v", oscalModelField)

		// Trim the [] braces away.
		oscalModelName := strings.Trim(oscalModelString, "[]")

		// Format the string from 'component-definition' to 'ComponentDefinition'.
		formattedOscalModelName := FmtFieldName(oscalModelName)

		// Format struct tags.
		tagList := make([]string, 0)
		for _, tag := range tags {
			tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", tag, oscalModelName))
		}

		// Construct the struct string.
		structString := fmt.Sprintf("type OscalModel struct {\n\t%s %s `%s`\n}\n", formattedOscalModelName, formattedOscalModelName, strings.Join(tagList, " "))

		// Format the Go struct.
		formattedStruct, err := format.Source([]byte(structString))
		if err != nil {
			fmt.Printf("error formatting:\n%s", structString)
			os.Exit(1)
		}

		return string(formattedStruct)

	} else {
		fmt.Println("Top-level 'required' field not found. Please verify the OSCAL JSON schema file is valid.")
		os.Exit(1)
	}

	return ""
}

func generateStruct(structMap map[string][]string, pkgName string) string {
	existing := make(map[string]bool)
	var typesString string

	// Begin generation of struct
	for i, v := range structMap {

		for index, value := range v {
			if index == 0 {
				_, ok := existing[value]
				if !ok {
					typesString += fmt.Sprintf("\ntype %v struct {", value)
				} else {
					// Handle duplicate struct names
					name := strings.Split(i, "_")
					typesString += fmt.Sprintf("\ntype %v struct {", strings.Trim(name[0], "#")+"_"+value)
				}
			} else {
				typesString += fmt.Sprintf("\n\t%s", value)
			}
		}
		typesString += fmt.Sprintf("\n}\n")
	}

	return typesString
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
