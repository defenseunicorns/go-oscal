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
	Description            string                  `json:"description" yaml:"description"`
	IncorporatesComponents []IncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	Name                   string                  `json:"name" yaml:"name"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
}

type ComponentDefinition struct {
	BackMatter                 BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Capabilities               []Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Components                 []DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	ImportComponentDefinitions []ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Metadata                   Metadata                    `json:"metadata" yaml:"metadata"`
	UUID                       string                      `json:"uuid" yaml:"uuid"`
}

type ControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   []Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property               `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                   `json:"source" yaml:"source"`
	UUID                    string                   `json:"uuid" yaml:"uuid"`
}

type DefinedComponent struct {
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                  `json:"description" yaml:"description"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              []Protocol              `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose                string                  `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles       []ResponsibleRole       `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Title                  string                  `json:"title" yaml:"title"`
	Type                   string                  `json:"type" yaml:"type"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
}

type ImplementedRequirement struct {
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type Statement struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type PortRange struct {
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type Protocol struct {
	Name       string      `json:"name" yaml:"name"`
	PortRanges []PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	Title      string      `json:"title,omitempty" yaml:"title,omitempty"`
	UUID       string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type SetParameter struct {
	ParamId string   `json:"param-id" yaml:"param-id"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Values  []string `json:"values" yaml:"values"`
}

type Action struct {
	Date               string             `json:"date,omitempty" yaml:"date,omitempty"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	System             string             `json:"system" yaml:"system"`
	Type               string             `json:"type" yaml:"type"`
	UUID               string             `json:"uuid" yaml:"uuid"`
}

type Address struct {
	AddrLines  []string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string   `json:"city,omitempty" yaml:"city,omitempty"`
	Country    string   `json:"country,omitempty" yaml:"country,omitempty"`
	PostalCode string   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	State      string   `json:"state,omitempty" yaml:"state,omitempty"`
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
}

type BackMatter struct {
	Resources []Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type Link struct {
	Href             string `json:"href" yaml:"href"`
	MediaType        string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Rel              string `json:"rel,omitempty" yaml:"rel,omitempty"`
	ResourceFragment string `json:"resource-fragment,omitempty" yaml:"resource-fragment,omitempty"`
	Text             string `json:"text,omitempty" yaml:"text,omitempty"`
}

type Metadata struct {
	Actions            []Action               `json:"actions,omitempty" yaml:"actions,omitempty"`
	DocumentIds        []DocumentId           `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	LastModified       string                 `json:"last-modified" yaml:"last-modified"`
	Links              []Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Locations          []Location             `json:"locations,omitempty" yaml:"locations,omitempty"`
	OscalVersion       string                 `json:"oscal-version" yaml:"oscal-version"`
	Parties            []Party                `json:"parties,omitempty" yaml:"parties,omitempty"`
	Props              []Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Published          string                 `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks            string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Revisions          []RevisionHistoryEntry `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	Roles              []Role                 `json:"roles,omitempty" yaml:"roles,omitempty"`
	Title              string                 `json:"title" yaml:"title"`
	Version            string                 `json:"version" yaml:"version"`
}

type Property struct {
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Group   string `json:"group,omitempty" yaml:"group,omitempty"`
	Name    string `json:"name" yaml:"name"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Value   string `json:"value" yaml:"value"`
}

type ResponsibleParty struct {
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids" yaml:"party-uuids"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
}

type ResponsibleRole struct {
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
}

type DocumentId struct {
	Identifier string `json:"identifier" yaml:"identifier"`
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type TelephoneNumber struct {
	Number string `json:"number" yaml:"number"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Citation struct {
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Text  string     `json:"text" yaml:"text"`
}

type PartyExternalIdentifier struct {
	ID     string `json:"id" yaml:"id"`
	Scheme string `json:"scheme" yaml:"scheme"`
}

type Location struct {
	Address          Address           `json:"address,omitempty" yaml:"address,omitempty"`
	EmailAddresses   []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TelephoneNumbers []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	Urls             []string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type Party struct {
	Addresses             []Address                 `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	EmailAddresses        []string                  `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	ExternalIds           []PartyExternalIdentifier `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Links                 []Link                    `json:"links,omitempty" yaml:"links,omitempty"`
	LocationUuids         []string                  `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations []string                  `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Name                  string                    `json:"name,omitempty" yaml:"name,omitempty"`
	Props                 []Property                `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName             string                    `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	TelephoneNumbers      []TelephoneNumber         `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Type                  string                    `json:"type" yaml:"type"`
	UUID                  string                    `json:"uuid" yaml:"uuid"`
}

type Resource struct {
	Base64      Base64         `json:"base64,omitempty" yaml:"base64,omitempty"`
	Citation    Citation       `json:"citation,omitempty" yaml:"citation,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds []DocumentId   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props       []Property     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Rlinks      []ResourceLink `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Title       string         `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string         `json:"uuid" yaml:"uuid"`
}

type RevisionHistoryEntry struct {
	LastModified string     `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Links        []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	OscalVersion string     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Published    string     `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks      string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string     `json:"title,omitempty" yaml:"title,omitempty"`
	Version      string     `json:"version" yaml:"version"`
}

type ResourceLink struct {
	Hashes    []Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	Href      string `json:"href" yaml:"href"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

type Role struct {
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	ID          string     `json:"id" yaml:"id"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName   string     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title       string     `json:"title" yaml:"title"`
}
