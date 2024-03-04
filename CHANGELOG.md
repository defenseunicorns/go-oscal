# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [unreleased] - yyyy-mm-dd

### Fix
- duplicate issues being created on OSCAL release to upgrade

### Add
- workflow to update the change log with version and date after release. 

## [v0.2.1] - 2024-03-01

### Fix

- utils/timestamp.GetTimeStamp now returns `time.Now()` as type `time.Time` in accordance with the changes to `oscalTypes_*`. 
  - `time.Time` is marshaled and un-marshaled to `RFC3339` by default in compliance with the OSCAL complete schema expected timestamp format. 

## [v0.2.0] - 2024-02-29

### Added
- OSCAL 1.1.2 types
  - Validation and Revision now supported for OSCAL 1.1.2
- `doctor` command to prep new oscal_complete_schema.json from [OSCAL Releases](https://github.com/usnistgov/OSCAL/releases/tag/v1.1.2) for use in the `validate` and `revise` commands. Behavior ported from [oscal-json-doctor](https://github.com/defenseunicorns/oscal-json-doctor) in order to keep all needed functionality within the go-oscal repo. Create to fix https://github.com/usnistgov/OSCAL/issues/1908 and https://github.com/usnistgov/OSCAL/issues/1908
- documentation for `generate` and `doctor` commands
- documentation for upgrading to a new version of OSCAL (docs/upgrading-oscal-version.md)
- Added support for custom types.
- Added validation and type support for Oscal Complete Schema 1.1.2

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

- update fields that are of format `date-time` in schema are now represented as `time.Time` in Types.

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

## [v0.1.0] - 2024-01-05
note: Bumping minor release in order to create a clean release baseline. 

## What's Changed
* deps Update dependency go to v1.21.5 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/72
* fix: CODEOWNER creation and dependabot removal by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/75
* deps Update actions/setup-go action to v5 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/76
* deps Update github/codeql-action action to v2.22.9 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/80
* feat: OSCAL Version Tracking by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/81
* fix!: renovate validator to update config by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/89
* fix: updated release docs by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/85
* #55. Create Command Tests by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/77
* #93. Fix test_output.go beind dumped into the cmd/generate directory … by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/94
* deps Update github/codeql-action action to v3 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/98
* 33 upgrade transformupdate across model versions by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/97
* deps Update module github.com/google/uuid to v1.5.0 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/96
* deps Update actions/download-artifact action to v4 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/100
* deps Update actions/upload-artifact action to v4 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/101
* #102. Create UpdateLastModified method that is used when revising a p… by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/103
* deps Update actions/download-artifact action to v4.1.0 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/104
* deps Update github/codeql-action action to v3.22.12 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/105
* deps Update anchore/sbom-action action to v0.15.2 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/106

## [v0.0.2] - 2024-01-05

## What's Changed
* deps Update dependency go to v1.21.5 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/72
* fix: CODEOWNER creation and dependabot removal by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/75
* deps Update actions/setup-go action to v5 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/76
* deps Update github/codeql-action action to v2.22.9 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/80
* feat: OSCAL Version Tracking by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/81
* fix!: renovate validator to update config by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/89
* fix: updated release docs by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/85
* #55. Create Command Tests by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/77
* #93. Fix test_output.go beind dumped into the cmd/generate directory … by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/94
* deps Update github/codeql-action action to v3 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/98
* 33 upgrade transformupdate across model versions by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/97
* deps Update module github.com/google/uuid to v1.5.0 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/96
* deps Update actions/download-artifact action to v4 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/100
* deps Update actions/upload-artifact action to v4 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/101
* #102. Create UpdateLastModified method that is used when revising a p… by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/103
* deps Update actions/download-artifact action to v4.1.0 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/104
* deps Update github/codeql-action action to v3.22.12 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/105
* deps Update anchore/sbom-action action to v0.15.2 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/106

## [v0.0.1] - 2023-12-5

## What's Changed
* Add initial, experimental logic and notes about context and usage by @lucasrod16 in https://github.com/defenseunicorns/go-oscal/pull/1
* Update --input flag to --input-file in readme by @lucasrod16 in https://github.com/defenseunicorns/go-oscal/pull/3
* Updating Makefile for minor optimizations by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/6
* Initial logic for generating map for struct creation by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/5
* Code cleanup and update documentation by @lucasrod16 in https://github.com/defenseunicorns/go-oscal/pull/11
* Add header comment to generated structs by @lucasrod16 in https://github.com/defenseunicorns/go-oscal/pull/14
* Unit tests by @lucasrod16 in https://github.com/defenseunicorns/go-oscal/pull/15
* Add top-level struct generation by @lucasrod16 in https://github.com/defenseunicorns/go-oscal/pull/17
* Bump actions/setup-go from 3 to 4 by @dependabot in https://github.com/defenseunicorns/go-oscal/pull/22
* Bump github.com/spf13/cobra from 1.6.1 to 1.7.0 by @dependabot in https://github.com/defenseunicorns/go-oscal/pull/23
* Bump actions/checkout from 3 to 4 by @dependabot in https://github.com/defenseunicorns/go-oscal/pull/24
* Oscal schema updates support by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/28
* Test: additional models testing and explicit failure by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/31
* Types Export for latest supported OSCAL models by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/32
* Bump github.com/swaggest/jsonschema-go from 0.3.60 to 0.3.62 by @dependabot in https://github.com/defenseunicorns/go-oscal/pull/34
* Feat: deterministic generation by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/36
* Feat: generate command migration & convert command initial commit by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/38
* Security assessment results by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/40
* Initial commit for generation of all oscal models by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/47
* Fix: add omitempty for marshall process by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/48
* Feat: Add uuid generation functions for library use by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/51
* 37 object validation for models by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/50
* 53 validate better error messaging for failed validation by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/57
* Feat: update ValidateCommand parameters to allow streamlined export by @brandtkeller in https://github.com/defenseunicorns/go-oscal/pull/59
* feat: add norm vuln scanning and updating. by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/60
* deps Update module github.com/spf13/cobra to v1.8.0 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/62
* deps Update dependency go to v1.21.4 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/61
* 53 validate better error messaging for failed validation by @mike-winberry in https://github.com/defenseunicorns/go-oscal/pull/67
* deps Update module github.com/swaggest/jsonschema-go to v0.3.64 by @renovate in https://github.com/defenseunicorns/go-oscal/pull/66
* feat: initial release process by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/71
* fix!: added tool installs action to include syft by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/73
* fix!: move the install to correct phase by @CloudBeard in https://github.com/defenseunicorns/go-oscal/pull/74

## New Contributors
* @lucasrod16 made their first contribution in https://github.com/defenseunicorns/go-oscal/pull/1
* @brandtkeller made their first contribution in https://github.com/defenseunicorns/go-oscal/pull/6
* @dependabot made their first contribution in https://github.com/defenseunicorns/go-oscal/pull/22
* @mike-winberry made their first contribution in https://github.com/defenseunicorns/go-oscal/pull/50
* @CloudBeard made their first contribution in https://github.com/defenseunicorns/go-oscal/pull/60
* @renovate made their first contribution in https://github.com/defenseunicorns/go-oscal/pull/62


[unreleased]: https://github.com/defenseunicorns/go-oscal/compare/v0.2.1...HEAD
[v0.2.1]: https://github.com/defenseunicorns/go-oscal/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/defenseunicorns/go-oscal/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/defenseunicorns/go-oscal/compare/v0.0.2...v0.1.0
[v0.0.2]: https://github.com/defenseunicorns/go-oscal/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/defenseunicorns/go-oscal/releases/tag/v0.0.1
