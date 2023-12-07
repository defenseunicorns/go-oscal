package cmd

import (
	"io"
	"log"
	"os"

	"github.com/defenseunicorns/go-oscal/src/cmd/convert"
	"github.com/defenseunicorns/go-oscal/src/cmd/generate"
	"github.com/defenseunicorns/go-oscal/src/cmd/validate"
	"github.com/defenseunicorns/go-oscal/src/internal/utils"
	"github.com/spf13/cobra"
)

var logFile string // -l, --log-file

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-oscal",
	Args:  cobra.ExactArgs(0),
	Short: "Generate Go data types from OSCAL JSON schema",
	Long:  "Generate Go data types from OSCAL JSON schema",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetOutput(os.Stderr)
		log.SetPrefix("go-oscal: ")
		file, err := utils.ReadLogFile(logFile)
		if err != nil {
			log.Fatalf("failed to create log file: %s\n", err)
		}
		log.SetOutput(io.MultiWriter(os.Stderr, file))
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&logFile, "log-file", "l", "", "the name of the file to write the log to (outputs to STDERR by default)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	commands := []*cobra.Command{
		generate.GenerateCmd,
		convert.ConvertCmd,
		validate.ValidateCmd,
	}

	RootCmd.AddCommand(commands...)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
