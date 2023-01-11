# Notes

## Context

The core logic/functionality as of this commit is based on a modified version of [gojson](https://github.com/ChimeraCoder/gojson).

There are multiple one-off, go modules available that are designed to generate Go structs from JSON schemas.

A few others looked at so far:

- [generate](https://github.com/a-h/generate.git)

- [go-jsonschema](https://github.com/atombender/go-jsonschema.git)

- [go-json-schema-tools](https://github.com/schorsch/go-json-schema-tools.git)

The tools listed above do not support and/or work with the [JSON schema specification](https://github.com/usnistgov/OSCAL/blob/d3a2b990e24210c253642451e30ea6db99bd045b/json/schema/oscal_component_schema.json#L2) that the latest version (v1.0.4) of OSCAL uses.

## Questions

- Is it necessary to generate Go structs from both OSCAL JSON *and* XML schemas? Or are the schemas "interchangeable", i.e. there would only be a need to generate from one schema format and the structs could be used for both data formats?

- Do we want/need to narrow the scope of the generated structs? The exposed/provided types generated from the schema could be limited to exposing only specific portions of the schema that we need to use. What are the trade-offs? Will it greatly increase code complexity?
