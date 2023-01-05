package oscal

type OscalComponentDefinition struct {
	Comment              string                          `json:"$comment"`
	ID                   string                          `json:"$id"`
	Schema               string                          `json:"$schema"`
	AdditionalProperties bool                            `json:"additionalProperties"`
	Definitions          OscalComponentDefinition_sub108 `json:"definitions"`
	MaxProperties        int64                           `json:"maxProperties"`
	Properties           OscalComponentDefinition_sub109 `json:"properties"`
	Required             []string                        `json:"required"`
	Type                 string                          `json:"type"`
}

type OscalComponentDefinition_sub45 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub44 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub58 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub57 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub69 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub68 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub71 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub70 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub73 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub72 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub76 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub75 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub9 struct {
	AdditionalProperties bool                          `json:"additionalProperties"`
	Description          string                        `json:"description"`
	Properties           OscalComponentDefinition_sub8 `json:"properties"`
	Required             []string                      `json:"required"`
	Title                string                        `json:"title"`
	Type                 string                        `json:"type"`
}

type OscalComponentDefinition_sub93 struct {
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub92 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub66 struct {
	Addr_lines  OscalComponentDefinition_sub4 `json:"addr-lines"`
	City        OscalComponentDefinition_sub2 `json:"city"`
	Country     OscalComponentDefinition_sub2 `json:"country"`
	Postal_code OscalComponentDefinition_sub2 `json:"postal-code"`
	State       OscalComponentDefinition_sub2 `json:"state"`
	Type        OscalComponentDefinition_sub2 `json:"type"`
}

type OscalComponentDefinition_sub88 struct {
	Address           OscalComponentDefinition_sub3  `json:"address"`
	Email_addresses   OscalComponentDefinition_sub4  `json:"email-addresses"`
	Links             OscalComponentDefinition_sub4  `json:"links"`
	Props             OscalComponentDefinition_sub4  `json:"props"`
	Remarks           OscalComponentDefinition_sub3  `json:"remarks"`
	Telephone_numbers OscalComponentDefinition_sub4  `json:"telephone-numbers"`
	Title             OscalComponentDefinition_sub5  `json:"title"`
	Urls              OscalComponentDefinition_sub87 `json:"urls"`
	UUID              OscalComponentDefinition_sub2  `json:"uuid"`
}

type OscalComponentDefinition_sub95 struct {
	Addresses               OscalComponentDefinition_sub4  `json:"addresses"`
	Email_addresses         OscalComponentDefinition_sub4  `json:"email-addresses"`
	External_ids            OscalComponentDefinition_sub94 `json:"external-ids"`
	Links                   OscalComponentDefinition_sub4  `json:"links"`
	Location_uuids          OscalComponentDefinition_sub4  `json:"location-uuids"`
	Member_of_organizations OscalComponentDefinition_sub54 `json:"member-of-organizations"`
	Name                    OscalComponentDefinition_sub2  `json:"name"`
	Props                   OscalComponentDefinition_sub4  `json:"props"`
	Remarks                 OscalComponentDefinition_sub3  `json:"remarks"`
	Short_name              OscalComponentDefinition_sub2  `json:"short-name"`
	Telephone_numbers       OscalComponentDefinition_sub4  `json:"telephone-numbers"`
	Type                    OscalComponentDefinition_sub16 `json:"type"`
	UUID                    OscalComponentDefinition_sub2  `json:"uuid"`
}

type OscalComponentDefinition_sub83 struct {
	Algorithm OscalComponentDefinition_sub2  `json:"algorithm"`
	Value     OscalComponentDefinition_sub61 `json:"value"`
}

type OscalComponentDefinition_sub64 struct {
	Authorized_privileges OscalComponentDefinition_sub4 `json:"authorized-privileges"`
	Description           OscalComponentDefinition_sub5 `json:"description"`
	Links                 OscalComponentDefinition_sub4 `json:"links"`
	Props                 OscalComponentDefinition_sub4 `json:"props"`
	Remarks               OscalComponentDefinition_sub3 `json:"remarks"`
	Role_ids              OscalComponentDefinition_sub4 `json:"role-ids"`
	Short_name            OscalComponentDefinition_sub2 `json:"short-name"`
	Title                 OscalComponentDefinition_sub5 `json:"title"`
	UUID                  OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub25 struct {
	Back_matter                  OscalComponentDefinition_sub3 `json:"back-matter"`
	Capabilities                 OscalComponentDefinition_sub4 `json:"capabilities"`
	Components                   OscalComponentDefinition_sub4 `json:"components"`
	Import_component_definitions OscalComponentDefinition_sub4 `json:"import-component-definitions"`
	Metadata                     OscalComponentDefinition_sub3 `json:"metadata"`
	UUID                         OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub75 struct {
	Base64       OscalComponentDefinition_sub69 `json:"base64"`
	Citation     OscalComponentDefinition_sub71 `json:"citation"`
	Description  OscalComponentDefinition_sub5  `json:"description"`
	Document_ids OscalComponentDefinition_sub4  `json:"document-ids"`
	Props        OscalComponentDefinition_sub4  `json:"props"`
	Remarks      OscalComponentDefinition_sub3  `json:"remarks"`
	Rlinks       OscalComponentDefinition_sub74 `json:"rlinks"`
	Title        OscalComponentDefinition_sub5  `json:"title"`
	UUID         OscalComponentDefinition_sub2  `json:"uuid"`
}

type OscalComponentDefinition_sub17 struct {
	Choice   OscalComponentDefinition_sub15 `json:"choice"`
	How_many OscalComponentDefinition_sub16 `json:"how-many"`
}

type OscalComponentDefinition_sub6 struct {
	Class       OscalComponentDefinition_sub2 `json:"class"`
	Constraints OscalComponentDefinition_sub4 `json:"constraints"`
	Depends_on  OscalComponentDefinition_sub2 `json:"depends-on"`
	Guidelines  OscalComponentDefinition_sub4 `json:"guidelines"`
	ID          OscalComponentDefinition_sub2 `json:"id"`
	Label       OscalComponentDefinition_sub5 `json:"label"`
	Links       OscalComponentDefinition_sub4 `json:"links"`
	Props       OscalComponentDefinition_sub4 `json:"props"`
	Remarks     OscalComponentDefinition_sub3 `json:"remarks"`
	Select      OscalComponentDefinition_sub3 `json:"select"`
	Usage       OscalComponentDefinition_sub5 `json:"usage"`
	Values      OscalComponentDefinition_sub4 `json:"values"`
}

type OscalComponentDefinition_sub21 struct {
	Class OscalComponentDefinition_sub2  `json:"class"`
	ID    OscalComponentDefinition_sub2  `json:"id"`
	Links OscalComponentDefinition_sub4  `json:"links"`
	Name  OscalComponentDefinition_sub2  `json:"name"`
	Ns    OscalComponentDefinition_sub20 `json:"ns"`
	Parts OscalComponentDefinition_sub4  `json:"parts"`
	Props OscalComponentDefinition_sub4  `json:"props"`
	Prose OscalComponentDefinition_sub5  `json:"prose"`
	Title OscalComponentDefinition_sub5  `json:"title"`
}

type OscalComponentDefinition_sub97 struct {
	Class   OscalComponentDefinition_sub2  `json:"class"`
	Name    OscalComponentDefinition_sub2  `json:"name"`
	Ns      OscalComponentDefinition_sub20 `json:"ns"`
	Remarks OscalComponentDefinition_sub3  `json:"remarks"`
	UUID    OscalComponentDefinition_sub2  `json:"uuid"`
	Value   OscalComponentDefinition_sub2  `json:"value"`
}

type OscalComponentDefinition_sub109 struct {
	Component_definition OscalComponentDefinition_sub3 `json:"component-definition"`
}

type OscalComponentDefinition_sub36 struct {
	Component_uuid OscalComponentDefinition_sub2 `json:"component-uuid"`
	Description    OscalComponentDefinition_sub5 `json:"description"`
}

type OscalComponentDefinition_sub44 struct {
	Component_uuid      OscalComponentDefinition_sub2 `json:"component-uuid"`
	Links               OscalComponentDefinition_sub4 `json:"links"`
	Props               OscalComponentDefinition_sub4 `json:"props"`
	Remarks             OscalComponentDefinition_sub3 `json:"remarks"`
	Responsible_parties OscalComponentDefinition_sub4 `json:"responsible-parties"`
}

type OscalComponentDefinition_sub32 struct {
	Control_id        OscalComponentDefinition_sub2 `json:"control-id"`
	Description       OscalComponentDefinition_sub5 `json:"description"`
	Links             OscalComponentDefinition_sub4 `json:"links"`
	Props             OscalComponentDefinition_sub4 `json:"props"`
	Remarks           OscalComponentDefinition_sub3 `json:"remarks"`
	Responsible_roles OscalComponentDefinition_sub4 `json:"responsible-roles"`
	Set_parameters    OscalComponentDefinition_sub4 `json:"set-parameters"`
	Statements        OscalComponentDefinition_sub4 `json:"statements"`
	UUID              OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub23 struct {
	Control_implementations OscalComponentDefinition_sub4 `json:"control-implementations"`
	Description             OscalComponentDefinition_sub5 `json:"description"`
	Incorporates_components OscalComponentDefinition_sub4 `json:"incorporates-components"`
	Links                   OscalComponentDefinition_sub4 `json:"links"`
	Name                    OscalComponentDefinition_sub2 `json:"name"`
	Props                   OscalComponentDefinition_sub4 `json:"props"`
	Remarks                 OscalComponentDefinition_sub3 `json:"remarks"`
	UUID                    OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub30 struct {
	Control_implementations OscalComponentDefinition_sub4 `json:"control-implementations"`
	Description             OscalComponentDefinition_sub5 `json:"description"`
	Links                   OscalComponentDefinition_sub4 `json:"links"`
	Props                   OscalComponentDefinition_sub4 `json:"props"`
	Protocols               OscalComponentDefinition_sub4 `json:"protocols"`
	Purpose                 OscalComponentDefinition_sub5 `json:"purpose"`
	Remarks                 OscalComponentDefinition_sub3 `json:"remarks"`
	Responsible_roles       OscalComponentDefinition_sub4 `json:"responsible-roles"`
	Title                   OscalComponentDefinition_sub5 `json:"title"`
	Type                    OscalComponentDefinition_sub2 `json:"type"`
	UUID                    OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub40 struct {
	Description         OscalComponentDefinition_sub5 `json:"description"`
	Functions_performed OscalComponentDefinition_sub4 `json:"functions-performed"`
	Title               OscalComponentDefinition_sub5 `json:"title"`
}

type OscalComponentDefinition_sub104 struct {
	Description OscalComponentDefinition_sub5 `json:"description"`
	ID          OscalComponentDefinition_sub2 `json:"id"`
	Links       OscalComponentDefinition_sub4 `json:"links"`
	Props       OscalComponentDefinition_sub4 `json:"props"`
	Remarks     OscalComponentDefinition_sub3 `json:"remarks"`
	Short_name  OscalComponentDefinition_sub2 `json:"short-name"`
	Title       OscalComponentDefinition_sub5 `json:"title"`
}

type OscalComponentDefinition_sub47 struct {
	Description            OscalComponentDefinition_sub5  `json:"description"`
	Implemented_components OscalComponentDefinition_sub46 `json:"implemented-components"`
	Links                  OscalComponentDefinition_sub4  `json:"links"`
	Props                  OscalComponentDefinition_sub4  `json:"props"`
	Remarks                OscalComponentDefinition_sub3  `json:"remarks"`
	Responsible_parties    OscalComponentDefinition_sub4  `json:"responsible-parties"`
	UUID                   OscalComponentDefinition_sub2  `json:"uuid"`
}

type OscalComponentDefinition_sub28 struct {
	Description              OscalComponentDefinition_sub5  `json:"description"`
	Implemented_requirements OscalComponentDefinition_sub4  `json:"implemented-requirements"`
	Links                    OscalComponentDefinition_sub4  `json:"links"`
	Props                    OscalComponentDefinition_sub4  `json:"props"`
	Set_parameters           OscalComponentDefinition_sub4  `json:"set-parameters"`
	Source                   OscalComponentDefinition_sub27 `json:"source"`
	UUID                     OscalComponentDefinition_sub2  `json:"uuid"`
}

type OscalComponentDefinition_sub59 struct {
	Description       OscalComponentDefinition_sub5  `json:"description"`
	Links             OscalComponentDefinition_sub4  `json:"links"`
	Props             OscalComponentDefinition_sub4  `json:"props"`
	Protocols         OscalComponentDefinition_sub4  `json:"protocols"`
	Purpose           OscalComponentDefinition_sub5  `json:"purpose"`
	Remarks           OscalComponentDefinition_sub3  `json:"remarks"`
	Responsible_roles OscalComponentDefinition_sub4  `json:"responsible-roles"`
	Status            OscalComponentDefinition_sub58 `json:"status"`
	Title             OscalComponentDefinition_sub5  `json:"title"`
	Type              OscalComponentDefinition_sub2  `json:"type"`
	UUID              OscalComponentDefinition_sub2  `json:"uuid"`
}

type OscalComponentDefinition_sub38 struct {
	Description       OscalComponentDefinition_sub5 `json:"description"`
	Links             OscalComponentDefinition_sub4 `json:"links"`
	Props             OscalComponentDefinition_sub4 `json:"props"`
	Remarks           OscalComponentDefinition_sub3 `json:"remarks"`
	Responsible_roles OscalComponentDefinition_sub4 `json:"responsible-roles"`
	Statement_id      OscalComponentDefinition_sub2 `json:"statement-id"`
	UUID              OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub11 struct {
	Description OscalComponentDefinition_sub5  `json:"description"`
	Tests       OscalComponentDefinition_sub10 `json:"tests"`
}

type OscalComponentDefinition_sub16 struct {
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
	Pattern     string   `json:"pattern"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
}

type OscalComponentDefinition_sub20 struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub27 struct {
	Description string `json:"description"`
	Format      string `json:"format"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub49 struct {
	Description string `json:"description"`
	Minimum     int64  `json:"minimum"`
	MultipleOf  int64  `json:"multipleOf"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub2 struct {
	Description string `json:"description"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub5 struct {
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub90 struct {
	Document_ids        OscalComponentDefinition_sub4 `json:"document-ids"`
	Last_modified       OscalComponentDefinition_sub3 `json:"last-modified"`
	Links               OscalComponentDefinition_sub4 `json:"links"`
	Locations           OscalComponentDefinition_sub4 `json:"locations"`
	Oscal_version       OscalComponentDefinition_sub3 `json:"oscal-version"`
	Parties             OscalComponentDefinition_sub4 `json:"parties"`
	Props               OscalComponentDefinition_sub4 `json:"props"`
	Published           OscalComponentDefinition_sub3 `json:"published"`
	Remarks             OscalComponentDefinition_sub3 `json:"remarks"`
	Responsible_parties OscalComponentDefinition_sub4 `json:"responsible-parties"`
	Revisions           OscalComponentDefinition_sub4 `json:"revisions"`
	Roles               OscalComponentDefinition_sub4 `json:"roles"`
	Title               OscalComponentDefinition_sub5 `json:"title"`
	Version             OscalComponentDefinition_sub3 `json:"version"`
}

type OscalComponentDefinition_sub50 struct {
	End       OscalComponentDefinition_sub49 `json:"end"`
	Start     OscalComponentDefinition_sub49 `json:"start"`
	Transport OscalComponentDefinition_sub16 `json:"transport"`
}

type OscalComponentDefinition_sub8 struct {
	Expression OscalComponentDefinition_sub2 `json:"expression"`
	Remarks    OscalComponentDefinition_sub3 `json:"remarks"`
}

type OscalComponentDefinition_sub68 struct {
	Filename   OscalComponentDefinition_sub27 `json:"filename"`
	Media_type OscalComponentDefinition_sub2  `json:"media-type"`
	Value      OscalComponentDefinition_sub61 `json:"value"`
}

type OscalComponentDefinition_sub72 struct {
	Hashes     OscalComponentDefinition_sub4  `json:"hashes"`
	Href       OscalComponentDefinition_sub27 `json:"href"`
	Media_type OscalComponentDefinition_sub2  `json:"media-type"`
}

type OscalComponentDefinition_sub85 struct {
	Href       OscalComponentDefinition_sub27 `json:"href"`
	Media_type OscalComponentDefinition_sub2  `json:"media-type"`
	Rel        OscalComponentDefinition_sub2  `json:"rel"`
	Text       OscalComponentDefinition_sub5  `json:"text"`
}

type OscalComponentDefinition_sub34 struct {
	Href OscalComponentDefinition_sub27 `json:"href"`
}

type OscalComponentDefinition_sub62 struct {
	ID              OscalComponentDefinition_sub61 `json:"id"`
	Identifier_type OscalComponentDefinition_sub20 `json:"identifier-type"`
}

type OscalComponentDefinition_sub92 struct {
	ID     OscalComponentDefinition_sub61 `json:"id"`
	Scheme OscalComponentDefinition_sub20 `json:"scheme"`
}

type OscalComponentDefinition_sub101 struct {
	ID                   string                          `json:"$id"`
	AdditionalProperties bool                            `json:"additionalProperties"`
	Description          string                          `json:"description"`
	Properties           OscalComponentDefinition_sub100 `json:"properties"`
	Required             []string                        `json:"required"`
	Title                string                          `json:"title"`
	Type                 string                          `json:"type"`
}

type OscalComponentDefinition_sub103 struct {
	ID                   string                          `json:"$id"`
	AdditionalProperties bool                            `json:"additionalProperties"`
	Description          string                          `json:"description"`
	Properties           OscalComponentDefinition_sub102 `json:"properties"`
	Required             []string                        `json:"required"`
	Title                string                          `json:"title"`
	Type                 string                          `json:"type"`
}

type OscalComponentDefinition_sub105 struct {
	ID                   string                          `json:"$id"`
	AdditionalProperties bool                            `json:"additionalProperties"`
	Description          string                          `json:"description"`
	Properties           OscalComponentDefinition_sub104 `json:"properties"`
	Required             []string                        `json:"required"`
	Title                string                          `json:"title"`
	Type                 string                          `json:"type"`
}

type OscalComponentDefinition_sub107 struct {
	ID                   string                          `json:"$id"`
	AdditionalProperties bool                            `json:"additionalProperties"`
	Description          string                          `json:"description"`
	Properties           OscalComponentDefinition_sub106 `json:"properties"`
	Required             []string                        `json:"required"`
	Title                string                          `json:"title"`
	Type                 string                          `json:"type"`
}

type OscalComponentDefinition_sub12 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub11 `json:"properties"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub14 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub13 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub18 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub17 `json:"properties"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub22 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub21 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub24 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub23 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub26 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub25 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub29 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub28 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub31 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub30 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub33 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub32 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub35 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub34 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub37 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub36 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub39 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub38 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub41 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub40 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub43 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub42 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub48 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub47 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub51 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub50 `json:"properties"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub53 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub52 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub56 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub55 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub60 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub59 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub7 struct {
	ID                   string                        `json:"$id"`
	AdditionalProperties bool                          `json:"additionalProperties"`
	Description          string                        `json:"description"`
	Properties           OscalComponentDefinition_sub6 `json:"properties"`
	Required             []string                      `json:"required"`
	Title                string                        `json:"title"`
	Type                 string                        `json:"type"`
}

type OscalComponentDefinition_sub63 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub62 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub65 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub64 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub67 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub66 `json:"properties"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub79 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub78 `json:"properties"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub81 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub80 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub84 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub83 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub86 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub85 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub89 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub88 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub91 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub90 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub96 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub95 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub98 struct {
	ID                   string                         `json:"$id"`
	AdditionalProperties bool                           `json:"additionalProperties"`
	Description          string                         `json:"description"`
	Properties           OscalComponentDefinition_sub97 `json:"properties"`
	Required             []string                       `json:"required"`
	Title                string                         `json:"title"`
	Type                 string                         `json:"type"`
}

type OscalComponentDefinition_sub1 struct {
	ID                   string `json:"$id"`
	AdditionalProperties bool   `json:"additionalProperties"`
	Description          string `json:"description"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
}

type OscalComponentDefinition_sub82 struct {
	ID          string `json:"$id"`
	Description string `json:"description"`
	Format      string `json:"format"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub19 struct {
	ID          string `json:"$id"`
	Description string `json:"description"`
	Pattern     string `json:"pattern"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub99 struct {
	ID          string `json:"$id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type OscalComponentDefinition_sub80 struct {
	Identifier OscalComponentDefinition_sub61 `json:"identifier"`
	Scheme     OscalComponentDefinition_sub20 `json:"scheme"`
}

type OscalComponentDefinition_sub54 struct {
	Items    OscalComponentDefinition_sub2 `json:"items"`
	MinItems int64                         `json:"minItems"`
	Type     string                        `json:"type"`
}

type OscalComponentDefinition_sub87 struct {
	Items    OscalComponentDefinition_sub20 `json:"items"`
	MinItems int64                          `json:"minItems"`
	Type     string                         `json:"type"`
}

type OscalComponentDefinition_sub4 struct {
	Items    OscalComponentDefinition_sub3 `json:"items"`
	MinItems int64                         `json:"minItems"`
	Type     string                        `json:"type"`
}

type OscalComponentDefinition_sub46 struct {
	Items    OscalComponentDefinition_sub45 `json:"items"`
	MinItems int64                          `json:"minItems"`
	Type     string                         `json:"type"`
}

type OscalComponentDefinition_sub15 struct {
	Items    OscalComponentDefinition_sub5 `json:"items"`
	MinItems int64                         `json:"minItems"`
	Type     string                        `json:"type"`
}

type OscalComponentDefinition_sub74 struct {
	Items    OscalComponentDefinition_sub73 `json:"items"`
	MinItems int64                          `json:"minItems"`
	Type     string                         `json:"type"`
}

type OscalComponentDefinition_sub77 struct {
	Items    OscalComponentDefinition_sub76 `json:"items"`
	MinItems int64                          `json:"minItems"`
	Type     string                         `json:"type"`
}

type OscalComponentDefinition_sub10 struct {
	Items    OscalComponentDefinition_sub9 `json:"items"`
	MinItems int64                         `json:"minItems"`
	Type     string                        `json:"type"`
}

type OscalComponentDefinition_sub94 struct {
	Items    OscalComponentDefinition_sub93 `json:"items"`
	MinItems int64                          `json:"minItems"`
	Type     string                         `json:"type"`
}

type OscalComponentDefinition_sub102 struct {
	Last_modified OscalComponentDefinition_sub3 `json:"last-modified"`
	Links         OscalComponentDefinition_sub4 `json:"links"`
	Oscal_version OscalComponentDefinition_sub3 `json:"oscal-version"`
	Props         OscalComponentDefinition_sub4 `json:"props"`
	Published     OscalComponentDefinition_sub3 `json:"published"`
	Remarks       OscalComponentDefinition_sub3 `json:"remarks"`
	Title         OscalComponentDefinition_sub5 `json:"title"`
	Version       OscalComponentDefinition_sub3 `json:"version"`
}

type OscalComponentDefinition_sub100 struct {
	Links       OscalComponentDefinition_sub4 `json:"links"`
	Party_uuids OscalComponentDefinition_sub4 `json:"party-uuids"`
	Props       OscalComponentDefinition_sub4 `json:"props"`
	Remarks     OscalComponentDefinition_sub3 `json:"remarks"`
	Role_id     OscalComponentDefinition_sub2 `json:"role-id"`
}

type OscalComponentDefinition_sub70 struct {
	Links OscalComponentDefinition_sub4 `json:"links"`
	Props OscalComponentDefinition_sub4 `json:"props"`
	Text  OscalComponentDefinition_sub5 `json:"text"`
}

type OscalComponentDefinition_sub52 struct {
	Name        OscalComponentDefinition_sub2 `json:"name"`
	Port_ranges OscalComponentDefinition_sub4 `json:"port-ranges"`
	Title       OscalComponentDefinition_sub5 `json:"title"`
	UUID        OscalComponentDefinition_sub2 `json:"uuid"`
}

type OscalComponentDefinition_sub106 struct {
	Number OscalComponentDefinition_sub61 `json:"number"`
	Type   OscalComponentDefinition_sub2  `json:"type"`
}

type OscalComponentDefinition_sub108 struct {
	Oscal_component_definition_oscal_catalog_common_include_all                       OscalComponentDefinition_sub1   `json:"oscal-component-definition-oscal-catalog-common:include-all"`
	Oscal_component_definition_oscal_catalog_common_parameter                         OscalComponentDefinition_sub7   `json:"oscal-component-definition-oscal-catalog-common:parameter"`
	Oscal_component_definition_oscal_catalog_common_parameter_constraint              OscalComponentDefinition_sub12  `json:"oscal-component-definition-oscal-catalog-common:parameter-constraint"`
	Oscal_component_definition_oscal_catalog_common_parameter_guideline               OscalComponentDefinition_sub14  `json:"oscal-component-definition-oscal-catalog-common:parameter-guideline"`
	Oscal_component_definition_oscal_catalog_common_parameter_selection               OscalComponentDefinition_sub18  `json:"oscal-component-definition-oscal-catalog-common:parameter-selection"`
	Oscal_component_definition_oscal_catalog_common_parameter_value                   OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-catalog-common:parameter-value"`
	Oscal_component_definition_oscal_catalog_common_part                              OscalComponentDefinition_sub22  `json:"oscal-component-definition-oscal-catalog-common:part"`
	Oscal_component_definition_oscal_component_definition_capability                  OscalComponentDefinition_sub24  `json:"oscal-component-definition-oscal-component-definition:capability"`
	Oscal_component_definition_oscal_component_definition_component_definition        OscalComponentDefinition_sub26  `json:"oscal-component-definition-oscal-component-definition:component-definition"`
	Oscal_component_definition_oscal_component_definition_control_implementation      OscalComponentDefinition_sub29  `json:"oscal-component-definition-oscal-component-definition:control-implementation"`
	Oscal_component_definition_oscal_component_definition_defined_component           OscalComponentDefinition_sub31  `json:"oscal-component-definition-oscal-component-definition:defined-component"`
	Oscal_component_definition_oscal_component_definition_implemented_requirement     OscalComponentDefinition_sub33  `json:"oscal-component-definition-oscal-component-definition:implemented-requirement"`
	Oscal_component_definition_oscal_component_definition_import_component_definition OscalComponentDefinition_sub35  `json:"oscal-component-definition-oscal-component-definition:import-component-definition"`
	Oscal_component_definition_oscal_component_definition_incorporates_component      OscalComponentDefinition_sub37  `json:"oscal-component-definition-oscal-component-definition:incorporates-component"`
	Oscal_component_definition_oscal_component_definition_statement                   OscalComponentDefinition_sub39  `json:"oscal-component-definition-oscal-component-definition:statement"`
	Oscal_component_definition_oscal_implementation_common_authorized_privilege       OscalComponentDefinition_sub41  `json:"oscal-component-definition-oscal-implementation-common:authorized-privilege"`
	Oscal_component_definition_oscal_implementation_common_function_performed         OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-implementation-common:function-performed"`
	Oscal_component_definition_oscal_implementation_common_implementation_status      OscalComponentDefinition_sub43  `json:"oscal-component-definition-oscal-implementation-common:implementation-status"`
	Oscal_component_definition_oscal_implementation_common_inventory_item             OscalComponentDefinition_sub48  `json:"oscal-component-definition-oscal-implementation-common:inventory-item"`
	Oscal_component_definition_oscal_implementation_common_port_range                 OscalComponentDefinition_sub51  `json:"oscal-component-definition-oscal-implementation-common:port-range"`
	Oscal_component_definition_oscal_implementation_common_protocol                   OscalComponentDefinition_sub53  `json:"oscal-component-definition-oscal-implementation-common:protocol"`
	Oscal_component_definition_oscal_implementation_common_set_parameter              OscalComponentDefinition_sub56  `json:"oscal-component-definition-oscal-implementation-common:set-parameter"`
	Oscal_component_definition_oscal_implementation_common_system_component           OscalComponentDefinition_sub60  `json:"oscal-component-definition-oscal-implementation-common:system-component"`
	Oscal_component_definition_oscal_implementation_common_system_id                  OscalComponentDefinition_sub63  `json:"oscal-component-definition-oscal-implementation-common:system-id"`
	Oscal_component_definition_oscal_implementation_common_system_user                OscalComponentDefinition_sub65  `json:"oscal-component-definition-oscal-implementation-common:system-user"`
	Oscal_component_definition_oscal_metadata_addr_line                               OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-metadata:addr-line"`
	Oscal_component_definition_oscal_metadata_address                                 OscalComponentDefinition_sub67  `json:"oscal-component-definition-oscal-metadata:address"`
	Oscal_component_definition_oscal_metadata_back_matter                             OscalComponentDefinition_sub79  `json:"oscal-component-definition-oscal-metadata:back-matter"`
	Oscal_component_definition_oscal_metadata_document_id                             OscalComponentDefinition_sub81  `json:"oscal-component-definition-oscal-metadata:document-id"`
	Oscal_component_definition_oscal_metadata_email_address                           OscalComponentDefinition_sub82  `json:"oscal-component-definition-oscal-metadata:email-address"`
	Oscal_component_definition_oscal_metadata_hash                                    OscalComponentDefinition_sub84  `json:"oscal-component-definition-oscal-metadata:hash"`
	Oscal_component_definition_oscal_metadata_last_modified                           OscalComponentDefinition_sub82  `json:"oscal-component-definition-oscal-metadata:last-modified"`
	Oscal_component_definition_oscal_metadata_link                                    OscalComponentDefinition_sub86  `json:"oscal-component-definition-oscal-metadata:link"`
	Oscal_component_definition_oscal_metadata_location                                OscalComponentDefinition_sub89  `json:"oscal-component-definition-oscal-metadata:location"`
	Oscal_component_definition_oscal_metadata_location_uuid                           OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-metadata:location-uuid"`
	Oscal_component_definition_oscal_metadata_metadata                                OscalComponentDefinition_sub91  `json:"oscal-component-definition-oscal-metadata:metadata"`
	Oscal_component_definition_oscal_metadata_oscal_version                           OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-metadata:oscal-version"`
	Oscal_component_definition_oscal_metadata_party                                   OscalComponentDefinition_sub96  `json:"oscal-component-definition-oscal-metadata:party"`
	Oscal_component_definition_oscal_metadata_party_uuid                              OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-metadata:party-uuid"`
	Oscal_component_definition_oscal_metadata_property                                OscalComponentDefinition_sub98  `json:"oscal-component-definition-oscal-metadata:property"`
	Oscal_component_definition_oscal_metadata_published                               OscalComponentDefinition_sub82  `json:"oscal-component-definition-oscal-metadata:published"`
	Oscal_component_definition_oscal_metadata_remarks                                 OscalComponentDefinition_sub99  `json:"oscal-component-definition-oscal-metadata:remarks"`
	Oscal_component_definition_oscal_metadata_responsible_party                       OscalComponentDefinition_sub101 `json:"oscal-component-definition-oscal-metadata:responsible-party"`
	Oscal_component_definition_oscal_metadata_responsible_role                        OscalComponentDefinition_sub101 `json:"oscal-component-definition-oscal-metadata:responsible-role"`
	Oscal_component_definition_oscal_metadata_revision                                OscalComponentDefinition_sub103 `json:"oscal-component-definition-oscal-metadata:revision"`
	Oscal_component_definition_oscal_metadata_role                                    OscalComponentDefinition_sub105 `json:"oscal-component-definition-oscal-metadata:role"`
	Oscal_component_definition_oscal_metadata_role_id                                 OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-metadata:role-id"`
	Oscal_component_definition_oscal_metadata_telephone_number                        OscalComponentDefinition_sub107 `json:"oscal-component-definition-oscal-metadata:telephone-number"`
	Oscal_component_definition_oscal_metadata_version                                 OscalComponentDefinition_sub19  `json:"oscal-component-definition-oscal-metadata:version"`
}

type OscalComponentDefinition_sub55 struct {
	Param_id OscalComponentDefinition_sub2  `json:"param-id"`
	Remarks  OscalComponentDefinition_sub3  `json:"remarks"`
	Values   OscalComponentDefinition_sub54 `json:"values"`
}

type OscalComponentDefinition_sub13 struct {
	Prose OscalComponentDefinition_sub5 `json:"prose"`
}

type OscalComponentDefinition_sub3 struct {
	Ref string `json:"$ref"`
}

type OscalComponentDefinition_sub57 struct {
	Remarks OscalComponentDefinition_sub3  `json:"remarks"`
	State   OscalComponentDefinition_sub16 `json:"state"`
}

type OscalComponentDefinition_sub42 struct {
	Remarks OscalComponentDefinition_sub3 `json:"remarks"`
	State   OscalComponentDefinition_sub2 `json:"state"`
}

type OscalComponentDefinition_sub78 struct {
	Resources OscalComponentDefinition_sub77 `json:"resources"`
}

type OscalComponentDefinition_sub61 struct {
	Type string `json:"type"`
}
