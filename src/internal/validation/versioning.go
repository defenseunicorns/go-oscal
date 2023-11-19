package validation

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

//go:embed schema/*.json
var schemas embed.FS

const SCHEMA_PREFIX = "oscal_complete_schema-"
const DEFAULT_OSCAL_VERSION = "1.0.4"

var versionRegexp = regexp.MustCompile(`^\d+([-\.]\d+){2}$`)

var supportedVersion = map[string]bool{
	"1.0.4": true,
	"1.0.5": true,
	"1.0.6": true,
	"1.1.0": true,
	"1.1.1": true,
}

// CoerceToJSONForTypeSafety takes a yaml byte array and coerces it to a json interface{}
// This is necessary because the jsonschema library does not support yaml date types that are not declared as strings ie "2021-01-01" vs 2021-01-01
func CoerceToJSONForTypeSafety[T interface{} | []byte](version string, ymlData T) (model interface{}, err error) {
	ymlBytes, ok := reflect.ValueOf(ymlData).Interface().([]byte)
	if ok {
		err = yaml.Unmarshal(ymlBytes, &model)
		if err != nil {
			return nil, err
		}
	} else {
		model = ymlData
	}
	jsonBytes, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func IsValidSchemaVersion[T interface{} | []byte](version string, docBytes T) bool {
	component, err := CoerceToJSONForTypeSafety(version, docBytes)
	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}

	compiler := jsonschema.NewCompiler()
	schemaPath := SCHEMA_PREFIX + strings.ReplaceAll(version, ".", "-") + ".json"
	schemaBytes, err := schemas.ReadFile("schema/" + schemaPath)

	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}
	compiler.AddResource(schemaPath, strings.NewReader(string(schemaBytes)))
	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}
	sch, err := compiler.Compile(schemaPath)
	if err != nil {
		log.Printf("%#v\n", err)
		return false
	}
	err = sch.Validate(component)
	if err != nil {
		log.Printf("%#v\n", err)
		return false
		// b, _ := json.MarshalIndent(err.(*jsonschema.ValidationError).DetailedOutput(), "", "  ")
		// fmt.Println(string(b))
		// return false
	}
	return true
}

// GetVersion returns formatted OSCAL version if valid version is passed, returns error if not.
func GetVersion(userVersion string) (string, error) {
	if userVersion == "" {
		return DEFAULT_OSCAL_VERSION, nil
	}
	builtVersion := formatUserVersion(userVersion)

	if !isVersionRegexp(builtVersion) {
		return builtVersion, fmt.Errorf("version %s is not a valid version", userVersion)
	}

	if !supportedVersion[builtVersion] {
		return builtVersion, fmt.Errorf("version %s is not supported", userVersion)
	}

	return builtVersion, nil
}

func isVersionRegexp(v string) bool {
	return versionRegexp.MatchString(v)
}

func formatUserVersion(v string) string {
	builtVersion := v
	if builtVersion[0] == 'v' {
		builtVersion = builtVersion[1:]
	}
	builtVersion = strings.ReplaceAll(builtVersion, "-", ".")
	return builtVersion
}
