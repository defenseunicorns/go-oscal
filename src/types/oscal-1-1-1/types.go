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
package oscalTypes_1_1_1

import (
	"time"
)

type OscalModels = OscalCompleteSchema
type OscalCompleteSchema struct {
	Catalog                   *Catalog                   `json:"catalog,omitempty" yaml:"catalog,omitempty"`
	Profile                   *Profile                   `json:"profile,omitempty" yaml:"profile,omitempty"`
	ComponentDefinition       *ComponentDefinition       `json:"component-definition,omitempty" yaml:"component-definition,omitempty"`
	SystemSecurityPlan        *SystemSecurityPlan        `json:"system-security-plan,omitempty" yaml:"system-security-plan,omitempty"`
	AssessmentPlan            *AssessmentPlan            `json:"assessment-plan,omitempty" yaml:"assessment-plan,omitempty"`
	AssessmentResults         *AssessmentResults         `json:"assessment-results,omitempty" yaml:"assessment-results,omitempty"`
	PlanOfActionAndMilestones *PlanOfActionAndMilestones `json:"plan-of-action-and-milestones,omitempty" yaml:"plan-of-action-and-milestones,omitempty"`
}

type Catalog struct {
	UUID       string       `json:"uuid" yaml:"uuid"`
	Metadata   Metadata     `json:"metadata" yaml:"metadata"`
	Params     *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Controls   *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     *[]Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	BackMatter *BackMatter  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type Profile struct {
	Metadata   Metadata    `json:"metadata" yaml:"metadata"`
	Imports    []Import    `json:"imports" yaml:"imports"`
	Merge      *Merge      `json:"merge,omitempty" yaml:"merge,omitempty"`
	Modify     *Modify     `json:"modify,omitempty" yaml:"modify,omitempty"`
	BackMatter *BackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	UUID       string      `json:"uuid" yaml:"uuid"`
}

type ComponentDefinition struct {
	BackMatter                 *BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	UUID                       string                       `json:"uuid" yaml:"uuid"`
	Metadata                   Metadata                     `json:"metadata" yaml:"metadata"`
	ImportComponentDefinitions *[]ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Components                 *[]DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	Capabilities               *[]Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
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

type AssessmentPlan struct {
	ReviewedControls   ReviewedControls                  `json:"reviewed-controls" yaml:"reviewed-controls"`
	Tasks              *[]Task                           `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	UUID               string                            `json:"uuid" yaml:"uuid"`
	ImportSsp          ImportSsp                         `json:"import-ssp" yaml:"import-ssp"`
	TermsAndConditions *AssessmentPlanTermsAndConditions `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	AssessmentAssets   *AssessmentAssets                 `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	BackMatter         *BackMatter                       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Metadata           Metadata                          `json:"metadata" yaml:"metadata"`
	LocalDefinitions   *LocalDefinitions                 `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	AssessmentSubjects *[]AssessmentSubject              `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
}

type AssessmentResults struct {
	UUID             string            `json:"uuid" yaml:"uuid"`
	Metadata         Metadata          `json:"metadata" yaml:"metadata"`
	ImportAp         ImportAp          `json:"import-ap" yaml:"import-ap"`
	LocalDefinitions *LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Results          []Result          `json:"results" yaml:"results"`
	BackMatter       *BackMatter       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type PlanOfActionAndMilestones struct {
	Risks            *[]Risk                                    `json:"risks,omitempty" yaml:"risks,omitempty"`
	Findings         *[]Finding                                 `json:"findings,omitempty" yaml:"findings,omitempty"`
	UUID             string                                     `json:"uuid" yaml:"uuid"`
	Metadata         Metadata                                   `json:"metadata" yaml:"metadata"`
	ImportSsp        *ImportSsp                                 `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	LocalDefinitions *PlanOfActionAndMilestonesLocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     *[]Observation                             `json:"observations,omitempty" yaml:"observations,omitempty"`
	SystemId         *SystemId                                  `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	PoamItems        []PoamItem                                 `json:"poam-items" yaml:"poam-items"`
	BackMatter       *BackMatter                                `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
}

type Metadata struct {
	Remarks            string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Links              *[]Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Roles              *[]Role                 `json:"roles,omitempty" yaml:"roles,omitempty"`
	ResponsibleParties *[]ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Actions            *[]Action               `json:"actions,omitempty" yaml:"actions,omitempty"`
	Revisions          *[]RevisionHistoryEntry `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	DocumentIds        *[]DocumentId           `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props              *[]Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Locations          *[]Location             `json:"locations,omitempty" yaml:"locations,omitempty"`
	Parties            *[]Party                `json:"parties,omitempty" yaml:"parties,omitempty"`
	LastModified       time.Time               `json:"last-modified" yaml:"last-modified"`
	Version            string                  `json:"version" yaml:"version"`
	OscalVersion       string                  `json:"oscal-version" yaml:"oscal-version"`
	Title              string                  `json:"title" yaml:"title"`
	Published          *time.Time              `json:"published,omitempty" yaml:"published,omitempty"`
}

type Parameter struct {
	Constraints *[]ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Guidelines  *[]ParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Label       string                 `json:"label,omitempty" yaml:"label,omitempty"`
	Usage       string                 `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      *[]string              `json:"values,omitempty" yaml:"values,omitempty"`
	ID          string                 `json:"id" yaml:"id"`
	Class       string                 `json:"class,omitempty" yaml:"class,omitempty"`
	DependsOn   string                 `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Props       *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Select      *ParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Remarks     string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Control struct {
	Class    string       `json:"class,omitempty" yaml:"class,omitempty"`
	Title    string       `json:"title" yaml:"title"`
	Params   *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Props    *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Links    *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Parts    *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Controls *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	ID       string       `json:"id" yaml:"id"`
}

type Group struct {
	Groups   *[]Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Class    string       `json:"class,omitempty" yaml:"class,omitempty"`
	Title    string       `json:"title" yaml:"title"`
	Params   *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Links    *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Parts    *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	ID       string       `json:"id,omitempty" yaml:"id,omitempty"`
	Props    *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Controls *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
}

type BackMatter struct {
	Resources *[]Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type Import struct {
	ExcludeControls *[]SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	Href            string               `json:"href" yaml:"href"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
}

type Merge struct {
	Combine *CombinationRule     `json:"combine,omitempty" yaml:"combine,omitempty"`
	Flat    *FlatWithoutGrouping `json:"flat,omitempty" yaml:"flat,omitempty"`
	AsIs    bool                 `json:"as-is,omitempty" yaml:"as-is,omitempty"`
	Custom  *CustomGrouping      `json:"custom,omitempty" yaml:"custom,omitempty"`
}

type Modify struct {
	SetParameters *[]ParameterSetting `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Alters        *[]Alteration       `json:"alters,omitempty" yaml:"alters,omitempty"`
}

type ImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type DefinedComponent struct {
	Title                  string                      `json:"title" yaml:"title"`
	Protocols              *[]Protocol                 `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Remarks                string                      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ControlImplementations *[]ControlImplementationSet `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	UUID                   string                      `json:"uuid" yaml:"uuid"`
	Type                   string                      `json:"type" yaml:"type"`
	Description            string                      `json:"description" yaml:"description"`
	Purpose                string                      `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Props                  *[]Property                 `json:"props,omitempty" yaml:"props,omitempty"`
	Links                  *[]Link                     `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles       *[]ResponsibleRole          `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type Capability struct {
	Name                   string                      `json:"name" yaml:"name"`
	Description            string                      `json:"description" yaml:"description"`
	Props                  *[]Property                 `json:"props,omitempty" yaml:"props,omitempty"`
	Links                  *[]Link                     `json:"links,omitempty" yaml:"links,omitempty"`
	IncorporatesComponents *[]IncorporatesComponent    `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	ControlImplementations *[]ControlImplementationSet `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Remarks                string                      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                      `json:"uuid" yaml:"uuid"`
}

type ImportProfile struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemCharacteristics struct {
	SystemName               string                `json:"system-name" yaml:"system-name"`
	SystemNameShort          string                `json:"system-name-short,omitempty" yaml:"system-name-short,omitempty"`
	Props                    *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	SecuritySensitivityLevel string                `json:"security-sensitivity-level,omitempty" yaml:"security-sensitivity-level,omitempty"`
	SystemInformation        SystemInformation     `json:"system-information" yaml:"system-information"`
	NetworkArchitecture      *NetworkArchitecture  `json:"network-architecture,omitempty" yaml:"network-architecture,omitempty"`
	Remarks                  string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description              string                `json:"description" yaml:"description"`
	Links                    *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	DateAuthorized           string                `json:"date-authorized,omitempty" yaml:"date-authorized,omitempty"`
	SecurityImpactLevel      *SecurityImpactLevel  `json:"security-impact-level,omitempty" yaml:"security-impact-level,omitempty"`
	DataFlow                 *DataFlow             `json:"data-flow,omitempty" yaml:"data-flow,omitempty"`
	ResponsibleParties       *[]ResponsibleParty   `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	SystemIds                []SystemId            `json:"system-ids" yaml:"system-ids"`
	Status                   Status                `json:"status" yaml:"status"`
	AuthorizationBoundary    AuthorizationBoundary `json:"authorization-boundary" yaml:"authorization-boundary"`
}

type SystemImplementation struct {
	Components              []SystemComponent         `json:"components" yaml:"components"`
	InventoryItems          *[]InventoryItem          `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks                 string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Props                   *[]Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   *[]Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	LeveragedAuthorizations *[]LeveragedAuthorization `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Users                   []SystemUser              `json:"users" yaml:"users"`
}

type ControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	SetParameters           *[]SetParameter          `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
}

type ReviewedControls struct {
	Description                string                         `json:"description,omitempty" yaml:"description,omitempty"`
	Props                      *[]Property                    `json:"props,omitempty" yaml:"props,omitempty"`
	Links                      *[]Link                        `json:"links,omitempty" yaml:"links,omitempty"`
	ControlSelections          []AssessedControls             `json:"control-selections" yaml:"control-selections"`
	ControlObjectiveSelections *[]ReferencedControlObjectives `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	Remarks                    string                         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Task struct {
	Title                string                `json:"title" yaml:"title"`
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Props                *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Dependencies         *[]TaskDependency     `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	AssociatedActivities *[]AssociatedActivity `json:"associated-activities,omitempty" yaml:"associated-activities,omitempty"`
	UUID                 string                `json:"uuid" yaml:"uuid"`
	Type                 string                `json:"type" yaml:"type"`
	Links                *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Timing               *EventTiming          `json:"timing,omitempty" yaml:"timing,omitempty"`
	Tasks                *[]Task               `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Subjects             *[]AssessmentSubject  `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	ResponsibleRoles     *[]ResponsibleRole    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImportSsp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentPlanTermsAndConditions struct {
	Parts *[]AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type AssessmentAssets struct {
	Components          *[]SystemComponent   `json:"components,omitempty" yaml:"components,omitempty"`
	AssessmentPlatforms []AssessmentPlatform `json:"assessment-platforms" yaml:"assessment-platforms"`
}

type LocalDefinitions struct {
	Components           *[]SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems       *[]InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Users                *[]SystemUser      `json:"users,omitempty" yaml:"users,omitempty"`
	ObjectivesAndMethods *[]LocalObjective  `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Activities           *[]Activity        `json:"activities,omitempty" yaml:"activities,omitempty"`
	Remarks              string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
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

type ImportAp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Result struct {
	Risks            *[]Risk                  `json:"risks,omitempty" yaml:"risks,omitempty"`
	Findings         *[]Finding               `json:"findings,omitempty" yaml:"findings,omitempty"`
	Description      string                   `json:"description" yaml:"description"`
	Start            time.Time                `json:"start" yaml:"start"`
	LocalDefinitions *LocalDefinitions        `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Remarks          string                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	End              *time.Time               `json:"end,omitempty" yaml:"end,omitempty"`
	Props            *[]Property              `json:"props,omitempty" yaml:"props,omitempty"`
	AssessmentLog    *AssessmentLog           `json:"assessment-log,omitempty" yaml:"assessment-log,omitempty"`
	UUID             string                   `json:"uuid" yaml:"uuid"`
	Links            *[]Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	Attestations     *[]AttestationStatements `json:"attestations,omitempty" yaml:"attestations,omitempty"`
	Observations     *[]Observation           `json:"observations,omitempty" yaml:"observations,omitempty"`
	Title            string                   `json:"title" yaml:"title"`
	ReviewedControls ReviewedControls         `json:"reviewed-controls" yaml:"reviewed-controls"`
}

type Risk struct {
	UUID                string                `json:"uuid" yaml:"uuid"`
	Deadline            *time.Time            `json:"deadline,omitempty" yaml:"deadline,omitempty"`
	RiskLog             *RiskLog              `json:"risk-log,omitempty" yaml:"risk-log,omitempty"`
	Title               string                `json:"title" yaml:"title"`
	Description         string                `json:"description" yaml:"description"`
	Statement           string                `json:"statement" yaml:"statement"`
	Origins             *[]Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Characterizations   *[]Characterization   `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	Props               *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Links               *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	ThreatIds           *[]ThreatId           `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	Status              string                `json:"status" yaml:"status"`
	MitigatingFactors   *[]MitigatingFactor   `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Remediations        *[]Response           `json:"remediations,omitempty" yaml:"remediations,omitempty"`
	RelatedObservations *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
}

type Finding struct {
	Description                 string                `json:"description" yaml:"description"`
	Props                       *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations         *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks                *[]AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Title                       string                `json:"title" yaml:"title"`
	Links                       *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     *[]Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Target                      FindingTarget         `json:"target" yaml:"target"`
	ImplementationStatementUuid string                `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
	Remarks                     string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                        string                `json:"uuid" yaml:"uuid"`
}

type PlanOfActionAndMilestonesLocalDefinitions struct {
	Components       *[]SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems   *[]InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	AssessmentAssets *AssessmentAssets  `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Observation struct {
	Collected        time.Time           `json:"collected" yaml:"collected"`
	Expires          *time.Time          `json:"expires,omitempty" yaml:"expires,omitempty"`
	Remarks          string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string              `json:"uuid" yaml:"uuid"`
	Description      string              `json:"description" yaml:"description"`
	Props            *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Methods          []string            `json:"methods" yaml:"methods"`
	Origins          *[]Origin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Title            string              `json:"title,omitempty" yaml:"title,omitempty"`
	Links            *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Types            *[]string           `json:"types,omitempty" yaml:"types,omitempty"`
	Subjects         *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	RelevantEvidence *[]RelevantEvidence `json:"relevant-evidence,omitempty" yaml:"relevant-evidence,omitempty"`
}

type SystemId struct {
	IdentifierType string `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
	ID             string `json:"id" yaml:"id"`
}

type PoamItem struct {
	UUID                string                `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Title               string                `json:"title" yaml:"title"`
	RelatedFindings     *[]RelatedFinding     `json:"related-findings,omitempty" yaml:"related-findings,omitempty"`
	Remarks             string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description         string                `json:"description" yaml:"description"`
	Props               *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Links               *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             *[]PoamItemOrigin     `json:"origins,omitempty" yaml:"origins,omitempty"`
	RelatedObservations *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks        *[]AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
}

type Link struct {
	Href             string `json:"href" yaml:"href"`
	Rel              string `json:"rel,omitempty" yaml:"rel,omitempty"`
	MediaType        string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	ResourceFragment string `json:"resource-fragment,omitempty" yaml:"resource-fragment,omitempty"`
	Text             string `json:"text,omitempty" yaml:"text,omitempty"`
}

type Role struct {
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ID          string      `json:"id" yaml:"id"`
	Title       string      `json:"title" yaml:"title"`
	ShortName   string      `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type ResponsibleParty struct {
	Props      *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks    string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string      `json:"role-id" yaml:"role-id"`
	PartyUuids []string    `json:"party-uuids" yaml:"party-uuids"`
}

type Action struct {
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID               string              `json:"uuid" yaml:"uuid"`
	Date               *time.Time          `json:"date,omitempty" yaml:"date,omitempty"`
	Type               string              `json:"type" yaml:"type"`
	System             string              `json:"system" yaml:"system"`
}

type RevisionHistoryEntry struct {
	LastModified *time.Time  `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Version      string      `json:"version" yaml:"version"`
	OscalVersion string      `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links        *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks      string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string      `json:"title,omitempty" yaml:"title,omitempty"`
	Published    *time.Time  `json:"published,omitempty" yaml:"published,omitempty"`
}

type DocumentId struct {
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Identifier string `json:"identifier" yaml:"identifier"`
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

type Location struct {
	Address          *Address           `json:"address,omitempty" yaml:"address,omitempty"`
	TelephoneNumbers *[]TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
	Title            string             `json:"title,omitempty" yaml:"title,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	EmailAddresses   *[]string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Urls             *[]string          `json:"urls,omitempty" yaml:"urls,omitempty"`
}

type Party struct {
	LocationUuids         *[]string                  `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	Type                  string                     `json:"type" yaml:"type"`
	TelephoneNumbers      *[]TelephoneNumber         `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	ShortName             string                     `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	ExternalIds           *[]PartyExternalIdentifier `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Props                 *[]Property                `json:"props,omitempty" yaml:"props,omitempty"`
	Links                 *[]Link                    `json:"links,omitempty" yaml:"links,omitempty"`
	EmailAddresses        *[]string                  `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Addresses             *[]Address                 `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	UUID                  string                     `json:"uuid" yaml:"uuid"`
	Name                  string                     `json:"name,omitempty" yaml:"name,omitempty"`
	MemberOfOrganizations *[]string                  `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Remarks               string                     `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ParameterConstraint struct {
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Tests       *[]ConstraintTest `json:"tests,omitempty" yaml:"tests,omitempty"`
}

type ParameterGuideline struct {
	Prose string `json:"prose" yaml:"prose"`
}

type ParameterSelection struct {
	Choice  *[]string `json:"choice,omitempty" yaml:"choice,omitempty"`
	HowMany string    `json:"how-many,omitempty" yaml:"how-many,omitempty"`
}

type Part struct {
	ID    string      `json:"id,omitempty" yaml:"id,omitempty"`
	Name  string      `json:"name" yaml:"name"`
	Ns    string      `json:"ns,omitempty" yaml:"ns,omitempty"`
	Props *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Class string      `json:"class,omitempty" yaml:"class,omitempty"`
	Title string      `json:"title,omitempty" yaml:"title,omitempty"`
	Prose string      `json:"prose,omitempty" yaml:"prose,omitempty"`
	Parts *[]Part     `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type Resource struct {
	UUID        string          `json:"uuid" yaml:"uuid"`
	Props       *[]Property     `json:"props,omitempty" yaml:"props,omitempty"`
	Base64      *Base64         `json:"base64,omitempty" yaml:"base64,omitempty"`
	Title       string          `json:"title,omitempty" yaml:"title,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds *[]DocumentId   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Citation    *Citation       `json:"citation,omitempty" yaml:"citation,omitempty"`
	Rlinks      *[]ResourceLink `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Remarks     string          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SelectControlById struct {
	WithChildControls string      `json:"with-child-controls,omitempty" yaml:"with-child-controls,omitempty"`
	WithIds           *[]string   `json:"with-ids,omitempty" yaml:"with-ids,omitempty"`
	Matching          *[]Matching `json:"matching,omitempty" yaml:"matching,omitempty"`
}

type IncludeAll = map[string]interface{}

type CombinationRule struct {
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

type FlatWithoutGrouping = map[string]interface{}

type CustomGrouping struct {
	Groups         *[]CustomGroupingGroup `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls *[]InsertControls      `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
}

type ParameterSetting struct {
	DependsOn   string                 `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Links       *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Constraints *[]ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Guidelines  *[]ParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Values      *[]string              `json:"values,omitempty" yaml:"values,omitempty"`
	ParamId     string                 `json:"param-id" yaml:"param-id"`
	Class       string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Props       *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Label       string                 `json:"label,omitempty" yaml:"label,omitempty"`
	Usage       string                 `json:"usage,omitempty" yaml:"usage,omitempty"`
	Select      *ParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
}

type Alteration struct {
	ControlId string      `json:"control-id" yaml:"control-id"`
	Removes   *[]Removal  `json:"removes,omitempty" yaml:"removes,omitempty"`
	Adds      *[]Addition `json:"adds,omitempty" yaml:"adds,omitempty"`
}

type Protocol struct {
	UUID       string       `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name       string       `json:"name" yaml:"name"`
	Title      string       `json:"title,omitempty" yaml:"title,omitempty"`
	PortRanges *[]PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
}

type ControlImplementationSet struct {
	ImplementedRequirements []ImplementedRequirementControlImplementation `json:"implemented-requirements" yaml:"implemented-requirements"`
	UUID                    string                                        `json:"uuid" yaml:"uuid"`
	Source                  string                                        `json:"source" yaml:"source"`
	Description             string                                        `json:"description" yaml:"description"`
	Props                   *[]Property                                   `json:"props,omitempty" yaml:"props,omitempty"`
	Links                   *[]Link                                       `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters           *[]SetParameter                               `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type ResponsibleRole struct {
	Remarks    string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string      `json:"role-id" yaml:"role-id"`
	Props      *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links      *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids *[]string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
}

type IncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type SystemInformation struct {
	Props            *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link           `json:"links,omitempty" yaml:"links,omitempty"`
	InformationTypes []InformationType `json:"information-types" yaml:"information-types"`
}

type NetworkArchitecture struct {
	Diagrams    *[]Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description string      `json:"description" yaml:"description"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type SecurityImpactLevel struct {
	SecurityObjectiveIntegrity       string `json:"security-objective-integrity" yaml:"security-objective-integrity"`
	SecurityObjectiveAvailability    string `json:"security-objective-availability" yaml:"security-objective-availability"`
	SecurityObjectiveConfidentiality string `json:"security-objective-confidentiality" yaml:"security-objective-confidentiality"`
}

type DataFlow struct {
	Description string      `json:"description" yaml:"description"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Diagrams    *[]Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type Status struct {
	State   string `json:"state" yaml:"state"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AuthorizationBoundary struct {
	Description string      `json:"description" yaml:"description"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Diagrams    *[]Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemComponent struct {
	Props            *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Status           SystemComponentStatus `json:"status" yaml:"status"`
	Protocols        *[]Protocol           `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Remarks          string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string                `json:"uuid" yaml:"uuid"`
	Title            string                `json:"title" yaml:"title"`
	Description      string                `json:"description" yaml:"description"`
	Type             string                `json:"type" yaml:"type"`
	Purpose          string                `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	ResponsibleRoles *[]ResponsibleRole    `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type InventoryItem struct {
	Props                 *[]Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Links                 *[]Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties    *[]ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	ImplementedComponents *[]ImplementedComponent `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Remarks               string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                  string                  `json:"uuid" yaml:"uuid"`
	Description           string                  `json:"description" yaml:"description"`
}

type LeveragedAuthorization struct {
	UUID           string      `json:"uuid" yaml:"uuid"`
	Title          string      `json:"title" yaml:"title"`
	Props          *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links          *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuid      string      `json:"party-uuid" yaml:"party-uuid"`
	DateAuthorized string      `json:"date-authorized" yaml:"date-authorized"`
	Remarks        string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemUser struct {
	Links                *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	RoleIds              *[]string              `json:"role-ids,omitempty" yaml:"role-ids,omitempty"`
	Remarks              string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                 string                 `json:"uuid" yaml:"uuid"`
	Title                string                 `json:"title,omitempty" yaml:"title,omitempty"`
	ShortName            string                 `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Description          string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Props                *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	AuthorizedPrivileges *[]AuthorizedPrivilege `json:"authorized-privileges,omitempty" yaml:"authorized-privileges,omitempty"`
}

type SetParameter struct {
	Values  []string `json:"values" yaml:"values"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ParamId string   `json:"param-id" yaml:"param-id"`
}

type ImplementedRequirement struct {
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	ControlId        string             `json:"control-id" yaml:"control-id"`
	SetParameters    *[]SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Statements       *[]Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	ByComponents     *[]ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type AssessedControls struct {
	Remarks         string                               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description     string                               `json:"description,omitempty" yaml:"description,omitempty"`
	Props           *[]Property                          `json:"props,omitempty" yaml:"props,omitempty"`
	Links           *[]Link                              `json:"links,omitempty" yaml:"links,omitempty"`
	IncludeAll      *IncludeAll                          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]AssessedControlsSelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	ExcludeControls *[]AssessedControlsSelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
}

type ReferencedControlObjectives struct {
	Props             *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Links             *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	IncludeAll        *IncludeAll            `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives *[]SelectObjectiveById `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	ExcludeObjectives *[]SelectObjectiveById `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	Remarks           string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description       string                 `json:"description,omitempty" yaml:"description,omitempty"`
}

type TaskDependency struct {
	Remarks  string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TaskUuid string `json:"task-uuid" yaml:"task-uuid"`
}

type AssociatedActivity struct {
	ActivityUuid     string              `json:"activity-uuid" yaml:"activity-uuid"`
	Props            *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole  `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         []AssessmentSubject `json:"subjects" yaml:"subjects"`
	Remarks          string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type EventTiming struct {
	OnDate          *OnDateCondition      `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange *OnDateRangeCondition `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
	AtFrequency     *FrequencyCondition   `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
}

type AssessmentPart struct {
	Ns    string            `json:"ns,omitempty" yaml:"ns,omitempty"`
	Prose string            `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string            `json:"title,omitempty" yaml:"title,omitempty"`
	Props *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Parts *[]AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Links *[]Link           `json:"links,omitempty" yaml:"links,omitempty"`
	UUID  string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name  string            `json:"name" yaml:"name"`
	Class string            `json:"class,omitempty" yaml:"class,omitempty"`
}

type AssessmentPlatform struct {
	UUID           string           `json:"uuid" yaml:"uuid"`
	Title          string           `json:"title,omitempty" yaml:"title,omitempty"`
	Props          *[]Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Links          *[]Link          `json:"links,omitempty" yaml:"links,omitempty"`
	UsesComponents *[]UsesComponent `json:"uses-components,omitempty" yaml:"uses-components,omitempty"`
	Remarks        string           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LocalObjective struct {
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Parts       []Part      `json:"parts" yaml:"parts"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ControlId   string      `json:"control-id" yaml:"control-id"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
}

type Activity struct {
	RelatedControls  *ReviewedControls  `json:"related-controls,omitempty" yaml:"related-controls,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title            string             `json:"title,omitempty" yaml:"title,omitempty"`
	Description      string             `json:"description" yaml:"description"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
	Steps            *[]Step            `json:"steps,omitempty" yaml:"steps,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
}

type SelectSubjectById struct {
	SubjectUuid string      `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string      `json:"type" yaml:"type"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

type AttestationStatements struct {
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Parts              []AssessmentPart    `json:"parts" yaml:"parts"`
}

type RiskLog struct {
	Entries []RiskLogEntry `json:"entries" yaml:"entries"`
}

type Origin struct {
	Actors       []OriginActor  `json:"actors" yaml:"actors"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

type Characterization struct {
	Props  *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links  *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Origin Origin      `json:"origin" yaml:"origin"`
	Facets []Facet     `json:"facets" yaml:"facets"`
}

type ThreatId struct {
	ID     string `json:"id" yaml:"id"`
	System string `json:"system" yaml:"system"`
	Href   string `json:"href,omitempty" yaml:"href,omitempty"`
}

type MitigatingFactor struct {
	UUID               string              `json:"uuid" yaml:"uuid"`
	ImplementationUuid string              `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Description        string              `json:"description" yaml:"description"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Subjects           *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
}

type Response struct {
	Lifecycle      string           `json:"lifecycle" yaml:"lifecycle"`
	Title          string           `json:"title" yaml:"title"`
	Props          *[]Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID           string           `json:"uuid" yaml:"uuid"`
	Description    string           `json:"description" yaml:"description"`
	Links          *[]Link          `json:"links,omitempty" yaml:"links,omitempty"`
	Origins        *[]Origin        `json:"origins,omitempty" yaml:"origins,omitempty"`
	RequiredAssets *[]RequiredAsset `json:"required-assets,omitempty" yaml:"required-assets,omitempty"`
	Tasks          *[]Task          `json:"tasks,omitempty" yaml:"tasks,omitempty"`
}

type RelatedObservation struct {
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

type AssociatedRisk struct {
	RiskUuid string `json:"risk-uuid" yaml:"risk-uuid"`
}

type FindingTarget struct {
	Type                 string                `json:"type" yaml:"type"`
	Links                *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Status               ObjectiveStatus       `json:"status" yaml:"status"`
	TargetId             string                `json:"target-id" yaml:"target-id"`
	Title                string                `json:"title,omitempty" yaml:"title,omitempty"`
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Props                *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	ImplementationStatus *ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SubjectReference struct {
	SubjectUuid string      `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string      `json:"type" yaml:"type"`
	Title       string      `json:"title,omitempty" yaml:"title,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type RelevantEvidence struct {
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Href        string      `json:"href,omitempty" yaml:"href,omitempty"`
	Description string      `json:"description" yaml:"description"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type RelatedFinding struct {
	FindingUuid string `json:"finding-uuid" yaml:"finding-uuid"`
}

type PoamItemOrigin struct {
	Actors []OriginActor `json:"actors" yaml:"actors"`
}

type Address struct {
	Type       string    `json:"type,omitempty" yaml:"type,omitempty"`
	AddrLines  *[]string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string    `json:"city,omitempty" yaml:"city,omitempty"`
	State      string    `json:"state,omitempty" yaml:"state,omitempty"`
	PostalCode string    `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	Country    string    `json:"country,omitempty" yaml:"country,omitempty"`
}

type TelephoneNumber struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Number string `json:"number" yaml:"number"`
}

type PartyExternalIdentifier struct {
	Scheme string `json:"scheme" yaml:"scheme"`
	ID     string `json:"id" yaml:"id"`
}

type ConstraintTest struct {
	Remarks    string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Expression string `json:"expression" yaml:"expression"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Citation struct {
	Text  string      `json:"text" yaml:"text"`
	Props *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type ResourceLink struct {
	Href      string  `json:"href" yaml:"href"`
	MediaType string  `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Hashes    *[]Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
}

type Matching struct {
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

type CustomGroupingGroup struct {
	ID             string                 `json:"id,omitempty" yaml:"id,omitempty"`
	Title          string                 `json:"title" yaml:"title"`
	Props          *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Links          *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Groups         *[]CustomGroupingGroup `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls *[]InsertControls      `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
	Class          string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Params         *[]Parameter           `json:"params,omitempty" yaml:"params,omitempty"`
	Parts          *[]Part                `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type InsertControls struct {
	Order           string               `json:"order,omitempty" yaml:"order,omitempty"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	ExcludeControls *[]SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
}

type Removal struct {
	ByName     string `json:"by-name,omitempty" yaml:"by-name,omitempty"`
	ByClass    string `json:"by-class,omitempty" yaml:"by-class,omitempty"`
	ById       string `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	ByItemName string `json:"by-item-name,omitempty" yaml:"by-item-name,omitempty"`
	ByNs       string `json:"by-ns,omitempty" yaml:"by-ns,omitempty"`
}

type Addition struct {
	Position string       `json:"position,omitempty" yaml:"position,omitempty"`
	ById     string       `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	Title    string       `json:"title,omitempty" yaml:"title,omitempty"`
	Params   *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Props    *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	Links    *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Parts    *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type PortRange struct {
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type ImplementedRequirementControlImplementation struct {
	Statements       *[]ControlStatementImplementation `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                            `json:"uuid" yaml:"uuid"`
	ControlId        string                            `json:"control-id" yaml:"control-id"`
	Links            *[]Link                           `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole                `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Description      string                            `json:"description" yaml:"description"`
	Props            *[]Property                       `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters    *[]SetParameter                   `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Remarks          string                            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type InformationType struct {
	Title                 string                           `json:"title" yaml:"title"`
	Categorizations       *[]InformationTypeCategorization `json:"categorizations,omitempty" yaml:"categorizations,omitempty"`
	Props                 *[]Property                      `json:"props,omitempty" yaml:"props,omitempty"`
	Links                 *[]Link                          `json:"links,omitempty" yaml:"links,omitempty"`
	IntegrityImpact       *Impact                          `json:"integrity-impact,omitempty" yaml:"integrity-impact,omitempty"`
	AvailabilityImpact    *Impact                          `json:"availability-impact,omitempty" yaml:"availability-impact,omitempty"`
	UUID                  string                           `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Description           string                           `json:"description" yaml:"description"`
	ConfidentialityImpact *Impact                          `json:"confidentiality-impact,omitempty" yaml:"confidentiality-impact,omitempty"`
}

type Diagram struct {
	UUID        string      `json:"uuid" yaml:"uuid"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Caption     string      `json:"caption,omitempty" yaml:"caption,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SystemComponentStatus struct {
	State   string `json:"state" yaml:"state"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedComponent struct {
	ComponentUuid      string              `json:"component-uuid" yaml:"component-uuid"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AuthorizedPrivilege struct {
	Description        string   `json:"description,omitempty" yaml:"description,omitempty"`
	FunctionsPerformed []string `json:"functions-performed" yaml:"functions-performed"`
	Title              string   `json:"title" yaml:"title"`
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
	UUID                 string                                          `json:"uuid" yaml:"uuid"`
	Description          string                                          `json:"description" yaml:"description"`
	Props                *[]Property                                     `json:"props,omitempty" yaml:"props,omitempty"`
	Links                *[]Link                                         `json:"links,omitempty" yaml:"links,omitempty"`
	SetParameters        *[]SetParameter                                 `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Inherited            *[]InheritedControlImplementation               `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	ComponentUuid        string                                          `json:"component-uuid" yaml:"component-uuid"`
	ImplementationStatus *ImplementationStatus                           `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Export               *Export                                         `json:"export,omitempty" yaml:"export,omitempty"`
	Satisfied            *[]SatisfiedControlImplementationResponsibility `json:"satisfied,omitempty" yaml:"satisfied,omitempty"`
	ResponsibleRoles     *[]ResponsibleRole                              `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks              string                                          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessedControlsSelectControlById struct {
	ControlId    string    `json:"control-id" yaml:"control-id"`
	StatementIds *[]string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

type SelectObjectiveById struct {
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

type OnDateCondition struct {
	Date time.Time `json:"date" yaml:"date"`
}

type OnDateRangeCondition struct {
	Start time.Time `json:"start" yaml:"start"`
	End   time.Time `json:"end" yaml:"end"`
}

type FrequencyCondition struct {
	Period int    `json:"period" yaml:"period"`
	Unit   string `json:"unit" yaml:"unit"`
}

type UsesComponent struct {
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ComponentUuid      string              `json:"component-uuid" yaml:"component-uuid"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
}

type Step struct {
	Description      string             `json:"description" yaml:"description"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ReviewedControls *ReviewedControls  `json:"reviewed-controls,omitempty" yaml:"reviewed-controls,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
	Title            string             `json:"title,omitempty" yaml:"title,omitempty"`
}

type AssessmentLogEntry struct {
	UUID         string         `json:"uuid" yaml:"uuid"`
	Title        string         `json:"title,omitempty" yaml:"title,omitempty"`
	Links        *[]Link        `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy     *[]LoggedBy    `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Description  string         `json:"description,omitempty" yaml:"description,omitempty"`
	Start        time.Time      `json:"start" yaml:"start"`
	End          *time.Time     `json:"end,omitempty" yaml:"end,omitempty"`
	Props        *[]Property    `json:"props,omitempty" yaml:"props,omitempty"`
}

type RiskLogEntry struct {
	Links            *[]Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy         *[]LoggedBy              `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	RelatedResponses *[]RiskResponseReference `json:"related-responses,omitempty" yaml:"related-responses,omitempty"`
	Title            string                   `json:"title,omitempty" yaml:"title,omitempty"`
	Description      string                   `json:"description,omitempty" yaml:"description,omitempty"`
	Start            time.Time                `json:"start" yaml:"start"`
	End              *time.Time               `json:"end,omitempty" yaml:"end,omitempty"`
	UUID             string                   `json:"uuid" yaml:"uuid"`
	Props            *[]Property              `json:"props,omitempty" yaml:"props,omitempty"`
	StatusChange     string                   `json:"status-change,omitempty" yaml:"status-change,omitempty"`
	Remarks          string                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OriginActor struct {
	Type      string      `json:"type" yaml:"type"`
	ActorUuid string      `json:"actor-uuid" yaml:"actor-uuid"`
	RoleId    string      `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	Props     *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links     *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
}

type RelatedTask struct {
	Subjects           *[]AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	IdentifiedSubject  *IdentifiedSubject   `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Remarks            string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TaskUuid           string               `json:"task-uuid" yaml:"task-uuid"`
	Props              *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Links              *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleParties *[]ResponsibleParty  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type Facet struct {
	Name    string      `json:"name" yaml:"name"`
	System  string      `json:"system" yaml:"system"`
	Value   string      `json:"value" yaml:"value"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type RequiredAsset struct {
	UUID        string              `json:"uuid" yaml:"uuid"`
	Subjects    *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title       string              `json:"title,omitempty" yaml:"title,omitempty"`
	Description string              `json:"description" yaml:"description"`
	Props       *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Links       *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Remarks     string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
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

type Hash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type ControlStatementImplementation struct {
	Description      string             `json:"description" yaml:"description"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	StatementId      string             `json:"statement-id" yaml:"statement-id"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type InformationTypeCategorization struct {
	System             string    `json:"system" yaml:"system"`
	InformationTypeIds *[]string `json:"information-type-ids,omitempty" yaml:"information-type-ids,omitempty"`
}

type Impact struct {
	Links                   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Base                    string      `json:"base" yaml:"base"`
	Selected                string      `json:"selected,omitempty" yaml:"selected,omitempty"`
	AdjustmentJustification string      `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Props                   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
}

type InheritedControlImplementation struct {
	ProvidedUuid     string             `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Description      string             `json:"description" yaml:"description"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
}

type Export struct {
	Description      string                                 `json:"description,omitempty" yaml:"description,omitempty"`
	Props            *[]Property                            `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link                                `json:"links,omitempty" yaml:"links,omitempty"`
	Provided         *[]ProvidedControlImplementation       `json:"provided,omitempty" yaml:"provided,omitempty"`
	Responsibilities *[]ControlImplementationResponsibility `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
	Remarks          string                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type SatisfiedControlImplementationResponsibility struct {
	ResponsibilityUuid string             `json:"responsibility-uuid,omitempty" yaml:"responsibility-uuid,omitempty"`
	Description        string             `json:"description" yaml:"description"`
	Props              *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links              *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles   *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks            string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID               string             `json:"uuid" yaml:"uuid"`
}

type LoggedBy struct {
	RoleId    string `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	PartyUuid string `json:"party-uuid" yaml:"party-uuid"`
}

type RiskResponseReference struct {
	ResponseUuid string         `json:"response-uuid" yaml:"response-uuid"`
	Props        *[]Property    `json:"props,omitempty" yaml:"props,omitempty"`
	Links        *[]Link        `json:"links,omitempty" yaml:"links,omitempty"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type IdentifiedSubject struct {
	SubjectPlaceholderUuid string              `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

type ProvidedControlImplementation struct {
	UUID             string             `json:"uuid" yaml:"uuid"`
	Description      string             `json:"description" yaml:"description"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ControlImplementationResponsibility struct {
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID             string             `json:"uuid" yaml:"uuid"`
	ProvidedUuid     string             `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Description      string             `json:"description" yaml:"description"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
}
