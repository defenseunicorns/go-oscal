// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    oscalModels, err := UnmarshalOscalModels(bytes)
//    bytes, err = oscalModels.Marshal()

package oscalTypes_1_1_2

import "time"

import "encoding/json"

func UnmarshalOscalModels(data []byte) (OscalModels, error) {
	var r OscalModels
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OscalModels) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OscalModels struct {
	Schema                    *string                        `json:"$schema,omitempty"`
	Catalog                   *Catalog                       `json:"catalog,omitempty"`
	Profile                   *Profile                       `json:"profile,omitempty"`
	ComponentDefinition       *ComponentDefinition           `json:"component-definition,omitempty"`
	SystemSecurityPlan        *SystemSecurityPlanSSP         `json:"system-security-plan,omitempty"`
	AssessmentPlan            *SecurityAssessmentPlanSAP     `json:"assessment-plan,omitempty"`
	AssessmentResults         *SecurityAssessmentResultsSAR  `json:"assessment-results,omitempty"`
	PlanOfActionAndMilestones *PlanOfActionAndMilestonesPOAM `json:"plan-of-action-and-milestones,omitempty"`
}

// An assessment plan, such as those provided by a FedRAMP assessor.
type SecurityAssessmentPlanSAP struct {
	AssessmentAssets                                                                            *AssessmentAssets                    `json:"assessment-assets,omitempty"`
	AssessmentSubjects                                                                          []SubjectOfAssessment                `json:"assessment-subjects,omitempty"`
	BackMatter                                                                                  *BackMatter                          `json:"back-matter,omitempty"`
	ImportSSP                                                                                   ImportSystemSecurityPlan             `json:"import-ssp"`
	// Used to define data objects that are used in the assessment plan, that do not appear in                                       
	// the referenced SSP.                                                                                                           
	LocalDefinitions                                                                            *AssessmentPlanLocalDefinitions      `json:"local-definitions,omitempty"`
	Metadata                                                                                    DocumentMetadata                     `json:"metadata"`
	ReviewedControls                                                                            ReviewedControlsAndControlObjectives `json:"reviewed-controls"`
	Tasks                                                                                       []Task                               `json:"tasks,omitempty"`
	// Used to define various terms and conditions under which an assessment, described by the                                       
	// plan, can be performed. Each child part defines a different type of term or condition.                                        
	TermsAndConditions                                                                          *AssessmentPlanTermsAndConditions    `json:"terms-and-conditions,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                     
	// to reference this assessment plan in this or other OSCAL instances. The locally defined                                       
	// UUID of the assessment plan can be used to reference the data item locally or globally                                        
	// (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject, which                                        
	// means it should be consistently used to identify the same subject across revisions of the                                     
	// document.                                                                                                                     
	UUID                                                                                        string                               `json:"uuid"`
}

// Identifies the assets used to perform this assessment, such as the assessment team,
// scanning tools, and assumptions.
type AssessmentAssets struct {
	AssessmentPlatforms []AssessmentPlatformElement `json:"assessment-platforms"`
	Components          []AssessmentAssetsComponent `json:"components,omitempty"`
}

// Used to represent the toolset used to perform aspects of the assessment.
type AssessmentPlatformElement struct {
	Links                                                                                       []LinkElement   `json:"links,omitempty"`
	Props                                                                                       []Property      `json:"props,omitempty"`
	Remarks                                                                                     *string         `json:"remarks,omitempty"`
	// The title or name for the assessment platform.                                                           
	Title                                                                                       *string         `json:"title,omitempty"`
	UsesComponents                                                                              []UsesComponent `json:"uses-components,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                
	// to reference this assessment platform elsewhere in this or other OSCAL instances. The                    
	// locally defined UUID of the assessment platform can be used to reference the data item                   
	// locally or globally (e.g., in an imported OSCAL instance). This UUID should be assigned                  
	// per-subject, which means it should be consistently used to identify the same subject                     
	// across revisions of the document.                                                                        
	UUID                                                                                        string          `json:"uuid"`
}

// A reference to a local or remote resource, that has a specific relation to the containing
// object.
type LinkElement struct {
	// A resolvable URL reference to a resource.                                                        
	Href                                                                                        string  `json:"href"`
	// A label that indicates the nature of a resource, as a data serialization or format.              
	MediaType                                                                                   *string `json:"media-type,omitempty"`
	// Describes the type of relationship provided by the link's hypertext reference. This can          
	// be an indicator of the link's purpose.                                                           
	Rel                                                                                         *string `json:"rel,omitempty"`
	// In case where the href points to a back-matter/resource, this value will indicate the URI        
	// fragment to append to any rlink associated with the resource. This value MUST be URI             
	// encoded.                                                                                         
	ResourceFragment                                                                            *string `json:"resource-fragment,omitempty"`
	// A textual label to associate with the link, which may be used for presentation in a tool.        
	Text                                                                                        *string `json:"text,omitempty"`
}

// An attribute, characteristic, or quality of the containing object expressed as a
// namespace qualified name/value pair.
type Property struct {
	// A textual label that provides a sub-type or characterization of the property's name.        
	Class                                                                                  *string `json:"class,omitempty"`
	// An identifier for relating distinct sets of properties.                                     
	Group                                                                                  *string `json:"group,omitempty"`
	// A textual label, within a namespace, that uniquely identifies a specific attribute,         
	// characteristic, or quality of the property's containing object.                             
	Name                                                                                   string  `json:"name"`
	// A namespace qualifying the property's name. This allows different organizations to          
	// associate distinct semantics with the same name.                                            
	NS                                                                                     *string `json:"ns,omitempty"`
	Remarks                                                                                *string `json:"remarks,omitempty"`
	// A unique identifier for a property.                                                         
	UUID                                                                                   *string `json:"uuid,omitempty"`
	// Indicates the value of the attribute, characteristic, or quality.                           
	Value                                                                                  string  `json:"value"`
}

// The set of components that are used by the assessment platform.
type UsesComponent struct {
	// A machine-oriented identifier reference to a component that is implemented as part of an                   
	// inventory item.                                                                                            
	ComponentUUID                                                                              string             `json:"component-uuid"`
	Links                                                                                      []LinkElement      `json:"links,omitempty"`
	Props                                                                                      []Property         `json:"props,omitempty"`
	Remarks                                                                                    *string            `json:"remarks,omitempty"`
	ResponsibleParties                                                                         []ResponsibleParty `json:"responsible-parties,omitempty"`
}

// A reference to a set of persons and/or organizations that have responsibility for
// performing the referenced role in the context of the containing object.
type ResponsibleParty struct {
	Links                                         []LinkElement `json:"links,omitempty"`
	PartyUuids                                    []string      `json:"party-uuids"`
	Props                                         []Property    `json:"props,omitempty"`
	Remarks                                       *string       `json:"remarks,omitempty"`
	// A reference to a role performed by a party.              
	RoleID                                        string        `json:"role-id"`
}

// A defined component that can be part of an implemented system.
type AssessmentAssetsComponent struct {
	// A description of the component, including information about its function.                                             
	Description                                                                                 string                       `json:"description"`
	Links                                                                                       []LinkElement                `json:"links,omitempty"`
	Props                                                                                       []Property                   `json:"props,omitempty"`
	Protocols                                                                                   []ServiceProtocolInformation `json:"protocols,omitempty"`
	// A summary of the technological or business purpose of the component.                                                  
	Purpose                                                                                     *string                      `json:"purpose,omitempty"`
	Remarks                                                                                     *string                      `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole            `json:"responsible-roles,omitempty"`
	// Describes the operational status of the system component.                                                             
	Status                                                                                      ComponentStatus              `json:"status"`
	// A human readable name for the system component.                                                                       
	Title                                                                                       string                       `json:"title"`
	// A category describing the purpose of the component.                                                                   
	Type                                                                                        string                       `json:"type"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                             
	// to reference this component elsewhere in this or other OSCAL instances. The locally                                   
	// defined UUID of the component can be used to reference the data item locally or globally                              
	// (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject, which                                
	// means it should be consistently used to identify the same subject across revisions of the                             
	// document.                                                                                                             
	UUID                                                                                        string                       `json:"uuid"`
}

// Information about the protocol used to provide a service.
type ServiceProtocolInformation struct {
	// The common name of the protocol, which should be the appropriate "service name" from the             
	// IANA Service Name and Transport Protocol Port Number Registry.                                       
	Name                                                                                        string      `json:"name"`
	PortRanges                                                                                  []PortRange `json:"port-ranges,omitempty"`
	// A human readable name for the protocol (e.g., Transport Layer Security).                             
	Title                                                                                       *string     `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used            
	// to reference this service protocol information elsewhere in this or other OSCAL                      
	// instances. The locally defined UUID of the service protocol can be used to reference the             
	// data item locally or globally (e.g., in an imported OSCAL instance). This UUID should be             
	// assigned per-subject, which means it should be consistently used to identify the same                
	// subject across revisions of the document.                                                            
	UUID                                                                                        *string     `json:"uuid,omitempty"`
}

// Where applicable this is the IPv4 port range on which the service operates.
type PortRange struct {
	// Indicates the ending port number in a port range             
	End                                                  *int64     `json:"end,omitempty"`
	// Indicates the starting port number in a port range           
	Start                                                *int64     `json:"start,omitempty"`
	// Indicates the transport type.                                
	Transport                                            *Transport `json:"transport,omitempty"`
}

// A reference to a role with responsibility for performing a function relative to the
// containing object, optionally associated with a set of persons and/or organizations that
// perform that role.
type ResponsibleRole struct {
	Links                                                        []LinkElement `json:"links,omitempty"`
	PartyUuids                                                   []string      `json:"party-uuids,omitempty"`
	Props                                                        []Property    `json:"props,omitempty"`
	Remarks                                                      *string       `json:"remarks,omitempty"`
	// A human-oriented identifier reference to a role performed.              
	RoleID                                                       string        `json:"role-id"`
}

// Describes the operational status of the system component.
type ComponentStatus struct {
	Remarks                   *string     `json:"remarks,omitempty"`
	// The operational status.            
	State                     PurpleState `json:"state"`
}

// Identifies system elements being assessed, such as components, inventory items, and
// locations. In the assessment plan, this identifies a planned assessment subject. In the
// assessment results this is an actual assessment subject, and reflects any changes from
// the plan. exactly what will be the focus of this assessment. Any subjects not identified
// in this way are out-of-scope.
type SubjectOfAssessment struct {
	// A human-readable description of the collection of subjects being included in this                                  
	// assessment.                                                                                                        
	Description                                                                                 *string                   `json:"description,omitempty"`
	ExcludeSubjects                                                                             []SelectAssessmentSubject `json:"exclude-subjects,omitempty"`
	IncludeAll                                                                                  *IncludeAll               `json:"include-all,omitempty"`
	IncludeSubjects                                                                             []SelectAssessmentSubject `json:"include-subjects,omitempty"`
	Links                                                                                       []LinkElement             `json:"links,omitempty"`
	Props                                                                                       []Property                `json:"props,omitempty"`
	Remarks                                                                                     *string                   `json:"remarks,omitempty"`
	// Indicates the type of assessment subject, such as a component, inventory, item, location,                          
	// or party represented by this selection statement.                                                                  
	Type                                                                                        string                    `json:"type"`
}

// Identifies a set of assessment subjects to include/exclude by UUID.
type SelectAssessmentSubject struct {
	Links                                                                                      []LinkElement `json:"links,omitempty"`
	Props                                                                                      []Property    `json:"props,omitempty"`
	Remarks                                                                                    *string       `json:"remarks,omitempty"`
	// A machine-oriented identifier reference to a component, inventory-item, location, party,              
	// user, or resource using it's UUID.                                                                    
	SubjectUUID                                                                                string        `json:"subject-uuid"`
	// Used to indicate the type of object pointed to by the uuid-ref within a subject.                      
	Type                                                                                       string        `json:"type"`
}

// Include all controls from the imported catalog or profile resources.
type IncludeAll struct {
}

// A collection of resources that may be referenced from within the OSCAL document instance.
type BackMatter struct {
	Resources []Resource `json:"resources,omitempty"`
}

// A resource associated with content in the containing document instance. A resource may be
// directly included in the document using base64 encoding or may point to one or more
// equivalent internet resources.
type Resource struct {
	// A resource encoded using the Base64 alphabet defined by RFC 2045.                                           
	Base64                                                                                    *Base64              `json:"base64,omitempty"`
	// An optional citation consisting of end note text using structured markup.                                   
	Citation                                                                                  *Citation            `json:"citation,omitempty"`
	// An optional short summary of the resource used to indicate the purpose of the resource.                     
	Description                                                                               *string              `json:"description,omitempty"`
	DocumentIDS                                                                               []DocumentIdentifier `json:"document-ids,omitempty"`
	Props                                                                                     []Property           `json:"props,omitempty"`
	Remarks                                                                                   *string              `json:"remarks,omitempty"`
	Rlinks                                                                                    []ResourceLink       `json:"rlinks,omitempty"`
	// An optional name given to the resource, which may be used by a tool for display and                         
	// navigation.                                                                                                 
	Title                                                                                     *string              `json:"title,omitempty"`
	// A unique identifier for a resource.                                                                         
	UUID                                                                                      string               `json:"uuid"`
}

// A resource encoded using the Base64 alphabet defined by RFC 2045.
type Base64 struct {
	// Name of the file before it was encoded as Base64 to be embedded in a resource. This is        
	// the name that will be assigned to the file when the file is decoded.                          
	Filename                                                                                 *string `json:"filename,omitempty"`
	// A label that indicates the nature of a resource, as a data serialization or format.           
	MediaType                                                                                *string `json:"media-type,omitempty"`
	Value                                                                                    string  `json:"value"`
}

// An optional citation consisting of end note text using structured markup.
type Citation struct {
	Links                      []LinkElement `json:"links,omitempty"`
	Props                      []Property    `json:"props,omitempty"`
	// A line of citation text.              
	Text                       string        `json:"text"`
}

// A document identifier qualified by an identifier scheme.
type DocumentIdentifier struct {
	Identifier                                                                                 string  `json:"identifier"`
	// Qualifies the kind of document identifier using a URI. If the scheme is not provided the        
	// value of the element will be interpreted as a string of characters.                             
	Scheme                                                                                     *string `json:"scheme,omitempty"`
}

// A URL-based pointer to an external resource with an optional hash for verification and
// change detection.
type ResourceLink struct {
	Hashes                                                                                []Hash  `json:"hashes,omitempty"`
	// A resolvable URL pointing to the referenced resource.                                      
	Href                                                                                  string  `json:"href"`
	// A label that indicates the nature of a resource, as a data serialization or format.        
	MediaType                                                                             *string `json:"media-type,omitempty"`
}

// A representation of a cryptographic digest generated over a resource using a specified
// hash algorithm.
type Hash struct {
	// The digest method by which a hash is derived.       
	Algorithm                                       string `json:"algorithm"`
	Value                                           string `json:"value"`
}

// Used by the assessment plan and POA&M to import information about the system.
type ImportSystemSecurityPlan struct {
	// A resolvable URL reference to the system security plan for the system being assessed.        
	Href                                                                                    string  `json:"href"`
	Remarks                                                                                 *string `json:"remarks,omitempty"`
}

// Used to define data objects that are used in the assessment plan, that do not appear in
// the referenced SSP.
type AssessmentPlanLocalDefinitions struct {
	Activities           []Activity                           `json:"activities,omitempty"`
	Components           []AssessmentAssetsComponent          `json:"components,omitempty"`
	InventoryItems       []InventoryItem                      `json:"inventory-items,omitempty"`
	ObjectivesAndMethods []AssessmentSpecificControlObjective `json:"objectives-and-methods,omitempty"`
	Remarks              *string                              `json:"remarks,omitempty"`
	Users                []SystemUser                         `json:"users,omitempty"`
}

// Identifies an assessment or related process that can be performed. In the assessment
// plan, this is an intended activity which may be associated with an assessment task. In
// the assessment results, this an activity that was actually performed as part of an
// assessment.
type Activity struct {
	// A human-readable description of this included activity.                                                                        
	Description                                                                                 string                                `json:"description"`
	Links                                                                                       []LinkElement                         `json:"links,omitempty"`
	Props                                                                                       []Property                            `json:"props,omitempty"`
	RelatedControls                                                                             *ReviewedControlsAndControlObjectives `json:"related-controls,omitempty"`
	Remarks                                                                                     *string                               `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole                     `json:"responsible-roles,omitempty"`
	Steps                                                                                       []Step                                `json:"steps,omitempty"`
	// The title for this included activity.                                                                                          
	Title                                                                                       *string                               `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                      
	// to reference this assessment activity elsewhere in this or other OSCAL instances. The                                          
	// locally defined UUID of the activity can be used to reference the data item locally or                                         
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                                      
	// which means it should be consistently used to identify the same subject across revisions                                       
	// of the document.                                                                                                               
	UUID                                                                                        string                                `json:"uuid"`
}

// Identifies the controls being assessed and their control objectives.
type ReviewedControlsAndControlObjectives struct {
	ControlObjectiveSelections                            []ReferencedControlObjectives `json:"control-objective-selections,omitempty"`
	ControlSelections                                     []AssessedControls            `json:"control-selections"`
	// A human-readable description of control objectives.                              
	Description                                           *string                       `json:"description,omitempty"`
	Links                                                 []LinkElement                 `json:"links,omitempty"`
	Props                                                 []Property                    `json:"props,omitempty"`
	Remarks                                               *string                       `json:"remarks,omitempty"`
}

// Identifies the control objectives of the assessment. In the assessment plan, these are
// the planned objectives. In the assessment results, these are the assessed objectives, and
// reflects any changes from the plan.
type ReferencedControlObjectives struct {
	// A human-readable description of this collection of control objectives.                  
	Description                                                              *string           `json:"description,omitempty"`
	ExcludeObjectives                                                        []SelectObjective `json:"exclude-objectives,omitempty"`
	IncludeAll                                                               *IncludeAll       `json:"include-all,omitempty"`
	IncludeObjectives                                                        []SelectObjective `json:"include-objectives,omitempty"`
	Links                                                                    []LinkElement     `json:"links,omitempty"`
	Props                                                                    []Property        `json:"props,omitempty"`
	Remarks                                                                  *string           `json:"remarks,omitempty"`
}

// Used to select a control objective for inclusion/exclusion based on the control
// objective's identifier.
type SelectObjective struct {
	// Points to an assessment objective.       
	ObjectiveID                          string `json:"objective-id"`
}

// Identifies the controls being assessed. In the assessment plan, these are the planned
// controls. In the assessment results, these are the actual controls, and reflects any
// changes from the plan.
type AssessedControls struct {
	// A human-readable description of in-scope controls specified for assessment.                                 
	Description                                                                   *string                          `json:"description,omitempty"`
	ExcludeControls                                                               []ControlSelectionExcludeControl `json:"exclude-controls,omitempty"`
	IncludeAll                                                                    *IncludeAll                      `json:"include-all,omitempty"`
	IncludeControls                                                               []ControlSelectionExcludeControl `json:"include-controls,omitempty"`
	Links                                                                         []LinkElement                    `json:"links,omitempty"`
	Props                                                                         []Property                       `json:"props,omitempty"`
	Remarks                                                                       *string                          `json:"remarks,omitempty"`
}

// Used to select a control for inclusion/exclusion based on one or more control
// identifiers. A set of statement identifiers can be used to target the inclusion/exclusion
// to only specific control statements providing more granularity over the specific
// statements that are within the asessment scope.
type ControlSelectionExcludeControl struct {
	// A reference to a control with a corresponding id value. When referencing an externally         
	// defined control, the Control Identifier Reference must be used in the context of the           
	// external / imported OSCAL instance (e.g., uri-reference).                                      
	ControlID                                                                                string   `json:"control-id"`
	StatementIDS                                                                             []string `json:"statement-ids,omitempty"`
}

// Identifies an individual step in a series of steps related to an activity, such as an
// assessment test or examination procedure.
type Step struct {
	// A human-readable description of this step.                                                                                     
	Description                                                                                 string                                `json:"description"`
	Links                                                                                       []LinkElement                         `json:"links,omitempty"`
	Props                                                                                       []Property                            `json:"props,omitempty"`
	Remarks                                                                                     *string                               `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole                     `json:"responsible-roles,omitempty"`
	ReviewedControls                                                                            *ReviewedControlsAndControlObjectives `json:"reviewed-controls,omitempty"`
	// The title for this step.                                                                                                       
	Title                                                                                       *string                               `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                      
	// to reference this step elsewhere in this or other OSCAL instances. The locally defined                                         
	// UUID of the step (in a series of steps) can be used to reference the data item locally or                                      
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                                      
	// which means it should be consistently used to identify the same subject across revisions                                       
	// of the document.                                                                                                               
	UUID                                                                                        string                                `json:"uuid"`
}

// A single managed inventory item within the system.
type InventoryItem struct {
	// A summary of the inventory item stating its purpose within the system.                                          
	Description                                                                                 string                 `json:"description"`
	ImplementedComponents                                                                       []ImplementedComponent `json:"implemented-components,omitempty"`
	Links                                                                                       []LinkElement          `json:"links,omitempty"`
	Props                                                                                       []Property             `json:"props,omitempty"`
	Remarks                                                                                     *string                `json:"remarks,omitempty"`
	ResponsibleParties                                                                          []ResponsibleParty     `json:"responsible-parties,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                       
	// to reference this inventory item elsewhere in this or other OSCAL instances. The locally                        
	// defined UUID of the inventory item can be used to reference the data item locally or                            
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                       
	// which means it should be consistently used to identify the same subject across revisions                        
	// of the document.                                                                                                
	UUID                                                                                        string                 `json:"uuid"`
}

