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

package systemSecurityPlan

type OscalModel struct {
	SystemSecurityPlan SystemSecurityPlan `json:"system-security-plan" yaml:"system-security-plan"`
}

type BackMatter struct {
	Resources []Resources `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type TelephoneNumber struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Number string `json:"number" yaml:"number"`
}

type InformationTypes struct {
	Links                 []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	ConfidentialityImpact []ConfidentialityImpact `json:"confidentiality-impact" yaml:"confidentiality-impact"`
	UUID                  string                  `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Description           string                  `json:"description" yaml:"description"`
	Props                 []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	AvailabilityImpact    []AvailabilityImpact    `json:"availability-impact" yaml:"availability-impact"`
	Title                 string                  `json:"title" yaml:"title"`
	Categorizations       []Categorizations       `json:"categorizations,omitempty" yaml:"categorizations,omitempty"`
	IntegrityImpact       []IntegrityImpact       `json:"integrity-impact" yaml:"integrity-impact"`
}

type ResponsibleParty struct {
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string     `json:"role-id" yaml:"role-id"`
	PartyUuids []string   `json:"party-uuids" yaml:"party-uuids"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type Metadata struct {
	Published          string             `json:"published,omitempty" yaml:"published,omitempty"`
	Version            string             `json:"version" yaml:"version"`
	Revisions          []Revision         `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Locations          []Location         `json:"locations,omitempty" yaml:"locations,omitempty"`
	Parties            []Party            `json:"parties,omitempty" yaml:"parties,omitempty"`
	Title              string             `json:"title" yaml:"title"`
	LastModified       string             `json:"last-modified" yaml:"last-modified"`
	OscalVersion       string             `json:"oscal-version" yaml:"oscal-version"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Roles              []Role             `json:"roles,omitempty" yaml:"roles,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	DocumentIds        []DocumentId       `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
}

type ConfidentialityImpact struct {
	Props                   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Base                    string     `json:"base" yaml:"base"`
	Selected                string     `json:"selected,omitempty" yaml:"selected,omitempty"`
	AdjustmentJustification string     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
}

type SystemCharacteristics struct {
	Remarks                  string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description              string                `json:"description" yaml:"description"`
	SecurityImpactLevel      SecurityImpactLevel   `json:"security-impact-level" yaml:"security-impact-level"`
	SystemIds                []SystemId            `json:"system-ids" yaml:"system-ids"`
	SystemInformation        SystemInformation     `json:"system-information" yaml:"system-information"`
	DateAuthorized           string                `json:"date-authorized,omitempty" yaml:"date-authorized,omitempty"`
	SecuritySensitivityLevel string                `json:"security-sensitivity-level" yaml:"security-sensitivity-level"`
	NetworkArchitecture      NetworkArchitecture   `json:"network-architecture,omitempty" yaml:"network-architecture,omitempty"`
	DataFlow                 DataFlow              `json:"data-flow,omitempty" yaml:"data-flow,omitempty"`
	SystemName               string                `json:"system-name" yaml:"system-name"`
	SystemNameShort          string                `json:"system-name-short,omitempty" yaml:"system-name-short,omitempty"`
	Status                   Status                `json:"status" yaml:"status"`
	AuthorizationBoundary    AuthorizationBoundary `json:"authorization-boundary" yaml:"authorization-boundary"`
	ResponsibleParties       []ResponsibleParty    `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Props                    []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Links                    []Link                `json:"links,omitempty" yaml:"links,omitempty"`
}

type Status1 struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type SystemComponent struct {
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Status           []Status          `json:"status" yaml:"status"`
	Protocols        []Protocol        `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	Type             string            `json:"type" yaml:"type"`
	Title            string            `json:"title" yaml:"title"`
	Description      string            `json:"description" yaml:"description"`
	Purpose          string            `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SetParameter struct {
	Values  []string `json:"values" yaml:"values"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ParamId string   `json:"param-id" yaml:"param-id"`
}

