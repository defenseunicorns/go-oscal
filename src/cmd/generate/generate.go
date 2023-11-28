package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/oscal"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"

	"github.com/spf13/cobra"
)

var opts = &oscal.BaseFlags{}

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate Golang data types from OSCAL schema",
	Long:  "Generate Golang data types from OSCAL schema",
	// Example: generateHelp,
	RunE: func(cmd *cobra.Command, componentDefinitionPaths []string) error {
		// Remove the transfer of objects across packages
		// aggregate file IO to root

		bytes, err := os.ReadFile(opts.InputFile)
		if err != nil {
			return err
		}

		// oscalMap, err := oscal.ParseJson(opts)
		// if err != nil {
		// 	return err
		// }

		tagList := formatTags()

		// Generate the Go structs.
		output, err := oscal.Generate(bytes, opts.Pkg, tagList)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error parsing", err)
			os.Exit(1)
		}
		// Write the Go struct output to either stdout or a file.
		utils.WriteOutput(output, opts.OutputFile)
		return nil
	},
}

func GenerateCommand() *cobra.Command {

	// insert flag options here
	return GenerateCmd
}

func init() {
	GenerateCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to a oscal json schema file")
	GenerateCmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	GenerateCmd.Flags().StringVarP(&opts.Pkg, "pkg", "p", "main", "the name of the package for the generated code")
	GenerateCmd.Flags().StringVarP(&opts.Tags, "tags", "t", "json", "comma separated list of the tags to put on the struct")
}

// formatTags formats Go struct tags.
func formatTags() (tagList []string) {
	if opts.Tags == "" {
		tagList = append(tagList, "json")
	} else {
		tagList = strings.Split(opts.Tags, ",")
	}

	return
}
