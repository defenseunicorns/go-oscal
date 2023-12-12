package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	V105 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-5"
	V106 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-6"
	V110 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-0"
	V111 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-1"
	"gopkg.in/yaml.v3"
)

var (
	versionRegexp    = regexp.MustCompile(`^\d+([-\.]\d+){2}$`)
	supportedVersion = map[string]bool{
		"1.0.4": true,
		"1.0.5": true,
		"1.0.6": true,
		"1.1.0": true,
		"1.1.1": true,
	}
)

// IsValidOscalVersion returns true if the version is supported, false if not.
func IsValidOscalVersion(version string) error {
	if !versionRegexp.MatchString(version) {
		return fmt.Errorf("version %s is not a valid version", version)
	}

	if !supportedVersion[version] {
		return fmt.Errorf("version %s is not supported", version)
	}
	return nil
}

// FormatOscalVersion takes a version string and returns a formatted version string
func FormatOscalVersion(version string) (formattedVersion string) {
	formattedVersion = strings.ToLower(version)
	formattedVersion = strings.ReplaceAll(formattedVersion, "-", ".")
	formattedVersion = strings.Trim(formattedVersion, "v")
	return formattedVersion
}

// GetOscalVersionFromMap returns formatted OSCAL version if valid version is passed, returns error if not.
func GetOscalVersionFromMap(model map[string]interface{}) (version string, err error) {
	var submodel map[string]interface{}
	var ok bool
	for _, value := range model {
		submodel, ok = value.(map[string]interface{})
		if !ok {
			continue
		}
		break

	}

	metadata, ok := submodel["metadata"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("required field: metadata not found")
	}

	if _, ok := metadata["oscal-version"]; !ok {
		return "", fmt.Errorf("required field: oscal-version not found")
	}

	version, ok = reflect.ValueOf(metadata["oscal-version"]).Interface().(string)
	if !ok {
		return "", fmt.Errorf("required field: oscal-version not found")
	}

	if err = IsValidOscalVersion(version); err != nil {
		return "", err
	}

	return version, nil
}

// ReplaceOscalVersionInMap returns the model with the oscal version replaced
func ReplaceOscalVersionInMap(model map[string]interface{}, version string) (upgradedModel map[string]interface{}, err error) {
	upgradedModel = model
	var submodel map[string]interface{}
	var ok bool
	for _, value := range model {
		submodel, ok = value.(map[string]interface{})
		if !ok {
			continue
		}
		break
	}

	metadata, ok := submodel["metadata"].(map[string]interface{})
	if !ok {
		return model, fmt.Errorf("required field: metadata not found")
	}
	metadata["oscal-version"] = version
	return model, nil
}

// GetVersionedStruct returns the versioned struct based on the version passed
func GetVersionedStruct(bytes []byte, version string) (model interface{}, err error) {
	switch version {
	case "1.0.4":
		model = V104.OscalModels{}
		break
	case "1.0.5":
		model = V105.OscalModels{}
		break
	case "1.0.6":
		model = V106.OscalModels{}
		break
	case "1.1.0":
		model = V110.OscalModels{}
		break
	case "1.1.1":
		model = V111.OscalModels{}
		break
	}
	err = yaml.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
