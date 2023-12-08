package validate

import (
	"bytes"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
)

func TestValidateCmd(t *testing.T) {
	t.Parallel()
	ValidateCmd.SetOut(new(bytes.Buffer))
	logOutput := gooscaltest.RedirectLog(t)

	t.Run("returns an error if no inputfile is provided", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", ""})
		err := ValidateCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if the inputfile is not a json or yaml file", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", "test.txt"})
		err := ValidateCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if the inputfile is not a valid oscal document", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", gooscaltest.InvalidCatalogPath})
		err := ValidateCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns an error if it fails to read the input file", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", "test.json"})
		err := ValidateCmd.Execute()
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("logs a success message if the input file is valid", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", gooscaltest.ValidComponentPath})
		ValidateCmd.Execute()
		if !bytes.Contains(gooscaltest.ReadLog(t, logOutput), []byte("Successfully validated")) {
			t.Errorf("Expected log output to contain 'Successfully validated', got %s", logOutput.String())
		}
	})
}
