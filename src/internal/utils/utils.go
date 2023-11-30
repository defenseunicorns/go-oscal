package utils

import (
	"fmt"
	"log"
	"os"
)

// WriteOutput writes to the provided output file or to stdout by default.
func WriteOutput(output []byte, outputFile string) {
	if outputFile != "" {
		err := os.WriteFile(outputFile, output, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(string(output))
	}
}

// WriteFatalError writes an error to the provided output file or to stdout by default and then exits.
func WriteFatalError(err error, outputFile string) {
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(err.Error()), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(err.Error())
	}
	os.Exit(1)
}
