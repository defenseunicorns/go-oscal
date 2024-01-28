package oscal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/swaggest/jsonschema-go"
)

var RefsToIgnore map[string]bool = map[string]bool{
	"#json-schema-directive": true,
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

func getRef(schema jsonschema.Schema) string {
	if schema.Ref != nil {
		return *schema.Ref
	} else if schema.ID != nil {
		return *schema.ID
	} else if schema.Title != nil {
		split := strings.Split(*schema.Title, " ")
		name := ""
		for _, s := range split {
			name += FmtFieldName(s)
		}
		return "#/definitions/" + name
	}
	return ""
}

// returns the type of the schema
func getSimpleType(schema jsonschema.Schema) string {
	if schema.Type != nil {
		return string(*schema.Type.SimpleTypes)
	}
	return ""
}

func isArrayType(t string) bool {
	lower := strings.ToLower(t)
	return lower == "array"
}

func isObjectType(t string) bool {
	lower := strings.ToLower(t)
	return lower == "object"
}

func isPrimitiveType(t string) bool {
	lower := strings.ToLower(t)
	switch lower {
	case "string":
		return true
	case "boolean":
		return true
	case "number":
		return true
	case "integer":
		return true
	}
	return false
}

func getGoType(t string) string {
	lower := strings.ToLower(t)
	switch lower {
	case "string":
		return "string"
	case "boolean":
		return "bool"
	case "number":
		return "float64"
	case "integer":
		return "int"
	case "array":
		return "[]"
	}
	return ""
}

func populateSchema(oscalSchema []byte) (jsonschema.Schema, error) {
	schema := jsonschema.Schema{}
	err := schema.UnmarshalJSON(oscalSchema)
	return schema, err
}

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
		if typeObj := *item.TypeObject; typeObj.ID != nil {
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
