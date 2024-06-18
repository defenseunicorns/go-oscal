package validation

import (
	"time"

	"github.com/defenseunicorns/go-oscal/src/pkg/files"
	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
)

type ValidationResult struct {
	// Valid is true if the validation result is valid
	Valid bool `json:"valid" yaml:"valid"`
	// TimeStamp is the time the validation result was created
	TimeStamp time.Time `json:"timeStamp" yaml:"timeStamp"`
	// Errors is a slice of ValidatorErrors
	Errors []ValidatorError `json:"errors,omitempty" yaml:"errors,omitempty"`
	// Metadata is the metadata of the validation result
	Metadata ValidationResultMetadata `json:"metadata" yaml:"metadata"`
}

type ValidationResultMetadata struct {
	// DocumentPath is the path to the document
	DocumentPath string `json:"documentPath,omitempty" yaml:"documentPath,omitempty"`
	// DocumentType is the type of the document
	DocumentType string `json:"documentType,omitempty" yaml:"documentType,omitempty"`
	// DocumentVersion is the version of the document
	DocumentVersion string `json:"documentVersion,omitempty" yaml:"documentVersion,omitempty"`
	// SchemaVersion is the version of the schema
	SchemaVersion string `json:"schemaVersion,omitempty" yaml:"schemaVersion,omitempty"`
}

// NewValidationResult creates a new ValidationResult from a Validator and a slice of ValidatorErrors
func NewValidationResult(validator *Validator, errors []ValidatorError) ValidationResult {
	documentVersion, err := versioning.GetOscalVersionFromMap(validator.GetJsonModel())
	if err != nil {
		documentVersion = ""
	}
	return ValidationResult{
		Valid:     len(errors) == 0,
		TimeStamp: time.Now(),
		Errors:    errors,
		Metadata: ValidationResultMetadata{
			DocumentPath:    validator.GetDocumentPath(),
			DocumentType:    validator.GetModelType(),
			DocumentVersion: documentVersion,
			SchemaVersion:   validator.GetSchemaVersion(),
		},
	}
}

// WriteValidationResult writes a ValidationResult to a file
func WriteValidationResult(validationResult ValidationResult, outputFile string) (err error) {
	validationResultBytes, err := model.MarshalByExtension(validationResult, outputFile)
	if err != nil {
		return err
	}
	return files.WriteOutput(validationResultBytes, outputFile)
}

// WriteValidationResults writes a slice of ValidationResults to a file
func WriteValidationResults(validationResults []ValidationResult, outputFile string) (err error) {
	resultMap := map[string][]ValidationResult{
		"results": validationResults,
	}
	validationResultBytes, err := model.MarshalByExtension(resultMap, outputFile)
	if err != nil {
		return err
	}
	return files.WriteOutput(validationResultBytes, outputFile)
}