// The set of components that are implemented in a given system inventory item.
type ImplementedComponent struct {
	// A machine-oriented identifier reference to a component that is implemented as part of an                   
	// inventory item.                                                                                            
	ComponentUUID                                                                              string             `json:"component-uuid"`
	Links                                                                                      []LinkElement      `json:"links,omitempty"`
	Props                                                                                      []Property         `json:"props,omitempty"`
	Remarks                                                                                    *string            `json:"remarks,omitempty"`
	ResponsibleParties                                                                         []ResponsibleParty `json:"responsible-parties,omitempty"`
}

// A local definition of a control objective for this assessment. Uses catalog syntax for
// control objective and assessment actions.
type AssessmentSpecificControlObjective struct {
	// A reference to a control with a corresponding id value. When referencing an externally              
	// defined control, the Control Identifier Reference must be used in the context of the                
	// external / imported OSCAL instance (e.g., uri-reference).                                           
	ControlID                                                                                string        `json:"control-id"`
	// A human-readable description of this control objective.                                             
	Description                                                                              *string       `json:"description,omitempty"`
	Links                                                                                    []LinkElement `json:"links,omitempty"`
	Parts                                                                                    []PartElement `json:"parts"`
	Props                                                                                    []Property    `json:"props,omitempty"`
	Remarks                                                                                  *string       `json:"remarks,omitempty"`
}

// An annotated, markup-based textual element of a control's or catalog group's definition,
// or a child of another part.
type PartElement struct {
	// An optional textual providing a sub-type or characterization of the part's name, or a                 
	// category to which the part belongs.                                                                   
	Class                                                                                      *string       `json:"class,omitempty"`
	// A unique identifier for the part.                                                                     
	ID                                                                                         *string       `json:"id,omitempty"`
	Links                                                                                      []LinkElement `json:"links,omitempty"`
	// A textual label that uniquely identifies the part's semantic type, which exists in a                  
	// value space qualified by the ns.                                                                      
	Name                                                                                       string        `json:"name"`
	// An optional namespace qualifying the part's name. This allows different organizations to              
	// associate distinct semantics with the same name.                                                      
	NS                                                                                         *string       `json:"ns,omitempty"`
	Parts                                                                                      []PartElement `json:"parts,omitempty"`
	Props                                                                                      []Property    `json:"props,omitempty"`
	// Permits multiple paragraphs, lists, tables etc.                                                       
	Prose                                                                                      *string       `json:"prose,omitempty"`
	// An optional name given to the part, which may be used by a tool for display and                       
	// navigation.                                                                                           
	Title                                                                                      *string       `json:"title,omitempty"`
}

// A type of user that interacts with the system based on an associated role.
type SystemUser struct {
	AuthorizedPrivileges                                                                        []Privilege   `json:"authorized-privileges,omitempty"`
	// A summary of the user's purpose within the system.                                                     
	Description                                                                                 *string       `json:"description,omitempty"`
	Links                                                                                       []LinkElement `json:"links,omitempty"`
	Props                                                                                       []Property    `json:"props,omitempty"`
	Remarks                                                                                     *string       `json:"remarks,omitempty"`
	RoleIDS                                                                                     []string      `json:"role-ids,omitempty"`
	// A short common name, abbreviation, or acronym for the user.                                            
	ShortName                                                                                   *string       `json:"short-name,omitempty"`
	// A name given to the user, which may be used by a tool for display and navigation.                      
	Title                                                                                       *string       `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used              
	// to reference this user class elsewhere in this or other OSCAL instances. The locally                   
	// defined UUID of the system user can be used to reference the data item locally or                      
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,              
	// which means it should be consistently used to identify the same subject across revisions               
	// of the document.                                                                                       
	UUID                                                                                        string        `json:"uuid"`
}

// Identifies a specific system privilege held by the user, along with an associated
// description and/or rationale for the privilege.
type Privilege struct {
	// A summary of the privilege's purpose within the system.         
	Description                                               *string  `json:"description,omitempty"`
	FunctionsPerformed                                        []string `json:"functions-performed"`
	// A human readable name for the privilege.                        
	Title                                                     string   `json:"title"`
}

// Provides information about the containing document, and defines concepts that are shared
// across the document.
type DocumentMetadata struct {
	Actions                                                                                 []Action               `json:"actions,omitempty"`
	DocumentIDS                                                                             []DocumentIdentifier   `json:"document-ids,omitempty"`
	LastModified                                                                            time.Time              `json:"last-modified"`
	Links                                                                                   []LinkElement          `json:"links,omitempty"`
	Locations                                                                               []Location             `json:"locations,omitempty"`
	OscalVersion                                                                            string                 `json:"oscal-version"`
	Parties                                                                                 []PartyElement         `json:"parties,omitempty"`
	Props                                                                                   []Property             `json:"props,omitempty"`
	Published                                                                               *time.Time             `json:"published,omitempty"`
	Remarks                                                                                 *string                `json:"remarks,omitempty"`
	ResponsibleParties                                                                      []ResponsibleParty     `json:"responsible-parties,omitempty"`
	Revisions                                                                               []RevisionHistoryEntry `json:"revisions,omitempty"`
	Roles                                                                                   []Role                 `json:"roles,omitempty"`
	// A name given to the document, which may be used by a tool for display and navigation.                       
	Title                                                                                   string                 `json:"title"`
	Version                                                                                 string                 `json:"version"`
}

// An action applied by a role within a given party to the content.
type Action struct {
	// The date and time when the action occurred.                                                                
	Date                                                                                       *time.Time         `json:"date,omitempty"`
	Links                                                                                      []LinkElement      `json:"links,omitempty"`
	Props                                                                                      []Property         `json:"props,omitempty"`
	Remarks                                                                                    *string            `json:"remarks,omitempty"`
	ResponsibleParties                                                                         []ResponsibleParty `json:"responsible-parties,omitempty"`
	// Specifies the action type system used.                                                                     
	System                                                                                     string             `json:"system"`
	// The type of action documented by the assembly, such as an approval.                                        
	Type                                                                                       string             `json:"type"`
	// A unique identifier that can be used to reference this defined action elsewhere in an                      
	// OSCAL document. A UUID should be consistently used for a given location across revisions                   
	// of the document.                                                                                           
	UUID                                                                                       string             `json:"uuid"`
}

// A physical point of presence, which may be associated with people, organizations, or
// other concepts within the current or linked OSCAL document.
type Location struct {
	Address                                                                                 *Address          `json:"address,omitempty"`
	EmailAddresses                                                                          []string          `json:"email-addresses,omitempty"`
	Links                                                                                   []LinkElement     `json:"links,omitempty"`
	Props                                                                                   []Property        `json:"props,omitempty"`
	Remarks                                                                                 *string           `json:"remarks,omitempty"`
	TelephoneNumbers                                                                        []TelephoneNumber `json:"telephone-numbers,omitempty"`
	// A name given to the location, which may be used by a tool for display and navigation.                  
	Title                                                                                   *string           `json:"title,omitempty"`
	Urls                                                                                    []string          `json:"urls,omitempty"`
	// A unique ID for the location, for reference.                                                           
	UUID                                                                                    string            `json:"uuid"`
}

// A postal address for the location.
type Address struct {
	AddrLines                                                                 []string `json:"addr-lines,omitempty"`
	// City, town or geographical region for the mailing address.                      
	City                                                                      *string  `json:"city,omitempty"`
	// The ISO 3166-1 alpha-2 country code for the mailing address.                    
	Country                                                                   *string  `json:"country,omitempty"`
	// Postal or ZIP code for mailing address.                                         
	PostalCode                                                                *string  `json:"postal-code,omitempty"`
	// State, province or analogous geographical region for a mailing address.         
	State                                                                     *string  `json:"state,omitempty"`
	// Indicates the type of address.                                                  
	Type                                                                      *string  `json:"type,omitempty"`
}

// A telephone service number as defined by ITU-T E.164.
type TelephoneNumber struct {
	Number                                string  `json:"number"`
	// Indicates the type of phone number.        
	Type                                  *string `json:"type,omitempty"`
}

// An organization or person, which may be associated with roles or other concepts within
// the current or linked OSCAL document.
type PartyElement struct {
	Addresses                                                                                 []Address                 `json:"addresses,omitempty"`
	EmailAddresses                                                                            []string                  `json:"email-addresses,omitempty"`
	ExternalIDS                                                                               []PartyExternalIdentifier `json:"external-ids,omitempty"`
	Links                                                                                     []LinkElement             `json:"links,omitempty"`
	LocationUuids                                                                             []string                  `json:"location-uuids,omitempty"`
	MemberOfOrganizations                                                                     []string                  `json:"member-of-organizations,omitempty"`
	// The full name of the party. This is typically the legal name associated with the party.                          
	Name                                                                                      *string                   `json:"name,omitempty"`
	Props                                                                                     []Property                `json:"props,omitempty"`
	Remarks                                                                                   *string                   `json:"remarks,omitempty"`
	// A short common name, abbreviation, or acronym for the party.                                                     
	ShortName                                                                                 *string                   `json:"short-name,omitempty"`
	TelephoneNumbers                                                                          []TelephoneNumber         `json:"telephone-numbers,omitempty"`
	// A category describing the kind of party the object describes.                                                    
	Type                                                                                      PartyType                 `json:"type"`
	// A unique identifier for the party.                                                                               
	UUID                                                                                      string                    `json:"uuid"`
}

// An identifier for a person or organization using a designated scheme. e.g. an Open
// Researcher and Contributor ID (ORCID).
type PartyExternalIdentifier struct {
	ID                                           string `json:"id"`
	// Indicates the type of external identifier.       
	Scheme                                       string `json:"scheme"`
}

// An entry in a sequential list of revisions to the containing document, expected to be in
// reverse chronological order (i.e. latest first).
type RevisionHistoryEntry struct {
	LastModified                                                                         *time.Time    `json:"last-modified,omitempty"`
	Links                                                                                []LinkElement `json:"links,omitempty"`
	OscalVersion                                                                         *string       `json:"oscal-version,omitempty"`
	Props                                                                                []Property    `json:"props,omitempty"`
	Published                                                                            *time.Time    `json:"published,omitempty"`
	Remarks                                                                              *string       `json:"remarks,omitempty"`
	// A name given to the document revision, which may be used by a tool for display and              
	// navigation.                                                                                     
	Title                                                                                *string       `json:"title,omitempty"`
	Version                                                                              string        `json:"version"`
}

// Defines a function, which might be assigned to a party in a specific situation.
type Role struct {
	// A summary of the role's purpose and associated responsibilities.                               
	Description                                                                         *string       `json:"description,omitempty"`
	// A unique identifier for the role.                                                              
	ID                                                                                  string        `json:"id"`
	Links                                                                               []LinkElement `json:"links,omitempty"`
	Props                                                                               []Property    `json:"props,omitempty"`
	Remarks                                                                             *string       `json:"remarks,omitempty"`
	// A short common name, abbreviation, or acronym for the role.                                    
	ShortName                                                                           *string       `json:"short-name,omitempty"`
	// A name given to the role, which may be used by a tool for display and navigation.              
	Title                                                                               string        `json:"title"`
}

// Represents a scheduled event or milestone, which may be associated with a series of
// assessment actions.
type Task struct {
	AssociatedActivities                                                                        []AssociatedActivity  `json:"associated-activities,omitempty"`
	Dependencies                                                                                []TaskDependency      `json:"dependencies,omitempty"`
	// A human-readable description of this task.                                                                     
	Description                                                                                 *string               `json:"description,omitempty"`
	Links                                                                                       []LinkElement         `json:"links,omitempty"`
	Props                                                                                       []Property            `json:"props,omitempty"`
	Remarks                                                                                     *string               `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole     `json:"responsible-roles,omitempty"`
	Subjects                                                                                    []SubjectOfAssessment `json:"subjects,omitempty"`
	Tasks                                                                                       []Task                `json:"tasks,omitempty"`
	// The timing under which the task is intended to occur.                                                          
	Timing                                                                                      *EventTiming          `json:"timing,omitempty"`
	// The title for this task.                                                                                       
	Title                                                                                       string                `json:"title"`
	// The type of task.                                                                                              
	Type                                                                                        string                `json:"type"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                      
	// to reference this task elsewhere in this or other OSCAL instances. The locally defined                         
	// UUID of the task can be used to reference the data item locally or globally (e.g., in an                       
	// imported OSCAL instance). This UUID should be assigned per-subject, which means it should                      
	// be consistently used to identify the same subject across revisions of the document.                            
	UUID                                                                                        string                `json:"uuid"`
}

// Identifies an individual activity to be performed as part of a task.
type AssociatedActivity struct {
	// A machine-oriented identifier reference to an activity defined in the list of activities.                      
	ActivityUUID                                                                                string                `json:"activity-uuid"`
	Links                                                                                       []LinkElement         `json:"links,omitempty"`
	Props                                                                                       []Property            `json:"props,omitempty"`
	Remarks                                                                                     *string               `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole     `json:"responsible-roles,omitempty"`
	Subjects                                                                                    []SubjectOfAssessment `json:"subjects"`
}

