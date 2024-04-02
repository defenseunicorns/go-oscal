package files

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// RemoveTrailingWhitespace removes trailing whitespace from each line in a byte array.
// This is useful for removing trailing whitespace from a yaml file before it is marshalled into a struct.
// Prevents strings from being collapsed into a single line by yaml parser.
func RemoveTrailingWhitespace(bytes []byte) []byte {
	str := string(bytes)
	lines := strings.Split(str, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimRight(l, " ")
	}
	return []byte(strings.Join(lines, "\n"))
}

// WriteOutput writes to the provided output file or to stdout by default.
func WriteOutput(output []byte, outputFile string) (err error) {
	if outputFile == "" {
		return errors.New("no output file provided")
	}
	err = CreateFileDirs(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file path: %s\n", err)
	}
	err = os.WriteFile(outputFile, output, 0644)
	if err != nil {
		return err
	}
	return nil
}

// CreateFileDirs creates the directories for a file path if they do not exist.
func CreateFileDirs(fullPath string) (err error) {
	pathSeperator := string(os.PathSeparator)
	splitPath := strings.Split(fullPath, pathSeperator)
	builtPath := ""
	// If the path is just a file name, return nil
	if len(splitPath) == 1 {
		if splitPath[0] == "" {
			return errors.New("no output file provided")
		}
		return nil
	}
	for _, location := range splitPath[:len(splitPath)-1] {
		builtPath += location + pathSeperator
		_, err := os.Stat(builtPath)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(builtPath, 0755)
				if err != nil {
					return fmt.Errorf("failed to create file path: %s\n", err)
				}
			} else {
				return err
			}
		}
	}
	return nil
}

// isJsonOrYaml returns error if the file is not a json or yaml file.
func IsJsonOrYaml(path string) (err error) {
	if !strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, ".yaml") {
		return errors.New("please specify a json or yaml file")
	}
	return nil
}

func IsJson(path string) (err error) {
	if !strings.HasSuffix(path, ".json") {
		return errors.New("please specify a json file")
	}
	return nil
}
