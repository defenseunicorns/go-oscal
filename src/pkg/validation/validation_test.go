package validation

import (
	"os"
	"sync"
	"testing"
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

// func TestIsValidOscal(t *testing.T) {
// 	t.Parallel()
// 	getByteMap(t)

// 	t.Run("returns nil when a valid component definition of the correct version is passed", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[validComponentPath])
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}

// 	})

// 	t.Run("returns error when an invalid component definition of the correct version is passed", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[invalidComponentPath])
// 		if err == nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("returns nil when a valid assessment result of the correct version is passed", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[validAssessmentResultPath])
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("returns nil when a valid catalog of the correct version is passed", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[validCatalogPath])
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("returns nil when a valid profile of the correct version is passed", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[validProfilePath])
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("Handles json as well as yaml", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[validCatalogJsonPath])
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("returns error if there are multiple top level fields or OSCALModels", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := IsValidOscal(byteMap[multipleDocPath])
// 		if err == nil {
// 			t.Errorf("expected error, got %v", err)
// 		}
// 	})

// 	t.Run("Handles interface{} as well as []byte", func(t *testing.T) {
// 		t.Parallel()
// 		var modelInterface interface{}
// 		err := yaml.Unmarshal(byteMap[validComponentPath], &modelInterface)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 		_, err = IsValidOscal(modelInterface)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("Handles OscalModels struct as well as interface{} and []byte", func(t *testing.T) {
// 		t.Parallel()
// 		var oscalModels V104.OscalModels
// 		err := yaml.Unmarshal(byteMap[validComponentPath], &oscalModels)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 		_, err = IsValidOscal(oscalModels)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})
// }

// func TestValidateOscalAgainstVersion(t *testing.T) {
// 	t.Parallel()
// 	getByteMap(t)

// 	var validComponent V104.OscalModels
// 	err := yaml.Unmarshal(byteMap[validComponentPath], &validComponent)
// 	if err != nil {
// 		t.Errorf("expected no error, got %v", err)
// 	}

// 	t.Run("returns nil when a valid oscal map of the correct version is passed", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := ValidateOscalAgainstVersion(validVersion, validComponent)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}
// 	})

// 	t.Run("returns error when the provided oscal map fails validation for the provided version", func(t *testing.T) {
// 		t.Parallel()
// 		_, err := ValidateOscalAgainstVersion("1.0.5", map[string]interface{}{})
// 		if err == nil {
// 			t.Errorf("expected error, got %v", err)
// 		}
// 	})
// }

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
