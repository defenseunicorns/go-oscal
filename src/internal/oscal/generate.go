package oscal

import (
	"fmt"
	"go/format"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/swaggest/jsonschema-go"
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
func Generate(oscalSchema []byte, pkgName string, tags []string) ([]byte, error) {

	// unmarshall to the schema object
	schema := jsonschema.Schema{}
	schema.UnmarshalJSON(oscalSchema)

	// Identify the OSCAL model under generation
	model, err := getOscalModel(schema.Required)
	if err != nil {
		return nil, err
	}

	// Get the $id for the OSCAL model under generation
	modelId, err := setOscalModelRef(model)
	if err != nil {
		return nil, err
	}

	// Generate a map with unique Id as key and existing schemaOrBool as value
	idMap, err := generateUniqueIdMap(schema)
	if err != nil {
		return nil, err
	}

	// Instantiate variable for storing the data
	modelTypeMap := make(map[string][]string)

	generateModelTypes(idMap, modelId, strings.Split(modelId, "_")[2], tags, modelTypeMap)

	// Construct header comment and package name.
	structString := fmt.Sprintf("%s\n\npackage %s\n", headerComment, pkgName)

	// Construct top-level OscalModel struct.
	toplevel, err := generateOscalModelStruct(model, pkgName, tags)
	if err != nil {
		return nil, err
	}
	structString += toplevel

	// Construct structs for oscal models.
	structString += generateStruct(modelTypeMap)

	formattedStruct, err := format.Source([]byte(structString))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, structString)
		return nil, err
	}

	return formattedStruct, nil
}

// setOscalModelRef determines which OSCAL model $ref to use based on the model name.
// TODO: Look at removing this required translation to allow generation of all OSCAL models and de-duplication
func setOscalModelRef(oscalModel string) (string, error) {
	// Check which OSCAL model we're working with, and set the $ref accordingly.
	if oscalModel == "system-security-plan" {
		return "#assembly_oscal-ssp_system-security-plan", nil
	}

	// Infinite loop problem - Investigate
	if oscalModel == "assessment-plan" {
		return "", fmt.Errorf("Unsupported OSCAL model. Currently supported OSCAL models are Component Definition and System Security Plan.")
	}

	// Infinite loop problem - Investigate
	if oscalModel == "assessment-results" {
		return "#assembly_oscal-ar_assessment-results", nil
	}

	// Infinite loop problem - Investigate
	if oscalModel == "plan-of-action-and-milestones" {
		return "", fmt.Errorf("Unsupported OSCAL model. Currently supported OSCAL models are Component Definition and System Security Plan.")
	}

	if oscalModel == "component-definition" {
		return "#assembly_oscal-component-definition_component-definition", nil
	}

	return "", fmt.Errorf("Unsupported OSCAL model. Currently supported OSCAL models are Component Definition and System Security Plan.")
}

// getOscalModel determines which OSCAL model we're working with.
func getOscalModel(required []string) (string, error) {
	// error if more than one required field - not supported
	if len(required) > 1 {
		return "", fmt.Errorf("Processing more than one model is unsupported")
	}
	if len(required) == 1 {
		return required[0], nil
	}

	return "", fmt.Errorf("Top-level 'required' field not found or is not populated. Please verify the OSCAL JSON schema file is valid.")
}

// TODO: Look into full definition naming when processing multiple schemas - unsupported at this time
// create a map with $id primary keys and the object values
func generateUniqueIdMap(schema jsonschema.Schema) (map[string]jsonschema.SchemaOrBool, error) {
	result := make(map[string]jsonschema.SchemaOrBool)

	for definition, item := range schema.Definitions {
		if strings.Contains(definition, "Datatype") {
			result["#/definitions/"+definition] = item
		} else if definition == "json-schema-definition" {
			// Skipping - TODO: determine relevance in the golang structs
			continue
		} else {
			result[*item.TypeObject.ID] = item
		}

	}

	return result, nil
}

// getRequiredFields collects the required fields in each OSCAL model definition.
func getRequiredFields(obj map[string]jsonschema.SchemaOrBool, structId string) map[string]bool {
	requiredFields := make(map[string]bool)

	if required := obj[structId].TypeObject.Required; required != nil {
		for _, v := range required {
			requiredFields[v] = true
		}
	}

	return requiredFields
}

