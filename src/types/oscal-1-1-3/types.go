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
package oscalTypes_1_1_3

import (
	"time"
)

type OscalModels = OscalCompleteSchema
type OscalCompleteSchema struct {
	AssessmentPlan            *OscalCompleteOscalApAssessmentPlan                       `json:"assessment-plan,omitempty" yaml:"assessment-plan,omitempty"`
	AssessmentResults         *OscalCompleteOscalArAssessmentResults                    `json:"assessment-results,omitempty" yaml:"assessment-results,omitempty"`
	Catalog                   *OscalCompleteOscalCatalogCatalog                         `json:"catalog,omitempty" yaml:"catalog,omitempty"`
	ComponentDefinition       *OscalCompleteOscalComponentDefinitionComponentDefinition `json:"component-definition,omitempty" yaml:"component-definition,omitempty"`
	PlanOfActionAndMilestones *OscalCompleteOscalPoamPlanOfActionAndMilestones          `json:"plan-of-action-and-milestones,omitempty" yaml:"plan-of-action-and-milestones,omitempty"`
	Profile                   *OscalCompleteOscalProfileProfile                         `json:"profile,omitempty" yaml:"profile,omitempty"`
	SystemSecurityPlan        *OscalCompleteOscalSspSystemSecurityPlan                  `json:"system-security-plan,omitempty" yaml:"system-security-plan,omitempty"`
}

