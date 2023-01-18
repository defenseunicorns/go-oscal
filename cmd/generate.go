package cmd

import (
	"fmt"
	"os"

	"github.com/defenseunicorns/go-oscal/internal/oscal"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Args:  cobra.MaximumNArgs(0),
	Short: "Generate OSCAL component definition structs from OSCAL schema types",
	Run: func(cmd *cobra.Command, args []string) {
		generateDefinitionStructs()
	},
}

// Add generateCmd to the parent rootCmd
func init() {
	rootCmd.AddCommand(generateCmd)
}

// generateDefinitionStructs validates user input,
// generates OSCAL component definition structs,
// and either writes them to a file or prints to STDOUT
func generateDefinitionStructs() {
	_, convertFloats, _, tagList := validateInput()

	output, err := oscal.GenerateComponentDefinitionStructs(name, pkg, tagList, subStruct, convertFloats)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	}

	// If the --output-file flag was provided, write the structs to the provided file
	// if not, print to STDOUT
	if outputFileName != "" {
		writeOutputFile(output)
	} else {
		printStdout(output)
	}
}
