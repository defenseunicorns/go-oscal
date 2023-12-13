package convert

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/upgrading"
	"github.com/spf13/cobra"
)

type ConvertOptions struct {
	InputFile  string // short -f, long --file, required
	OutputFile string // short -o, long --output, default to stdout
	Version    string // short -v, long --version, default to latest
}

var opts = &ConvertOptions{}

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert existing oscal document to a newer version of oscal",
	Long:  "Convert a given model from an older oscal version to the current (or specified) oscal version ",
	// Example: convertHelp,
	RunE: func(cmd *cobra.Command, componentDefinitionPaths []string) error {

		upgrader, err := ConvertCommand(opts)
		if err != nil {
			return err
		}

		var outputExt string
		if opts.OutputFile == "" {
			outputExt = "json"
		} else {
			split := strings.Split(opts.OutputFile, ".")
			outputExt = split[len(split)-1]
		}

		upgradeBytes, err := upgrader.GetUpgradedBytes(outputExt)
		if err != nil {
			return fmt.Errorf("Failed to get upgraded bytes: %s\n", err)
		}

		// Write the upgraded model to the output file or log
		if opts.OutputFile == "" {
			log.Println(string(upgradeBytes))
		} else {
			err = utils.WriteOutput(upgradeBytes, opts.OutputFile)
			if err != nil {
				return fmt.Errorf("Failed to write to output file: %s\n", err)
			}
		}

		log.Printf("Successfully upgraded %s from %s to version %s\n", upgrader.GetModelType(), upgrader.GetModelVersion(), upgrader.GetSchemaVersion())

		return nil
	},
}

func ConvertCommand(opts *ConvertOptions) (upgrader upgrading.Upgrader, err error) {
	// Validate inputfile was provided and that is json or yaml
	if opts.InputFile == "" {
		return upgrader, errors.New("Please specify an input file with the -f flag")
	} else {
		if err := utils.IsJsonOrYaml(opts.InputFile); err != nil {
			return upgrader, fmt.Errorf("invalid input file: %s\n", err)
		}
	}

	// Validate outputfile is json or yaml, defaults to stdout
	if opts.OutputFile == "" {
		log.Printf("No output file specified, result will be logged\n")
	} else {
		if err := utils.IsJsonOrYaml(opts.OutputFile); err != nil {
			return upgrader, fmt.Errorf("invalid output file: %s\n", err)
		}
	}

	// Validate version was provided
	if opts.Version == "" {
		return upgrader, errors.New("Please specify a version to convert to with the -v flag")
	}

	// Read the input file
	bytes, err := os.ReadFile(opts.InputFile)
	if err != nil {
		return upgrader, fmt.Errorf("reading input file: %s\n", err)
	}

	// Create Upgrader
	upgrader, err = upgrading.NewUpgrader(bytes, opts.Version)
	if err != nil {
		return upgrader, fmt.Errorf("Failed to create upgrader: %s\n", err)
	}

	// Run the upgrade
	err = upgrader.Upgrade()
	if err != nil {
		return upgrader, fmt.Errorf("Failed to upgrade %s version %s: %s\n", upgrader.GetModelType(), upgrader.GetSchemaVersion(), err)
	}

	return upgrader, nil
}

func init() {
	ConvertCmd.Flags().StringVarP(&opts.InputFile, "file", "f", "", "input file to convert")
	ConvertCmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "output file to write to")
	ConvertCmd.Flags().StringVarP(&opts.Version, "version", "v", "", "version to convert to")
}
