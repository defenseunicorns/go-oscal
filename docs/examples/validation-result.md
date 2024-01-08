# Examples: Validation Result

## YAML

### Example Command
`./bin/go-oscal validate -f testdata/validation/invalid-catalog.yaml -r test_out/validation-result.yaml`

### Output:

```yaml
valid: false
timeStamp: 2024-01-08T09:16:06.164456554-08:00
errors:
    - basicerror:
        keywordlocation: /oneOf/0/properties/catalog/$ref/properties/metadata/$ref/required
        absolutekeywordlocation: http://csrc.nist.gov/ns/oscal/1.0/1.0.4/oscal-complete-schema.json#/definitions/oscal-complete-oscal-metadata:metadata/required
        instancelocation: /catalog/metadata
        error: 'missing properties: ''title'''
    - basicerror:
        keywordlocation: /oneOf/0/properties/catalog/$ref/properties/metadata/$ref/properties/links/items/$ref/properties/rel/pattern
        absolutekeywordlocation: http://csrc.nist.gov/ns/oscal/1.0/1.0.4/oscal-complete-schema.json#/definitions/oscal-complete-oscal-metadata:link/properties/rel/pattern
        instancelocation: /catalog/metadata/links/0/rel
        error: does not match pattern '^(\\p{L}|_)(\\p{L}|\\p{N}|[.\\-_])*$'
      failedValue: https://something.com
    - basicerror:
        keywordlocation: /oneOf/0/properties/catalog/$ref/properties/metadata/$ref/properties/parties/items/$ref/properties/uuid/pattern
        absolutekeywordlocation: http://csrc.nist.gov/ns/oscal/1.0/1.0.4/oscal-complete-schema.json#/definitions/oscal-complete-oscal-metadata:party/properties/uuid/pattern
        instancelocation: /catalog/metadata/parties/0/uuid
        error: does not match pattern '^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[45][0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$'
      failedValue: invalid-uuid
metadata:
    documentType: catalog
    documentVersion: 1.0.4
    schemaVersion: 2.0.4
```

## JSON

### Example Command

`./bin/go-oscal validate -f testdata/validation/invalid-catalog.yaml -r test_out/validation-result.json`

### Output:
```json
{
    "valid": false,
    "timeStamp": "2024-01-08T09:21:30.693531237-08:00",
    "errors": [
        {
            "keywordLocation": "/oneOf/0/properties/catalog/$ref/properties/metadata/$ref/required",
            "absoluteKeywordLocation": "http://csrc.nist.gov/ns/oscal/1.0/1.0.4/oscal-complete-schema.json#/definitions/oscal-complete-oscal-metadata:metadata/required",
            "instanceLocation": "/catalog/metadata",
            "error": "missing properties: 'title'"
        },
        {
            "keywordLocation": "/oneOf/0/properties/catalog/$ref/properties/metadata/$ref/properties/parties/items/$ref/properties/uuid/pattern",
            "absoluteKeywordLocation": "http://csrc.nist.gov/ns/oscal/1.0/1.0.4/oscal-complete-schema.json#/definitions/oscal-complete-oscal-metadata:party/properties/uuid/pattern",
            "instanceLocation": "/catalog/metadata/parties/0/uuid",
            "error": "does not match pattern '^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[45][0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$'",
            "failedValue": "invalid-uuid"
        },
        {
            "keywordLocation": "/oneOf/0/properties/catalog/$ref/properties/metadata/$ref/properties/links/items/$ref/properties/rel/pattern",
            "absoluteKeywordLocation": "http://csrc.nist.gov/ns/oscal/1.0/1.0.4/oscal-complete-schema.json#/definitions/oscal-complete-oscal-metadata:link/properties/rel/pattern",
            "instanceLocation": "/catalog/metadata/links/0/rel",
            "error": "does not match pattern '^(\\\\p{L}|_)(\\\\p{L}|\\\\p{N}|[.\\\\-_])*$'",
            "failedValue": "https://something.com"
        }
    ],
    "metadata": {
        "documentType": "catalog",
        "documentVersion": "1.0.4",
        "schemaVersion": "1.0.4"
    }
}
```
