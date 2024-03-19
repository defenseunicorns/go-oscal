type Catalog struct {
	BackMatter *BackMatter  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Controls   *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     *[]Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Metadata   Metadata     `json:"metadata" yaml:"metadata"`
	Params     *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	UUID       string       `json:"uuid" yaml:"uuid"`
}