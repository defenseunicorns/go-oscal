# 2. Use pointers for non-primitive types

Date: 2024-03-25

## Status

Accepted

## Context

Initially, types were generated without pointers. When using `encoding/json`'s `json.Marshal` and `json.Unmarshal`, all fields, even those tagged with `omitempty`, were instantiated with empty values. This can cause problems when reading, editing, and writing `json` objects, as they no longer reflect the initial and intended state of the original `json`. This issue was not present in `yaml` documents using `gopkg.in/yaml.v3`. More context is available in this [proposal](https://github.com/golang/go/issues/11939).

## Decision

To better support consumers who implement OSCAL and the go-oscal types in `json`, we have decided that embedded `struct` objects and non-primitives in the go-oscal types should be pointers. This is the most conventional solution to prevent the instantiation of unused (`omitempty`) `json` fields. A more in-depth explanation of the `json` `omitempty` tag can be found [here](https://www.sohamkamani.com/golang/omitempty/).

## Consequences
- **Nil Checks**: Additional `nil` checks are required before accessing methods or fields on embedded structs.
- **Initialization**: Pointers to embedded structs must be properly initialized before use.
- **Interface Compatibility**: Interfaces that expect embedded structs may need to be adjusted to work with pointers.
- **Serialization and Deserialization**: Correct handling of pointers is required to ensure proper serialization and deserialization.
- **Memory Management**: Using pointers may increase memory usage compared to embedded structs.
