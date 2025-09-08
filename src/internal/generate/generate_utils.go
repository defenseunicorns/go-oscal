package generate

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/swaggest/jsonschema-go"
)

var Imports []string = []string{"time"}

var GoPrimitive map[string]bool = map[string]bool{
	"string":  true,
	"bool":    true,
	"float64": true,
	"int":     true,
}

var PrimitiveAndCustomTypes map[string]string = map[string]string{
	"string":    "string",
	"boolean":   "bool",
	"number":    "float64",
	"integer":   "int",
	"date-time": "time.Time",
}

var ComplexTypes map[string]string = map[string]string{
	"array": "[]",
}

var RefsToIgnore map[string]bool = map[string]bool{
	"#json-schema-directive": true,
}

var KeysToIgnore map[string]bool = map[string]bool{
	"$schema": true,
}

var Aliases map[string][]string = map[string][]string{
	"OscalCompleteSchema": {"OscalModels"},
}

const headerComment string = `/*
	This file was auto-generated with go-oscal.

	To regenerate:
	
	go-oscal generate \
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

// addPointerIfOptionalNonPrimitive adds a pointer to the type if the field is optional
func addPointerIfOptionalNonPrimitive(required bool, t string) string {

	if required || GoPrimitive[t] {
		return t
	}
	return "*" + t
}

// buildStructs builds the structs for each definition in the schema
func buildTagString(tags []string, field string, required bool) string {
	tagStrings := []string{}
	tagSuffix := ",omitempty"

	// No tags, return empty string
	if len(tags) == 0 {
		return ""
	}

	// If required, remove omitempty
	if required {
		tagSuffix = ""
	}

	// Build tag string for each tag
	for _, tag := range tags {
		tagStrings = append(tagStrings, fmt.Sprintf("%s:\"%s%s\"", tag, field, tagSuffix))
	}
	return "`" + strings.Join(tagStrings, " ") + "`"
}

// getRef builds a ref string from a schema
func getRef(schema jsonschema.Schema) (string, error) {
	if schema.Ref != nil {
		return *schema.Ref, nil
	} else if schema.ID != nil {
		return *schema.ID, nil
	} else if schema.Title != nil {
		return getRefWithName(getNameFromTitle(*schema.Title)), nil
	}
	return "", fmt.Errorf("unable to get ref from schema")
}

func getNameFromTitle(title string) string {
	split := strings.Split(title, " ")
	name := ""
	for _, s := range split {
		name += FmtFieldName(s)
	}
	return name
}

func getRefWithName(name string) string {
	return "#/definitions/" + name
}

// getCustomTypeKey returns the custom type key if the schema has a custom type
func getCustomTypeKey(schema jsonschema.Schema) string {
	// If the schema has a format, check if it's a custom type
	if schema.Format != nil {
		if _, ok := PrimitiveAndCustomTypes[*schema.Format]; ok {
			return *schema.Format
		}
	}
	return ""
}

// getJsonOrCustomType returns the json type of the schema
func getJsonOrCustomType(schema jsonschema.Schema) string {
	// if the schema has a custom type, return the custom type
	if importType := getCustomTypeKey(schema); importType != "" {
		return importType
	}
	if schema.Type != nil {
		return string(*schema.Type.SimpleTypes)
	}
	return ""
}

// isPrimitiveOrCustomJsonType returns true if the type is a primitive type or a custom type
func isPrimitiveOrCustomJsonType(t string) bool {
	lower := strings.ToLower(t)
	return PrimitiveAndCustomTypes[lower] != ""
}

// getGoType returns the Go type for a given JSON type
func getGoType(t string) string {
	lower := strings.ToLower(t)

	if PrimitiveAndCustomTypes[lower] != "" {
		return PrimitiveAndCustomTypes[lower]
	}

	return ComplexTypes[lower]
}

// populateSchema unmarshals the OSCAL JSON schema file into a jsonschema.Schema object
func populateSchema(oscalSchema []byte) (jsonschema.Schema, error) {
	schema := jsonschema.Schema{}
	err := schema.UnmarshalJSON(oscalSchema)
	return schema, err
}

// getNameFromRef returns the name of the struct from a ref
func getNameFromRef(ref string) string {
	pattern := regexp.MustCompile("[/_]")
	splitPathSeperator := pattern.Split(ref, -1)
	splitFileExt := strings.Split(splitPathSeperator[len(splitPathSeperator)-1], ".")
	return FmtFieldName(splitFileExt[0])
}

// getDefinitionMap generates a map of definitions from the OSCAL JSON schema file.
// The key is the $id or $ref, and the value is the schemaOrBool object.
func getDefinitionMap(schema jsonschema.Schema) (map[string]jsonschema.Schema, error) {
	result := make(map[string]jsonschema.Schema)
	for definition, item := range schema.Definitions {
		typeObj := *item.TypeObject

		// If the object has no Parent, set the parent to the schema
		if typeObj.Parent == nil {
			typeObj.Parent = &schema
		}

		// If the object has an ID, use that as the key, otherwise use the definition name
		if typeObj.ID != nil {
			result[*item.TypeObject.ID] = typeObj
		} else {
			result["#/definitions/"+definition] = typeObj
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no definitions found in the OSCAL JSON schema file, please verify the OSCAL JSON schema file is valid")
	}
	return result, nil
}

/*
	FmtFieldName formats a string as a struct key

Example:

	FmtFieldName("foo_id")

	Output: FooID
*/
func FmtFieldName(s string) string {

	// Handle definitions that include #/definitions/<modeltype>:<fieldname>
	if strings.Contains(s, ":") {
		s = strings.Split(s, ":")[1]
	}

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

// buildImportString returns a string of imports for the Go file
func buildImportString() string {
	imports := "import (\n"
	for _, imp := range Imports {
		imports += fmt.Sprintf("\t\"%s\"\n", imp)
	}
	imports += ")\n"
	return imports
}

// extractVersion parses a URL/string and returns the first x.y.z version it finds.
func extractVersion(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	// Match MAJOR.MINOR.PATCH (e.g., 1.1.3)
	re := regexp.MustCompile(`\b\d+\.\d+\.\d+\b`)

	// Search the path first; fall back to the whole string just in case.
	if m := re.FindString(u.Path); m != "" {
		return m, nil
	}
	if m := re.FindString(raw); m != "" {
		return m, nil
	}

	return "", errors.New("version not found")
}
