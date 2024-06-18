package validate

import (
	"encoding/json"
	"fmt"
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
		validationResponse, err := validation.ValidationCommand(inputfile)
		// Return any non-validation errors if they exist
		if err != nil {
			return err
		}

		// Write validation result if it was specified and exists before returning ValidateCommand error
		if validationResultFile != "" {
			err := validation.WriteValidationResult(validationResponse.Result, validationResultFile)
			if err != nil {
				log.Printf("Failed to write validation result to %s: %s\n", validationResultFile, err)
			}
		}

		// Log any go-oscal related warnings (ie version warnings)
		for _, warning := range validationResponse.Warnings {
			log.Print(warning)
		}

		// Return any validation errors
		if validationResponse.JsonSchemaError != nil {
			jsonResult, err := json.MarshalIndent(validationResponse.Result, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to format validation result as JSON: %v", err)
			}
			return fmt.Errorf("invalid OSCAL document, results: %s", string(jsonResult))
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
