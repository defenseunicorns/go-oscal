package cmd

import (
	"fmt"
	"log"
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

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateDefinitionStructs() {
	_, convertFloats, _, tagList := validateInput()

	output, err := oscal.GenerateComponentDefinitionStructs(name, pkg, tagList, subStruct, convertFloats)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	}

	if outputFileName != "" {
		if err := os.WriteFile(outputFileName, output, 0644); err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(string(output))
	}
}
