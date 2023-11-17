package validation

import (
	"fmt"
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
		invalidComponentPath = "../../../testdata/validation/test-data.yaml"
		validComponentPath   = "../../../testdata/validation/valid-test-data.yaml"
	)
	invalidComponentDefinition, err := readTestComponentDefinitionFile(t, invalidComponentPath)
	if err != nil {
		t.Fatal(err)
	}
	validComponentDefinition, err := readTestComponentDefinitionFile(t, validComponentPath)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("returns true when a valid component definition of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", validComponentDefinition)
		if !valid {
			t.Errorf("expected true, got %v", valid)
		}

	})

	t.Run("returns false when an invalid component definition of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", invalidComponentDefinition)
		if valid {
			t.Errorf("expected false, got %v", valid)
		}
	})
}

func readTestComponentDefinitionFile(t *testing.T, path string) (componentDefinition interface{}, err error) {
	t.Helper()

	componentDefinitionBytes, err := os.ReadFile(path)
	if err != nil {
		return componentDefinition, fmt.Errorf("failed to read in test component definition file: %v", err)
	}

	if err := yaml.Unmarshal(componentDefinitionBytes, &componentDefinition); err != nil {
		return componentDefinition, fmt.Errorf("failed to unmarshal test component definition file: %v", err)
	}

	return componentDefinition, err
}