// formatStructTags formats Go struct tags.
func formatStructTags(obj map[string]jsonschema.SchemaOrBool, structId string, key string, tags []string) []string {
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
func buildStructData(prop map[string]jsonschema.SchemaOrBool, obj map[string]jsonschema.SchemaOrBool, structId string, tags []string, structData []string, modelMap map[string][]string) ([]string, error) {
	keys := make([]string, 0, len(prop))

	for k := range prop {
		keys = append(keys, k)
	}

	// sort the string array
	sort.Strings(keys)

	for _, key := range keys {

		simple, err := prop[key].ToSimpleMap()
		if err != nil {
			return nil, err
		}
		valueName := FmtFieldName(key)
		tagList := formatStructTags(obj, structId, key, tags)

		if prop[key].TypeObject.Type != nil {
			switch value := simple["type"].(string); value {
			case "string":
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "bool":
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "integer":
				value = "int"
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, value, strings.Join(tagList, " ")))
			case "array":
				// If type array and ref is populated
				if ref := prop[key].TypeObject.Items.SchemaOrBool.TypeObject.Ref; ref != nil {
					var objectType string
					if strings.Contains(*ref, "Datatype") {
						objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "/")[2], tags, modelMap)
						if err != nil {
							return nil, err
						}
					} else {
						objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "_")[2], tags, modelMap)
						if err != nil {
							return nil, err
						}
					}

					structData = append(structData, fmt.Sprintf("%s []%s `%s`", valueName, objectType, strings.Join(tagList, " ")))
				} else {
					// TODO: would anyof/allof ever be evaluated here?

					// if array and no ref populated - check for items
					if items := prop[key].TypeObject.Items.SchemaOrBool; items != nil {
						singular := *prop[key].TypeObject.Items.SchemaOrBool.TypeObject.Title
						obj[valueName] = *prop[key].TypeObject.Items.SchemaOrBool
						objectType, err := generateModelTypes(obj, valueName, singular, tags, modelMap)
						if err != nil {
							return nil, err
						}
						structData = append(structData, fmt.Sprintf("%s []%s `%s`", valueName, objectType, strings.Join(tagList, " ")))
					} else {
						// TODO: Error Handling
						fmt.Println("Did not find ref/items/allof/anyof")
					}

				}
			case "object":
				obj[valueName] = prop[key]
				objectType, err := generateModelTypes(obj, valueName, valueName, tags, modelMap)
				if err != nil {
					return nil, err
				}
				structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, objectType, strings.Join(tagList, " ")))
			default:
				// TODO: Error handling
				fmt.Printf("type not defined: %v", value)
			}
		} else if ref := prop[key].TypeObject.Ref; ref != nil {
			var objectType string
			if strings.Contains(*ref, "Datatype") {
				objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "/")[2], tags, modelMap)
				if err != nil {
					return nil, err
				}
			} else {
				objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "_")[2], tags, modelMap)
				if err != nil {
					return nil, err
				}
			}
			structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, objectType, strings.Join(tagList, " ")))
		} else if anyof := prop[key].TypeObject.AnyOf; anyof != nil {
			// TODO: support OR operation and enums - for now we will locate the ref
			for _, schema := range prop[key].TypeObject.AnyOf {
				if ref := schema.TypeObject.Ref; ref != nil {
					var objectType string
					if strings.Contains(*ref, "Datatype") {
						objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "/")[2], tags, modelMap)
						if err != nil {
							return nil, err
						}
					} else {
						objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "_")[2], tags, modelMap)
						if err != nil {
							return nil, err
						}
					}
					structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, objectType, strings.Join(tagList, " ")))
				}
			}

		} else if allof := prop[key].TypeObject.AllOf; allof != nil {

			for _, schema := range prop[key].TypeObject.AllOf {
				// TODO: support AND operation and enums - for now we will locate the ref
				if ref := schema.TypeObject.Ref; ref != nil {
					var objectType string
					if strings.Contains(*ref, "Datatype") {
						objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "/")[2], tags, modelMap)
						if err != nil {
							return nil, err
						}
					} else {
						objectType, err = generateModelTypes(obj, *ref, strings.Split(*ref, "_")[2], tags, modelMap)
						if err != nil {
							return nil, err
						}
					}
					structData = append(structData, fmt.Sprintf("%s %s `%s`", valueName, objectType, strings.Join(tagList, " ")))
				}
			}
		} else {
			// TODO: Error Handling
			fmt.Printf("no type or ref for: %s\n", key)

		}
	}

	return structData, nil
}

