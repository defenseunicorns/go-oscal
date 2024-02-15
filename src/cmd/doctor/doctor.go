package doctor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string

var DoctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "doctor an oscal document",
	Long:  "Doctor an OSCAL document to ensure it is valid and conforms to the OSCAL schema version specified in the document.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Doctor the input file
		doctoredSchema, doctorErr := DoctorCommand(inputFile)

		// Return the error from the doctor if there was one
		if doctorErr != nil {
			return doctorErr
		}

		if outputFile == "" {
			// If no output file is specified, log the doctor output
			log.Println(string(doctoredSchema))
			return nil
		} else {
			// Write the doctor output to the specified file
			writeErr := utils.WriteOutput(doctoredSchema, outputFile)
			if writeErr != nil {
				return fmt.Errorf("failed to write doctor output to file: %s", writeErr)
			}
		}

		// No errors, log success
		log.Printf("Successfully doctored %s and wrote to %s\n", inputFile, outputFile)
		return nil
	},
}

func DoctorCommand(inputfile string) (doctoredSchema []byte, err error) {
	// Read the input file
	schemaBytes, err := os.ReadFile(inputfile)
	if err != nil {
		return doctoredSchema, err
	}

	err = utils.IsJson(inputfile)
	if err != nil {
		return doctoredSchema, err
	}

	// Unmarshal the input file
	var mapSchema map[string]interface{}
	err = json.Unmarshal(schemaBytes, &mapSchema)
	if err != nil {
		return doctoredSchema, err
	}

	err = utils.DoctorJSON(mapSchema, mapSchema, "")
	if err != nil {
		return doctoredSchema, err
	}

	// Marshal the doctored schema
	doctoredSchema, err = json.MarshalIndent(mapSchema, "", "  ")
	if err != nil {
		return doctoredSchema, err
	}
	return doctoredSchema, nil
}

func init() {
	DoctorCmd.Flags().StringVarP(&inputFile, "input-file", "f", "", "the path to a oscal json schema file")
	DoctorCmd.Flags().StringVarP(&outputFile, "output-file", "o", "", "the path to write the doctored oscal json schema file")
}
