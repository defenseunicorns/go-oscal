package convert_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/cmd/convert"
	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

func TestConvertCmd(t *testing.T) {
	t.Parallel()
	gooscaltest.GetByteMap(t)
	convert.ConvertCmd.SetOut(new(bytes.Buffer))
	logBytes := gooscaltest.RedirectLog(t)

	tempDir := t.TempDir()

	t.Run("Converts a valid catalog to json", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", gooscaltest.ValidCatalogPath,
			"-v", "1.0.6",
			"-o", tempDir + "/catalog.json",
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
			"-v", "1.0.6",
			"-o", tempDir + "/catalog.yaml",
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

	t.Run("throws an error if input file is not json or yaml", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", "test.txt",
			"-v", "1.0.6",
			"-o", tempDir + "/catalog.yaml",
		})
		if err := convert.ConvertCmd.Execute(); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("throws an error if version is not supported", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", gooscaltest.ValidCatalogPath,
			"-v", "1.0.3",
			"-o", tempDir + "/catalog.yaml",
		})
		if err := convert.ConvertCmd.Execute(); err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("throws an error if the output file provided is not json or yaml", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", gooscaltest.ValidCatalogPath,
			"-v", "1.0.6",
			"-o", tempDir + "/catalog.txt",
		})
		if err := convert.ConvertCmd.Execute(); err == nil {
			t.Error("expected error, got nil")
		}
		convert.ConvertCmd.SetArgs([]string{})
	})

	t.Run("outputs to the log if no output file is provided", func(t *testing.T) {
		convert.ConvertCmd.SetArgs([]string{
			"-f", gooscaltest.ValidCatalogPath,
			"-v", "1.0.6",
			"-o", "",
		})
		if err := convert.ConvertCmd.Execute(); err != nil {
			t.Error(err)
		}
		actual := gooscaltest.ReadLog(t, logBytes)
		if !strings.Contains(string(actual), "Successfully upgraded") {
			t.Errorf("expected to find log message, got %s", string(actual))
		}
	})
}
