# Notes

- create a subcommand that generates new structs for each oscal data model in the schema


- Latest Notes
    - Likely need to perform a conversion from json schema into a map[string]interface{}
    - Basically having a map that we can begin to traverse recursively in order to:
        - Start with some top-level definition entry - create a name - this will be the struct name
        - Read each property - create a name
            - If property contains a "type" key
                - Evaluate if it is bool/string - if so - 
                - Evaluate arrays - TODO understand what this entails
                    - should check for array.items.type
                        - produces a slice of that type
                    - else check for "$ref"
                        - Would use the id to check for an existing struct type otherwise creating
                    - else error
            - if the property only contains a "$ref"
                - Check for existing struct definition first
                - Need to recurse and return a unique struct name by first finding that object in the map via id

    - Once fully generated - then loop over provided types and begin producing file
