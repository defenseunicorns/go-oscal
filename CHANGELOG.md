# Change Log
All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).
 
## [Unreleased] - yyyy-mm-dd

### Added
- `doctor` command to prep new oscal_complete_schema.json from [OSCAL Releases](https://github.com/usnistgov/OSCAL/releases/tag/v1.1.2) for use in the `validate` and `revise` commands. Behavior ported from [oscal-json-doctor](https://github.com/defenseunicorns/oscal-json-doctor) in order to keep all needed functionality within the go-oscal repo. 
- documentation for `generate` and `doctor` commands
- documentation for upgrading to a new version of OSCAL (docs/upgrading-oscal-version.md)

### Changed
- [breaking] Types now belong to unique packages in the format `oscalTypes_X_X_X` (ie: `oscalTypes_1_0_4`)
- Now uses the schema information to generate types
- update `OscalModels` is now an alias for `OscalCompleteSchema`
- update makefile to include `upgrade`, `generate-latest-schema`, and `doctor-latest-schema` methods
- remove make commands related to generating types for individual schemas in favor of using complete schemas only for consistency.
- moved `src/pkg/validation/schema` to `src/internal/schemas`
- renamed `src/internal/oscal` to `src/internal/generate` for consistency. 

### Breaking Types Changes
#### oscal types versions: 1.0.4, 1.0.5, 1.0.6, 1.1.0, 1.1.1

- update `AssessmentPlan.TermsAndConditions` is now of type `AssessmentPlanTermsAndConditions` per the schema `title`
  - Renamed from `TermsAndConditions`

- create `PlanOfActionAndMilestonesLocalDefinitions struct`
  - update `PlanOfActionAndMilestones.LocalDefinitions` is now of type `PlanOfActionAndMilestonesLocalDefinitions`

- rename `LocalDefinitions1` to `LocalDefinitions`
  - fix `AssessmentPlan`, `AssessmentResults`, and `Results` now referencing the correct definition for their `LocalDefinitions` field. Previously were referencing the `PlanOfActionAndMilestonesLocalDefinitions` definition.

- rename `Timing struct` to `EventTiming struct` per the schema `title`. 
  - update `Task.Timing` to be of type `EventTiming`

- update `Capability.ControlImplementations` and `DefinedComponent.ControlImplementations` to be of type `[]ControlImplementationSet`
  - create `ControlImplementationSet struct`

- update `PoamItem.Origins` to be of type `[]PoamItemOrigin`
  - rename `Origins1` to `PoamItemOrigin`
  - fix `PoamItem.Origins` type that was previously linked to the incorrect definition

- update `Merge.Combine` to be of type `CombinationRule`
  - rename `Combine` to `CombinationRule` per schema title

- update `Merge.Custom` to be of type `CustomGrouping`
  - rename `Custom` to `CustomGrouping` per schema title

- rename `AssemblyOscalSspControlImplementation` to `ControlImplementation`
  - fix `SystemSecurityPlan.ControlImplementation` now referencing the correct definition

- update `SystemComponent.Status` is now of type `SystemComponentStatus`
  - create `SystemComponentStatus`

- fix `IncludeAll struct` is now a type alias for `map[string]interface{}` in order to support marshaling and field access for what is an `anonymous` schema object 

- rename `AtFrequency` to `FrequencyCondition` per schema title

- rename `OnDate` to `OnDateCondition` per schema title

- rename `WithinDateRange` to `OnDateRangeCondition` per schema title

- update `ImplementedRequirement struct` with the previous definition for `AssemblyOscalSspImplementedRequirement` which was previously unlinked
  - create `ImplementedRequirementControlImplementation struct` with the definition previously associated with `ImplementedRequirement` which is now linked to `ControlImplementationSet` used by `Catalog` and `DefinedComponent`

- update `FindingTarget.Status` to be of type `ObjectiveStatus`
  - fix incorrect definition for `FindingTarget.Status` that had missing field `Reason`
  - create `ObjectiveStatus`

- Rename `Statement` to `ControlStatementImplementation`
  - rename `AssemblyOscalSspStatement` to `Statement`
  - fix `ImplementedRequirement.Statements` not being of the correct definition

- create `RiskResponseReference struct` used by previously missing `RiskLogEntry` type

- update `RiskLog.Entries` to be of type `RiskLogEntry`
  - create `RiskLogEntry struct` type.
  - fix `RiskLog.Entries` was previously linked to the wrong definition of `AssessmentLogEntry` which differs from the `RiskLogEntry`

#### oscal types versions: 1.0.4, 1.0.5, 1.0.6

- rename `AssemblyOscalProfileSelectControlById` to `SelectControlById`
  - fix `Import` and `InsertControls` `IncludeControls` and `ExcludeControls` now referencing correct definition

- update `AssessedControls` fields `IncludeControls` and `ExcludeControls` to be of type `SelectControl` 
 - create `SelectControl struct` which has the definition previously assigned to `SelectControlById`

- fix `Flat struct` is now a type alias for `map[string]interface{}` in order to support marshaling and field access for what is an `anonymous` schema object 

- rename `AssemblyOscalProfileGroup` to `ControlGroup`
  - fix improper links to `Group` (should have been `AssemblyOscalProfileGroup` previously, now `ControlGroup`) in `CustomGrouping` (previously `Custom`) and `Merge` structs.

- rename `AvailabilityImpact` to `AvailabilityImpactLevel` per schema title
  - update `InformationType.AvailabilityImpact` to be of type `AvailabilityImpactLevel`

- rename `ConfidentialityImpact` to `ConfidentialityImpactLevel` per schema title
  - update `InformationType.ConfidentialityImpact` to be of type `ConfidentialityImpactLevel`

- rename `IntegrityImpact` to `IntegrityImpactLevel` per schema title
  - update `InformationType.IntegrityImpact` to be of type `IntegrityImpactLevel`

- fix `CustomGrouping.Groups` (previously `Custom.Groups`) to be of type `ControlGroup` as the definition differs from that of previously declared `Group struct`

#### oscal types versions: 1.1.0, 1.1.1

- rename `Flat` to `FlatWithoutGroup` due to schema title change
  - update `FlatWithoutGroup` is now a type alias for `map[string]interface{}` in order to support marshaling and field access for what is an `anonymous` schema object
  - update `Merge.Flat` to be of type `FlatWithoutGrouping`
  
- rename `SelectControlById` to `AssessedControlSelectControlById`
  - update `AssessedControls` fields `IncludeControls` and `ExcludeControls` to be of type `AssessedControlsSelectControlById` 
  
- rename `AssemblyOscalProfileSelectControlById` to `SelectControlById`
  - fix `Import` and `InsertControls` `IncludeControls` and `ExcludeControls` now referencing correct definition
  
- rename `AssemblyOscalProfileGroup` to `CustomGroupingGroup`
  - fix improper links to `Group` (should have been `AssemblyOscalProfileGroup` previously, now `CustomGroupingGroup`) in `CustomGrouping` (previously `Custom`) and `Merge` structs.
  
- Fix `CustomGrouping.Groups` (previously `Custom.Groups`) to be of type `CustomGroupingGroup` as the definition differs from that of previously declared `Group struct`