func generateModelTypes(obj map[string]jsonschema.SchemaOrBool, structId string, structName string, tags []string, modelMap map[string][]string) (string, error) {
	// use structId to search for existing data in modelMap
	if mapId, ok := modelMap[structId]; ok && mapId != nil {
		return modelMap[structId][0], nil
	}

	// If the object has properties
	if obj[structId].TypeObject.Properties != nil {
		properties := obj[structId].TypeObject.Properties
		formattedStructName := []string{FmtFieldName(structName)}
		modelMap[structId] = formattedStructName
		structData, err := buildStructData(properties, obj, structId, tags, formattedStructName, modelMap)
		if err != nil {
			return "", err
		}
		modelMap[structId] = structData
	} else if objType := obj[structId].TypeObject.Type; objType != nil {
		simple, err := obj[structId].ToSimpleMap()
		if err != nil {
			return "", err
		}
		switch value := simple["type"].(string); value {
		case "string":
			return "string", nil
		case "bool":
			return "bool", nil
		case "array":
			return "array", nil
		case "object":
			// Edge case here for empty objects (ie `include-all`)
			formattedStructName := []string{FmtFieldName(structName)}
			modelMap[structId] = formattedStructName
			return formattedStructName[0], nil
		default:
			//TODO: Error Handling
			fmt.Printf("type not defined: %v", value)
		}
	} else if ref := obj[structId].TypeObject.Ref; ref != nil {
		switch *ref {
		case "#/definitions/Base64Datatype":
			return "string", nil
		case "#/definitions/DateTimeWithTimezoneDatatype":
			return "string", nil
		case "#/definitions/EmailAddressDatatype":
			return "string", nil
		case "#/definitions/IntegerDatatype":
			return "int", nil
		case "#/definitions/NonNegativeIntegerDatatype":
			return "int", nil
		case "#/definitions/StringDatatype":
			return "string", nil
		case "#/definitions/TokenDatatype":
			return "string", nil
		case "#/definitions/URIDatatype":
			return "string", nil
		case "#/definitions/URIReferenceDatatype":
			return "string", nil
		case "#/definitions/UUIDDatatype":
			return "string", nil
		}

		reference, err := obj[*ref].ToSimpleMap()
		if err != nil {
			return "", err
		}

		switch value := reference["type"].(string); value {
		case "string":
			return "string", nil
		case "bool":
			return "bool", nil
		case "array":
			return "array", nil
		default:
			fmt.Printf("type not defined: %v", value)
		}

	} else if allof := obj[structId].TypeObject.AllOf; allof != nil {
		for _, schema := range allof {
			// TODO: support AND operation and enums - for now we will locate the ref
			if ref := schema.TypeObject.Ref; ref != nil {
				switch *ref {
				case "#/definitions/Base64Datatype":
					return "string", nil
				case "#/definitions/DateTimeWithTimezoneDatatype":
					return "string", nil
				case "#/definitions/EmailAddressDatatype":
					return "string", nil
				case "#/definitions/IntegerDatatype":
					return "int", nil
				case "#/definitions/NonNegativeIntegerDatatype":
					return "int", nil
				case "#/definitions/StringDatatype":
					return "string", nil
				case "#/definitions/TokenDatatype":
					return "string", nil
				case "#/definitions/URIDatatype":
					return "string", nil
				case "#/definitions/URIReferenceDatatype":
					return "string", nil
				case "#/definitions/UUIDDatatype":
					return "string", nil
				}
			}
		}
	} else if anyof := obj[structId].TypeObject.AnyOf; anyof != nil {
		for _, schema := range anyof {
			// TODO: support AND operation and enums - for now we will locate the ref
			if ref := schema.TypeObject.Ref; ref != nil {
				switch *ref {
				case "#/definitions/Base64Datatype":
					return "string", nil
				case "#/definitions/DateTimeWithTimezoneDatatype":
					return "string", nil
				case "#/definitions/EmailAddressDatatype":
					return "string", nil
				case "#/definitions/IntegerDatatype":
					return "int", nil
				case "#/definitions/NonNegativeIntegerDatatype":
					return "int", nil
				case "#/definitions/StringDatatype":
					return "string", nil
				case "#/definitions/TokenDatatype":
					return "string", nil
				case "#/definitions/URIDatatype":
					return "string", nil
				case "#/definitions/URIReferenceDatatype":
					return "string", nil
				case "#/definitions/UUIDDatatype":
					return "string", nil
				}
			}
		}
	} else {
		// TODO: Error Handling
		fmt.Println("did not find properties or type")
		fmt.Printf("structId: %s\n", structId)
	}

	formattedStructName := FmtFieldName(structName)

	return formattedStructName, nil
}

// generateOscalModelStruct generates the top-level struct for OSCAL data models.
func generateOscalModelStruct(oscalModel string, pkgName string, tags []string) (string, error) {

	// Example: Format the string from 'oscal-model' to 'OscalModel'.
	formattedOscalModelName := FmtFieldName(oscalModel)

	// Format struct tags.
	tagList := []string{}
	for _, tag := range tags {
		tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", tag, oscalModel))
	}

	// Construct the struct string.
	structString := fmt.Sprintf("type Oscal%sModel struct {\n\t%s %s `%s`\n}\n", formattedOscalModelName, formattedOscalModelName, formattedOscalModelName, strings.Join(tagList, " "))

	// Format the Go struct.
	formattedStruct, err := format.Source([]byte(structString))
	if err != nil {
		return "", fmt.Errorf("error formatting:\n%s", structString)
	}

	return string(formattedStruct), nil
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

	// create string array and populate it from the map
	keys := make([]string, 0, len(structMap))

	for k := range structMap {
		keys = append(keys, k)
	}

	// sort the string array
	sort.Strings(keys)

	// Begin generation of the structs
	for _, k := range keys {
		for index, value := range structMap[k] {
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
