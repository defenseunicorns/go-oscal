# Changelog

## [0.6.0](https://github.com/defenseunicorns/go-oscal/compare/v0.5.0...v0.6.0) (2024-08-07)


### ⚠ BREAKING CHANGES

* **validation:** #274 refactor to accommodate updating to jsonschema v6 ([#282](https://github.com/defenseunicorns/go-oscal/issues/282))

### refactor

* **validation:** [#274](https://github.com/defenseunicorns/go-oscal/issues/274) refactor to accommodate updating to jsonschema v6 ([#282](https://github.com/defenseunicorns/go-oscal/issues/282)) ([5e1a894](https://github.com/defenseunicorns/go-oscal/commit/5e1a8941fab19b04030db994e10d5063d08f05bd))


### Miscellaneous

* **deps:** update actions/download-artifact action to v4.1.8 ([#275](https://github.com/defenseunicorns/go-oscal/issues/275)) ([06c015f](https://github.com/defenseunicorns/go-oscal/commit/06c015f994948c030747466bde50d8beb14d4d16))
* **deps:** update actions/github-script digest to 35b1cdd ([#286](https://github.com/defenseunicorns/go-oscal/issues/286)) ([b81dbfb](https://github.com/defenseunicorns/go-oscal/commit/b81dbfbc5ab7326ecdbcf816f3b7c8603b7d4885))
* **deps:** update actions/setup-go action to v5.0.2 ([#279](https://github.com/defenseunicorns/go-oscal/issues/279)) ([40b2efe](https://github.com/defenseunicorns/go-oscal/commit/40b2efe131c04c3ff121a19bc2647bcda3c8c01c))
* **deps:** update actions/setup-node action to v4.0.3 ([#276](https://github.com/defenseunicorns/go-oscal/issues/276)) ([54f6ab5](https://github.com/defenseunicorns/go-oscal/commit/54f6ab5a132982dcbe092c6da7ff8881f045f7f5))
* **deps:** update actions/upload-artifact action to v4.3.4 ([#277](https://github.com/defenseunicorns/go-oscal/issues/277)) ([c149bc7](https://github.com/defenseunicorns/go-oscal/commit/c149bc7e0757c0e4f5a577435541fcca1b039543))
* **deps:** update actions/upload-artifact action to v4.3.6 ([#287](https://github.com/defenseunicorns/go-oscal/issues/287)) ([31eb0c5](https://github.com/defenseunicorns/go-oscal/commit/31eb0c5175edbd235237561e58992c33f891057e))
* **deps:** update anchore/sbom-action action to v0.16.1 ([#278](https://github.com/defenseunicorns/go-oscal/issues/278)) ([ad3e6aa](https://github.com/defenseunicorns/go-oscal/commit/ad3e6aae883ef6a819bcc25c8aa635e9175f2ab6))
* **deps:** update anchore/sbom-action action to v0.17.0 ([#281](https://github.com/defenseunicorns/go-oscal/issues/281)) ([5000d31](https://github.com/defenseunicorns/go-oscal/commit/5000d311d277b2d918e84f0a556108823eadd80f))
* **deps:** update dependency commitlint to v19.4.0 ([#290](https://github.com/defenseunicorns/go-oscal/issues/290)) ([d2351fd](https://github.com/defenseunicorns/go-oscal/commit/d2351fd7504d9eeddd2068b323e2dcd59222a1b3))
* **deps:** update github/codeql-action action to v3.25.11 ([#273](https://github.com/defenseunicorns/go-oscal/issues/273)) ([6910af5](https://github.com/defenseunicorns/go-oscal/commit/6910af5c5e4d0d151b3e07992bdaa9baf69549dd))
* **deps:** update github/codeql-action action to v3.25.12 ([#280](https://github.com/defenseunicorns/go-oscal/issues/280)) ([147c634](https://github.com/defenseunicorns/go-oscal/commit/147c6340b3a2a273c9909e68b698101ac82a69b9))
* **deps:** update github/codeql-action action to v3.25.13 ([#283](https://github.com/defenseunicorns/go-oscal/issues/283)) ([e7dc32a](https://github.com/defenseunicorns/go-oscal/commit/e7dc32a6830213af63c0be625897ccaf2db91e44))
* **deps:** update github/codeql-action action to v3.25.15 ([#284](https://github.com/defenseunicorns/go-oscal/issues/284)) ([462e363](https://github.com/defenseunicorns/go-oscal/commit/462e3636c39819bfc842d03bab16572ce3ab5163))
* **deps:** update github/codeql-action action to v3.26.0 ([#289](https://github.com/defenseunicorns/go-oscal/issues/289)) ([27c7494](https://github.com/defenseunicorns/go-oscal/commit/27c74942664ba43b775730f973536cb324658c07))
* **deps:** update module github.com/swaggest/jsonschema-go to v0.3.72 ([#272](https://github.com/defenseunicorns/go-oscal/issues/272)) ([f61f486](https://github.com/defenseunicorns/go-oscal/commit/f61f486f5c8f4ed003081438994b63a24b377895))
* **deps:** update ossf/scorecard-action action to v2.4.0 ([#285](https://github.com/defenseunicorns/go-oscal/issues/285)) ([2a1e32e](https://github.com/defenseunicorns/go-oscal/commit/2a1e32e82228af59ad39fbb37cb1e7282df3bfdc))

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