// Used to indicate that a task is dependent on another task.
type TaskDependency struct {
	Remarks                                                     *string `json:"remarks,omitempty"`
	// A machine-oriented identifier reference to a unique task.        
	TaskUUID                                                    string  `json:"task-uuid"`
}

// The timing under which the task is intended to occur.
type EventTiming struct {
	// The task is intended to occur at the specified frequency.                           
	AtFrequency                                                      *FrequencyCondition   `json:"at-frequency,omitempty"`
	// The task is intended to occur on the specified date.                                
	OnDate                                                           *OnDateCondition      `json:"on-date,omitempty"`
	// The task is intended to occur within the specified date range.                      
	WithinDateRange                                                  *OnDateRangeCondition `json:"within-date-range,omitempty"`
}

// The task is intended to occur at the specified frequency.
type FrequencyCondition struct {
	// The task must occur after the specified period has elapsed.         
	Period                                                        int64    `json:"period"`
	// The unit of time for the period.                                    
	Unit                                                          TimeUnit `json:"unit"`
}

// The task is intended to occur on the specified date.
type OnDateCondition struct {
	// The task must occur on the specified date.          
	Date                                         time.Time `json:"date"`
}

// The task is intended to occur within the specified date range.
type OnDateRangeCondition struct {
	// The task must occur on or before the specified date.          
	End                                                    time.Time `json:"end"`
	// The task must occur on or after the specified date.           
	Start                                                  time.Time `json:"start"`
}

// Used to define various terms and conditions under which an assessment, described by the
// plan, can be performed. Each child part defines a different type of term or condition.
type AssessmentPlanTermsAndConditions struct {
	Parts []AssessmentPart `json:"parts,omitempty"`
}

// A partition of an assessment plan or results or a child of another part.
type AssessmentPart struct {
	// A textual label that provides a sub-type or characterization of the part's name. This can                 
	// be used to further distinguish or discriminate between the semantics of multiple parts of                 
	// the same control with the same name and ns.                                                               
	Class                                                                                       *string          `json:"class,omitempty"`
	Links                                                                                       []LinkElement    `json:"links,omitempty"`
	// A textual label that uniquely identifies the part's semantic type.                                        
	Name                                                                                        string           `json:"name"`
	// A namespace qualifying the part's name. This allows different organizations to associate                  
	// distinct semantics with the same name.                                                                    
	NS                                                                                          *string          `json:"ns,omitempty"`
	Parts                                                                                       []AssessmentPart `json:"parts,omitempty"`
	Props                                                                                       []Property       `json:"props,omitempty"`
	// Permits multiple paragraphs, lists, tables etc.                                                           
	Prose                                                                                       *string          `json:"prose,omitempty"`
	// A name given to the part, which may be used by a tool for display and navigation.                         
	Title                                                                                       *string          `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                 
	// to reference this part elsewhere in this or other OSCAL instances. The locally defined                    
	// UUID of the part can be used to reference the data item locally or globally (e.g., in an                  
	// ported OSCAL instance). This UUID should be assigned per-subject, which means it should                   
	// be consistently used to identify the same subject across revisions of the document.                       
	UUID                                                                                        *string          `json:"uuid,omitempty"`
}

// Security assessment results, such as those provided by a FedRAMP assessor in the FedRAMP
// Security Assessment Report.
type SecurityAssessmentResultsSAR struct {
	BackMatter                                                                                  *BackMatter                        `json:"back-matter,omitempty"`
	ImportAp                                                                                    ImportAssessmentPlan               `json:"import-ap"`
	// Used to define data objects that are used in the assessment plan, that do not appear in                                     
	// the referenced SSP.                                                                                                         
	LocalDefinitions                                                                            *AssessmentResultsLocalDefinitions `json:"local-definitions,omitempty"`
	Metadata                                                                                    DocumentMetadata                   `json:"metadata"`
	Results                                                                                     []AssessmentResult                 `json:"results"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                   
	// to reference this assessment results instance in this or other OSCAL instances. The                                         
	// locally defined UUID of the assessment result can be used to reference the data item                                        
	// locally or globally (e.g., in an imported OSCAL instance). This UUID should be assigned                                     
	// per-subject, which means it should be consistently used to identify the same subject                                        
	// across revisions of the document.                                                                                           
	UUID                                                                                        string                             `json:"uuid"`
}

// Used by assessment-results to import information about the original plan for assessing
// the system.
type ImportAssessmentPlan struct {
	// A resolvable URL reference to the assessment plan governing the assessment activities.        
	Href                                                                                     string  `json:"href"`
	Remarks                                                                                  *string `json:"remarks,omitempty"`
}

// Used to define data objects that are used in the assessment plan, that do not appear in
// the referenced SSP.
type AssessmentResultsLocalDefinitions struct {
	Activities           []Activity                           `json:"activities,omitempty"`
	ObjectivesAndMethods []AssessmentSpecificControlObjective `json:"objectives-and-methods,omitempty"`
	Remarks              *string                              `json:"remarks,omitempty"`
}

// Used by the assessment results and POA&M. In the assessment results, this identifies all
// of the assessment observations and findings, initial and residual risks, deviations, and
// disposition. In the POA&M, this identifies initial and residual risks, deviations, and
// disposition.
type AssessmentResult struct {
	// A log of all assessment-related actions taken.                                                                                
	AssessmentLog                                                                               *AssessmentLog                       `json:"assessment-log,omitempty"`
	Attestations                                                                                []AttestationStatements              `json:"attestations,omitempty"`
	// A human-readable description of this set of test results.                                                                     
	Description                                                                                 string                               `json:"description"`
	// Date/time stamp identifying the end of the evidence collection reflected in these                                             
	// results. In a continuous motoring scenario, this may contain the same value as start if                                       
	// appropriate.                                                                                                                  
	End                                                                                         *time.Time                           `json:"end,omitempty"`
	Findings                                                                                    []Finding                            `json:"findings,omitempty"`
	Links                                                                                       []LinkElement                        `json:"links,omitempty"`
	// Used to define data objects that are used in the assessment plan, that do not appear in                                       
	// the referenced SSP.                                                                                                           
	LocalDefinitions                                                                            *ResultLocalDefinitions              `json:"local-definitions,omitempty"`
	Observations                                                                                []Observation                        `json:"observations,omitempty"`
	Props                                                                                       []Property                           `json:"props,omitempty"`
	Remarks                                                                                     *string                              `json:"remarks,omitempty"`
	ReviewedControls                                                                            ReviewedControlsAndControlObjectives `json:"reviewed-controls"`
	Risks                                                                                       []IdentifiedRisk                     `json:"risks,omitempty"`
	// Date/time stamp identifying the start of the evidence collection reflected in these                                           
	// results.                                                                                                                      
	Start                                                                                       time.Time                            `json:"start"`
	// The title for this set of results.                                                                                            
	Title                                                                                       string                               `json:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                     
	// to reference this set of results in this or other OSCAL instances. The locally defined                                        
	// UUID of the assessment result can be used to reference the data item locally or globally                                      
	// (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject, which                                        
	// means it should be consistently used to identify the same subject across revisions of the                                     
	// document.                                                                                                                     
	UUID                                                                                        string                               `json:"uuid"`
}

// A log of all assessment-related actions taken.
type AssessmentLog struct {
	Entries []AssessmentLogEntry `json:"entries"`
}

// Identifies the result of an action and/or task that occurred as part of executing an
// assessment plan or an assessment event that occurred in producing the assessment results.
type AssessmentLogEntry struct {
	// A human-readable description of this event.                                                              
	Description                                                                                 *string         `json:"description,omitempty"`
	// Identifies the end date and time of an event. If the event is a point in time, the start                 
	// and end will be the same date and time.                                                                  
	End                                                                                         *time.Time      `json:"end,omitempty"`
	Links                                                                                       []LinkElement   `json:"links,omitempty"`
	LoggedBy                                                                                    []LoggedBy      `json:"logged-by,omitempty"`
	Props                                                                                       []Property      `json:"props,omitempty"`
	RelatedTasks                                                                                []TaskReference `json:"related-tasks,omitempty"`
	Remarks                                                                                     *string         `json:"remarks,omitempty"`
	// Identifies the start date and time of an event.                                                          
	Start                                                                                       time.Time       `json:"start"`
	// The title for this event.                                                                                
	Title                                                                                       *string         `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                
	// to reference an assessment event in this or other OSCAL instances. The locally defined                   
	// UUID of the assessment log entry can be used to reference the data item locally or                       
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                
	// which means it should be consistently used to identify the same subject across revisions                 
	// of the document.                                                                                         
	UUID                                                                                        string          `json:"uuid"`
}

// Used to indicate who created a log entry in what role.
type LoggedBy struct {
	// A machine-oriented identifier reference to the party who is making the log entry.        
	PartyUUID                                                                           string  `json:"party-uuid"`
	// A point to the role-id of the role in which the party is making the log entry.           
	RoleID                                                                              *string `json:"role-id,omitempty"`
}

// Identifies an individual task for which the containing object is a consequence of.
type TaskReference struct {
	// Used to detail assessment subjects that were identfied by this task.                      
	IdentifiedSubject                                                      *IdentifiedSubject    `json:"identified-subject,omitempty"`
	Links                                                                  []LinkElement         `json:"links,omitempty"`
	Props                                                                  []Property            `json:"props,omitempty"`
	Remarks                                                                *string               `json:"remarks,omitempty"`
	ResponsibleParties                                                     []ResponsibleParty    `json:"responsible-parties,omitempty"`
	Subjects                                                               []SubjectOfAssessment `json:"subjects,omitempty"`
	// A machine-oriented identifier reference to a unique task.                                 
	TaskUUID                                                               string                `json:"task-uuid"`
}

// Used to detail assessment subjects that were identfied by this task.
type IdentifiedSubject struct {
	// A machine-oriented identifier reference to a unique assessment subject placeholder                      
	// defined by this task.                                                                                   
	SubjectPlaceholderUUID                                                               string                `json:"subject-placeholder-uuid"`
	Subjects                                                                             []SubjectOfAssessment `json:"subjects"`
}

// A set of textual statements, typically written by the assessor.
type AttestationStatements struct {
	Parts              []AssessmentPart   `json:"parts"`
	ResponsibleParties []ResponsibleParty `json:"responsible-parties,omitempty"`
}

