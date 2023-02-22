package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/oscal"

	"github.com/spf13/cobra"
)

var (
	pkg            string
	inputFileName  string
	outputFileName string
	format         string
	tags           string
)

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
	rootCmd.Flags().StringVarP(&pkg, "pkg", "p", "main", "the name of the package for the generated code")
	rootCmd.Flags().StringVarP(&inputFileName, "input-file", "i", "", "the name of the input file containing JSON (if input not provided via STDIN)")
	rootCmd.Flags().StringVarP(&outputFileName, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	rootCmd.Flags().StringVar(&format, "fmt", "json", "the format of the input data (json or yaml)")
	rootCmd.Flags().StringVar(&tags, "tags", format, "comma seperated list of the tags to put on the struct, default is the same as fmt")
}

func run() {
	checkForEmptyInput()
	checkInputDataFormat()
	tagList := formatTags()

	// Read provided input.
	var input = os.Stdin
	if inputFileName != "" {
		f, err := os.Open(inputFileName)
		if err != nil {
			log.Fatalf("reading input file: %s", err)
		}
		defer f.Close()
		input = f
	}

	// Parse the input.
	parser, convertFloats := parseInput()

	// Generate the Go structs.
	output := generateStructs(input, parser, pkg, tagList, convertFloats)

	// Write the Go struct output to either stdout or a file.
	writeOutput(output)
}

// checkforEmptyInput checks whether the user provided any input or not.
func checkForEmptyInput() {
	// If there was no input provided, exit and notify the user to provide input.
	stats, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalf("reading STDIN: %s", err)
	}
	if stats.Size() == 0 && inputFileName == "" {
		fmt.Println("You must provide input either via STDIN or the --input-file flag.")
		fmt.Println("Use the --help flag to see available flags.")
		os.Exit(1)
	}
}

// checkInputDataFormat checks if the input data format is json or yaml.
// If it isn't, it exits the program.
func checkInputDataFormat() {
	if format != "json" && format != "yaml" {
		fmt.Fprintln(os.Stderr, "fmt must be json or yaml")
		os.Exit(1)
	}
}

// formatTags formats Go struct tags.
func formatTags() (tagList []string) {
	if tags == "" || tags == format {
		tagList = append(tagList, format)
	} else {
		tagList = strings.Split(tags, ",")
	}

	return
}

// parseInput determines which parsing method we will use
// based on the format of the input data.
func parseInput() (parser oscal.Parser, convertFloats bool) {
	switch format {
	case "json":
		parser = oscal.ParseJson
		convertFloats = true
	case "yaml":
		parser = oscal.ParseYaml
	}

	return
}

// generateStructs reads input and generates Go structs from a OSCAL JSON schema file.
func generateStructs(input io.Reader, parser oscal.Parser, pkg string, tagList []string, convertFloats bool) (output []byte) {
	output, err := oscal.Generate(input, parser, pkg, tagList, convertFloats)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	}

	return
}

// writeOutput writes the generated Go structs to either a file if provided,
// or stdout by default.
func writeOutput(output []byte) {
	if outputFileName != "" {
		err := os.WriteFile(outputFileName, output, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(string(output))
	}
}
