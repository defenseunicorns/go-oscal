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
package oscalTypes_1_1_0

import (
	"time"
)

const Version = "1.1.0"

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

// An assessment plan, such as those provided by a FedRAMP assessor.
type AssessmentPlan struct {
	AssessmentAssets   *AssessmentAssets    `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	AssessmentSubjects *[]AssessmentSubject `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
	BackMatter         *BackMatter          `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp          ImportSsp            `json:"import-ssp" yaml:"import-ssp"`
	// Used to define data objects that are used in the assessment plan, that do
	// not appear in the referenced SSP.
	LocalDefinitions *LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata          `json:"metadata" yaml:"metadata"`
	ReviewedControls ReviewedControls  `json:"reviewed-controls" yaml:"reviewed-controls"`
	Tasks            *[]Task           `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	// Used to define various terms and conditions under which an assessment,
	// described by the plan, can be performed. Each child part defines a different
	// type of term or condition.
	TermsAndConditions *AssessmentPlanTermsAndConditions `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this assessment plan in this or other OSCAL
	// instances. The locally defined UUID of the assessment plan can be used to
	// reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Security assessment results, such as those provided by a FedRAMP assessor in
// the FedRAMP Security Assessment Report.
type AssessmentResults struct {
	BackMatter *BackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportAp   ImportAp    `json:"import-ap" yaml:"import-ap"`
	// Used to define data objects that are used in the assessment plan, that do
	// not appear in the referenced SSP.
	LocalDefinitions *LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata          `json:"metadata" yaml:"metadata"`
	Results          []Result          `json:"results" yaml:"results"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this assessment results instance in this or
	// other OSCAL instances. The locally defined UUID of the assessment result can
	// be used to reference the data item locally or globally (e.g., in an imported
	// OSCAL instance). This UUID should be assigned per-subject, which means it
	// should be consistently used to identify the same subject across revisions of
	// the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A structured, organized collection of control information.
type Catalog struct {
	BackMatter *BackMatter  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Controls   *[]Control   `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     *[]Group     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Metadata   Metadata     `json:"metadata" yaml:"metadata"`
	Params     *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	// Provides a globally unique means to identify a given catalog instance.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A collection of component descriptions, which may optionally be grouped by
// capability.
type ComponentDefinition struct {
	BackMatter                 *BackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Capabilities               *[]Capability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Components                 *[]DefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	ImportComponentDefinitions *[]ImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Metadata                   Metadata                     `json:"metadata" yaml:"metadata"`
	// Provides a globally unique means to identify a given component definition
	// instance.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A plan of action and milestones which identifies initial and residual risks,
// deviations, and disposition, such as those required by FedRAMP.
type PlanOfActionAndMilestones struct {
	BackMatter       *BackMatter                                `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Findings         *[]Finding                                 `json:"findings,omitempty" yaml:"findings,omitempty"`
	ImportSsp        *ImportSsp                                 `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	LocalDefinitions *PlanOfActionAndMilestonesLocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata                                   `json:"metadata" yaml:"metadata"`
	Observations     *[]Observation                             `json:"observations,omitempty" yaml:"observations,omitempty"`
	PoamItems        []PoamItem                                 `json:"poam-items" yaml:"poam-items"`
	Risks            *[]Risk                                    `json:"risks,omitempty" yaml:"risks,omitempty"`
	SystemId         *SystemId                                  `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	// A machine-oriented, globally unique identifier with instancescope that can
	// be used to reference this POA&M instance in this OSCAL instance. This UUID
	// should be assigned per-subject, which means it should be consistently used
	// to identify the same subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Each OSCAL profile is defined by a profile element.
type Profile struct {
	BackMatter *BackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Imports    []Import    `json:"imports" yaml:"imports"`
	Merge      *Merge      `json:"merge,omitempty" yaml:"merge,omitempty"`
	Metadata   Metadata    `json:"metadata" yaml:"metadata"`
	Modify     *Modify     `json:"modify,omitempty" yaml:"modify,omitempty"`
	// Provides a globally unique means to identify a given profile instance.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A system security plan, such as those described in NIST SP 800-18.
type SystemSecurityPlan struct {
	BackMatter            *BackMatter           `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ControlImplementation ControlImplementation `json:"control-implementation" yaml:"control-implementation"`
	ImportProfile         ImportProfile         `json:"import-profile" yaml:"import-profile"`
	Metadata              Metadata              `json:"metadata" yaml:"metadata"`
	SystemCharacteristics SystemCharacteristics `json:"system-characteristics" yaml:"system-characteristics"`
	SystemImplementation  SystemImplementation  `json:"system-implementation" yaml:"system-implementation"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this system security plan (SSP) elsewhere in
	// this or other OSCAL instances. The locally defined UUID of the SSP can be
	// used to reference the data item locally or globally (e.g., in an imported
	// OSCAL instance).This UUID should be assigned per-subject, which means it
	// should be consistently used to identify the same subject across revisions of
	// the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Identifies the assets used to perform this assessment, such as the assessment
// team, scanning tools, and assumptions.
type AssessmentAssets struct {
	AssessmentPlatforms []AssessmentPlatform `json:"assessment-platforms" yaml:"assessment-platforms"`
	Components          *[]SystemComponent   `json:"components,omitempty" yaml:"components,omitempty"`
}

// Identifies system elements being assessed, such as components, inventory
// items, and locations. In the assessment plan, this identifies a planned
// assessment subject. In the assessment results this is an actual assessment
// subject, and reflects any changes from the plan. exactly what will be the
// focus of this assessment. Any subjects not identified in this way are
// out-of-scope.
type AssessmentSubject struct {
	// A human-readable description of the collection of subjects being included in
	// this assessment.
	Description     string               `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeSubjects *[]SelectSubjectById `json:"exclude-subjects,omitempty" yaml:"exclude-subjects,omitempty"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeSubjects *[]SelectSubjectById `json:"include-subjects,omitempty" yaml:"include-subjects,omitempty"`
	Links           *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// Indicates the type of assessment subject, such as a component, inventory,
	// item, location, or party represented by this selection statement.
	Type string `json:"type" yaml:"type"`
}

// A collection of resources that may be referenced from within the OSCAL
// document instance.
type BackMatter struct {
	Resources *[]Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

// Used by the assessment plan and POA&M to import information about the system.
type ImportSsp struct {
	// A resolvable URL reference to the system security plan for the system being
	// assessed.
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Used to define data objects that are used in the assessment plan, that do not
// appear in the referenced SSP.
type LocalDefinitions struct {
	Activities           *[]Activity        `json:"activities,omitempty" yaml:"activities,omitempty"`
	Components           *[]SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems       *[]InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	ObjectivesAndMethods *[]LocalObjective  `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Remarks              string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                *[]SystemUser      `json:"users,omitempty" yaml:"users,omitempty"`
}

// Provides information about the containing document, and defines concepts that
// are shared across the document.
type Metadata struct {
	Actions            *[]Action               `json:"actions,omitempty" yaml:"actions,omitempty"`
	DocumentIds        *[]DocumentId           `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	LastModified       time.Time               `json:"last-modified" yaml:"last-modified"`
	Links              *[]Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Locations          *[]Location             `json:"locations,omitempty" yaml:"locations,omitempty"`
	OscalVersion       string                  `json:"oscal-version" yaml:"oscal-version"`
	Parties            *[]Party                `json:"parties,omitempty" yaml:"parties,omitempty"`
	Props              *[]Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Published          *time.Time              `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks            string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Revisions          *[]RevisionHistoryEntry `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	Roles              *[]Role                 `json:"roles,omitempty" yaml:"roles,omitempty"`
	// A name given to the document, which may be used by a tool for display and
	// navigation.
	Title   string `json:"title" yaml:"title"`
	Version string `json:"version" yaml:"version"`
}

// Identifies the controls being assessed and their control objectives.
type ReviewedControls struct {
	ControlObjectiveSelections *[]ReferencedControlObjectives `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	ControlSelections          []AssessedControls             `json:"control-selections" yaml:"control-selections"`
	// A human-readable description of control objectives.
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Represents a scheduled event or milestone, which may be associated with a
// series of assessment actions.
type Task struct {
	AssociatedActivities *[]AssociatedActivity `json:"associated-activities,omitempty" yaml:"associated-activities,omitempty"`
	Dependencies         *[]TaskDependency     `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	// A human-readable description of this task.
	Description      string               `json:"description,omitempty" yaml:"description,omitempty"`
	Links            *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole   `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         *[]AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Tasks            *[]Task              `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	// The timing under which the task is intended to occur.
	Timing *EventTiming `json:"timing,omitempty" yaml:"timing,omitempty"`
	// The title for this task.
	Title string `json:"title" yaml:"title"`
	// The type of task.
	Type string `json:"type" yaml:"type"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this task elsewhere in this or other OSCAL
	// instances. The locally defined UUID of the task can be used to reference the
	// data item locally or globally (e.g., in an imported OSCAL instance). This
	// UUID should be assigned per-subject, which means it should be consistently
	// used to identify the same subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Used to define various terms and conditions under which an assessment,
// described by the plan, can be performed. Each child part defines a different
// type of term or condition.
type AssessmentPlanTermsAndConditions struct {
	Parts *[]AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
}

// Used by assessment-results to import information about the original plan for
// assessing the system.
type ImportAp struct {
	// A resolvable URL reference to the assessment plan governing the assessment
	// activities.
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Used by the assessment results and POA&M. In the assessment results, this
// identifies all of the assessment observations and findings, initial and
// residual risks, deviations, and disposition. In the POA&M, this identifies
// initial and residual risks, deviations, and disposition.
type Result struct {
	// A log of all assessment-related actions taken.
	AssessmentLog *AssessmentLog           `json:"assessment-log,omitempty" yaml:"assessment-log,omitempty"`
	Attestations  *[]AttestationStatements `json:"attestations,omitempty" yaml:"attestations,omitempty"`
	// A human-readable description of this set of test results.
	Description string `json:"description" yaml:"description"`
	// Date/time stamp identifying the end of the evidence collection reflected in
	// these results. In a continuous motoring scenario, this may contain the same
	// value as start if appropriate.
	End      *time.Time `json:"end,omitempty" yaml:"end,omitempty"`
	Findings *[]Finding `json:"findings,omitempty" yaml:"findings,omitempty"`
	Links    *[]Link    `json:"links,omitempty" yaml:"links,omitempty"`
	// Used to define data objects that are used in the assessment plan, that do
	// not appear in the referenced SSP.
	LocalDefinitions *LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     *[]Observation    `json:"observations,omitempty" yaml:"observations,omitempty"`
	Props            *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ReviewedControls ReviewedControls  `json:"reviewed-controls" yaml:"reviewed-controls"`
	Risks            *[]Risk           `json:"risks,omitempty" yaml:"risks,omitempty"`
	// Date/time stamp identifying the start of the evidence collection reflected
	// in these results.
	Start time.Time `json:"start" yaml:"start"`
	// The title for this set of results.
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this set of results in this or other OSCAL
	// instances. The locally defined UUID of the assessment result can be used to
	// reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A structured object representing a requirement or guideline, which when
// implemented will reduce an aspect of risk related to an information system
// and its information.
type Control struct {
	// A textual label that provides a sub-type or characterization of the control.
	Class    string     `json:"class,omitempty" yaml:"class,omitempty"`
	Controls *[]Control `json:"controls,omitempty" yaml:"controls,omitempty"`
	// Identifies a control such that it can be referenced in the defining catalog
	// and other OSCAL instances (e.g., profiles).
	ID     string       `json:"id" yaml:"id"`
	Links  *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts  *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props  *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	// A name given to the control, which may be used by a tool for display and
	// navigation.
	Title string `json:"title" yaml:"title"`
}

// A group of controls, or of groups of controls.
type Group struct {
	// A textual label that provides a sub-type or characterization of the group.
	Class    string     `json:"class,omitempty" yaml:"class,omitempty"`
	Controls *[]Control `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups   *[]Group   `json:"groups,omitempty" yaml:"groups,omitempty"`
	// Identifies the group for the purpose of cross-linking within the defining
	// instance or from other instances that reference the catalog.
	ID     string       `json:"id,omitempty" yaml:"id,omitempty"`
	Links  *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts  *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props  *[]Property  `json:"props,omitempty" yaml:"props,omitempty"`
	// A name given to the group, which may be used by a tool for display and
	// navigation.
	Title string `json:"title" yaml:"title"`
}

// Parameters provide a mechanism for the dynamic assignment of value(s) in a
// control.
type Parameter struct {
	// A textual label that provides a characterization of the type, purpose, use
	// or scope of the parameter.
	Class       string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints *[]ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	// (deprecated) Another parameter invoking this one. This construct has been
	// deprecated and should not be used.
	DependsOn  string                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines *[]ParameterGuideline `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	// A unique identifier for the parameter.
	ID string `json:"id" yaml:"id"`
	// A short, placeholder name for the parameter, which can be used as a
	// substitute for a value if no value is assigned.
	Label   string              `json:"label,omitempty" yaml:"label,omitempty"`
	Links   *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props   *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Select  *ParameterSelection `json:"select,omitempty" yaml:"select,omitempty"`
	// Describes the purpose and use of a parameter.
	Usage  string    `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values *[]string `json:"values,omitempty" yaml:"values,omitempty"`
}

// A grouping of other components and/or capabilities.
type Capability struct {
	ControlImplementations *[]ControlImplementationSet `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	// A summary of the capability.
	Description            string                   `json:"description" yaml:"description"`
	IncorporatesComponents *[]IncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	Links                  *[]Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	// The capability's human-readable name.
	Name    string      `json:"name" yaml:"name"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// Provides a globally unique means to identify a given capability.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A defined component that can be part of an implemented system.
type DefinedComponent struct {
	ControlImplementations *[]ControlImplementationSet `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	// A description of the component, including information about its function.
	Description string      `json:"description" yaml:"description"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols   *[]Protocol `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	// A summary of the technological or business purpose of the component.
	Purpose          string             `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A human readable name for the component.
	Title string `json:"title" yaml:"title"`
	// A category describing the purpose of the component.
	Type string `json:"type" yaml:"type"`
	// Provides a globally unique means to identify a given component.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Loads a component definition from another resource.
type ImportComponentDefinition struct {
	// A link to a resource that defines a set of components and/or capabilities to
	// import into this collection.
	Href string `json:"href" yaml:"href"`
}

// Describes an individual finding.
type Finding struct {
	// A human-readable description of this finding.
	Description string `json:"description" yaml:"description"`
	// A machine-oriented identifier reference to the implementation statement in
	// the SSP to which this finding is related.
	ImplementationStatementUuid string                `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
	Links                       *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     *[]Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props                       *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations         *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks                *[]AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks                     string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Target                      FindingTarget         `json:"target" yaml:"target"`
	// The title for this finding.
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this finding in this or other OSCAL instances.
	// The locally defined UUID of the finding can be used to reference the data
	// item locally or globally (e.g., in an imported OSCAL instance). This UUID
	// should be assigned per-subject, which means it should be consistently used
	// to identify the same subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Allows components, and inventory-items to be defined within the POA&M for
// circumstances where no OSCAL-based SSP exists, or is not delivered with the
// POA&M.
type PlanOfActionAndMilestonesLocalDefinitions struct {
	AssessmentAssets *AssessmentAssets  `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	Components       *[]SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems   *[]InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Describes an individual observation.
type Observation struct {
	// Date/time stamp identifying when the finding information was collected.
	Collected time.Time `json:"collected" yaml:"collected"`
	// A human-readable description of this assessment observation.
	Description string `json:"description" yaml:"description"`
	// Date/time identifying when the finding information is out-of-date and no
	// longer valid. Typically used with continuous assessment scenarios.
	Expires          *time.Time          `json:"expires,omitempty" yaml:"expires,omitempty"`
	Links            *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Methods          []string            `json:"methods" yaml:"methods"`
	Origins          *[]Origin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props            *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	RelevantEvidence *[]RelevantEvidence `json:"relevant-evidence,omitempty" yaml:"relevant-evidence,omitempty"`
	Remarks          string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects         *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	// The title for this observation.
	Title string    `json:"title,omitempty" yaml:"title,omitempty"`
	Types *[]string `json:"types,omitempty" yaml:"types,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this observation elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the observation can be used to
	// reference the data item locally or globally (e.g., in an imorted OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Describes an individual POA&M item.
type PoamItem struct {
	// A human-readable description of POA&M item.
	Description         string                `json:"description" yaml:"description"`
	Links               *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             *[]PoamItemOrigin     `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedFindings     *[]RelatedFinding     `json:"related-findings,omitempty" yaml:"related-findings,omitempty"`
	RelatedObservations *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks        *[]AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks             string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// The title or name for this POA&M item .
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with instance scope that can
	// be used to reference this POA&M item entry in this OSCAL instance. This UUID
	// should be assigned per-subject, which means it should be consistently used
	// to identify the same subject across revisions of the document.
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// An identified risk.
type Risk struct {
	Characterizations *[]Characterization `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	// The date/time by which the risk must be resolved.
	Deadline *time.Time `json:"deadline,omitempty" yaml:"deadline,omitempty"`
	// A human-readable summary of the identified risk, to include a statement of
	// how the risk impacts the system.
	Description         string                `json:"description" yaml:"description"`
	Links               *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	MitigatingFactors   *[]MitigatingFactor   `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Origins             *[]Origin             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations *[]RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	Remediations        *[]Response           `json:"remediations,omitempty" yaml:"remediations,omitempty"`
	// A log of all risk-related tasks taken.
	RiskLog *RiskLog `json:"risk-log,omitempty" yaml:"risk-log,omitempty"`
	// An summary of impact for how the risk affects the system.
	Statement string      `json:"statement" yaml:"statement"`
	Status    string      `json:"status" yaml:"status"`
	ThreatIds *[]ThreatId `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	// The title for this risk.
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this risk elsewhere in this or other OSCAL
	// instances. The locally defined UUID of the risk can be used to reference the
	// data item locally or globally (e.g., in an imported OSCAL instance). This
	// UUID should be assigned per-subject, which means it should be consistently
	// used to identify the same subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A human-oriented, globally unique identifier with cross-instance scope that
// can be used to reference this system identification property elsewhere in
// this or other OSCAL instances. When referencing an externally defined system
// identification, the system identification must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference). This string should
// be assigned per-subject, which means it should be consistently used to
// identify the same system across revisions of the document.
type SystemId struct {
	ID string `json:"id" yaml:"id"`
	// Identifies the identification system from which the provided identifier was
	// assigned.
	IdentifierType string `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
}

// Designates a referenced source catalog or profile that provides a source of
// control information for use in creating a new overlay or baseline.
type Import struct {
	ExcludeControls *[]SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	// A resolvable URL reference to the base catalog or profile that this profile
	// is tailoring.
	Href            string               `json:"href" yaml:"href"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
}

// Provides structuring directives that instruct how controls are organized
// after profile resolution.
type Merge struct {
	// Indicates that the controls selected should retain their original grouping
	// as defined in the import source.
	AsIs bool `json:"as-is,omitempty" yaml:"as-is,omitempty"`
	// A Combine element defines how to resolve duplicate instances of the same
	// control (e.g., controls with the same ID).
	Combine *CombinationRule `json:"combine,omitempty" yaml:"combine,omitempty"`
	// Provides an alternate grouping structure that selected controls will be
	// placed in.
	Custom *CustomGrouping `json:"custom,omitempty" yaml:"custom,omitempty"`
	// Directs that controls appear without any grouping structure.
	Flat *FlatWithoutGrouping `json:"flat,omitempty" yaml:"flat,omitempty"`
}

// Set parameters or amend controls in resolution.
type Modify struct {
	Alters        *[]Alteration       `json:"alters,omitempty" yaml:"alters,omitempty"`
	SetParameters *[]ParameterSetting `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

// Describes how the system satisfies a set of controls.
type ControlImplementation struct {
	// A statement describing important things to know about how this set of
	// control satisfaction documentation is approached.
	Description             string                   `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	SetParameters           *[]SetParameter          `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

// Used to import the OSCAL profile representing the system's control baseline.
type ImportProfile struct {
	// A resolvable URL reference to the profile or catalog to use as the system's
	// control baseline.
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Contains the characteristics of the system, such as its name, purpose, and
// security impact level.
type SystemCharacteristics struct {
	AuthorizationBoundary AuthorizationBoundary `json:"authorization-boundary" yaml:"authorization-boundary"`
	DataFlow              *DataFlow             `json:"data-flow,omitempty" yaml:"data-flow,omitempty"`
	DateAuthorized        string                `json:"date-authorized,omitempty" yaml:"date-authorized,omitempty"`
	// A summary of the system.
	Description         string               `json:"description" yaml:"description"`
	Links               *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	NetworkArchitecture *NetworkArchitecture `json:"network-architecture,omitempty" yaml:"network-architecture,omitempty"`
	Props               *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks             string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties  *[]ResponsibleParty  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	SecurityImpactLevel *SecurityImpactLevel `json:"security-impact-level,omitempty" yaml:"security-impact-level,omitempty"`
	// The overall information system sensitivity categorization, such as defined
	// by FIPS-199.
	SecuritySensitivityLevel string            `json:"security-sensitivity-level,omitempty" yaml:"security-sensitivity-level,omitempty"`
	Status                   Status            `json:"status" yaml:"status"`
	SystemIds                []SystemId        `json:"system-ids" yaml:"system-ids"`
	SystemInformation        SystemInformation `json:"system-information" yaml:"system-information"`
	// The full name of the system.
	SystemName string `json:"system-name" yaml:"system-name"`
	// A short name for the system, such as an acronym, that is suitable for
	// display in a data table or summary list.
	SystemNameShort string `json:"system-name-short,omitempty" yaml:"system-name-short,omitempty"`
}

// Provides information as to how the system is implemented.
type SystemImplementation struct {
	Components              []SystemComponent         `json:"components" yaml:"components"`
	InventoryItems          *[]InventoryItem          `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	LeveragedAuthorizations *[]LeveragedAuthorization `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Links                   *[]Link                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]Property               `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                 string                    `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                   []SystemUser              `json:"users" yaml:"users"`
}

// Used to represent the toolset used to perform aspects of the assessment.
type AssessmentPlatform struct {
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// The title or name for the assessment platform.
	Title          string           `json:"title,omitempty" yaml:"title,omitempty"`
	UsesComponents *[]UsesComponent `json:"uses-components,omitempty" yaml:"uses-components,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this assessment platform elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the assessment platform
	// can be used to reference the data item locally or globally (e.g., in an
	// imported OSCAL instance). This UUID should be assigned per-subject, which
	// means it should be consistently used to identify the same subject across
	// revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A defined component that can be part of an implemented system.
type SystemComponent struct {
	// A description of the component, including information about its function.
	Description string      `json:"description" yaml:"description"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols   *[]Protocol `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	// A summary of the technological or business purpose of the component.
	Purpose          string             `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// Describes the operational status of the system component.
	Status SystemComponentStatus `json:"status" yaml:"status"`
	// A human readable name for the system component.
	Title string `json:"title" yaml:"title"`
	// A category describing the purpose of the component.
	Type string `json:"type" yaml:"type"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this component elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the component can be used to
	// reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Identifies a set of assessment subjects to include/exclude by UUID.
type SelectSubjectById struct {
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A machine-oriented identifier reference to a component, inventory-item,
	// location, party, user, or resource using it's UUID.
	SubjectUuid string `json:"subject-uuid" yaml:"subject-uuid"`
	// Used to indicate the type of object pointed to by the uuid-ref within a
	// subject.
	Type string `json:"type" yaml:"type"`
}

// Include all controls from the imported catalog or profile resources.
type IncludeAll = map[string]interface{}

// A reference to a local or remote resource, that has a specific relation to
// the containing object.
type Link struct {
	// A resolvable URL reference to a resource.
	Href string `json:"href" yaml:"href"`
	// A label that indicates the nature of a resource, as a data serialization or
	// format.
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	// Describes the type of relationship provided by the link's hypertext
	// reference. This can be an indicator of the link's purpose.
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty"`
	// In case where the href points to a back-matter/resource, this value will
	// indicate the URI fragment to append to any rlink associated with the
	// resource. This value MUST be URI encoded.
	ResourceFragment string `json:"resource-fragment,omitempty" yaml:"resource-fragment,omitempty"`
	// A textual label to associate with the link, which may be used for
	// presentation in a tool.
	Text string `json:"text,omitempty" yaml:"text,omitempty"`
}

// An attribute, characteristic, or quality of the containing object expressed
// as a namespace qualified name/value pair.
type Property struct {
	// A textual label that provides a sub-type or characterization of the
	// property's name.
	Class string `json:"class,omitempty" yaml:"class,omitempty"`
	// An identifier for relating distinct sets of properties.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`
	// A textual label, within a namespace, that uniquely identifies a specific
	// attribute, characteristic, or quality of the property's containing object.
	Name string `json:"name" yaml:"name"`
	// A namespace qualifying the property's name. This allows different
	// organizations to associate distinct semantics with the same name.
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A unique identifier for a property.
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	// Indicates the value of the attribute, characteristic, or quality.
	Value string `json:"value" yaml:"value"`
}

// A resource associated with content in the containing document instance. A
// resource may be directly included in the document using base64 encoding or
// may point to one or more equivalent internet resources.
type Resource struct {
	// A resource encoded using the Base64 alphabet defined by RFC 2045.
	Base64 *Base64 `json:"base64,omitempty" yaml:"base64,omitempty"`
	// An optional citation consisting of end note text using structured markup.
	Citation *Citation `json:"citation,omitempty" yaml:"citation,omitempty"`
	// An optional short summary of the resource used to indicate the purpose of
	// the resource.
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds *[]DocumentId   `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props       *[]Property     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Rlinks      *[]ResourceLink `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	// An optional name given to the resource, which may be used by a tool for
	// display and navigation.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A unique identifier for a resource.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Identifies an assessment or related process that can be performed. In the
// assessment plan, this is an intended activity which may be associated with an
// assessment task. In the assessment results, this an activity that was
// actually performed as part of an assessment.
type Activity struct {
	// A human-readable description of this included activity.
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedControls  *ReviewedControls  `json:"related-controls,omitempty" yaml:"related-controls,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Steps            *[]Step            `json:"steps,omitempty" yaml:"steps,omitempty"`
	// The title for this included activity.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this assessment activity elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the activity can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A single managed inventory item within the system.
type InventoryItem struct {
	// A summary of the inventory item stating its purpose within the system.
	Description           string                  `json:"description" yaml:"description"`
	ImplementedComponents *[]ImplementedComponent `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Links                 *[]Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 *[]Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties    *[]ResponsibleParty     `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this inventory item elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the inventory item can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A local definition of a control objective for this assessment. Uses catalog
// syntax for control objective and assessment actions.
type LocalObjective struct {
	// A reference to a control with a corresponding id value. When referencing an
	// externally defined control, the Control Identifier Reference must be used in
	// the context of the external / imported OSCAL instance (e.g., uri-reference).
	ControlId string `json:"control-id" yaml:"control-id"`
	// A human-readable description of this control objective.
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Parts       []Part      `json:"parts" yaml:"parts"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// A type of user that interacts with the system based on an associated role.
type SystemUser struct {
	AuthorizedPrivileges *[]AuthorizedPrivilege `json:"authorized-privileges,omitempty" yaml:"authorized-privileges,omitempty"`
	// A summary of the user's purpose within the system.
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleIds     *[]string   `json:"role-ids,omitempty" yaml:"role-ids,omitempty"`
	// A short common name, abbreviation, or acronym for the user.
	ShortName string `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	// A name given to the user, which may be used by a tool for display and
	// navigation.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this user class elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the system user can be used to
	// reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// An action applied by a role within a given party to the content.
type Action struct {
	// The date and time when the action occurred.
	Date               *time.Time          `json:"date,omitempty" yaml:"date,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	// Specifies the action type system used.
	System string `json:"system" yaml:"system"`
	// The type of action documented by the assembly, such as an approval.
	Type string `json:"type" yaml:"type"`
	// A unique identifier that can be used to reference this defined action
	// elsewhere in an OSCAL document. A UUID should be consistently used for a
	// given location across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A document identifier qualified by an identifier scheme.
type DocumentId struct {
	Identifier string `json:"identifier" yaml:"identifier"`
	// Qualifies the kind of document identifier using a URI. If the scheme is not
	// provided the value of the element will be interpreted as a string of
	// characters.
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

// A physical point of presence, which may be associated with people,
// organizations, or other concepts within the current or linked OSCAL document.
type Location struct {
	Address          *Address           `json:"address,omitempty" yaml:"address,omitempty"`
	EmailAddresses   *[]string          `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TelephoneNumbers *[]TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	// A name given to the location, which may be used by a tool for display and
	// navigation.
	Title string    `json:"title,omitempty" yaml:"title,omitempty"`
	Urls  *[]string `json:"urls,omitempty" yaml:"urls,omitempty"`
	// A unique ID for the location, for reference.
	UUID string `json:"uuid" yaml:"uuid"`
}

// An organization or person, which may be associated with roles or other
// concepts within the current or linked OSCAL document.
type Party struct {
	Addresses             *[]Address                 `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	EmailAddresses        *[]string                  `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	ExternalIds           *[]PartyExternalIdentifier `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Links                 *[]Link                    `json:"links,omitempty" yaml:"links,omitempty"`
	LocationUuids         *[]string                  `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations *[]string                  `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	// The full name of the party. This is typically the legal name associated with
	// the party.
	Name    string      `json:"name,omitempty" yaml:"name,omitempty"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A short common name, abbreviation, or acronym for the party.
	ShortName        string             `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	TelephoneNumbers *[]TelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	// A category describing the kind of party the object describes.
	Type string `json:"type" yaml:"type"`
	// A unique identifier for the party.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A reference to a set of persons and/or organizations that have responsibility
// for performing the referenced role in the context of the containing object.
type ResponsibleParty struct {
	Links      *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string    `json:"party-uuids" yaml:"party-uuids"`
	Props      *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A reference to a role performed by a party.
	RoleId string `json:"role-id" yaml:"role-id"`
}

// An entry in a sequential list of revisions to the containing document,
// expected to be in reverse chronological order (i.e. latest first).
type RevisionHistoryEntry struct {
	LastModified *time.Time  `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Links        *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	OscalVersion string      `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Published    *time.Time  `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks      string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A name given to the document revision, which may be used by a tool for
	// display and navigation.
	Title   string `json:"title,omitempty" yaml:"title,omitempty"`
	Version string `json:"version" yaml:"version"`
}

// Defines a function, which might be assigned to a party in a specific
// situation.
type Role struct {
	// A summary of the role's purpose and associated responsibilities.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// A unique identifier for the role.
	ID      string      `json:"id" yaml:"id"`
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A short common name, abbreviation, or acronym for the role.
	ShortName string `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	// A name given to the role, which may be used by a tool for display and
	// navigation.
	Title string `json:"title" yaml:"title"`
}

// Identifies the control objectives of the assessment. In the assessment plan,
// these are the planned objectives. In the assessment results, these are the
// assessed objectives, and reflects any changes from the plan.
type ReferencedControlObjectives struct {
	// A human-readable description of this collection of control objectives.
	Description       string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeObjectives *[]SelectObjectiveById `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	IncludeAll        *IncludeAll            `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives *[]SelectObjectiveById `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	Links             *[]Link                `json:"links,omitempty" yaml:"links,omitempty"`
	Props             *[]Property            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks           string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Identifies the controls being assessed. In the assessment plan, these are the
// planned controls. In the assessment results, these are the actual controls,
// and reflects any changes from the plan.
type AssessedControls struct {
	// A human-readable description of in-scope controls specified for assessment.
	Description     string                               `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls *[]AssessedControlsSelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      *IncludeAll                          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]AssessedControlsSelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           *[]Link                              `json:"links,omitempty" yaml:"links,omitempty"`
	Props           *[]Property                          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string                               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Identifies an individual activity to be performed as part of a task.
type AssociatedActivity struct {
	// A machine-oriented identifier reference to an activity defined in the list
	// of activities.
	ActivityUuid     string              `json:"activity-uuid" yaml:"activity-uuid"`
	Links            *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole  `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

// Used to indicate that a task is dependent on another task.
type TaskDependency struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A machine-oriented identifier reference to a unique task.
	TaskUuid string `json:"task-uuid" yaml:"task-uuid"`
}

// A reference to a role with responsibility for performing a function relative
// to the containing object, optionally associated with a set of persons and/or
// organizations that perform that role.
type ResponsibleRole struct {
	Links      *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids *[]string   `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Props      *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A human-oriented identifier reference to a role performed.
	RoleId string `json:"role-id" yaml:"role-id"`
}

// The timing under which the task is intended to occur.
type EventTiming struct {
	// The task is intended to occur at the specified frequency.
	AtFrequency *FrequencyCondition `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	// The task is intended to occur on the specified date.
	OnDate *OnDateCondition `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	// The task is intended to occur within the specified date range.
	WithinDateRange *OnDateRangeCondition `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

// A partition of an assessment plan or results or a child of another part.
type AssessmentPart struct {
	// A textual label that provides a sub-type or characterization of the part's
	// name. This can be used to further distinguish or discriminate between the
	// semantics of multiple parts of the same control with the same name and ns.
	Class string  `json:"class,omitempty" yaml:"class,omitempty"`
	Links *[]Link `json:"links,omitempty" yaml:"links,omitempty"`
	// A textual label that uniquely identifies the part's semantic type.
	Name string `json:"name" yaml:"name"`
	// A namespace qualifying the part's name. This allows different organizations
	// to associate distinct semantics with the same name.
	Ns    string            `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts *[]AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
	// Permits multiple paragraphs, lists, tables etc.
	Prose string `json:"prose,omitempty" yaml:"prose,omitempty"`
	// A name given to the part, which may be used by a tool for display and
	// navigation.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this part elsewhere in this or other OSCAL
	// instances. The locally defined UUID of the part can be used to reference the
	// data item locally or globally (e.g., in an ported OSCAL instance). This UUID
	// should be assigned per-subject, which means it should be consistently used
	// to identify the same subject across revisions of the document.
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// A log of all assessment-related actions taken.
type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

// A set of textual statements, typically written by the assessor.
type AttestationStatements struct {
	Parts              []AssessmentPart    `json:"parts" yaml:"parts"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

// An annotated, markup-based textual element of a control's or catalog group's
// definition, or a child of another part.
type Part struct {
	// An optional textual providing a sub-type or characterization of the part's
	// name, or a category to which the part belongs.
	Class string `json:"class,omitempty" yaml:"class,omitempty"`
	// A unique identifier for the part.
	ID    string  `json:"id,omitempty" yaml:"id,omitempty"`
	Links *[]Link `json:"links,omitempty" yaml:"links,omitempty"`
	// A textual label that uniquely identifies the part's semantic type, which
	// exists in a value space qualified by the ns.
	Name string `json:"name" yaml:"name"`
	// An optional namespace qualifying the part's name. This allows different
	// organizations to associate distinct semantics with the same name.
	Ns    string      `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts *[]Part     `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// Permits multiple paragraphs, lists, tables etc.
	Prose string `json:"prose,omitempty" yaml:"prose,omitempty"`
	// An optional name given to the part, which may be used by a tool for display
	// and navigation.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

// A formal or informal expression of a constraint or test.
type ParameterConstraint struct {
	// A textual summary of the constraint to be applied.
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Tests       *[]ConstraintTest `json:"tests,omitempty" yaml:"tests,omitempty"`
}

// A prose statement that provides a recommendation for the use of a parameter.
type ParameterGuideline struct {
	// Prose permits multiple paragraphs, lists, tables etc.
	Prose string `json:"prose" yaml:"prose"`
}

// Presenting a choice among alternatives.
type ParameterSelection struct {
	Choice *[]string `json:"choice,omitempty" yaml:"choice,omitempty"`
	// Describes the number of selections that must occur. Without this setting,
	// only one value should be assumed to be permitted.
	HowMany string `json:"how-many,omitempty" yaml:"how-many,omitempty"`
}

// Defines how the component or capability supports a set of controls.
type ControlImplementationSet struct {
	// A description of how the specified set of controls are implemented for the
	// containing component or capability.
	Description             string                                        `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirementControlImplementation `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   *[]Link                                       `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]Property                                   `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           *[]SetParameter                               `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	// A reference to an OSCAL catalog or profile providing the referenced control
	// or subcontrol definition.
	Source string `json:"source" yaml:"source"`
	// Provides a means to identify a set of control implementations that are
	// supported by a given component or capability.
	UUID string `json:"uuid" yaml:"uuid"`
}

// The collection of components comprising this capability.
type IncorporatesComponent struct {
	// A machine-oriented identifier reference to a component.
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	// A description of the component, including information about its function.
	Description string `json:"description" yaml:"description"`
}

// Information about the protocol used to provide a service.
type Protocol struct {
	// The common name of the protocol, which should be the appropriate "service
	// name" from the IANA Service Name and Transport Protocol Port Number
	// Registry.
	Name       string       `json:"name" yaml:"name"`
	PortRanges *[]PortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	// A human readable name for the protocol (e.g., Transport Layer Security).
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this service protocol information elsewhere in
	// this or other OSCAL instances. The locally defined UUID of the service
	// protocol can be used to reference the data item locally or globally (e.g.,
	// in an imported OSCAL instance). This UUID should be assigned per-subject,
	// which means it should be consistently used to identify the same subject
	// across revisions of the document.
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// Identifies the source of the finding, such as a tool, interviewed person, or
// activity.
type Origin struct {
	Actors       []OriginActor  `json:"actors" yaml:"actors"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

// Relates the finding to a set of referenced observations that were used to
// determine the finding.
type RelatedObservation struct {
	// A machine-oriented identifier reference to an observation defined in the
	// list of observations.
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

// Relates the finding to a set of referenced risks that were used to determine
// the finding.
type AssociatedRisk struct {
	// A machine-oriented identifier reference to a risk defined in the list of
	// risks.
	RiskUuid string `json:"risk-uuid" yaml:"risk-uuid"`
}

// Captures an assessor's conclusions regarding the degree to which an objective
// is satisfied.
type FindingTarget struct {
	// A human-readable description of the assessor's conclusions regarding the
	// degree to which an objective is satisfied.
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	ImplementationStatus *ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Links                *[]Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                *[]Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A determination of if the objective is satisfied or not within a given
	// system.
	Status ObjectiveStatus `json:"status" yaml:"status"`
	// A machine-oriented identifier reference for a specific target qualified by
	// the type.
	TargetId string `json:"target-id" yaml:"target-id"`
	// The title for this objective status.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// Identifies the type of the target.
	Type string `json:"type" yaml:"type"`
}

// Links this observation to relevant evidence.
type RelevantEvidence struct {
	// A human-readable description of this evidence.
	Description string `json:"description" yaml:"description"`
	// A resolvable URL reference to relevant evidence.
	Href    string      `json:"href,omitempty" yaml:"href,omitempty"`
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// A human-oriented identifier reference to a resource. Use type to indicate
// whether the identified resource is a component, inventory item, location,
// user, or something else.
type SubjectReference struct {
	Links   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A machine-oriented identifier reference to a component, inventory-item,
	// location, party, user, or resource using it's UUID.
	SubjectUuid string `json:"subject-uuid" yaml:"subject-uuid"`
	// The title or name for the referenced subject.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// Used to indicate the type of object pointed to by the uuid-ref within a
	// subject.
	Type string `json:"type" yaml:"type"`
}

// Identifies the source of the finding, such as a tool or person.
type PoamItemOrigin struct {
	Actors []OriginActor `json:"actors" yaml:"actors"`
}

// Relates the poam-item to referenced finding(s).
type RelatedFinding struct {
	// A machine-oriented identifier reference to a finding defined in the list of
	// findings.
	FindingUuid string `json:"finding-uuid" yaml:"finding-uuid"`
}

// A collection of descriptive data about the containing object from a specific
// origin.
type Characterization struct {
	Facets []Facet     `json:"facets" yaml:"facets"`
	Links  *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Origin Origin      `json:"origin" yaml:"origin"`
	Props  *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
}

// Describes an existing mitigating factor that may affect the overall
// determination of the risk, with an optional link to an implementation
// statement in the SSP.
type MitigatingFactor struct {
	// A human-readable description of this mitigating factor.
	Description string `json:"description" yaml:"description"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this implementation statement elsewhere in
	// this or other OSCAL instancess. The locally defined UUID of the
	// implementation statement can be used to reference the data item locally or
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned
	// per-subject, which means it should be consistently used to identify the same
	// subject across revisions of the document.
	ImplementationUuid string              `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Subjects           *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this mitigating factor elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the mitigating factor can
	// be used to reference the data item locally or globally (e.g., in an imported
	// OSCAL instance). This UUID should be assigned per-subject, which means it
	// should be consistently used to identify the same subject across revisions of
	// the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Describes either recommended or an actual plan for addressing the risk.
type Response struct {
	// A human-readable description of this response plan.
	Description string `json:"description" yaml:"description"`
	// Identifies whether this is a recommendation, such as from an assessor or
	// tool, or an actual plan accepted by the system owner.
	Lifecycle      string           `json:"lifecycle" yaml:"lifecycle"`
	Links          *[]Link          `json:"links,omitempty" yaml:"links,omitempty"`
	Origins        *[]Origin        `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props          *[]Property      `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RequiredAssets *[]RequiredAsset `json:"required-assets,omitempty" yaml:"required-assets,omitempty"`
	Tasks          *[]Task          `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	// The title for this response activity.
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this remediation elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the risk response can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A log of all risk-related tasks taken.
type RiskLog struct {
	Entries []RiskLogEntry `json:"entries" yaml:"entries"`
}

// A pointer, by ID, to an externally-defined threat.
type ThreatId struct {
	// An optional location for the threat data, from which this ID originates.
	Href string `json:"href,omitempty" yaml:"href,omitempty"`
	ID   string `json:"id" yaml:"id"`
	// Specifies the source of the threat information.
	System string `json:"system" yaml:"system"`
}

// Select a control or controls from an imported control set.
type SelectControlById struct {
	Matching *[]Matching `json:"matching,omitempty" yaml:"matching,omitempty"`
	// When a control is included, whether its child (dependent) controls are also
	// included.
	WithChildControls string    `json:"with-child-controls,omitempty" yaml:"with-child-controls,omitempty"`
	WithIds           *[]string `json:"with-ids,omitempty" yaml:"with-ids,omitempty"`
}

// A Combine element defines how to resolve duplicate instances of the same
// control (e.g., controls with the same ID).
type CombinationRule struct {
	// Declare how clashing controls should be handled.
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

// Provides an alternate grouping structure that selected controls will be
// placed in.
type CustomGrouping struct {
	Groups         *[]CustomGroupingGroup `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls *[]InsertControls      `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
}

// Directs that controls appear without any grouping structure.
type FlatWithoutGrouping = map[string]interface{}

// Specifies changes to be made to an included control when a profile is
// resolved.
type Alteration struct {
	Adds *[]Addition `json:"adds,omitempty" yaml:"adds,omitempty"`
	// A reference to a control with a corresponding id value. When referencing an
	// externally defined control, the Control Identifier Reference must be used in
	// the context of the external / imported OSCAL instance (e.g., uri-reference).
	ControlId string     `json:"control-id" yaml:"control-id"`
	Removes   *[]Removal `json:"removes,omitempty" yaml:"removes,omitempty"`
}

// A parameter setting, to be propagated to points of insertion.
type ParameterSetting struct {
	// A textual label that provides a characterization of the parameter.
	Class       string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints *[]ParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	// **(deprecated)** Another parameter invoking this one. This construct has
	// been deprecated and should not be used.
	DependsOn  string                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines *[]ParameterGuideline `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	// A short, placeholder name for the parameter, which can be used as a
	// substitute for a value if no value is assigned.
	Label string  `json:"label,omitempty" yaml:"label,omitempty"`
	Links *[]Link `json:"links,omitempty" yaml:"links,omitempty"`
	// An identifier for the parameter.
	ParamId string              `json:"param-id" yaml:"param-id"`
	Props   *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Select  *ParameterSelection `json:"select,omitempty" yaml:"select,omitempty"`
	// Describes the purpose and use of a parameter.
	Usage  string    `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values *[]string `json:"values,omitempty" yaml:"values,omitempty"`
}

// Describes how the system satisfies the requirements of an individual control.
type ImplementedRequirement struct {
	ByComponents *[]ByComponent `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	// A reference to a control with a corresponding id value. When referencing an
	// externally defined control, the Control Identifier Reference must be used in
	// the context of the external / imported OSCAL instance (e.g., uri-reference).
	ControlId        string             `json:"control-id" yaml:"control-id"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    *[]SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       *[]Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this control requirement elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the control requirement
	// can be used to reference the data item locally or globally (e.g., in an
	// imported OSCAL instance). This UUID should be assigned per-subject, which
	// means it should be consistently used to identify the same subject across
	// revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Identifies the parameter that will be set by the enclosed value.
type SetParameter struct {
	// A human-oriented reference to a parameter within a control, who's catalog
	// has been imported into the current implementation context.
	ParamId string   `json:"param-id" yaml:"param-id"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Values  []string `json:"values" yaml:"values"`
}

// A description of this system's authorization boundary, optionally
// supplemented by diagrams that illustrate the authorization boundary.
type AuthorizationBoundary struct {
	// A summary of the system's authorization boundary.
	Description string      `json:"description" yaml:"description"`
	Diagrams    *[]Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// A description of the logical flow of information within the system and across
// its boundaries, optionally supplemented by diagrams that illustrate these
// flows.
type DataFlow struct {
	// A summary of the system's data flow.
	Description string      `json:"description" yaml:"description"`
	Diagrams    *[]Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// A description of the system's network architecture, optionally supplemented
// by diagrams that illustrate the network architecture.
type NetworkArchitecture struct {
	// A summary of the system's network architecture.
	Description string      `json:"description" yaml:"description"`
	Diagrams    *[]Diagram  `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// The overall level of expected impact resulting from unauthorized disclosure,
// modification, or loss of access to information.
type SecurityImpactLevel struct {
	// A target-level of availability for the system, based on the sensitivity of
	// information within the system.
	SecurityObjectiveAvailability string `json:"security-objective-availability" yaml:"security-objective-availability"`
	// A target-level of confidentiality for the system, based on the sensitivity
	// of information within the system.
	SecurityObjectiveConfidentiality string `json:"security-objective-confidentiality" yaml:"security-objective-confidentiality"`
	// A target-level of integrity for the system, based on the sensitivity of
	// information within the system.
	SecurityObjectiveIntegrity string `json:"security-objective-integrity" yaml:"security-objective-integrity"`
}

// Describes the operational status of the system.
type Status struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// The current operating status.
	State string `json:"state" yaml:"state"`
}

// Contains details about all information types that are stored, processed, or
// transmitted by the system, such as privacy information, and those defined in
// NIST SP 800-60.
type SystemInformation struct {
	InformationTypes []InformationType `json:"information-types" yaml:"information-types"`
	Links            *[]Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
}

// A description of another authorized system from which this system inherits
// capabilities that satisfy security requirements. Another term for this
// concept is a common control provider.
type LeveragedAuthorization struct {
	DateAuthorized string  `json:"date-authorized" yaml:"date-authorized"`
	Links          *[]Link `json:"links,omitempty" yaml:"links,omitempty"`
	// A machine-oriented identifier reference to the party that manages the
	// leveraged system.
	PartyUuid string      `json:"party-uuid" yaml:"party-uuid"`
	Props     *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks   string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A human readable name for the leveraged authorization in the context of the
	// system.
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope and
	// can be used to reference this leveraged authorization elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the leveraged
	// authorization can be used to reference the data item locally or globally
	// (e.g., in an imported OSCAL instance). This UUID should be assigned
	// per-subject, which means it should be consistently used to identify the same
	// subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// The set of components that are used by the assessment platform.
type UsesComponent struct {
	// A machine-oriented identifier reference to a component that is implemented
	// as part of an inventory item.
	ComponentUuid      string              `json:"component-uuid" yaml:"component-uuid"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

// Describes the operational status of the system component.
type SystemComponentStatus struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// The operational status.
	State string `json:"state" yaml:"state"`
}

// A resource encoded using the Base64 alphabet defined by RFC 2045.
type Base64 struct {
	// Name of the file before it was encoded as Base64 to be embedded in a
	// resource. This is the name that will be assigned to the file when the file
	// is decoded.
	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`
	// A label that indicates the nature of a resource, as a data serialization or
	// format.
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

// An optional citation consisting of end note text using structured markup.
type Citation struct {
	Links *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// A line of citation text.
	Text string `json:"text" yaml:"text"`
}

// A URL-based pointer to an external resource with an optional hash for
// verification and change detection.
type ResourceLink struct {
	Hashes *[]Hash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	// A resolvable URL pointing to the referenced resource.
	Href string `json:"href" yaml:"href"`
	// A label that indicates the nature of a resource, as a data serialization or
	// format.
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

// Identifies an individual step in a series of steps related to an activity,
// such as an assessment test or examination procedure.
type Step struct {
	// A human-readable description of this step.
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ReviewedControls *ReviewedControls  `json:"reviewed-controls,omitempty" yaml:"reviewed-controls,omitempty"`
	// The title for this step.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this step elsewhere in this or other OSCAL
	// instances. The locally defined UUID of the step (in a series of steps) can
	// be used to reference the data item locally or globally (e.g., in an imported
	// OSCAL instance). This UUID should be assigned per-subject, which means it
	// should be consistently used to identify the same subject across revisions of
	// the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// The set of components that are implemented in a given system inventory item.
type ImplementedComponent struct {
	// A machine-oriented identifier reference to a component that is implemented
	// as part of an inventory item.
	ComponentUuid      string              `json:"component-uuid" yaml:"component-uuid"`
	Links              *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

// Identifies a specific system privilege held by the user, along with an
// associated description and/or rationale for the privilege.
type AuthorizedPrivilege struct {
	// A summary of the privilege's purpose within the system.
	Description        string   `json:"description,omitempty" yaml:"description,omitempty"`
	FunctionsPerformed []string `json:"functions-performed" yaml:"functions-performed"`
	// A human readable name for the privilege.
	Title string `json:"title" yaml:"title"`
}

// A postal address for the location.
type Address struct {
	AddrLines *[]string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	// City, town or geographical region for the mailing address.
	City string `json:"city,omitempty" yaml:"city,omitempty"`
	// The ISO 3166-1 alpha-2 country code for the mailing address.
	Country string `json:"country,omitempty" yaml:"country,omitempty"`
	// Postal or ZIP code for mailing address.
	PostalCode string `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	// State, province or analogous geographical region for a mailing address.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
	// Indicates the type of address.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// A telephone service number as defined by ITU-T E.164.
type TelephoneNumber struct {
	Number string `json:"number" yaml:"number"`
	// Indicates the type of phone number.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// An identifier for a person or organization using a designated scheme. e.g. an
// Open Researcher and Contributor ID (ORCID).
type PartyExternalIdentifier struct {
	ID string `json:"id" yaml:"id"`
	// Indicates the type of external identifier.
	Scheme string `json:"scheme" yaml:"scheme"`
}

// Used to select a control objective for inclusion/exclusion based on the
// control objective's identifier.
type SelectObjectiveById struct {
	// Points to an assessment objective.
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

// Used to select a control for inclusion/exclusion based on one or more control
// identifiers. A set of statement identifiers can be used to target the
// inclusion/exclusion to only specific control statements providing more
// granularity over the specific statements that are within the asessment scope.
type AssessedControlsSelectControlById struct {
	// A reference to a control with a corresponding id value. When referencing an
	// externally defined control, the Control Identifier Reference must be used in
	// the context of the external / imported OSCAL instance (e.g., uri-reference).
	ControlId    string    `json:"control-id" yaml:"control-id"`
	StatementIds *[]string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

// The task is intended to occur at the specified frequency.
type FrequencyCondition struct {
	// The task must occur after the specified period has elapsed.
	Period int `json:"period" yaml:"period"`
	// The unit of time for the period.
	Unit string `json:"unit" yaml:"unit"`
}

// The task is intended to occur on the specified date.
type OnDateCondition struct {
	// The task must occur on the specified date.
	Date time.Time `json:"date" yaml:"date"`
}

// The task is intended to occur within the specified date range.
type OnDateRangeCondition struct {
	// The task must occur on or before the specified date.
	End time.Time `json:"end" yaml:"end"`
	// The task must occur on or after the specified date.
	Start time.Time `json:"start" yaml:"start"`
}

// Identifies the result of an action and/or task that occurred as part of
// executing an assessment plan or an assessment event that occurred in
// producing the assessment results.
type AssessmentLogEntry struct {
	// A human-readable description of this event.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Identifies the end date and time of an event. If the event is a point in
	// time, the start and end will be the same date and time.
	End          *time.Time     `json:"end,omitempty" yaml:"end,omitempty"`
	Links        *[]Link        `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy     *[]LoggedBy    `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props        *[]Property    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// Identifies the start date and time of an event.
	Start time.Time `json:"start" yaml:"start"`
	// The title for this event.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference an assessment event in this or other OSCAL
	// instances. The locally defined UUID of the assessment log entry can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// A test expression which is expected to be evaluated by a tool.
type ConstraintTest struct {
	// A formal (executable) expression of a constraint.
	Expression string `json:"expression" yaml:"expression"`
	Remarks    string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

// Describes how the containing component or capability implements an individual
// control.
type ImplementedRequirementControlImplementation struct {
	// A reference to a control with a corresponding id value. When referencing an
	// externally defined control, the Control Identifier Reference must be used in
	// the context of the external / imported OSCAL instance (e.g., uri-reference).
	ControlId string `json:"control-id" yaml:"control-id"`
	// A suggestion from the supplier (e.g., component vendor or author) for how
	// the specified control may be implemented if the containing component or
	// capability is instantiated in a system security plan.
	Description      string                            `json:"description" yaml:"description"`
	Links            *[]Link                           `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property                       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole                `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    *[]SetParameter                   `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       *[]ControlStatementImplementation `json:"statements,omitempty" yaml:"statements,omitempty"`
	// Provides a globally unique means to identify a given control implementation
	// by a component.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Where applicable this is the IPv4 port range on which the service operates.
type PortRange struct {
	// Indicates the ending port number in a port range
	End int `json:"end,omitempty" yaml:"end,omitempty"`
	// Indicates the starting port number in a port range
	Start int `json:"start,omitempty" yaml:"start,omitempty"`
	// Indicates the transport type.
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

// The actor that produces an observation, a finding, or a risk. One or more
// actor type can be used to specify a person that is using a tool.
type OriginActor struct {
	// A machine-oriented identifier reference to the tool or person based on the
	// associated type.
	ActorUuid string      `json:"actor-uuid" yaml:"actor-uuid"`
	Links     *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props     *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// For a party, this can optionally be used to specify the role the actor was
	// performing.
	RoleId string `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	// The kind of actor.
	Type string `json:"type" yaml:"type"`
}

// Identifies an individual task for which the containing object is a
// consequence of.
type RelatedTask struct {
	// Used to detail assessment subjects that were identfied by this task.
	IdentifiedSubject  *IdentifiedSubject   `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Links              *[]Link              `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]Property          `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]ResponsibleParty  `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Subjects           *[]AssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	// A machine-oriented identifier reference to a unique task.
	TaskUuid string `json:"task-uuid" yaml:"task-uuid"`
}

// Indicates the degree to which the a given control is implemented.
type ImplementationStatus struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// Identifies the implementation status of the control or control objective.
	State string `json:"state" yaml:"state"`
}

// A determination of if the objective is satisfied or not within a given
// system.
type ObjectiveStatus struct {
	// The reason the objective was given it's status.
	Reason  string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// An indication as to whether the objective is satisfied or not.
	State string `json:"state" yaml:"state"`
}

// An individual characteristic that is part of a larger set produced by the
// same actor.
type Facet struct {
	Links *[]Link `json:"links,omitempty" yaml:"links,omitempty"`
	// The name of the risk metric within the specified system.
	Name    string      `json:"name" yaml:"name"`
	Props   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// Specifies the naming system under which this risk metric is organized, which
	// allows for the same names to be used in different systems controlled by
	// different parties. This avoids the potential of a name clash.
	System string `json:"system" yaml:"system"`
	// Indicates the value of the facet.
	Value string `json:"value" yaml:"value"`
}

// Identifies an asset required to achieve remediation.
type RequiredAsset struct {
	// A human-readable description of this required asset.
	Description string              `json:"description" yaml:"description"`
	Links       *[]Link             `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects    *[]SubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	// The title for this required asset.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this required asset elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the asset can be used to
	// reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Identifies an individual risk response that occurred as part of managing an
// identified risk.
type RiskLogEntry struct {
	// A human-readable description of what was done regarding the risk.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Identifies the end date and time of the event. If the event is a point in
	// time, the start and end will be the same date and time.
	End              *time.Time               `json:"end,omitempty" yaml:"end,omitempty"`
	Links            *[]Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy         *[]LoggedBy              `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props            *[]Property              `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedResponses *[]RiskResponseReference `json:"related-responses,omitempty" yaml:"related-responses,omitempty"`
	Remarks          string                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// Identifies the start date and time of the event.
	Start        time.Time `json:"start" yaml:"start"`
	StatusChange string    `json:"status-change,omitempty" yaml:"status-change,omitempty"`
	// The title for this risk log entry.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this risk log entry elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the risk log entry can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Selecting a set of controls by matching their IDs with a wildcard pattern.
type Matching struct {
	// A glob expression matching the IDs of one or more controls to be selected.
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

// A group of (selected) controls or of groups of controls.
type CustomGroupingGroup struct {
	// A textual label that provides a sub-type or characterization of the group.
	Class  string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Groups *[]CustomGroupingGroup `json:"groups,omitempty" yaml:"groups,omitempty"`
	// Identifies the group.
	ID             string            `json:"id,omitempty" yaml:"id,omitempty"`
	InsertControls *[]InsertControls `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
	Links          *[]Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Params         *[]Parameter      `json:"params,omitempty" yaml:"params,omitempty"`
	Parts          *[]Part           `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props          *[]Property       `json:"props,omitempty" yaml:"props,omitempty"`
	// A name to be given to the group for use in display.
	Title string `json:"title" yaml:"title"`
}

// Specifies which controls to use in the containing context.
type InsertControls struct {
	ExcludeControls *[]SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      *IncludeAll          `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	// A designation of how a selection of controls in a profile is to be ordered.
	Order string `json:"order,omitempty" yaml:"order,omitempty"`
}

// Specifies contents to be added into controls, in resolution.
type Addition struct {
	// Target location of the addition.
	ById   string       `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	Links  *[]Link      `json:"links,omitempty" yaml:"links,omitempty"`
	Params *[]Parameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts  *[]Part      `json:"parts,omitempty" yaml:"parts,omitempty"`
	// Where to add the new content with respect to the targeted element (beside it
	// or inside it).
	Position string      `json:"position,omitempty" yaml:"position,omitempty"`
	Props    *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// A name given to the control, which may be used by a tool for display and
	// navigation.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}

// Specifies objects to be removed from a control based on specific aspects of
// the object that must all match.
type Removal struct {
	// Identify items to remove by matching their class.
	ByClass string `json:"by-class,omitempty" yaml:"by-class,omitempty"`
	// Identify items to remove indicated by their id.
	ById string `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	// Identify items to remove by the name of the item's information object name,
	// e.g. title or prop.
	ByItemName string `json:"by-item-name,omitempty" yaml:"by-item-name,omitempty"`
	// Identify items remove by matching their assigned name.
	ByName string `json:"by-name,omitempty" yaml:"by-name,omitempty"`
	// Identify items to remove by the item's ns, which is the namespace associated
	// with a part, or prop.
	ByNs string `json:"by-ns,omitempty" yaml:"by-ns,omitempty"`
}

// Defines how the referenced component implements a set of controls.
type ByComponent struct {
	// A machine-oriented identifier reference to the component that is implemeting
	// a given control.
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	// An implementation statement that describes how a control or a control
	// statement is implemented within the referenced system component.
	Description string `json:"description" yaml:"description"`
	// Identifies content intended for external consumption, such as with leveraged
	// organizations.
	Export               *Export                                         `json:"export,omitempty" yaml:"export,omitempty"`
	ImplementationStatus *ImplementationStatus                           `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Inherited            *[]InheritedControlImplementation               `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	Links                *[]Link                                         `json:"links,omitempty" yaml:"links,omitempty"`
	Props                *[]Property                                     `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                                          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     *[]ResponsibleRole                              `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Satisfied            *[]SatisfiedControlImplementationResponsibility `json:"satisfied,omitempty" yaml:"satisfied,omitempty"`
	SetParameters        *[]SetParameter                                 `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this by-component entry elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the by-component entry
	// can be used to reference the data item locally or globally (e.g., in an
	// imported OSCAL instance). This UUID should be assigned per-subject, which
	// means it should be consistently used to identify the same subject across
	// revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Identifies which statements within a control are addressed.
type Statement struct {
	ByComponents     *[]ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A human-oriented identifier reference to a control statement.
	StatementId string `json:"statement-id" yaml:"statement-id"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this control statement elsewhere in this or
	// other OSCAL instances. The UUID of the control statement in the source OSCAL
	// instance is sufficient to reference the data item locally or globally (e.g.,
	// in an imported OSCAL instance).
	UUID string `json:"uuid" yaml:"uuid"`
}

// A graphic that provides a visual representation the system, or some aspect of
// it.
type Diagram struct {
	// A brief caption to annotate the diagram.
	Caption string `json:"caption,omitempty" yaml:"caption,omitempty"`
	// A summary of the diagram.
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this diagram elsewhere in this or other OSCAL
	// instances. The locally defined UUID of the diagram can be used to reference
	// the data item locally or globally (e.g., in an imported OSCAL instance).
	// This UUID should be assigned per-subject, which means it should be
	// consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Contains details about one information type that is stored, processed, or
// transmitted by the system, such as privacy information, and those defined in
// NIST SP 800-60.
type InformationType struct {
	AvailabilityImpact    *Impact                          `json:"availability-impact,omitempty" yaml:"availability-impact,omitempty"`
	Categorizations       *[]InformationTypeCategorization `json:"categorizations,omitempty" yaml:"categorizations,omitempty"`
	ConfidentialityImpact *Impact                          `json:"confidentiality-impact,omitempty" yaml:"confidentiality-impact,omitempty"`
	// A summary of how this information type is used within the system.
	Description     string      `json:"description" yaml:"description"`
	IntegrityImpact *Impact     `json:"integrity-impact,omitempty" yaml:"integrity-impact,omitempty"`
	Links           *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props           *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// A human readable name for the information type. This title should be
	// meaningful within the context of the system.
	Title string `json:"title" yaml:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this information type elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the information type can
	// be used to reference the data item locally or globally (e.g., in an imported
	// OSCAL instance). This UUID should be assigned per-subject, which means it
	// should be consistently used to identify the same subject across revisions of
	// the document.
	UUID string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

// A representation of a cryptographic digest generated over a resource using a
// specified hash algorithm.
type Hash struct {
	// The digest method by which a hash is derived.
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

// Used to indicate who created a log entry in what role.
type LoggedBy struct {
	// A machine-oriented identifier reference to the party who is making the log
	// entry.
	PartyUuid string `json:"party-uuid" yaml:"party-uuid"`
	// A point to the role-id of the role in which the party is making the log
	// entry.
	RoleId string `json:"role-id,omitempty" yaml:"role-id,omitempty"`
}

// Identifies which statements within a control are addressed.
type ControlStatementImplementation struct {
	// A summary of how the containing control statement is implemented by the
	// component or capability.
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A human-oriented identifier reference to a control statement.
	StatementId string `json:"statement-id" yaml:"statement-id"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this control statement elsewhere in this or
	// other OSCAL instances. The UUID of the control statement in the source OSCAL
	// instance is sufficient to reference the data item locally or globally (e.g.,
	// in an imported OSCAL instance).
	UUID string `json:"uuid" yaml:"uuid"`
}

// Used to detail assessment subjects that were identfied by this task.
type IdentifiedSubject struct {
	// A machine-oriented identifier reference to a unique assessment subject
	// placeholder defined by this task.
	SubjectPlaceholderUuid string              `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []AssessmentSubject `json:"subjects" yaml:"subjects"`
}

// Identifies an individual risk response that this log entry is for.
type RiskResponseReference struct {
	Links        *[]Link        `json:"links,omitempty" yaml:"links,omitempty"`
	Props        *[]Property    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks *[]RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string         `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A machine-oriented identifier reference to a unique risk response.
	ResponseUuid string `json:"response-uuid" yaml:"response-uuid"`
}

// Identifies content intended for external consumption, such as with leveraged
// organizations.
type Export struct {
	// An implementation statement that describes the aspects of the control or
	// control statement implementation that can be available to another system
	// leveraging this system.
	Description      string                                 `json:"description,omitempty" yaml:"description,omitempty"`
	Links            *[]Link                                `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property                            `json:"props,omitempty" yaml:"props,omitempty"`
	Provided         *[]ProvidedControlImplementation       `json:"provided,omitempty" yaml:"provided,omitempty"`
	Remarks          string                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Responsibilities *[]ControlImplementationResponsibility `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
}

// Describes a control implementation inherited by a leveraging system.
type InheritedControlImplementation struct {
	// An implementation statement that describes the aspects of a control or
	// control statement implementation that a leveraging system is inheriting from
	// a leveraged system.
	Description string      `json:"description" yaml:"description"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// A machine-oriented identifier reference to an inherited control
	// implementation that a leveraging system is inheriting from a leveraged
	// system.
	ProvidedUuid     string             `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this inherited entry elsewhere in this or
	// other OSCAL instances. The locally defined UUID of the inherited control
	// implementation can be used to reference the data item locally or globally
	// (e.g., in an imported OSCAL instance). This UUID should be assigned
	// per-subject, which means it should be consistently used to identify the same
	// subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Describes how this system satisfies a responsibility imposed by a leveraged
// system.
type SatisfiedControlImplementationResponsibility struct {
	// An implementation statement that describes the aspects of a control or
	// control statement implementation that a leveraging system is implementing
	// based on a requirement from a leveraged system.
	Description string      `json:"description" yaml:"description"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	// A machine-oriented identifier reference to a control implementation that
	// satisfies a responsibility imposed by a leveraged system.
	ResponsibilityUuid string             `json:"responsibility-uuid,omitempty" yaml:"responsibility-uuid,omitempty"`
	ResponsibleRoles   *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this satisfied control implementation entry
	// elsewhere in this or other OSCAL instances. The locally defined UUID of the
	// control implementation can be used to reference the data item locally or
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned
	// per-subject, which means it should be consistently used to identify the same
	// subject across revisions of the document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// The expected level of impact resulting from the described information.
type Impact struct {
	AdjustmentJustification string      `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string      `json:"base" yaml:"base"`
	Links                   *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                string      `json:"selected,omitempty" yaml:"selected,omitempty"`
}

// A set of information type identifiers qualified by the given identification
// system used, such as NIST SP 800-60.
type InformationTypeCategorization struct {
	InformationTypeIds *[]string `json:"information-type-ids,omitempty" yaml:"information-type-ids,omitempty"`
	// Specifies the information type identification system used.
	System string `json:"system" yaml:"system"`
}

// Describes a capability which may be inherited by a leveraging system.
type ProvidedControlImplementation struct {
	// An implementation statement that describes the aspects of the control or
	// control statement implementation that can be provided to another system
	// leveraging this system.
	Description      string             `json:"description" yaml:"description"`
	Links            *[]Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this provided entry elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the provided entry can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}

// Describes a control implementation responsibility imposed on a leveraging
// system.
type ControlImplementationResponsibility struct {
	// An implementation statement that describes the aspects of the control or
	// control statement implementation that a leveraging system must implement to
	// satisfy the control provided by a leveraged system.
	Description string      `json:"description" yaml:"description"`
	Links       *[]Link     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]Property `json:"props,omitempty" yaml:"props,omitempty"`
	// A machine-oriented identifier reference to an inherited control
	// implementation that a leveraging system is inheriting from a leveraged
	// system.
	ProvidedUuid     string             `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Remarks          string             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope
	// that can be used to reference this responsibility elsewhere in this or other
	// OSCAL instances. The locally defined UUID of the responsibility can be used
	// to reference the data item locally or globally (e.g., in an imported OSCAL
	// instance). This UUID should be assigned per-subject, which means it should
	// be consistently used to identify the same subject across revisions of the
	// document.
	UUID string `json:"uuid" yaml:"uuid"`
}
