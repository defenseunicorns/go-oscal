package revision_test

import (
	"testing"

	"github.com/defenseunicorns/go-oscal/src/gooscaltest"
	"github.com/defenseunicorns/go-oscal/src/pkg/revision"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
)

func TestRevisionCommand(t *testing.T) {
	// TestRevisionCommand tests the revision command
	t.Parallel()

	t.Run("returns an error if no input file is provided", func(t *testing.T) {
		opts := &revision.RevisionOptions{}
		_, err := revision.RevisionCommand(opts)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error if the input file is not json or yaml", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile: "invalid",
		}
		_, err := revision.RevisionCommand(opts)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error if the output file is not json or yaml", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile:  gooscaltest.ValidCatalogPath,
			OutputFile: "invalid",
		}
		_, err := revision.RevisionCommand(opts)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error if no version is provided", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile:  gooscaltest.ValidCatalogPath,
			OutputFile: gooscaltest.ValidCatalogPath,
		}
		_, err := revision.RevisionCommand(opts)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error if the input file is not found", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile:  "invalid.json",
			OutputFile: gooscaltest.ValidCatalogPath,
			Version:    "1.0.6",
		}
		_, err := revision.RevisionCommand(opts)
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns an error and ValidationResponse if the input file fails validation", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile:  gooscaltest.InvalidCatalogPath,
			OutputFile: gooscaltest.ValidCatalogPath,
			Version:    "1.0.6",
		}
		revisionResponse, err := revision.RevisionCommand(opts)

		if revisionResponse.Result.TimeStamp == (validation.ValidationResult{}).TimeStamp {
			t.Error("expected a validation response, got nil")
		}

		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("returns a warning if the version is not the latest", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile:  gooscaltest.ValidCatalogPath,
			OutputFile: gooscaltest.ValidCatalogPath,
			Version:    "1.0.6",
		}
		revisionResponse, err := revision.RevisionCommand(opts)

		if len(revisionResponse.Warnings) == 0 {
			t.Error("expected a warning, got none")
		}

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("returns the upgraded bytes if the revision is successful", func(t *testing.T) {
		opts := &revision.RevisionOptions{
			InputFile:  gooscaltest.ValidCatalogPath,
			OutputFile: gooscaltest.ValidCatalogPath,
			Version:    "1.0.6",
		}
		revisionResponse, err := revision.RevisionCommand(opts)

		if len(revisionResponse.RevisedBytes) == 0 {
			t.Error("expected upgraded bytes, got none")
		}

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
