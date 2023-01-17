package oscal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

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

// Generate a struct definition for OSCAL component definition
func GenerateComponentDefinitionStructs(structName, pkgName string, tags []string, subStruct bool, convertFloats bool) ([]byte, error) {
	var componentDefinition = Properties_25{}
	var componentDefinitionMap map[string]interface{}

	var subStructMap map[string]string = nil

	// If --sub-struct flag is used, create string-string map
	if subStruct {
		subStructMap = make(map[string]string)
	}

	data, err := json.Marshal(componentDefinition)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &componentDefinitionMap); err != nil {
		return nil, err
	}

	src := fmt.Sprintf("package %s\ntype %s %s}",
		pkgName,
		structName,
		generateTypes(componentDefinitionMap, structName, tags, 0, subStructMap, convertFloats))

	// Append map keys to a string slice
	keys := make([]string, 0, len(subStructMap))
	for key := range subStructMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap[key], key)
	}

	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}

	return formatted, err
}

// Generate a struct definition given a JSON string representation of an object and a name structName.
func Generate(input io.Reader, parser Parser, structName, pkgName string, tags []string, subStruct bool, convertFloats bool) ([]byte, error) {
	var subStructMap map[string]string = nil

	// If --sub-struct flag is used, create string-string map
	if subStruct {
		subStructMap = make(map[string]string)
	}

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

	// Infer the data type of the interface slice and format the Go statements
	case []interface{}:
		src := fmt.Sprintf("package %s\n\ntype %s %s\n",
			pkgName,
			structName,
			structDataType(interfaceResult, structName, tags, subStructMap, convertFloats))

		formatted, err := format.Source([]byte(src))
		if err != nil {
			err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
		}
		return formatted, err

	default:
		return nil, fmt.Errorf("unexpected type: %T", interfaceResult)
	}

	src := fmt.Sprintf("package %s\ntype %s %s}",
		pkgName,
		structName,
		generateTypes(result, structName, tags, 0, subStructMap, convertFloats))

	// Append subStructMap keys to a string slice
	keys := make([]string, 0, len(subStructMap))
	for key := range subStructMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		src = fmt.Sprintf("%v\n\ntype %v %v", src, subStructMap[key], key)
	}

	formatted, err := format.Source([]byte(src))
	if err != nil {
		err = fmt.Errorf("error formatting: %s, was formatting\n%s", err, src)
	}

	return formatted, err
}

// convertKeysToStrings converts interface{} map keys to strings
func convertKeysToStrings(obj map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for key, value := range obj {
		res[fmt.Sprintf("%v", key)] = value
	}

	return res
}

// Generate go struct entries for a map[string]interface{} structure
func generateTypes(obj map[string]interface{}, structName string, tags []string, depth int, subStructMap map[string]string, convertFloats bool) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := obj[key]
		valueType := structDataType(value, structName, tags, subStructMap, convertFloats)

		// If a nested value, recurse
		switch value := value.(type) {
		// case []interface{}:
		// 	if len(value) > 0 {
		// 		sub := ""
		// 		if v, ok := value[0].(map[interface{}]interface{}); ok {
		// 			sub = generateTypes(convertKeysToStrings(v), structName, tags, depth+1, subStructMap, convertFloats) + "}"

		// 		} else if v, ok := value[0].(map[string]interface{}); ok {
		// 			sub = generateTypes(v, structName, tags, depth+1, subStructMap, convertFloats) + "}"
		// 		}

		// 		if sub != "" {
		// 			subName := sub

		// 			if subStructMap != nil {
		// 				if val, ok := subStructMap[sub]; ok {
		// 					subName = val

		// 				} else {
		// 					// TODO: dynamically grab/build the sub-struct name instead of building it as "OscalComponentDefinition_sub1"
		// 					// Reason for this is so that the types in this repo can be referenced by their field name,
		// 					// rather than the index number it was generated with.
		// 					// Example of how types would be referenced now:
		// 					// var componentDefinitions = &OscalComponentDefinition_sub108{}
		// 					structName := FmtFieldName(key)
		// 					subName = fmt.Sprintf("%v_%v", structName, len(subStructMap)+1)

		// 					subStructMap[sub] = subName
		// 				}
		// 			}

		// 			valueType = "[]" + subName
		// 		}
		// 	}

		// case map[interface{}]interface{}:
		// 	sub := generateTypes(convertKeysToStrings(value), structName, tags, depth+1, subStructMap, convertFloats) + "}"
		// 	subName := sub

		// 	if subStructMap != nil {
		// 		if val, ok := subStructMap[sub]; ok {
		// 			subName = val
		// 		} else {
		// 			// TODO: dynamically grab/build the sub-struct name instead of building it as "OscalComponentDefinition_sub1"
		// 			// Reason for this is so that the types in this repo can be referenced by their field name,
		// 			// rather than the index number it was generated with.
		// 			// Example of how types would be referenced now:
		// 			// var componentDefinitions = &OscalComponentDefinition_sub108{}
		// 			structName := FmtFieldName(key)
		// 			subName = fmt.Sprintf("%v_%v", structName, len(subStructMap)+1)

		// 			subStructMap[sub] = subName
		// 		}
		// 	}
		// 	valueType = subName

		case map[string]interface{}:
			sub := generateTypes(value, structName, tags, depth+1, subStructMap, convertFloats) + "}"
			subName := sub

			if subStructMap != nil {
				if val, ok := subStructMap[sub]; ok {
					subName = val
				} else {
					// TODO: dynamically grab/build the sub-struct name instead of building it as "OscalComponentDefinition_sub1"
					// Reason for this is so that the types in this repo can be referenced by their field name,
					// rather than the index number it was generated with.
					// Example of how types would be referenced now:
					// var componentDefinitions = &OscalComponentDefinition_sub108{}
					structName := FmtFieldName(key)
					subName = fmt.Sprintf("%v_%v", structName, len(subStructMap)+1)

					subStructMap[sub] = subName

				}
			}

			valueType = subName
		}

		fieldName := FmtFieldName(key)

		tagList := make([]string, 0)
		for _, t := range tags {
			tagList = append(tagList, fmt.Sprintf("%s:\"%s\"", t, key))
		}

		structure += fmt.Sprintf("\n%s %s `%s`",
			fieldName,
			valueType,
			strings.Join(tagList, " "))
	}

	return structure
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
	return s
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

