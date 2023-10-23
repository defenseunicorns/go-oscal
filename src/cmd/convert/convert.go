package convert

import (
	"github.com/spf13/cobra"
)

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert existing oscal document to a newer version of oscal",
	Long:  "Convert a given model from an older oscal version to the current (or specified) oscal version ",
	// Example: convertHelp,
	RunE: func(cmd *cobra.Command, componentDefinitionPaths []string) error {

		// Logic here for the conversion process
		// This should really only be the process of opening the document to a []byte and then passing it
		// to a library that will convert it to the targeted object version.
		// Logic here will then handle writing the output to a file or STDOUT.

		return nil
	},
}

func ConvertCommand() *cobra.Command {

	// insert flag options here
	return ConvertCmd
}

// func init() {
// 	GenerateCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to a oscal json schema file")
// 	GenerateCmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
// 	GenerateCmd.Flags().StringVarP(&opts.Pkg, "pkg", "p", "main", "the name of the package for the generated code")
// 	GenerateCmd.Flags().StringVarP(&opts.Tags, "tags", "t", "json", "comma separated list of the tags to put on the struct")
// }
