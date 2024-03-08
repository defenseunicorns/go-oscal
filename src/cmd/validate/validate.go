package validate

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/spf13/cobra"
)

type ValidationResponse struct {
	Validator validation.Validator
	Result    validation.ValidationResult
	Warnings  []string
}

var inputfile string
var validationResultFile string

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate an oscal document",
	Long:  "Validate an OSCAL document against the OSCAL schema version specified in the document.",
	RunE: func(cmd *cobra.Command, args []string) error {
		validationResponse, validationErr := ValidateCommand(inputfile)

		// Write validation result if it was specified and exists before returning ValidateCommand error
		if validationResultFile != "" {
			err := validation.WriteValidationResult(validationResponse.Result, validationResultFile)
			if err != nil {
				log.Printf("Failed to write validation result to %s: %s\n", validationResultFile, err)
			}
		}
		// Return the error from the validation if there was one
		if validationErr != nil {
			return validationErr
		}

		if len(validationResponse.Warnings) > 0 {
			for _, warning := range validationResponse.Warnings {
				log.Print(warning)
			}
		}

		// No errors, log success
		log.Printf("Successfully validated %s is valid OSCAL version %s %s\n", inputfile, validationResponse.Validator.GetSchemaVersion(), validationResponse.Validator.GetModelType())
		return nil
	},
}

func ValidateCommand(inputFile string) (validationResponse ValidationResponse, err error) {
	// Validate the input file
	if inputFile == "" {
		return validationResponse, errors.New("please specify an input file with the -f flag")
	}

	// Validate the input file is a json or yaml file
	if !strings.HasSuffix(inputFile, "json") && !strings.HasSuffix(inputFile, "yaml") {
		return validationResponse, errors.New("please specify a json or yaml file")
	}

	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		return validationResponse, fmt.Errorf("reading input file: %s", err)
	}

	validator, err := validation.NewValidator(bytes)
	if err != nil {
		return validationResponse, fmt.Errorf("failed to create validator: %s", err)
	}

	version := validator.GetSchemaVersion()
	err = utils.VersionWarning(version)
	if err != nil {
		validationResponse.Warnings = append(validationResponse.Warnings, err.Error())
	}

	validator.SetDocumentPath(inputFile)

	validationError := validator.Validate()

	// Write validation result if it was specified and exists before returning ValidateCommand error
	validationResult, _ := validator.GetValidationResult()

	validationResponse = ValidationResponse{
		Validator: validator,
		Result:    validationResult,
		Warnings:  validationResponse.Warnings,
	}

	if validationError != nil {
		return validationResponse, fmt.Errorf("failed to validate %s version %s: %s", validator.GetModelType(), validator.GetSchemaVersion(), err)
	}

	return validationResponse, nil
}

func init() {
	ValidateCmd.Flags().StringVarP(&inputfile, "input-file", "f", "", "the path to a oscal json schema file")
	ValidateCmd.Flags().StringVarP(&validationResultFile, "validation-result", "r", "", "the path to a validation result file")
}