type OscalCompleteOscalApAssessmentPlan struct {
	AssessmentAssets   *OscalCompleteOscalAssessmentCommonAssessmentAssets    `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	AssessmentSubjects *[]OscalCompleteOscalAssessmentCommonAssessmentSubject `json:"assessment-subjects,omitempty" yaml:"assessment-subjects,omitempty"`
	BackMatter         *OscalCompleteOscalMetadataBackMatter                  `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportSsp          OscalCompleteOscalAssessmentCommonImportSsp            `json:"import-ssp" yaml:"import-ssp"`
	LocalDefinitions   *LocalDefinitions                                      `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata           OscalCompleteOscalMetadataMetadata                     `json:"metadata" yaml:"metadata"`
	ReviewedControls   OscalCompleteOscalAssessmentCommonReviewedControls     `json:"reviewed-controls" yaml:"reviewed-controls"`
	Tasks              *[]OscalCompleteOscalAssessmentCommonTask              `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	TermsAndConditions *AssessmentPlanTermsAndConditions                      `json:"terms-and-conditions,omitempty" yaml:"terms-and-conditions,omitempty"`
	UUID               string                                                 `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalArAssessmentResults struct {
	BackMatter       *OscalCompleteOscalMetadataBackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ImportAp         OscalCompleteOscalArImportAp          `json:"import-ap" yaml:"import-ap"`
	LocalDefinitions *LocalDefinitions                     `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         OscalCompleteOscalMetadataMetadata    `json:"metadata" yaml:"metadata"`
	Results          []OscalCompleteOscalArResult          `json:"results" yaml:"results"`
	UUID             string                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalCatalogCatalog struct {
	BackMatter *OscalCompleteOscalMetadataBackMatter       `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Controls   *[]OscalCompleteOscalCatalogControl         `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups     *[]OscalCompleteOscalCatalogGroup           `json:"groups,omitempty" yaml:"groups,omitempty"`
	Metadata   OscalCompleteOscalMetadataMetadata          `json:"metadata" yaml:"metadata"`
	Params     *[]OscalCompleteOscalControlCommonParameter `json:"params,omitempty" yaml:"params,omitempty"`
	UUID       string                                      `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalComponentDefinitionComponentDefinition struct {
	BackMatter                 *OscalCompleteOscalMetadataBackMatter                             `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Capabilities               *[]OscalCompleteOscalComponentDefinitionCapability                `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Components                 *[]OscalCompleteOscalComponentDefinitionDefinedComponent          `json:"components,omitempty" yaml:"components,omitempty"`
	ImportComponentDefinitions *[]OscalCompleteOscalComponentDefinitionImportComponentDefinition `json:"import-component-definitions,omitempty" yaml:"import-component-definitions,omitempty"`
	Metadata                   OscalCompleteOscalMetadataMetadata                                `json:"metadata" yaml:"metadata"`
	UUID                       string                                                            `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalPoamPlanOfActionAndMilestones struct {
	BackMatter       *OscalCompleteOscalMetadataBackMatter            `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Findings         *[]OscalCompleteOscalAssessmentCommonFinding     `json:"findings,omitempty" yaml:"findings,omitempty"`
	ImportSsp        *OscalCompleteOscalAssessmentCommonImportSsp     `json:"import-ssp,omitempty" yaml:"import-ssp,omitempty"`
	LocalDefinitions *OscalCompleteOscalPoamLocalDefinitions          `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Metadata         OscalCompleteOscalMetadataMetadata               `json:"metadata" yaml:"metadata"`
	Observations     *[]OscalCompleteOscalAssessmentCommonObservation `json:"observations,omitempty" yaml:"observations,omitempty"`
	PoamItems        []OscalCompleteOscalPoamPoamItem                 `json:"poam-items" yaml:"poam-items"`
	Risks            *[]OscalCompleteOscalAssessmentCommonRisk        `json:"risks,omitempty" yaml:"risks,omitempty"`
	SystemId         *OscalCompleteOscalImplementationCommonSystemId  `json:"system-id,omitempty" yaml:"system-id,omitempty"`
	UUID             string                                           `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalProfileProfile struct {
	BackMatter *OscalCompleteOscalMetadataBackMatter `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	Imports    []OscalCompleteOscalProfileImport     `json:"imports" yaml:"imports"`
	Merge      *OscalCompleteOscalProfileMerge       `json:"merge,omitempty" yaml:"merge,omitempty"`
	Metadata   OscalCompleteOscalMetadataMetadata    `json:"metadata" yaml:"metadata"`
	Modify     *OscalCompleteOscalProfileModify      `json:"modify,omitempty" yaml:"modify,omitempty"`
	UUID       string                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalSspSystemSecurityPlan struct {
	BackMatter            *OscalCompleteOscalMetadataBackMatter      `json:"back-matter,omitempty" yaml:"back-matter,omitempty"`
	ControlImplementation OscalCompleteOscalSspControlImplementation `json:"control-implementation" yaml:"control-implementation"`
	ImportProfile         OscalCompleteOscalSspImportProfile         `json:"import-profile" yaml:"import-profile"`
	Metadata              OscalCompleteOscalMetadataMetadata         `json:"metadata" yaml:"metadata"`
	SystemCharacteristics OscalCompleteOscalSspSystemCharacteristics `json:"system-characteristics" yaml:"system-characteristics"`
	SystemImplementation  OscalCompleteOscalSspSystemImplementation  `json:"system-implementation" yaml:"system-implementation"`
	UUID                  string                                     `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalAssessmentCommonAssessmentAssets struct {
	AssessmentPlatforms []AssessmentPlatform                                     `json:"assessment-platforms" yaml:"assessment-platforms"`
	Components          *[]OscalCompleteOscalImplementationCommonSystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
}

type OscalCompleteOscalAssessmentCommonAssessmentSubject struct {
	Description     string                                                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeSubjects *[]OscalCompleteOscalAssessmentCommonSelectSubjectById `json:"exclude-subjects,omitempty" yaml:"exclude-subjects,omitempty"`
	IncludeAll      *OscalCompleteOscalControlCommonIncludeAll             `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeSubjects *[]OscalCompleteOscalAssessmentCommonSelectSubjectById `json:"include-subjects,omitempty" yaml:"include-subjects,omitempty"`
	Links           *[]OscalCompleteOscalMetadataLink                      `json:"links,omitempty" yaml:"links,omitempty"`
	Props           *[]OscalCompleteOscalMetadataProperty                  `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string                                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Type            string                                                 `json:"type" yaml:"type"`
}

type OscalCompleteOscalMetadataBackMatter struct {
	Resources *[]Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type OscalCompleteOscalAssessmentCommonImportSsp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type LocalDefinitions struct {
	Activities           *[]OscalCompleteOscalAssessmentCommonActivity            `json:"activities,omitempty" yaml:"activities,omitempty"`
	Components           *[]OscalCompleteOscalImplementationCommonSystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems       *[]OscalCompleteOscalImplementationCommonInventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	ObjectivesAndMethods *[]OscalCompleteOscalAssessmentCommonLocalObjective      `json:"objectives-and-methods,omitempty" yaml:"objectives-and-methods,omitempty"`
	Remarks              string                                                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                *[]OscalCompleteOscalImplementationCommonSystemUser      `json:"users,omitempty" yaml:"users,omitempty"`
}

type OscalCompleteOscalMetadataMetadata struct {
	Actions            *[]OscalCompleteOscalMetadataAction           `json:"actions,omitempty" yaml:"actions,omitempty"`
	DocumentIds        *[]OscalCompleteOscalMetadataDocumentId       `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	LastModified       time.Time                                     `json:"last-modified" yaml:"last-modified"`
	Links              *[]OscalCompleteOscalMetadataLink             `json:"links,omitempty" yaml:"links,omitempty"`
	Locations          *[]Location                                   `json:"locations,omitempty" yaml:"locations,omitempty"`
	OscalVersion       string                                        `json:"oscal-version" yaml:"oscal-version"`
	Parties            *[]Party                                      `json:"parties,omitempty" yaml:"parties,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty         `json:"props,omitempty" yaml:"props,omitempty"`
	Published          *time.Time                                    `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks            string                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]OscalCompleteOscalMetadataResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Revisions          *[]RevisionHistoryEntry                       `json:"revisions,omitempty" yaml:"revisions,omitempty"`
	Roles              *[]Role                                       `json:"roles,omitempty" yaml:"roles,omitempty"`
	Title              string                                        `json:"title" yaml:"title"`
	Version            string                                        `json:"version" yaml:"version"`
}

type OscalCompleteOscalAssessmentCommonReviewedControls struct {
	ControlObjectiveSelections *[]ReferencedControlObjectives        `json:"control-objective-selections,omitempty" yaml:"control-objective-selections,omitempty"`
	ControlSelections          []AssessedControls                    `json:"control-selections" yaml:"control-selections"`
	Description                string                                `json:"description,omitempty" yaml:"description,omitempty"`
	Links                      *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                      *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                    string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalAssessmentCommonTask struct {
	AssociatedActivities *[]AssociatedActivity                                  `json:"associated-activities,omitempty" yaml:"associated-activities,omitempty"`
	Dependencies         *[]TaskDependency                                      `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Description          string                                                 `json:"description,omitempty" yaml:"description,omitempty"`
	Links                *[]OscalCompleteOscalMetadataLink                      `json:"links,omitempty" yaml:"links,omitempty"`
	Props                *[]OscalCompleteOscalMetadataProperty                  `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     *[]OscalCompleteOscalMetadataResponsibleRole           `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects             *[]OscalCompleteOscalAssessmentCommonAssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Tasks                *[]OscalCompleteOscalAssessmentCommonTask              `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Timing               *EventTiming                                           `json:"timing,omitempty" yaml:"timing,omitempty"`
	Title                string                                                 `json:"title" yaml:"title"`
	Type                 string                                                 `json:"type" yaml:"type"`
	UUID                 string                                                 `json:"uuid" yaml:"uuid"`
}

type AssessmentPlanTermsAndConditions struct {
	Parts *[]OscalCompleteOscalAssessmentCommonAssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
}

type OscalCompleteOscalArImportAp struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalArResult struct {
	AssessmentLog    *AssessmentLog                                     `json:"assessment-log,omitempty" yaml:"assessment-log,omitempty"`
	Attestations     *[]AttestationStatements                           `json:"attestations,omitempty" yaml:"attestations,omitempty"`
	Description      string                                             `json:"description" yaml:"description"`
	End              *time.Time                                         `json:"end,omitempty" yaml:"end,omitempty"`
	Findings         *[]OscalCompleteOscalAssessmentCommonFinding       `json:"findings,omitempty" yaml:"findings,omitempty"`
	Links            *[]OscalCompleteOscalMetadataLink                  `json:"links,omitempty" yaml:"links,omitempty"`
	LocalDefinitions *LocalDefinitions                                  `json:"local-definitions,omitempty" yaml:"local-definitions,omitempty"`
	Observations     *[]OscalCompleteOscalAssessmentCommonObservation   `json:"observations,omitempty" yaml:"observations,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty              `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                             `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ReviewedControls OscalCompleteOscalAssessmentCommonReviewedControls `json:"reviewed-controls" yaml:"reviewed-controls"`
	Risks            *[]OscalCompleteOscalAssessmentCommonRisk          `json:"risks,omitempty" yaml:"risks,omitempty"`
	Start            time.Time                                          `json:"start" yaml:"start"`
	Title            string                                             `json:"title" yaml:"title"`
	UUID             string                                             `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalCatalogControl struct {
	Class    string                                      `json:"class,omitempty" yaml:"class,omitempty"`
	Controls *[]OscalCompleteOscalCatalogControl         `json:"controls,omitempty" yaml:"controls,omitempty"`
	ID       string                                      `json:"id" yaml:"id"`
	Links    *[]OscalCompleteOscalMetadataLink           `json:"links,omitempty" yaml:"links,omitempty"`
	Params   *[]OscalCompleteOscalControlCommonParameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    *[]OscalCompleteOscalControlCommonPart      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props    *[]OscalCompleteOscalMetadataProperty       `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string                                      `json:"title" yaml:"title"`
}

type OscalCompleteOscalCatalogGroup struct {
	Class    string                                      `json:"class,omitempty" yaml:"class,omitempty"`
	Controls *[]OscalCompleteOscalCatalogControl         `json:"controls,omitempty" yaml:"controls,omitempty"`
	Groups   *[]OscalCompleteOscalCatalogGroup           `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID       string                                      `json:"id,omitempty" yaml:"id,omitempty"`
	Links    *[]OscalCompleteOscalMetadataLink           `json:"links,omitempty" yaml:"links,omitempty"`
	Params   *[]OscalCompleteOscalControlCommonParameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    *[]OscalCompleteOscalControlCommonPart      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props    *[]OscalCompleteOscalMetadataProperty       `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string                                      `json:"title" yaml:"title"`
}

type OscalCompleteOscalControlCommonParameter struct {
	Class       string                                                `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints *[]OscalCompleteOscalControlCommonParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   string                                                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  *[]OscalCompleteOscalControlCommonParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	ID          string                                                `json:"id" yaml:"id"`
	Label       string                                                `json:"label,omitempty" yaml:"label,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Select      *OscalCompleteOscalControlCommonParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       string                                                `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      *[]string                                             `json:"values,omitempty" yaml:"values,omitempty"`
}

type OscalCompleteOscalComponentDefinitionCapability struct {
	ControlImplementations *[]OscalCompleteOscalComponentDefinitionControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                                                        `json:"description" yaml:"description"`
	IncorporatesComponents *[]OscalCompleteOscalComponentDefinitionIncorporatesComponent `json:"incorporates-components,omitempty" yaml:"incorporates-components,omitempty"`
	Links                  *[]OscalCompleteOscalMetadataLink                             `json:"links,omitempty" yaml:"links,omitempty"`
	Name                   string                                                        `json:"name" yaml:"name"`
	Props                  *[]OscalCompleteOscalMetadataProperty                         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                string                                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID                   string                                                        `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalComponentDefinitionDefinedComponent struct {
	ControlImplementations *[]OscalCompleteOscalComponentDefinitionControlImplementation `json:"control-implementations,omitempty" yaml:"control-implementations,omitempty"`
	Description            string                                                        `json:"description" yaml:"description"`
	Links                  *[]OscalCompleteOscalMetadataLink                             `json:"links,omitempty" yaml:"links,omitempty"`
	Props                  *[]OscalCompleteOscalMetadataProperty                         `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols              *[]OscalCompleteOscalImplementationCommonProtocol             `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose                string                                                        `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks                string                                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles       *[]OscalCompleteOscalMetadataResponsibleRole                  `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Title                  string                                                        `json:"title" yaml:"title"`
	Type                   string                                                        `json:"type" yaml:"type"`
	UUID                   string                                                        `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalComponentDefinitionImportComponentDefinition struct {
	Href string `json:"href" yaml:"href"`
}

type OscalCompleteOscalAssessmentCommonFinding struct {
	Description                 string                                          `json:"description" yaml:"description"`
	ImplementationStatementUuid string                                          `json:"implementation-statement-uuid,omitempty" yaml:"implementation-statement-uuid,omitempty"`
	Links                       *[]OscalCompleteOscalMetadataLink               `json:"links,omitempty" yaml:"links,omitempty"`
	Origins                     *[]OscalCompleteOscalAssessmentCommonOrigin     `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props                       *[]OscalCompleteOscalMetadataProperty           `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations         *[]RelatedObservation                           `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks                *[]AssociatedRisk                               `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks                     string                                          `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Target                      OscalCompleteOscalAssessmentCommonFindingTarget `json:"target" yaml:"target"`
	Title                       string                                          `json:"title" yaml:"title"`
	UUID                        string                                          `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalPoamLocalDefinitions struct {
	AssessmentAssets *OscalCompleteOscalAssessmentCommonAssessmentAssets      `json:"assessment-assets,omitempty" yaml:"assessment-assets,omitempty"`
	Components       *[]OscalCompleteOscalImplementationCommonSystemComponent `json:"components,omitempty" yaml:"components,omitempty"`
	InventoryItems   *[]OscalCompleteOscalImplementationCommonInventoryItem   `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	Remarks          string                                                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalAssessmentCommonObservation struct {
	Collected        time.Time                                             `json:"collected" yaml:"collected"`
	Description      string                                                `json:"description" yaml:"description"`
	Expires          *time.Time                                            `json:"expires,omitempty" yaml:"expires,omitempty"`
	Links            *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Methods          []string                                              `json:"methods" yaml:"methods"`
	Origins          *[]OscalCompleteOscalAssessmentCommonOrigin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	RelevantEvidence *[]RelevantEvidence                                   `json:"relevant-evidence,omitempty" yaml:"relevant-evidence,omitempty"`
	Remarks          string                                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects         *[]OscalCompleteOscalAssessmentCommonSubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title            string                                                `json:"title,omitempty" yaml:"title,omitempty"`
	Types            *[]string                                             `json:"types,omitempty" yaml:"types,omitempty"`
	UUID             string                                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalPoamPoamItem struct {
	Description         string                                `json:"description" yaml:"description"`
	Links               *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Origins             *[]Origin                             `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedFindings     *[]RelatedFinding                     `json:"related-findings,omitempty" yaml:"related-findings,omitempty"`
	RelatedObservations *[]RelatedObservation                 `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	RelatedRisks        *[]AssociatedRisk                     `json:"related-risks,omitempty" yaml:"related-risks,omitempty"`
	Remarks             string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title               string                                `json:"title" yaml:"title"`
	UUID                string                                `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type OscalCompleteOscalAssessmentCommonRisk struct {
	Characterizations   *[]OscalCompleteOscalAssessmentCommonCharacterization `json:"characterizations,omitempty" yaml:"characterizations,omitempty"`
	Deadline            *time.Time                                            `json:"deadline,omitempty" yaml:"deadline,omitempty"`
	Description         string                                                `json:"description" yaml:"description"`
	Links               *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	MitigatingFactors   *[]MitigatingFactor                                   `json:"mitigating-factors,omitempty" yaml:"mitigating-factors,omitempty"`
	Origins             *[]OscalCompleteOscalAssessmentCommonOrigin           `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props               *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedObservations *[]RelatedObservation                                 `json:"related-observations,omitempty" yaml:"related-observations,omitempty"`
	Remediations        *[]OscalCompleteOscalAssessmentCommonResponse         `json:"remediations,omitempty" yaml:"remediations,omitempty"`
	RiskLog             *RiskLog                                              `json:"risk-log,omitempty" yaml:"risk-log,omitempty"`
	Statement           string                                                `json:"statement" yaml:"statement"`
	Status              string                                                `json:"status" yaml:"status"`
	ThreatIds           *[]OscalCompleteOscalAssessmentCommonThreatId         `json:"threat-ids,omitempty" yaml:"threat-ids,omitempty"`
	Title               string                                                `json:"title" yaml:"title"`
	UUID                string                                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalImplementationCommonSystemId struct {
	ID             string `json:"id" yaml:"id"`
	IdentifierType string `json:"identifier-type,omitempty" yaml:"identifier-type,omitempty"`
}

type OscalCompleteOscalProfileImport struct {
	ExcludeControls *[]OscalCompleteOscalProfileSelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	Href            string                                        `json:"href" yaml:"href"`
	IncludeAll      *OscalCompleteOscalControlCommonIncludeAll    `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]OscalCompleteOscalProfileSelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
}

type OscalCompleteOscalProfileMerge struct {
	AsIs    bool                 `json:"as-is,omitempty" yaml:"as-is,omitempty"`
	Combine *CombinationRule     `json:"combine,omitempty" yaml:"combine,omitempty"`
	Custom  *CustomGrouping      `json:"custom,omitempty" yaml:"custom,omitempty"`
	Flat    *FlatWithoutGrouping `json:"flat,omitempty" yaml:"flat,omitempty"`
}

type OscalCompleteOscalProfileModify struct {
	Alters        *[]Alteration       `json:"alters,omitempty" yaml:"alters,omitempty"`
	SetParameters *[]ParameterSetting `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type OscalCompleteOscalSspControlImplementation struct {
	Description             string                                                `json:"description" yaml:"description"`
	ImplementedRequirements []OscalCompleteOscalSspImplementedRequirement         `json:"implemented-requirements" yaml:"implemented-requirements"`
	SetParameters           *[]OscalCompleteOscalImplementationCommonSetParameter `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
}

type OscalCompleteOscalSspImportProfile struct {
	Href    string `json:"href" yaml:"href"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalSspSystemCharacteristics struct {
	AuthorizationBoundary    OscalCompleteOscalSspAuthorizationBoundary       `json:"authorization-boundary" yaml:"authorization-boundary"`
	DataFlow                 *OscalCompleteOscalSspDataFlow                   `json:"data-flow,omitempty" yaml:"data-flow,omitempty"`
	DateAuthorized           string                                           `json:"date-authorized,omitempty" yaml:"date-authorized,omitempty"`
	Description              string                                           `json:"description" yaml:"description"`
	Links                    *[]OscalCompleteOscalMetadataLink                `json:"links,omitempty" yaml:"links,omitempty"`
	NetworkArchitecture      *OscalCompleteOscalSspNetworkArchitecture        `json:"network-architecture,omitempty" yaml:"network-architecture,omitempty"`
	Props                    *[]OscalCompleteOscalMetadataProperty            `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                  string                                           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties       *[]OscalCompleteOscalMetadataResponsibleParty    `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	SecurityImpactLevel      *OscalCompleteOscalSspSecurityImpactLevel        `json:"security-impact-level,omitempty" yaml:"security-impact-level,omitempty"`
	SecuritySensitivityLevel string                                           `json:"security-sensitivity-level,omitempty" yaml:"security-sensitivity-level,omitempty"`
	Status                   OscalCompleteOscalSspStatus                      `json:"status" yaml:"status"`
	SystemIds                []OscalCompleteOscalImplementationCommonSystemId `json:"system-ids" yaml:"system-ids"`
	SystemInformation        OscalCompleteOscalSspSystemInformation           `json:"system-information" yaml:"system-information"`
	SystemName               string                                           `json:"system-name" yaml:"system-name"`
	SystemNameShort          string                                           `json:"system-name-short,omitempty" yaml:"system-name-short,omitempty"`
}

type OscalCompleteOscalSspSystemImplementation struct {
	Components              []OscalCompleteOscalImplementationCommonSystemComponent `json:"components" yaml:"components"`
	InventoryItems          *[]OscalCompleteOscalImplementationCommonInventoryItem  `json:"inventory-items,omitempty" yaml:"inventory-items,omitempty"`
	LeveragedAuthorizations *[]LeveragedAuthorization                               `json:"leveraged-authorizations,omitempty" yaml:"leveraged-authorizations,omitempty"`
	Links                   *[]OscalCompleteOscalMetadataLink                       `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]OscalCompleteOscalMetadataProperty                   `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks                 string                                                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Users                   []OscalCompleteOscalImplementationCommonSystemUser      `json:"users" yaml:"users"`
}

type AssessmentPlatform struct {
	Links          *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props          *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string                                `json:"title,omitempty" yaml:"title,omitempty"`
	UsesComponents *[]UsesComponent                      `json:"uses-components,omitempty" yaml:"uses-components,omitempty"`
	UUID           string                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalImplementationCommonSystemComponent struct {
	Description      string                                            `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink                 `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty             `json:"props,omitempty" yaml:"props,omitempty"`
	Protocols        *[]OscalCompleteOscalImplementationCommonProtocol `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Purpose          string                                            `json:"purpose,omitempty" yaml:"purpose,omitempty"`
	Remarks          string                                            `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole      `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Status           Status                                            `json:"status" yaml:"status"`
	Title            string                                            `json:"title" yaml:"title"`
	Type             string                                            `json:"type" yaml:"type"`
	UUID             string                                            `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalAssessmentCommonSelectSubjectById struct {
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string                                `json:"subject-uuid" yaml:"subject-uuid"`
	Type        string                                `json:"type" yaml:"type"`
}

type OscalCompleteOscalControlCommonIncludeAll = map[string]interface{}

type OscalCompleteOscalMetadataLink struct {
	Href             string `json:"href" yaml:"href"`
	MediaType        string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Rel              string `json:"rel,omitempty" yaml:"rel,omitempty"`
	ResourceFragment string `json:"resource-fragment,omitempty" yaml:"resource-fragment,omitempty"`
	Text             string `json:"text,omitempty" yaml:"text,omitempty"`
}

type OscalCompleteOscalMetadataProperty struct {
	Class   string `json:"class,omitempty" yaml:"class,omitempty"`
	Group   string `json:"group,omitempty" yaml:"group,omitempty"`
	Name    string `json:"name" yaml:"name"`
	Ns      string `json:"ns,omitempty" yaml:"ns,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID    string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Value   string `json:"value" yaml:"value"`
}

type Resource struct {
	Base64      *Base64                                 `json:"base64,omitempty" yaml:"base64,omitempty"`
	Citation    *Citation                               `json:"citation,omitempty" yaml:"citation,omitempty"`
	Description string                                  `json:"description,omitempty" yaml:"description,omitempty"`
	DocumentIds *[]OscalCompleteOscalMetadataDocumentId `json:"document-ids,omitempty" yaml:"document-ids,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty   `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                  `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Rlinks      *[]ResourceLink                         `json:"rlinks,omitempty" yaml:"rlinks,omitempty"`
	Title       string                                  `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string                                  `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalAssessmentCommonActivity struct {
	Description      string                                              `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty               `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedControls  *OscalCompleteOscalAssessmentCommonReviewedControls `json:"related-controls,omitempty" yaml:"related-controls,omitempty"`
	Remarks          string                                              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole        `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Steps            *[]Step                                             `json:"steps,omitempty" yaml:"steps,omitempty"`
	Title            string                                              `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                                              `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalImplementationCommonInventoryItem struct {
	Description           string                                        `json:"description" yaml:"description"`
	ImplementedComponents *[]ImplementedComponent                       `json:"implemented-components,omitempty" yaml:"implemented-components,omitempty"`
	Links                 *[]OscalCompleteOscalMetadataLink             `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 *[]OscalCompleteOscalMetadataProperty         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               string                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties    *[]OscalCompleteOscalMetadataResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	UUID                  string                                        `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalAssessmentCommonLocalObjective struct {
	ControlId   string                                `json:"control-id" yaml:"control-id"`
	Description string                                `json:"description,omitempty" yaml:"description,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Parts       []OscalCompleteOscalControlCommonPart `json:"parts" yaml:"parts"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalImplementationCommonSystemUser struct {
	AuthorizedPrivileges *[]OscalCompleteOscalImplementationCommonAuthorizedPrivilege `json:"authorized-privileges,omitempty" yaml:"authorized-privileges,omitempty"`
	Description          string                                                       `json:"description,omitempty" yaml:"description,omitempty"`
	Links                *[]OscalCompleteOscalMetadataLink                            `json:"links,omitempty" yaml:"links,omitempty"`
	Props                *[]OscalCompleteOscalMetadataProperty                        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleIds              *[]string                                                    `json:"role-ids,omitempty" yaml:"role-ids,omitempty"`
	ShortName            string                                                       `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title                string                                                       `json:"title,omitempty" yaml:"title,omitempty"`
	UUID                 string                                                       `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalMetadataAction struct {
	Date               *time.Time                                    `json:"date,omitempty" yaml:"date,omitempty"`
	Links              *[]OscalCompleteOscalMetadataLink             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]OscalCompleteOscalMetadataResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	System             string                                        `json:"system" yaml:"system"`
	Type               string                                        `json:"type" yaml:"type"`
	UUID               string                                        `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalMetadataDocumentId struct {
	Identifier string `json:"identifier" yaml:"identifier"`
	Scheme     string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type Location struct {
	Address          *OscalCompleteOscalMetadataAddress           `json:"address,omitempty" yaml:"address,omitempty"`
	EmailAddresses   *[]string                                    `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	Links            *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TelephoneNumbers *[]OscalCompleteOscalMetadataTelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Title            string                                       `json:"title,omitempty" yaml:"title,omitempty"`
	Urls             *[]string                                    `json:"urls,omitempty" yaml:"urls,omitempty"`
	UUID             string                                       `json:"uuid" yaml:"uuid"`
}

type Party struct {
	Addresses             *[]OscalCompleteOscalMetadataAddress         `json:"addresses,omitempty" yaml:"addresses,omitempty"`
	EmailAddresses        *[]string                                    `json:"email-addresses,omitempty" yaml:"email-addresses,omitempty"`
	ExternalIds           *[]PartyExternalIdentifier                   `json:"external-ids,omitempty" yaml:"external-ids,omitempty"`
	Links                 *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	LocationUuids         *[]string                                    `json:"location-uuids,omitempty" yaml:"location-uuids,omitempty"`
	MemberOfOrganizations *[]string                                    `json:"member-of-organizations,omitempty" yaml:"member-of-organizations,omitempty"`
	Name                  string                                       `json:"name,omitempty" yaml:"name,omitempty"`
	Props                 *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks               string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName             string                                       `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	TelephoneNumbers      *[]OscalCompleteOscalMetadataTelephoneNumber `json:"telephone-numbers,omitempty" yaml:"telephone-numbers,omitempty"`
	Type                  string                                       `json:"type" yaml:"type"`
	UUID                  string                                       `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalMetadataResponsibleParty struct {
	Links      *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids []string                              `json:"party-uuids" yaml:"party-uuids"`
	Props      *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string                                `json:"role-id" yaml:"role-id"`
}

type RevisionHistoryEntry struct {
	LastModified *time.Time                            `json:"last-modified,omitempty" yaml:"last-modified,omitempty"`
	Links        *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	OscalVersion string                                `json:"oscal-version,omitempty" yaml:"oscal-version,omitempty"`
	Props        *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Published    *time.Time                            `json:"published,omitempty" yaml:"published,omitempty"`
	Remarks      string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title        string                                `json:"title,omitempty" yaml:"title,omitempty"`
	Version      string                                `json:"version" yaml:"version"`
}

type Role struct {
	Description string                                `json:"description,omitempty" yaml:"description,omitempty"`
	ID          string                                `json:"id" yaml:"id"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ShortName   string                                `json:"short-name,omitempty" yaml:"short-name,omitempty"`
	Title       string                                `json:"title" yaml:"title"`
}

type ReferencedControlObjectives struct {
	Description       string                                                   `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeObjectives *[]OscalCompleteOscalAssessmentCommonSelectObjectiveById `json:"exclude-objectives,omitempty" yaml:"exclude-objectives,omitempty"`
	IncludeAll        *OscalCompleteOscalControlCommonIncludeAll               `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeObjectives *[]OscalCompleteOscalAssessmentCommonSelectObjectiveById `json:"include-objectives,omitempty" yaml:"include-objectives,omitempty"`
	Links             *[]OscalCompleteOscalMetadataLink                        `json:"links,omitempty" yaml:"links,omitempty"`
	Props             *[]OscalCompleteOscalMetadataProperty                    `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks           string                                                   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssessedControls struct {
	Description     string                                                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExcludeControls *[]OscalCompleteOscalAssessmentCommonSelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      *OscalCompleteOscalControlCommonIncludeAll             `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]OscalCompleteOscalAssessmentCommonSelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Links           *[]OscalCompleteOscalMetadataLink                      `json:"links,omitempty" yaml:"links,omitempty"`
	Props           *[]OscalCompleteOscalMetadataProperty                  `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks         string                                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type AssociatedActivity struct {
	ActivityUuid     string                                                `json:"activity-uuid" yaml:"activity-uuid"`
	Links            *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole          `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Subjects         []OscalCompleteOscalAssessmentCommonAssessmentSubject `json:"subjects" yaml:"subjects"`
}

type TaskDependency struct {
	Remarks  string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	TaskUuid string `json:"task-uuid" yaml:"task-uuid"`
}

type OscalCompleteOscalMetadataResponsibleRole struct {
	Links      *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuids *[]string                             `json:"party-uuids,omitempty" yaml:"party-uuids,omitempty"`
	Props      *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks    string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RoleId     string                                `json:"role-id" yaml:"role-id"`
}

type EventTiming struct {
	AtFrequency     *FrequencyCondition   `json:"at-frequency,omitempty" yaml:"at-frequency,omitempty"`
	OnDate          *OnDateCondition      `json:"on-date,omitempty" yaml:"on-date,omitempty"`
	WithinDateRange *OnDateRangeCondition `json:"within-date-range,omitempty" yaml:"within-date-range,omitempty"`
}

type OscalCompleteOscalAssessmentCommonAssessmentPart struct {
	Class string                                              `json:"class,omitempty" yaml:"class,omitempty"`
	Links *[]OscalCompleteOscalMetadataLink                   `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string                                              `json:"name" yaml:"name"`
	Ns    string                                              `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts *[]OscalCompleteOscalAssessmentCommonAssessmentPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props *[]OscalCompleteOscalMetadataProperty               `json:"props,omitempty" yaml:"props,omitempty"`
	Prose string                                              `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string                                              `json:"title,omitempty" yaml:"title,omitempty"`
	UUID  string                                              `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries" yaml:"entries"`
}

type AttestationStatements struct {
	Parts              []OscalCompleteOscalAssessmentCommonAssessmentPart `json:"parts" yaml:"parts"`
	ResponsibleParties *[]OscalCompleteOscalMetadataResponsibleParty      `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type OscalCompleteOscalControlCommonPart struct {
	Class string                                 `json:"class,omitempty" yaml:"class,omitempty"`
	ID    string                                 `json:"id,omitempty" yaml:"id,omitempty"`
	Links *[]OscalCompleteOscalMetadataLink      `json:"links,omitempty" yaml:"links,omitempty"`
	Name  string                                 `json:"name" yaml:"name"`
	Ns    string                                 `json:"ns,omitempty" yaml:"ns,omitempty"`
	Parts *[]OscalCompleteOscalControlCommonPart `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props *[]OscalCompleteOscalMetadataProperty  `json:"props,omitempty" yaml:"props,omitempty"`
	Prose string                                 `json:"prose,omitempty" yaml:"prose,omitempty"`
	Title string                                 `json:"title,omitempty" yaml:"title,omitempty"`
}

type OscalCompleteOscalControlCommonParameterConstraint struct {
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Tests       *[]ConstraintTest `json:"tests,omitempty" yaml:"tests,omitempty"`
}

type OscalCompleteOscalControlCommonParameterGuideline struct {
	Prose string `json:"prose" yaml:"prose"`
}

type OscalCompleteOscalControlCommonParameterSelection struct {
	Choice  *[]string `json:"choice,omitempty" yaml:"choice,omitempty"`
	HowMany string    `json:"how-many,omitempty" yaml:"how-many,omitempty"`
}

type OscalCompleteOscalComponentDefinitionControlImplementation struct {
	Description             string                                                        `json:"description" yaml:"description"`
	ImplementedRequirements []OscalCompleteOscalComponentDefinitionImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
	Links                   *[]OscalCompleteOscalMetadataLink                             `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]OscalCompleteOscalMetadataProperty                         `json:"props,omitempty" yaml:"props,omitempty"`
	SetParameters           *[]OscalCompleteOscalImplementationCommonSetParameter         `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Source                  string                                                        `json:"source" yaml:"source"`
	UUID                    string                                                        `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalComponentDefinitionIncorporatesComponent struct {
	ComponentUuid string `json:"component-uuid" yaml:"component-uuid"`
	Description   string `json:"description" yaml:"description"`
}

type OscalCompleteOscalImplementationCommonProtocol struct {
	Name       string                                             `json:"name,omitempty" yaml:"name,omitempty"`
	PortRanges *[]OscalCompleteOscalImplementationCommonPortRange `json:"port-ranges,omitempty" yaml:"port-ranges,omitempty"`
	Title      string                                             `json:"title,omitempty" yaml:"title,omitempty"`
	UUID       string                                             `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type OscalCompleteOscalAssessmentCommonOrigin struct {
	Actors       []OscalCompleteOscalAssessmentCommonOriginActor  `json:"actors" yaml:"actors"`
	RelatedTasks *[]OscalCompleteOscalAssessmentCommonRelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
}

type RelatedObservation struct {
	ObservationUuid string `json:"observation-uuid" yaml:"observation-uuid"`
}

type AssociatedRisk struct {
	RiskUuid string `json:"risk-uuid" yaml:"risk-uuid"`
}

type OscalCompleteOscalAssessmentCommonFindingTarget struct {
	Description          string                                                      `json:"description,omitempty" yaml:"description,omitempty"`
	ImplementationStatus *OscalCompleteOscalImplementationCommonImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Links                *[]OscalCompleteOscalMetadataLink                           `json:"links,omitempty" yaml:"links,omitempty"`
	Props                *[]OscalCompleteOscalMetadataProperty                       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                                                      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Status               ObjectiveStatus                                             `json:"status" yaml:"status"`
	TargetId             string                                                      `json:"target-id" yaml:"target-id"`
	Title                string                                                      `json:"title,omitempty" yaml:"title,omitempty"`
	Type                 string                                                      `json:"type" yaml:"type"`
}

type RelevantEvidence struct {
	Description string                                `json:"description" yaml:"description"`
	Href        string                                `json:"href,omitempty" yaml:"href,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalAssessmentCommonSubjectReference struct {
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	SubjectUuid string                                `json:"subject-uuid" yaml:"subject-uuid"`
	Title       string                                `json:"title,omitempty" yaml:"title,omitempty"`
	Type        string                                `json:"type" yaml:"type"`
}

type Origin struct {
	Actors []OscalCompleteOscalAssessmentCommonOriginActor `json:"actors" yaml:"actors"`
}

type RelatedFinding struct {
	FindingUuid string `json:"finding-uuid" yaml:"finding-uuid"`
}

type OscalCompleteOscalAssessmentCommonCharacterization struct {
	Facets []Facet                                  `json:"facets" yaml:"facets"`
	Links  *[]OscalCompleteOscalMetadataLink        `json:"links,omitempty" yaml:"links,omitempty"`
	Origin OscalCompleteOscalAssessmentCommonOrigin `json:"origin" yaml:"origin"`
	Props  *[]OscalCompleteOscalMetadataProperty    `json:"props,omitempty" yaml:"props,omitempty"`
}

type MitigatingFactor struct {
	Description        string                                                `json:"description" yaml:"description"`
	ImplementationUuid string                                                `json:"implementation-uuid,omitempty" yaml:"implementation-uuid,omitempty"`
	Links              *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Subjects           *[]OscalCompleteOscalAssessmentCommonSubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	UUID               string                                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalAssessmentCommonResponse struct {
	Description    string                                      `json:"description" yaml:"description"`
	Lifecycle      string                                      `json:"lifecycle" yaml:"lifecycle"`
	Links          *[]OscalCompleteOscalMetadataLink           `json:"links,omitempty" yaml:"links,omitempty"`
	Origins        *[]OscalCompleteOscalAssessmentCommonOrigin `json:"origins,omitempty" yaml:"origins,omitempty"`
	Props          *[]OscalCompleteOscalMetadataProperty       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string                                      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	RequiredAssets *[]RequiredAsset                            `json:"required-assets,omitempty" yaml:"required-assets,omitempty"`
	Tasks          *[]OscalCompleteOscalAssessmentCommonTask   `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Title          string                                      `json:"title" yaml:"title"`
	UUID           string                                      `json:"uuid" yaml:"uuid"`
}

type RiskLog struct {
	Entries []RiskLogEntry `json:"entries" yaml:"entries"`
}

type OscalCompleteOscalAssessmentCommonThreatId struct {
	Href   string `json:"href,omitempty" yaml:"href,omitempty"`
	ID     string `json:"id" yaml:"id"`
	System string `json:"system" yaml:"system"`
}

type OscalCompleteOscalProfileSelectControlById struct {
	Matching          *[]OscalCompleteOscalProfileMatching `json:"matching,omitempty" yaml:"matching,omitempty"`
	WithChildControls string                               `json:"with-child-controls,omitempty" yaml:"with-child-controls,omitempty"`
	WithIds           *[]string                            `json:"with-ids,omitempty" yaml:"with-ids,omitempty"`
}

type CombinationRule struct {
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

type CustomGrouping struct {
	Groups         *[]OscalCompleteOscalProfileGroup          `json:"groups,omitempty" yaml:"groups,omitempty"`
	InsertControls *[]OscalCompleteOscalProfileInsertControls `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
}

type FlatWithoutGrouping = map[string]interface{}

type Alteration struct {
	Adds      *[]Addition `json:"adds,omitempty" yaml:"adds,omitempty"`
	ControlId string      `json:"control-id" yaml:"control-id"`
	Removes   *[]Removal  `json:"removes,omitempty" yaml:"removes,omitempty"`
}

type ParameterSetting struct {
	Class       string                                                `json:"class,omitempty" yaml:"class,omitempty"`
	Constraints *[]OscalCompleteOscalControlCommonParameterConstraint `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	DependsOn   string                                                `json:"depends-on,omitempty" yaml:"depends-on,omitempty"`
	Guidelines  *[]OscalCompleteOscalControlCommonParameterGuideline  `json:"guidelines,omitempty" yaml:"guidelines,omitempty"`
	Label       string                                                `json:"label,omitempty" yaml:"label,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	ParamId     string                                                `json:"param-id" yaml:"param-id"`
	Props       *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Select      *OscalCompleteOscalControlCommonParameterSelection    `json:"select,omitempty" yaml:"select,omitempty"`
	Usage       string                                                `json:"usage,omitempty" yaml:"usage,omitempty"`
	Values      *[]string                                             `json:"values,omitempty" yaml:"values,omitempty"`
}

type OscalCompleteOscalSspImplementedRequirement struct {
	ByComponents     *[]OscalCompleteOscalSspByComponent                   `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	ControlId        string                                                `json:"control-id" yaml:"control-id"`
	Links            *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole          `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    *[]OscalCompleteOscalImplementationCommonSetParameter `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       *[]OscalCompleteOscalSspStatement                     `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalImplementationCommonSetParameter struct {
	ParamId string   `json:"param-id" yaml:"param-id"`
	Remarks string   `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Values  []string `json:"values" yaml:"values"`
}

type OscalCompleteOscalSspAuthorizationBoundary struct {
	Description string                                `json:"description" yaml:"description"`
	Diagrams    *[]OscalCompleteOscalSspDiagram       `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalSspDataFlow struct {
	Description string                                `json:"description" yaml:"description"`
	Diagrams    *[]OscalCompleteOscalSspDiagram       `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalSspNetworkArchitecture struct {
	Description string                                `json:"description" yaml:"description"`
	Diagrams    *[]OscalCompleteOscalSspDiagram       `json:"diagrams,omitempty" yaml:"diagrams,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalSspSecurityImpactLevel struct {
	SecurityObjectiveAvailability    string `json:"security-objective-availability" yaml:"security-objective-availability"`
	SecurityObjectiveConfidentiality string `json:"security-objective-confidentiality" yaml:"security-objective-confidentiality"`
	SecurityObjectiveIntegrity       string `json:"security-objective-integrity" yaml:"security-objective-integrity"`
}

type OscalCompleteOscalSspStatus struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type OscalCompleteOscalSspSystemInformation struct {
	InformationTypes []InformationType                     `json:"information-types" yaml:"information-types"`
	Links            *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
}

type LeveragedAuthorization struct {
	DateAuthorized string                                `json:"date-authorized" yaml:"date-authorized"`
	Links          *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	PartyUuid      string                                `json:"party-uuid" yaml:"party-uuid"`
	Props          *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks        string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Title          string                                `json:"title" yaml:"title"`
	UUID           string                                `json:"uuid" yaml:"uuid"`
}

type UsesComponent struct {
	ComponentUuid      string                                        `json:"component-uuid" yaml:"component-uuid"`
	Links              *[]OscalCompleteOscalMetadataLink             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]OscalCompleteOscalMetadataResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type Status struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type Base64 struct {
	Filename  string `json:"filename,omitempty" yaml:"filename,omitempty"`
	MediaType string `json:"media-type,omitempty" yaml:"media-type,omitempty"`
	Value     string `json:"value" yaml:"value"`
}

type Citation struct {
	Links *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Text  string                                `json:"text" yaml:"text"`
}

type ResourceLink struct {
	Hashes    *[]OscalCompleteOscalMetadataHash `json:"hashes,omitempty" yaml:"hashes,omitempty"`
	Href      string                            `json:"href" yaml:"href"`
	MediaType string                            `json:"media-type,omitempty" yaml:"media-type,omitempty"`
}

type Step struct {
	Description      string                                              `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink                   `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty               `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                              `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole        `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	ReviewedControls *OscalCompleteOscalAssessmentCommonReviewedControls `json:"reviewed-controls,omitempty" yaml:"reviewed-controls,omitempty"`
	Title            string                                              `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                                              `json:"uuid" yaml:"uuid"`
}

type ImplementedComponent struct {
	ComponentUuid      string                                        `json:"component-uuid" yaml:"component-uuid"`
	Links              *[]OscalCompleteOscalMetadataLink             `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty         `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]OscalCompleteOscalMetadataResponsibleParty `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
}

type OscalCompleteOscalImplementationCommonAuthorizedPrivilege struct {
	Description        string   `json:"description,omitempty" yaml:"description,omitempty"`
	FunctionsPerformed []string `json:"functions-performed" yaml:"functions-performed"`
	Title              string   `json:"title" yaml:"title"`
}

type OscalCompleteOscalMetadataAddress struct {
	AddrLines  *[]string `json:"addr-lines,omitempty" yaml:"addr-lines,omitempty"`
	City       string    `json:"city,omitempty" yaml:"city,omitempty"`
	Country    string    `json:"country,omitempty" yaml:"country,omitempty"`
	PostalCode string    `json:"postal-code,omitempty" yaml:"postal-code,omitempty"`
	State      string    `json:"state,omitempty" yaml:"state,omitempty"`
	Type       string    `json:"type,omitempty" yaml:"type,omitempty"`
}

type OscalCompleteOscalMetadataTelephoneNumber struct {
	Number string `json:"number" yaml:"number"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}

type PartyExternalIdentifier struct {
	ID     string `json:"id" yaml:"id"`
	Scheme string `json:"scheme" yaml:"scheme"`
}

type OscalCompleteOscalAssessmentCommonSelectObjectiveById struct {
	ObjectiveId string `json:"objective-id" yaml:"objective-id"`
}

type OscalCompleteOscalAssessmentCommonSelectControlById struct {
	ControlId    string    `json:"control-id" yaml:"control-id"`
	StatementIds *[]string `json:"statement-ids,omitempty" yaml:"statement-ids,omitempty"`
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

type AssessmentLogEntry struct {
	Description  string                                           `json:"description,omitempty" yaml:"description,omitempty"`
	End          *time.Time                                       `json:"end,omitempty" yaml:"end,omitempty"`
	Links        *[]OscalCompleteOscalMetadataLink                `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy     *[]OscalCompleteOscalAssessmentCommonLoggedBy    `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props        *[]OscalCompleteOscalMetadataProperty            `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks *[]OscalCompleteOscalAssessmentCommonRelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string                                           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start        time.Time                                        `json:"start" yaml:"start"`
	Title        string                                           `json:"title,omitempty" yaml:"title,omitempty"`
	UUID         string                                           `json:"uuid" yaml:"uuid"`
}

type ConstraintTest struct {
	Expression string `json:"expression" yaml:"expression"`
	Remarks    string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
}

type OscalCompleteOscalComponentDefinitionImplementedRequirement struct {
	ControlId        string                                                `json:"control-id" yaml:"control-id"`
	Description      string                                                `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole          `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	SetParameters    *[]OscalCompleteOscalImplementationCommonSetParameter `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	Statements       *[]OscalCompleteOscalComponentDefinitionStatement     `json:"statements,omitempty" yaml:"statements,omitempty"`
	UUID             string                                                `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalImplementationCommonPortRange struct {
	End       int    `json:"end,omitempty" yaml:"end,omitempty"`
	Start     int    `json:"start,omitempty" yaml:"start,omitempty"`
	Transport string `json:"transport,omitempty" yaml:"transport,omitempty"`
}

type OscalCompleteOscalAssessmentCommonOriginActor struct {
	ActorUuid string                                `json:"actor-uuid" yaml:"actor-uuid"`
	Links     *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props     *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	RoleId    string                                `json:"role-id,omitempty" yaml:"role-id,omitempty"`
	Type      string                                `json:"type" yaml:"type"`
}

type OscalCompleteOscalAssessmentCommonRelatedTask struct {
	IdentifiedSubject  *IdentifiedSubject                                     `json:"identified-subject,omitempty" yaml:"identified-subject,omitempty"`
	Links              *[]OscalCompleteOscalMetadataLink                      `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty                  `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string                                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleParties *[]OscalCompleteOscalMetadataResponsibleParty          `json:"responsible-parties,omitempty" yaml:"responsible-parties,omitempty"`
	Subjects           *[]OscalCompleteOscalAssessmentCommonAssessmentSubject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	TaskUuid           string                                                 `json:"task-uuid" yaml:"task-uuid"`
}

type OscalCompleteOscalImplementationCommonImplementationStatus struct {
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type ObjectiveStatus struct {
	Reason  string `json:"reason,omitempty" yaml:"reason,omitempty"`
	Remarks string `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	State   string `json:"state" yaml:"state"`
}

type Facet struct {
	Links   *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Name    string                                `json:"name" yaml:"name"`
	Props   *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	System  string                                `json:"system" yaml:"system"`
	Value   string                                `json:"value" yaml:"value"`
}

type RequiredAsset struct {
	Description string                                                `json:"description" yaml:"description"`
	Links       *[]OscalCompleteOscalMetadataLink                     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty                 `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Subjects    *[]OscalCompleteOscalAssessmentCommonSubjectReference `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Title       string                                                `json:"title,omitempty" yaml:"title,omitempty"`
	UUID        string                                                `json:"uuid" yaml:"uuid"`
}

type RiskLogEntry struct {
	Description      string                                        `json:"description,omitempty" yaml:"description,omitempty"`
	End              *time.Time                                    `json:"end,omitempty" yaml:"end,omitempty"`
	Links            *[]OscalCompleteOscalMetadataLink             `json:"links,omitempty" yaml:"links,omitempty"`
	LoggedBy         *[]OscalCompleteOscalAssessmentCommonLoggedBy `json:"logged-by,omitempty" yaml:"logged-by,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty         `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedResponses *[]RiskResponseReference                      `json:"related-responses,omitempty" yaml:"related-responses,omitempty"`
	Remarks          string                                        `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Start            time.Time                                     `json:"start" yaml:"start"`
	StatusChange     string                                        `json:"status-change,omitempty" yaml:"status-change,omitempty"`
	Title            string                                        `json:"title,omitempty" yaml:"title,omitempty"`
	UUID             string                                        `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalProfileMatching struct {
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}

type OscalCompleteOscalProfileGroup struct {
	Class          string                                      `json:"class,omitempty" yaml:"class,omitempty"`
	Groups         *[]OscalCompleteOscalProfileGroup           `json:"groups,omitempty" yaml:"groups,omitempty"`
	ID             string                                      `json:"id,omitempty" yaml:"id,omitempty"`
	InsertControls *[]OscalCompleteOscalProfileInsertControls  `json:"insert-controls,omitempty" yaml:"insert-controls,omitempty"`
	Links          *[]OscalCompleteOscalMetadataLink           `json:"links,omitempty" yaml:"links,omitempty"`
	Params         *[]OscalCompleteOscalControlCommonParameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts          *[]OscalCompleteOscalControlCommonPart      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Props          *[]OscalCompleteOscalMetadataProperty       `json:"props,omitempty" yaml:"props,omitempty"`
	Title          string                                      `json:"title" yaml:"title"`
}

type OscalCompleteOscalProfileInsertControls struct {
	ExcludeControls *[]OscalCompleteOscalProfileSelectControlById `json:"exclude-controls,omitempty" yaml:"exclude-controls,omitempty"`
	IncludeAll      *OscalCompleteOscalControlCommonIncludeAll    `json:"include-all,omitempty" yaml:"include-all,omitempty"`
	IncludeControls *[]OscalCompleteOscalProfileSelectControlById `json:"include-controls,omitempty" yaml:"include-controls,omitempty"`
	Order           string                                        `json:"order,omitempty" yaml:"order,omitempty"`
}

type Addition struct {
	ById     string                                      `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	Links    *[]OscalCompleteOscalMetadataLink           `json:"links,omitempty" yaml:"links,omitempty"`
	Params   *[]OscalCompleteOscalControlCommonParameter `json:"params,omitempty" yaml:"params,omitempty"`
	Parts    *[]OscalCompleteOscalControlCommonPart      `json:"parts,omitempty" yaml:"parts,omitempty"`
	Position string                                      `json:"position,omitempty" yaml:"position,omitempty"`
	Props    *[]OscalCompleteOscalMetadataProperty       `json:"props,omitempty" yaml:"props,omitempty"`
	Title    string                                      `json:"title,omitempty" yaml:"title,omitempty"`
}

type Removal struct {
	ByClass    string `json:"by-class,omitempty" yaml:"by-class,omitempty"`
	ById       string `json:"by-id,omitempty" yaml:"by-id,omitempty"`
	ByItemName string `json:"by-item-name,omitempty" yaml:"by-item-name,omitempty"`
	ByName     string `json:"by-name,omitempty" yaml:"by-name,omitempty"`
	ByNs       string `json:"by-ns,omitempty" yaml:"by-ns,omitempty"`
}

type OscalCompleteOscalSspByComponent struct {
	ComponentUuid        string                                                      `json:"component-uuid" yaml:"component-uuid"`
	Description          string                                                      `json:"description" yaml:"description"`
	Export               *Export                                                     `json:"export,omitempty" yaml:"export,omitempty"`
	ImplementationStatus *OscalCompleteOscalImplementationCommonImplementationStatus `json:"implementation-status,omitempty" yaml:"implementation-status,omitempty"`
	Inherited            *[]InheritedControlImplementation                           `json:"inherited,omitempty" yaml:"inherited,omitempty"`
	Links                *[]OscalCompleteOscalMetadataLink                           `json:"links,omitempty" yaml:"links,omitempty"`
	Props                *[]OscalCompleteOscalMetadataProperty                       `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks              string                                                      `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles     *[]OscalCompleteOscalMetadataResponsibleRole                `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	Satisfied            *[]SatisfiedControlImplementationResponsibility             `json:"satisfied,omitempty" yaml:"satisfied,omitempty"`
	SetParameters        *[]OscalCompleteOscalImplementationCommonSetParameter       `json:"set-parameters,omitempty" yaml:"set-parameters,omitempty"`
	UUID                 string                                                      `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalSspStatement struct {
	ByComponents     *[]OscalCompleteOscalSspByComponent          `json:"by-components,omitempty" yaml:"by-components,omitempty"`
	Links            *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string                                       `json:"statement-id" yaml:"statement-id"`
	UUID             string                                       `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalSspDiagram struct {
	Caption     string                                `json:"caption,omitempty" yaml:"caption,omitempty"`
	Description string                                `json:"description,omitempty" yaml:"description,omitempty"`
	Links       *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props       *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks     string                                `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	UUID        string                                `json:"uuid" yaml:"uuid"`
}

type InformationType struct {
	AvailabilityImpact    *OscalCompleteOscalSspImpact          `json:"availability-impact,omitempty" yaml:"availability-impact,omitempty"`
	Categorizations       *[]InformationTypeCategorization      `json:"categorizations,omitempty" yaml:"categorizations,omitempty"`
	ConfidentialityImpact *OscalCompleteOscalSspImpact          `json:"confidentiality-impact,omitempty" yaml:"confidentiality-impact,omitempty"`
	Description           string                                `json:"description" yaml:"description"`
	IntegrityImpact       *OscalCompleteOscalSspImpact          `json:"integrity-impact,omitempty" yaml:"integrity-impact,omitempty"`
	Links                 *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                 *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Title                 string                                `json:"title" yaml:"title"`
	UUID                  string                                `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type OscalCompleteOscalMetadataHash struct {
	Algorithm string `json:"algorithm" yaml:"algorithm"`
	Value     string `json:"value" yaml:"value"`
}

type OscalCompleteOscalAssessmentCommonLoggedBy struct {
	PartyUuid string `json:"party-uuid" yaml:"party-uuid"`
	RoleId    string `json:"role-id,omitempty" yaml:"role-id,omitempty"`
}

type OscalCompleteOscalComponentDefinitionStatement struct {
	Description      string                                       `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	StatementId      string                                       `json:"statement-id" yaml:"statement-id"`
	UUID             string                                       `json:"uuid" yaml:"uuid"`
}

type IdentifiedSubject struct {
	SubjectPlaceholderUuid string                                                `json:"subject-placeholder-uuid" yaml:"subject-placeholder-uuid"`
	Subjects               []OscalCompleteOscalAssessmentCommonAssessmentSubject `json:"subjects" yaml:"subjects"`
}

type RiskResponseReference struct {
	Links        *[]OscalCompleteOscalMetadataLink                `json:"links,omitempty" yaml:"links,omitempty"`
	Props        *[]OscalCompleteOscalMetadataProperty            `json:"props,omitempty" yaml:"props,omitempty"`
	RelatedTasks *[]OscalCompleteOscalAssessmentCommonRelatedTask `json:"related-tasks,omitempty" yaml:"related-tasks,omitempty"`
	Remarks      string                                           `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponseUuid string                                           `json:"response-uuid" yaml:"response-uuid"`
}

type Export struct {
	Description      string                                 `json:"description,omitempty" yaml:"description,omitempty"`
	Links            *[]OscalCompleteOscalMetadataLink      `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty  `json:"props,omitempty" yaml:"props,omitempty"`
	Provided         *[]ProvidedControlImplementation       `json:"provided,omitempty" yaml:"provided,omitempty"`
	Remarks          string                                 `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	Responsibilities *[]ControlImplementationResponsibility `json:"responsibilities,omitempty" yaml:"responsibilities,omitempty"`
}

type InheritedControlImplementation struct {
	Description      string                                       `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     string                                       `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string                                       `json:"uuid" yaml:"uuid"`
}

type SatisfiedControlImplementationResponsibility struct {
	Description        string                                       `json:"description" yaml:"description"`
	Links              *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props              *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks            string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibilityUuid string                                       `json:"responsibility-uuid,omitempty" yaml:"responsibility-uuid,omitempty"`
	ResponsibleRoles   *[]OscalCompleteOscalMetadataResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID               string                                       `json:"uuid" yaml:"uuid"`
}

type OscalCompleteOscalSspImpact struct {
	AdjustmentJustification string                                `json:"adjustment-justification,omitempty" yaml:"adjustment-justification,omitempty"`
	Base                    string                                `json:"base" yaml:"base"`
	Links                   *[]OscalCompleteOscalMetadataLink     `json:"links,omitempty" yaml:"links,omitempty"`
	Props                   *[]OscalCompleteOscalMetadataProperty `json:"props,omitempty" yaml:"props,omitempty"`
	Selected                string                                `json:"selected,omitempty" yaml:"selected,omitempty"`
}

type InformationTypeCategorization struct {
	InformationTypeIds *[]string `json:"information-type-ids,omitempty" yaml:"information-type-ids,omitempty"`
	System             string    `json:"system" yaml:"system"`
}

type ProvidedControlImplementation struct {
	Description      string                                       `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	Remarks          string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string                                       `json:"uuid" yaml:"uuid"`
}

type ControlImplementationResponsibility struct {
	Description      string                                       `json:"description" yaml:"description"`
	Links            *[]OscalCompleteOscalMetadataLink            `json:"links,omitempty" yaml:"links,omitempty"`
	Props            *[]OscalCompleteOscalMetadataProperty        `json:"props,omitempty" yaml:"props,omitempty"`
	ProvidedUuid     string                                       `json:"provided-uuid,omitempty" yaml:"provided-uuid,omitempty"`
	Remarks          string                                       `json:"remarks,omitempty" yaml:"remarks,omitempty"`
	ResponsibleRoles *[]OscalCompleteOscalMetadataResponsibleRole `json:"responsible-roles,omitempty" yaml:"responsible-roles,omitempty"`
	UUID             string                                       `json:"uuid" yaml:"uuid"`
}
