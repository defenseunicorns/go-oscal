package revision_test

import (
	"strings"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	revision "github.com/defenseunicorns/go-oscal/src/pkg/revision"
)

func TestRevisor(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)

	t.Run("NewReviser", func(t *testing.T) {
		t.Parallel()
		t.Run("returns a new Upgrader", func(t *testing.T) {
			t.Parallel()
			reviser, err := revision.NewReviser(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.6")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			expected := "1.0.4"
			actual := reviser.GetModelVersion()
			if actual != expected {
				t.Errorf("expected %s, got %s", expected, actual)
			}
		})
	})

	t.Run("Upgrade", func(t *testing.T) {
		t.Parallel()

		t.Run("returns an error when it is unable to upgrade the current model", func(t *testing.T) {
			t.Parallel()
			reviser, err := revision.NewReviser(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.5")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			err = reviser.Revise()
			if err == nil {
				t.Errorf("expected error, got %v", err)
			}

		})

		t.Run("success", func(t *testing.T) {
			t.Parallel()
			reviser, err := revision.NewReviser(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.6")
			err = reviser.Revise()
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			t.Run("sets the upgradedJsonMap when successful", func(t *testing.T) {
				t.Parallel()
				if reviser.GetRevisedJsonMap() == nil {
					t.Errorf("expected upgradedJsonMap, got nil")
				}
			})

			t.Run("sets the upgradedJsonMap to the upgraded model version", func(t *testing.T) {
				t.Parallel()
				expected := "1.0.6"
				actual, err := utils.GetOscalVersionFromMap(reviser.GetRevisedJsonMap())
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if actual != expected {
					t.Errorf("expected %s, got %s", expected, actual)
				}
			})

			t.Run("maintains the original version", func(t *testing.T) {
				expected := "1.0.4"
				actual, err := utils.GetOscalVersionFromMap(reviser.GetJsonModel())
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
		reviser, err := revision.NewReviser(gooscaltest.ByteMap[gooscaltest.ValidComponentPath], "1.0.6")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		err = reviser.Revise()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		t.Run("returns an error if the extension is not json or yaml", func(t *testing.T) {
			t.Parallel()
			_, err := reviser.GetRevisedBytes("txt")
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		})

		t.Run("returns the upgraded model as json", func(t *testing.T) {
			t.Parallel()
			bytes, err := reviser.GetRevisedBytes("json")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !strings.Contains(string(bytes), "\"oscal-version\":\"1.0.6\"") {
				t.Errorf("expected version 1.0.6, got %s", string(bytes))
			}
		})

		t.Run("returns the upgraded model as yaml", func(t *testing.T) {
			t.Parallel()
			bytes, err := reviser.GetRevisedBytes("yaml")
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if !strings.Contains(string(bytes), "oscal-version: 1.0.6") {
				t.Errorf("expected version 1.0.6, got %s", string(bytes))
			}
		})
	})

}
