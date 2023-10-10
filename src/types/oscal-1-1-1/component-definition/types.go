/*
	This file was auto-generated with go-oscal.

	To regenerate:

	go-oscal \
		--input-file <path_to_oscal_json_schema_file> \
		--output-file <name_of_go_types_file> // the path to this file must already exist \
		--tags json,yaml // the tags to add to the Go structs \
		--pkg <name_of_your_go_package> // defaults to "main"

	For more information on how to use go-oscal: go-oscal --help

	Source: https://github.com/defenseunicorns/go-oscal
*/

package compdeftypes

type OscalComponentDefinitionModel struct {
	ComponentDefinition ComponentDefinition `json:"component-definition" yaml:"component-definition"`
}

type Capability struct {
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
	Name                   string                  `json:"name" yaml:"name"`
	Description            string                  `json:"description" yaml:"description"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	IncorporatesComponents []IncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
}

type ComponentDefinition struct {
	UUID                       string                      `json:"uuid" yaml:"uuid"`
	Metadata                   Metadata                    `json:"metadata" yaml:"metadata"`
	ImportComponentDefinitions []ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Components                 []DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	Capabilities               []Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	BackMatter                 BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type ControlImplementation struct {
	Source                  string                   `json:"source" yaml:"source"`
	Description             string                   `json:"description" yaml:"description"`
	Props                   []Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	UUID                    string                   `json:"uuid" yaml:"uuid"`
}

type DefinedComponent struct {
	UUID                   string                  `json:"uuid" yaml:"uuid"`
	Title                  string                  `json:"title" yaml:"title"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Protocols              []Protocol              `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Type                   string                  `json:"type" yaml:"type"`
	Description            string                  `json:"description" yaml:"description"`
	Purpose                string                  `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles       []ResponsibleRole       `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type ImplementedRequirement struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type Statement struct {
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	Description      string            `json:"description" yaml:"description"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
}

type PortRange struct {
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type Protocol struct {
	PortRanges []PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	UUID       string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string      `json:"name" yaml:"name"`
	Title      string      `json:"title,omitempty" yaml:"title,omitempty"`
}

type SetParameter struct {
	ParamId string   `json:"param-id" yaml:"param-id"`
	Values  []string `json:"values" yaml:"values"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Action struct {
	Date               string             `json:"date,omitempty" yaml:"date,omitempty"`
	Type               string             `json:"type" yaml:"type"`
	System             string             `json:"system" yaml:"system"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID               string             `json:"uuid" yaml:"uuid"`
}

type Address struct {
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
	AddrLines  []string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string   `json:"city,omitempty" yaml:"city,omitempty"`
	State      string   `json:"state,omitempty" yaml:"state,omitempty"`
	PostalCode string   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	Country    string   `json:"country,omitempty" yaml:"country,omitempty"`
}

type BackMatter struct {
	Resources []Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type Link struct {
	ResourceFragment string `json:"resource-fragment,omitempty" yaml:"resource-fragment,omitempty"`
	Text             string `json:"text,omitempty" yaml:"text,omitempty"`
	Href             string `json:"href" yaml:"href"`
	Rel              string `json:"rel,omitempty" yaml:"rel,omitempty"`
	MediaType        string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

type Metadata struct {
	LastModified       string                 `json:"last-modified" yaml:"last-modified"`
	DocumentIds        []DocumentId           `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Links              []Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Parties            []Party                `json:"parties,omitempty" yaml:"parties,omitempty"`
	Actions            []Action               `json:"actions,omitempty" yaml:"actions,omitempty"`
	Remarks            string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Props              []Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Title              string                 `json:"title" yaml:"title"`
	Published          string                 `json:"published,omitempty" yaml:"published,omitempty"`
	OscalVersion       string                 `json:"oscal-version" yaml:"oscal-version"`
	Roles              []Role                 `json:"roles,omitempty" yaml:"roles,omitempty"`
	Locations          []Location             `json:"locations,omitempty" yaml:"locations,omitempty"`
	ResponsibleParties []ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Version            string                 `json:"version" yaml:"version"`
	Revisions          []RevisionHistoryEntry `json:"revisions,omitempty" yaml:"revisions,omitempty"`
}

type Property struct {
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Group   string `json:"group,omitempty" yaml:"group,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Name    string `json:"name" yaml:"name"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Value   string `json:"value" yaml:"value"`
}

type ResponsibleParty struct {
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
	PartyUuids []string   `json:"party-uuids" yaml:"party-uuids"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type ResponsibleRole struct {
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
}

type DocumentId struct {
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Identifier string `json:"identifier" yaml:"identifier"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type TelephoneNumber struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Number string `json:"number" yaml:"number"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Citation struct {
	Text  string     `json:"text" yaml:"text"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type PartyExternalIdentifier struct {
	Scheme string `json:"scheme" yaml:"scheme"`
	ID     string `json:"id" yaml:"id"`
}

type Location struct {
	EmailAddresses   []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	TelephoneNumbers []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Urls             []string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	Address          Address           `json:"address,omitempty" yaml:"address,omitempty"`
}

type Party struct {
	ShortName             string                    `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Addresses             []Address                 `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	MemberOfOrganizations []string                  `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	UUID                  string                    `json:"uuid" yaml:"uuid"`
	Props                 []Property                `json:"props,omitempty" yaml:"props,omitempty"`
	EmailAddresses        []string                  `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	TelephoneNumbers      []TelephoneNumber         `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	ExternalIds           []PartyExternalIdentifier `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Links                 []Link                    `json:"links,omitempty" yaml:"links,omitempty"`
	Type                  string                    `json:"type" yaml:"type"`
	Name                  string                    `json:"name,omitempty" yaml:"name,omitempty"`
	LocationUuids         []string                  `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	Remarks               string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Resource struct {
	UUID        string         `json:"uuid" yaml:"uuid"`
	Title       string         `json:"title,omitempty" yaml:"title,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Props       []Property     `json:"props,omitempty" yaml:"props,omitempty"`
	Citation    Citation       `json:"citation,omitempty" yaml:"citation,omitempty"`
	Rlinks      []ResourceLink `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	DocumentIds []DocumentId   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Base64      Base64         `json:"base64,omitempty" yaml:"base64,omitempty"`
	Remarks     string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type RevisionHistoryEntry struct {
	Props        []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links        []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks      string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string     `json:"title,omitempty" yaml:"title,omitempty"`
	Published    string     `json:"published,omitempty" yaml:"published,omitempty"`
	LastModified string     `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Version      string     `json:"version" yaml:"version"`
	OscalVersion string     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
}

type ResourceLink struct {
	Href      string `json:"href" yaml:"href"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Hashes    []Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
}

type Role struct {
	Title       string     `json:"title" yaml:"title"`
	ShortName   string     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ID          string     `json:"id" yaml:"id"`
}
