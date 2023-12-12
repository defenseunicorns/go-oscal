package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

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

// OpenLogFile opens the provided input file or returns an error if the file does not exist.
func OpenLogFile(path string) (file *os.File, err error) {
	err = CreateFileDirs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file path: %s\n", err)
	}
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %s\n", err)
	}
	return file, nil
}

// CreateFileDirs creates the directories for a file path if they do not exist.
func CreateFileDirs(fullPath string) (err error) {
	pathSeperator := string(os.PathSeparator)
	splitPath := strings.Split(fullPath, pathSeperator)
	builtPath := ""
	if len(splitPath) == 1 {
		return fmt.Errorf("invalid path: %s\n", fullPath)
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
