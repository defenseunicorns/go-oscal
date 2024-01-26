type assembly_oscal-catalog_catalog struct {
	BackMatter assembly_oscal-metadata_back-matter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Controls []assembly_oscal-catalog_control `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups []assembly_oscal-catalog_group `json:"groups,omitempty" yaml:"groups,omitempty"`
	Metadata assembly_oscal-metadata_metadata `json:"metadata" yaml:"metadata"`
	Params []assembly_oscal-control-common_parameter `json:"params,omitempty" yaml:"params,omitempty"`
	UUID string `json:"uuid" yaml:"uuid"`
}