type Export struct {
	Provided         []Provided         `json:"provided,omitempty" yaml:"provided,omitempty"`
	Responsibilities []Responsibilities `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description      string             `json:"description,omitempty" yaml:"description,omitempty"`
	Props            []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link             `json:"links,omitempty" yaml:"links,omitempty"`
}

type ExternalIds struct {
	Scheme string `json:"scheme" yaml:"scheme"`
	ID     string `json:"id" yaml:"id"`
}

type NetworkArchitecture struct {
	Description string     `json:"description" yaml:"description"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Diagrams    []Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Inherited struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	ProvidedUuid     string            `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Description      string            `json:"description" yaml:"description"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type Satisfied struct {
	Links              []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles   []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks            string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID               string            `json:"uuid" yaml:"uuid"`
	ResponsibilityUuid string            `json:"responsibility-uuid,omitempty" yaml:"responsibility-uuid,omitempty"`
	Description        string            `json:"description" yaml:"description"`
	Props              []Property        `json:"props,omitempty" yaml:"props,omitempty"`
}

type ByComponent struct {
	Description          string               `json:"description" yaml:"description"`
	Props                []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters        []SetParameter       `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementationStatus ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Export               []Export             `json:"export,omitempty" yaml:"export,omitempty"`
	Inherited            []Inherited          `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	Remarks              string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ComponentUuid        string               `json:"component-uuid" yaml:"component-uuid"`
	UUID                 string               `json:"uuid" yaml:"uuid"`
	Links                []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Satisfied            []Satisfied          `json:"satisfied,omitempty" yaml:"satisfied,omitempty"`
	ResponsibleRoles     []ResponsibleRole    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type Location struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	EmailAddresses   []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Address          Address           `json:"address" yaml:"address"`
	TelephoneNumbers []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Urls             []string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Party struct {
	UUID                  string            `json:"uuid" yaml:"uuid"`
	Name                  string            `json:"name,omitempty" yaml:"name,omitempty"`
	ExternalIds           []ExternalIds     `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Props                 []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	EmailAddresses        []string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	TelephoneNumbers      []TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Type                  string            `json:"type" yaml:"type"`
	ShortName             string            `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Links                 []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Addresses             []Address         `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	LocationUuids         []string          `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations []string          `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Remarks               string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemInformation struct {
	Links            []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	InformationTypes []InformationTypes `json:"information-types" yaml:"information-types"`
	Props            []Property         `json:"props,omitempty" yaml:"props,omitempty"`
}

type ControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
}

type Revision struct {
	Title        string     `json:"title,omitempty" yaml:"title,omitempty"`
	Published    string     `json:"published,omitempty" yaml:"published,omitempty"`
	LastModified string     `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Version      string     `json:"version" yaml:"version"`
	OscalVersion string     `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links        []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks      string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Address struct {
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`
	AddrLines  []string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string   `json:"city,omitempty" yaml:"city,omitempty"`
	State      string   `json:"state,omitempty" yaml:"state,omitempty"`
	PostalCode string   `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	Country    string   `json:"country,omitempty" yaml:"country,omitempty"`
}

type SystemId struct {
	IdentifierType string `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
	ID             string `json:"id" yaml:"id"`
}

type SystemSecurityPlan struct {
	SystemImplementation  SystemImplementation  `json:"system-implementation" yaml:"system-implementation"`
	ControlImplementation ControlImplementation `json:"control-implementation" yaml:"control-implementation"`
	BackMatter            BackMatter            `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	UUID                  string                `json:"uuid" yaml:"uuid"`
	Metadata              Metadata              `json:"metadata" yaml:"metadata"`
	ImportProfile         ImportProfile         `json:"import-profile" yaml:"import-profile"`
	SystemCharacteristics SystemCharacteristics `json:"system-characteristics" yaml:"system-characteristics"`
}

type PortRange struct {
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type Responsibilities struct {
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	ProvidedUuid     string            `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Description      string            `json:"description" yaml:"description"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
}

type Diagram struct {
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Caption     string     `json:"caption,omitempty" yaml:"caption,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID        string     `json:"uuid" yaml:"uuid"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type Link struct {
	Href      string `json:"href" yaml:"href"`
	Rel       string `json:"rel,omitempty" yaml:"rel,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Text      string `json:"text,omitempty" yaml:"text,omitempty"`
}

type ImplementationStatus struct {
	State   string `json:"state" yaml:"state"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Statement struct {
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ByComponents     []ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type Status struct {
	State   string `json:"state" yaml:"state"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LeveragedAuthorizations struct {
	Title          string     `json:"title" yaml:"title"`
	Props          []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links          []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuid      string     `json:"party-uuid" yaml:"party-uuid"`
	DateAuthorized string     `json:"date-authorized" yaml:"date-authorized"`
	Remarks        string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID           string     `json:"uuid" yaml:"uuid"`
}

type Rlinks struct {
	Href      string `json:"href" yaml:"href"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Hashes    []Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
}

type SystemUser struct {
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Links                []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	RoleIds              []string              `json:"role-ids,omitempty" yaml:"role-ids,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                 string                `json:"uuid" yaml:"uuid"`
	ShortName            string                `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	AuthorizedPrivileges []AuthorizedPrivilege `json:"authorized-privileges,omitempty" yaml:"authorized-privileges,omitempty"`
	Title                string                `json:"title,omitempty" yaml:"title,omitempty"`
	Props                []Property            `json:"props,omitempty" yaml:"props,omitempty"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
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

type IntegrityImpact struct {
	Selected                string     `json:"selected,omitempty" yaml:"selected,omitempty"`
	AdjustmentJustification string     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Props                   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Base                    string     `json:"base" yaml:"base"`
}

type AuthorizationBoundary struct {
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description string     `json:"description" yaml:"description"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Diagrams    []Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
}

type ResponsibleRole struct {
	RoleId     string     `json:"role-id" yaml:"role-id"`
	Props      []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Remarks    string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedComponents struct {
	ComponentUuid      string             `json:"component-uuid" yaml:"component-uuid"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Provided struct {
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	Description      string            `json:"description" yaml:"description"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type DocumentId struct {
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Identifier string `json:"identifier" yaml:"identifier"`
}

type AvailabilityImpact struct {
	AdjustmentJustification string     `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Props                   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Base                    string     `json:"base" yaml:"base"`
	Selected                string     `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type Categorizations struct {
	System             string   `json:"system" yaml:"system"`
	InformationTypeIds []string `json:"information-type-ids,omitempty" yaml:"information-type-ids,omitempty"`
}

type SystemImplementation struct {
	LeveragedAuthorizations []LeveragedAuthorizations `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Users                   []SystemUser              `json:"users" yaml:"users"`
	Components              []SystemComponent         `json:"components" yaml:"components"`
	InventoryItems          []InventoryItem           `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks                 string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Props                   []Property                `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   []Link                    `json:"links,omitempty" yaml:"links,omitempty"`
}

type Resources struct {
	Remarks     string       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Citation    []Citation   `json:"citation,omitempty" yaml:"citation,omitempty"`
	Title       string       `json:"title,omitempty" yaml:"title,omitempty"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Props       []Property   `json:"props,omitempty" yaml:"props,omitempty"`
	DocumentIds []DocumentId `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Rlinks      []Rlinks     `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Base64      []Base64     `json:"base64,omitempty" yaml:"base64,omitempty"`
	UUID        string       `json:"uuid" yaml:"uuid"`
}

type ImportProfile struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedRequirement struct {
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ByComponents     []ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
}

type SecurityImpactLevel struct {
	SecurityObjectiveConfidentiality string `json:"security-objective-confidentiality" yaml:"security-objective-confidentiality"`
	SecurityObjectiveIntegrity       string `json:"security-objective-integrity" yaml:"security-objective-integrity"`
	SecurityObjectiveAvailability    string `json:"security-objective-availability" yaml:"security-objective-availability"`
}

type DataFlow struct {
	Description string     `json:"description" yaml:"description"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Diagrams    []Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Property struct {
	Name    string `json:"name" yaml:"name"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Value   string `json:"value" yaml:"value"`
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Citation struct {
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Text  string     `json:"text" yaml:"text"`
}

type AuthorizedPrivilege struct {
	Description        string   `json:"description,omitempty" yaml:"description,omitempty"`
	FunctionsPerformed []string `json:"functions-performed" yaml:"functions-performed"`
	Title              string   `json:"title" yaml:"title"`
}

type Protocol struct {
	UUID       string      `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string      `json:"name" yaml:"name"`
	Title      string      `json:"title,omitempty" yaml:"title,omitempty"`
	PortRanges []PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
}

type InventoryItem struct {
	Props                 []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Links                 []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties    []ResponsibleParty      `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	ImplementedComponents []ImplementedComponents `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Remarks               string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                  string                  `json:"uuid" yaml:"uuid"`
	Description           string                  `json:"description" yaml:"description"`
}
