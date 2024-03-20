# go-oscal

[![Go Reference](https://pkg.go.dev/badge/github.com/defenseunicorns/go-oscal.svg)](https://pkg.go.dev/github.com/defenseunicorns/go-oscal)
[![Go Report Card](https://goreportcard.com/badge/github.com/defenseunicorns/go-oscal)](https://goreportcard.com/report/github.com/defenseunicorns/go-oscal)
[![License](https://img.shields.io/github/license/defenseunicorns/go-oscal)](https://github.com/defenseunicorns/go-oscal/blob/main/LICENSE)

go-oscal is a Go library for working with the Open Security Controls Assessment Language (OSCAL). It provides functionality to parse, validate, manipulate, and generate OSCAL content in JSON and YAML formats.

## Table Of Contents

- [go-oscal](#go-oscal)
  - [Table Of Contents](#table-of-contents)
  - [Usage](#usage)
    - [CLI](#cli)
      - [Commands](#commands)
    - [Import](#import)
      - [Using Types](#using-types)
  - [Additional Resources and Projects](#additional-resources-and-projects)

## Usage
### CLI
- Clone the repository and change into the `go-oscal` directory
- Build the CLI
  
```bash
go build .
```
#### Commands
- [root](./docs/commands/root.md)
- [validate](./docs/commands/validate.md)
- [revise](./docs/commands/revise.md)
- [doctor](./docs/commands/doctor.md)
- [generate](./docs/commands/generate.md)


### Import
```bash
go get github.com/defenseunicorns/go-oscal
```

#### Using Types
```golang
// Types can be imported by their version
oscalTypes_1_1_2 "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-2"

result := oscalTypes_1_1_2.Result{
  Findings: []oscalTypes_1_1_2.Finding{
    {
      Target: oscalTypes_1_1_2.FindingTarget{
        TargetId: "ID-1",
        Status: oscalTypes_1_1_2.ObjectiveStatus{
          State: "satisfied",
        },
      },
    },
  },
}
```
## Additional Resources and Projects
- [lula](https://github.com/defenseunicorns/lula)
- [OSCAL](https://github.com/usnistgov/OSCAL)
- [fedramp automation](https://github.com/GSA/fedramp-automation)
- [Awesome Oscal](https://github.com/oscal-club/awesome-oscal)
