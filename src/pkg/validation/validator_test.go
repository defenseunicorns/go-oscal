package validation

import (
	"os"
	"reflect"
	"sync"
	"testing"

	"gopkg.in/yaml.v3"
)

var (
	unsupportedVersion              = "1.0.7"
	validVersion                    = "1.0.4"
	tooManyVersionNumbers           = "1.0.4.1"
	mutex                           = sync.Mutex{}
	byteMap                         = map[string][]byte{}
	validComponentPath              = "../../../testdata/validation/valid-component-definition.yaml"
	noVersionComponentPath          = "../../../testdata/validation/no-version-component-definition.yaml"
	invalidVersionComponentPath     = "../../../testdata/validation/invalid-version-component-definition.yaml"
	unsupportedVersionComponentPath = "../../../testdata/validation/unsupported-version-component-definition.yaml"
	validAssessmentResultPath       = "../../../testdata/validation/assessment-result.yaml"
	validCatalogPath                = "../../../testdata/validation/catalog.yaml"
	validCatalogJsonPath            = "../../../testdata/validation/catalog.json"
	invalidCatalogPath              = "../../../testdata/validation/invalid-catalog.yaml"
	validProfilePath                = "../../../testdata/validation/profile.yaml"
	validSSP                        = "../../../testdata/validation/system-security-plan.yaml"
	validPlanOfActionAndMilestones  = "../../../testdata/validation/plan-of-action-and-milestones.json"
	validAsessmentPlan              = "../../../testdata/validation/assessment-plan.json"
	multipleDocPath                 = "../../../testdata/validation/multiple.yaml"

	pathSlice = []string{validComponentPath, noVersionComponentPath, invalidVersionComponentPath, unsupportedVersionComponentPath, validAssessmentResultPath, validCatalogPath, validCatalogJsonPath, invalidCatalogPath, validProfilePath, validSSP, multipleDocPath, validPlanOfActionAndMilestones, validAsessmentPlan}
)

func TestValidator(t *testing.T) {
	t.Parallel()
	getByteMap(t)

	t.Run("NewValidator", func(t *testing.T) {
		t.Parallel()

		t.Run("returns *validator when a valid model is passed", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator == nil {
				t.Errorf("expected validator, got nil")
			}
		})

		t.Run("creates a json map from the byte slice", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			jsonModel := validator.GetJsonModel()
			if _, ok := reflect.ValueOf(jsonModel).Interface().(map[string]interface{}); !ok {
				t.Errorf("expected json map, got nil")
			}
		})

		t.Run("sets the version of the validator to the oscal-version of the model", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetVersion() != validVersion {
				t.Errorf("expected version %s, got %s", validVersion, validator.GetVersion())
			}
		})

		t.Run("sets the model type of the validator to the type of the model", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetModelType() != "component-definition" {
				t.Errorf("expected model type component-definition, got %s", validator.GetModelType())
			}
		})

		t.Run("handles interface{} as well as []byte", func(t *testing.T) {
			t.Parallel()
			var modelInterface interface{}
			err := yaml.Unmarshal(byteMap[validComponentPath], &modelInterface)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			_, err = NewValidator(modelInterface)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("returns error when the model has no version", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(byteMap[noVersionComponentPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when the model has an invalid version", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(byteMap[invalidVersionComponentPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when the model has an unsupported version", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(byteMap[unsupportedVersionComponentPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error if there are multiple top level fields or OSCALModels", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(byteMap[multipleDocPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}

		})
	})

	t.Run("NewValidatorDesiredVersion", func(t *testing.T) {
		t.Parallel()

		t.Run("sets the version of the validator to the desired version", func(t *testing.T) {
			t.Parallel()
			inputVersion := "1.1.0"
			validator, err := NewValidatorDesiredVersion(byteMap[validComponentPath], inputVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetVersion() != inputVersion {
				t.Errorf("expected version %s, got %s", validVersion, validator.GetVersion())
			}
		})

		t.Run("formats the version to the correct format", func(t *testing.T) {
			t.Parallel()
			inputVersion := "V1-1-0"
			validator, err := NewValidatorDesiredVersion(byteMap[validComponentPath], inputVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetVersion() != "1.1.0" {
				t.Errorf("expected version %s, got %s", validVersion, validator.GetVersion())
			}
		})

		t.Run("returns error when an unsupported version is passed", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidatorDesiredVersion(byteMap[validComponentPath], unsupportedVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when the version is invalid", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidatorDesiredVersion(byteMap[validComponentPath], tooManyVersionNumbers)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

	})

	t.Run("Validate", func(t *testing.T) {
		t.Parallel()

		t.Run("returns nil when a valid component definition of the correct version is passed", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("returns error when it fails to validate against the supported schema", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[invalidCatalogPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when it fails to validate against the desired version", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidatorDesiredVersion(byteMap[validComponentPath], "1.0.5")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("supports catalog", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validCatalogPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("supports assessment-result", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validAssessmentResultPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("supports profile", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validProfilePath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("supports system-security-plan", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validSSP])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("supports plan-of-action-and-milestones", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validPlanOfActionAndMilestones])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("supports assessment-plan", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(byteMap[validAsessmentPlan])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
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
