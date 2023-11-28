package validation

import (
	"os"
	"sync"
	"testing"

	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
	"gopkg.in/yaml.v3"
)

var (
	unsupportedVersion        = "1.0.7"
	tooManyVersionNumbers     = "1.0.4.1"
	validVersion              = "1.0.4"
	invalidComponentPath      = "../../../testdata/validation/invalid-component-definition.yaml"
	validComponentPath        = "../../../testdata/validation/valid-component-definition.yaml"
	validAssessmentResultPath = "../../../testdata/validation/assessment-result.yaml"
	validCatalogPath          = "../../../testdata/validation/catalog.yaml"
	validCatalogJsonPath      = "../../../testdata/validation/catalog.json"
	validProfilePath          = "../../../testdata/validation/profile.yaml"
	multipleDocPath           = "../../../testdata/validation/multiple.yaml"
	pathSlice                 = []string{invalidComponentPath, validComponentPath, validAssessmentResultPath, validCatalogPath, validCatalogJsonPath, validProfilePath, multipleDocPath}
	byteMap                   = map[string][]byte{}
	mutex                     = sync.Mutex{}
)

func TestGetOscalVersionFromMap(t *testing.T) {
	t.Parallel()
	getByteMap(t)

	validModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": validVersion,
			},
		},
	}

	failRegexModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": tooManyVersionNumbers,
			},
		},
	}

	unsupportedVersionModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": unsupportedVersion,
			},
		},
	}

	invalidOscalModel := map[string]interface{}{
		"component-definition": map[string]interface{}{},
		"oscal-version":        validVersion,
	}

	noOscalVersionModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": "",
			},
		},
	}

	t.Run("returns valid version when user version is in proper format", func(t *testing.T) {
		t.Parallel()
		version, err := GetOscalVersionFromMap(validModel)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != validVersion {
			t.Errorf("expected %s, got %s", validVersion, version)
		}
	})
	t.Run("throws error with invalid version structure", func(t *testing.T) {
		t.Parallel()
		_, err := GetOscalVersionFromMap(failRegexModel)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error when version is not supported", func(t *testing.T) {
		t.Parallel()
		_, err := GetOscalVersionFromMap(unsupportedVersionModel)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error when version is not found", func(t *testing.T) {
		t.Parallel()
		_, err := GetOscalVersionFromMap(invalidOscalModel)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("throws error for empty version", func(t *testing.T) {
		t.Parallel()
		_, err := GetOscalVersionFromMap(noOscalVersionModel)
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})
}

func TestIsValidOscal(t *testing.T) {
	t.Parallel()
	getByteMap(t)

	t.Run("returns nil when a valid component definition of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[validComponentPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

	})

	t.Run("returns error when an invalid component definition of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[invalidComponentPath])
		if err == nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("returns nil when a valid assessment result of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[validAssessmentResultPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("returns nil when a valid catalog of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[validCatalogPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("returns nil when a valid profile of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[validProfilePath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Handles json as well as yaml", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[validCatalogJsonPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("returns error if there are multiple top level fields or OSCALModels", func(t *testing.T) {
		t.Parallel()
		_, err := IsValidOscal(byteMap[multipleDocPath])
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("Handles interface{} as well as []byte", func(t *testing.T) {
		t.Parallel()
		var modelInterface interface{}
		err := yaml.Unmarshal(byteMap[validComponentPath], &modelInterface)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		_, err = IsValidOscal(modelInterface)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Handles OscalModels struct as well as interface{} and []byte", func(t *testing.T) {
		t.Parallel()
		var oscalModels V104.OscalModels
		err := yaml.Unmarshal(byteMap[validComponentPath], &oscalModels)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		_, err = IsValidOscal(oscalModels)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func TestValidateOscalAgainstVersion(t *testing.T) {
	t.Parallel()
	getByteMap(t)

	var validComponent V104.OscalModels
	err := yaml.Unmarshal(byteMap[validComponentPath], &validComponent)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	t.Run("returns nil when a valid oscal map of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		_, err := ValidateOscalAgainstVersion(validVersion, validComponent)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("returns error when the provided oscal map fails validation for the provided version", func(t *testing.T) {
		t.Parallel()
		_, err := ValidateOscalAgainstVersion("1.0.5", map[string]interface{}{})
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})
}

func getByteMap(t *testing.T) {
	mutex.Lock()
	if len(byteMap) == 0 {
		for _, path := range pathSlice {
			bytes, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			byteMap[path] = bytes
		}
	}
	mutex.Unlock()
}
