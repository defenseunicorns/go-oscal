package revise

import (
	"fmt"
	"log"

	"github.com/defenseunicorns/go-oscal/src/pkg/files"
	"github.com/defenseunicorns/go-oscal/src/pkg/revision"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
	"github.com/spf13/cobra"
)

var opts = &revision.RevisionOptions{}

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

		revisionResponse, revisionErr := revision.RevisionCommand(opts)

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
			err := files.WriteOutput(revisionResponse.RevisedBytes, opts.OutputFile)
			if err != nil {
				return fmt.Errorf("failed to write to output file: %s", err)
			}
		}

		log.Printf("Successfully upgraded %s from %s to version %s\n", revisionResponse.Reviser.GetModelType(), revisionResponse.Reviser.GetModelVersion(), revisionResponse.Reviser.GetSchemaVersion())

		return nil
	},
}

func init() {
	ReviseCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "input file to convert")
	ReviseCmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", "", "output file to write to")
	ReviseCmd.Flags().StringVarP(&opts.ValidationResult, "validation-result", "r", "", "validation result file to write to")
	ReviseCmd.Flags().StringVarP(&opts.Version, "version", "v", versioning.GetLatestSupportedVersion(), "version to convert to")
	ReviseCmd.MarkFlagRequired("input-file")
}
