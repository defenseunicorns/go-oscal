# How We Track OSCAL Version Changes
We use a combination of Renovate and GitHub Actions to track OSCAL changes.

## Renovate 
A YAML file is used to track the changes to OSCAL version. located in update/oscal-version.yaml as
```yaml
---
oscal: X.X.X
```

The Renovate config file is setup to track GitHub tag changes on the USNISTGOV/OSCAL repo. Using the following configuration:
```json
  "regexManagers": [
    {
      "fileMatch": ["^update/oscal-version\\.yaml$"],
      "matchStrings": ["oscal: v(?<currentValue>\\d+\\.\\d+\\.\\d+)"],
      "datasource": "github-tags",
      "depName": "usnistgov/OSCAL",
      "versioning": "semver"
    }
  ]

```

When Renovate sees a tag change it will create a PR to update the version stored in the YAML file. This PR will trigger the GitHub Action to create an issue.

## GitHub Action

The GitHub Action located in .github/workflows/create-issue-oscal-version.yaml is triggered when a PR is created affecting this file location. The GitHub Action uses this configuration: 
```yaml
name: Create Issue on PR for OSCAL Version Updates

permissions:
  contents: read
  issues: write

on:
  pull_request:
    paths:
      - 'update/oscal-version.yaml'

jobs:
  createIssue:
    runs-on: ubuntu-latest
    steps:
      - name: Create Issue
        uses: actions/github-script@60a0d83039c74a4aee543508d2ffcb1c3799cdea # v7.0.1
        with:
          github-token: ${{secrets.GITHUB_TOKEN}}
          script: |
            const title = `PR #${context.payload.pull_request.number} OSCAL Version Update`;
            const body = `A new OSCAL version has been released. Please review the changes https://github.com/usnistgov/OSCAL/releases`;
            const projectID = '12';
            github.rest.issues.create({
              owner: 'defenseunicorns',
              repo: 'go-oscal',
              title: title,
              body: body,
              labels: ['enhancement']
            });

```

The GitHub Action creates an issue on the go-oscal repo to check out the releases for OSCAL to see what has changed. Depending on the change the level of effort for the update can vary.