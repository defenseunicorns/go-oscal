# Changelog

## [0.6.0](https://github.com/defenseunicorns/go-oscal/compare/v0.5.0...v0.6.0) (2024-06-18)


### ⚠ BREAKING CHANGES

* **validation:** #266 validation.ValidationCommand no longer returns error if validation was run ([#267](https://github.com/defenseunicorns/go-oscal/issues/267))
* **utils:** #162 extract utils into individual packages ([#191](https://github.com/defenseunicorns/go-oscal/issues/191))
* **generate:** embedded optional structs changed to pointers ([#170](https://github.com/defenseunicorns/go-oscal/issues/170))
* renovate validator to update config
* renovate validator to update config
* renovate validator to update config
* move the install to correct phase
* call said action
* added tool installs action to include syft

### refactor

* **utils:** [#162](https://github.com/defenseunicorns/go-oscal/issues/162) extract utils into individual packages ([#191](https://github.com/defenseunicorns/go-oscal/issues/191)) ([a0cf9ae](https://github.com/defenseunicorns/go-oscal/commit/a0cf9ae9e3e4c303279a3fdfe78063e50fa5f154))


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
* **validate:** support for writing multiple validation results ([#261](https://github.com/defenseunicorns/go-oscal/issues/261)) ([e760ba6](https://github.com/defenseunicorns/go-oscal/commit/e760ba6baec1dff6a3645ed33126859a8c90f867))


### Bug Fixes

* add omitempty for marshall process ([34b6d6b](https://github.com/defenseunicorns/go-oscal/commit/34b6d6b5f5b06399d7c5b63eea482d3e23307382))
* added tool installs action to include syft ([785013f](https://github.com/defenseunicorns/go-oscal/commit/785013ff912f46891e565b73d68ec652f7ed53b4))
* call said action ([ef059dd](https://github.com/defenseunicorns/go-oscal/commit/ef059ddfb76aedcef10b2e010238ccff9bf98602))
* cleanup comments and whitespaces ([849c0c8](https://github.com/defenseunicorns/go-oscal/commit/849c0c873c0ea6fa5abeb2e836601048c8663018))
* cleanup dependencies from testing ([6460c45](https://github.com/defenseunicorns/go-oscal/commit/6460c45bb7cee72f929851ab04e5afb1839b1d02))
* CODEOWNER creation and dependabot removal ([e3b8b7d](https://github.com/defenseunicorns/go-oscal/commit/e3b8b7d15bf3dc42afe24e94a91e57a1a60f7e2f))
* **commitlint:** [#242](https://github.com/defenseunicorns/go-oscal/issues/242) fix pinned dependencies ([#243](https://github.com/defenseunicorns/go-oscal/issues/243)) ([9b77990](https://github.com/defenseunicorns/go-oscal/commit/9b779909d87cdc507d367961a28f4f38c143a1e7))
* **commitlint:** set to ci and see if that effects the pinned dep ([#245](https://github.com/defenseunicorns/go-oscal/issues/245)) ([4844328](https://github.com/defenseunicorns/go-oscal/commit/4844328a3241bc0c463122b54c01b857e700c6ca))
* deterministic generation of SSP ([00ec917](https://github.com/defenseunicorns/go-oscal/commit/00ec917590864d484e5bc0f342e5717b4cbb3af9))
* generate unit tests ([954e69a](https://github.com/defenseunicorns/go-oscal/commit/954e69aba1b7f60ece662f0bed7d45c75914eaf6))
* **generate:** embedded optional structs changed to pointers ([#170](https://github.com/defenseunicorns/go-oscal/issues/170)) ([22bb613](https://github.com/defenseunicorns/go-oscal/commit/22bb6137a43c5dac08576e1648c04b52bf26d6d0))
* **go-releaser:** replace deprecated commands with updated commands ([#256](https://github.com/defenseunicorns/go-oscal/issues/256)) ([38e2365](https://github.com/defenseunicorns/go-oscal/commit/38e23654659e87f8d5afae4023aaa57d23ab80cf))
* move the install to correct phase ([db34f87](https://github.com/defenseunicorns/go-oscal/commit/db34f877982da9f5e8b05d14cec84445282e4770))
* renovate validator to update config ([bc09f3b](https://github.com/defenseunicorns/go-oscal/commit/bc09f3bd471638ff41b7fa5589a7afc64b4d5ef1))
* renovate validator to update config ([bc09f3b](https://github.com/defenseunicorns/go-oscal/commit/bc09f3bd471638ff41b7fa5589a7afc64b4d5ef1))
* renovate validator to update config ([4752c0f](https://github.com/defenseunicorns/go-oscal/commit/4752c0f39f1f7c09b2b88e317719f3008f787666))
* renovate.json ([#192](https://github.com/defenseunicorns/go-oscal/issues/192)) ([c4f3344](https://github.com/defenseunicorns/go-oscal/commit/c4f334474e0cb8e45df6f06b69e577127cd97e39))
* **renovate.json:** [#195](https://github.com/defenseunicorns/go-oscal/issues/195) update with semantic commits (type: chore) ([#198](https://github.com/defenseunicorns/go-oscal/issues/198)) ([678e80f](https://github.com/defenseunicorns/go-oscal/commit/678e80fa142dc12535471e7cdd705bd1af3270bc))
* updated release docs ([5deeb8e](https://github.com/defenseunicorns/go-oscal/commit/5deeb8ea6583bc7c7f83c8238fc5cb74ecfbe964))
* **validation:** [#266](https://github.com/defenseunicorns/go-oscal/issues/266) validation.ValidationCommand no longer returns error if validation was run ([#267](https://github.com/defenseunicorns/go-oscal/issues/267)) ([3264df9](https://github.com/defenseunicorns/go-oscal/commit/3264df9c539ae92a2044fb54138199c6cc77f7c8))
* **validation:** validation.ExtractErrors now properly handles empty top level ([#263](https://github.com/defenseunicorns/go-oscal/issues/263)) ([2e17507](https://github.com/defenseunicorns/go-oscal/commit/2e17507d1e8b57e8faf2a0093cce7492579bfd95))


### Miscellaneous

* **actions:** fix code scanning alerts ([#232](https://github.com/defenseunicorns/go-oscal/issues/232)) ([6b94a90](https://github.com/defenseunicorns/go-oscal/commit/6b94a9098f062b51a74477501f4ba5f4e430b164))
* **deps:** update actions/checkout action to v4.1.1 ([#193](https://github.com/defenseunicorns/go-oscal/issues/193)) ([614ba1b](https://github.com/defenseunicorns/go-oscal/commit/614ba1b5cbbadf3fa17d052dd1a72c7aa35f52fb))
* **deps:** update actions/checkout action to v4.1.3 ([#205](https://github.com/defenseunicorns/go-oscal/issues/205)) ([4a33828](https://github.com/defenseunicorns/go-oscal/commit/4a33828eec01d8032a9cb914e4ff04826f8bcfd2))
* **deps:** update actions/checkout action to v4.1.4 ([#209](https://github.com/defenseunicorns/go-oscal/issues/209)) ([01e92be](https://github.com/defenseunicorns/go-oscal/commit/01e92bee27828248a7ded3cbcb1f68efe69441a0))
* **deps:** update actions/checkout action to v4.1.5 ([#214](https://github.com/defenseunicorns/go-oscal/issues/214)) ([33ee71a](https://github.com/defenseunicorns/go-oscal/commit/33ee71a7bb3a41b57089443f4cd61b199387b2a1))
* **deps:** update actions/checkout action to v4.1.6 ([#223](https://github.com/defenseunicorns/go-oscal/issues/223)) ([d8c4c7e](https://github.com/defenseunicorns/go-oscal/commit/d8c4c7e7e14110d75e77649ad55c93d56266915b))
* **deps:** update actions/checkout action to v4.1.7 ([#259](https://github.com/defenseunicorns/go-oscal/issues/259)) ([08973c9](https://github.com/defenseunicorns/go-oscal/commit/08973c964152b0818c32554f741ed5ebde5faf92))
* **deps:** update actions/download-artifact action to v4.1.5 ([#194](https://github.com/defenseunicorns/go-oscal/issues/194)) ([78f56a5](https://github.com/defenseunicorns/go-oscal/commit/78f56a536c5ce4d862af0d3e231e2a5e19acff32))
* **deps:** update actions/download-artifact action to v4.1.7 ([#206](https://github.com/defenseunicorns/go-oscal/issues/206)) ([1453390](https://github.com/defenseunicorns/go-oscal/commit/1453390d673da0ac3a4635a5626ec8b8d6983293))
* **deps:** update actions/setup-go action to v5.0.1 ([#212](https://github.com/defenseunicorns/go-oscal/issues/212)) ([f978b2d](https://github.com/defenseunicorns/go-oscal/commit/f978b2d5cc83a4505aa5291d492ea73039e548bf))
* **deps:** update actions/setup-node action to v4 ([#203](https://github.com/defenseunicorns/go-oscal/issues/203)) ([45aecb3](https://github.com/defenseunicorns/go-oscal/commit/45aecb3bfcb6de79e105b429f0b3be548909b55e))
* **deps:** update actions/upload-artifact action to v4.3.2 ([#197](https://github.com/defenseunicorns/go-oscal/issues/197)) ([898b46f](https://github.com/defenseunicorns/go-oscal/commit/898b46f4a5d3ad895bdc6a43496e23633d098c58))
* **deps:** update actions/upload-artifact action to v4.3.3 ([#207](https://github.com/defenseunicorns/go-oscal/issues/207)) ([b9093f7](https://github.com/defenseunicorns/go-oscal/commit/b9093f78477b196bd29e1d4621f0a7180fa7a9a7))
* **deps:** update anchore/sbom-action action to v0.15.10 ([#199](https://github.com/defenseunicorns/go-oscal/issues/199)) ([32acfd7](https://github.com/defenseunicorns/go-oscal/commit/32acfd767a2f9d7217d1819d11526f3d7e54d012))
* **deps:** update anchore/sbom-action action to v0.15.11 ([#211](https://github.com/defenseunicorns/go-oscal/issues/211)) ([aec5c41](https://github.com/defenseunicorns/go-oscal/commit/aec5c41def73b31a923968d72ffad035f406a811))
* **deps:** update anchore/sbom-action action to v0.16.0 ([#229](https://github.com/defenseunicorns/go-oscal/issues/229)) ([0c00239](https://github.com/defenseunicorns/go-oscal/commit/0c002395f54fbf7a01f96de197e6006b86ce3152))
* **deps:** update dependency go to v1.22.2 ([#200](https://github.com/defenseunicorns/go-oscal/issues/200)) ([d48fa51](https://github.com/defenseunicorns/go-oscal/commit/d48fa51714946ce870f672dc7a9327863bacfec3))
* **deps:** update github/codeql-action action to v3.25.1 ([#202](https://github.com/defenseunicorns/go-oscal/issues/202)) ([25ebdef](https://github.com/defenseunicorns/go-oscal/commit/25ebdef38e875a88fd6b81a6c6e4a18471ed30e8))
* **deps:** update github/codeql-action action to v3.25.10 ([#258](https://github.com/defenseunicorns/go-oscal/issues/258)) ([004addd](https://github.com/defenseunicorns/go-oscal/commit/004addd22026d4cba26a0eb1dabc09698d412ef3))
* **deps:** update github/codeql-action action to v3.25.2 ([#208](https://github.com/defenseunicorns/go-oscal/issues/208)) ([412b27a](https://github.com/defenseunicorns/go-oscal/commit/412b27ac842d05c25b888e4767ffb7d9466946e3))
* **deps:** update github/codeql-action action to v3.25.3 ([#210](https://github.com/defenseunicorns/go-oscal/issues/210)) ([b07d331](https://github.com/defenseunicorns/go-oscal/commit/b07d331e1744fa66dc17b1a7d0ed6b1475c58947))
* **deps:** update github/codeql-action action to v3.25.4 ([#213](https://github.com/defenseunicorns/go-oscal/issues/213)) ([ad3cf0b](https://github.com/defenseunicorns/go-oscal/commit/ad3cf0b9f62b188ff1aef927e75935a49f54b4a3))
* **deps:** update github/codeql-action action to v3.25.5 ([#218](https://github.com/defenseunicorns/go-oscal/issues/218)) ([b0fbd94](https://github.com/defenseunicorns/go-oscal/commit/b0fbd94452a75de2ac3300f039fa41613145b3e6))
* **deps:** update github/codeql-action action to v3.25.6 ([#228](https://github.com/defenseunicorns/go-oscal/issues/228)) ([d95a749](https://github.com/defenseunicorns/go-oscal/commit/d95a749a904a565a28243dc5183938d40424d6fe))
* **deps:** update github/codeql-action action to v3.25.7 ([#244](https://github.com/defenseunicorns/go-oscal/issues/244)) ([a49ede0](https://github.com/defenseunicorns/go-oscal/commit/a49ede08aaac118c4060cb5ea2323748cb94540d))
* **deps:** update github/codeql-action action to v3.25.8 ([#247](https://github.com/defenseunicorns/go-oscal/issues/247)) ([5fce8ec](https://github.com/defenseunicorns/go-oscal/commit/5fce8ec71aaaee59a4f46261ba7dfd14743882d8))
* **deps:** update go version to 1.22.3 ([#226](https://github.com/defenseunicorns/go-oscal/issues/226)) ([ad903f8](https://github.com/defenseunicorns/go-oscal/commit/ad903f8a726be8d253ba321cea26b33992a3f29c))
* **deps:** update googleapis/release-please-action digest to 7987652 ([#257](https://github.com/defenseunicorns/go-oscal/issues/257)) ([71db26b](https://github.com/defenseunicorns/go-oscal/commit/71db26b5b56be2c6e8609a4584728359d1276135))
* **deps:** update googleapis/release-please-action digest to f3969c0 ([#222](https://github.com/defenseunicorns/go-oscal/issues/222)) ([8f8302e](https://github.com/defenseunicorns/go-oscal/commit/8f8302eab0710d7a6d6e6174bfcbd49f5dfb1c1b))
* **deps:** update goreleaser/goreleaser-action action to v5.1.0 ([#216](https://github.com/defenseunicorns/go-oscal/issues/216)) ([03728e3](https://github.com/defenseunicorns/go-oscal/commit/03728e3e1625e4220c6e6b84e8b8ea398309090b))
* **deps:** update goreleaser/goreleaser-action action to v6 ([#249](https://github.com/defenseunicorns/go-oscal/issues/249)) ([9886d60](https://github.com/defenseunicorns/go-oscal/commit/9886d6024d88bae2f5968f8ac978a006b46a1513))
* **deps:** update module github.com/santhosh-tekuri/jsonschema/v5 to v6 ([#248](https://github.com/defenseunicorns/go-oscal/issues/248)) ([456d3a8](https://github.com/defenseunicorns/go-oscal/commit/456d3a870b587a1c59ffdb0d578e154de23a00ed))
* **deps:** update module github.com/santhosh-tekuri/jsonschema/v5 to v6 ([#251](https://github.com/defenseunicorns/go-oscal/issues/251)) ([a26a473](https://github.com/defenseunicorns/go-oscal/commit/a26a473b438b5bcfeceef3d4c39fe0bee2b853a0))
* **deps:** update module github.com/spf13/cobra to v1.8.1 ([#264](https://github.com/defenseunicorns/go-oscal/issues/264)) ([b05ed96](https://github.com/defenseunicorns/go-oscal/commit/b05ed96813fe321dc264e56f3155c9cdf373899f))
* **deps:** update module github.com/swaggest/jsonschema-go to v0.3.70 ([#201](https://github.com/defenseunicorns/go-oscal/issues/201)) ([5a06269](https://github.com/defenseunicorns/go-oscal/commit/5a06269b2eb2b61b83d00b20c355d6a5a39800de))
* **deps:** update ossf/scorecard-action action to v2.3.3 ([#215](https://github.com/defenseunicorns/go-oscal/issues/215)) ([33cbc44](https://github.com/defenseunicorns/go-oscal/commit/33cbc4452be5a73cc64e340b5b680add5aa75b99))
* **gh-actions/go-releaser:** [#166](https://github.com/defenseunicorns/go-oscal/issues/166). Update/Automate release process ([#169](https://github.com/defenseunicorns/go-oscal/issues/169)) ([dd8a9f8](https://github.com/defenseunicorns/go-oscal/commit/dd8a9f89b8349dd273c92c27816b2672c265a74e))
* **go-releaser:** [#187](https://github.com/defenseunicorns/go-oscal/issues/187) no longer creates draft and executes on tag creation ([#188](https://github.com/defenseunicorns/go-oscal/issues/188)) ([c419739](https://github.com/defenseunicorns/go-oscal/commit/c419739d9bc355c5848ed2cce905deef1ee7061f))
* **main:** release 0.3.0 ([#180](https://github.com/defenseunicorns/go-oscal/issues/180)) ([daa7dab](https://github.com/defenseunicorns/go-oscal/commit/daa7dabfee40d7178c4d36b5efb1c0afb6bf9241))
* **main:** release 0.3.1 ([#185](https://github.com/defenseunicorns/go-oscal/issues/185)) ([bd57d04](https://github.com/defenseunicorns/go-oscal/commit/bd57d041502ccb5455d21f5f9b54f0e48391c699))
* **main:** release 0.3.2 ([#189](https://github.com/defenseunicorns/go-oscal/issues/189)) ([20c2a22](https://github.com/defenseunicorns/go-oscal/commit/20c2a2269af69e762c4328204677a8513dfd82a4))
* **main:** release 0.4.0 ([#204](https://github.com/defenseunicorns/go-oscal/issues/204)) ([16e60be](https://github.com/defenseunicorns/go-oscal/commit/16e60be412ef49bab8a4edb48710f7d210fb6cca))
* **main:** release 0.4.1 ([#224](https://github.com/defenseunicorns/go-oscal/issues/224)) ([9b7e4ea](https://github.com/defenseunicorns/go-oscal/commit/9b7e4ea0d74383229d68f39d009e85cfbaf188a3))
* **main:** release 0.4.2 ([#230](https://github.com/defenseunicorns/go-oscal/issues/230)) ([170d01d](https://github.com/defenseunicorns/go-oscal/commit/170d01de8de2d2ed51c73771ab54dcc4fd140d4b))
* **main:** release 0.4.3 ([#252](https://github.com/defenseunicorns/go-oscal/issues/252)) ([1921e9f](https://github.com/defenseunicorns/go-oscal/commit/1921e9f6f2d8474abe9fbd688c216b82da344a52))
* **main:** release 0.5.0 ([#265](https://github.com/defenseunicorns/go-oscal/issues/265)) ([562091d](https://github.com/defenseunicorns/go-oscal/commit/562091d45373efe4b358944ac7a115a2c1363f7b))
* **release-please-action:** [#178](https://github.com/defenseunicorns/go-oscal/issues/178) fix release-please-config for v4 ([#179](https://github.com/defenseunicorns/go-oscal/issues/179)) ([c9d746f](https://github.com/defenseunicorns/go-oscal/commit/c9d746f3d7508058c461bd528d4052e517d32a6a))
* **release-please:** update config with draft set to true for consistency with go-releaser ([#182](https://github.com/defenseunicorns/go-oscal/issues/182)) ([7d23ba5](https://github.com/defenseunicorns/go-oscal/commit/7d23ba558b05458b6d143b91542336255192a7cc))
* **workflows:** add triage label workflow ([#184](https://github.com/defenseunicorns/go-oscal/issues/184)) ([280247a](https://github.com/defenseunicorns/go-oscal/commit/280247ac32c0c66abfe41219257af751df6f6009))

## [0.5.0](https://github.com/defenseunicorns/go-oscal/compare/v0.4.3...v0.5.0) (2024-06-18)


### ⚠ BREAKING CHANGES

* **validation:** #266 validation.ValidationCommand no longer returns error if validation was run ([#267](https://github.com/defenseunicorns/go-oscal/issues/267))

### Bug Fixes

* **validation:** [#266](https://github.com/defenseunicorns/go-oscal/issues/266) validation.ValidationCommand no longer returns error if validation was run ([#267](https://github.com/defenseunicorns/go-oscal/issues/267)) ([3264df9](https://github.com/defenseunicorns/go-oscal/commit/3264df9c539ae92a2044fb54138199c6cc77f7c8))


### Miscellaneous

* **deps:** update github/codeql-action action to v3.25.10 ([#258](https://github.com/defenseunicorns/go-oscal/issues/258)) ([004addd](https://github.com/defenseunicorns/go-oscal/commit/004addd22026d4cba26a0eb1dabc09698d412ef3))

## [0.4.3](https://github.com/defenseunicorns/go-oscal/compare/v0.4.2...v0.4.3) (2024-06-17)


### Features

* **validate:** support for writing multiple validation results ([#261](https://github.com/defenseunicorns/go-oscal/issues/261)) ([e760ba6](https://github.com/defenseunicorns/go-oscal/commit/e760ba6baec1dff6a3645ed33126859a8c90f867))


### Bug Fixes

* **go-releaser:** replace deprecated commands with updated commands ([#256](https://github.com/defenseunicorns/go-oscal/issues/256)) ([38e2365](https://github.com/defenseunicorns/go-oscal/commit/38e23654659e87f8d5afae4023aaa57d23ab80cf))
* **validation:** validation.ExtractErrors now properly handles empty top level ([#263](https://github.com/defenseunicorns/go-oscal/issues/263)) ([2e17507](https://github.com/defenseunicorns/go-oscal/commit/2e17507d1e8b57e8faf2a0093cce7492579bfd95))


### Miscellaneous

* **deps:** update actions/checkout action to v4.1.7 ([#259](https://github.com/defenseunicorns/go-oscal/issues/259)) ([08973c9](https://github.com/defenseunicorns/go-oscal/commit/08973c964152b0818c32554f741ed5ebde5faf92))
* **deps:** update googleapis/release-please-action digest to 7987652 ([#257](https://github.com/defenseunicorns/go-oscal/issues/257)) ([71db26b](https://github.com/defenseunicorns/go-oscal/commit/71db26b5b56be2c6e8609a4584728359d1276135))
* **deps:** update module github.com/santhosh-tekuri/jsonschema/v5 to v6 ([#251](https://github.com/defenseunicorns/go-oscal/issues/251)) ([a26a473](https://github.com/defenseunicorns/go-oscal/commit/a26a473b438b5bcfeceef3d4c39fe0bee2b853a0))
* **deps:** update module github.com/spf13/cobra to v1.8.1 ([#264](https://github.com/defenseunicorns/go-oscal/issues/264)) ([b05ed96](https://github.com/defenseunicorns/go-oscal/commit/b05ed96813fe321dc264e56f3155c9cdf373899f))

## [0.4.2](https://github.com/defenseunicorns/go-oscal/compare/v0.4.1...v0.4.2) (2024-06-06)


### Bug Fixes

* **commitlint:** [#242](https://github.com/defenseunicorns/go-oscal/issues/242) fix pinned dependencies ([#243](https://github.com/defenseunicorns/go-oscal/issues/243)) ([9b77990](https://github.com/defenseunicorns/go-oscal/commit/9b779909d87cdc507d367961a28f4f38c143a1e7))
* **commitlint:** set to ci and see if that effects the pinned dep ([#245](https://github.com/defenseunicorns/go-oscal/issues/245)) ([4844328](https://github.com/defenseunicorns/go-oscal/commit/4844328a3241bc0c463122b54c01b857e700c6ca))


### Miscellaneous

* **actions:** fix code scanning alerts ([#232](https://github.com/defenseunicorns/go-oscal/issues/232)) ([6b94a90](https://github.com/defenseunicorns/go-oscal/commit/6b94a9098f062b51a74477501f4ba5f4e430b164))
* **deps:** update github/codeql-action action to v3.25.7 ([#244](https://github.com/defenseunicorns/go-oscal/issues/244)) ([a49ede0](https://github.com/defenseunicorns/go-oscal/commit/a49ede08aaac118c4060cb5ea2323748cb94540d))
* **deps:** update github/codeql-action action to v3.25.8 ([#247](https://github.com/defenseunicorns/go-oscal/issues/247)) ([5fce8ec](https://github.com/defenseunicorns/go-oscal/commit/5fce8ec71aaaee59a4f46261ba7dfd14743882d8))
* **deps:** update goreleaser/goreleaser-action action to v6 ([#249](https://github.com/defenseunicorns/go-oscal/issues/249)) ([9886d60](https://github.com/defenseunicorns/go-oscal/commit/9886d6024d88bae2f5968f8ac978a006b46a1513))
* **deps:** update module github.com/santhosh-tekuri/jsonschema/v5 to v6 ([#248](https://github.com/defenseunicorns/go-oscal/issues/248)) ([456d3a8](https://github.com/defenseunicorns/go-oscal/commit/456d3a870b587a1c59ffdb0d578e154de23a00ed))

## [0.4.1](https://github.com/defenseunicorns/go-oscal/compare/v0.4.0...v0.4.1) (2024-05-23)


### Miscellaneous

* **deps:** update anchore/sbom-action action to v0.16.0 ([#229](https://github.com/defenseunicorns/go-oscal/issues/229)) ([0c00239](https://github.com/defenseunicorns/go-oscal/commit/0c002395f54fbf7a01f96de197e6006b86ce3152))
* **deps:** update github/codeql-action action to v3.25.6 ([#228](https://github.com/defenseunicorns/go-oscal/issues/228)) ([d95a749](https://github.com/defenseunicorns/go-oscal/commit/d95a749a904a565a28243dc5183938d40424d6fe))
* **deps:** update go version to 1.22.3 ([#226](https://github.com/defenseunicorns/go-oscal/issues/226)) ([ad903f8](https://github.com/defenseunicorns/go-oscal/commit/ad903f8a726be8d253ba321cea26b33992a3f29c))

## [0.4.0](https://github.com/defenseunicorns/go-oscal/compare/v0.3.2...v0.4.0) (2024-05-20)


### ⚠ BREAKING CHANGES

* **utils:** #162 extract utils into individual packages ([#191](https://github.com/defenseunicorns/go-oscal/issues/191))

### refactor

* **utils:** [#162](https://github.com/defenseunicorns/go-oscal/issues/162) extract utils into individual packages ([#191](https://github.com/defenseunicorns/go-oscal/issues/191)) ([a0cf9ae](https://github.com/defenseunicorns/go-oscal/commit/a0cf9ae9e3e4c303279a3fdfe78063e50fa5f154))


### Miscellaneous

* **deps:** update actions/checkout action to v4.1.3 ([#205](https://github.com/defenseunicorns/go-oscal/issues/205)) ([4a33828](https://github.com/defenseunicorns/go-oscal/commit/4a33828eec01d8032a9cb914e4ff04826f8bcfd2))
* **deps:** update actions/checkout action to v4.1.4 ([#209](https://github.com/defenseunicorns/go-oscal/issues/209)) ([01e92be](https://github.com/defenseunicorns/go-oscal/commit/01e92bee27828248a7ded3cbcb1f68efe69441a0))
* **deps:** update actions/checkout action to v4.1.5 ([#214](https://github.com/defenseunicorns/go-oscal/issues/214)) ([33ee71a](https://github.com/defenseunicorns/go-oscal/commit/33ee71a7bb3a41b57089443f4cd61b199387b2a1))
* **deps:** update actions/checkout action to v4.1.6 ([#223](https://github.com/defenseunicorns/go-oscal/issues/223)) ([d8c4c7e](https://github.com/defenseunicorns/go-oscal/commit/d8c4c7e7e14110d75e77649ad55c93d56266915b))
* **deps:** update actions/download-artifact action to v4.1.7 ([#206](https://github.com/defenseunicorns/go-oscal/issues/206)) ([1453390](https://github.com/defenseunicorns/go-oscal/commit/1453390d673da0ac3a4635a5626ec8b8d6983293))
* **deps:** update actions/setup-go action to v5.0.1 ([#212](https://github.com/defenseunicorns/go-oscal/issues/212)) ([f978b2d](https://github.com/defenseunicorns/go-oscal/commit/f978b2d5cc83a4505aa5291d492ea73039e548bf))
* **deps:** update actions/upload-artifact action to v4.3.3 ([#207](https://github.com/defenseunicorns/go-oscal/issues/207)) ([b9093f7](https://github.com/defenseunicorns/go-oscal/commit/b9093f78477b196bd29e1d4621f0a7180fa7a9a7))
* **deps:** update anchore/sbom-action action to v0.15.11 ([#211](https://github.com/defenseunicorns/go-oscal/issues/211)) ([aec5c41](https://github.com/defenseunicorns/go-oscal/commit/aec5c41def73b31a923968d72ffad035f406a811))
* **deps:** update github/codeql-action action to v3.25.2 ([#208](https://github.com/defenseunicorns/go-oscal/issues/208)) ([412b27a](https://github.com/defenseunicorns/go-oscal/commit/412b27ac842d05c25b888e4767ffb7d9466946e3))
* **deps:** update github/codeql-action action to v3.25.3 ([#210](https://github.com/defenseunicorns/go-oscal/issues/210)) ([b07d331](https://github.com/defenseunicorns/go-oscal/commit/b07d331e1744fa66dc17b1a7d0ed6b1475c58947))
* **deps:** update github/codeql-action action to v3.25.4 ([#213](https://github.com/defenseunicorns/go-oscal/issues/213)) ([ad3cf0b](https://github.com/defenseunicorns/go-oscal/commit/ad3cf0b9f62b188ff1aef927e75935a49f54b4a3))
* **deps:** update github/codeql-action action to v3.25.5 ([#218](https://github.com/defenseunicorns/go-oscal/issues/218)) ([b0fbd94](https://github.com/defenseunicorns/go-oscal/commit/b0fbd94452a75de2ac3300f039fa41613145b3e6))
* **deps:** update googleapis/release-please-action digest to f3969c0 ([#222](https://github.com/defenseunicorns/go-oscal/issues/222)) ([8f8302e](https://github.com/defenseunicorns/go-oscal/commit/8f8302eab0710d7a6d6e6174bfcbd49f5dfb1c1b))
* **deps:** update goreleaser/goreleaser-action action to v5.1.0 ([#216](https://github.com/defenseunicorns/go-oscal/issues/216)) ([03728e3](https://github.com/defenseunicorns/go-oscal/commit/03728e3e1625e4220c6e6b84e8b8ea398309090b))
* **deps:** update ossf/scorecard-action action to v2.3.3 ([#215](https://github.com/defenseunicorns/go-oscal/issues/215)) ([33cbc44](https://github.com/defenseunicorns/go-oscal/commit/33cbc4452be5a73cc64e340b5b680add5aa75b99))

## [0.3.2](https://github.com/defenseunicorns/go-oscal/compare/v0.3.1...v0.3.2) (2024-04-19)


### Bug Fixes

* renovate.json ([#192](https://github.com/defenseunicorns/go-oscal/issues/192)) ([c4f3344](https://github.com/defenseunicorns/go-oscal/commit/c4f334474e0cb8e45df6f06b69e577127cd97e39))
* **renovate.json:** [#195](https://github.com/defenseunicorns/go-oscal/issues/195) update with semantic commits (type: chore) ([#198](https://github.com/defenseunicorns/go-oscal/issues/198)) ([678e80f](https://github.com/defenseunicorns/go-oscal/commit/678e80fa142dc12535471e7cdd705bd1af3270bc))


### Miscellaneous

* **deps:** update actions/checkout action to v4.1.1 ([#193](https://github.com/defenseunicorns/go-oscal/issues/193)) ([614ba1b](https://github.com/defenseunicorns/go-oscal/commit/614ba1b5cbbadf3fa17d052dd1a72c7aa35f52fb))
* **deps:** update actions/download-artifact action to v4.1.5 ([#194](https://github.com/defenseunicorns/go-oscal/issues/194)) ([78f56a5](https://github.com/defenseunicorns/go-oscal/commit/78f56a536c5ce4d862af0d3e231e2a5e19acff32))
* **deps:** update actions/setup-node action to v4 ([#203](https://github.com/defenseunicorns/go-oscal/issues/203)) ([45aecb3](https://github.com/defenseunicorns/go-oscal/commit/45aecb3bfcb6de79e105b429f0b3be548909b55e))
* **deps:** update actions/upload-artifact action to v4.3.2 ([#197](https://github.com/defenseunicorns/go-oscal/issues/197)) ([898b46f](https://github.com/defenseunicorns/go-oscal/commit/898b46f4a5d3ad895bdc6a43496e23633d098c58))
* **deps:** update anchore/sbom-action action to v0.15.10 ([#199](https://github.com/defenseunicorns/go-oscal/issues/199)) ([32acfd7](https://github.com/defenseunicorns/go-oscal/commit/32acfd767a2f9d7217d1819d11526f3d7e54d012))
* **deps:** update dependency go to v1.22.2 ([#200](https://github.com/defenseunicorns/go-oscal/issues/200)) ([d48fa51](https://github.com/defenseunicorns/go-oscal/commit/d48fa51714946ce870f672dc7a9327863bacfec3))
* **deps:** update github/codeql-action action to v3.25.1 ([#202](https://github.com/defenseunicorns/go-oscal/issues/202)) ([25ebdef](https://github.com/defenseunicorns/go-oscal/commit/25ebdef38e875a88fd6b81a6c6e4a18471ed30e8))
* **deps:** update module github.com/swaggest/jsonschema-go to v0.3.70 ([#201](https://github.com/defenseunicorns/go-oscal/issues/201)) ([5a06269](https://github.com/defenseunicorns/go-oscal/commit/5a06269b2eb2b61b83d00b20c355d6a5a39800de))
* **go-releaser:** [#187](https://github.com/defenseunicorns/go-oscal/issues/187) no longer creates draft and executes on tag creation ([#188](https://github.com/defenseunicorns/go-oscal/issues/188)) ([c419739](https://github.com/defenseunicorns/go-oscal/commit/c419739d9bc355c5848ed2cce905deef1ee7061f))

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