// Describes an individual finding.
type Finding struct {
	// A human-readable description of this finding.                                                                        
	Description                                                                                 string                      `json:"description"`
	// A machine-oriented identifier reference to the implementation statement in the SSP to                                
	// which this finding is related.                                                                                       
	ImplementationStatementUUID                                                                 *string                     `json:"implementation-statement-uuid,omitempty"`
	Links                                                                                       []LinkElement               `json:"links,omitempty"`
	Origins                                                                                     []FindingOrigin             `json:"origins,omitempty"`
	Props                                                                                       []Property                  `json:"props,omitempty"`
	RelatedObservations                                                                         []FindingRelatedObservation `json:"related-observations,omitempty"`
	RelatedRisks                                                                                []FindingRelatedRisk        `json:"related-risks,omitempty"`
	Remarks                                                                                     *string                     `json:"remarks,omitempty"`
	Target                                                                                      TargetClass                 `json:"target"`
	// The title for this finding.                                                                                          
	Title                                                                                       string                      `json:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                            
	// to reference this finding in this or other OSCAL instances. The locally defined UUID of                              
	// the finding can be used to reference the data item locally or globally (e.g., in an                                  
	// imported OSCAL instance). This UUID should be assigned per-subject, which means it should                            
	// be consistently used to identify the same subject across revisions of the document.                                  
	UUID                                                                                        string                      `json:"uuid"`
}

// Identifies the source of the finding, such as a tool, interviewed person, or activity.
type FindingOrigin struct {
	Actors       []OriginatingActor `json:"actors"`
	RelatedTasks []TaskReference    `json:"related-tasks,omitempty"`
}

// The actor that produces an observation, a finding, or a risk. One or more actor type can
// be used to specify a person that is using a tool.
type OriginatingActor struct {
	// A machine-oriented identifier reference to the tool or person based on the associated               
	// type.                                                                                               
	ActorUUID                                                                                string        `json:"actor-uuid"`
	Links                                                                                    []LinkElement `json:"links,omitempty"`
	Props                                                                                    []Property    `json:"props,omitempty"`
	// For a party, this can optionally be used to specify the role the actor was performing.              
	RoleID                                                                                   *string       `json:"role-id,omitempty"`
	// The kind of actor.                                                                                  
	Type                                                                                     ActorType     `json:"type"`
}

// Relates the finding to a set of referenced observations that were used to determine the
// finding.
type FindingRelatedObservation struct {
	// A machine-oriented identifier reference to an observation defined in the list of       
	// observations.                                                                          
	ObservationUUID                                                                    string `json:"observation-uuid"`
}

// Relates the finding to a set of referenced risks that were used to determine the finding.
type FindingRelatedRisk struct {
	// A machine-oriented identifier reference to a risk defined in the list of risks.       
	RiskUUID                                                                          string `json:"risk-uuid"`
}

// Captures an assessor's conclusions regarding the degree to which an objective is
// satisfied.
type TargetClass struct {
	// A human-readable description of the assessor's conclusions regarding the degree to which                      
	// an objective is satisfied.                                                                                    
	Description                                                                                *string               `json:"description,omitempty"`
	ImplementationStatus                                                                       *ImplementationStatus `json:"implementation-status,omitempty"`
	Links                                                                                      []LinkElement         `json:"links,omitempty"`
	Props                                                                                      []Property            `json:"props,omitempty"`
	Remarks                                                                                    *string               `json:"remarks,omitempty"`
	// A determination of if the objective is satisfied or not within a given system.                                
	Status                                                                                     StatusClass           `json:"status"`
	// A machine-oriented identifier reference for a specific target qualified by the type.                          
	TargetID                                                                                   string                `json:"target-id"`
	// The title for this objective status.                                                                          
	Title                                                                                      *string               `json:"title,omitempty"`
	// Identifies the type of the target.                                                                            
	Type                                                                                       FindingTargetType     `json:"type"`
}

// Indicates the degree to which the a given control is implemented.
type ImplementationStatus struct {
	Remarks                                                                     *string `json:"remarks,omitempty"`
	// Identifies the implementation status of the control or control objective.        
	State                                                                       string  `json:"state"`
}

// A determination of if the objective is satisfied or not within a given system.
type StatusClass struct {
	// The reason the objective was given it's status.                                    
	Reason                                                           *string              `json:"reason,omitempty"`
	Remarks                                                          *string              `json:"remarks,omitempty"`
	// An indication as to whether the objective is satisfied or not.                     
	State                                                            ObjectiveStatusState `json:"state"`
}

// Used to define data objects that are used in the assessment plan, that do not appear in
// the referenced SSP.
type ResultLocalDefinitions struct {
	AssessmentAssets *AssessmentAssets           `json:"assessment-assets,omitempty"`
	Components       []AssessmentAssetsComponent `json:"components,omitempty"`
	InventoryItems   []InventoryItem             `json:"inventory-items,omitempty"`
	Tasks            []Task                      `json:"tasks,omitempty"`
	Users            []SystemUser                `json:"users,omitempty"`
}

// Describes an individual observation.
type Observation struct {
	// Date/time stamp identifying when the finding information was collected.                                         
	Collected                                                                                   time.Time              `json:"collected"`
	// A human-readable description of this assessment observation.                                                    
	Description                                                                                 string                 `json:"description"`
	// Date/time identifying when the finding information is out-of-date and no longer valid.                          
	// Typically used with continuous assessment scenarios.                                                            
	Expires                                                                                     *time.Time             `json:"expires,omitempty"`
	Links                                                                                       []LinkElement          `json:"links,omitempty"`
	Methods                                                                                     []string               `json:"methods"`
	Origins                                                                                     []FindingOrigin        `json:"origins,omitempty"`
	Props                                                                                       []Property             `json:"props,omitempty"`
	RelevantEvidence                                                                            []RelevantEvidence     `json:"relevant-evidence,omitempty"`
	Remarks                                                                                     *string                `json:"remarks,omitempty"`
	Subjects                                                                                    []IdentifiesTheSubject `json:"subjects,omitempty"`
	// The title for this observation.                                                                                 
	Title                                                                                       *string                `json:"title,omitempty"`
	Types                                                                                       []string               `json:"types,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                       
	// to reference this observation elsewhere in this or other OSCAL instances. The locally                           
	// defined UUID of the observation can be used to reference the data item locally or                               
	// globally (e.g., in an imorted OSCAL instance). This UUID should be assigned per-subject,                        
	// which means it should be consistently used to identify the same subject across revisions                        
	// of the document.                                                                                                
	UUID                                                                                        string                 `json:"uuid"`
}

// Links this observation to relevant evidence.
type RelevantEvidence struct {
	// A human-readable description of this evidence.                
	Description                                        string        `json:"description"`
	// A resolvable URL reference to relevant evidence.              
	Href                                               *string       `json:"href,omitempty"`
	Links                                              []LinkElement `json:"links,omitempty"`
	Props                                              []Property    `json:"props,omitempty"`
	Remarks                                            *string       `json:"remarks,omitempty"`
}

// A human-oriented identifier reference to a resource. Use type to indicate whether the
// identified resource is a component, inventory item, location, user, or something else.
type IdentifiesTheSubject struct {
	Links                                                                                      []LinkElement `json:"links,omitempty"`
	Props                                                                                      []Property    `json:"props,omitempty"`
	Remarks                                                                                    *string       `json:"remarks,omitempty"`
	// A machine-oriented identifier reference to a component, inventory-item, location, party,              
	// user, or resource using it's UUID.                                                                    
	SubjectUUID                                                                                string        `json:"subject-uuid"`
	// The title or name for the referenced subject.                                                         
	Title                                                                                      *string       `json:"title,omitempty"`
	// Used to indicate the type of object pointed to by the uuid-ref within a subject.                      
	Type                                                                                       string        `json:"type"`
}

// An identified risk.
type IdentifiedRisk struct {
	Characterizations                                                                           []Characterization       `json:"characterizations,omitempty"`
	// The date/time by which the risk must be resolved.                                                                 
	Deadline                                                                                    *time.Time               `json:"deadline,omitempty"`
	// A human-readable summary of the identified risk, to include a statement of how the risk                           
	// impacts the system.                                                                                               
	Description                                                                                 string                   `json:"description"`
	Links                                                                                       []LinkElement            `json:"links,omitempty"`
	MitigatingFactors                                                                           []MitigatingFactor       `json:"mitigating-factors,omitempty"`
	Origins                                                                                     []FindingOrigin          `json:"origins,omitempty"`
	Props                                                                                       []Property               `json:"props,omitempty"`
	RelatedObservations                                                                         []RiskRelatedObservation `json:"related-observations,omitempty"`
	Remediations                                                                                []RiskResponse           `json:"remediations,omitempty"`
	// A log of all risk-related tasks taken.                                                                            
	RiskLog                                                                                     *RiskLog                 `json:"risk-log,omitempty"`
	// An summary of impact for how the risk affects the system.                                                         
	Statement                                                                                   string                   `json:"statement"`
	Status                                                                                      string                   `json:"status"`
	ThreatIDS                                                                                   []ThreatID               `json:"threat-ids,omitempty"`
	// The title for this risk.                                                                                          
	Title                                                                                       string                   `json:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                         
	// to reference this risk elsewhere in this or other OSCAL instances. The locally defined                            
	// UUID of the risk can be used to reference the data item locally or globally (e.g., in an                          
	// imported OSCAL instance). This UUID should be assigned per-subject, which means it should                         
	// be consistently used to identify the same subject across revisions of the document.                               
	UUID                                                                                        string                   `json:"uuid"`
}

// A collection of descriptive data about the containing object from a specific origin.
type Characterization struct {
	Facets []Facet       `json:"facets"`
	Links  []LinkElement `json:"links,omitempty"`
	Origin FindingOrigin `json:"origin"`
	Props  []Property    `json:"props,omitempty"`
}

// An individual characteristic that is part of a larger set produced by the same actor.
type Facet struct {
	Links                                                                                     []LinkElement `json:"links,omitempty"`
	// The name of the risk metric within the specified system.                                             
	Name                                                                                      string        `json:"name"`
	Props                                                                                     []Property    `json:"props,omitempty"`
	Remarks                                                                                   *string       `json:"remarks,omitempty"`
	// Specifies the naming system under which this risk metric is organized, which allows for              
	// the same names to be used in different systems controlled by different parties. This                 
	// avoids the potential of a name clash.                                                                
	System                                                                                    string        `json:"system"`
	// Indicates the value of the facet.                                                                    
	Value                                                                                     string        `json:"value"`
}

// Describes an existing mitigating factor that may affect the overall determination of the
// risk, with an optional link to an implementation statement in the SSP.
type MitigatingFactor struct {
	// A human-readable description of this mitigating factor.                                                         
	Description                                                                                 string                 `json:"description"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                       
	// to reference this implementation statement elsewhere in this or other OSCAL instancess.                         
	// The locally defined UUID of the implementation statement can be used to reference the                           
	// data item locally or globally (e.g., in an imported OSCAL instance). This UUID should be                        
	// assigned per-subject, which means it should be consistently used to identify the same                           
	// subject across revisions of the document.                                                                       
	ImplementationUUID                                                                          *string                `json:"implementation-uuid,omitempty"`
	Links                                                                                       []LinkElement          `json:"links,omitempty"`
	Props                                                                                       []Property             `json:"props,omitempty"`
	Subjects                                                                                    []IdentifiesTheSubject `json:"subjects,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                       
	// to reference this mitigating factor elsewhere in this or other OSCAL instances. The                             
	// locally defined UUID of the mitigating factor can be used to reference the data item                            
	// locally or globally (e.g., in an imported OSCAL instance). This UUID should be assigned                         
	// per-subject, which means it should be consistently used to identify the same subject                            
	// across revisions of the document.                                                                               
	UUID                                                                                        string                 `json:"uuid"`
}

// Relates the finding to a set of referenced observations that were used to determine the
// finding.
type RiskRelatedObservation struct {
	// A machine-oriented identifier reference to an observation defined in the list of       
	// observations.                                                                          
	ObservationUUID                                                                    string `json:"observation-uuid"`
}

// Describes either recommended or an actual plan for addressing the risk.
type RiskResponse struct {
	// A human-readable description of this response plan.                                                      
	Description                                                                                 string          `json:"description"`
	// Identifies whether this is a recommendation, such as from an assessor or tool, or an                     
	// actual plan accepted by the system owner.                                                                
	Lifecycle                                                                                   string          `json:"lifecycle"`
	Links                                                                                       []LinkElement   `json:"links,omitempty"`
	Origins                                                                                     []FindingOrigin `json:"origins,omitempty"`
	Props                                                                                       []Property      `json:"props,omitempty"`
	Remarks                                                                                     *string         `json:"remarks,omitempty"`
	RequiredAssets                                                                              []RequiredAsset `json:"required-assets,omitempty"`
	Tasks                                                                                       []Task          `json:"tasks,omitempty"`
	// The title for this response activity.                                                                    
	Title                                                                                       string          `json:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                
	// to reference this remediation elsewhere in this or other OSCAL instances. The locally                    
	// defined UUID of the risk response can be used to reference the data item locally or                      
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                
	// which means it should be consistently used to identify the same subject across revisions                 
	// of the document.                                                                                         
	UUID                                                                                        string          `json:"uuid"`
}

// Identifies an asset required to achieve remediation.
type RequiredAsset struct {
	// A human-readable description of this required asset.                                                            
	Description                                                                                 string                 `json:"description"`
	Links                                                                                       []LinkElement          `json:"links,omitempty"`
	Props                                                                                       []Property             `json:"props,omitempty"`
	Remarks                                                                                     *string                `json:"remarks,omitempty"`
	Subjects                                                                                    []IdentifiesTheSubject `json:"subjects,omitempty"`
	// The title for this required asset.                                                                              
	Title                                                                                       *string                `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                       
	// to reference this required asset elsewhere in this or other OSCAL instances. The locally                        
	// defined UUID of the asset can be used to reference the data item locally or globally                            
	// (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject, which                          
	// means it should be consistently used to identify the same subject across revisions of the                       
	// document.                                                                                                       
	UUID                                                                                        string                 `json:"uuid"`
}

// A log of all risk-related tasks taken.
type RiskLog struct {
	Entries []RiskLogEntry `json:"entries"`
}

// Identifies an individual risk response that occurred as part of managing an identified
// risk.
type RiskLogEntry struct {
	// A human-readable description of what was done regarding the risk.                                                
	Description                                                                                 *string                 `json:"description,omitempty"`
	// Identifies the end date and time of the event. If the event is a point in time, the start                        
	// and end will be the same date and time.                                                                          
	End                                                                                         *time.Time              `json:"end,omitempty"`
	Links                                                                                       []LinkElement           `json:"links,omitempty"`
	LoggedBy                                                                                    []LoggedBy              `json:"logged-by,omitempty"`
	Props                                                                                       []Property              `json:"props,omitempty"`
	RelatedResponses                                                                            []RiskResponseReference `json:"related-responses,omitempty"`
	Remarks                                                                                     *string                 `json:"remarks,omitempty"`
	// Identifies the start date and time of the event.                                                                 
	Start                                                                                       time.Time               `json:"start"`
	StatusChange                                                                                *string                 `json:"status-change,omitempty"`
	// The title for this risk log entry.                                                                               
	Title                                                                                       *string                 `json:"title,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                        
	// to reference this risk log entry elsewhere in this or other OSCAL instances. The locally                         
	// defined UUID of the risk log entry can be used to reference the data item locally or                             
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                        
	// which means it should be consistently used to identify the same subject across revisions                         
	// of the document.                                                                                                 
	UUID                                                                                        string                  `json:"uuid"`
}

// Identifies an individual risk response that this log entry is for.
type RiskResponseReference struct {
	Links                                                                []LinkElement   `json:"links,omitempty"`
	Props                                                                []Property      `json:"props,omitempty"`
	RelatedTasks                                                         []TaskReference `json:"related-tasks,omitempty"`
	Remarks                                                              *string         `json:"remarks,omitempty"`
	// A machine-oriented identifier reference to a unique risk response.                
	ResponseUUID                                                         string          `json:"response-uuid"`
}

// A pointer, by ID, to an externally-defined threat.
type ThreatID struct {
	// An optional location for the threat data, from which this ID originates.        
	Href                                                                       *string `json:"href,omitempty"`
	ID                                                                         string  `json:"id"`
	// Specifies the source of the threat information.                                 
	System                                                                     string  `json:"system"`
}

