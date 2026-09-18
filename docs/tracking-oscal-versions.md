# How We Track OSCAL Version Changes

We use Renovate and GitHub Actions to track OSCAL releases.

## Renovate

Renovate tracks the latest OSCAL version in two locations:

- `update/oscal-version.yaml`

  ```yaml
  ---
  oscal: vX.Y.Z
  ```

- `src/pkg/versioning/versioning.go`

  ```go
  const latestVersion = "X.Y.Z"
  ```

`TestGetLatestVersion` requires these values to match. The relevant Renovate configuration is:

```json
"customManagers": [
  {
    "customType": "regex",
    "fileMatch": [
      "^update/oscal-version\\.yaml$",
      "^src/pkg/versioning/versioning\\.go$"
    ],
    "matchStrings": [
      "oscal: v(?<currentValue>\\d+\\.\\d+\\.\\d+)",
      "\\s*latestVersion = \"(?<currentValue>\\d+\\.\\d+\\.\\d+)\""
    ],
    "datasourceTemplate": "github-tags",
    "depNameTemplate": "usnistgov/OSCAL"
  }
]
```

When Renovate sees a new OSCAL tag, it opens a PR that updates both declarations. Updating `update/oscal-version.yaml` triggers the GitHub Action below.

## GitHub Action

The GitHub Action in `.github/workflows/create-issue-oscal-version.yaml` runs for pull requests that change `update/oscal-version.yaml`:

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

The resulting issue prompts review of the release delta. Follow [upgrading an OSCAL version](./upgrading-oscal-version.md) to add the doctored schema and generated types.