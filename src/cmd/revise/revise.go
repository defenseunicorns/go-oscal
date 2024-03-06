package revise

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/pkg/revision"
	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/spf13/cobra"
)

type ReviseOptions struct {
	InputFile        string // short -f, long --file, required
	OutputFile       string // short -o, long --output, default to stdout
	Version          string // short -v, long --version, default to latest
	ValidationResult string // short -r, long --result, default to stdout
}

var opts = &ReviseOptions{}

var ReviseCmd = &cobra.Command{
	Use:   "revise",
	Short: "revise an existing oscal document to another version of oscal",
	Long:  "Revise a given model from one oscal version to the specified oscal version. The steps to revise are output to log, successful revision is output to stdout or the specified output file.",
	// Example: convertHelp,
	RunE: func(cmd *cobra.Command, componentDefinitionPaths []string) error {

		reviser, revisionErr := Revise(opts)

		// Write the validation result if it was specified and exists before returning Revise error (if there was one)
		result, err := reviser.GetValidationResult()
		if err == nil && opts.ValidationResult != "" {
			err = validation.WriteValidationResult(result, opts.ValidationResult)
			if err != nil {
				log.Printf("Failed to write validation result to %s: %s\n", opts.ValidationResult, err)
			}
		}

		// Return the error from the revision if there was one
		if revisionErr != nil {
			return revisionErr
		}

		var outputExt string
		if opts.OutputFile == "" {
			outputExt = "json"
		} else {
			split := strings.Split(opts.OutputFile, ".")
			outputExt = split[len(split)-1]
		}

		upgradeBytes, err := reviser.GetRevisedBytes(outputExt)
		if err != nil {
			return fmt.Errorf("Failed to get upgraded bytes: %s\n", err)
		}

		// Write the upgraded model to the output file or log
		if opts.OutputFile == "" {
			log.Println(string(upgradeBytes))
		} else {
			err = utils.WriteOutput(upgradeBytes, opts.OutputFile)
			if err != nil {
				return fmt.Errorf("Failed to write to output file: %s\n", err)
			}
		}

		log.Printf("Successfully upgraded %s from %s to version %s\n", reviser.GetModelType(), reviser.GetModelVersion(), reviser.GetSchemaVersion())

		return nil
	},
}

func Revise(opts *ReviseOptions) (reviser revision.Reviser, err error) {
	// Validate inputfile was provided and that is json or yaml
	if opts.InputFile == "" {
		return reviser, errors.New("Please specify an input file with the -f flag")
	} else {
		if err := utils.IsJsonOrYaml(opts.InputFile); err != nil {
			return reviser, fmt.Errorf("invalid input file: %s\n", err)
		}
	}

	// Validate outputfile is json or yaml, defaults to stdout
	if opts.OutputFile == "" {
		log.Printf("No output file specified, result will be logged\n")
	} else {
		if err := utils.IsJsonOrYaml(opts.OutputFile); err != nil {
			return reviser, fmt.Errorf("invalid output file: %s\n", err)
		}
	}

	// Validate version was provided
	if opts.Version == "" {
		return reviser, errors.New("Please specify a version to convert to with the -v flag")
	}

	// Read the input file
	bytes, err := os.ReadFile(opts.InputFile)
	if err != nil {
		return reviser, fmt.Errorf("reading input file: %s\n", err)
	}

	// Create Upgrader
	reviser, err = revision.NewReviser(bytes, opts.Version)
	if err != nil {
		return reviser, fmt.Errorf("Failed to create reviser: %s\n", err)
	}

	version := reviser.GetSchemaVersion()
	err = utils.VersionWarning(version)
	if err != nil {
		log.Print(err)
	}

	reviser.SetDocumentPath(opts.InputFile)

	// Run the upgrade
	err = reviser.Revise()
	if err != nil {
		return reviser, fmt.Errorf("failed to upgrade %s version %s: %s", reviser.GetModelType(), reviser.GetSchemaVersion(), err)
	}

	return reviser, nil
}

func init() {
	ReviseCmd.Flags().StringVarP(&opts.InputFile, "file", "f", "", "input file to convert")
	ReviseCmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "output file to write to")
	ReviseCmd.Flags().StringVarP(&opts.ValidationResult, "validation-result", "r", "", "validation result file to write to")
	ReviseCmd.Flags().StringVarP(&opts.Version, "version", "v", "", "version to convert to")
}
