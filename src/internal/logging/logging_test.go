package logging_test

import (
	"testing"

	"github.com/defenseunicorns/go-oscal/src/internal/logging"
)

func TestLogging(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	t.Run("OpenLogFile", func(t *testing.T) {
		t.Parallel()
		subdir := tmpDir + "openLogFile"
		testFile := "/test.log"

		t.Run("It opens the provided log file", func(t *testing.T) {
			t.Parallel()
			file, err := logging.OpenLogFile(tmpDir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			if file == nil {
				t.Error("expected a file, got", file)
			}
		})

		t.Run("creates the file path if it doesn't exist", func(t *testing.T) {
			t.Parallel()
			file, err := logging.OpenLogFile(subdir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			if file == nil {
				t.Error("expected a file, got", file)
			}
		})

		t.Run("returns an error if no log file is provided", func(t *testing.T) {
			t.Parallel()
			file, err := logging.OpenLogFile("")
			if err == nil {
				t.Error("expected an error, got", err)
			}
			if file != nil {
				t.Error("expected no file, got", file)
			}
		})
	})
}
