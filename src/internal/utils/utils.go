package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// WriteOutput writes to the provided output file or to stdout by default.
func WriteOutput(output []byte, outputFile string) {
	if outputFile != "" {
		err := os.WriteFile(outputFile, output, 0644)
		if err != nil {
			err = CreateFilePath(outputFile)
			if err != nil {
				log.Fatalf("writing output: %s", err)
			}
		}
	} else {
		log.Print(string(output))
	}
}

func ReadLogFile(path string) (file *os.File, err error) {
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		err = CreateFilePath(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create log file path: %s\n", err)
		}
		return ReadLogFile(path)

	}
	return file, nil
}

func CreateFilePath(path string) (err error) {
	pathSeperator := string(os.PathSeparator)
	splitPath := strings.Split(path, pathSeperator)
	builtPath := ""
	if len(splitPath) == 1 {
		return fmt.Errorf("invalid path: %s\n", path)
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
