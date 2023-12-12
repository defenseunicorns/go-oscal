package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/upgrader"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
		if opts.InputFile == "" {
			return errors.New("Please specify an input file with the -f flag")
		}
		if opts.Version == "" {
			return errors.New("Please specify a version to convert to with the -v flag")
		}
		if opts.OutputFile == "" {
			log.Printf("No output file specified, writing to stdout")
		}

		// Read the input file
		bytes, err := os.ReadFile(opts.InputFile)
		if err != nil {
			return fmt.Errorf("reading input file: %s\n", err)
		}

		// Create Upgrader
		upgrader, err := upgrader.NewUpgrader(bytes, opts.Version)
		if err != nil {
			return fmt.Errorf("Failed to create upgrader: %s\n", err)
		}

		err = upgrader.Upgrade()
		if err != nil {
			return fmt.Errorf("Failed to upgrade %s version %s: %s\n", upgrader.GetModelType(), upgrader.GetSchemaVersion(), err)
		}

		log.Printf("Successfully upgraded %s to version %s\n", opts.InputFile, opts.Version)

		var upgradeBytes []byte
		if strings.HasSuffix(opts.OutputFile, "json") {
			upgradeBytes, err = json.Marshal(upgrader.GetUpgradedJsonMap())
			if err != nil {
				return fmt.Errorf("Failed to marshal upgraded json map: %s\n", err)
			}
		} else {
			upgradeBytes, err = yaml.Marshal(upgrader.GetUpgradedJsonMap())
			if err != nil {
				return fmt.Errorf("Failed to marshal upgraded json map: %s\n", err)
			}
		}

		if opts.OutputFile == "" {
			log.Println(string(upgradeBytes))
		} else {
			err = utils.WriteOutput(upgradeBytes, opts.OutputFile)
			if err != nil {
				return fmt.Errorf("Failed to write to output file: %s\n", err)
			}
		}

		return nil
	},
}

func init() {
	ConvertCmd.Flags().StringVarP(&opts.InputFile, "file", "f", "", "input file to convert")
	ConvertCmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "output file to write to")
	ConvertCmd.Flags().StringVarP(&opts.Version, "version", "v", "", "version to convert to")
}
