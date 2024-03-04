package utils

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/schemas"
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

// GetLatestVersion returns the latest version of the OSCAL schema from the schemas directory
func GetLatestVersion() (latestVersion string, err error) {
	// calcVersionVal is a helper function to calculate the version value
	// by removing the dots and converting the version to an integer
	calcVersionVal := func(version string) (int, error) {
		split := strings.Split(version, ".")
		joined := strings.Join(split, "")
		return strconv.Atoi(joined)
	}

	// WalkDir walks the file tree rooted at schema.Schemas, calling walkFn for each file or directory in the tree, including root.
	err = fs.WalkDir(schemas.Schemas, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		// Extract the version from the path
		version := strings.TrimPrefix(path, schemas.SCHEMA_PREFIX)
		version = strings.TrimSuffix(version, ".json")
		version = FormatOscalVersion(version)

		// Check if the version is valid, if not, continue to next file
		if err := validVersionFormat(version); err != nil {
			return nil
		}

		// If latestVersion is empty, set it to the current version
		if latestVersion == "" {
			latestVersion = version
		} else {
			// Compare the current version with the latest version
			latestVal, err := calcVersionVal(latestVersion)
			if err != nil {
				return err
			}
			versionVal, err := calcVersionVal(version)
			if err != nil {
				return err
			}
			// If the current version is greater than the latest version, set the latest version to the current version
			if versionVal > latestVal {
				latestVersion = version
			}
		}
		return nil
	})

	return latestVersion, err
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
	UpdateLastModified(metadata)
	return model, nil
}

// UpdateLastModified updates the last-modified field in the metadata
func UpdateLastModified(metadata map[string]interface{}) {
	metadata["last-modified"] = GetTimestamp()
}

// VersionWarning returns an warning as an error if there are any known issues with the current version or it isn't the latest.
func VersionWarning(version string) error {
	latestVersion, err := GetLatestVersion()
	if err != nil {
		return err
	}
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
