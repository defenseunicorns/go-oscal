package validate

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/spf13/cobra"
)

type flags struct {
	InputFile string // -f --input-file
	LogFile   string // -l --logger-file

}

var opts = &flags{}

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate an oscal document",
	Long:  "Validate an OSCAL document against the OSCAL schema version specified in the document.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Move this to main command, use global logger with cobra
		logger := log.New(os.Stderr, "", log.LstdFlags)

		// Set the logger output to a file if specified
		if opts.LogFile != "" {
			file, err := os.OpenFile(opts.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				logger.Fatalln(err)
			}
			logger.SetOutput(file)
		}
		successMsg, err := ValidateCommand(opts.InputFile)
		if err != nil {
			logger.Fatalln(err)
		}
		logger.Println(successMsg)
		return err
	},
}

func ValidateCommand(inputFile string) (string, error) {

	// Validate the input file
	if inputFile == "" {
		return "", errors.New("Please specify an input file with the -f flag")
	}

	// Validate the input file is a json or yaml file
	if !strings.HasSuffix(inputFile, "json") && !strings.HasSuffix(inputFile, "yaml") {
		return "", errors.New("Please specify a json or yaml file")
	}

	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("Failed to read input file %s: %s\n", inputFile, err)
	}

	validator, err := validation.NewValidator(bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to create validator: %s\n", err)
	}

	err = validator.Validate()
	if err != nil {
		return "", fmt.Errorf("Failed to validate %s version %s: %s\n", validator.GetModelType(), validator.GetVersion(), err)
	}

	return fmt.Sprintf("Successfully validated %s is valid OSCAL version %s %s\n", inputFile, validator.GetVersion(), validator.GetModelType()), nil
}

func init() {
	ValidateCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to a oscal json schema file")
	ValidateCmd.Flags().StringVarP(&opts.LogFile, "logger-file", "l", "", "the name of the file to write logs to (outputs to STDERR by default)")
}
