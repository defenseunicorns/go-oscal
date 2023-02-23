package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/defenseunicorns/go-oscal/src/internal/oscal"

	"github.com/spf13/cobra"
)

// baseFlags represent command-line flags for the base go-oscal command.
type baseFlags struct {
	inputFiles []string // -f / --input-file
	outputFile string   // -o / --output-file
	pkg        string   // -p / --pkg
	tags       string   // -t / --tags
}

var opts = &baseFlags{}

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
	rootCmd.Flags().StringSliceVarP(&opts.inputFiles, "input-file", "f", []string{""}, "the path to a oscal json schema file")
	rootCmd.Flags().StringVarP(&opts.outputFile, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	rootCmd.Flags().StringVarP(&opts.pkg, "pkg", "p", "main", "the name of the package for the generated code")
	rootCmd.Flags().StringVarP(&opts.tags, "tags", "t", "json", "comma seperated list of the tags to put on the struct")
}

func run() error {
	oscalMap, err := parseOscalInterface()
	if err != nil {
		return err
	}

	tagList := formatTags()

	// Generate the Go structs.
	output := generateStructs(oscalMap, opts.pkg, tagList)

	// Write the Go struct output to either stdout or a file.
	writeOutput(output)

	return nil
}

// parseOscalInterface parses an interface containing oscal data
// to determine what type it is, and ensures we get a map[string]interface{}.
func parseOscalInterface() (map[string]interface{}, error) {
	var oscalMap map[string]interface{}

	interfaceResult, err := parseJson()
	if err != nil {
		return nil, err
	}

	// Parse input interface and perform necessary data transformations.
	switch interfaceResultType := interfaceResult.(type) {
	case map[interface{}]interface{}:
		oscalMap = oscal.ConvertKeysToStrings(interfaceResultType)
	case map[string]interface{}:
		oscalMap = interfaceResultType
	default:
		return nil, fmt.Errorf("unexpected type: %T", interfaceResult)
	}

	return oscalMap, nil
}

// parseJson reads user-provided oscal json schema files as input,
// stores them to an interface pointer, and returns the interface.
func parseJson() (interface{}, error) {
	var bytes []byte
	var err error
	var result interface{}

	// User specified schema files via -f/--input-file
	for _, filePath := range opts.inputFiles {
		bytes, err = os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
	}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	// TODO: We may need to unmarshal the oscal json data directly into a map[string]interface{} structure.
	// I don't think we will run into a situation where the data won't be in this format.
	// If so, there are other ways to check/validate that this is the case.
	// So we can probably skip unmarshalling the data into an interface{}, and then parsing it with parseOscalInterface().
	// The purpose of this would be to allow getting the data into the map[string]interface{} format,
	// and then merging the maps into one monolithic map to be able to process multiple oscal schemas at once.

	return result, nil
}

// formatTags formats Go struct tags.
func formatTags() (tagList []string) {
	if opts.tags == "" {
		tagList = append(tagList, "json")
	} else {
		tagList = strings.Split(opts.tags, ",")
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
	if opts.outputFile != "" {
		err := os.WriteFile(opts.outputFile, output, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	} else {
		fmt.Print(string(output))
	}
}
