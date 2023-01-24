# Notes

- create a subcommand that generates new structs for each oscal data model in the schema

- TODO
    - add comments to the top of the generated file to indicate it being generated and a reference to go-oscal?
        - add sub-comment to recommend how to make manual edits - this should aid in instances where users need to track modifications made to the generated types file (particularly during code review)
    - Enhance logging and errors
    - Abstract the logic to handle other OSCAL models/schemas