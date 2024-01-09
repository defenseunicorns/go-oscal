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

var inputfile string
var validationResultFile string

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate an oscal document",
	Long:  "Validate an OSCAL document against the OSCAL schema version specified in the document.",
	RunE: func(cmd *cobra.Command, args []string) error {
		validator, validationErr := ValidateCommand(inputfile)

		// Write validation result if it was specified and exists before returning ValidateCommand error
		validationResult, err := validator.GetValidationResult()
		if err == nil && validationResultFile != "" {
			err = validation.WriteValidationResult(validationResult, validationResultFile)
			if err != nil {
				log.Printf("Failed to write validation result to %s: %s\n", validationResultFile, err)
			}
		}

		// Return the error from the validation if there was one
		if validationErr != nil {
			return validationErr
		}

		// No errors, log success
		log.Printf("Successfully validated %s is valid OSCAL version %s %s\n", inputfile, validator.GetSchemaVersion(), validator.GetModelType())
		return nil
	},
}

func ValidateCommand(inputFile string) (validator validation.Validator, err error) {
	// Validate the input file
	if inputFile == "" {
		return validator, errors.New("Please specify an input file with the -f flag")
	}

	// Validate the input file is a json or yaml file
	if !strings.HasSuffix(inputFile, "json") && !strings.HasSuffix(inputFile, "yaml") {
		return validator, errors.New("Please specify a json or yaml file")
	}

	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return validator, fmt.Errorf("reading input file: %s\n", err)
	}

	validator, err = validation.NewValidator(bytes)
	if err != nil {
		return validator, fmt.Errorf("Failed to create validator: %s\n", err)
	}

	validator.SetDocumentPath(inputFile)

	err = validator.Validate()
	if err != nil {
		return validator, fmt.Errorf("Failed to validate %s version %s: %s\n", validator.GetModelType(), validator.GetSchemaVersion(), err)
	}

	return validator, nil
}

func init() {
	ValidateCmd.Flags().StringVarP(&inputfile, "input-file", "f", "", "the path to a oscal json schema file")
	ValidateCmd.Flags().StringVarP(&validationResultFile, "validation-result", "r", "", "the path to a validation result file")
}
