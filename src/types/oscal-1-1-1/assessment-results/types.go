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

package assessmentResultsTypes

type OscalAssessmentResultsModel struct {
	AssessmentResults AssessmentResults `json:"assessment-results" yaml:"assessment-results"`
}

type AssessmentResults struct {
	BackMatter       BackMatter       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportAp         ImportAp         `json:"import-ap" yaml:"import-ap"`
	LocalDefinitions LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata         `json:"metadata" yaml:"metadata"`
	Results          []Result         `json:"results" yaml:"results"`
	UUID             string           `json:"uuid" yaml:"uuid"`
}

type ImportAp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Result struct {
	AssessmentLog    AssessmentLog           `json:"assessment-log,omitempty" yaml:"assessment-log,omitempty"`
	Attestations     []AttestationStatements `json:"attestations,omitempty" yaml:"attestations,omitempty"`
	Description      string                  `json:"description" yaml:"description"`
	End              string                  `json:"end,omitempty" yaml:"end,omitempty"`
	Findings         []Finding               `json:"findings,omitempty" yaml:"findings,omitempty"`
	Links            []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	LocalDefinitions LocalDefinitions        `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     []Observation           `json:"observations,omitempty" yaml:"observations,omitempty"`
	Props            []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ReviewedControls ReviewedControls        `json:"reviewed-controls" yaml:"reviewed-controls"`
	Risks            []Risk                  `json:"risks,omitempty" yaml:"risks,omitempty"`
	Start            string                  `json:"start" yaml:"start"`
	Title            string                  `json:"title" yaml:"title"`
	UUID             string                  `json:"uuid" yaml:"uuid"`
}

type Activity struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedControls  ReviewedControls  `json:"related-controls,omitempty" yaml:"related-controls,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Steps            []Step            `json:"steps,omitempty" yaml:"steps,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type AssessmentPart struct {
	Class string           `json:"class,omitempty" yaml:"class,omitempty"`
	Links []Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string           `json:"name" yaml:"name"`
	Ns    string           `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts []AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props []Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Prose string           `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string           `json:"title,omitempty" yaml:"title,omitempty"`
	UUID  string           `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AssessmentSubject struct {
	Description     string              `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeSubjects []SelectSubjectById `json:"exclude-subjects,omitempty" yaml:"exclude-subjects,omitempty"`
	IncludeAll      IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeSubjects []SelectSubjectById `json:"include-subjects,omitempty" yaml:"include-subjects,omitempty"`
	Links           []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Type            string              `json:"type" yaml:"type"`
}

type Characterization struct {
	Facets []Facet    `json:"facets" yaml:"facets"`
	Links  []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Origin Origin     `json:"origin" yaml:"origin"`
	Props  []Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type Finding struct {
	Description                 string               `json:"description" yaml:"description"`
	ImplementationStatementUuid string               `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
	Links                       []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     []Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props                       []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations         []RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks                []AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks                     string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Target                      FindingTarget        `json:"target" yaml:"target"`
	Title                       string               `json:"title" yaml:"title"`
	UUID                        string               `json:"uuid" yaml:"uuid"`
}

type FindingTarget struct {
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	ImplementationStatus ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Links                []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Status               Status               `json:"status" yaml:"status"`
	TargetId             string               `json:"target-id" yaml:"target-id"`
	Title                string               `json:"title,omitempty" yaml:"title,omitempty"`
	Type                 string               `json:"type" yaml:"type"`
}

type LocalObjective struct {
	ControlId   string     `json:"control-id" yaml:"control-id"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Parts       []Part     `json:"parts" yaml:"parts"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LoggedBy struct {
	PartyUuid string `json:"party-uuid" yaml:"party-uuid"`
	RoleId    string `json:"role-id,omitempty" yaml:"role-id,omitempty"`
}

type Observation struct {
	Collected        string             `json:"collected" yaml:"collected"`
	Description      string             `json:"description" yaml:"description"`
	Expires          string             `json:"expires,omitempty" yaml:"expires,omitempty"`
	Links            []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Methods          []string           `json:"methods" yaml:"methods"`
	Origins          []Origin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props            []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	RelevantEvidence []RelevantEvidence `json:"relevant-evidence,omitempty" yaml:"relevant-evidence,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects         []SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title            string             `json:"title,omitempty" yaml:"title,omitempty"`
	Types            []string           `json:"types,omitempty" yaml:"types,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type Origin struct {
	Actors       []OriginActor `json:"actors" yaml:"actors"`
	RelatedTasks []RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

type OriginActor struct {
	ActorUuid string     `json:"actor-uuid" yaml:"actor-uuid"`
	Links     []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props     []Property `json:"props,omitempty" yaml:"props,omitempty"`
	RoleId    string     `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	Type      string     `json:"type" yaml:"type"`
}

type RelatedTask struct {
	IdentifiedSubject  IdentifiedSubject   `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Links              []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties []ResponsibleParty  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Subjects           []AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	TaskUuid           string              `json:"task-uuid" yaml:"task-uuid"`
}

type Response struct {
	Description    string          `json:"description" yaml:"description"`
	Lifecycle      string          `json:"lifecycle" yaml:"lifecycle"`
	Links          []Link          `json:"links,omitempty" yaml:"links,omitempty"`
	Origins        []Origin        `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props          []Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RequiredAssets []RequiredAsset `json:"required-assets,omitempty" yaml:"required-assets,omitempty"`
	Tasks          []Task          `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Title          string          `json:"title" yaml:"title"`
	UUID           string          `json:"uuid" yaml:"uuid"`
}

type ReviewedControls struct {
	ControlObjectiveSelections []ReferencedControlObjectives `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	ControlSelections          []AssessedControls            `json:"control-selections" yaml:"control-selections"`
	Description                string                        `json:"description,omitempty" yaml:"description,omitempty"`
	Links                      []Link                        `json:"links,omitempty" yaml:"links,omitempty"`
	Props                      []Property                    `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                    string                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Risk struct {
	Characterizations   []Characterization   `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	Deadline            string               `json:"deadline,omitempty" yaml:"deadline,omitempty"`
	Description         string               `json:"description" yaml:"description"`
	Links               []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	MitigatingFactors   []MitigatingFactor   `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Origins             []Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations []RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	Remediations        []Response           `json:"remediations,omitempty" yaml:"remediations,omitempty"`
	RiskLog             RiskLog              `json:"risk-log,omitempty" yaml:"risk-log,omitempty"`
	Statement           string               `json:"statement" yaml:"statement"`
	Status              string               `json:"status" yaml:"status"`
	ThreatIds           []ThreatId           `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	Title               string               `json:"title" yaml:"title"`
	UUID                string               `json:"uuid" yaml:"uuid"`
}

type SelectControlById struct {
	ControlId    string   `json:"control-id" yaml:"control-id"`
	StatementIds []string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

type SelectObjectiveById struct {
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

type SelectSubjectById struct {
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string     `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string     `json:"type" yaml:"type"`
}

type SubjectReference struct {
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string     `json:"subject-uuid" yaml:"subject-uuid"`
	Title       string     `json:"title,omitempty" yaml:"title,omitempty"`
	Type        string     `json:"type" yaml:"type"`
}

type Task struct {
	AssociatedActivities []AssociatedActivity `json:"associated-activities,omitempty" yaml:"associated-activities,omitempty"`
	Dependencies         []TaskDependency     `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	Links                []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     []ResponsibleRole    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects             []AssessmentSubject  `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Tasks                []Task               `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Timing               Timing               `json:"timing,omitempty" yaml:"timing,omitempty"`
	Title                string               `json:"title" yaml:"title"`
	Type                 string               `json:"type" yaml:"type"`
	UUID                 string               `json:"uuid" yaml:"uuid"`
}

type IncludeAll struct {
}

type Part struct {
	Class string     `json:"class,omitempty" yaml:"class,omitempty"`
	ID    string     `json:"id,omitempty" yaml:"id,omitempty"`
	Links []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string     `json:"name" yaml:"name"`
	Ns    string     `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts []Part     `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Prose string     `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string     `json:"title,omitempty" yaml:"title,omitempty"`
}

type ImplementationStatus struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
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

type ThreatId struct {
	Href   string `json:"href,omitempty" yaml:"href,omitempty"`
	ID     string `json:"id" yaml:"id"`
	System string `json:"system" yaml:"system"`
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

type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

type AssociatedActivity struct {
	ActivityUuid     string              `json:"activity-uuid" yaml:"activity-uuid"`
	Links            []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole   `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type AtFrequency struct {
	Period int    `json:"period" yaml:"period"`
	Unit   string `json:"unit" yaml:"unit"`
}

type AttestationStatements struct {
	Parts              []AssessmentPart   `json:"parts" yaml:"parts"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
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

type ReferencedControlObjectives struct {
	Description       string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeObjectives []SelectObjectiveById `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	IncludeAll        IncludeAll            `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives []SelectObjectiveById `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	Links             []Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Props             []Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks           string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessedControls struct {
	Description     string              `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls []SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls []SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           []Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           []Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type TaskDependency struct {
	Remarks  string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TaskUuid string `json:"task-uuid" yaml:"task-uuid"`
}

type AssessmentLogEntry struct {
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	End          string        `json:"end,omitempty" yaml:"end,omitempty"`
	Links        []Link        `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy     []LoggedBy    `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props        []Property    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks []RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start        string        `json:"start" yaml:"start"`
	Title        string        `json:"title,omitempty" yaml:"title,omitempty"`
	UUID         string        `json:"uuid" yaml:"uuid"`
}

type PartyExternalIdentifier struct {
	ID     string `json:"id" yaml:"id"`
	Scheme string `json:"scheme" yaml:"scheme"`
}

type Facet struct {
	Links   []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Name    string     `json:"name" yaml:"name"`
	Props   []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	System  string     `json:"system" yaml:"system"`
	Value   string     `json:"value" yaml:"value"`
}

type IdentifiedSubject struct {
	SubjectPlaceholderUuid string              `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type LocalDefinitions struct {
	Activities           []Activity       `json:"activities,omitempty" yaml:"activities,omitempty"`
	ObjectivesAndMethods []LocalObjective `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Remarks              string           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
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

type MitigatingFactor struct {
	Description        string             `json:"description" yaml:"description"`
	ImplementationUuid string             `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Links              []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Subjects           []SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	UUID               string             `json:"uuid" yaml:"uuid"`
}

type OnDate struct {
	Date string `json:"date" yaml:"date"`
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

type RelatedObservation struct {
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

type AssociatedRisk struct {
	RiskUuid string `json:"risk-uuid" yaml:"risk-uuid"`
}

type RelevantEvidence struct {
	Description string     `json:"description" yaml:"description"`
	Href        string     `json:"href,omitempty" yaml:"href,omitempty"`
	Links       []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type RequiredAsset struct {
	Description string             `json:"description" yaml:"description"`
	Links       []Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props       []Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects    []SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title       string             `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string             `json:"uuid" yaml:"uuid"`
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

type RiskLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
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

type Status struct {
	Reason  string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type Step struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ReviewedControls ReviewedControls  `json:"reviewed-controls,omitempty" yaml:"reviewed-controls,omitempty"`
	Title            string            `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type Timing struct {
	AtFrequency     AtFrequency     `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	OnDate          OnDate          `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange WithinDateRange `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

type WithinDateRange struct {
	End   string `json:"end" yaml:"end"`
	Start string `json:"start" yaml:"start"`
}
