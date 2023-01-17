package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/internal/oscal"

	"github.com/spf13/cobra"
)

var (
	name           string
	pkg            string
	inputFileName  string
	outputFileName string
	format         string
	tags           string
	subStruct      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-oscal",
	Args:  cobra.MaximumNArgs(0),
	Short: "Generate Go data types from OSCAL JSON schema",
	Long:  "Generate Go data types from OSCAL JSON schema",
	Run: func(cmd *cobra.Command, args []string) {
		generateSchemaStructs()
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
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "Example", "the name of the struct")
	rootCmd.PersistentFlags().StringVarP(&pkg, "pkg", "p", "main", "the name of the package for the generated code")
	rootCmd.Flags().StringVarP(&inputFileName, "input-file", "i", "", "the name of the input file containing JSON (if input not provided via STDIN)")
	rootCmd.PersistentFlags().StringVarP(&outputFileName, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	rootCmd.PersistentFlags().StringVar(&format, "fmt", "json", "the format of the input data (json or yaml)")
	rootCmd.PersistentFlags().StringVar(&tags, "tags", format, "comma seperated list of the tags to put on the struct, default is the same as fmt")
	rootCmd.PersistentFlags().BoolVar(&subStruct, "sub-struct", false, "create types for sub-structs")

	rootCmd.MarkFlagRequired("input-file")
}

func generateSchemaStructs() {
	input, convertFloats, parser, tagList := validateInput()

	output, err := oscal.Generate(input, parser, name, pkg, tagList, subStruct, convertFloats)
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

func validateInput() (io.Reader, bool, oscal.Parser, []string) {
	if format != "json" && format != "yaml" {
		fmt.Fprintln(os.Stderr, "fmt must be json or yaml")
		os.Exit(1)
	}

	tagList := make([]string, 0)
	if tags == "" || tags == format {
		tagList = append(tagList, format)
	} else {
		tagList = strings.Split(tags, ",")
	}

	var input io.Reader
	input = os.Stdin
	if inputFileName != "" {
		f, err := os.Open(inputFileName)
		if err != nil {
			log.Fatalf("reading input file: %s", err)
		}
		input = f
	}

	var convertFloats bool
	var parser oscal.Parser
	switch format {
	case "json":
		parser = oscal.ParseJson
		convertFloats = true
	case "yaml":
		parser = oscal.ParseYaml
	}

	return input, convertFloats, parser, tagList
}
