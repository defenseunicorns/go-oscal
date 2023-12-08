package utils_test

import (
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
)

func UtilsTest(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	t.Run("WriteOutput", func(t *testing.T) {
		t.Parallel()
		testOutput := "test output"
		subdir := tmpDir + "/output"
		testFile := "/test.txt"

		t.Run("It writes to the provided output file", func(t *testing.T) {
			t.Parallel()
			err := utils.WriteOutput([]byte(testOutput), tmpDir+testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}

			bytes, err := os.ReadFile(tmpDir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			if string(bytes) != testOutput {
				t.Error("expected", testOutput, "got", bytes)
			}
		})

		t.Run("creates the file path if it doesn't exist", func(t *testing.T) {
			t.Parallel()
			err := utils.WriteOutput([]byte(testOutput), subdir+"/output/test.txt")
			if err != nil {
				t.Error("expected no error, got", err)
			}

			bytes, err := os.ReadFile(subdir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			if string(bytes) != testOutput {
				t.Error("expected", testOutput, "got", bytes)
			}
		})

		t.Run("returns an error if no output file is provided", func(t *testing.T) {
			t.Parallel()
			err := utils.WriteOutput([]byte(testOutput), "")
			if err == nil {
				t.Error("expected an error, got", err)
			}
		})
	})

	t.Run("OpenLogFile", func(t *testing.T) {
		t.Parallel()
		subdir := tmpDir + "openLogFile"
		testFile := "/test.log"

		t.Run("It opens the provided log file", func(t *testing.T) {
			t.Parallel()
			file, err := utils.OpenLogFile(tmpDir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			if file == nil {
				t.Error("expected a file, got", file)
			}
		})

		t.Run("creates the file path if it doesn't exist", func(t *testing.T) {
			t.Parallel()
			file, err := utils.OpenLogFile(subdir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			if file == nil {
				t.Error("expected a file, got", file)
			}
		})

		t.Run("returns an error if no log file is provided", func(t *testing.T) {
			t.Parallel()
			file, err := utils.OpenLogFile("")
			if err == nil {
				t.Error("expected an error, got", err)
			}
			if file != nil {
				t.Error("expected no file, got", file)
			}
		})
	})

	t.Run("CreateFileDirs", func(t *testing.T) {
		t.Parallel()
		subdir := tmpDir + "/createFilePath"
		testFile := "/test.txt"

		t.Run("It creates the provided file path", func(t *testing.T) {
			t.Parallel()
			err := utils.CreateFileDirs(subdir + testFile)
			if err != nil {
				t.Error("expected no error, got", err)
			}
			_, err = os.Stat(subdir)
			if os.IsNotExist(err) {
				t.Error("expected", tmpDir, "to exist, got", err)
			}
		})

		t.Run("returns an error if no file path is provided", func(t *testing.T) {
			t.Parallel()
			err := utils.CreateFileDirs("")
			if err == nil {
				t.Error("expected an error, got", err)
			}
		})

		t.Run("returns an error if the file path is invalid", func(t *testing.T) {
			t.Parallel()
			err := utils.CreateFileDirs(tmpDir + "test.txt")
			if err == nil {
				t.Error("expected an error, got", err)
			}
		})
	})
}
