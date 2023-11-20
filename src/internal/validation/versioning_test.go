package validation

import (
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	V104 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-0-4"
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

var (
	invalidComponentPath = "../../../testdata/validation/invalid-component-definition.yaml"
	validComponentPath   = "../../../testdata/validation/valid-component-definition.yaml"
	assessmentResultPath = "../../../testdata/validation/assessment-result.yaml"
	validCatalogPath     = "../../../testdata/validation/catalog.yaml"
	catalogJsonPath      = "../../../testdata/validation/catalog.json"
	pathSlice            = []string{invalidComponentPath, validComponentPath, assessmentResultPath, validCatalogPath, catalogJsonPath}
	byteMap              = map[string][]byte{}
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

func TestGettingVersionFromModel(t *testing.T) {
	initByteMap(t)

	t.Run("returns valid version when a valid component definition is passed", func(t *testing.T) {
		version, err := GetOscalModelVersion(byteMap[validComponentPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != "1.0.4" {
			t.Errorf("expected %s, got %s", "1.0.4", version)
		}
	})

	t.Run("returns valid version when a valid assessment result is passed", func(t *testing.T) {
		version, err := GetOscalModelVersion(byteMap[assessmentResultPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != "1.0.4" {
			t.Errorf("expected %s, got %s", "1.0.4", version)
		}
	})

	t.Run("returns valid version when a valid catalog is passed", func(t *testing.T) {
		version, err := GetOscalModelVersion(byteMap[validCatalogPath])
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if version != "1.1.0" {
			t.Errorf("expected %s, got %s", "1.1.0", version)
		}
	})

	t.Run("returns an error when there is no metadata.oscal-version", func(t *testing.T) {
		_, err := GetOscalModelVersion(byteMap[invalidComponentPath])
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("returns an error when the []byte is json or yaml", func(t *testing.T) {
		_, err := GetOscalModelVersion(interface{}(
			[]byte(`not a map[string]interface{}`)))
		if err == nil {
			t.Errorf("expected error, got %v", err)
		}
	})

	t.Run("oscal struct behaviors", func(t *testing.T) {
		t.Run("it retrieves the oscal-version from a valid oscal struct", func(t *testing.T) {
			var oscalModel V104.OscalModels
			validComponentBytes := byteMap[validComponentPath]
			err := yaml.Unmarshal(validComponentBytes, &oscalModel)
			if err != nil {
				t.Errorf("expected no error, got %v when attempting to unmarshal", err)
			}
			oscalVersion, err := GetOscalModelVersion(oscalModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if oscalVersion != "1.0.4" {
				t.Errorf("expected %s, got %s", "1.0.4", oscalVersion)
			}
		})

		t.Run("it returns an error if the oscal struct is un-initialized", func(t *testing.T) {
			oscalModels := V104.OscalModels{}
			_, err := GetOscalModelVersion(oscalModels)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("it returns an error if the oscal struct oscal-version is empty", func(t *testing.T) {
			oscalModels := V104.OscalModels{
				ComponentDefinition: V104.ComponentDefinition{
					Metadata: V104.Metadata{
						OscalVersion: "",
					},
				},
			}
			_, err := GetOscalModelVersion(oscalModels)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})
	})
}

func TestIsValidSchemaVersion(t *testing.T) {
	initByteMap(t)

	t.Run("returns true when a valid component definition of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", byteMap[validComponentPath])
		if !valid {
			t.Errorf("expected true, got %v", valid)
		}

	})

	t.Run("returns false when an invalid component definition of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", byteMap[invalidComponentPath])
		if valid {
			t.Errorf("expected false, got %v", valid)
		}
	})

	t.Run("returns true when a valid assessment result of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.0.4", byteMap[assessmentResultPath])
		if !valid {
			t.Errorf("expected true, got %v", valid)
		}
	})

	t.Run("returns true when a valid catalog of the correct version is passed", func(t *testing.T) {
		valid := IsValidSchemaVersion("1.1.0", byteMap[validCatalogPath])
		if !valid {
			t.Errorf("expected true, got %v", valid)
		}
	})

	t.Run("Handles json as well as yaml", func(t *testing.T) {

		t.Run("returns true when a valid catalog of the correct version is passed", func(t *testing.T) {
			valid := IsValidSchemaVersion("1.1.0", byteMap[catalogJsonPath])
			if !valid {
				t.Errorf("expected true, got %v", valid)
			}
		})

	})
	t.Run("Handles interface{} as well as []byte", func(t *testing.T) {
		t.Run("returns true when a valid assessment result interface of the correct version is passed", func(t *testing.T) {
			assessmentResultInterface := unmarshalToYaml(t, byteMap[assessmentResultPath])
			valid := IsValidSchemaVersion("1.0.4", assessmentResultInterface)
			if !valid {
				t.Errorf("expected true, got %v", valid)
			}
		})
		t.Run("returns true when a valid catalog interface of the correct version is passed", func(t *testing.T) {
			catalogResultInterface := unmarshalToYaml(t, byteMap[catalogJsonPath])
			valid := IsValidSchemaVersion("1.1.0", catalogResultInterface)
			if !valid {
				t.Errorf("expected true, got %v", valid)
			}
		})
	})
}

func initByteMap(t *testing.T) {
	if byteMap[invalidComponentPath] == nil {
		for _, path := range pathSlice {
			bytes, err := utils.ReadFile(path)
			if err != nil {
				panic(err)
			}
			byteMap[path] = bytes
		}
	}
}

func unmarshalToYaml(t *testing.T, bytes []byte) (result interface{}) {
	t.Helper()
	err := yaml.Unmarshal(bytes, &result)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func convertToYaml(t *testing.T, version string, outputPath string, bytes []byte) {
	t.Helper()
	model := GetVersionedModel(version)
	err := yaml.Unmarshal(bytes, &model)
	if err != nil {
		t.Fatal(err)
	}

	yamlBytes, err := yaml.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile(outputPath, yamlBytes, 0644)

}

func readFile(t *testing.T, path string) []byte {
	t.Helper()

	docBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return docBytes
}
