## Differences
- Types now belong to unique packages in the format `oscalTypes_X_X_X` (ie: `oscalTypes_1_0_4`)
- Now uses the schema information to generate types
- update `OscalModels` is now an alias for `OscalCompleteSchema`

### 1.0.4, 1.0.5, 1.0.6, 1.1.0, 1.1.1

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
  - fixes `PoamItem.Origins` type that was previously linked to the incorrect definition

- update `Merge.Combine` to be of type `CombinationRule`
  - rename `Combine` to `CombinationRule` per schema title

- update `Merge.Custom` to be of type `CustomGrouping`
  - rename `Custom` to `CustomGrouping` per schema title

- rename `AssemblyOscalSspControlImplementation` to `ControlImplementation`
  - fixes `SystemSecurityPlan.ControlImplementation` now referencing the correct definition

- update `SystemComponent.Status` is now of type `SystemComponentStatus`
  - create `SystemComponentStatus`

- fix `IncludeAll struct` is now a type alias for `map[string]interface{}` in order to support marshaling and field access for what is an `anonymous` schema object 

- rename `AtFrequency` to `FrequencyCondition` per schema title

- rename `OnDate` to `OnDateCondition` per schema title

- rename `WithinDateRange` to `OnDateRangeCondition` per schema title

- update `ImplementedRequirement struct` with the previous definition for `AssemblyOscalSspImplementedRequirement` which was previously unlinked
  - create `ImplementedRequirementControlImplementation struct` with the definition previously associated with `ImplementedRequirement` which is now linked to `ControlImplementationSet` used by `Catalog` and `DefinedComponent`

- update `FindingTarget.Status` to be of type `ObjectiveStatus`
  - fixes incorrect definition for `FindingTarget.Status` that had missing field `Reason`
  - create `ObjectiveStatus`

- Rename `Statement` to `ControlStatementImplementation`
  - rename `AssemblyOscalSspStatement` to `Statement`
  - fix `ImplementedRequirement.Statements` not being of the correct definition

- create `RiskResponseReference struct` used by previously missing `RiskLogEntry` type

- update `RiskLog.Entries` to be of type `RiskLogEntry`
  - create `RiskLogEntry struct` type.
  - fix `RiskLog.Entries` was previously linked to the wrong definition of `AssessmentLogEntry` which differs from the `RiskLogEntry`

### 1.0.4, 1.0.5, 1.0.6

- rename `AssemblyOscalProfileSelectControlById` to `SelectControlById`
  - fixes `Import` and `InsertControls` `IncludeControls` and `ExcludeControls` now referencing correct definition

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

- Fix `CustomGrouping.Groups` (previously `Custom.Groups`) to be of type `ControlGroup` as the definition differs from that of previously declared `Group struct`

### 1.1.0, 1.1.1

- rename `Flat` to `FlatWithoutGroup` due to schema title change
  - update `FlatWithoutGroup` is now a type alias for `map[string]interface{}` in order to support marshaling and field access for what is an `anonymous` schema object
  - update `Merge.Flat` to be of type `FlatWithoutGrouping`
  
- rename `SelectControlById` to `AssessedControlSelectControlById`
  - update `AssessedControls` fields `IncludeControls` and `ExcludeControls` to be of type `AssessedControlsSelectControlById` 
  
- rename `AssemblyOscalProfileSelectControlById` to `SelectControlById`
  - fixes `Import` and `InsertControls` `IncludeControls` and `ExcludeControls` now referencing correct definition
  
- rename `AssemblyOscalProfileGroup` to `CustomGroupingGroup`
  - fix improper links to `Group` (should have been `AssemblyOscalProfileGroup` previously, now `CustomGroupingGroup`) in `CustomGrouping` (previously `Custom`) and `Merge` structs.
  
- Fix `CustomGrouping.Groups` (previously `Custom.Groups`) to be of type `CustomGroupingGroup` as the definition differs from that of previously declared `Group struct`