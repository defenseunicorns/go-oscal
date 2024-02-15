package schemas

import "embed"

//go:embed *.json
var Schemas embed.FS

const (
	SCHEMA_PREFIX = "oscal_complete_schema-"
)