// A structured, organized collection of control information.
type Catalog struct {
	BackMatter                                                               *BackMatter      `json:"back-matter,omitempty"`
	Controls                                                                 []Control        `json:"controls,omitempty"`
	Groups                                                                   []CatalogGroup   `json:"groups,omitempty"`
	Metadata                                                                 DocumentMetadata `json:"metadata"`
	Params                                                                   []Parameter      `json:"params,omitempty"`
	// Provides a globally unique means to identify a given catalog instance.                 
	UUID                                                                     string           `json:"uuid"`
}

// A structured object representing a requirement or guideline, which when implemented will
// reduce an aspect of risk related to an information system and its information.
type Control struct {
	// A textual label that provides a sub-type or characterization of the control.                       
	Class                                                                                   *string       `json:"class,omitempty"`
	Controls                                                                                []Control     `json:"controls,omitempty"`
	// Identifies a control such that it can be referenced in the defining catalog and other              
	// OSCAL instances (e.g., profiles).                                                                  
	ID                                                                                      string        `json:"id"`
	Links                                                                                   []LinkElement `json:"links,omitempty"`
	Params                                                                                  []Parameter   `json:"params,omitempty"`
	Parts                                                                                   []PartElement `json:"parts,omitempty"`
	Props                                                                                   []Property    `json:"props,omitempty"`
	// A name given to the control, which may be used by a tool for display and navigation.               
	Title                                                                                   string        `json:"title"`
}

// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
type Parameter struct {
	// A textual label that provides a characterization of the type, purpose, use or scope of                
	// the parameter.                                                                                        
	Class                                                                                      *string       `json:"class,omitempty"`
	Constraints                                                                                []Constraint  `json:"constraints,omitempty"`
	// (deprecated) Another parameter invoking this one. This construct has been deprecated and              
	// should not be used.                                                                                   
	DependsOn                                                                                  *string       `json:"depends-on,omitempty"`
	Guidelines                                                                                 []Guideline   `json:"guidelines,omitempty"`
	// A unique identifier for the parameter.                                                                
	ID                                                                                         string        `json:"id"`
	// A short, placeholder name for the parameter, which can be used as a substitute for a                  
	// value if no value is assigned.                                                                        
	Label                                                                                      *string       `json:"label,omitempty"`
	Links                                                                                      []LinkElement `json:"links,omitempty"`
	Props                                                                                      []Property    `json:"props,omitempty"`
	Remarks                                                                                    *string       `json:"remarks,omitempty"`
	Select                                                                                     *Selection    `json:"select,omitempty"`
	// Describes the purpose and use of a parameter.                                                         
	Usage                                                                                      *string       `json:"usage,omitempty"`
	Values                                                                                     []string      `json:"values,omitempty"`
}

// A formal or informal expression of a constraint or test.
type Constraint struct {
	// A textual summary of the constraint to be applied.                 
	Description                                          *string          `json:"description,omitempty"`
	Tests                                                []ConstraintTest `json:"tests,omitempty"`
}

// A test expression which is expected to be evaluated by a tool.
type ConstraintTest struct {
	// A formal (executable) expression of a constraint.        
	Expression                                          string  `json:"expression"`
	Remarks                                             *string `json:"remarks,omitempty"`
}

// A prose statement that provides a recommendation for the use of a parameter.
type Guideline struct {
	// Prose permits multiple paragraphs, lists, tables etc.       
	Prose                                                   string `json:"prose"`
}

// Presenting a choice among alternatives.
type Selection struct {
	Choice                                                                                     []string              `json:"choice,omitempty"`
	// Describes the number of selections that must occur. Without this setting, only one value                      
	// should be assumed to be permitted.                                                                            
	HowMany                                                                                    *ParameterCardinality `json:"how-many,omitempty"`
}

// A group of controls, or of groups of controls.
type CatalogGroup struct {
	// A textual label that provides a sub-type or characterization of the group.                          
	Class                                                                                   *string        `json:"class,omitempty"`
	Controls                                                                                []Control      `json:"controls,omitempty"`
	Groups                                                                                  []CatalogGroup `json:"groups,omitempty"`
	// Identifies the group for the purpose of cross-linking within the defining instance or               
	// from other instances that reference the catalog.                                                    
	ID                                                                                      *string        `json:"id,omitempty"`
	Links                                                                                   []LinkElement  `json:"links,omitempty"`
	Params                                                                                  []Parameter    `json:"params,omitempty"`
	Parts                                                                                   []PartElement  `json:"parts,omitempty"`
	Props                                                                                   []Property     `json:"props,omitempty"`
	// A name given to the group, which may be used by a tool for display and navigation.                  
	Title                                                                                   string         `json:"title"`
}

// A collection of component descriptions, which may optionally be grouped by capability.
type ComponentDefinition struct {
	BackMatter                                                                            *BackMatter                    `json:"back-matter,omitempty"`
	Capabilities                                                                          []Capability                   `json:"capabilities,omitempty"`
	Components                                                                            []ComponentDefinitionComponent `json:"components,omitempty"`
	ImportComponentDefinitions                                                            []ImportComponentDefinition    `json:"import-component-definitions,omitempty"`
	Metadata                                                                              DocumentMetadata               `json:"metadata"`
	// Provides a globally unique means to identify a given component definition instance.                               
	UUID                                                                                  string                         `json:"uuid"`
}

// A grouping of other components and/or capabilities.
type Capability struct {
	ControlImplementations                                             []ControlImplementationSet `json:"control-implementations,omitempty"`
	// A summary of the capability.                                                               
	Description                                                        string                     `json:"description"`
	IncorporatesComponents                                             []IncorporatesComponent    `json:"incorporates-components,omitempty"`
	Links                                                              []LinkElement              `json:"links,omitempty"`
	// The capability's human-readable name.                                                      
	Name                                                               string                     `json:"name"`
	Props                                                              []Property                 `json:"props,omitempty"`
	Remarks                                                            *string                    `json:"remarks,omitempty"`
	// Provides a globally unique means to identify a given capability.                           
	UUID                                                               string                     `json:"uuid"`
}

// Defines how the component or capability supports a set of controls.
type ControlImplementationSet struct {
	// A description of how the specified set of controls are implemented for the containing                                    
	// component or capability.                                                                                                 
	Description                                                                                 string                          `json:"description"`
	ImplementedRequirements                                                                     []ImplementedRequirementElement `json:"implemented-requirements"`
	Links                                                                                       []LinkElement                   `json:"links,omitempty"`
	Props                                                                                       []Property                      `json:"props,omitempty"`
	SetParameters                                                                               []SetParameterValue             `json:"set-parameters,omitempty"`
	// A reference to an OSCAL catalog or profile providing the referenced control or subcontrol                                
	// definition.                                                                                                              
	Source                                                                                      string                          `json:"source"`
	// Provides a means to identify a set of control implementations that are supported by a                                    
	// given component or capability.                                                                                           
	UUID                                                                                        string                          `json:"uuid"`
}

