package oscal

type OscalComponentDefinition struct {
	Comment              string          `json:"$comment"`
	ID                   string          `json:"$id"`
	Schema               string          `json:"$schema"`
	AdditionalProperties bool            `json:"additionalProperties"`
	Definitions          Definitions_108 `json:"definitions"`
	MaxProperties        int64           `json:"maxProperties"`
	Properties           Properties_109  `json:"properties"`
	Required             []string        `json:"required"`
	Type                 string          `json:"type"`
}

type Items_45 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_44 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Status_58 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_57 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Base64_69 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_68 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Citation_71 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_70 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Items_73 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_72 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Items_76 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_75 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Items_9 struct {
	AdditionalProperties bool         `json:"additionalProperties"`
	Description          string       `json:"description"`
	Properties           Properties_8 `json:"properties"`
	Required             []string     `json:"required"`
	Title                string       `json:"title"`
	Type                 string       `json:"type"`
}

type Items_93 struct {
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_92 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Properties_66 struct {
	Addr_lines  Constraints_4 `json:"addr-lines"`
	City        Class_2       `json:"city"`
	Country     Class_2       `json:"country"`
	Postal_code Class_2       `json:"postal-code"`
	State       Class_2       `json:"state"`
	Type        Class_2       `json:"type"`
}

type Properties_88 struct {
	Address           Items_3       `json:"address"`
	Email_addresses   Constraints_4 `json:"email-addresses"`
	Links             Constraints_4 `json:"links"`
	Props             Constraints_4 `json:"props"`
	Remarks           Items_3       `json:"remarks"`
	Telephone_numbers Constraints_4 `json:"telephone-numbers"`
	Title             Label_5       `json:"title"`
	Urls              Urls_87       `json:"urls"`
	UUID              Class_2       `json:"uuid"`
}

type Properties_95 struct {
	Addresses               Constraints_4   `json:"addresses"`
	Email_addresses         Constraints_4   `json:"email-addresses"`
	External_ids            External_ids_94 `json:"external-ids"`
	Links                   Constraints_4   `json:"links"`
	Location_uuids          Constraints_4   `json:"location-uuids"`
	Member_of_organizations Values_54       `json:"member-of-organizations"`
	Name                    Class_2         `json:"name"`
	Props                   Constraints_4   `json:"props"`
	Remarks                 Items_3         `json:"remarks"`
	Short_name              Class_2         `json:"short-name"`
	Telephone_numbers       Constraints_4   `json:"telephone-numbers"`
	Type                    How_many_16     `json:"type"`
	UUID                    Class_2         `json:"uuid"`
}

type Properties_83 struct {
	Algorithm Class_2 `json:"algorithm"`
	Value     ID_61   `json:"value"`
}

type Properties_64 struct {
	Authorized_privileges Constraints_4 `json:"authorized-privileges"`
	Description           Label_5       `json:"description"`
	Links                 Constraints_4 `json:"links"`
	Props                 Constraints_4 `json:"props"`
	Remarks               Items_3       `json:"remarks"`
	Role_ids              Constraints_4 `json:"role-ids"`
	Short_name            Class_2       `json:"short-name"`
	Title                 Label_5       `json:"title"`
	UUID                  Class_2       `json:"uuid"`
}

type Properties_25 struct {
	Back_matter                  Items_3       `json:"back-matter"`
	Capabilities                 Constraints_4 `json:"capabilities"`
	Components                   Constraints_4 `json:"components"`
	Import_component_definitions Constraints_4 `json:"import-component-definitions"`
	Metadata                     Items_3       `json:"metadata"`
	UUID                         Class_2       `json:"uuid"`
}

type Properties_75 struct {
	Base64       Base64_69     `json:"base64"`
	Citation     Citation_71   `json:"citation"`
	Description  Label_5       `json:"description"`
	Document_ids Constraints_4 `json:"document-ids"`
	Props        Constraints_4 `json:"props"`
	Remarks      Items_3       `json:"remarks"`
	Rlinks       Rlinks_74     `json:"rlinks"`
	Title        Label_5       `json:"title"`
	UUID         Class_2       `json:"uuid"`
}

type Properties_17 struct {
	Choice   Choice_15   `json:"choice"`
	How_many How_many_16 `json:"how-many"`
}

type Properties_6 struct {
	Class       Class_2       `json:"class"`
	Constraints Constraints_4 `json:"constraints"`
	Depends_on  Class_2       `json:"depends-on"`
	Guidelines  Constraints_4 `json:"guidelines"`
	ID          Class_2       `json:"id"`
	Label       Label_5       `json:"label"`
	Links       Constraints_4 `json:"links"`
	Props       Constraints_4 `json:"props"`
	Remarks     Items_3       `json:"remarks"`
	Select      Items_3       `json:"select"`
	Usage       Label_5       `json:"usage"`
	Values      Constraints_4 `json:"values"`
}

type Properties_21 struct {
	Class Class_2       `json:"class"`
	ID    Class_2       `json:"id"`
	Links Constraints_4 `json:"links"`
	Name  Class_2       `json:"name"`
	Ns    Ns_20         `json:"ns"`
	Parts Constraints_4 `json:"parts"`
	Props Constraints_4 `json:"props"`
	Prose Label_5       `json:"prose"`
	Title Label_5       `json:"title"`
}

type Properties_97 struct {
	Class   Class_2 `json:"class"`
	Name    Class_2 `json:"name"`
	Ns      Ns_20   `json:"ns"`
	Remarks Items_3 `json:"remarks"`
	UUID    Class_2 `json:"uuid"`
	Value   Class_2 `json:"value"`
}

type Properties_109 struct {
	Component_definition Items_3 `json:"component-definition"`
}

type Properties_36 struct {
	Component_uuid Class_2 `json:"component-uuid"`
	Description    Label_5 `json:"description"`
}

type Properties_44 struct {
	Component_uuid      Class_2       `json:"component-uuid"`
	Links               Constraints_4 `json:"links"`
	Props               Constraints_4 `json:"props"`
	Remarks             Items_3       `json:"remarks"`
	Responsible_parties Constraints_4 `json:"responsible-parties"`
}

type Properties_32 struct {
	Control_id        Class_2       `json:"control-id"`
	Description       Label_5       `json:"description"`
	Links             Constraints_4 `json:"links"`
	Props             Constraints_4 `json:"props"`
	Remarks           Items_3       `json:"remarks"`
	Responsible_roles Constraints_4 `json:"responsible-roles"`
	Set_parameters    Constraints_4 `json:"set-parameters"`
	Statements        Constraints_4 `json:"statements"`
	UUID              Class_2       `json:"uuid"`
}

type Properties_23 struct {
	Control_implementations Constraints_4 `json:"control-implementations"`
	Description             Label_5       `json:"description"`
	Incorporates_components Constraints_4 `json:"incorporates-components"`
	Links                   Constraints_4 `json:"links"`
	Name                    Class_2       `json:"name"`
	Props                   Constraints_4 `json:"props"`
	Remarks                 Items_3       `json:"remarks"`
	UUID                    Class_2       `json:"uuid"`
}

type Properties_30 struct {
	Control_implementations Constraints_4 `json:"control-implementations"`
	Description             Label_5       `json:"description"`
	Links                   Constraints_4 `json:"links"`
	Props                   Constraints_4 `json:"props"`
	Protocols               Constraints_4 `json:"protocols"`
	Purpose                 Label_5       `json:"purpose"`
	Remarks                 Items_3       `json:"remarks"`
	Responsible_roles       Constraints_4 `json:"responsible-roles"`
	Title                   Label_5       `json:"title"`
	Type                    Class_2       `json:"type"`
	UUID                    Class_2       `json:"uuid"`
}

type Properties_40 struct {
	Description         Label_5       `json:"description"`
	Functions_performed Constraints_4 `json:"functions-performed"`
	Title               Label_5       `json:"title"`
}

type Properties_104 struct {
	Description Label_5       `json:"description"`
	ID          Class_2       `json:"id"`
	Links       Constraints_4 `json:"links"`
	Props       Constraints_4 `json:"props"`
	Remarks     Items_3       `json:"remarks"`
	Short_name  Class_2       `json:"short-name"`
	Title       Label_5       `json:"title"`
}

type Properties_47 struct {
	Description            Label_5                   `json:"description"`
	Implemented_components Implemented_components_46 `json:"implemented-components"`
	Links                  Constraints_4             `json:"links"`
	Props                  Constraints_4             `json:"props"`
	Remarks                Items_3                   `json:"remarks"`
	Responsible_parties    Constraints_4             `json:"responsible-parties"`
	UUID                   Class_2                   `json:"uuid"`
}

type Properties_28 struct {
	Description              Label_5       `json:"description"`
	Implemented_requirements Constraints_4 `json:"implemented-requirements"`
	Links                    Constraints_4 `json:"links"`
	Props                    Constraints_4 `json:"props"`
	Set_parameters           Constraints_4 `json:"set-parameters"`
	Source                   Source_27     `json:"source"`
	UUID                     Class_2       `json:"uuid"`
}

type Properties_59 struct {
	Description       Label_5       `json:"description"`
	Links             Constraints_4 `json:"links"`
	Props             Constraints_4 `json:"props"`
	Protocols         Constraints_4 `json:"protocols"`
	Purpose           Label_5       `json:"purpose"`
	Remarks           Items_3       `json:"remarks"`
	Responsible_roles Constraints_4 `json:"responsible-roles"`
	Status            Status_58     `json:"status"`
	Title             Label_5       `json:"title"`
	Type              Class_2       `json:"type"`
	UUID              Class_2       `json:"uuid"`
}

type Properties_38 struct {
	Description       Label_5       `json:"description"`
	Links             Constraints_4 `json:"links"`
	Props             Constraints_4 `json:"props"`
	Remarks           Items_3       `json:"remarks"`
	Responsible_roles Constraints_4 `json:"responsible-roles"`
	Statement_id      Class_2       `json:"statement-id"`
	UUID              Class_2       `json:"uuid"`
}

type Properties_11 struct {
	Description Label_5  `json:"description"`
	Tests       Tests_10 `json:"tests"`
}

type How_many_16 struct {
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
	Pattern     string   `json:"pattern"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
}

type Ns_20 struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Source_27 struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type End_49 struct {
	Description string `json:"description"`
	Minimum     int64  `json:"minimum"`
	MultipleOf  int64  `json:"multipleOf"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Class_2 struct {
	Description string `json:"description"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Label_5 struct {
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Properties_90 struct {
	Document_ids        Constraints_4 `json:"document-ids"`
	Last_modified       Items_3       `json:"last-modified"`
	Links               Constraints_4 `json:"links"`
	Locations           Constraints_4 `json:"locations"`
	Oscal_version       Items_3       `json:"oscal-version"`
	Parties             Constraints_4 `json:"parties"`
	Props               Constraints_4 `json:"props"`
	Published           Items_3       `json:"published"`
	Remarks             Items_3       `json:"remarks"`
	Responsible_parties Constraints_4 `json:"responsible-parties"`
	Revisions           Constraints_4 `json:"revisions"`
	Roles               Constraints_4 `json:"roles"`
	Title               Label_5       `json:"title"`
	Version             Items_3       `json:"version"`
}

type Properties_50 struct {
	End       End_49      `json:"end"`
	Start     End_49      `json:"start"`
	Transport How_many_16 `json:"transport"`
}

type Properties_8 struct {
	Expression Class_2 `json:"expression"`
	Remarks    Items_3 `json:"remarks"`
}

type Properties_68 struct {
	Filename   Source_27 `json:"filename"`
	Media_type Class_2   `json:"media-type"`
	Value      ID_61     `json:"value"`
}

type Properties_72 struct {
	Hashes     Constraints_4 `json:"hashes"`
	Href       Source_27     `json:"href"`
	Media_type Class_2       `json:"media-type"`
}

type Properties_85 struct {
	Href       Source_27 `json:"href"`
	Media_type Class_2   `json:"media-type"`
	Rel        Class_2   `json:"rel"`
	Text       Label_5   `json:"text"`
}

type Properties_34 struct {
	Href Source_27 `json:"href"`
}

type Properties_62 struct {
	ID              ID_61 `json:"id"`
	Identifier_type Ns_20 `json:"identifier-type"`
}

type Properties_92 struct {
	ID     ID_61 `json:"id"`
	Scheme Ns_20 `json:"scheme"`
}

type Oscal_component_definition_oscal_metadata_responsible_party_101 struct {
	ID                   string         `json:"$id"`
	AdditionalProperties bool           `json:"additionalProperties"`
	Description          string         `json:"description"`
	Properties           Properties_100 `json:"properties"`
	Required             []string       `json:"required"`
	Title                string         `json:"title"`
	Type                 string         `json:"type"`
}

type Oscal_component_definition_oscal_metadata_revision_103 struct {
	ID                   string         `json:"$id"`
	AdditionalProperties bool           `json:"additionalProperties"`
	Description          string         `json:"description"`
	Properties           Properties_102 `json:"properties"`
	Required             []string       `json:"required"`
	Title                string         `json:"title"`
	Type                 string         `json:"type"`
}

type Oscal_component_definition_oscal_metadata_role_105 struct {
	ID                   string         `json:"$id"`
	AdditionalProperties bool           `json:"additionalProperties"`
	Description          string         `json:"description"`
	Properties           Properties_104 `json:"properties"`
	Required             []string       `json:"required"`
	Title                string         `json:"title"`
	Type                 string         `json:"type"`
}

type Oscal_component_definition_oscal_metadata_telephone_number_107 struct {
	ID                   string         `json:"$id"`
	AdditionalProperties bool           `json:"additionalProperties"`
	Description          string         `json:"description"`
	Properties           Properties_106 `json:"properties"`
	Required             []string       `json:"required"`
	Title                string         `json:"title"`
	Type                 string         `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_parameter_constraint_12 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_11 `json:"properties"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_parameter_guideline_14 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_13 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_parameter_selection_18 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_17 `json:"properties"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_part_22 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_21 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_capability_24 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_23 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_component_definition_26 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_25 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_control_implementation_29 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_28 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_defined_component_31 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_30 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_implemented_requirement_33 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_32 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_import_component_definition_35 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_34 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_incorporates_component_37 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_36 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_component_definition_statement_39 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_38 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_authorized_privilege_41 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_40 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_implementation_status_43 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_42 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_inventory_item_48 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_47 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_port_range_51 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_50 `json:"properties"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_protocol_53 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_52 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_set_parameter_56 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_55 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_system_component_60 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_59 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_parameter_7 struct {
	ID                   string       `json:"$id"`
	AdditionalProperties bool         `json:"additionalProperties"`
	Description          string       `json:"description"`
	Properties           Properties_6 `json:"properties"`
	Required             []string     `json:"required"`
	Title                string       `json:"title"`
	Type                 string       `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_system_id_63 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_62 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_implementation_common_system_user_65 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_64 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_address_67 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_66 `json:"properties"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_back_matter_79 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_78 `json:"properties"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_document_id_81 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_80 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_hash_84 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_83 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_link_86 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_85 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_location_89 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_88 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_metadata_91 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_90 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_party_96 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_95 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_metadata_property_98 struct {
	ID                   string        `json:"$id"`
	AdditionalProperties bool          `json:"additionalProperties"`
	Description          string        `json:"description"`
	Properties           Properties_97 `json:"properties"`
	Required             []string      `json:"required"`
	Title                string        `json:"title"`
	Type                 string        `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_include_all_1 struct {
	ID                   string `json:"$id"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Description          string `json:"description"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
}

type Oscal_component_definition_oscal_metadata_email_address_82 struct {
	ID          string `json:"$id"`
	Description string `json:"description"`
	Format      string `json:"format"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Oscal_component_definition_oscal_catalog_common_parameter_value_19 struct {
	ID          string `json:"$id"`
	Description string `json:"description"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Oscal_component_definition_oscal_metadata_remarks_99 struct {
	ID          string `json:"$id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type Properties_80 struct {
	Identifier ID_61 `json:"identifier"`
	Scheme     Ns_20 `json:"scheme"`
}

type Values_54 struct {
	Items    Class_2 `json:"items"`
	MinItems int64   `json:"minItems"`
	Type     string  `json:"type"`
}

type Constraints_4 struct {
	Items    Items_3 `json:"items"`
	MinItems int64   `json:"minItems"`
	Type     string  `json:"type"`
}

type Implemented_components_46 struct {
	Items    Items_45 `json:"items"`
	MinItems int64    `json:"minItems"`
	Type     string   `json:"type"`
}

type Rlinks_74 struct {
	Items    Items_73 `json:"items"`
	MinItems int64    `json:"minItems"`
	Type     string   `json:"type"`
}

type Resources_77 struct {
	Items    Items_76 `json:"items"`
	MinItems int64    `json:"minItems"`
	Type     string   `json:"type"`
}

type Tests_10 struct {
	Items    Items_9 `json:"items"`
	MinItems int64   `json:"minItems"`
	Type     string  `json:"type"`
}

type External_ids_94 struct {
	Items    Items_93 `json:"items"`
	MinItems int64    `json:"minItems"`
	Type     string   `json:"type"`
}

type Choice_15 struct {
	Items    Label_5 `json:"items"`
	MinItems int64   `json:"minItems"`
	Type     string  `json:"type"`
}

type Urls_87 struct {
	Items    Ns_20  `json:"items"`
	MinItems int64  `json:"minItems"`
	Type     string `json:"type"`
}

type Properties_102 struct {
	Last_modified Items_3       `json:"last-modified"`
	Links         Constraints_4 `json:"links"`
	Oscal_version Items_3       `json:"oscal-version"`
	Props         Constraints_4 `json:"props"`
	Published     Items_3       `json:"published"`
	Remarks       Items_3       `json:"remarks"`
	Title         Label_5       `json:"title"`
	Version       Items_3       `json:"version"`
}

type Properties_100 struct {
	Links       Constraints_4 `json:"links"`
	Party_uuids Constraints_4 `json:"party-uuids"`
	Props       Constraints_4 `json:"props"`
	Remarks     Items_3       `json:"remarks"`
	Role_id     Class_2       `json:"role-id"`
}

type Properties_70 struct {
	Links Constraints_4 `json:"links"`
	Props Constraints_4 `json:"props"`
	Text  Label_5       `json:"text"`
}

type Properties_52 struct {
	Name        Class_2       `json:"name"`
	Port_ranges Constraints_4 `json:"port-ranges"`
	Title       Label_5       `json:"title"`
	UUID        Class_2       `json:"uuid"`
}

type Properties_106 struct {
	Number ID_61   `json:"number"`
	Type   Class_2 `json:"type"`
}

type Definitions_108 struct {
	Oscal_component_definition_oscal_catalog_common_include_all                       Oscal_component_definition_oscal_catalog_common_include_all_1                        `json:"oscal-component-definition-oscal-catalog-common:include-all"`
	Oscal_component_definition_oscal_catalog_common_parameter                         Oscal_component_definition_oscal_catalog_common_parameter_7                          `json:"oscal-component-definition-oscal-catalog-common:parameter"`
	Oscal_component_definition_oscal_catalog_common_parameter_constraint              Oscal_component_definition_oscal_catalog_common_parameter_constraint_12              `json:"oscal-component-definition-oscal-catalog-common:parameter-constraint"`
	Oscal_component_definition_oscal_catalog_common_parameter_guideline               Oscal_component_definition_oscal_catalog_common_parameter_guideline_14               `json:"oscal-component-definition-oscal-catalog-common:parameter-guideline"`
	Oscal_component_definition_oscal_catalog_common_parameter_selection               Oscal_component_definition_oscal_catalog_common_parameter_selection_18               `json:"oscal-component-definition-oscal-catalog-common:parameter-selection"`
	Oscal_component_definition_oscal_catalog_common_parameter_value                   Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-catalog-common:parameter-value"`
	Oscal_component_definition_oscal_catalog_common_part                              Oscal_component_definition_oscal_catalog_common_part_22                              `json:"oscal-component-definition-oscal-catalog-common:part"`
	Oscal_component_definition_oscal_component_definition_capability                  Oscal_component_definition_oscal_component_definition_capability_24                  `json:"oscal-component-definition-oscal-component-definition:capability"`
	Oscal_component_definition_oscal_component_definition_component_definition        Oscal_component_definition_oscal_component_definition_component_definition_26        `json:"oscal-component-definition-oscal-component-definition:component-definition"`
	Oscal_component_definition_oscal_component_definition_control_implementation      Oscal_component_definition_oscal_component_definition_control_implementation_29      `json:"oscal-component-definition-oscal-component-definition:control-implementation"`
	Oscal_component_definition_oscal_component_definition_defined_component           Oscal_component_definition_oscal_component_definition_defined_component_31           `json:"oscal-component-definition-oscal-component-definition:defined-component"`
	Oscal_component_definition_oscal_component_definition_implemented_requirement     Oscal_component_definition_oscal_component_definition_implemented_requirement_33     `json:"oscal-component-definition-oscal-component-definition:implemented-requirement"`
	Oscal_component_definition_oscal_component_definition_import_component_definition Oscal_component_definition_oscal_component_definition_import_component_definition_35 `json:"oscal-component-definition-oscal-component-definition:import-component-definition"`
	Oscal_component_definition_oscal_component_definition_incorporates_component      Oscal_component_definition_oscal_component_definition_incorporates_component_37      `json:"oscal-component-definition-oscal-component-definition:incorporates-component"`
	Oscal_component_definition_oscal_component_definition_statement                   Oscal_component_definition_oscal_component_definition_statement_39                   `json:"oscal-component-definition-oscal-component-definition:statement"`
	Oscal_component_definition_oscal_implementation_common_authorized_privilege       Oscal_component_definition_oscal_implementation_common_authorized_privilege_41       `json:"oscal-component-definition-oscal-implementation-common:authorized-privilege"`
	Oscal_component_definition_oscal_implementation_common_function_performed         Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-implementation-common:function-performed"`
	Oscal_component_definition_oscal_implementation_common_implementation_status      Oscal_component_definition_oscal_implementation_common_implementation_status_43      `json:"oscal-component-definition-oscal-implementation-common:implementation-status"`
	Oscal_component_definition_oscal_implementation_common_inventory_item             Oscal_component_definition_oscal_implementation_common_inventory_item_48             `json:"oscal-component-definition-oscal-implementation-common:inventory-item"`
	Oscal_component_definition_oscal_implementation_common_port_range                 Oscal_component_definition_oscal_implementation_common_port_range_51                 `json:"oscal-component-definition-oscal-implementation-common:port-range"`
	Oscal_component_definition_oscal_implementation_common_protocol                   Oscal_component_definition_oscal_implementation_common_protocol_53                   `json:"oscal-component-definition-oscal-implementation-common:protocol"`
	Oscal_component_definition_oscal_implementation_common_set_parameter              Oscal_component_definition_oscal_implementation_common_set_parameter_56              `json:"oscal-component-definition-oscal-implementation-common:set-parameter"`
	Oscal_component_definition_oscal_implementation_common_system_component           Oscal_component_definition_oscal_implementation_common_system_component_60           `json:"oscal-component-definition-oscal-implementation-common:system-component"`
	Oscal_component_definition_oscal_implementation_common_system_id                  Oscal_component_definition_oscal_implementation_common_system_id_63                  `json:"oscal-component-definition-oscal-implementation-common:system-id"`
	Oscal_component_definition_oscal_implementation_common_system_user                Oscal_component_definition_oscal_implementation_common_system_user_65                `json:"oscal-component-definition-oscal-implementation-common:system-user"`
	Oscal_component_definition_oscal_metadata_addr_line                               Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-metadata:addr-line"`
	Oscal_component_definition_oscal_metadata_address                                 Oscal_component_definition_oscal_metadata_address_67                                 `json:"oscal-component-definition-oscal-metadata:address"`
	Oscal_component_definition_oscal_metadata_back_matter                             Oscal_component_definition_oscal_metadata_back_matter_79                             `json:"oscal-component-definition-oscal-metadata:back-matter"`
	Oscal_component_definition_oscal_metadata_document_id                             Oscal_component_definition_oscal_metadata_document_id_81                             `json:"oscal-component-definition-oscal-metadata:document-id"`
	Oscal_component_definition_oscal_metadata_email_address                           Oscal_component_definition_oscal_metadata_email_address_82                           `json:"oscal-component-definition-oscal-metadata:email-address"`
	Oscal_component_definition_oscal_metadata_hash                                    Oscal_component_definition_oscal_metadata_hash_84                                    `json:"oscal-component-definition-oscal-metadata:hash"`
	Oscal_component_definition_oscal_metadata_last_modified                           Oscal_component_definition_oscal_metadata_email_address_82                           `json:"oscal-component-definition-oscal-metadata:last-modified"`
	Oscal_component_definition_oscal_metadata_link                                    Oscal_component_definition_oscal_metadata_link_86                                    `json:"oscal-component-definition-oscal-metadata:link"`
	Oscal_component_definition_oscal_metadata_location                                Oscal_component_definition_oscal_metadata_location_89                                `json:"oscal-component-definition-oscal-metadata:location"`
	Oscal_component_definition_oscal_metadata_location_uuid                           Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-metadata:location-uuid"`
	Oscal_component_definition_oscal_metadata_metadata                                Oscal_component_definition_oscal_metadata_metadata_91                                `json:"oscal-component-definition-oscal-metadata:metadata"`
	Oscal_component_definition_oscal_metadata_oscal_version                           Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-metadata:oscal-version"`
	Oscal_component_definition_oscal_metadata_party                                   Oscal_component_definition_oscal_metadata_party_96                                   `json:"oscal-component-definition-oscal-metadata:party"`
	Oscal_component_definition_oscal_metadata_party_uuid                              Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-metadata:party-uuid"`
	Oscal_component_definition_oscal_metadata_property                                Oscal_component_definition_oscal_metadata_property_98                                `json:"oscal-component-definition-oscal-metadata:property"`
	Oscal_component_definition_oscal_metadata_published                               Oscal_component_definition_oscal_metadata_email_address_82                           `json:"oscal-component-definition-oscal-metadata:published"`
	Oscal_component_definition_oscal_metadata_remarks                                 Oscal_component_definition_oscal_metadata_remarks_99                                 `json:"oscal-component-definition-oscal-metadata:remarks"`
	Oscal_component_definition_oscal_metadata_responsible_party                       Oscal_component_definition_oscal_metadata_responsible_party_101                      `json:"oscal-component-definition-oscal-metadata:responsible-party"`
	Oscal_component_definition_oscal_metadata_responsible_role                        Oscal_component_definition_oscal_metadata_responsible_party_101                      `json:"oscal-component-definition-oscal-metadata:responsible-role"`
	Oscal_component_definition_oscal_metadata_revision                                Oscal_component_definition_oscal_metadata_revision_103                               `json:"oscal-component-definition-oscal-metadata:revision"`
	Oscal_component_definition_oscal_metadata_role                                    Oscal_component_definition_oscal_metadata_role_105                                   `json:"oscal-component-definition-oscal-metadata:role"`
	Oscal_component_definition_oscal_metadata_role_id                                 Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-metadata:role-id"`
	Oscal_component_definition_oscal_metadata_telephone_number                        Oscal_component_definition_oscal_metadata_telephone_number_107                       `json:"oscal-component-definition-oscal-metadata:telephone-number"`
	Oscal_component_definition_oscal_metadata_version                                 Oscal_component_definition_oscal_catalog_common_parameter_value_19                   `json:"oscal-component-definition-oscal-metadata:version"`
}

type Properties_55 struct {
	Param_id Class_2   `json:"param-id"`
	Remarks  Items_3   `json:"remarks"`
	Values   Values_54 `json:"values"`
}

type Properties_13 struct {
	Prose Label_5 `json:"prose"`
}

type Items_3 struct {
	Ref string `json:"$ref"`
}

type Properties_42 struct {
	Remarks Items_3 `json:"remarks"`
	State   Class_2 `json:"state"`
}

type Properties_57 struct {
	Remarks Items_3     `json:"remarks"`
	State   How_many_16 `json:"state"`
}

type Properties_78 struct {
	Resources Resources_77 `json:"resources"`
}

type ID_61 struct {
	Type string `json:"type"`
}