// generate an appropriate struct data type entry
func structDataType(value interface{}, structName string, tags []string, subStructMap map[string]string, convertFloats bool) string {
	// Check if this is an array
	if objects, ok := value.([]interface{}); ok {
		types := make(map[reflect.Type]bool, 0)
		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}
		if len(types) == 1 {
			return "[]" + structDataType(mergeElements(objects).([]interface{})[0], structName, tags, subStructMap, convertFloats)
		}
		return "[]interface{}"

	} else if object, ok := value.(map[interface{}]interface{}); ok {
		return generateTypes(convertKeysToStrings(object), structName, tags, 0, subStructMap, convertFloats) + "}"

	} else if object, ok := value.(map[string]interface{}); ok {
		return generateTypes(object, structName, tags, 0, subStructMap, convertFloats) + "}"

	} else if reflect.TypeOf(value) == nil {
		return "interface{}"
	}
	typeName := reflect.TypeOf(value).Name()
	if typeName == "float64" && convertFloats {
		typeName = disambiguateFloatInt(value)
	}
	return typeName
}

// All numbers will initially be read as float64
// If the number appears to be an integer value, use int instead
func disambiguateFloatInt(value interface{}) string {
	const epsilon = .0001
	vfloat := value.(float64)

	if math.Abs(vfloat-math.Floor(vfloat+epsilon)) < epsilon {
		var tmp int64
		return reflect.TypeOf(tmp).Name()
	}

	return reflect.TypeOf(value).Name()
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

func mergeElements(i interface{}) interface{} {
	switch i := i.(type) {

	default:
		return i

	case []interface{}:
		length := len(i)
		if length == 0 {
			return i
		}
		for j := 1; j < length; j++ {
			i[0] = mergeObjects(i[0], i[j])
		}
		return i[0:1]
	}
}

func mergeObjects(o1, o2 interface{}) interface{} {
	if o1 == nil {
		return o2
	}

	if o2 == nil {
		return o1
	}

	if reflect.TypeOf(o1) != reflect.TypeOf(o2) {
		return nil
	}

	switch i := o1.(type) {

	default:
		return o1

	// case []interface{}:
	// 	if i2, ok := o2.([]interface{}); ok {
	// 		i3 := append(i, i2...)
	// 		return mergeElements(i3)
	// 	}
	// 	return mergeElements(i)

	case map[string]interface{}:
		if i2, ok := o2.(map[string]interface{}); ok {
			for k, v := range i2 {
				if v2, ok := i[k]; ok {
					i[k] = mergeObjects(v2, v)
				} else {
					i[k] = v
				}
			}
		}
		return i

		// case map[interface{}]interface{}:
		// 	if i2, ok := o2.(map[interface{}]interface{}); ok {
		// 		for k, v := range i2 {
		// 			if v2, ok := i[k]; ok {
		// 				i[k] = mergeObjects(v2, v)
		// 			} else {
		// 				i[k] = v
		// 			}
		// 		}
		// 	}
		// 	return i
	}
}
