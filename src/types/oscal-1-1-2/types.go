/*
This file was auto-generated with go-oscal.

To regenerate:

	go-oscal generate \
		--input-file <path_to_oscal_json_schema_file> \
		--output-file <name_of_go_types_file> // the path to this file must already exist \
		--tags json,yaml // the tags to add to the Go structs \
		--pkg <name_of_your_go_package> // defaults to "main"

For more information on how to use go-oscal: go-oscal --help

Source: https://github.com/defenseunicorns/go-oscal
*/
package oscalTypes_1_1_2

import (
	"time"
)

type OscalModels = OscalCompleteSchema
type OscalCompleteSchema struct {
	AssessmentPlan            *AssessmentPlan            `json:"assessment-plan,omitempty" yaml:"assessment-plan,omitempty"`
	AssessmentResults         *AssessmentResults         `json:"assessment-results,omitempty" yaml:"assessment-results,omitempty"`
	Catalog                   *Catalog                   `json:"catalog,omitempty" yaml:"catalog,omitempty"`
	ComponentDefinition       *ComponentDefinition       `json:"component-definition,omitempty" yaml:"component-definition,omitempty"`
	PlanOfActionAndMilestones *PlanOfActionAndMilestones `json:"plan-of-action-and-milestones,omitempty" yaml:"plan-of-action-and-milestones,omitempty"`
	Profile                   *Profile                   `json:"profile,omitempty" yaml:"profile,omitempty"`
	SystemSecurityPlan        *SystemSecurityPlan        `json:"system-security-plan,omitempty" yaml:"system-security-plan,omitempty"`
}

