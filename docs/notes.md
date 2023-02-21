# Notes

- TODO
    - Enhance logging and errors
    - Abstract the logic to handle other OSCAL models/schemas

Questions:

- To process multiple schemas, should we accept multiple json schema files as input, convert them to maps, and merge the maps into a single monolithic map to parse? Is there a better/simpler, more efficient way to handle this? There is also a [monolithic oscal json schema file](https://github.com/usnistgov/OSCAL/blob/main/json/schema/oscal_complete_schema.json) that contains all models that we could consider using instead of multiple, individual schema files. I could see this option potentially being more simple versus figuring out how to either process multiple maps or merge multiple maps into a single one. If we read one monolithic file that we read into memory, it would give one monolithic map to process and iterate over. Going to experiment and do further testing.