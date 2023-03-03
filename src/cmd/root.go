package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/oscal"

	"github.com/spf13/cobra"
)

var opts = &oscal.BaseFlags{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-oscal",
	Args:  cobra.ExactArgs(0),
	Short: "Generate Go data types from OSCAL JSON schema",
	Long:  "Generate Go data types from OSCAL JSON schema",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to a oscal json schema file")
	rootCmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	rootCmd.Flags().StringVarP(&opts.Pkg, "pkg", "p", "main", "the name of the package for the generated code")
	rootCmd.Flags().StringVarP(&opts.Tags, "tags", "t", "json", "comma seperated list of the tags to put on the struct")
}

func run() error {
	oscalMap, err := oscal.ParseJson(opts)
	if err != nil {
		return err
	}

	tagList := formatTags()

	// Generate the Go structs.
	output := generateStructs(oscalMap, opts.Pkg, tagList)

	// Write the Go struct output to either stdout or a file.
	writeOutput(output)

	return nil
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

// generateStructs reads input and generates Go structs from a OSCAL JSON schema file.
func generateStructs(oscalMap map[string]interface{}, pkg string, tagList []string) (output []byte) {
	output, err := oscal.Generate(oscalMap, pkg, tagList)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	}

	return
}

// writeOutput writes the generated Go structs to an output file
// if provided via -o/--output-file or stdout by default.
func writeOutput(output []byte) {
	if opts.OutputFile != "" {
		err := os.WriteFile(opts.OutputFile, output, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(string(output))
	}
}
