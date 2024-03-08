package validate

import (
	"log"

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
		// Run the validation
		validationResponse, validationErr := validation.ValidationCommand(inputfile)

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

func init() {
	ValidateCmd.Flags().StringVarP(&inputfile, "input-file", "f", "", "the path to a oscal json schema file")
	ValidateCmd.Flags().StringVarP(&validationResultFile, "validation-result", "r", "", "the path to a validation result file")
}