// Describes how the containing component or capability implements an individual control.
type ImplementedRequirementElement struct {
	// A reference to a control with a corresponding id value. When referencing an externally                                    
	// defined control, the Control Identifier Reference must be used in the context of the                                      
	// external / imported OSCAL instance (e.g., uri-reference).                                                                 
	ControlID                                                                                   string                           `json:"control-id"`
	// A suggestion from the supplier (e.g., component vendor or author) for how the specified                                   
	// control may be implemented if the containing component or capability is instantiated in a                                 
	// system security plan.                                                                                                     
	Description                                                                                 string                           `json:"description"`
	Links                                                                                       []LinkElement                    `json:"links,omitempty"`
	Props                                                                                       []Property                       `json:"props,omitempty"`
	Remarks                                                                                     *string                          `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole                `json:"responsible-roles,omitempty"`
	SetParameters                                                                               []SetParameterValue              `json:"set-parameters,omitempty"`
	Statements                                                                                  []ControlStatementImplementation `json:"statements,omitempty"`
	// Provides a globally unique means to identify a given control implementation by a                                          
	// component.                                                                                                                
	UUID                                                                                        string                           `json:"uuid"`
}

// Identifies the parameter that will be set by the enclosed value.
type SetParameterValue struct {
	// A human-oriented reference to a parameter within a control, who's catalog has been         
	// imported into the current implementation context.                                          
	ParamID                                                                              string   `json:"param-id"`
	Remarks                                                                              *string  `json:"remarks,omitempty"`
	Values                                                                               []string `json:"values"`
}

// Identifies which statements within a control are addressed.
type ControlStatementImplementation struct {
	// A summary of how the containing control statement is implemented by the component or                       
	// capability.                                                                                                
	Description                                                                                 string            `json:"description"`
	Links                                                                                       []LinkElement     `json:"links,omitempty"`
	Props                                                                                       []Property        `json:"props,omitempty"`
	Remarks                                                                                     *string           `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole `json:"responsible-roles,omitempty"`
	// A human-oriented identifier reference to a control statement.                                              
	StatementID                                                                                 string            `json:"statement-id"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                  
	// to reference this control statement elsewhere in this or other OSCAL instances. The UUID                   
	// of the control statement in the source OSCAL instance is sufficient to reference the data                  
	// item locally or globally (e.g., in an imported OSCAL instance).                                            
	UUID                                                                                        string            `json:"uuid"`
}

// The collection of components comprising this capability.
type IncorporatesComponent struct {
	// A machine-oriented identifier reference to a component.                         
	ComponentUUID                                                               string `json:"component-uuid"`
	// A description of the component, including information about its function.       
	Description                                                                 string `json:"description"`
}

// A defined component that can be part of an implemented system.
type ComponentDefinitionComponent struct {
	ControlImplementations                                                      []ControlImplementationSet   `json:"control-implementations,omitempty"`
	// A description of the component, including information about its function.                             
	Description                                                                 string                       `json:"description"`
	Links                                                                       []LinkElement                `json:"links,omitempty"`
	Props                                                                       []Property                   `json:"props,omitempty"`
	Protocols                                                                   []ServiceProtocolInformation `json:"protocols,omitempty"`
	// A summary of the technological or business purpose of the component.                                  
	Purpose                                                                     *string                      `json:"purpose,omitempty"`
	Remarks                                                                     *string                      `json:"remarks,omitempty"`
	ResponsibleRoles                                                            []ResponsibleRole            `json:"responsible-roles,omitempty"`
	// A human readable name for the component.                                                              
	Title                                                                       string                       `json:"title"`
	// A category describing the purpose of the component.                                                   
	Type                                                                        string                       `json:"type"`
	// Provides a globally unique means to identify a given component.                                       
	UUID                                                                        string                       `json:"uuid"`
}

// Loads a component definition from another resource.
type ImportComponentDefinition struct {
	// A link to a resource that defines a set of components and/or capabilities to import into       
	// this collection.                                                                               
	Href                                                                                       string `json:"href"`
}

// A plan of action and milestones which identifies initial and residual risks, deviations,
// and disposition, such as those required by FedRAMP.
type PlanOfActionAndMilestonesPOAM struct {
	BackMatter                                                                              *BackMatter                                `json:"back-matter,omitempty"`
	Findings                                                                                []Finding                                  `json:"findings,omitempty"`
	ImportSSP                                                                               *ImportSystemSecurityPlan                  `json:"import-ssp,omitempty"`
	LocalDefinitions                                                                        *PlanOfActionAndMilestonesLocalDefinitions `json:"local-definitions,omitempty"`
	Metadata                                                                                DocumentMetadata                           `json:"metadata"`
	Observations                                                                            []Observation                              `json:"observations,omitempty"`
	PoamItems                                                                               []POAMItem                                 `json:"poam-items"`
	Risks                                                                                   []IdentifiedRisk                           `json:"risks,omitempty"`
	SystemID                                                                                *SystemIdentification                      `json:"system-id,omitempty"`
	// A machine-oriented, globally unique identifier with instancescope that can be used to                                           
	// reference this POA&M instance in this OSCAL instance. This UUID should be assigned                                              
	// per-subject, which means it should be consistently used to identify the same subject                                            
	// across revisions of the document.                                                                                               
	UUID                                                                                    string                                     `json:"uuid"`
}

// Allows components, and inventory-items to be defined within the POA&M for circumstances
// where no OSCAL-based SSP exists, or is not delivered with the POA&M.
type PlanOfActionAndMilestonesLocalDefinitions struct {
	AssessmentAssets *AssessmentAssets           `json:"assessment-assets,omitempty"`
	Components       []AssessmentAssetsComponent `json:"components,omitempty"`
	InventoryItems   []InventoryItem             `json:"inventory-items,omitempty"`
	Remarks          *string                     `json:"remarks,omitempty"`
}

// Describes an individual POA&M item.
type POAMItem struct {
	// A human-readable description of POA&M item.                                                                        
	Description                                                                              string                       `json:"description"`
	Links                                                                                    []LinkElement                `json:"links,omitempty"`
	Origins                                                                                  []PoamItemOrigin             `json:"origins,omitempty"`
	Props                                                                                    []Property                   `json:"props,omitempty"`
	RelatedFindings                                                                          []RelatedFinding             `json:"related-findings,omitempty"`
	RelatedObservations                                                                      []PoamItemRelatedObservation `json:"related-observations,omitempty"`
	RelatedRisks                                                                             []PoamItemRelatedRisk        `json:"related-risks,omitempty"`
	Remarks                                                                                  *string                      `json:"remarks,omitempty"`
	// The title or name for this POA&M item .                                                                            
	Title                                                                                    string                       `json:"title"`
	// A machine-oriented, globally unique identifier with instance scope that can be used to                             
	// reference this POA&M item entry in this OSCAL instance. This UUID should be assigned                               
	// per-subject, which means it should be consistently used to identify the same subject                               
	// across revisions of the document.                                                                                  
	UUID                                                                                     *string                      `json:"uuid,omitempty"`
}

// Identifies the source of the finding, such as a tool or person.
type PoamItemOrigin struct {
	Actors []OriginatingActor `json:"actors"`
}

// Relates the poam-item to referenced finding(s).
type RelatedFinding struct {
	// A machine-oriented identifier reference to a finding defined in the list of findings.       
	FindingUUID                                                                             string `json:"finding-uuid"`
}

// Relates the poam-item to a set of referenced observations that were used to determine the
// finding.
type PoamItemRelatedObservation struct {
	// A machine-oriented identifier reference to an observation defined in the list of       
	// observations.                                                                          
	ObservationUUID                                                                    string `json:"observation-uuid"`
}

// Relates the finding to a set of referenced risks that were used to determine the finding.
type PoamItemRelatedRisk struct {
	// A machine-oriented identifier reference to a risk defined in the list of risks.       
	RiskUUID                                                                          string `json:"risk-uuid"`
}

// A human-oriented, globally unique identifier with cross-instance scope that can be used
// to reference this system identification property elsewhere in this or other OSCAL
// instances. When referencing an externally defined system identification, the system
// identification must be used in the context of the external / imported OSCAL instance
// (e.g., uri-reference). This string should be assigned per-subject, which means it should
// be consistently used to identify the same system across revisions of the document.
type SystemIdentification struct {
	ID                                                                                      string  `json:"id"`
	// Identifies the identification system from which the provided identifier was assigned.        
	IdentifierType                                                                          *string `json:"identifier-type,omitempty"`
}

// Each OSCAL profile is defined by a profile element.
type Profile struct {
	BackMatter                                                               *BackMatter      `json:"back-matter,omitempty"`
	Imports                                                                  []ImportResource `json:"imports"`
	Merge                                                                    *MergeControls   `json:"merge,omitempty"`
	Metadata                                                                 DocumentMetadata `json:"metadata"`
	Modify                                                                   *ModifyControls  `json:"modify,omitempty"`
	// Provides a globally unique means to identify a given profile instance.                 
	UUID                                                                     string           `json:"uuid"`
}

// Designates a referenced source catalog or profile that provides a source of control
// information for use in creating a new overlay or baseline.
type ImportResource struct {
	ExcludeControls                                                                             []ImportExcludeControl `json:"exclude-controls,omitempty"`
	// A resolvable URL reference to the base catalog or profile that this profile is tailoring.                       
	Href                                                                                        string                 `json:"href"`
	IncludeAll                                                                                  *IncludeAll            `json:"include-all,omitempty"`
	IncludeControls                                                                             []ImportExcludeControl `json:"include-controls,omitempty"`
}

// Select a control or controls from an imported control set.
type ImportExcludeControl struct {
	Matching                                                                                []MatchControlsByPattern             `json:"matching,omitempty"`
	// When a control is included, whether its child (dependent) controls are also included.                                     
	WithChildControls                                                                       *IncludeContainedControlsWithControl `json:"with-child-controls,omitempty"`
	WithIDS                                                                                 []string                             `json:"with-ids,omitempty"`
}

// Selecting a set of controls by matching their IDs with a wildcard pattern.
type MatchControlsByPattern struct {
	// A glob expression matching the IDs of one or more controls to be selected.        
	Pattern                                                                      *string `json:"pattern,omitempty"`
}

// Provides structuring directives that instruct how controls are organized after profile
// resolution.
type MergeControls struct {
	// Indicates that the controls selected should retain their original grouping as defined in                     
	// the import source.                                                                                           
	AsIs                                                                                       *bool                `json:"as-is,omitempty"`
	// A Combine element defines how to resolve duplicate instances of the same control (e.g.,                      
	// controls with the same ID).                                                                                  
	Combine                                                                                    *CombinationRule     `json:"combine,omitempty"`
	// Provides an alternate grouping structure that selected controls will be placed in.                           
	Custom                                                                                     *CustomGrouping      `json:"custom,omitempty"`
	// Directs that controls appear without any grouping structure.                                                 
	Flat                                                                                       *FlatWithoutGrouping `json:"flat,omitempty"`
}

// A Combine element defines how to resolve duplicate instances of the same control (e.g.,
// controls with the same ID).
type CombinationRule struct {
	// Declare how clashing controls should be handled.                   
	Method                                             *CombinationMethod `json:"method,omitempty"`
}

// Provides an alternate grouping structure that selected controls will be placed in.
type CustomGrouping struct {
	Groups         []CustomGroup    `json:"groups,omitempty"`
	InsertControls []InsertControls `json:"insert-controls,omitempty"`
}

// A group of (selected) controls or of groups of controls.
type CustomGroup struct {
	// A textual label that provides a sub-type or characterization of the group.                 
	Class                                                                        *string          `json:"class,omitempty"`
	Groups                                                                       []CustomGroup    `json:"groups,omitempty"`
	// Identifies the group.                                                                      
	ID                                                                           *string          `json:"id,omitempty"`
	InsertControls                                                               []InsertControls `json:"insert-controls,omitempty"`
	Links                                                                        []LinkElement    `json:"links,omitempty"`
	Params                                                                       []Parameter      `json:"params,omitempty"`
	Parts                                                                        []PartElement    `json:"parts,omitempty"`
	Props                                                                        []Property       `json:"props,omitempty"`
	// A name to be given to the group for use in display.                                        
	Title                                                                        string           `json:"title"`
}

// Specifies which controls to use in the containing context.
type InsertControls struct {
	ExcludeControls                                                               []ImportExcludeControl `json:"exclude-controls,omitempty"`
	IncludeAll                                                                    *IncludeAll            `json:"include-all,omitempty"`
	IncludeControls                                                               []ImportExcludeControl `json:"include-controls,omitempty"`
	// A designation of how a selection of controls in a profile is to be ordered.                       
	Order                                                                         *Order                 `json:"order,omitempty"`
}

// Directs that controls appear without any grouping structure.
type FlatWithoutGrouping struct {
}

// Set parameters or amend controls in resolution.
type ModifyControls struct {
	Alters        []Alteration       `json:"alters,omitempty"`
	SetParameters []ParameterSetting `json:"set-parameters,omitempty"`
}

// Specifies changes to be made to an included control when a profile is resolved.
type Alteration struct {
	Adds                                                                                     []Addition `json:"adds,omitempty"`
	// A reference to a control with a corresponding id value. When referencing an externally           
	// defined control, the Control Identifier Reference must be used in the context of the             
	// external / imported OSCAL instance (e.g., uri-reference).                                        
	ControlID                                                                                string     `json:"control-id"`
	Removes                                                                                  []Removal  `json:"removes,omitempty"`
}

// Specifies contents to be added into controls, in resolution.
type Addition struct {
	// Target location of the addition.                                                                    
	ByID                                                                                     *string       `json:"by-id,omitempty"`
	Links                                                                                    []LinkElement `json:"links,omitempty"`
	Params                                                                                   []Parameter   `json:"params,omitempty"`
	Parts                                                                                    []PartElement `json:"parts,omitempty"`
	// Where to add the new content with respect to the targeted element (beside it or inside              
	// it).                                                                                                
	Position                                                                                 *Position     `json:"position,omitempty"`
	Props                                                                                    []Property    `json:"props,omitempty"`
	// A name given to the control, which may be used by a tool for display and navigation.                
	Title                                                                                    *string       `json:"title,omitempty"`
}

// Specifies objects to be removed from a control based on specific aspects of the object
// that must all match.
type Removal struct {
	// Identify items to remove by matching their class.                                                           
	ByClass                                                                                     *string            `json:"by-class,omitempty"`
	// Identify items to remove indicated by their id.                                                             
	ByID                                                                                        *string            `json:"by-id,omitempty"`
	// Identify items to remove by the name of the item's information object name, e.g. title or                   
	// prop.                                                                                                       
	ByItemName                                                                                  *ItemNameReference `json:"by-item-name,omitempty"`
	// Identify items remove by matching their assigned name.                                                      
	ByName                                                                                      *string            `json:"by-name,omitempty"`
	// Identify items to remove by the item's ns, which is the namespace associated with a part,                   
	// or prop.                                                                                                    
	ByNS                                                                                        *string            `json:"by-ns,omitempty"`
}

// A parameter setting, to be propagated to points of insertion.
type ParameterSetting struct {
	// A textual label that provides a characterization of the parameter.                                    
	Class                                                                                      *string       `json:"class,omitempty"`
	Constraints                                                                                []Constraint  `json:"constraints,omitempty"`
	// **(deprecated)** Another parameter invoking this one. This construct has been deprecated              
	// and should not be used.                                                                               
	DependsOn                                                                                  *string       `json:"depends-on,omitempty"`
	Guidelines                                                                                 []Guideline   `json:"guidelines,omitempty"`
	// A short, placeholder name for the parameter, which can be used as a substitute for a                  
	// value if no value is assigned.                                                                        
	Label                                                                                      *string       `json:"label,omitempty"`
	Links                                                                                      []LinkElement `json:"links,omitempty"`
	// An identifier for the parameter.                                                                      
	ParamID                                                                                    string        `json:"param-id"`
	Props                                                                                      []Property    `json:"props,omitempty"`
	Select                                                                                     *Selection    `json:"select,omitempty"`
	// Describes the purpose and use of a parameter.                                                         
	Usage                                                                                      *string       `json:"usage,omitempty"`
	Values                                                                                     []string      `json:"values,omitempty"`
}

// A system security plan, such as those described in NIST SP 800-18.
type SystemSecurityPlanSSP struct {
	BackMatter                                                                                  *BackMatter                `json:"back-matter,omitempty"`
	ControlImplementation                                                                       ControlImplementationClass `json:"control-implementation"`
	ImportProfile                                                                               ImportProfile              `json:"import-profile"`
	Metadata                                                                                    DocumentMetadata           `json:"metadata"`
	SystemCharacteristics                                                                       SystemCharacteristics      `json:"system-characteristics"`
	SystemImplementation                                                                        SystemImplementation       `json:"system-implementation"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                           
	// to reference this system security plan (SSP) elsewhere in this or other OSCAL instances.                            
	// The locally defined UUID of the SSP can be used to reference the data item locally or                               
	// globally (e.g., in an imported OSCAL instance).This UUID should be assigned per-subject,                            
	// which means it should be consistently used to identify the same subject across revisions                            
	// of the document.                                                                                                    
	UUID                                                                                        string                     `json:"uuid"`
}

// Describes how the system satisfies a set of controls.
type ControlImplementationClass struct {
	// A statement describing important things to know about how this set of control                          
	// satisfaction documentation is approached.                                                              
	Description                                                                     string                    `json:"description"`
	ImplementedRequirements                                                         []ControlBasedRequirement `json:"implemented-requirements"`
	SetParameters                                                                   []SetParameterValue       `json:"set-parameters,omitempty"`
}

// Describes how the system satisfies the requirements of an individual control.
type ControlBasedRequirement struct {
	ByComponents                                                                                []ComponentControlImplementation `json:"by-components,omitempty"`
	// A reference to a control with a corresponding id value. When referencing an externally                                    
	// defined control, the Control Identifier Reference must be used in the context of the                                      
	// external / imported OSCAL instance (e.g., uri-reference).                                                                 
	ControlID                                                                                   string                           `json:"control-id"`
	Links                                                                                       []LinkElement                    `json:"links,omitempty"`
	Props                                                                                       []Property                       `json:"props,omitempty"`
	Remarks                                                                                     *string                          `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole                `json:"responsible-roles,omitempty"`
	SetParameters                                                                               []SetParameterValue              `json:"set-parameters,omitempty"`
	Statements                                                                                  []SpecificControlStatement       `json:"statements,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                 
	// to reference this control requirement elsewhere in this or other OSCAL instances. The                                     
	// locally defined UUID of the control requirement can be used to reference the data item                                    
	// locally or globally (e.g., in an imported OSCAL instance). This UUID should be assigned                                   
	// per-subject, which means it should be consistently used to identify the same subject                                      
	// across revisions of the document.                                                                                         
	UUID                                                                                        string                           `json:"uuid"`
}

// Defines how the referenced component implements a set of controls.
type ComponentControlImplementation struct {
	// A machine-oriented identifier reference to the component that is implemeting a given                                                    
	// control.                                                                                                                                
	ComponentUUID                                                                               string                                         `json:"component-uuid"`
	// An implementation statement that describes how a control or a control statement is                                                      
	// implemented within the referenced system component.                                                                                     
	Description                                                                                 string                                         `json:"description"`
	// Identifies content intended for external consumption, such as with leveraged                                                            
	// organizations.                                                                                                                          
	Export                                                                                      *Export                                        `json:"export,omitempty"`
	ImplementationStatus                                                                        *ImplementationStatus                          `json:"implementation-status,omitempty"`
	Inherited                                                                                   []InheritedControlImplementation               `json:"inherited,omitempty"`
	Links                                                                                       []LinkElement                                  `json:"links,omitempty"`
	Props                                                                                       []Property                                     `json:"props,omitempty"`
	Remarks                                                                                     *string                                        `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole                              `json:"responsible-roles,omitempty"`
	Satisfied                                                                                   []SatisfiedControlImplementationResponsibility `json:"satisfied,omitempty"`
	SetParameters                                                                               []SetParameterValue                            `json:"set-parameters,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                               
	// to reference this by-component entry elsewhere in this or other OSCAL instances. The                                                    
	// locally defined UUID of the by-component entry can be used to reference the data item                                                   
	// locally or globally (e.g., in an imported OSCAL instance). This UUID should be assigned                                                 
	// per-subject, which means it should be consistently used to identify the same subject                                                    
	// across revisions of the document.                                                                                                       
	UUID                                                                                        string                                         `json:"uuid"`
}

// Identifies content intended for external consumption, such as with leveraged
// organizations.
type Export struct {
	// An implementation statement that describes the aspects of the control or control                                              
	// statement implementation that can be available to another system leveraging this system.                                      
	Description                                                                                *string                               `json:"description,omitempty"`
	Links                                                                                      []LinkElement                         `json:"links,omitempty"`
	Props                                                                                      []Property                            `json:"props,omitempty"`
	Provided                                                                                   []ProvidedControlImplementation       `json:"provided,omitempty"`
	Remarks                                                                                    *string                               `json:"remarks,omitempty"`
	Responsibilities                                                                           []ControlImplementationResponsibility `json:"responsibilities,omitempty"`
}

// Describes a capability which may be inherited by a leveraging system.
type ProvidedControlImplementation struct {
	// An implementation statement that describes the aspects of the control or control                           
	// statement implementation that can be provided to another system leveraging this system.                    
	Description                                                                                 string            `json:"description"`
	Links                                                                                       []LinkElement     `json:"links,omitempty"`
	Props                                                                                       []Property        `json:"props,omitempty"`
	Remarks                                                                                     *string           `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole `json:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                  
	// to reference this provided entry elsewhere in this or other OSCAL instances. The locally                   
	// defined UUID of the provided entry can be used to reference the data item locally or                       
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                  
	// which means it should be consistently used to identify the same subject across revisions                   
	// of the document.                                                                                           
	UUID                                                                                        string            `json:"uuid"`
}

// Describes a control implementation responsibility imposed on a leveraging system.
type ControlImplementationResponsibility struct {
	// An implementation statement that describes the aspects of the control or control                           
	// statement implementation that a leveraging system must implement to satisfy the control                    
	// provided by a leveraged system.                                                                            
	Description                                                                                 string            `json:"description"`
	Links                                                                                       []LinkElement     `json:"links,omitempty"`
	Props                                                                                       []Property        `json:"props,omitempty"`
	// A machine-oriented identifier reference to an inherited control implementation that a                      
	// leveraging system is inheriting from a leveraged system.                                                   
	ProvidedUUID                                                                                *string           `json:"provided-uuid,omitempty"`
	Remarks                                                                                     *string           `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole `json:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                  
	// to reference this responsibility elsewhere in this or other OSCAL instances. The locally                   
	// defined UUID of the responsibility can be used to reference the data item locally or                       
	// globally (e.g., in an imported OSCAL instance). This UUID should be assigned per-subject,                  
	// which means it should be consistently used to identify the same subject across revisions                   
	// of the document.                                                                                           
	UUID                                                                                        string            `json:"uuid"`
}

// Describes a control implementation inherited by a leveraging system.
type InheritedControlImplementation struct {
	// An implementation statement that describes the aspects of a control or control statement                   
	// implementation that a leveraging system is inheriting from a leveraged system.                             
	Description                                                                                 string            `json:"description"`
	Links                                                                                       []LinkElement     `json:"links,omitempty"`
	Props                                                                                       []Property        `json:"props,omitempty"`
	// A machine-oriented identifier reference to an inherited control implementation that a                      
	// leveraging system is inheriting from a leveraged system.                                                   
	ProvidedUUID                                                                                *string           `json:"provided-uuid,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole `json:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                  
	// to reference this inherited entry elsewhere in this or other OSCAL instances. The locally                  
	// defined UUID of the inherited control implementation can be used to reference the data                     
	// item locally or globally (e.g., in an imported OSCAL instance). This UUID should be                        
	// assigned per-subject, which means it should be consistently used to identify the same                      
	// subject across revisions of the document.                                                                  
	UUID                                                                                        string            `json:"uuid"`
}

