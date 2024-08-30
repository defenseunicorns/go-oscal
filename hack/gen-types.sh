#!/bin/bash

# Set the source and destination directories
SCHEMA_DIR="src/internal/schemas"
OUTPUT_DIR="src/types/quicktype"

# Loop through all JSON files in the schema directory
for schema_file in "$SCHEMA_DIR"/oscal_complete_schema-*.json; do
    # Extract the version number from the filename
    version=$(echo "$schema_file" | sed -n 's/.*oscal_complete_schema-\(.*\)\.json/\1/p' | tr '.' '-')
    
    # Create the output directory if it doesn't exist
    mkdir -p "$OUTPUT_DIR/oscal-$version"
    
    # Generate the Go types using quicktype
    cat "$schema_file" | npx quicktype -s schema -o "$OUTPUT_DIR/oscal-$version/types.go" --package "oscalTypes_${version//-/_}" --top-level OscalModels    
    echo "Generated types for OSCAL version $version"
done

echo "Type generation complete"
