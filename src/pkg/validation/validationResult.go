package validation

import (
	"time"

	"github.com/defenseunicorns/go-oscal/src/pkg/files"
	"github.com/defenseunicorns/go-oscal/src/pkg/model"
	"github.com/defenseunicorns/go-oscal/src/pkg/versioning"
)

type ValidationResult struct {
	Valid     bool                     `json:"valid" yaml:"valid"`
	TimeStamp time.Time                `json:"timeStamp" yaml:"timeStamp"`
	Errors    []ValidatorError         `json:"errors,omitempty" yaml:"errors,omitempty"`
	Metadata  ValidationResultMetadata `json:"metadata" yaml:"metadata"`
}

type ValidationResultMetadata struct {
	DocumentPath    string `json:"documentPath,omitempty" yaml:"documentPath,omitempty"`
	DocumentType    string `json:"documentType,omitempty" yaml:"documentType,omitempty"`
	DocumentVersion string `json:"documentVersion,omitempty" yaml:"documentVersion,omitempty"`
	SchemaVersion   string `json:"schemaVersion,omitempty" yaml:"schemaVersion,omitempty"`
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
