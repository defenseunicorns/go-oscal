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

- TODO
    - evaluate back-matter resources key under properties
    - add tags
    - identify omitempty for tags
        - Review objects for "required" fields 
        - have the ability to establish omitempty on all fields that are not required?
    - add comments to the top of the generated file to indicate it being generated and a reference to go-oscal?
        - add sub-comment to recommend how to make manual edits - this should aid in instances where users need to track modifications made to the generated types file (particularly during code review)
    - Enhance logging and errors
    - Abstract the logic to handle other OSCAL models/schemas
    - Enhance structId to account for potential duplicate conflicts
    - Format `integer` to `int` as value type in port range struct