// Describes how this system satisfies a responsibility imposed by a leveraged system.
type SatisfiedControlImplementationResponsibility struct {
	// An implementation statement that describes the aspects of a control or control statement                   
	// implementation that a leveraging system is implementing based on a requirement from a                      
	// leveraged system.                                                                                          
	Description                                                                                 string            `json:"description"`
	Links                                                                                       []LinkElement     `json:"links,omitempty"`
	Props                                                                                       []Property        `json:"props,omitempty"`
	Remarks                                                                                     *string           `json:"remarks,omitempty"`
	// A machine-oriented identifier reference to a control implementation that satisfies a                       
	// responsibility imposed by a leveraged system.                                                              
	ResponsibilityUUID                                                                          *string           `json:"responsibility-uuid,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole `json:"responsible-roles,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                  
	// to reference this satisfied control implementation entry elsewhere in this or other OSCAL                  
	// instances. The locally defined UUID of the control implementation can be used to                           
	// reference the data item locally or globally (e.g., in an imported OSCAL instance). This                    
	// UUID should be assigned per-subject, which means it should be consistently used to                         
	// identify the same subject across revisions of the document.                                                
	UUID                                                                                        string            `json:"uuid"`
}

// Identifies which statements within a control are addressed.
type SpecificControlStatement struct {
	ByComponents                                                                                []ComponentControlImplementation `json:"by-components,omitempty"`
	Links                                                                                       []LinkElement                    `json:"links,omitempty"`
	Props                                                                                       []Property                       `json:"props,omitempty"`
	Remarks                                                                                     *string                          `json:"remarks,omitempty"`
	ResponsibleRoles                                                                            []ResponsibleRole                `json:"responsible-roles,omitempty"`
	// A human-oriented identifier reference to a control statement.                                                             
	StatementID                                                                                 string                           `json:"statement-id"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                 
	// to reference this control statement elsewhere in this or other OSCAL instances. The UUID                                  
	// of the control statement in the source OSCAL instance is sufficient to reference the data                                 
	// item locally or globally (e.g., in an imported OSCAL instance).                                                           
	UUID                                                                                        string                           `json:"uuid"`
}

// Used to import the OSCAL profile representing the system's control baseline.
type ImportProfile struct {
	// A resolvable URL reference to the profile or catalog to use as the system's control        
	// baseline.                                                                                  
	Href                                                                                  string  `json:"href"`
	Remarks                                                                               *string `json:"remarks,omitempty"`
}

// Contains the characteristics of the system, such as its name, purpose, and security
// impact level.
type SystemCharacteristics struct {
	AuthorizationBoundary                                                                     AuthorizationBoundary       `json:"authorization-boundary"`
	DataFlow                                                                                  *DataFlow                   `json:"data-flow,omitempty"`
	DateAuthorized                                                                            *string                     `json:"date-authorized,omitempty"`
	// A summary of the system.                                                                                           
	Description                                                                               string                      `json:"description"`
	Links                                                                                     []LinkElement               `json:"links,omitempty"`
	NetworkArchitecture                                                                       *NetworkArchitecture        `json:"network-architecture,omitempty"`
	Props                                                                                     []Property                  `json:"props,omitempty"`
	Remarks                                                                                   *string                     `json:"remarks,omitempty"`
	ResponsibleParties                                                                        []ResponsibleParty          `json:"responsible-parties,omitempty"`
	SecurityImpactLevel                                                                       *SecurityImpactLevel        `json:"security-impact-level,omitempty"`
	// The overall information system sensitivity categorization, such as defined by FIPS-199.                            
	SecuritySensitivityLevel                                                                  *string                     `json:"security-sensitivity-level,omitempty"`
	Status                                                                                    SystemCharacteristicsStatus `json:"status"`
	SystemIDS                                                                                 []SystemIdentification      `json:"system-ids"`
	SystemInformation                                                                         SystemInformation           `json:"system-information"`
	// The full name of the system.                                                                                       
	SystemName                                                                                string                      `json:"system-name"`
	// A short name for the system, such as an acronym, that is suitable for display in a data                            
	// table or summary list.                                                                                             
	SystemNameShort                                                                           *string                     `json:"system-name-short,omitempty"`
}

// A description of this system's authorization boundary, optionally supplemented by
// diagrams that illustrate the authorization boundary.
type AuthorizationBoundary struct {
	// A summary of the system's authorization boundary.              
	Description                                         string        `json:"description"`
	Diagrams                                            []Diagram     `json:"diagrams,omitempty"`
	Links                                               []LinkElement `json:"links,omitempty"`
	Props                                               []Property    `json:"props,omitempty"`
	Remarks                                             *string       `json:"remarks,omitempty"`
}

// A graphic that provides a visual representation the system, or some aspect of it.
type Diagram struct {
	// A brief caption to annotate the diagram.                                                                
	Caption                                                                                      *string       `json:"caption,omitempty"`
	// A summary of the diagram.                                                                               
	Description                                                                                  *string       `json:"description,omitempty"`
	Links                                                                                        []LinkElement `json:"links,omitempty"`
	Props                                                                                        []Property    `json:"props,omitempty"`
	Remarks                                                                                      *string       `json:"remarks,omitempty"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used               
	// to reference this diagram elsewhere in this or other OSCAL instances. The locally defined               
	// UUID of the diagram can be used to reference the data item locally or globally (e.g., in                
	// an imported OSCAL instance). This UUID should be assigned per-subject, which means it                   
	// should be consistently used to identify the same subject across revisions of the document.              
	UUID                                                                                         string        `json:"uuid"`
}

// A description of the logical flow of information within the system and across its
// boundaries, optionally supplemented by diagrams that illustrate these flows.
type DataFlow struct {
	// A summary of the system's data flow.              
	Description                            string        `json:"description"`
	Diagrams                               []Diagram     `json:"diagrams,omitempty"`
	Links                                  []LinkElement `json:"links,omitempty"`
	Props                                  []Property    `json:"props,omitempty"`
	Remarks                                *string       `json:"remarks,omitempty"`
}

// A description of the system's network architecture, optionally supplemented by diagrams
// that illustrate the network architecture.
type NetworkArchitecture struct {
	// A summary of the system's network architecture.              
	Description                                       string        `json:"description"`
	Diagrams                                          []Diagram     `json:"diagrams,omitempty"`
	Links                                             []LinkElement `json:"links,omitempty"`
	Props                                             []Property    `json:"props,omitempty"`
	Remarks                                           *string       `json:"remarks,omitempty"`
}

// The overall level of expected impact resulting from unauthorized disclosure,
// modification, or loss of access to information.
type SecurityImpactLevel struct {
	// A target-level of availability for the system, based on the sensitivity of information          
	// within the system.                                                                              
	SecurityObjectiveAvailability                                                               string `json:"security-objective-availability"`
	// A target-level of confidentiality for the system, based on the sensitivity of information       
	// within the system.                                                                              
	SecurityObjectiveConfidentiality                                                            string `json:"security-objective-confidentiality"`
	// A target-level of integrity for the system, based on the sensitivity of information             
	// within the system.                                                                              
	SecurityObjectiveIntegrity                                                                  string `json:"security-objective-integrity"`
}

// Describes the operational status of the system.
type SystemCharacteristicsStatus struct {
	Remarks                         *string     `json:"remarks,omitempty"`
	// The current operating status.            
	State                           FluffyState `json:"state"`
}

// Contains details about all information types that are stored, processed, or transmitted
// by the system, such as privacy information, and those defined in NIST SP 800-60.
type SystemInformation struct {
	InformationTypes []InformationType `json:"information-types"`
	Links            []LinkElement     `json:"links,omitempty"`
	Props            []Property        `json:"props,omitempty"`
}