type AssessmentPlan struct {
	UUID               string                            `json:"uuid" yaml:"uuid"`
	Metadata           Metadata                          `json:"metadata" yaml:"metadata"`
	ImportSsp          ImportSsp                         `json:"import-ssp" yaml:"import-ssp"`
	LocalDefinitions   *LocalDefinitions                 `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	TermsAndConditions *AssessmentPlanTermsAndConditions `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	ReviewedControls   ReviewedControls                  `json:"reviewed-controls" yaml:"reviewed-controls"`
	AssessmentSubjects *[]AssessmentSubject              `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
	AssessmentAssets   *AssessmentAssets                 `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	Tasks              *[]Task                           `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	BackMatter         *BackMatter                       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type AssessmentResults struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	Metadata         Metadata          `json:"metadata" yaml:"metadata"`
	ImportAp         ImportAp          `json:"import-ap" yaml:"import-ap"`
	LocalDefinitions *LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Results          []Result          `json:"results" yaml:"results"`
	BackMatter       *BackMatter       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type Catalog struct {
	UUID       string       `json:"uuid" yaml:"uuid"`
	Metadata   Metadata     `json:"metadata" yaml:"metadata"`
	Params     *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Controls   *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     *[]Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	BackMatter *BackMatter  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type ComponentDefinition struct {
	UUID                       string                       `json:"uuid" yaml:"uuid"`
	Metadata                   Metadata                     `json:"metadata" yaml:"metadata"`
	ImportComponentDefinitions *[]ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Components                 *[]DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	BackMatter                 *BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type PlanOfActionAndMilestones struct {
	UUID             string                                     `json:"uuid" yaml:"uuid"`
	Metadata         Metadata                                   `json:"metadata" yaml:"metadata"`
	ImportSsp        *ImportSsp                                 `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	SystemId         *SystemId                                  `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	LocalDefinitions *PlanOfActionAndMilestonesLocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     *[]Observation                             `json:"observations,omitempty" yaml:"observations,omitempty"`
	Risks            *[]Risk                                    `json:"risks,omitempty" yaml:"risks,omitempty"`
	Findings         *[]Finding                                 `json:"findings,omitempty" yaml:"findings,omitempty"`
	PoamItems        []PoamItem                                 `json:"poam-items" yaml:"poam-items"`
	BackMatter       *BackMatter                                `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type Profile struct {
	UUID       string      `json:"uuid" yaml:"uuid"`
	Metadata   Metadata    `json:"metadata" yaml:"metadata"`
	Imports    []Import    `json:"imports" yaml:"imports"`
	Merge      *Merge      `json:"merge,omitempty" yaml:"merge,omitempty"`
	Modify     *Modify     `json:"modify,omitempty" yaml:"modify,omitempty"`
	BackMatter *BackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type SystemSecurityPlan struct {
	UUID                  string                `json:"uuid" yaml:"uuid"`
	Metadata              Metadata              `json:"metadata" yaml:"metadata"`
	ImportProfile         ImportProfile         `json:"import-profile" yaml:"import-profile"`
	SystemCharacteristics SystemCharacteristics `json:"system-characteristics" yaml:"system-characteristics"`
	SystemImplementation  SystemImplementation  `json:"system-implementation" yaml:"system-implementation"`
	ControlImplementation ControlImplementation `json:"control-implementation" yaml:"control-implementation"`
	BackMatter            *BackMatter           `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type Metadata struct {
	Title        string                  `json:"title" yaml:"title"`
	Published    *time.Time              `json:"published,omitempty" yaml:"published,omitempty"`
	LastModified time.Time               `json:"last-modified" yaml:"last-modified"`
	Version      string                  `json:"version" yaml:"version"`
	OscalVersion string                  `json:"oscal-version" yaml:"oscal-version"`
	Revisions    *[]RevisionHistoryEntry `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	DocumentIds  *[]DocumentId           `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props        *[]Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Links        *[]Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Roles        *[]Role                 `json:"roles,omitempty" yaml:"roles,omitempty"`
	Locations    *[]Location             `json:"locations,omitempty" yaml:"locations,omitempty"`
	Actions      *[]Action               `json:"actions,omitempty" yaml:"actions,omitempty"`
	Remarks      string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImportSsp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LocalDefinitions struct {
	Components     *[]SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems *[]InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks        string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentPlanTermsAndConditions struct {
	Parts *[]AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type ReviewedControls struct {
	Description                string                         `json:"description,omitempty" yaml:"description,omitempty"`
	Props                      *[]Property                    `json:"props,omitempty" yaml:"props,omitempty"`
	Links                      *[]Link                        `json:"links,omitempty" yaml:"links,omitempty"`
	ControlSelections          []AssessedControls             `json:"control-selections" yaml:"control-selections"`
	ControlObjectiveSelections *[]ReferencedControlObjectives `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	Remarks                    string                         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentSubject struct {
	Type            string               `json:"type" yaml:"type"`
	Description     string               `json:"description,omitempty" yaml:"description,omitempty"`
	Props           *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Links           *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeSubjects *[]SelectSubjectById `json:"include-subjects,omitempty" yaml:"include-subjects,omitempty"`
	ExcludeSubjects *[]SelectSubjectById `json:"exclude-subjects,omitempty" yaml:"exclude-subjects,omitempty"`
	Remarks         string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentAssets struct {
	Components          *[]SystemComponent   `json:"components,omitempty" yaml:"components,omitempty"`
	AssessmentPlatforms []AssessmentPlatform `json:"assessment-platforms" yaml:"assessment-platforms"`
}

type Task struct {
	UUID             string               `json:"uuid" yaml:"uuid"`
	Type             string               `json:"type" yaml:"type"`
	Title            string               `json:"title" yaml:"title"`
	Description      string               `json:"description,omitempty" yaml:"description,omitempty"`
	Props            *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Timing           *EventTiming         `json:"timing,omitempty" yaml:"timing,omitempty"`
	Tasks            *[]Task              `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Subjects         *[]AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	ResponsibleRoles *[]ResponsibleRole   `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type BackMatter struct {
	Resources *[]Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type ImportAp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Result struct {
	Title       string    `json:"title" yaml:"title"`
	Description string    `json:"description" yaml:"description"`
	Start       time.Time `json:"start" yaml:"start"`
}

type Parameter struct {
	ID          string                 `json:"id" yaml:"id"`
	Class       string                 `json:"class,omitempty" yaml:"class,omitempty"`
	DependsOn   string                 `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Props       *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Label       string                 `json:"label,omitempty" yaml:"label,omitempty"`
	Usage       string                 `json:"usage,omitempty" yaml:"usage,omitempty"`
	Constraints *[]ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Guidelines  *[]ParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Values      *[]string              `json:"values,omitempty" yaml:"values,omitempty"`
	Select      *ParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Remarks     string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Control struct {
	ID       string       `json:"id" yaml:"id"`
	Class    string       `json:"class,omitempty" yaml:"class,omitempty"`
	Title    string       `json:"title" yaml:"title"`
	Params   *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Props    *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Links    *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Parts    *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Controls *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
}

type Group struct {
	ID     string       `json:"id,omitempty" yaml:"id,omitempty"`
	Class  string       `json:"class,omitempty" yaml:"class,omitempty"`
	Title  string       `json:"title" yaml:"title"`
	Params *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Props  *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Links  *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Parts  *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Groups *[]Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type DefinedComponent struct {
	UUID                   string                      `json:"uuid" yaml:"uuid"`
	Type                   string                      `json:"type" yaml:"type"`
	Title                  string                      `json:"title" yaml:"title"`
	Description            string                      `json:"description" yaml:"description"`
	Purpose                string                      `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Props                  *[]Property                 `json:"props,omitempty" yaml:"props,omitempty"`
	Links                  *[]Link                     `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles       *[]ResponsibleRole          `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Protocols              *[]Protocol                 `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	ControlImplementations *[]ControlImplementationSet `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemId struct {
	ID             string `json:"id" yaml:"id"`
	IdentifierType string `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
}

type PlanOfActionAndMilestonesLocalDefinitions struct {
	AssessmentAssets *AssessmentAssets  `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	Components       *[]SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems   *[]InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Observation struct {
	Title       string      `json:"title,omitempty" yaml:"title,omitempty"`
	Description string      `json:"description" yaml:"description"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Methods     []string    `json:"methods" yaml:"methods"`
}

type Risk struct {
	Title             string              `json:"title" yaml:"title"`
	Description       string              `json:"description" yaml:"description"`
	Statement         string              `json:"statement" yaml:"statement"`
	Props             *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links             *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Status            string              `json:"status" yaml:"status"`
	Origins           *[]Origin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	ThreatIds         *[]ThreatId         `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	Characterizations *[]Characterization `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	MitigatingFactors *[]MitigatingFactor `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Deadline          *time.Time          `json:"deadline,omitempty" yaml:"deadline,omitempty"`
}

type Finding struct {
	Title                       string        `json:"title" yaml:"title"`
	Description                 string        `json:"description" yaml:"description"`
	Props                       *[]Property   `json:"props,omitempty" yaml:"props,omitempty"`
	Links                       *[]Link       `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     *[]Origin     `json:"origins,omitempty" yaml:"origins,omitempty"`
	Target                      FindingTarget `json:"target" yaml:"target"`
	ImplementationStatementUuid string        `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
}

type PoamItem struct {
	UUID                string                `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Title               string                `json:"title" yaml:"title"`
	Description         string                `json:"description" yaml:"description"`
	Props               *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Links               *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             *[]PoamItemOrigin     `json:"origins,omitempty" yaml:"origins,omitempty"`
	RelatedFindings     *[]RelatedFinding     `json:"related-findings,omitempty" yaml:"related-findings,omitempty"`
	RelatedObservations *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	Remarks             string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Import struct {
	Href       string      `json:"href" yaml:"href"`
	IncludeAll *IncludeAll `json:"include-all,omitempty" yaml:"include-all,omitempty"`
}

type Merge struct {
	Combine *CombinationRule     `json:"combine,omitempty" yaml:"combine,omitempty"`
	Flat    *FlatWithoutGrouping `json:"flat,omitempty" yaml:"flat,omitempty"`
	AsIs    bool                 `json:"as-is,omitempty" yaml:"as-is,omitempty"`
}

type Modify struct {
	SetParameters *[]ParameterSetting `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Alters        *[]Alteration       `json:"alters,omitempty" yaml:"alters,omitempty"`
}

type ImportProfile struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemCharacteristics struct {
	SystemIds  []SystemId `json:"system-ids" yaml:"system-ids"`
	SystemName string     `json:"system-name" yaml:"system-name"`
}

type SystemImplementation struct {
	Props                   *[]Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   *[]Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	LeveragedAuthorizations *[]LeveragedAuthorization `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Users                   []SystemUser              `json:"users" yaml:"users"`
	Components              []SystemComponent         `json:"components" yaml:"components"`
	InventoryItems          *[]InventoryItem          `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks                 string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	SetParameters           *[]SetParameter          `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
}

type RevisionHistoryEntry struct {
	LastModified *time.Time  `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Links        *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	OscalVersion string      `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Published    *time.Time  `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks      string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string      `json:"title,omitempty" yaml:"title,omitempty"`
	Version      string      `json:"version" yaml:"version"`
}

type DocumentId struct {
	Identifier string `json:"identifier" yaml:"identifier"`
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type Property struct {
	Name    string `json:"name" yaml:"name"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Value   string `json:"value" yaml:"value"`
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Group   string `json:"group,omitempty" yaml:"group,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Link struct {
	Href             string `json:"href" yaml:"href"`
	Rel              string `json:"rel,omitempty" yaml:"rel,omitempty"`
	MediaType        string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	ResourceFragment string `json:"resource-fragment,omitempty" yaml:"resource-fragment,omitempty"`
	Text             string `json:"text,omitempty" yaml:"text,omitempty"`
}

type Role struct {
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	ID          string      `json:"id" yaml:"id"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName   string      `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title       string      `json:"title" yaml:"title"`
}

type Location struct {
	Address          *Address           `json:"address,omitempty" yaml:"address,omitempty"`
	EmailAddresses   *[]string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TelephoneNumbers *[]TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Title            string             `json:"title,omitempty" yaml:"title,omitempty"`
	Urls             *[]string          `json:"urls,omitempty" yaml:"urls,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type Action struct {
	UUID    string      `json:"uuid" yaml:"uuid"`
	Date    *time.Time  `json:"date,omitempty" yaml:"date,omitempty"`
	Type    string      `json:"type" yaml:"type"`
	System  string      `json:"system" yaml:"system"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemComponent struct {
	UUID             string             `json:"uuid" yaml:"uuid"`
	Type             string             `json:"type" yaml:"type"`
	Title            string             `json:"title" yaml:"title"`
	Description      string             `json:"description" yaml:"description"`
	Purpose          string             `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Status           Status             `json:"status" yaml:"status"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Protocols        *[]Protocol        `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type InventoryItem struct {
	UUID                  string                  `json:"uuid" yaml:"uuid"`
	Description           string                  `json:"description" yaml:"description"`
	Props                 *[]Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Links                 *[]Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	ImplementedComponents *[]ImplementedComponent `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Remarks               string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentPart struct {
	UUID  string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name  string            `json:"name" yaml:"name"`
	Ns    string            `json:"ns,omitempty" yaml:"ns,omitempty"`
	Class string            `json:"class,omitempty" yaml:"class,omitempty"`
	Title string            `json:"title,omitempty" yaml:"title,omitempty"`
	Props *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Parts *[]AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Links *[]Link           `json:"links,omitempty" yaml:"links,omitempty"`
}

type AssessedControls struct {
	Description     string               `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls *[]SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ReferencedControlObjectives struct {
	Description       string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeObjectives *[]SelectObjectiveById `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	IncludeAll        *IncludeAll            `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives *[]SelectObjectiveById `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	Links             *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Props             *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks           string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type IncludeAll = map[string]interface{}

type SelectSubjectById struct {
	SubjectUuid string      `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string      `json:"type" yaml:"type"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentPlatform struct {
	Links          *[]Link          `json:"links,omitempty" yaml:"links,omitempty"`
	Props          *[]Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string           `json:"title,omitempty" yaml:"title,omitempty"`
	UsesComponents *[]UsesComponent `json:"uses-components,omitempty" yaml:"uses-components,omitempty"`
	UUID           string           `json:"uuid" yaml:"uuid"`
}

type EventTiming struct {
	AtFrequency     *FrequencyCondition   `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	OnDate          *OnDateCondition      `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange *OnDateRangeCondition `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

type ResponsibleRole struct {
	RoleId     string      `json:"role-id" yaml:"role-id"`
	Props      *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids *[]string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Remarks    string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Resource struct {
	Base64      *Base64         `json:"base64,omitempty" yaml:"base64,omitempty"`
	Citation    *Citation       `json:"citation,omitempty" yaml:"citation,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds *[]DocumentId   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props       *[]Property     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Rlinks      *[]ResourceLink `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Title       string          `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string          `json:"uuid" yaml:"uuid"`
}

type ParameterConstraint struct {
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Tests       *[]ConstraintTest `json:"tests,omitempty" yaml:"tests,omitempty"`
}

type ParameterGuideline struct {
	Prose string `json:"prose" yaml:"prose"`
}

type ParameterSelection struct {
	HowMany string `json:"how-many,omitempty" yaml:"how-many,omitempty"`
}

type Part struct {
	ID    string      `json:"id,omitempty" yaml:"id,omitempty"`
	Name  string      `json:"name" yaml:"name"`
	Ns    string      `json:"ns,omitempty" yaml:"ns,omitempty"`
	Class string      `json:"class,omitempty" yaml:"class,omitempty"`
	Title string      `json:"title,omitempty" yaml:"title,omitempty"`
	Props *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Parts *[]Part     `json:"parts,omitempty" yaml:"parts,omitempty"`
	Links *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type Protocol struct {
	UUID       string       `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string       `json:"name" yaml:"name"`
	Title      string       `json:"title,omitempty" yaml:"title,omitempty"`
	PortRanges *[]PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
}

type ControlImplementationSet struct {
	Description             string                                        `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirementControlImplementation `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   *[]Link                                       `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]Property                                   `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           *[]SetParameter                               `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                                        `json:"source" yaml:"source"`
	UUID                    string                                        `json:"uuid" yaml:"uuid"`
}

type Origin struct {
	Actors       []OriginActor  `json:"actors" yaml:"actors"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

type ThreatId struct {
	Href   string `json:"href,omitempty" yaml:"href,omitempty"`
	ID     string `json:"id" yaml:"id"`
	System string `json:"system" yaml:"system"`
}

type Characterization struct {
	Props  *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links  *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Origin Origin      `json:"origin" yaml:"origin"`
	Facets []Facet     `json:"facets" yaml:"facets"`
}

type MitigatingFactor struct {
	Description        string              `json:"description" yaml:"description"`
	ImplementationUuid string              `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Subjects           *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	UUID               string              `json:"uuid" yaml:"uuid"`
}

type FindingTarget struct {
	Type                 string                `json:"type" yaml:"type"`
	TargetId             string                `json:"target-id" yaml:"target-id"`
	Title                string                `json:"title,omitempty" yaml:"title,omitempty"`
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Props                *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Links                *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Status               ObjectiveStatus       `json:"status" yaml:"status"`
	ImplementationStatus *ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type PoamItemOrigin struct {
	Actors []OriginActor `json:"actors" yaml:"actors"`
}

type RelatedFinding struct {
	FindingUuid string `json:"finding-uuid" yaml:"finding-uuid"`
}

type RelatedObservation struct {
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

type CombinationRule struct {
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

type FlatWithoutGrouping = map[string]interface{}

type ParameterSetting struct {
	Class       string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints *[]ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   string                 `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  *[]ParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Label       string                 `json:"label,omitempty" yaml:"label,omitempty"`
	Links       *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	ParamId     string                 `json:"param-id" yaml:"param-id"`
	Props       *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Select      *ParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       string                 `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      *[]string              `json:"values,omitempty" yaml:"values,omitempty"`
}

type Alteration struct {
	Adds      *[]Addition `json:"adds,omitempty" yaml:"adds,omitempty"`
	ControlId string      `json:"control-id" yaml:"control-id"`
	Removes   *[]Removal  `json:"removes,omitempty" yaml:"removes,omitempty"`
}

type LeveragedAuthorization struct {
	DateAuthorized string      `json:"date-authorized" yaml:"date-authorized"`
	Links          *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuid      string      `json:"party-uuid" yaml:"party-uuid"`
	Props          *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string      `json:"title" yaml:"title"`
	UUID           string      `json:"uuid" yaml:"uuid"`
}

type SystemUser struct {
	Title     string `json:"title,omitempty" yaml:"title,omitempty"`
	ShortName string `json:"short-name,omitempty" yaml:"short-name,omitempty"`
}

type SetParameter struct {
	Values []string `json:"values" yaml:"values"`
}

type ImplementedRequirement struct {
	UUID             string             `json:"uuid" yaml:"uuid"`
	ControlId        string             `json:"control-id" yaml:"control-id"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters    *[]SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Statements       *[]Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	ByComponents     *[]ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Address struct {
	AddrLines *[]string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City      string    `json:"city,omitempty" yaml:"city,omitempty"`
}

type TelephoneNumber struct {
	Number string `json:"number" yaml:"number"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}

type Status struct {
	State   string `json:"state" yaml:"state"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedComponent struct {
	ComponentUuid      string              `json:"component-uuid" yaml:"component-uuid"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type SelectControlById struct {
	StatementIds *[]string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

type SelectObjectiveById struct {
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

type UsesComponent struct {
	ComponentUuid      string              `json:"component-uuid" yaml:"component-uuid"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type FrequencyCondition struct {
	Period int    `json:"period" yaml:"period"`
	Unit   string `json:"unit" yaml:"unit"`
}

type OnDateCondition struct {
	Date time.Time `json:"date" yaml:"date"`
}

type OnDateRangeCondition struct {
	End   time.Time `json:"end" yaml:"end"`
	Start time.Time `json:"start" yaml:"start"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Citation struct {
	Links *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Text  string      `json:"text" yaml:"text"`
}

type ResourceLink struct {
	Hashes    *[]Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	Href      string  `json:"href" yaml:"href"`
	MediaType string  `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

type ConstraintTest struct {
	Expression string `json:"expression" yaml:"expression"`
	Remarks    string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type PortRange struct {
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type ImplementedRequirementControlImplementation struct {
	ControlId        string                            `json:"control-id" yaml:"control-id"`
	Description      string                            `json:"description" yaml:"description"`
	Links            *[]Link                           `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property                       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole                `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    *[]SetParameter                   `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       *[]ControlStatementImplementation `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                            `json:"uuid" yaml:"uuid"`
}

type OriginActor struct {
	Type      string      `json:"type" yaml:"type"`
	ActorUuid string      `json:"actor-uuid" yaml:"actor-uuid"`
	RoleId    string      `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	Props     *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links     *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type RelatedTask struct {
	TaskUuid          string               `json:"task-uuid" yaml:"task-uuid"`
	Props             *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Links             *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Subjects          *[]AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	IdentifiedSubject *IdentifiedSubject   `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Remarks           string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Facet struct {
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Name    string      `json:"name" yaml:"name"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	System  string      `json:"system" yaml:"system"`
	Value   string      `json:"value" yaml:"value"`
}

type SubjectReference struct {
	SubjectUuid string      `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string      `json:"type" yaml:"type"`
	Title       string      `json:"title,omitempty" yaml:"title,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ObjectiveStatus struct {
	Reason  string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type ImplementationStatus struct {
	State   string `json:"state" yaml:"state"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Addition struct {
	ById     string       `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	Links    *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params   *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Position string       `json:"position,omitempty" yaml:"position,omitempty"`
	Props    *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string       `json:"title,omitempty" yaml:"title,omitempty"`
}

type Removal struct {
	ByClass    string `json:"by-class,omitempty" yaml:"by-class,omitempty"`
	ById       string `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	ByItemName string `json:"by-item-name,omitempty" yaml:"by-item-name,omitempty"`
	ByName     string `json:"by-name,omitempty" yaml:"by-name,omitempty"`
	ByNs       string `json:"by-ns,omitempty" yaml:"by-ns,omitempty"`
}

type Statement struct {
	StatementId      string             `json:"statement-id" yaml:"statement-id"`
	UUID             string             `json:"uuid" yaml:"uuid"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ByComponents     *[]ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ByComponent struct {
	ComponentUuid        string                `json:"component-uuid" yaml:"component-uuid"`
	UUID                 string                `json:"uuid" yaml:"uuid"`
	Description          string                `json:"description" yaml:"description"`
	Props                *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Links                *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters        *[]SetParameter       `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementationStatus *ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Export               *Export               `json:"export,omitempty" yaml:"export,omitempty"`
	ResponsibleRoles     *[]ResponsibleRole    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ResponsibleParty struct {
	RoleId     string      `json:"role-id" yaml:"role-id"`
	PartyUuids []string    `json:"party-uuids" yaml:"party-uuids"`
	Props      *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks    string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type ControlStatementImplementation struct {
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string             `json:"statement-id" yaml:"statement-id"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type IdentifiedSubject struct {
	SubjectPlaceholderUuid string              `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type Export struct {
	Description      string                                 `json:"description,omitempty" yaml:"description,omitempty"`
	Links            *[]Link                                `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property                            `json:"props,omitempty" yaml:"props,omitempty"`
	Provided         *[]ProvidedControlImplementation       `json:"provided,omitempty" yaml:"provided,omitempty"`
	Remarks          string                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Responsibilities *[]ControlImplementationResponsibility `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
}

type ProvidedControlImplementation struct {
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type ControlImplementationResponsibility struct {
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     string             `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}
