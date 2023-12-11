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
	err = os.WriteFile(outputFile, output, 0644)
	if err != nil {
		err = CreateFileDirs(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file path: %s\n", err)
		}
	}
	return nil
}

func OpenLogFile(path string) (file *os.File, err error) {
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		err = CreateFileDirs(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create log file path: %s\n", err)
		}
		return OpenLogFile(path)
	}
	return file, nil
}

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
		if os.IsNotExist(err) {
			err = os.Mkdir(builtPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to create log file path: %s\n", err)
			}
		}
		if os.IsPermission(err) {
			return err
		}
	}
	return nil
}
