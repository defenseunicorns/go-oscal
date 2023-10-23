package cmd

import (
	"log"

	"github.com/defenseunicorns/go-oscal/src/cmd/convert"
	"github.com/defenseunicorns/go-oscal/src/cmd/generate"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-oscal",
	Args:  cobra.ExactArgs(0),
	Short: "Generate Go data types from OSCAL JSON schema",
	Long:  "Generate Go data types from OSCAL JSON schema",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	commands := []*cobra.Command{
		generate.GenerateCommand(),
		convert.ConvertCommand(),
	}

	rootCmd.AddCommand(commands...)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
