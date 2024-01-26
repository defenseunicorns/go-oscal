package diffTypes

type RiskLogEntry struct {
	Description      string                  `json:"description,omitempty" yaml:"description,omitempty"`
	End              string                  `json:"end,omitempty" yaml:"end,omitempty"`
	Links            []Link                  `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy         []LoggedBy              `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props            []Property              `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedResponses []RiskResponseReference `json:"related-responses,omitempty" yaml:"related-responses,omitempty"`
	Remarks          string                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start            string                  `json:"start" yaml:"start"`
	StatusChange     string                  `json:"status-change,omitempty" yaml:"status-change,omitempty"`
	Title            string                  `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                  `json:"uuid" yaml:"uuid"`
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
	Timing               EventTiming          `json:"timing,omitempty" yaml:"timing,omitempty"`
	Title                string               `json:"title" yaml:"title"`
	Type                 string               `json:"type" yaml:"type"`
	UUID                 string               `json:"uuid" yaml:"uuid"`
}

type EventTiming struct {
	AtFrequency     FrequencyCondition   `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	OnDate          OnDateCondition      `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange OnDateRangeCondition `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

type AssessmentPlanTermsAndConditions struct {
	Parts []AssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type RiskLog struct {
	Entries []RiskLogEntry `json:"entries" yaml:"entries"`
}

type Capability_ControlImplementation struct {
	Description             string                                         `json:"description" yaml:"description"`
	ImplementedRequirements []ControlImplementation_ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   []Link                                         `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property                                     `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           []SetParameter                                 `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                                         `json:"source" yaml:"source"`
	UUID                    string                                         `json:"uuid" yaml:"uuid"`
}

type AssessedControls_SelectControlById struct {
	ControlId    string   `json:"control-id" yaml:"control-id"`
	StatementIds []string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
}

