package validate

import (
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/spf13/cobra"
)

type flags struct {
	InputFile string // -f --input-file
	LogFile   string // -l --log-file

}

var opts = &flags{}

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate an oscal document",
	Long:  "Validate an OSCAL document against the OSCAL schema version specified in the document.",
	Run: func(cmd *cobra.Command, args []string) {
		ValidateCommand()

	},
}

func ValidateCommand() {
	var oscalModel string
	logger := log.New(os.Stderr, "", 0)

	// Set the logger output to a file if specified
	if opts.LogFile != "" {
		file, err := os.Create(opts.LogFile)
		if err != nil {
			log.Fatalln(err)
		}
		logger.SetOutput(file)
	}

	// Validate the input file
	if opts.InputFile == "" {
		log.Fatalln("Please specify an input file with the -f flag")
	}

	// Validate the input file is a json or yaml file
	if !strings.HasSuffix(opts.InputFile, "json") && !strings.HasSuffix(opts.InputFile, "yaml") {
		logger.Fatalln("Please specify a json or yaml file")
	}

	// Read the input file
	bytes, err := os.ReadFile(opts.InputFile)
	if err != nil {
		log.Fatalf("reading input file: %s\n", err)
	}

	oscalMap, err := validation.IsValidOscal(bytes)

	if err != nil {
		log.Fatalf("validating oscal: %s\n", err)
	}

	for key := range oscalMap {
		oscalModel = key
	}

	version := oscalMap[oscalModel].(map[string]interface{})["metadata"].(map[string]interface{})["oscal-version"].(string)

	log.Printf("Successfully validated %s is valid OSCAL version %s %s\n", opts.InputFile, version, oscalModel)
}

func init() {
	ValidateCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to a oscal json schema file")
	ValidateCmd.Flags().StringVarP(&opts.LogFile, "log-file", "l", "", "the name of the file to write logs to (outputs to STDERR by default)")
}
