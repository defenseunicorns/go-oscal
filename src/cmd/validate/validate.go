package validate

import (
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
		ValidateCommand(inputfile)

	},
}

func ValidateCommand(inputFile string) {

	// Validate the input file
	if inputFile == "" {
		log.Fatalln("Please specify an input file with the -f flag")
	}

	// Validate the input file is a json or yaml file
	if !strings.HasSuffix(inputFile, "json") && !strings.HasSuffix(inputFile, "yaml") {
		log.Fatalln("Please specify a json or yaml file")
	}

	// Read the input file
	bytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("reading input file: %s\n", err)
	}

	validator, err := validation.NewValidator(bytes)
	if err != nil {
		log.Fatalf("Failed to create validator: %s\n", err)
	}

	err = validator.Validate()
	if err != nil {
		log.Fatalf("Failed to validate %s version %s: %s\n", validator.GetModelType(), validator.GetVersion(), err)
	}

	log.Printf("Successfully validated %s is valid OSCAL version %s %s\n", inputFile, validator.GetVersion(), validator.GetModelType())
}

func init() {
	ValidateCmd.Flags().StringVarP(&inputfile, "input-file", "f", "", "the path to a oscal json schema file")
}
