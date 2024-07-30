package validation

import (
	"reflect"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
	"gopkg.in/yaml.v3"
)

var (
	unsupportedVersion    = "-2.0.7"
	validVersion          = "1.0.4"
	tooManyVersionNumbers = "1.0.4.1"
	upgradeVersion        = "1.0.6"
)

func TestValidator(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)

	t.Run("NewValidator", func(t *testing.T) {
		t.Parallel()

		t.Run("returns *validator when a valid model is passed", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("creates a json map from the byte slice", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
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
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetSchemaVersion() != validVersion {
				t.Errorf("expected version %s, got %s", validVersion, validator.GetSchemaVersion())
			}
		})

		t.Run("sets the model type of the validator to the type of the model", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
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
			err := yaml.Unmarshal(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], &modelInterface)
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
			_, err := NewValidator(gooscaltest.ByteMap[gooscaltest.NoVersionComponentPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when the model has an invalid version", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(gooscaltest.ByteMap[gooscaltest.InvalidVersionComponentPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when the model has an unsupported version", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(gooscaltest.ByteMap[gooscaltest.UnsupportedVersionComponentPath])
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error if there are multiple top level fields or OSCALModels", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidator(gooscaltest.ByteMap[gooscaltest.MultipleDocPath])
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
			validator, err := NewValidatorDesiredVersion(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], inputVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetSchemaVersion() != inputVersion {
				t.Errorf("expected version %s, got %s", validVersion, validator.GetSchemaVersion())
			}
		})

		t.Run("formats the version to the correct format", func(t *testing.T) {
			t.Parallel()
			inputVersion := "V1-1-0"
			validator, err := NewValidatorDesiredVersion(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], inputVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if validator.GetSchemaVersion() != "1.1.0" {
				t.Errorf("expected version %s, got %s", validVersion, validator.GetSchemaVersion())
			}
		})

		t.Run("returns error when an unsupported version is passed", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidatorDesiredVersion(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], unsupportedVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when the version is invalid", func(t *testing.T) {
			t.Parallel()
			_, err := NewValidatorDesiredVersion(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], tooManyVersionNumbers)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

	})

	t.Run("GetValidationResult", func(t *testing.T) {
		t.Parallel()

		t.Run("returns a ValidationResult if validation has occured", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = validator.Validate()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			result, err := validator.GetValidationResult()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !result.Valid {
				t.Errorf("expected result.Valid to be true, got false")
			}

		})

		t.Run("throws an error if the validator has not yet been through validation", func(t *testing.T) {
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			_, err = validator.GetValidationResult()
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})
	})
	t.Run("IsLatestOscalVersion", func(t *testing.T) {
		t.Parallel()

		latestVersion := versioning.GetLatestSupportedVersion()

		t.Run("returns false if the model is not the latest version of the OSCAL schema", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			isLatest, err := validator.IsLatestOscalVersion()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if isLatest {
				t.Errorf("expected false, got true")
			}
		})

		t.Run("returns true if the model is the latest version of the OSCAL schema", func(t *testing.T) {
			t.Parallel()
			validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			jsonMap := validator.GetJsonModel()
			jsonMap, err = versioning.ReplaceOscalVersionInMap(jsonMap, latestVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			validator, err = NewValidator(jsonMap)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			isLatest, err := validator.IsLatestOscalVersion()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !isLatest {
				t.Errorf("expected true, got false")
			}
		})
	})
}

func TestValidator_Validate(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)

	t.Run("returns nil when a valid component definition of the correct version is passed", func(t *testing.T) {
		t.Parallel()
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidComponentPath])
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.InvalidCatalogPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		err = validator.Validate()
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
		result, err := validator.GetValidationResult()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result.Valid {
			t.Errorf("expected false, got true")
		}
	})

	t.Run("returns error when it fails to validate against the desired version", func(t *testing.T) {
		t.Parallel()
		validator, err := NewValidatorDesiredVersion(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.5")
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidCatalogPath])
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidAssessmentResultPath])
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidProfilePath])
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidSSP])
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidPlanOfActionAndMilestones])
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
		validator, err := NewValidator(gooscaltest.ByteMap[gooscaltest.ValidAsessmentPlan])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		err = validator.Validate()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
