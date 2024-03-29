package utils_test

import (
	"os"
	"testing"

	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
)

func TestRemoveTrailingWhitespace(t *testing.T) {
	t.Parallel()

	t.Run("It removes trailing whitespace from each line in a byte array", func(t *testing.T) {
		t.Parallel()
		testBytes := []byte("test  \n  test  \n")
		expected := "test\n  test\n"
		result := string(utils.RemoveTrailingWhitespace(testBytes))
		if result != expected {
			t.Error("expected", expected, "got", result)
		}
	})

	t.Run("It returns the same byte array if there is no trailing whitespace", func(t *testing.T) {
		t.Parallel()
		testBytes := []byte("test\n  test\n")
		expected := "test\n  test\n"
		result := string(utils.RemoveTrailingWhitespace(testBytes))
		if result != expected {
			t.Error("expected", expected, "got", result)
		}
	})

	t.Run("It does not remove leading whitespace", func(t *testing.T) {
		t.Parallel()
		testBytes := []byte("  test\n  test\n")
		expected := "  test\n  test\n"
		result := string(utils.RemoveTrailingWhitespace(testBytes))
		if result != expected {
			t.Error("expected", expected, "got", result)
		}
	})
}

func TestFileUtils(t *testing.T) {
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
			err := utils.WriteOutput([]byte(testOutput), subdir+testFile)
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
	})

	t.Run("isJsonOrYaml", func(t *testing.T) {
		t.Parallel()
		t.Run("returns nil if the file path ends with the .json extension", func(t *testing.T) {
			t.Parallel()
			err := utils.IsJsonOrYaml(tmpDir + "/test.json")
			if err != nil {
				t.Error("expected no error, got", err)
			}

		})

		t.Run("returns nil if the file path ends with the .yaml extension", func(t *testing.T) {
			t.Parallel()
			err := utils.IsJsonOrYaml(tmpDir + "/test.yaml")
			if err != nil {
				t.Error("expected no error, got", err)
			}
		})

		t.Run("returns an error if the file path does not end with the .json or .yaml extension", func(t *testing.T) {
			t.Parallel()
			err := utils.IsJsonOrYaml(tmpDir + "/test.txt")
			if err == nil {
				t.Error("expected an error, got", err)
			}
		})
	})

	t.Run("isJson", func(t *testing.T) {
		t.Parallel()
		t.Run("returns nil if the file path ends with the .json extension", func(t *testing.T) {
			t.Parallel()
			err := utils.IsJson(tmpDir + "/test.json")
			if err != nil {
				t.Error("expected no error, got", err)
			}
		})

		t.Run("returns an error if the file path does not end with the .json extension", func(t *testing.T) {
			t.Parallel()
			err := utils.IsJson(tmpDir + "/test.yaml")
			if err == nil {
				t.Error("expected an error, got", err)
			}
		})
	})
}
