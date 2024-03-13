package revision

import (
	"errors"
	"fmt"
	"os"

	"github.com/defenseunicorns/go-oscal/src/pkg/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

type RevisionResponse struct {
	Reviser      Reviser
	Result       validation.ValidationResult
	Warnings     []string
	RevisedBytes []byte
}

type RevisionOptions struct {
	InputFile        string // filepath json or yaml
	OutputFile       string // filepath json or yaml
	Version          string // OSCAL version X.X.X
	ValidationResult string // filepath json or yaml
}

// RevisionCommand Runs the revision and returns the revision response
// For Consumers, this method assumes no ReviseOptions defaults. All defaults are handled in the cobra command.
func RevisionCommand(opts *RevisionOptions) (revisionResponse RevisionResponse, err error) {
	// Validate inputfile was provided and that is json or yaml
	if err := utils.IsJsonOrYaml(opts.InputFile); err != nil {
		return revisionResponse, fmt.Errorf("invalid input file: %s", err)
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

	bytes = utils.RemoveTrailingWhitespace(bytes)

	// Create the reviser and set it in the response
	reviser, err := NewReviser(bytes, opts.Version)
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
