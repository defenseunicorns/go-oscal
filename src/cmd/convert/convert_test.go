package convert_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/cmd/convert"
	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

func TestConvertCmd(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)

	convert.ConvertCmd.SetOut(new(bytes.Buffer))
	log.SetOutput(new(bytes.Buffer))

	tempDir := t.TempDir()

	t.Run("Converts a valid catalog to json", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", gooscaltest.ValidCatalogPath,
			"-o", tempDir + "/catalog.json",
			"-v", "1.0.6",
		})
		if err := convert.ConvertCmd.Execute(); err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		actual, err := os.ReadFile(tempDir + "/catalog.json")
		if err != nil {
			t.Error(err)
		}
		validator, err := validation.NewValidator(actual)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		err = validator.Validate()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Converts a valid catalog to yaml", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", gooscaltest.ValidCatalogPath,
			"-o", tempDir + "/catalog.yaml",
			"-v", "1.0.6",
		})
		if err := convert.ConvertCmd.Execute(); err != nil {
			t.Error(err)
		}
		actual, err := os.ReadFile(tempDir + "/catalog.yaml")
		if err != nil {
			t.Error(err)
		}
		validator, err := validation.NewValidator(actual)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		err = validator.Validate()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

}
