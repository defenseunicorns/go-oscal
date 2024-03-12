package utils

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
	"github.com/goccy/go-yaml"
)

const (
	latestVersion = "1.1.2"
)

var (
	versionRegexp = regexp.MustCompile(`^\d+([-\.]\d+){2}$`)
)

// IsValidOscalVersion returns true if the version is supported, false if not.
func IsValidOscalVersion(version string) error {
	if err := validVersionFormat(version); err != nil {
		return err
	}

	version = FormatOscalVersion(version)
	schemaPath := schemas.SCHEMA_PREFIX + strings.ReplaceAll(version, ".", "-") + ".json"

	if _, err := schemas.Schemas.Open(schemaPath); os.IsNotExist(err) {
		return fmt.Errorf("version %s is not supported", version)
	}
	return nil
}

func validVersionFormat(version string) error {
	if !versionRegexp.MatchString(version) {
		return fmt.Errorf("version %s is not a valid version", version)
	}
	return nil
}

// GetLatestSupportedVersion returns the latest version of the OSCAL schema supported by the go-oscal release
func GetLatestSupportedVersion() string {
	return latestVersion
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

func ReplaceKeyInMapSlice(mapSlice yaml.MapSlice, path []string, value interface{}) (yaml.MapSlice, error) {
	if len(path) == 0 {
		return mapSlice, fmt.Errorf("path cannot be empty")
	}

	key := path[0]
	for i, item := range mapSlice {
		if item.Key == key {
			if len(path) == 1 {
				mapSlice[i].Value = value
				return mapSlice, nil
			}

			nextPath := path[1:]
			nextValue, ok := item.Value.(yaml.MapSlice)
			if !ok {
				return mapSlice, fmt.Errorf("expected map, got %T", item.Value)
			}

			mapSlice[i].Value, _ = ReplaceKeyInMapSlice(nextValue, nextPath, value)
			return mapSlice, nil
		}
	}

	return mapSlice, fmt.Errorf("key %s not found", key)
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
	UpdateLastModified(metadata)
	return model, nil
}

// UpdateLastModified updates the last-modified field in the metadata
func UpdateLastModified(metadata map[string]interface{}) {
	metadata["last-modified"] = GetTimestamp()
}

// VersionWarning returns an warning as an error if there are any known issues with the current version or it isn't the latest.
func VersionWarning(version string) error {
	switch version {
	case "1.0.5":
		return fmt.Errorf("WARNING: 1.0.5 has known issues. Please upgrade to version 1.0.6 or higher")
	default:
		if latestVersion != version {
			return fmt.Errorf("WARNING: Currently using OSCAL version %s. The latest version is %s", version, latestVersion)
		}
	}
	return nil
}
