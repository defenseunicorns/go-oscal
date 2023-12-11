package upgrade

import (
	"errors"
	"fmt"
	"os"

	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/spf13/cobra"
)

type UpgradeCmdOptions struct {
	InputFile  string `short:"f" long:"file" description:"Path to the OSCAL document to upgrade" required:"true"`
	OutPutFile string `short:"o" long:"output" description:"Path to the OSCAL document to output" required:"false"`
	Version    string `short:"v" long:"version" description:"Version to upgrade to" required:"true"`
}

var opts UpgradeCmdOptions

var UpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade an oscal document",
	Long:  "Runs validation of the version provided against the Oscal Document Provided. Returns list of fields that need changes before upgrading to the desired version.",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := UpgradeCmdCommand(&opts)
		if err != nil {
			return err
		}
		return nil
	},
}

func UpgradeCmdCommand(opts *UpgradeCmdOptions) (err error) {
	// Validate the input file
	if opts.InputFile == "" {
		return errors.New("Please specify an input file with the -f flag")
	}

	if opts.Version == "" {
		return errors.New("Please specify a version with the -v flag")
	}

	if err := utils.IsJsonOrYaml(opts.InputFile); err != nil {
		return err
	}

	// Read the input file
	bytes, err := os.ReadFile(opts.InputFile)
	if err != nil {
		return fmt.Errorf("reading input file: %s\n", err)
	}

	// Create Validator
	validator, err := validation.NewValidatorDesiredVersion(bytes, opts.Version)
	if err != nil {
		return fmt.Errorf("creating validator: %s\n", err)
	}

	// Validate the input file
	if err := validator.Validate(); err != nil {
		return fmt.Errorf("validating input file: %s\n", err)
	}

	// TODO: Upgrade the input file

	return nil
}

func init() {
	UpgradeCmd.Flags().StringVarP(&opts.InputFile, "file", "f", "", "Path to the OSCAL document to upgrade")
	UpgradeCmd.Flags().StringVarP(&opts.OutPutFile, "output", "o", "", "Path to the OSCAL document to output")
	UpgradeCmd.Flags().StringVarP(&opts.Version, "version", "v", "", "Version to upgrade to")
}
