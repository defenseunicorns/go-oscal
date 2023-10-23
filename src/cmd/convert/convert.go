package convert

import (
	"log"
	"os"

	"github.com/defenseunicorns/go-oscal/src/pkg/oscal"
	"github.com/spf13/cobra"
)

var inputFile, outputFile, targetVersion string

const latestVersion = "1.1.1"

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

		_, err := os.Stat(inputFile)

		if os.IsNotExist(err) {
			log.Fatal("file does not exist")
		}

		data, err := os.ReadFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}

		if targetVersion == "" {
			targetVersion = latestVersion
		}

		object, err := oscal.ConvertOscal(data, targetVersion)
		if err != nil {
			log.Fatal(err)
		}
		log.Default().Println(object)
		// Call library function here to convert the data to the desired object version.

		return nil
	},
}

func ConvertCommand() *cobra.Command {

	// insert flag options here
	return ConvertCmd
}

func init() {
	ConvertCmd.Flags().StringVarP(&inputFile, "input-file", "f", "", "the path to a oscal file")
	ConvertCmd.Flags().StringVarP(&outputFile, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	ConvertCmd.Flags().StringVarP(&targetVersion, "target-version", "t", "", "the desired version of the output file - otherwise latest")
}
