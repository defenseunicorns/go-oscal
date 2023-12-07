package validate

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
)

func TestValidateCmd(t *testing.T) {
	t.Parallel()
	logOutput := new(bytes.Buffer)
	log.SetOutput(logOutput)

	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
	})

	t.Run("Logs an error if no inputfile is provided", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", ""})
		ValidateCmd.Execute()
		if !bytes.Contains(logOutput.Bytes(), []byte("Please specify an input file with the -f flag")) {
			t.Errorf("Expected log output to contain 'Please specify an input file with the -f flag', got %s", logOutput.String())
		}
	})

	t.Run("Logs an error if the inputfile is not a json or yaml file", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", "test.txt"})
		ValidateCmd.Execute()
		if !bytes.Contains(logOutput.Bytes(), []byte("Please specify a json or yaml file")) {
			t.Errorf("Expected log output to contain 'Please specify a json or yaml file', got %s", logOutput.String())
		}
	})

	t.Run("Logs an error if the inputfile is not a valid oscal document", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", gooscaltest.InvalidCatalogPath})
		ValidateCmd.Execute()
		if !bytes.Contains(logOutput.Bytes(), []byte("Failed to validate")) {
			t.Errorf("Expected log output to contain 'Failed to validate', got %s", logOutput.String())
		}
	})

	t.Run("logs an error if it fails to read the input file", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", "test.json"})
		ValidateCmd.Execute()
		if !bytes.Contains(logOutput.Bytes(), []byte("reading input file")) {
			t.Errorf("Expected log output to contain 'reading input file', got %s", logOutput.String())
		}
	})

	t.Run("logs an error if it fails to create the validator", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", gooscaltest.NoVersionComponentPath})
		ValidateCmd.Execute()
		if !bytes.Contains(logOutput.Bytes(), []byte("Failed to create validator")) {
			t.Errorf("Expected log output to contain 'Failed to create validator', got %s", logOutput.String())
		}
	})

	t.Run("logs a success message if the input file is valid", func(t *testing.T) {
		ValidateCmd.SetArgs([]string{"-f", gooscaltest.ValidComponentPath})
		ValidateCmd.Execute()
		if !bytes.Contains(logOutput.Bytes(), []byte("Successfully validated")) {
			t.Errorf("Expected log output to contain 'Successfully validated', got %s", logOutput.String())
		}
	})
}
