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

// baseFlags represents command-line flags for the base go-oscal command.
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
	oscalMap, err := parseJson()
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

/*
parseJson reads user-provided oscal json schema files as input,
stores each one of them to a map[string]interface{} pointer,
merges the maps into a single map, and returns the map[string]interface{}.
*/
func parseJson() (map[string]interface{}, error) {
	base := map[string]interface{}{}

	// User specified schema files via -f/--input-file
	for _, filePath := range opts.inputFiles {
		currentMap := map[string]interface{}{}

		bytes, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(bytes, &currentMap); err != nil {
			return nil, err
		}
		// Merge with the previous map
		base = mergeMaps(base, currentMap)
	}

	return base, nil
}

// copy-pasted logic from Helm that merges maps together.
// this will need to be tweaked as it overwrites duplicate values with the last map that's processed
func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
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
