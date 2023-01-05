# Notes

The core logic/functionality as of this commit is based on a modified version of [gojson](https://github.com/ChimeraCoder/gojson).

There is a quite a bit more modification/customization that will probably be required to seamlessly generate Go structs for OSCAL data types, so using the `gojson` tool directly is likely not a great choice. It doesn't support generating structs from XML (a supported OSCAL data format), which is a capability that we should aim to have at some point. There also hasn't been a commit made to the repository in 3.5 years, which is a concern in regards to security and reliability due to lack of active maintenance/development.

There are multiple one-off, go modules available that are designed to generate Go structs from JSON schemas.

A few others looked at so far:

- [generate](https://github.com/a-h/generate.git)

- [go-jsonschema](https://github.com/atombender/go-jsonschema.git)

- [go-json-schema-tools](https://github.com/schorsch/go-json-schema-tools.git)

The tools listed above do not support and/or work with the [JSON schema specification](https://github.com/usnistgov/OSCAL/blob/d3a2b990e24210c253642451e30ea6db99bd045b/json/schema/oscal_component_schema.json#L2) that the latest version (v1.0.4) of OSCAL uses.
