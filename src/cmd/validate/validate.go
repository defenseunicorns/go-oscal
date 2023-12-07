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

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate an oscal document",
	Long:  "Validate an OSCAL document against the OSCAL schema version specified in the document.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := ValidateCommand(inputfile)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Successfully validated %s is valid OSCAL version %s %s\n", inputfile, validator.GetVersion(), validator.GetModelType())
		}
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

	err = validator.Validate()
	if err != nil {
		return validator, fmt.Errorf("Failed to validate %s version %s: %s\n", validator.GetModelType(), validator.GetVersion(), err)
	}

	return validator, nil
}

func init() {
	ValidateCmd.Flags().StringVarP(&inputfile, "input-file", "f", "", "the path to a oscal json schema file")
}