type CustomGrouping_Group struct {
	Class          string                 `json:"class,omitempty" yaml:"class,omitempty"`
	Groups         []CustomGrouping_Group `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID             string                 `json:"id,omitempty" yaml:"id,omitempty"`
	InsertControls []InsertControls       `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
	Links          []Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Params         []Parameter            `json:"params,omitempty" yaml:"params,omitempty"`
	Parts          []Part                 `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props          []Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Title          string                 `json:"title" yaml:"title"`
}

type ImplementedRequirement_Statement struct {
	Description      string            `json:"description" yaml:"description"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type RiskResponseReference struct {
	Links        []Link        `json:"links,omitempty" yaml:"links,omitempty"`
	Props        []Property    `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks []RelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponseUuid string        `json:"response-uuid" yaml:"response-uuid"`
}

type SystemComponent_Status struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type PlanOfActionAndMilestones struct {
	BackMatter       BackMatter                                 `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Findings         []Finding                                  `json:"findings,omitempty" yaml:"findings,omitempty"`
	ImportSsp        ImportSsp                                  `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	LocalDefinitions PlanOfActionAndMilestones_LocalDefinitions `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         Metadata                                   `json:"metadata" yaml:"metadata"`
	Observations     []Observation                              `json:"observations,omitempty" yaml:"observations,omitempty"`
	PoamItems        []PoamItem                                 `json:"poam-items" yaml:"poam-items"`
	Risks            []Risk                                     `json:"risks,omitempty" yaml:"risks,omitempty"`
	SystemId         SystemId                                   `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	UUID             string                                     `json:"uuid" yaml:"uuid"`
}

type DefinedComponent struct {
	ControlImplementations []DefinedComponent_ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                                   `json:"description" yaml:"description"`
	Links                  []Link                                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props                  []Property                               `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              []Protocol                               `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose                string                                   `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks                string                                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles       []ResponsibleRole                        `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Title                  string                                   `json:"title" yaml:"title"`
	Type                   string                                   `json:"type" yaml:"type"`
	UUID                   string                                   `json:"uuid" yaml:"uuid"`
}

type ControlImplementation struct {
	Description             string                   `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	SetParameters           []SetParameter           `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type OnDateCondition struct {
	Date string `json:"date" yaml:"date"`
}

type Capability struct {
	ControlImplementations []Capability_ControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                             `json:"description" yaml:"description"`
	IncorporatesComponents []IncorporatesComponent            `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	Links                  []Link                             `json:"links,omitempty" yaml:"links,omitempty"`
	Name                   string                             `json:"name" yaml:"name"`
	Props                  []Property                         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                string                             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                             `json:"uuid" yaml:"uuid"`
}

type FrequencyCondition struct {
	Period int    `json:"period" yaml:"period"`
	Unit   string `json:"unit" yaml:"unit"`
}

type PlanOfActionAndMilestones_LocalDefinitions struct {
	AssessmentAssets AssessmentAssets  `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	Components       []SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems   []InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ImplementedRequirement struct {
	ByComponents     []ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	ControlId        string            `json:"control-id" yaml:"control-id"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    []SetParameter    `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       []Statement       `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type FindingTarget struct {
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	ImplementationStatus ImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Links                []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Props                []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Status               ObjectiveStatus      `json:"status" yaml:"status"`
	TargetId             string               `json:"target-id" yaml:"target-id"`
	Title                string               `json:"title,omitempty" yaml:"title,omitempty"`
	Type                 string               `json:"type" yaml:"type"`
}

type Statement struct {
	ByComponents     []ByComponent     `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Links            []Link            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string            `json:"statement-id" yaml:"statement-id"`
	UUID             string            `json:"uuid" yaml:"uuid"`
}

type SelectControlById struct {
	Matching          []Matching `json:"matching,omitempty" yaml:"matching,omitempty"`
	WithChildControls string     `json:"with-child-controls,omitempty" yaml:"with-child-controls,omitempty"`
	WithIds           []string   `json:"with-ids,omitempty" yaml:"with-ids,omitempty"`
}

type PoamItem_Origin struct {
	Actors []OriginActor `json:"actors" yaml:"actors"`
}

type AssessedControls struct {
	Description     string                               `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls []AssessedControls_SelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      IncludeAll                           `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls []AssessedControls_SelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           []Link                               `json:"links,omitempty" yaml:"links,omitempty"`
	Props           []Property                           `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string                               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type ControlImplementation_ImplementedRequirement struct {
	ControlId        string                             `json:"control-id" yaml:"control-id"`
	Description      string                             `json:"description" yaml:"description"`
	Links            []Link                             `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property                         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole                  `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    []SetParameter                     `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       []ImplementedRequirement_Statement `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                             `json:"uuid" yaml:"uuid"`
}

type Merge struct {
	AsIs    bool                `json:"as-is,omitempty" yaml:"as-is,omitempty"`
	Combine CombinationRule     `json:"combine,omitempty" yaml:"combine,omitempty"`
	Custom  CustomGrouping      `json:"custom,omitempty" yaml:"custom,omitempty"`
	Flat    FlatWithoutGrouping `json:"flat,omitempty" yaml:"flat,omitempty"`
}

type CombinationRule struct {
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

type CustomGrouping struct {
	Groups         []CustomGrouping_Group `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls []InsertControls       `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
}

type DefinedComponent_ControlImplementation struct {
	Description             string                                         `json:"description" yaml:"description"`
	ImplementedRequirements []ControlImplementation_ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   []Link                                         `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   []Property                                     `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           []SetParameter                                 `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                                         `json:"source" yaml:"source"`
	UUID                    string                                         `json:"uuid" yaml:"uuid"`
}

type LocalDefinitions struct {
	Activities           []Activity        `json:"activities,omitempty" yaml:"activities,omitempty"`
	Components           []SystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems       []InventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	ObjectivesAndMethods []LocalObjective  `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Remarks              string            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                []SystemUser      `json:"users,omitempty" yaml:"users,omitempty"`
}

type AssessmentPlan struct {
	AssessmentAssets   AssessmentAssets                 `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	AssessmentSubjects []AssessmentSubject              `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
	BackMatter         BackMatter                       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp          ImportSsp                        `json:"import-ssp" yaml:"import-ssp"`
	LocalDefinitions   LocalDefinitions                 `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata           Metadata                         `json:"metadata" yaml:"metadata"`
	ReviewedControls   ReviewedControls                 `json:"reviewed-controls" yaml:"reviewed-controls"`
	Tasks              []Task                           `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	TermsAndConditions AssessmentPlanTermsAndConditions `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	UUID               string                           `json:"uuid" yaml:"uuid"`
}

type FlatWithoutGrouping struct {
}

type SystemComponent struct {
	Description      string                 `json:"description" yaml:"description"`
	Links            []Link                 `json:"links,omitempty" yaml:"links,omitempty"`
	Props            []Property             `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols        []Protocol             `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose          string                 `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks          string                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles []ResponsibleRole      `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Status           SystemComponent_Status `json:"status" yaml:"status"`
	Title            string                 `json:"title" yaml:"title"`
	Type             string                 `json:"type" yaml:"type"`
	UUID             string                 `json:"uuid" yaml:"uuid"`
}

type PoamItem struct {
	Description         string               `json:"description" yaml:"description"`
	Links               []Link               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             []PoamItem_Origin    `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               []Property           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedFindings     []RelatedFinding     `json:"related-findings,omitempty" yaml:"related-findings,omitempty"`
	RelatedObservations []RelatedObservation `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks        []AssociatedRisk     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks             string               `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title               string               `json:"title" yaml:"title"`
	UUID                string               `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type OnDateRangeCondition struct {
	End   string `json:"end" yaml:"end"`
	Start string `json:"start" yaml:"start"`
}

type ObjectiveStatus struct {
	Reason  string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}