// Contains details about one information type that is stored, processed, or transmitted by
// the system, such as privacy information, and those defined in NIST SP 800-60.
type InformationType struct {
	AvailabilityImpact                                                                          *ImpactLevel                    `json:"availability-impact,omitempty"`
	Categorizations                                                                             []InformationTypeCategorization `json:"categorizations,omitempty"`
	ConfidentialityImpact                                                                       *ImpactLevel                    `json:"confidentiality-impact,omitempty"`
	// A summary of how this information type is used within the system.                                                        
	Description                                                                                 string                          `json:"description"`
	IntegrityImpact                                                                             *ImpactLevel                    `json:"integrity-impact,omitempty"`
	Links                                                                                       []LinkElement                   `json:"links,omitempty"`
	Props                                                                                       []Property                      `json:"props,omitempty"`
	// A human readable name for the information type. This title should be meaningful within                                   
	// the context of the system.                                                                                               
	Title                                                                                       string                          `json:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope that can be used                                
	// to reference this information type elsewhere in this or other OSCAL instances. The                                       
	// locally defined UUID of the information type can be used to reference the data item                                      
	// locally or globally (e.g., in an imported OSCAL instance). This UUID should be assigned                                  
	// per-subject, which means it should be consistently used to identify the same subject                                     
	// across revisions of the document.                                                                                        
	UUID                                                                                        *string                         `json:"uuid,omitempty"`
}

// The expected level of impact resulting from the described information.
type ImpactLevel struct {
	AdjustmentJustification *string       `json:"adjustment-justification,omitempty"`
	Base                    string        `json:"base"`
	Links                   []LinkElement `json:"links,omitempty"`
	Props                   []Property    `json:"props,omitempty"`
	Selected                *string       `json:"selected,omitempty"`
}

// A set of information type identifiers qualified by the given identification system used,
// such as NIST SP 800-60.
type InformationTypeCategorization struct {
	InformationTypeIDS                                           []string `json:"information-type-ids,omitempty"`
	// Specifies the information type identification system used.         
	System                                                       string   `json:"system"`
}

// Provides information as to how the system is implemented.
type SystemImplementation struct {
	Components              []AssessmentAssetsComponent `json:"components"`
	InventoryItems          []InventoryItem             `json:"inventory-items,omitempty"`
	LeveragedAuthorizations []LeveragedAuthorization    `json:"leveraged-authorizations,omitempty"`
	Links                   []LinkElement               `json:"links,omitempty"`
	Props                   []Property                  `json:"props,omitempty"`
	Remarks                 *string                     `json:"remarks,omitempty"`
	Users                   []SystemUser                `json:"users"`
}

// A description of another authorized system from which this system inherits capabilities
// that satisfy security requirements. Another term for this concept is a common control
// provider.
type LeveragedAuthorization struct {
	DateAuthorized                                                                              string        `json:"date-authorized"`
	Links                                                                                       []LinkElement `json:"links,omitempty"`
	// A machine-oriented identifier reference to the party that manages the leveraged system.                
	PartyUUID                                                                                   string        `json:"party-uuid"`
	Props                                                                                       []Property    `json:"props,omitempty"`
	Remarks                                                                                     *string       `json:"remarks,omitempty"`
	// A human readable name for the leveraged authorization in the context of the system.                    
	Title                                                                                       string        `json:"title"`
	// A machine-oriented, globally unique identifier with cross-instance scope and can be used               
	// to reference this leveraged authorization elsewhere in this or other OSCAL instances. The              
	// locally defined UUID of the leveraged authorization can be used to reference the data                  
	// item locally or globally (e.g., in an imported OSCAL instance). This UUID should be                    
	// assigned per-subject, which means it should be consistently used to identify the same                  
	// subject across revisions of the document.                                                              
	UUID                                                                                        string        `json:"uuid"`
}

// Indicates the transport type.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type Transport string

const (
	TCP Transport = "TCP"
	UDP Transport = "UDP"
)

// The operational status.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type PurpleState string

const (
	PurpleDisposition      PurpleState = "disposition"
	PurpleOperational      PurpleState = "operational"
	PurpleOther            PurpleState = "other"
	PurpleUnderDevelopment PurpleState = "under-development"
)

// A category describing the kind of party the object describes.
//
// A label that indicates the nature of a resource, as a data serialization or format.
//
// A non-empty string with leading and trailing whitespace disallowed. Whitespace is: U+9,
// U+10, U+32 or [
// ]+
//
// In case where the href points to a back-matter/resource, this value will indicate the URI
// fragment to append to any rlink associated with the resource. This value MUST be URI
// encoded.
//
// Indicates the value of the attribute, characteristic, or quality.
//
// A formal (executable) expression of a constraint.
//
// City, town or geographical region for the mailing address.
//
// The ISO 3166-1 alpha-2 country code for the mailing address.
//
// Postal or ZIP code for mailing address.
//
// State, province or analogous geographical region for a mailing address.
//
// The full name of the party. This is typically the legal name associated with the party.
//
// A short common name, abbreviation, or acronym for the party.
//
// A short common name, abbreviation, or acronym for the role.
//
// A glob expression matching the IDs of one or more controls to be selected.
//
// A parameter value or set of values.
//
// The capability's human-readable name.
//
// The common name of the protocol, which should be the appropriate "service name" from the
// IANA Service Name and Transport Protocol Port Number Registry.
//
// A target-level of availability for the system, based on the sensitivity of information
// within the system.
//
// A target-level of confidentiality for the system, based on the sensitivity of information
// within the system.
//
// A target-level of integrity for the system, based on the sensitivity of information
// within the system.
//
// The overall information system sensitivity categorization, such as defined by FIPS-199.
//
// A human-oriented, globally unique identifier qualified by the given identification system
// used, such as NIST SP 800-60. This identifier has cross-instance scope and can be used to
// reference this system elsewhere in this or other OSCAL instances. This id should be
// assigned per-subject, which means it should be consistently used to identify the same
// subject across revisions of the document.
//
// The full name of the system.
//
// A short name for the system, such as an acronym, that is suitable for display in a data
// table or summary list.
//
// A short common name, abbreviation, or acronym for the user.
//
// Indicates the value of the facet.
//
// A single line of an address.
//
// The OSCAL model version the document was authored against and will conform to as valid.
//
// Used to distinguish a specific revision of an OSCAL document from other previous and
// future versions.
//
// The prescribed base (Confidentiality, Integrity, or Availability) security impact level.
//
// The selected (Confidentiality, Integrity, or Availability) security impact level.
//
// Describes a function performed for a given authorized privilege by this user class.
//
// The digest method by which a hash is derived.
//
// Indicates the type of phone number.
//
// A category describing the purpose of the component.
//
// Identifies how the observation was made.
type PartyType string

const (
	Organization PartyType = "organization"
	Person       PartyType = "person"
)

// The unit of time for the period.
//
// A label that indicates the nature of a resource, as a data serialization or format.
//
// A non-empty string with leading and trailing whitespace disallowed. Whitespace is: U+9,
// U+10, U+32 or [
// ]+
//
// In case where the href points to a back-matter/resource, this value will indicate the URI
// fragment to append to any rlink associated with the resource. This value MUST be URI
// encoded.
//
// Indicates the value of the attribute, characteristic, or quality.
//
// A formal (executable) expression of a constraint.
//
// City, town or geographical region for the mailing address.
//
// The ISO 3166-1 alpha-2 country code for the mailing address.
//
// Postal or ZIP code for mailing address.
//
// State, province or analogous geographical region for a mailing address.
//
// The full name of the party. This is typically the legal name associated with the party.
//
// A short common name, abbreviation, or acronym for the party.
//
// A short common name, abbreviation, or acronym for the role.
//
// A glob expression matching the IDs of one or more controls to be selected.
//
// A parameter value or set of values.
//
// The capability's human-readable name.
//
// The common name of the protocol, which should be the appropriate "service name" from the
// IANA Service Name and Transport Protocol Port Number Registry.
//
// A target-level of availability for the system, based on the sensitivity of information
// within the system.
//
// A target-level of confidentiality for the system, based on the sensitivity of information
// within the system.
//
// A target-level of integrity for the system, based on the sensitivity of information
// within the system.
//
// The overall information system sensitivity categorization, such as defined by FIPS-199.
//
// A human-oriented, globally unique identifier qualified by the given identification system
// used, such as NIST SP 800-60. This identifier has cross-instance scope and can be used to
// reference this system elsewhere in this or other OSCAL instances. This id should be
// assigned per-subject, which means it should be consistently used to identify the same
// subject across revisions of the document.
//
// The full name of the system.
//
// A short name for the system, such as an acronym, that is suitable for display in a data
// table or summary list.
//
// A short common name, abbreviation, or acronym for the user.
//
// Indicates the value of the facet.
//
// A single line of an address.
//
// The OSCAL model version the document was authored against and will conform to as valid.
//
// Used to distinguish a specific revision of an OSCAL document from other previous and
// future versions.
//
// The prescribed base (Confidentiality, Integrity, or Availability) security impact level.
//
// The selected (Confidentiality, Integrity, or Availability) security impact level.
//
// Describes a function performed for a given authorized privilege by this user class.
//
// The digest method by which a hash is derived.
//
// Indicates the type of phone number.
//
// A category describing the purpose of the component.
//
// Identifies how the observation was made.
type TimeUnit string

const (
	Days    TimeUnit = "days"
	Hours   TimeUnit = "hours"
	Minutes TimeUnit = "minutes"
	Months  TimeUnit = "months"
	Seconds TimeUnit = "seconds"
	Years   TimeUnit = "years"
)

// The kind of actor.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type ActorType string

const (
	AssessmentPlatform ActorType = "assessment-platform"
	Party              ActorType = "party"
	Tool               ActorType = "tool"
)

// An indication as to whether the objective is satisfied or not.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type ObjectiveStatusState string

const (
	NotSatisfied ObjectiveStatusState = "not-satisfied"
	Satisfied    ObjectiveStatusState = "satisfied"
)

// Identifies the type of the target.
//
// A label that indicates the nature of a resource, as a data serialization or format.
//
// A non-empty string with leading and trailing whitespace disallowed. Whitespace is: U+9,
// U+10, U+32 or [
// ]+
//
// In case where the href points to a back-matter/resource, this value will indicate the URI
// fragment to append to any rlink associated with the resource. This value MUST be URI
// encoded.
//
// Indicates the value of the attribute, characteristic, or quality.
//
// A formal (executable) expression of a constraint.
//
// City, town or geographical region for the mailing address.
//
// The ISO 3166-1 alpha-2 country code for the mailing address.
//
// Postal or ZIP code for mailing address.
//
// State, province or analogous geographical region for a mailing address.
//
// The full name of the party. This is typically the legal name associated with the party.
//
// A short common name, abbreviation, or acronym for the party.
//
// A short common name, abbreviation, or acronym for the role.
//
// A glob expression matching the IDs of one or more controls to be selected.
//
// A parameter value or set of values.
//
// The capability's human-readable name.
//
// The common name of the protocol, which should be the appropriate "service name" from the
// IANA Service Name and Transport Protocol Port Number Registry.
//
// A target-level of availability for the system, based on the sensitivity of information
// within the system.
//
// A target-level of confidentiality for the system, based on the sensitivity of information
// within the system.
//
// A target-level of integrity for the system, based on the sensitivity of information
// within the system.
//
// The overall information system sensitivity categorization, such as defined by FIPS-199.
//
// A human-oriented, globally unique identifier qualified by the given identification system
// used, such as NIST SP 800-60. This identifier has cross-instance scope and can be used to
// reference this system elsewhere in this or other OSCAL instances. This id should be
// assigned per-subject, which means it should be consistently used to identify the same
// subject across revisions of the document.
//
// The full name of the system.
//
// A short name for the system, such as an acronym, that is suitable for display in a data
// table or summary list.
//
// A short common name, abbreviation, or acronym for the user.
//
// Indicates the value of the facet.
//
// A single line of an address.
//
// The OSCAL model version the document was authored against and will conform to as valid.
//
// Used to distinguish a specific revision of an OSCAL document from other previous and
// future versions.
//
// The prescribed base (Confidentiality, Integrity, or Availability) security impact level.
//
// The selected (Confidentiality, Integrity, or Availability) security impact level.
//
// Describes a function performed for a given authorized privilege by this user class.
//
// The digest method by which a hash is derived.
//
// Indicates the type of phone number.
//
// A category describing the purpose of the component.
//
// Identifies how the observation was made.
type FindingTargetType string

const (
	ObjectiveID FindingTargetType = "objective-id"
	StatementID FindingTargetType = "statement-id"
)

// Describes the number of selections that must occur. Without this setting, only one value
// should be assumed to be permitted.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type ParameterCardinality string

const (
	One       ParameterCardinality = "one"
	OneOrMore ParameterCardinality = "one-or-more"
)

// When a control is included, whether its child (dependent) controls are also included.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type IncludeContainedControlsWithControl string

const (
	No  IncludeContainedControlsWithControl = "no"
	Yes IncludeContainedControlsWithControl = "yes"
)

// Declare how clashing controls should be handled.
//
// A label that indicates the nature of a resource, as a data serialization or format.
//
// A non-empty string with leading and trailing whitespace disallowed. Whitespace is: U+9,
// U+10, U+32 or [
// ]+
//
// In case where the href points to a back-matter/resource, this value will indicate the URI
// fragment to append to any rlink associated with the resource. This value MUST be URI
// encoded.
//
// Indicates the value of the attribute, characteristic, or quality.
//
// A formal (executable) expression of a constraint.
//
// City, town or geographical region for the mailing address.
//
// The ISO 3166-1 alpha-2 country code for the mailing address.
//
// Postal or ZIP code for mailing address.
//
// State, province or analogous geographical region for a mailing address.
//
// The full name of the party. This is typically the legal name associated with the party.
//
// A short common name, abbreviation, or acronym for the party.
//
// A short common name, abbreviation, or acronym for the role.
//
// A glob expression matching the IDs of one or more controls to be selected.
//
// A parameter value or set of values.
//
// The capability's human-readable name.
//
// The common name of the protocol, which should be the appropriate "service name" from the
// IANA Service Name and Transport Protocol Port Number Registry.
//
// A target-level of availability for the system, based on the sensitivity of information
// within the system.
//
// A target-level of confidentiality for the system, based on the sensitivity of information
// within the system.
//
// A target-level of integrity for the system, based on the sensitivity of information
// within the system.
//
// The overall information system sensitivity categorization, such as defined by FIPS-199.
//
// A human-oriented, globally unique identifier qualified by the given identification system
// used, such as NIST SP 800-60. This identifier has cross-instance scope and can be used to
// reference this system elsewhere in this or other OSCAL instances. This id should be
// assigned per-subject, which means it should be consistently used to identify the same
// subject across revisions of the document.
//
// The full name of the system.
//
// A short name for the system, such as an acronym, that is suitable for display in a data
// table or summary list.
//
// A short common name, abbreviation, or acronym for the user.
//
// Indicates the value of the facet.
//
// A single line of an address.
//
// The OSCAL model version the document was authored against and will conform to as valid.
//
// Used to distinguish a specific revision of an OSCAL document from other previous and
// future versions.
//
// The prescribed base (Confidentiality, Integrity, or Availability) security impact level.
//
// The selected (Confidentiality, Integrity, or Availability) security impact level.
//
// Describes a function performed for a given authorized privilege by this user class.
//
// The digest method by which a hash is derived.
//
// Indicates the type of phone number.
//
// A category describing the purpose of the component.
//
// Identifies how the observation was made.
type CombinationMethod string

const (
	CombinationMethodKeep CombinationMethod = "keep"
	Merge                 CombinationMethod = "merge"
	UseFirst              CombinationMethod = "use-first"
)

// A designation of how a selection of controls in a profile is to be ordered.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type Order string

const (
	Ascending  Order = "ascending"
	Descending Order = "descending"
	OrderKeep  Order = "keep"
)

// Where to add the new content with respect to the targeted element (beside it or inside
// it).
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type Position string

const (
	After    Position = "after"
	Before   Position = "before"
	Ending   Position = "ending"
	Starting Position = "starting"
)

// Identify items to remove by the name of the item's information object name, e.g. title or
// prop.
//
// Name of the file before it was encoded as Base64 to be embedded in a resource. This is
// the name that will be assigned to the file when the file is decoded.
//
// A non-colonized name as defined by XML Schema Part 2: Datatypes Second Edition.
// https://www.w3.org/TR/xmlschema11-2/#NCName.
//
// A textual label that provides a sub-type or characterization of the property's name.
//
// An identifier for relating distinct sets of properties.
//
// A textual label, within a namespace, that uniquely identifies a specific attribute,
// characteristic, or quality of the property's containing object.
//
// A textual label that provides a sub-type or characterization of the control.
//
// Identifies a control such that it can be referenced in the defining catalog and other
// OSCAL instances (e.g., profiles).
//
// A textual label that provides a characterization of the type, purpose, use or scope of
// the parameter.
//
// (deprecated) Another parameter invoking this one. This construct has been deprecated and
// should not be used.
//
// A unique identifier for the parameter.
//
// An optional textual providing a sub-type or characterization of the part's name, or a
// category to which the part belongs.
//
// A unique identifier for the part.
//
// A textual label that uniquely identifies the part's semantic type, which exists in a
// value space qualified by the ns.
//
// A textual label that provides a sub-type or characterization of the group.
//
// Identifies the group for the purpose of cross-linking within the defining instance or
// from other instances that reference the catalog.
//
// A reference to a role performed by a party.
//
// The type of action documented by the assembly, such as an approval.
//
// A unique identifier for the role.
//
// Identifies the group.
//
// Target location of the addition.
//
// A reference to a control with a corresponding id value. When referencing an externally
// defined control, the Control Identifier Reference must be used in the context of the
// external / imported OSCAL instance (e.g., uri-reference).
//
// Identify items to remove by matching their class.
//
// Identify items to remove indicated by their id.
//
// Identify items remove by matching their assigned name.
//
// Identify items to remove by the item's ns, which is the namespace associated with a part,
// or prop.
//
// A textual label that provides a characterization of the parameter.
//
// **(deprecated)** Another parameter invoking this one. This construct has been deprecated
// and should not be used.
//
// An identifier for the parameter.
//
// A human-oriented identifier reference to a role performed.
//
// A human-oriented reference to a parameter within a control, who's catalog has been
// imported into the current implementation context.
//
// A human-oriented identifier reference to a control statement.
//
// Points to an assessment objective.
//
// Used to constrain the selection to only specificity identified statements.
//
// A textual label that provides a sub-type or characterization of the part's name. This can
// be used to further distinguish or discriminate between the semantics of multiple parts of
// the same control with the same name and ns.
//
// A point to the role-id of the role in which the party is making the log entry.
//
// For a party, this can optionally be used to specify the role the actor was performing.
//
// A machine-oriented identifier reference for a specific target qualified by the type.
//
// The name of the risk metric within the specified system.
//
// Selecting a control by its ID given as a literal.
//
// Reference to a role by UUID.
//
// Describes the type of relationship provided by the link's hypertext reference. This can
// be an indicator of the link's purpose.
//
// Indicates the type of address.
//
// Identifies the implementation status of the control or control objective.
//
// Used to indicate the type of object pointed to by the uuid-ref within a subject.
//
// Indicates the type of assessment subject, such as a component, inventory, item, location,
// or party represented by this selection statement.
//
// The type of task.
//
// A textual label that uniquely identifies the part's semantic type.
//
// The reason the objective was given it's status.
//
// Identifies the nature of the observation. More than one may be used to further qualify
// and enable filtering.
//
// Identifies whether this is a recommendation, such as from an assessor or tool, or an
// actual plan accepted by the system owner.
//
// Describes the status of the associated risk.
type ItemNameReference string

const (
	Link    ItemNameReference = "link"
	Map     ItemNameReference = "map"
	Mapping ItemNameReference = "mapping"
	Param   ItemNameReference = "param"
	Part    ItemNameReference = "part"
	Prop    ItemNameReference = "prop"
)

// The current operating status.
//
// A label that indicates the nature of a resource, as a data serialization or format.
//
// A non-empty string with leading and trailing whitespace disallowed. Whitespace is: U+9,
// U+10, U+32 or [
// ]+
//
// In case where the href points to a back-matter/resource, this value will indicate the URI
// fragment to append to any rlink associated with the resource. This value MUST be URI
// encoded.
//
// Indicates the value of the attribute, characteristic, or quality.
//
// A formal (executable) expression of a constraint.
//
// City, town or geographical region for the mailing address.
//
// The ISO 3166-1 alpha-2 country code for the mailing address.
//
// Postal or ZIP code for mailing address.
//
// State, province or analogous geographical region for a mailing address.
//
// The full name of the party. This is typically the legal name associated with the party.
//
// A short common name, abbreviation, or acronym for the party.
//
// A short common name, abbreviation, or acronym for the role.
//
// A glob expression matching the IDs of one or more controls to be selected.
//
// A parameter value or set of values.
//
// The capability's human-readable name.
//
// The common name of the protocol, which should be the appropriate "service name" from the
// IANA Service Name and Transport Protocol Port Number Registry.
//
// A target-level of availability for the system, based on the sensitivity of information
// within the system.
//
// A target-level of confidentiality for the system, based on the sensitivity of information
// within the system.
//
// A target-level of integrity for the system, based on the sensitivity of information
// within the system.
//
// The overall information system sensitivity categorization, such as defined by FIPS-199.
//
// A human-oriented, globally unique identifier qualified by the given identification system
// used, such as NIST SP 800-60. This identifier has cross-instance scope and can be used to
// reference this system elsewhere in this or other OSCAL instances. This id should be
// assigned per-subject, which means it should be consistently used to identify the same
// subject across revisions of the document.
//
// The full name of the system.
//
// A short name for the system, such as an acronym, that is suitable for display in a data
// table or summary list.
//
// A short common name, abbreviation, or acronym for the user.
//
// Indicates the value of the facet.
//
// A single line of an address.
//
// The OSCAL model version the document was authored against and will conform to as valid.
//
// Used to distinguish a specific revision of an OSCAL document from other previous and
// future versions.
//
// The prescribed base (Confidentiality, Integrity, or Availability) security impact level.
//
// The selected (Confidentiality, Integrity, or Availability) security impact level.
//
// Describes a function performed for a given authorized privilege by this user class.
//
// The digest method by which a hash is derived.
//
// Indicates the type of phone number.
//
// A category describing the purpose of the component.
//
// Identifies how the observation was made.
type FluffyState string

const (
	FluffyDisposition      FluffyState = "disposition"
	FluffyOperational      FluffyState = "operational"
	FluffyOther            FluffyState = "other"
	FluffyUnderDevelopment FluffyState = "under-development"
	UnderMajorModification FluffyState = "under-major-modification"
)
