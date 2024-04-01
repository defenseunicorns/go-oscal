# Changelog

## [0.4.0](https://github.com/defenseunicorns/go-oscal/compare/v0.3.1...v0.4.0) (2024-04-01)


### ⚠ BREAKING CHANGES

* **generate:** embedded optional structs changed to pointers ([#170](https://github.com/defenseunicorns/go-oscal/issues/170))
* renovate validator to update config
* renovate validator to update config
* renovate validator to update config
* move the install to correct phase
* call said action
* added tool installs action to include syft

### Features

* add norm vuln scanning and updating. ([4012ecb](https://github.com/defenseunicorns/go-oscal/commit/4012ecb6401220eee3e1183a54daa50f08778e25))
* Add uuid generation functions for library use ([09d26c5](https://github.com/defenseunicorns/go-oscal/commit/09d26c5ca19a26a239935c14ff63615e464e0247))
* deterministic generation ([6aaefd0](https://github.com/defenseunicorns/go-oscal/commit/6aaefd09abb16d01485de57c812413282c78300a))
* generate command migration & convert command initial commit ([a5145fb](https://github.com/defenseunicorns/go-oscal/commit/a5145fbdfce25a22e9be50b3a021cc52ec3333c3))
* generate types for model versions 1.0.4 -&gt; 1.1.1 ([4a4c311](https://github.com/defenseunicorns/go-oscal/commit/4a4c311b34d6a296b719bc449a1576e1d94bf3f0))
* OSCAL Version Tracking ([36fe3f2](https://github.com/defenseunicorns/go-oscal/commit/36fe3f2650a2b7ff8fff2d9b3a26d2edb7ddb25e))
* oscal-renovate-version-tracking ([36fe3f2](https://github.com/defenseunicorns/go-oscal/commit/36fe3f2650a2b7ff8fff2d9b3a26d2edb7ddb25e))
* support for Assessment Plan model ([5f7912b](https://github.com/defenseunicorns/go-oscal/commit/5f7912b8ea77da88bcda05e7283cadc1828c06b8))
* support for POAM model ([447400d](https://github.com/defenseunicorns/go-oscal/commit/447400dd42da4913a7e695980c795726e82de3cd))
* Support for Security Assessment Results model ([666ceb5](https://github.com/defenseunicorns/go-oscal/commit/666ceb54f49fa67e8984774a1774f02ee30338f7))
* test for new uuid generation ([3eea2b6](https://github.com/defenseunicorns/go-oscal/commit/3eea2b636cdb81c82fa310ce8d92b755a5505954))
* tracking OSCAL version updates ([7741c27](https://github.com/defenseunicorns/go-oscal/commit/7741c27eb1b1114fc71fc08de94fbaa6c8f10af1))
* update ValidateCommand parameters to allow streamlined export ([ad8478f](https://github.com/defenseunicorns/go-oscal/commit/ad8478ffc397ab674e358a0453b460fc5b51a7d4))
* update ValidateCommand parameters to allow streamlined export ([ad8478f](https://github.com/defenseunicorns/go-oscal/commit/ad8478ffc397ab674e358a0453b460fc5b51a7d4))
* update ValidateCommand parameters to allow streamlined export ([36ddaa4](https://github.com/defenseunicorns/go-oscal/commit/36ddaa4c138edc206bc9f1f4c6755ed84b6811f1))
* upgrade-transformupdate-across-model-versions ([349dbd9](https://github.com/defenseunicorns/go-oscal/commit/349dbd9fe36fe91e9ad9f04aa388166604a1a85a))


### Bug Fixes

* add omitempty for marshall process ([34b6d6b](https://github.com/defenseunicorns/go-oscal/commit/34b6d6b5f5b06399d7c5b63eea482d3e23307382))
* added tool installs action to include syft ([785013f](https://github.com/defenseunicorns/go-oscal/commit/785013ff912f46891e565b73d68ec652f7ed53b4))
* call said action ([ef059dd](https://github.com/defenseunicorns/go-oscal/commit/ef059ddfb76aedcef10b2e010238ccff9bf98602))
* cleanup comments and whitespaces ([849c0c8](https://github.com/defenseunicorns/go-oscal/commit/849c0c873c0ea6fa5abeb2e836601048c8663018))
* cleanup dependencies from testing ([6460c45](https://github.com/defenseunicorns/go-oscal/commit/6460c45bb7cee72f929851ab04e5afb1839b1d02))
* CODEOWNER creation and dependabot removal ([e3b8b7d](https://github.com/defenseunicorns/go-oscal/commit/e3b8b7d15bf3dc42afe24e94a91e57a1a60f7e2f))
* deterministic generation of SSP ([00ec917](https://github.com/defenseunicorns/go-oscal/commit/00ec917590864d484e5bc0f342e5717b4cbb3af9))
* generate unit tests ([954e69a](https://github.com/defenseunicorns/go-oscal/commit/954e69aba1b7f60ece662f0bed7d45c75914eaf6))
* **generate:** embedded optional structs changed to pointers ([#170](https://github.com/defenseunicorns/go-oscal/issues/170)) ([22bb613](https://github.com/defenseunicorns/go-oscal/commit/22bb6137a43c5dac08576e1648c04b52bf26d6d0))
* move the install to correct phase ([db34f87](https://github.com/defenseunicorns/go-oscal/commit/db34f877982da9f5e8b05d14cec84445282e4770))
* renovate validator to update config ([bc09f3b](https://github.com/defenseunicorns/go-oscal/commit/bc09f3bd471638ff41b7fa5589a7afc64b4d5ef1))
* renovate validator to update config ([bc09f3b](https://github.com/defenseunicorns/go-oscal/commit/bc09f3bd471638ff41b7fa5589a7afc64b4d5ef1))
* renovate validator to update config ([4752c0f](https://github.com/defenseunicorns/go-oscal/commit/4752c0f39f1f7c09b2b88e317719f3008f787666))
* updated release docs ([5deeb8e](https://github.com/defenseunicorns/go-oscal/commit/5deeb8ea6583bc7c7f83c8238fc5cb74ecfbe964))


### Miscellaneous

* **gh-actions/go-releaser:** [#166](https://github.com/defenseunicorns/go-oscal/issues/166). Update/Automate release process ([#169](https://github.com/defenseunicorns/go-oscal/issues/169)) ([dd8a9f8](https://github.com/defenseunicorns/go-oscal/commit/dd8a9f89b8349dd273c92c27816b2672c265a74e))
* **main:** release 0.3.0 ([#180](https://github.com/defenseunicorns/go-oscal/issues/180)) ([daa7dab](https://github.com/defenseunicorns/go-oscal/commit/daa7dabfee40d7178c4d36b5efb1c0afb6bf9241))
* **main:** release 0.3.1 ([#185](https://github.com/defenseunicorns/go-oscal/issues/185)) ([bd57d04](https://github.com/defenseunicorns/go-oscal/commit/bd57d041502ccb5455d21f5f9b54f0e48391c699))
* **release-please-action:** [#178](https://github.com/defenseunicorns/go-oscal/issues/178) fix release-please-config for v4 ([#179](https://github.com/defenseunicorns/go-oscal/issues/179)) ([c9d746f](https://github.com/defenseunicorns/go-oscal/commit/c9d746f3d7508058c461bd528d4052e517d32a6a))
* **release-please:** update config with draft set to true for consistency with go-releaser ([#182](https://github.com/defenseunicorns/go-oscal/issues/182)) ([7d23ba5](https://github.com/defenseunicorns/go-oscal/commit/7d23ba558b05458b6d143b91542336255192a7cc))
* **workflows:** add triage label workflow ([#184](https://github.com/defenseunicorns/go-oscal/issues/184)) ([280247a](https://github.com/defenseunicorns/go-oscal/commit/280247ac32c0c66abfe41219257af751df6f6009))

## [0.3.1](https://github.com/defenseunicorns/go-oscal/compare/v0.3.0...v0.3.1) (2024-03-29)


### Miscellaneous

* **release-please:** update config with draft set to true for consistency with go-releaser ([#182](https://github.com/defenseunicorns/go-oscal/issues/182)) ([7d23ba5](https://github.com/defenseunicorns/go-oscal/commit/7d23ba558b05458b6d143b91542336255192a7cc))
* **workflows:** add triage label workflow ([#184](https://github.com/defenseunicorns/go-oscal/issues/184)) ([280247a](https://github.com/defenseunicorns/go-oscal/commit/280247ac32c0c66abfe41219257af751df6f6009))

## [0.3.0](https://github.com/defenseunicorns/go-oscal/compare/v0.2.5...v0.3.0) (2024-03-26)


### ⚠ BREAKING CHANGES

* **generate:** embedded optional structs changed to pointers ([#170](https://github.com/defenseunicorns/go-oscal/issues/170))

### Bug Fixes

* **generate:** embedded optional structs changed to pointers ([#170](https://github.com/defenseunicorns/go-oscal/issues/170)) ([22bb613](https://github.com/defenseunicorns/go-oscal/commit/22bb6137a43c5dac08576e1648c04b52bf26d6d0))


### Miscellaneous

* **gh-actions/go-releaser:** [#166](https://github.com/defenseunicorns/go-oscal/issues/166). Update/Automate release process ([#169](https://github.com/defenseunicorns/go-oscal/issues/169)) ([dd8a9f8](https://github.com/defenseunicorns/go-oscal/commit/dd8a9f89b8349dd273c92c27816b2672c265a74e))
* **release-please-action:** [#178](https://github.com/defenseunicorns/go-oscal/issues/178) fix release-please-config for v4 ([#179](https://github.com/defenseunicorns/go-oscal/issues/179)) ([c9d746f](https://github.com/defenseunicorns/go-oscal/commit/c9d746f3d7508058c461bd528d4052e517d32a6a))
