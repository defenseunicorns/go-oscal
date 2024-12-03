# Changelog

## [0.7.0](https://github.com/defenseunicorns/go-oscal/compare/v0.6.2...v0.7.0) (2024-12-03)


### ⚠ BREAKING CHANGES

* **validation:** #274 refactor to accommodate updating to jsonschema v6 ([#282](https://github.com/defenseunicorns/go-oscal/issues/282))
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
* **validation:** [#274](https://github.com/defenseunicorns/go-oscal/issues/274) refactor to accommodate updating to jsonschema v6 ([#282](https://github.com/defenseunicorns/go-oscal/issues/282)) ([5e1a894](https://github.com/defenseunicorns/go-oscal/commit/5e1a8941fab19b04030db994e10d5063d08f05bd))


### Features

* add norm vuln scanning and updating. ([4012ecb](https://github.com/defenseunicorns/go-oscal/commit/4012ecb6401220eee3e1183a54daa50f08778e25))
* Add uuid generation functions for library use ([09d26c5](https://github.com/defenseunicorns/go-oscal/commit/09d26c5ca19a26a239935c14ff63615e464e0247))
* **deps:** update dependency usnistgov/oscal to v1.1.3 ([#331](https://github.com/defenseunicorns/go-oscal/issues/331)) ([37c9d23](https://github.com/defenseunicorns/go-oscal/commit/37c9d2300a3e5aec62ecf5ef848b58cc22237217))
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
* **deps:** update actions/checkout action to v4.2.2 ([#316](https://github.com/defenseunicorns/go-oscal/issues/316)) ([b00c2d7](https://github.com/defenseunicorns/go-oscal/commit/b00c2d7d3b42cba94727f972fa58a0c81353b94e))
* **deps:** update actions/download-artifact action to v4.1.5 ([#194](https://github.com/defenseunicorns/go-oscal/issues/194)) ([78f56a5](https://github.com/defenseunicorns/go-oscal/commit/78f56a536c5ce4d862af0d3e231e2a5e19acff32))
* **deps:** update actions/download-artifact action to v4.1.7 ([#206](https://github.com/defenseunicorns/go-oscal/issues/206)) ([1453390](https://github.com/defenseunicorns/go-oscal/commit/1453390d673da0ac3a4635a5626ec8b8d6983293))
* **deps:** update actions/download-artifact action to v4.1.8 ([#275](https://github.com/defenseunicorns/go-oscal/issues/275)) ([06c015f](https://github.com/defenseunicorns/go-oscal/commit/06c015f994948c030747466bde50d8beb14d4d16))
* **deps:** update actions/github-script digest to 35b1cdd ([#286](https://github.com/defenseunicorns/go-oscal/issues/286)) ([b81dbfb](https://github.com/defenseunicorns/go-oscal/commit/b81dbfbc5ab7326ecdbcf816f3b7c8603b7d4885))
* **deps:** update actions/github-script digest to 4020e46 ([#323](https://github.com/defenseunicorns/go-oscal/issues/323)) ([140f52d](https://github.com/defenseunicorns/go-oscal/commit/140f52d2585d2196d249ae5fbadd2227f51a38a2))
* **deps:** update actions/github-script digest to 58d7008 ([#309](https://github.com/defenseunicorns/go-oscal/issues/309)) ([7b5356b](https://github.com/defenseunicorns/go-oscal/commit/7b5356b5faca82f39a0cc8e740016c8eaa43c4a3))
* **deps:** update actions/github-script digest to 660ec11 ([#310](https://github.com/defenseunicorns/go-oscal/issues/310)) ([d2c288a](https://github.com/defenseunicorns/go-oscal/commit/d2c288a0f81b21f7dafd24e97e2430ab38bde1f8))
* **deps:** update actions/setup-go action to v5.0.1 ([#212](https://github.com/defenseunicorns/go-oscal/issues/212)) ([f978b2d](https://github.com/defenseunicorns/go-oscal/commit/f978b2d5cc83a4505aa5291d492ea73039e548bf))
* **deps:** update actions/setup-go action to v5.0.2 ([#279](https://github.com/defenseunicorns/go-oscal/issues/279)) ([40b2efe](https://github.com/defenseunicorns/go-oscal/commit/40b2efe131c04c3ff121a19bc2647bcda3c8c01c))
* **deps:** update actions/setup-go action to v5.1.0 ([#320](https://github.com/defenseunicorns/go-oscal/issues/320)) ([27d4a9d](https://github.com/defenseunicorns/go-oscal/commit/27d4a9d93b8a4ccc8d2eacb90f3f599d9d907152))
* **deps:** update actions/setup-node action to v4 ([#203](https://github.com/defenseunicorns/go-oscal/issues/203)) ([45aecb3](https://github.com/defenseunicorns/go-oscal/commit/45aecb3bfcb6de79e105b429f0b3be548909b55e))
* **deps:** update actions/setup-node action to v4.0.3 ([#276](https://github.com/defenseunicorns/go-oscal/issues/276)) ([54f6ab5](https://github.com/defenseunicorns/go-oscal/commit/54f6ab5a132982dcbe092c6da7ff8881f045f7f5))
* **deps:** update actions/setup-node action to v4.0.4 ([#313](https://github.com/defenseunicorns/go-oscal/issues/313)) ([9ee6ce7](https://github.com/defenseunicorns/go-oscal/commit/9ee6ce70e7c40e77a89cdfc236f0e45396b0cb72))
* **deps:** update actions/setup-node action to v4.1.0 ([#321](https://github.com/defenseunicorns/go-oscal/issues/321)) ([3201da1](https://github.com/defenseunicorns/go-oscal/commit/3201da1338087642b59c4ef49dc1d3a783f06cb2))
* **deps:** update actions/upload-artifact action to v4.3.2 ([#197](https://github.com/defenseunicorns/go-oscal/issues/197)) ([898b46f](https://github.com/defenseunicorns/go-oscal/commit/898b46f4a5d3ad895bdc6a43496e23633d098c58))
* **deps:** update actions/upload-artifact action to v4.3.3 ([#207](https://github.com/defenseunicorns/go-oscal/issues/207)) ([b9093f7](https://github.com/defenseunicorns/go-oscal/commit/b9093f78477b196bd29e1d4621f0a7180fa7a9a7))
* **deps:** update actions/upload-artifact action to v4.3.4 ([#277](https://github.com/defenseunicorns/go-oscal/issues/277)) ([c149bc7](https://github.com/defenseunicorns/go-oscal/commit/c149bc7e0757c0e4f5a577435541fcca1b039543))
* **deps:** update actions/upload-artifact action to v4.3.6 ([#287](https://github.com/defenseunicorns/go-oscal/issues/287)) ([31eb0c5](https://github.com/defenseunicorns/go-oscal/commit/31eb0c5175edbd235237561e58992c33f891057e))
* **deps:** update actions/upload-artifact action to v4.4.0 ([#304](https://github.com/defenseunicorns/go-oscal/issues/304)) ([f43c01d](https://github.com/defenseunicorns/go-oscal/commit/f43c01d9e3cbed0c4b1717d9276fe23df102a88a))
* **deps:** update actions/upload-artifact action to v4.4.3 ([#317](https://github.com/defenseunicorns/go-oscal/issues/317)) ([cb8f4f5](https://github.com/defenseunicorns/go-oscal/commit/cb8f4f55e07d11259fd2e442d1fcffc6e0154e96))
* **deps:** update anchore/sbom-action action to v0.15.10 ([#199](https://github.com/defenseunicorns/go-oscal/issues/199)) ([32acfd7](https://github.com/defenseunicorns/go-oscal/commit/32acfd767a2f9d7217d1819d11526f3d7e54d012))
* **deps:** update anchore/sbom-action action to v0.15.11 ([#211](https://github.com/defenseunicorns/go-oscal/issues/211)) ([aec5c41](https://github.com/defenseunicorns/go-oscal/commit/aec5c41def73b31a923968d72ffad035f406a811))
* **deps:** update anchore/sbom-action action to v0.16.0 ([#229](https://github.com/defenseunicorns/go-oscal/issues/229)) ([0c00239](https://github.com/defenseunicorns/go-oscal/commit/0c002395f54fbf7a01f96de197e6006b86ce3152))
* **deps:** update anchore/sbom-action action to v0.16.1 ([#278](https://github.com/defenseunicorns/go-oscal/issues/278)) ([ad3e6aa](https://github.com/defenseunicorns/go-oscal/commit/ad3e6aae883ef6a819bcc25c8aa635e9175f2ab6))
* **deps:** update anchore/sbom-action action to v0.17.0 ([#281](https://github.com/defenseunicorns/go-oscal/issues/281)) ([5000d31](https://github.com/defenseunicorns/go-oscal/commit/5000d311d277b2d918e84f0a556108823eadd80f))
* **deps:** update anchore/sbom-action action to v0.17.1 ([#292](https://github.com/defenseunicorns/go-oscal/issues/292)) ([c917c71](https://github.com/defenseunicorns/go-oscal/commit/c917c7183ddd58f28336dd5f095833686d54eb6e))
* **deps:** update anchore/sbom-action action to v0.17.2 ([#298](https://github.com/defenseunicorns/go-oscal/issues/298)) ([2614471](https://github.com/defenseunicorns/go-oscal/commit/2614471e55bd3eb17a0c1b1b547ef94918ca68c1))
* **deps:** update anchore/sbom-action action to v0.17.5 ([#318](https://github.com/defenseunicorns/go-oscal/issues/318)) ([08686bd](https://github.com/defenseunicorns/go-oscal/commit/08686bdbbd554cc2a43073cf65d2ef6dd1422868))
* **deps:** update anchore/sbom-action action to v0.17.6 ([#322](https://github.com/defenseunicorns/go-oscal/issues/322)) ([008abdc](https://github.com/defenseunicorns/go-oscal/commit/008abdc721d06e0d934d604e9d01f757e27ec61b))
* **deps:** update anchore/sbom-action action to v0.17.7 ([#324](https://github.com/defenseunicorns/go-oscal/issues/324)) ([f5087af](https://github.com/defenseunicorns/go-oscal/commit/f5087afe82546b35a375879c2e610e43713c69bb))
* **deps:** update anchore/sbom-action action to v0.17.8 ([#330](https://github.com/defenseunicorns/go-oscal/issues/330)) ([e6ce561](https://github.com/defenseunicorns/go-oscal/commit/e6ce5617fef9939725ce5bf938ec952117985b2f))
* **deps:** update dependency @commitlint/config-conventional to v19.4.1 ([#300](https://github.com/defenseunicorns/go-oscal/issues/300)) ([3261916](https://github.com/defenseunicorns/go-oscal/commit/32619160c3a4d34e42f4b1c9b35d5f4a0eb24137))
* **deps:** update dependency @commitlint/config-conventional to v19.5.0 ([#306](https://github.com/defenseunicorns/go-oscal/issues/306)) ([9744bc4](https://github.com/defenseunicorns/go-oscal/commit/9744bc45788f3d9fc0444a4acd428a643e1b959c))
* **deps:** update dependency @commitlint/config-conventional to v19.6.0 ([#327](https://github.com/defenseunicorns/go-oscal/issues/327)) ([f2ed44e](https://github.com/defenseunicorns/go-oscal/commit/f2ed44e77dfb20ec6d6e9b8323e4691fcb913aca))
* **deps:** update dependency commitlint to v19.4.0 ([#290](https://github.com/defenseunicorns/go-oscal/issues/290)) ([d2351fd](https://github.com/defenseunicorns/go-oscal/commit/d2351fd7504d9eeddd2068b323e2dcd59222a1b3))
* **deps:** update dependency commitlint to v19.4.1 ([#299](https://github.com/defenseunicorns/go-oscal/issues/299)) ([1c8da28](https://github.com/defenseunicorns/go-oscal/commit/1c8da2852a0cda838e7cfa62965471f2340aa931))
* **deps:** update dependency commitlint to v19.5.0 ([#307](https://github.com/defenseunicorns/go-oscal/issues/307)) ([f5b95fa](https://github.com/defenseunicorns/go-oscal/commit/f5b95fa10d8373b9924872cbae83583a1ade83ef))
* **deps:** update dependency commitlint to v19.6.0 ([#328](https://github.com/defenseunicorns/go-oscal/issues/328)) ([8e518c5](https://github.com/defenseunicorns/go-oscal/commit/8e518c5cc51f22fa5a9bbbb97fba72b73a268147))
* **deps:** update dependency go to v1.22.2 ([#200](https://github.com/defenseunicorns/go-oscal/issues/200)) ([d48fa51](https://github.com/defenseunicorns/go-oscal/commit/d48fa51714946ce870f672dc7a9327863bacfec3))
* **deps:** update github/codeql-action action to v3.25.1 ([#202](https://github.com/defenseunicorns/go-oscal/issues/202)) ([25ebdef](https://github.com/defenseunicorns/go-oscal/commit/25ebdef38e875a88fd6b81a6c6e4a18471ed30e8))
* **deps:** update github/codeql-action action to v3.25.10 ([#258](https://github.com/defenseunicorns/go-oscal/issues/258)) ([004addd](https://github.com/defenseunicorns/go-oscal/commit/004addd22026d4cba26a0eb1dabc09698d412ef3))
* **deps:** update github/codeql-action action to v3.25.11 ([#273](https://github.com/defenseunicorns/go-oscal/issues/273)) ([6910af5](https://github.com/defenseunicorns/go-oscal/commit/6910af5c5e4d0d151b3e07992bdaa9baf69549dd))
* **deps:** update github/codeql-action action to v3.25.12 ([#280](https://github.com/defenseunicorns/go-oscal/issues/280)) ([147c634](https://github.com/defenseunicorns/go-oscal/commit/147c6340b3a2a273c9909e68b698101ac82a69b9))
* **deps:** update github/codeql-action action to v3.25.13 ([#283](https://github.com/defenseunicorns/go-oscal/issues/283)) ([e7dc32a](https://github.com/defenseunicorns/go-oscal/commit/e7dc32a6830213af63c0be625897ccaf2db91e44))
* **deps:** update github/codeql-action action to v3.25.15 ([#284](https://github.com/defenseunicorns/go-oscal/issues/284)) ([462e363](https://github.com/defenseunicorns/go-oscal/commit/462e3636c39819bfc842d03bab16572ce3ab5163))
* **deps:** update github/codeql-action action to v3.25.2 ([#208](https://github.com/defenseunicorns/go-oscal/issues/208)) ([412b27a](https://github.com/defenseunicorns/go-oscal/commit/412b27ac842d05c25b888e4767ffb7d9466946e3))
* **deps:** update github/codeql-action action to v3.25.3 ([#210](https://github.com/defenseunicorns/go-oscal/issues/210)) ([b07d331](https://github.com/defenseunicorns/go-oscal/commit/b07d331e1744fa66dc17b1a7d0ed6b1475c58947))
* **deps:** update github/codeql-action action to v3.25.4 ([#213](https://github.com/defenseunicorns/go-oscal/issues/213)) ([ad3cf0b](https://github.com/defenseunicorns/go-oscal/commit/ad3cf0b9f62b188ff1aef927e75935a49f54b4a3))
* **deps:** update github/codeql-action action to v3.25.5 ([#218](https://github.com/defenseunicorns/go-oscal/issues/218)) ([b0fbd94](https://github.com/defenseunicorns/go-oscal/commit/b0fbd94452a75de2ac3300f039fa41613145b3e6))
* **deps:** update github/codeql-action action to v3.25.6 ([#228](https://github.com/defenseunicorns/go-oscal/issues/228)) ([d95a749](https://github.com/defenseunicorns/go-oscal/commit/d95a749a904a565a28243dc5183938d40424d6fe))
* **deps:** update github/codeql-action action to v3.25.7 ([#244](https://github.com/defenseunicorns/go-oscal/issues/244)) ([a49ede0](https://github.com/defenseunicorns/go-oscal/commit/a49ede08aaac118c4060cb5ea2323748cb94540d))
* **deps:** update github/codeql-action action to v3.25.8 ([#247](https://github.com/defenseunicorns/go-oscal/issues/247)) ([5fce8ec](https://github.com/defenseunicorns/go-oscal/commit/5fce8ec71aaaee59a4f46261ba7dfd14743882d8))
* **deps:** update github/codeql-action action to v3.26.0 ([#289](https://github.com/defenseunicorns/go-oscal/issues/289)) ([27c7494](https://github.com/defenseunicorns/go-oscal/commit/27c74942664ba43b775730f973536cb324658c07))
* **deps:** update github/codeql-action action to v3.26.1 ([#294](https://github.com/defenseunicorns/go-oscal/issues/294)) ([5499248](https://github.com/defenseunicorns/go-oscal/commit/5499248009134e4b1de19e4d6fbcbc7a11e94e8b))
* **deps:** update github/codeql-action action to v3.26.12 ([#315](https://github.com/defenseunicorns/go-oscal/issues/315)) ([f072f80](https://github.com/defenseunicorns/go-oscal/commit/f072f803f40e6317cb9db5354f2aa1125814b0dc))
* **deps:** update github/codeql-action action to v3.26.2 ([#296](https://github.com/defenseunicorns/go-oscal/issues/296)) ([32c1368](https://github.com/defenseunicorns/go-oscal/commit/32c1368f9302fe62b8e3a3c7a8b7009789e2df99))
* **deps:** update github/codeql-action action to v3.26.5 ([#297](https://github.com/defenseunicorns/go-oscal/issues/297)) ([b6d55a9](https://github.com/defenseunicorns/go-oscal/commit/b6d55a9b9db95fadef0ea23564315e307a2ae511))
* **deps:** update github/codeql-action action to v3.26.6 ([#301](https://github.com/defenseunicorns/go-oscal/issues/301)) ([7bf60c3](https://github.com/defenseunicorns/go-oscal/commit/7bf60c313e5f13f2682290de5eaefff4c3105d95))
* **deps:** update github/codeql-action action to v3.26.7 ([#308](https://github.com/defenseunicorns/go-oscal/issues/308)) ([38600bd](https://github.com/defenseunicorns/go-oscal/commit/38600bd8e12c06bd3df62dc40b9ba80dbf0fbf78))
* **deps:** update github/codeql-action action to v3.26.8 ([#314](https://github.com/defenseunicorns/go-oscal/issues/314)) ([2c34a83](https://github.com/defenseunicorns/go-oscal/commit/2c34a8332912e9b429f6a116692435431e5fcc4b))
* **deps:** update github/codeql-action action to v3.27.0 ([#319](https://github.com/defenseunicorns/go-oscal/issues/319)) ([0206e06](https://github.com/defenseunicorns/go-oscal/commit/0206e06ba9424be4219a286bf81393361f720e84))
* **deps:** update github/codeql-action action to v3.27.5 ([#326](https://github.com/defenseunicorns/go-oscal/issues/326)) ([21975f5](https://github.com/defenseunicorns/go-oscal/commit/21975f526d1a69834ff20e3e051d192068ec7d29))
* **deps:** update github/codeql-action action to v3.27.6 ([#333](https://github.com/defenseunicorns/go-oscal/issues/333)) ([41fbd9c](https://github.com/defenseunicorns/go-oscal/commit/41fbd9cd5126716b5fd3786d6b2fa19090e8371c))
* **deps:** update go version to 1.22.3 ([#226](https://github.com/defenseunicorns/go-oscal/issues/226)) ([ad903f8](https://github.com/defenseunicorns/go-oscal/commit/ad903f8a726be8d253ba321cea26b33992a3f29c))
* **deps:** update googleapis/release-please-action digest to 7987652 ([#257](https://github.com/defenseunicorns/go-oscal/issues/257)) ([71db26b](https://github.com/defenseunicorns/go-oscal/commit/71db26b5b56be2c6e8609a4584728359d1276135))
* **deps:** update googleapis/release-please-action digest to f3969c0 ([#222](https://github.com/defenseunicorns/go-oscal/issues/222)) ([8f8302e](https://github.com/defenseunicorns/go-oscal/commit/8f8302eab0710d7a6d6e6174bfcbd49f5dfb1c1b))
* **deps:** update goreleaser/goreleaser-action action to v5.1.0 ([#216](https://github.com/defenseunicorns/go-oscal/issues/216)) ([03728e3](https://github.com/defenseunicorns/go-oscal/commit/03728e3e1625e4220c6e6b84e8b8ea398309090b))
* **deps:** update goreleaser/goreleaser-action action to v6 ([#249](https://github.com/defenseunicorns/go-oscal/issues/249)) ([9886d60](https://github.com/defenseunicorns/go-oscal/commit/9886d6024d88bae2f5968f8ac978a006b46a1513))
* **deps:** update goreleaser/goreleaser-action action to v6.1.0 ([#325](https://github.com/defenseunicorns/go-oscal/issues/325)) ([4c4b417](https://github.com/defenseunicorns/go-oscal/commit/4c4b417fd885bf4b1cb33aeee8a977d63ba90424))
* **deps:** update module github.com/santhosh-tekuri/jsonschema/v5 to v6 ([#248](https://github.com/defenseunicorns/go-oscal/issues/248)) ([456d3a8](https://github.com/defenseunicorns/go-oscal/commit/456d3a870b587a1c59ffdb0d578e154de23a00ed))
* **deps:** update module github.com/santhosh-tekuri/jsonschema/v5 to v6 ([#251](https://github.com/defenseunicorns/go-oscal/issues/251)) ([a26a473](https://github.com/defenseunicorns/go-oscal/commit/a26a473b438b5bcfeceef3d4c39fe0bee2b853a0))
* **deps:** update module github.com/spf13/cobra to v1.8.1 ([#264](https://github.com/defenseunicorns/go-oscal/issues/264)) ([b05ed96](https://github.com/defenseunicorns/go-oscal/commit/b05ed96813fe321dc264e56f3155c9cdf373899f))
* **deps:** update module github.com/swaggest/jsonschema-go to v0.3.70 ([#201](https://github.com/defenseunicorns/go-oscal/issues/201)) ([5a06269](https://github.com/defenseunicorns/go-oscal/commit/5a06269b2eb2b61b83d00b20c355d6a5a39800de))
* **deps:** update module github.com/swaggest/jsonschema-go to v0.3.72 ([#272](https://github.com/defenseunicorns/go-oscal/issues/272)) ([f61f486](https://github.com/defenseunicorns/go-oscal/commit/f61f486f5c8f4ed003081438994b63a24b377895))
* **deps:** update ossf/scorecard-action action to v2.3.3 ([#215](https://github.com/defenseunicorns/go-oscal/issues/215)) ([33cbc44](https://github.com/defenseunicorns/go-oscal/commit/33cbc4452be5a73cc64e340b5b680add5aa75b99))
* **deps:** update ossf/scorecard-action action to v2.4.0 ([#285](https://github.com/defenseunicorns/go-oscal/issues/285)) ([2a1e32e](https://github.com/defenseunicorns/go-oscal/commit/2a1e32e82228af59ad39fbb37cb1e7282df3bfdc))
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
* **main:** release 0.6.0 ([#270](https://github.com/defenseunicorns/go-oscal/issues/270)) ([e18aae8](https://github.com/defenseunicorns/go-oscal/commit/e18aae8ff543369f4e9c24f31bb6110b6955790b))
* **main:** release 0.6.1 ([#291](https://github.com/defenseunicorns/go-oscal/issues/291)) ([35c9603](https://github.com/defenseunicorns/go-oscal/commit/35c9603a5559a7c296a1353431bb29e57d2b0f69))
* **main:** release 0.6.2 ([#329](https://github.com/defenseunicorns/go-oscal/issues/329)) ([e7a835e](https://github.com/defenseunicorns/go-oscal/commit/e7a835e26b4cbc38f0ba858c9f81df6f98e91e8c))
* **release-please-action:** [#178](https://github.com/defenseunicorns/go-oscal/issues/178) fix release-please-config for v4 ([#179](https://github.com/defenseunicorns/go-oscal/issues/179)) ([c9d746f](https://github.com/defenseunicorns/go-oscal/commit/c9d746f3d7508058c461bd528d4052e517d32a6a))
* **release-please:** update config with draft set to true for consistency with go-releaser ([#182](https://github.com/defenseunicorns/go-oscal/issues/182)) ([7d23ba5](https://github.com/defenseunicorns/go-oscal/commit/7d23ba558b05458b6d143b91542336255192a7cc))
* **renovate:** remove schedule ([#293](https://github.com/defenseunicorns/go-oscal/issues/293)) ([514a7ae](https://github.com/defenseunicorns/go-oscal/commit/514a7aeaeefc6ba9da44c639e8dad90efe15bf63))
* **workflows:** add triage label workflow ([#184](https://github.com/defenseunicorns/go-oscal/issues/184)) ([280247a](https://github.com/defenseunicorns/go-oscal/commit/280247ac32c0c66abfe41219257af751df6f6009))

## [0.6.2](https://github.com/defenseunicorns/go-oscal/compare/v0.6.1...v0.6.2) (2024-12-03)


### Features

* **deps:** update dependency usnistgov/oscal to v1.1.3 ([#331](https://github.com/defenseunicorns/go-oscal/issues/331)) ([37c9d23](https://github.com/defenseunicorns/go-oscal/commit/37c9d2300a3e5aec62ecf5ef848b58cc22237217))


### Miscellaneous

* **deps:** update anchore/sbom-action action to v0.17.8 ([#330](https://github.com/defenseunicorns/go-oscal/issues/330)) ([e6ce561](https://github.com/defenseunicorns/go-oscal/commit/e6ce5617fef9939725ce5bf938ec952117985b2f))
* **deps:** update github/codeql-action action to v3.27.6 ([#333](https://github.com/defenseunicorns/go-oscal/issues/333)) ([41fbd9c](https://github.com/defenseunicorns/go-oscal/commit/41fbd9cd5126716b5fd3786d6b2fa19090e8371c))

## [0.6.1](https://github.com/defenseunicorns/go-oscal/compare/v0.6.0...v0.6.1) (2024-11-21)


### Miscellaneous

* **deps:** update actions/checkout action to v4.2.2 ([#316](https://github.com/defenseunicorns/go-oscal/issues/316)) ([b00c2d7](https://github.com/defenseunicorns/go-oscal/commit/b00c2d7d3b42cba94727f972fa58a0c81353b94e))
* **deps:** update actions/github-script digest to 4020e46 ([#323](https://github.com/defenseunicorns/go-oscal/issues/323)) ([140f52d](https://github.com/defenseunicorns/go-oscal/commit/140f52d2585d2196d249ae5fbadd2227f51a38a2))
* **deps:** update actions/github-script digest to 58d7008 ([#309](https://github.com/defenseunicorns/go-oscal/issues/309)) ([7b5356b](https://github.com/defenseunicorns/go-oscal/commit/7b5356b5faca82f39a0cc8e740016c8eaa43c4a3))
* **deps:** update actions/github-script digest to 660ec11 ([#310](https://github.com/defenseunicorns/go-oscal/issues/310)) ([d2c288a](https://github.com/defenseunicorns/go-oscal/commit/d2c288a0f81b21f7dafd24e97e2430ab38bde1f8))
* **deps:** update actions/setup-go action to v5.1.0 ([#320](https://github.com/defenseunicorns/go-oscal/issues/320)) ([27d4a9d](https://github.com/defenseunicorns/go-oscal/commit/27d4a9d93b8a4ccc8d2eacb90f3f599d9d907152))
* **deps:** update actions/setup-node action to v4.0.4 ([#313](https://github.com/defenseunicorns/go-oscal/issues/313)) ([9ee6ce7](https://github.com/defenseunicorns/go-oscal/commit/9ee6ce70e7c40e77a89cdfc236f0e45396b0cb72))
* **deps:** update actions/setup-node action to v4.1.0 ([#321](https://github.com/defenseunicorns/go-oscal/issues/321)) ([3201da1](https://github.com/defenseunicorns/go-oscal/commit/3201da1338087642b59c4ef49dc1d3a783f06cb2))
* **deps:** update actions/upload-artifact action to v4.4.0 ([#304](https://github.com/defenseunicorns/go-oscal/issues/304)) ([f43c01d](https://github.com/defenseunicorns/go-oscal/commit/f43c01d9e3cbed0c4b1717d9276fe23df102a88a))
* **deps:** update actions/upload-artifact action to v4.4.3 ([#317](https://github.com/defenseunicorns/go-oscal/issues/317)) ([cb8f4f5](https://github.com/defenseunicorns/go-oscal/commit/cb8f4f55e07d11259fd2e442d1fcffc6e0154e96))
* **deps:** update anchore/sbom-action action to v0.17.1 ([#292](https://github.com/defenseunicorns/go-oscal/issues/292)) ([c917c71](https://github.com/defenseunicorns/go-oscal/commit/c917c7183ddd58f28336dd5f095833686d54eb6e))
* **deps:** update anchore/sbom-action action to v0.17.2 ([#298](https://github.com/defenseunicorns/go-oscal/issues/298)) ([2614471](https://github.com/defenseunicorns/go-oscal/commit/2614471e55bd3eb17a0c1b1b547ef94918ca68c1))
* **deps:** update anchore/sbom-action action to v0.17.5 ([#318](https://github.com/defenseunicorns/go-oscal/issues/318)) ([08686bd](https://github.com/defenseunicorns/go-oscal/commit/08686bdbbd554cc2a43073cf65d2ef6dd1422868))
* **deps:** update anchore/sbom-action action to v0.17.6 ([#322](https://github.com/defenseunicorns/go-oscal/issues/322)) ([008abdc](https://github.com/defenseunicorns/go-oscal/commit/008abdc721d06e0d934d604e9d01f757e27ec61b))
* **deps:** update anchore/sbom-action action to v0.17.7 ([#324](https://github.com/defenseunicorns/go-oscal/issues/324)) ([f5087af](https://github.com/defenseunicorns/go-oscal/commit/f5087afe82546b35a375879c2e610e43713c69bb))
* **deps:** update dependency @commitlint/config-conventional to v19.4.1 ([#300](https://github.com/defenseunicorns/go-oscal/issues/300)) ([3261916](https://github.com/defenseunicorns/go-oscal/commit/32619160c3a4d34e42f4b1c9b35d5f4a0eb24137))
* **deps:** update dependency @commitlint/config-conventional to v19.5.0 ([#306](https://github.com/defenseunicorns/go-oscal/issues/306)) ([9744bc4](https://github.com/defenseunicorns/go-oscal/commit/9744bc45788f3d9fc0444a4acd428a643e1b959c))
* **deps:** update dependency @commitlint/config-conventional to v19.6.0 ([#327](https://github.com/defenseunicorns/go-oscal/issues/327)) ([f2ed44e](https://github.com/defenseunicorns/go-oscal/commit/f2ed44e77dfb20ec6d6e9b8323e4691fcb913aca))
* **deps:** update dependency commitlint to v19.4.1 ([#299](https://github.com/defenseunicorns/go-oscal/issues/299)) ([1c8da28](https://github.com/defenseunicorns/go-oscal/commit/1c8da2852a0cda838e7cfa62965471f2340aa931))
* **deps:** update dependency commitlint to v19.5.0 ([#307](https://github.com/defenseunicorns/go-oscal/issues/307)) ([f5b95fa](https://github.com/defenseunicorns/go-oscal/commit/f5b95fa10d8373b9924872cbae83583a1ade83ef))
* **deps:** update dependency commitlint to v19.6.0 ([#328](https://github.com/defenseunicorns/go-oscal/issues/328)) ([8e518c5](https://github.com/defenseunicorns/go-oscal/commit/8e518c5cc51f22fa5a9bbbb97fba72b73a268147))
* **deps:** update github/codeql-action action to v3.26.1 ([#294](https://github.com/defenseunicorns/go-oscal/issues/294)) ([5499248](https://github.com/defenseunicorns/go-oscal/commit/5499248009134e4b1de19e4d6fbcbc7a11e94e8b))
* **deps:** update github/codeql-action action to v3.26.12 ([#315](https://github.com/defenseunicorns/go-oscal/issues/315)) ([f072f80](https://github.com/defenseunicorns/go-oscal/commit/f072f803f40e6317cb9db5354f2aa1125814b0dc))
* **deps:** update github/codeql-action action to v3.26.2 ([#296](https://github.com/defenseunicorns/go-oscal/issues/296)) ([32c1368](https://github.com/defenseunicorns/go-oscal/commit/32c1368f9302fe62b8e3a3c7a8b7009789e2df99))
* **deps:** update github/codeql-action action to v3.26.5 ([#297](https://github.com/defenseunicorns/go-oscal/issues/297)) ([b6d55a9](https://github.com/defenseunicorns/go-oscal/commit/b6d55a9b9db95fadef0ea23564315e307a2ae511))
* **deps:** update github/codeql-action action to v3.26.6 ([#301](https://github.com/defenseunicorns/go-oscal/issues/301)) ([7bf60c3](https://github.com/defenseunicorns/go-oscal/commit/7bf60c313e5f13f2682290de5eaefff4c3105d95))
* **deps:** update github/codeql-action action to v3.26.7 ([#308](https://github.com/defenseunicorns/go-oscal/issues/308)) ([38600bd](https://github.com/defenseunicorns/go-oscal/commit/38600bd8e12c06bd3df62dc40b9ba80dbf0fbf78))
* **deps:** update github/codeql-action action to v3.26.8 ([#314](https://github.com/defenseunicorns/go-oscal/issues/314)) ([2c34a83](https://github.com/defenseunicorns/go-oscal/commit/2c34a8332912e9b429f6a116692435431e5fcc4b))
* **deps:** update github/codeql-action action to v3.27.0 ([#319](https://github.com/defenseunicorns/go-oscal/issues/319)) ([0206e06](https://github.com/defenseunicorns/go-oscal/commit/0206e06ba9424be4219a286bf81393361f720e84))
* **deps:** update github/codeql-action action to v3.27.5 ([#326](https://github.com/defenseunicorns/go-oscal/issues/326)) ([21975f5](https://github.com/defenseunicorns/go-oscal/commit/21975f526d1a69834ff20e3e051d192068ec7d29))
* **deps:** update goreleaser/goreleaser-action action to v6.1.0 ([#325](https://github.com/defenseunicorns/go-oscal/issues/325)) ([4c4b417](https://github.com/defenseunicorns/go-oscal/commit/4c4b417fd885bf4b1cb33aeee8a977d63ba90424))
* **renovate:** remove schedule ([#293](https://github.com/defenseunicorns/go-oscal/issues/293)) ([514a7ae](https://github.com/defenseunicorns/go-oscal/commit/514a7aeaeefc6ba9da44c639e8dad90efe15bf63))

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
