package oscal

import (
	"fmt"
	"reflect"
	// "log"

	compdef104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4/component-definition"
	compdef105 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-5/component-definition"

	// compdef106 "github.com/defeunseunicorns/go-oscal/src/types/oscal-1-0-6/component-definition"
	// compdef110 "github.com/defeunseunicorns/go-oscal/src/types/oscal-1-1-0/component-definition"
	// compdef111 "github.com/defeunseunicorns/go-oscal/src/types/oscal-1-1-1/component-definition"
	"gopkg.in/yaml.v3"
)

/*
	Conversion logic for component-definition model documents.
	The primary idea here is that we can write logic to perform the conversion of an older oscal version to a newer by iterating through versions.
	Instead of every possible mutation, we opinionate and perform an upgrade to the next version.
	If a document is upgrading across multiple versions, we perform a series of conversions.

	we can look at reduction of logic and complexity after proving out the process.
*/

// Convert accepts a slice of bytes and returns an object representing the oscal document.
func ConvertComponentDefinition(data []byte, currentVersion string, targetVersion string) (any, error) {
	// object placeholders for original oscal and new oscal
	// var original, new any

	// // create list of valid versions
	// validVersions := []string{"1.0.4", "1.0.5", "1.0.6", "1.1.0", "1.1.1"}

	// Do some error checking
	// if !contains(validVersions, currentVersion) {
	// 	return nil, fmt.Errorf("current version %s is not valid", currentVersion)
	// }
	// if !contains(validVersions, targetVersion) {
	// 	return nil, fmt.Errorf("target version %s is not valid", targetVersion)
	// }
	current, err := compDefForVersion(data, currentVersion)
	if err != nil {
		return nil, err
	}
	for currentVersion < targetVersion {

		switch currentVersion {
		case "1.0.4":
			current, currentVersion, err = convertTo105(current.(compdef104.OscalComponentDefinitionModel))
			if err != nil {
				return nil, err
			}
		case "1.0.5":
			// do nothing
		case "1.0.6":
			// do nothing
		case "1.1.0":
			// do nothing
		case "1.1.1":
			// do nothing
		}
	}

	return nil, nil
}

// Convert to 1.0.5 will assume current version of 1.0.4
// No changes to schema - only update oscalVersion
// Still have to copy everything until we establish common structs across models? Would that even apply here?
func convertTo105(original compdef104.OscalComponentDefinitionModel) (compdef105.OscalComponentDefinitionModel, string, error) {
	var newDef compdef105.OscalComponentDefinitionModel

	// need to copy contents of metadata into newDef
	/*

		type ComponentDefinition struct {
			BackMatter                 BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
			Capabilities               []Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
			Components                 []DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
			ImportComponentDefinitions []ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
			Metadata                   Metadata                    `json:"metadata" yaml:"metadata"`
			UUID                       string                      `json:"uuid" yaml:"uuid"`
		}

	*/

	v := reflect.ValueOf(original.ComponentDefinition)
	typeOfS := v.Type()

	fmt.Println(typeOfS)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\n", typeOfS.Field(i).Name)
		fmt.Printf("Value: %s\n", v.Field(i).Type())
		if v.Field(i).Type() == "string" {

		}

	}
	return newDef, "1.0.5", nil
}

func compDefForVersion(data []byte, version string) (any, error) {

	switch version {
	case "1.0.4":
		var temp compdef104.OscalComponentDefinitionModel
		err := yaml.Unmarshal(data, &temp)
		if err != nil {
			return nil, err
		}
		return temp, nil
	case "1.0.5":
		var temp compdef105.OscalComponentDefinitionModel
		err := yaml.Unmarshal(data, &temp)
		if err != nil {
			return nil, err
		}
		return temp, nil
	default:
		return nil, fmt.Errorf("Unsupported version: %s", version)
	}
}
