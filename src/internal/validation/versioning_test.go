package validation

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

const (
	nonExistentVersion     = "1.0.7"
	tooFewVersionNumbers   = "1.0"
	tooManyVersionNumbers  = "1.0.4.1"
	validVersion           = "1.0.4"
	validVersionWithDashes = "1-0-4"
	validVersionWithPrefix = "v1.0.4"
)

func TestOscalVersioning(t *testing.T) {
	t.Run("returns valid version when user version is in proper format", func(t *testing.T) {
		version, err := GetVersion(validVersion)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})

	t.Run("replaces dashes with periods when version given with dashes", func(t *testing.T) {
		version, err := GetVersion(validVersionWithDashes)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})

	t.Run("returns valid version when prefixed with v", func(t *testing.T) {
		version, err := GetVersion(validVersionWithPrefix)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})

	t.Run("uses the default oscal version when version is empty", func(t *testing.T) {
		version, err := GetVersion("")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != DEFAULT_OSCAL_VERSION {
			t.Errorf("expected %s, got %s", DEFAULT_OSCAL_VERSION, version)
		}
	})

	t.Run("throws error with invalid version structure", func(t *testing.T) {
		_, err := GetVersion(tooManyVersionNumbers)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error with too few version numbers", func(t *testing.T) {
		_, err := GetVersion(tooFewVersionNumbers)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error when version is not supported", func(t *testing.T) {
		_, err := GetVersion(nonExistentVersion)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})
}

// TODO: Add tests for supported versions other than V1.0.4
func TestIsValidSchemaVersion(t *testing.T) {
	var (
		invalidComponentPath = "../../../testdata/validation/invalid-component-definition.yaml"
		validComponentPath   = "../../../testdata/validation/valid-component-definition.yaml"
		assessmentResultPath = "../../../testdata/validation/assessment-result.yaml"
	)
	invalidComponentDefinitionYamlBytes := readFile(t, invalidComponentPath)
	validComponentDefinitionYamlBytes := readFile(t, validComponentPath)
	assessmentResultYamlBytes := readFile(t, assessmentResultPath)
	catalogJsonBytes := readFile(t, "../../../testdata/validation/catalog.json")

	t.Run("returns true when a valid component definition of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", validComponentDefinitionYamlBytes)
		if !valid {
			t.Errorf("expected true, got %v", valid)
		}

	})

	t.Run("returns false when an invalid component definition of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", invalidComponentDefinitionYamlBytes)
		if valid {
			t.Errorf("expected false, got %v", valid)
		}
	})

	t.Run("returns true when a valid assessment result of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", assessmentResultYamlBytes)
		if !valid {
			t.Errorf("expected true, got %v", valid)
		}
	})

	t.Run("Handles json as well as yaml", func(t *testing.T) {

		t.Run("returns true when a valid catalog of the correct version is passed", func(t *testing.T) {
			valid := IsValidSchemaVersion("1.1.0", catalogJsonBytes)
			if !valid {
				t.Errorf("expected true, got %v", valid)
			}
		})

	})
	t.Run("Handles interface{} as well as []byte", func(t *testing.T) {
		t.Run("returns true when a valid assessment result interface of the correct version is passed", func(t *testing.T) {
			assessmentResultInterface := unmarshalToYaml(t, assessmentResultYamlBytes)
			valid := IsValidSchemaVersion("1.0.4", assessmentResultInterface)
			if !valid {
				t.Errorf("expected true, got %v", valid)
			}
		})
		t.Run("returns true when a valid catalog interface of the correct version is passed", func(t *testing.T) {
			catalogResultInterface := unmarshalToYaml(t, catalogJsonBytes)
			valid := IsValidSchemaVersion("1.1.0", catalogResultInterface)
			if !valid {
				t.Errorf("expected true, got %v", valid)
			}
		})
	})
}

func unmarshalToYaml(t *testing.T, bytes []byte) (result interface{}) {
	t.Helper()
	err := yaml.Unmarshal(bytes, &result)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func readFile(t *testing.T, path string) []byte {
	t.Helper()

	docBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return docBytes
}
