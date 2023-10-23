# OSCAL Component Definition 1.1.0

## Changelog
```
SSP: Change certain elements from required to optional for non-RMF use cases.
SSP: improve constraints of links for cross-referencing components and indicating where components were imported from.
POAM: add related-findings assembly.
Profile: Remove with-parent-controls from the profile model.
Metadata: add group attribute to props
Metadata: add resource fragment to links (very useful for deep-linking into elements by UUID and point to sub-element UUID)
Metadata: add actions assembly to track approvals or request for changes status.
Metadata: correct how cross-references between controls and their parts are handled.
⚠️ Mapping: re previous discussion, mapping, by itself or within catalog, has been moved out of the v1.1.0 release.
```

## Mapping
- metadata `actions` field added (not required)
- property `group` field added (not required)
- links `resource-fragment` field added (not required)
- Revision type -> RevisionHistoryEntry type
