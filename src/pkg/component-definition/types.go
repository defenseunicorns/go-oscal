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

package main

type OscalModel struct {
	ComponentDefinition ComponentDefinition `json:"component-definition" yaml:"component-definition"`
}

type Link struct {
	Href      string `json:"href" yaml:"href"`
	Rel       string `json:"rel,omitempty" yaml:"rel,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Text      string `json:"text,omitempty" yaml:"text,omitempty"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Party struct {
	Props                 []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links                 []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	EmailAddresses        []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	TelephoneNumbers      []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	LocationUuids         []string          `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations []string          `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Type                  string            `json:"type" yaml:"type"`
	ExternalIds           []ExternalIds     `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Remarks               string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName             string            `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Addresses             []Address         `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	UUID                  string            `json:"uuid" yaml:"uuid"`
	Name                  string            `json:"name,omitempty" yaml:"name,omitempty"`
}

type PortRange struct {
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type Protocol struct {
	Title      string      `json:"title,omitempty" yaml:"title,omitempty"`
	PortRanges []PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	UUID       string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string      `json:"name" yaml:"name"`
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

type DefinedComponent struct {
	UUID                   string                  `json:"uuid" yaml:"uuid"`
	Type                   string                  `json:"type" yaml:"type"`
	Title                  string                  `json:"title" yaml:"title"`
	Purpose                string                  `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles       []ResponsibleRole       `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Description            string                  `json:"description" yaml:"description"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              []Protocol              `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Rlinks struct {
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Hashes    []Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	Href      string `json:"href" yaml:"href"`
}

type BackMatter struct {
	Resources []Resources `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type Property struct {
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Name    string `json:"name" yaml:"name"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Value   string `json:"value" yaml:"value"`
}

type ExternalIds struct {
	Scheme string `json:"scheme" yaml:"scheme"`
	ID     string `json:"id" yaml:"id"`
}

type Location struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	Address          Address           `json:"address" yaml:"address"`
	Urls             []string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	EmailAddresses   []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	TelephoneNumbers []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
}

type ResponsibleRole struct {
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
}

type ComponentDefinition struct {
	Metadata                   Metadata                    `json:"metadata" yaml:"metadata"`
	ImportComponentDefinitions []ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Components                 []DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	Capabilities               []Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	BackMatter                 BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	UUID                       string                      `json:"uuid" yaml:"uuid"`
}

type Role struct {
	ShortName   string     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ID          string     `json:"id" yaml:"id"`
	Title       string     `json:"title" yaml:"title"`
}

type Metadata struct {
	Published          string             `json:"published,omitempty" yaml:"published,omitempty"`
	OscalVersion       string             `json:"oscal-version" yaml:"oscal-version"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Roles              []Role             `json:"roles,omitempty" yaml:"roles,omitempty"`
	Parties            []Party            `json:"parties,omitempty" yaml:"parties,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title              string             `json:"title" yaml:"title"`
	LastModified       string             `json:"last-modified" yaml:"last-modified"`
	Revisions          []Revision         `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	DocumentIds        []DocumentId       `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Version            string             `json:"version" yaml:"version"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Locations          []Location         `json:"locations,omitempty" yaml:"locations,omitempty"`
}

type ControlImplementation struct {
	UUID                    string                   `json:"uuid" yaml:"uuid"`
	Source                  string                   `json:"source" yaml:"source"`
	Description             string                   `json:"description" yaml:"description"`
	Props                   []Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
}

type Citation struct {
	Text  string     `json:"text" yaml:"text"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type TelephoneNumber struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Number string `json:"number" yaml:"number"`
}

type ResponsibleParty struct {
	RoleId     string     `json:"role-id" yaml:"role-id"`
	PartyUuids []string   `json:"party-uuids" yaml:"party-uuids"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedRequirement struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Description      string            `json:"description" yaml:"description"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
}

type Revision struct {
	OscalVersion string     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links        []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks      string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string     `json:"title,omitempty" yaml:"title,omitempty"`
	Published    string     `json:"published,omitempty" yaml:"published,omitempty"`
	LastModified string     `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Version      string     `json:"version" yaml:"version"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type Address struct {
	Country    string   `json:"country,omitempty" yaml:"country,omitempty"`
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
	AddrLines  []string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string   `json:"city,omitempty" yaml:"city,omitempty"`
	State      string   `json:"state,omitempty" yaml:"state,omitempty"`
	PostalCode string   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
}

type DocumentId struct {
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Identifier string `json:"identifier" yaml:"identifier"`
}

type SetParameter struct {
	Values  []string `json:"values" yaml:"values"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ParamId string   `json:"param-id" yaml:"param-id"`
}

type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type Capability struct {
	Name                   string                  `json:"name" yaml:"name"`
	Description            string                  `json:"description" yaml:"description"`
	Props                  []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Links                  []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	IncorporatesComponents []IncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	ControlImplementations []ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                  `json:"uuid" yaml:"uuid"`
}

type Resources struct {
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Citation    []Citation   `json:"citation,omitempty" yaml:"citation,omitempty"`
	Rlinks      []Rlinks     `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	UUID        string       `json:"uuid" yaml:"uuid"`
	Props       []Property   `json:"props,omitempty" yaml:"props,omitempty"`
	DocumentIds []DocumentId `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Base64      []Base64     `json:"base64,omitempty" yaml:"base64,omitempty"`
	Remarks     string       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title       string       `json:"title,omitempty" yaml:"title,omitempty"`
}
