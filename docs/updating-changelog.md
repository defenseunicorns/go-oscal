# Updating the changelog
The changelog should be updated before merging and after a release. 


## Before merge
- changelog should be updated in accordance with [keep-a-changelog](https://keepachangelog.com/en/1.1.0/)
- currently no automation see issue [#153](https://github.com/defenseunicorns/go-oscal/issues/153) if you would like to help. 

## After release
- After a release is published the workflow [update-changelog](../.github/workflows/update-changelog-post-release.yaml) will run and create a PR. 
- This workflow handles adding the new `## [version] - yyyy-mm-dd` tag replacing the old `## [unreleased] - yyyy-mm-dd` subheader.
- This workflow also handles updating the links with the compare to last version. 
- Any additional changes to the changelog done at release should be added before merging. 
