package upgrading_test

import (
	"strings"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/upgrading"
)

func TestUpgrader(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)

	t.Run("NewUpgrader", func(t *testing.T) {
		t.Parallel()
		t.Run("returns a new Upgrader", func(t *testing.T) {
			t.Parallel()
			upgrader, err := upgrading.NewUpgrader(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.6")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			expected := "1.0.4"
			actual := upgrader.GetModelVersion()
			if actual != expected {
				t.Errorf("expected %s, got %s", expected, actual)
			}
		})
	})

	t.Run("Upgrade", func(t *testing.T) {
		t.Parallel()

		t.Run("returns an error when it is unable to upgrade the current model", func(t *testing.T) {
			t.Parallel()
			upgrader, err := upgrading.NewUpgrader(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.5")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = upgrader.Upgrade()
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}

		})

		t.Run("success", func(t *testing.T) {
			t.Parallel()
			upgrader, err := upgrading.NewUpgrader(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.6")
			err = upgrader.Upgrade()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			t.Run("sets the upgradedJsonMap when successful", func(t *testing.T) {
				t.Parallel()
				if upgrader.GetUpgradedJsonMap() == nil {
					t.Errorf("expected upgradedJsonMap, got nil")
				}
			})

			t.Run("sets the upgradedJsonMap to the upgraded model version", func(t *testing.T) {
				t.Parallel()
				expected := "1.0.6"
				actual, err := utils.GetOscalVersionFromMap(upgrader.GetUpgradedJsonMap())
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if actual != expected {
					t.Errorf("expected %s, got %s", expected, actual)
				}
			})

			t.Run("maintains the original version", func(t *testing.T) {
				expected := "1.0.4"
				actual, err := utils.GetOscalVersionFromMap(upgrader.GetJsonModel())
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if actual != expected {
					t.Errorf("expected %s, got %s", expected, actual)
				}
			})
		})
	})

	t.Run("GetUpgradedBytes", func(t *testing.T) {
		t.Parallel()
		upgrader, err := upgrading.NewUpgrader(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.6")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		err = upgrader.Upgrade()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		t.Run("returns an error if the extension is not json or yaml", func(t *testing.T) {
			t.Parallel()
			_, err := upgrader.GetUpgradedBytes("txt")
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		})

		t.Run("returns the upgraded model as json", func(t *testing.T) {
			t.Parallel()
			bytes, err := upgrader.GetUpgradedBytes("json")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !strings.Contains(string(bytes), "\"oscal-version\":\"1.0.6\"") {
				t.Errorf("expected version 1.0.6, got %s", string(bytes))
			}
		})

		t.Run("returns the upgraded model as yaml", func(t *testing.T) {
			t.Parallel()
			bytes, err := upgrader.GetUpgradedBytes("yaml")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !strings.Contains(string(bytes), "oscal-version: 1.0.6") {
				t.Errorf("expected version 1.0.6, got %s", string(bytes))
			}
		})
	})

}
