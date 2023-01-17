package oscal

type OscalComponentDefinition struct {
	Back_matter                  Back_matter_1  `json:"back-matter" yaml:"back-matter"`
	Capabilities                 Capabilities_2 `json:"capabilities" yaml:"capabilities"`
	Components                   Capabilities_2 `json:"components" yaml:"components"`
	Import_component_definitions Capabilities_2 `json:"import-component-definitions" yaml:"import-component-definitions"`
	Metadata                     Back_matter_1  `json:"metadata" yaml:"metadata"`
	UUID                         UUID_3         `json:"uuid" yaml:"uuid"`
}

type UUID_3 struct {
	Description string `json:"description" yaml:"description"`
	Pattern     string `json:"pattern" yaml:"pattern"`
	Title       string `json:"title" yaml:"title"`
	Type        string `json:"type" yaml:"type"`
}

type Capabilities_2 struct {
	Items    Back_matter_1 `json:"items" yaml:"items"`
	MinItems int64         `json:"minItems" yaml:"minItems"`
	Type     string        `json:"type" yaml:"type"`
}

type Back_matter_1 struct {
	Ref string `json:"$ref" yaml:"$ref"`
}
