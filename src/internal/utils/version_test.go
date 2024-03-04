package utils_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"time"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"gopkg.in/yaml.v3"
)

func TestVersionUtils(t *testing.T) {
	var validVersion = "1.0.4"

	t.Run("isValidOscalVersion", func(t *testing.T) {
		t.Parallel()
		var (
			unsupportedVersion = "1.0.7"
			invalidVersion     = "invalid-version-1-0-4"
		)

		t.Run("returns nil when a valid version is passed", func(t *testing.T) {
			t.Parallel()
			err := utils.IsValidOscalVersion(validVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})

		t.Run("returns error when an unsupported version is passed", func(t *testing.T) {
			t.Parallel()
			err := utils.IsValidOscalVersion(unsupportedVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("returns error when an invalid version is passed", func(t *testing.T) {
			t.Parallel()
			err := utils.IsValidOscalVersion(invalidVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})
	})

	t.Run("formatOscalVersion", func(t *testing.T) {
		t.Parallel()
		var (
			vPrefixed     = "v1.0.4"
			capVPrefixed  = "V1.0.4"
			dashedVersion = "1-0-4"
		)

		t.Run("returns formatted version when a valid version is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := utils.FormatOscalVersion(validVersion)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})

		t.Run("returns formatted version when a valid version prefixed with 'v' is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := utils.FormatOscalVersion(vPrefixed)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})

		t.Run("returns formatted version when a valid version prefixed with 'V' is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := utils.FormatOscalVersion(capVPrefixed)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})

		t.Run("returns formatted version when a valid version with dashes is passed", func(t *testing.T) {
			t.Parallel()
			formattedVersion := utils.FormatOscalVersion(dashedVersion)
			if formattedVersion != validVersion {
				t.Errorf("expected %s, got %s", validVersion, formattedVersion)
			}
		})
	})

	validModel := map[string]interface{}{
		"component-definition": map[string]interface{}{
			"metadata": map[string]interface{}{
				"oscal-version": validVersion,
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

	t.Run("getOscalVersionFromMap", func(t *testing.T) {
		t.Parallel()

		t.Run("returns the oscal version given a model.", func(t *testing.T) {
			t.Parallel()
			version, err := utils.GetOscalVersionFromMap(validModel)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if version != validVersion {
				t.Errorf("expected %s, got %s", validVersion, version)
			}
		})

		t.Run("throws error when version is not found", func(t *testing.T) {
			t.Parallel()
			_, err := utils.GetOscalVersionFromMap(invalidOscalModel)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

		t.Run("throws error for empty version", func(t *testing.T) {
			t.Parallel()
			_, err := utils.GetOscalVersionFromMap(noOscalVersionModel)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})
	})

	t.Run("replaceOscalVersionInMap", func(t *testing.T) {
		t.Parallel()

		var upgradeVersion = "1.0.6"
		t.Run("returns the model with the oscal version replaced", func(t *testing.T) {
			t.Parallel()
			upgradeModel := map[string]interface{}{
				"component-definition": map[string]interface{}{
					"metadata": map[string]interface{}{
						"oscal-version": validVersion,
					},
				},
			}
			model, err := utils.ReplaceOscalVersionInMap(upgradeModel, upgradeVersion)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			version, err := utils.GetOscalVersionFromMap(model)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if version != upgradeVersion {
				t.Errorf("expected %s, got %s", upgradeVersion, version)
			}
		})

		t.Run("throws error when version is not found", func(t *testing.T) {
			t.Parallel()
			_, err := utils.ReplaceOscalVersionInMap(invalidOscalModel, validVersion)
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}
		})

	})

	t.Run("UpdateLastModified", func(t *testing.T) {
		t.Parallel()

		t.Run("updates the last modified field", func(t *testing.T) {
			t.Parallel()

			initTime, error := time.Parse(time.RFC3339, "2020-09-09T00:00:00Z")
			if error != nil {
				t.Errorf("expected no error, got %v", error)
			}
			metadata := map[string]interface{}{
				"last-modified": initTime,
			}
			utils.UpdateLastModified(metadata)
			if metadata["last-modified"].(time.Time).Format(time.RFC3339) == initTime.Format(time.RFC3339) {
				t.Errorf("expected last-modified to be updated")
			}
		})
	})

	t.Run("VersionWarning", func(t *testing.T) {
		t.Parallel()
		logBytes := gooscaltest.RedirectLog(t)

		latestVersion, err := utils.GetLatestVersion()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		t.Run("prints a warning if the version is has known issues", func(t *testing.T) {
			utils.VersionWarning("1.0.5")
			if !strings.Contains(string(gooscaltest.ReadLog(t, logBytes)), "WARNING: 1.0.5 has known issues.") {
				t.Errorf("expected warning to be printed")
			}
		})

		t.Run("prints a warning if not on the latest version", func(t *testing.T) {
			utils.VersionWarning("1.0.6")
			if !strings.Contains(string(gooscaltest.ReadLog(t, logBytes)), fmt.Sprintf("WARNING: A new version of the OSCAL schema is available. Please upgrade to version %s", latestVersion)) {
				t.Errorf("expected warning to be printed")
			}
		})

		t.Run("does not print a warning if on the latest version", func(t *testing.T) {
			utils.VersionWarning(latestVersion)
			if strings.Contains(string(gooscaltest.ReadLog(t, logBytes)), fmt.Sprintf("WARNING: A new version of the OSCAL schema is available. Please upgrade to version %s", latestVersion)) {
				t.Errorf("expected no warning to be printed")
			}
		})
	})
}

func TestGetLatestVersion(t *testing.T) {
	t.Parallel()
	latestVersionPath := "../../../update/oscal-version.yaml"

	// Read the latest version from the file (updates from renovate PR when a new version of OSCAL is released)
	bytes, err := os.ReadFile(latestVersionPath)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Unmarshal the latest version from the file
	var updateVersion map[string]string
	err = yaml.Unmarshal(bytes, &updateVersion)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	// Format the latest version
	expected := utils.FormatOscalVersion(updateVersion["oscal"])

	t.Run("returns the latest version", func(t *testing.T) {
		t.Parallel()
		latestVersion, err := utils.GetLatestVersion()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if latestVersion != expected {
			t.Errorf("expected %s, got %s", expected, latestVersion)
		}
	})

}
