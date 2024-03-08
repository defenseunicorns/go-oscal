package revise

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/defenseunicorns/go-oscal/src/pkg/revision"
	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/spf13/cobra"
)

type RevisionResponse struct {
	Reviser      revision.Reviser
	Result       validation.ValidationResult
	Warnings     []string
	RevisedBytes []byte
}

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
		// Default .json so the output can be printed to stdout in json format.
		var OUTPUT_DEFAULT = ".json"

		// If output file is not specified, set it to json, so it will not throw an error and can be printed to stdout
		if opts.OutputFile == "" {
			opts.OutputFile = OUTPUT_DEFAULT
		}

		revisionResponse, revisionErr := RevisionCommand(opts)

		if opts.ValidationResult != "" {
			err := validation.WriteValidationResult(revisionResponse.Result, opts.ValidationResult)
			if err != nil {
				log.Printf("Failed to write validation result to %s: %s\n", opts.ValidationResult, err)
			}
		}

		// Return the error from the revision if there was one
		if revisionErr != nil {
			return revisionErr
		}

		// Log any warnings
		if len(revisionResponse.Warnings) > 0 {
			for _, warning := range revisionResponse.Warnings {
				log.Print(warning)
			}
		}

		// Write the upgraded model to the output file or log
		if opts.OutputFile == OUTPUT_DEFAULT {
			log.Println(string(revisionResponse.RevisedBytes))
		} else {
			err := utils.WriteOutput(revisionResponse.RevisedBytes, opts.OutputFile)
			if err != nil {
				return fmt.Errorf("failed to write to output file: %s", err)
			}
		}

		log.Printf("Successfully upgraded %s from %s to version %s\n", revisionResponse.Reviser.GetModelType(), revisionResponse.Reviser.GetModelVersion(), revisionResponse.Reviser.GetSchemaVersion())

		return nil
	},
}

// RevisionCommand Runs the revision and returns the revision response
// For Consumers, this method assumes no ReviseOptions defaults. All defaults are handled in the cobra command.
func RevisionCommand(opts *ReviseOptions) (revisionResponse RevisionResponse, err error) {
	// Validate inputfile was provided and that is json or yaml
	if opts.InputFile == "" {
		return revisionResponse, errors.New("please specify an input file with the -f flag")
	} else {
		if err := utils.IsJsonOrYaml(opts.InputFile); err != nil {
			return revisionResponse, fmt.Errorf("invalid input file: %s", err)
		}
	}

	// If output file is not json or yaml, return an error
	if err := utils.IsJsonOrYaml(opts.OutputFile); err != nil {
		return revisionResponse, fmt.Errorf("invalid output file: %s", err)
	}

	// If version is not specified, return an error
	if opts.Version == "" {
		return revisionResponse, errors.New("please specify a version to convert to with the -v flag")
	}

	// Read the input file
	bytes, err := os.ReadFile(opts.InputFile)
	if err != nil {
		return revisionResponse, fmt.Errorf("reading input file: %s", err)
	}

	// Create the reviser and set it in the response
	reviser, err := revision.NewReviser(bytes, opts.Version)
	if err != nil {
		return revisionResponse, fmt.Errorf("failed to create reviser: %s", err)
	}
	revisionResponse.Reviser = reviser

	// Get the schema version and set warnings if they exist
	version := reviser.GetSchemaVersion()
	err = utils.VersionWarning(version)
	if err != nil {
		revisionResponse.Warnings = append(revisionResponse.Warnings, err.Error())
	}

	// Set the document path
	reviser.SetDocumentPath(opts.InputFile)

	// Run the revision
	err = reviser.Revise()

	// Get the validation result and set it in the response
	validationResult, _ := reviser.GetValidationResult()
	revisionResponse.Result = validationResult

	// return the revision error if there was one
	if err != nil {
		return revisionResponse, fmt.Errorf("failed to upgrade %s version %s: %s", reviser.GetModelType(), reviser.GetSchemaVersion(), err)
	}

	// Get the upgraded bytes and set them in the response
	revisedBytes, err := reviser.GetRevisedBytes(opts.OutputFile)
	if err != nil {
		return revisionResponse, fmt.Errorf("failed to get upgraded bytes: %s", err)
	}
	revisionResponse.RevisedBytes = revisedBytes

	return revisionResponse, nil
}

func init() {
	ReviseCmd.Flags().StringVarP(&opts.InputFile, "file", "f", "", "input file to convert")
	ReviseCmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "output file to write to")
	ReviseCmd.Flags().StringVarP(&opts.ValidationResult, "validation-result", "r", "", "validation result file to write to")
	ReviseCmd.Flags().StringVarP(&opts.Version, "version", "v", utils.GetLatestSupportedVersion(), "version to convert to")
}
