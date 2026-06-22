# Changelog

## 0.14.0 (2026-06-22)

Full Changelog:
[v0.13.0...v0.14.0](https://github.com/Nimbleway/nimble-cli/compare/v0.13.0...v0.14.0)

### Features

- **api:** manual updates
  ([77d88f1](https://github.com/Nimbleway/nimble-cli/commit/77d88f15da069f3f39de3347d31b24c25cdd9631))

## 0.13.0 (2026-06-19)

Full Changelog:
[v0.12.0...v0.13.0](https://github.com/Nimbleway/nimble-cli/compare/v0.12.0...v0.13.0)

### Features

- **api:** Add jobs api
  ([d69d0ac](https://github.com/Nimbleway/nimble-cli/commit/d69d0ac34caa7e0567a8fd1914a0fe19c495d5fa))
- **api:** api update
  ([13d4adb](https://github.com/Nimbleway/nimble-cli/commit/13d4adb42b2196d8a4727792d4f60d065c926ca7))
- **api:** manually update openapi
  ([f78716a](https://github.com/Nimbleway/nimble-cli/commit/f78716af8d89d42fde571878d15d59d7513af0b7))
- **api:** Manually update OpenAPI spec
  ([b8c11fc](https://github.com/Nimbleway/nimble-cli/commit/b8c11fccd028c25113e039e934a54f6512200e6d))

### Chores

- **internal:** codegen related update
  ([b179733](https://github.com/Nimbleway/nimble-cli/commit/b179733b9e92e06ed546255cde6928a0aa60aea7))
- **internal:** codegen related update
  ([cf90407](https://github.com/Nimbleway/nimble-cli/commit/cf90407de9f8723cee73c2ee1e4a14a2ed391d00))

## 0.12.0 (2026-05-10)

Full Changelog:
[v0.11.0...v0.12.0](https://github.com/Nimbleway/nimble-cli/compare/v0.11.0...v0.12.0)

### Features

- **api:** Add media, serp and domain_knowledge
  ([c89cd69](https://github.com/Nimbleway/nimble-cli/commit/c89cd69547c488b70c25dfe4e26af10d07a827d7))

### Chores

- redact api-key headers in debug logs
  ([b9ad508](https://github.com/Nimbleway/nimble-cli/commit/b9ad5087d15ea17e2f08fe2e335da9f1a4a25d35))
- **test:** scope body requests to a single variant
  ([0992f16](https://github.com/Nimbleway/nimble-cli/commit/0992f1655b412298aea24b2602ad2d91f48ba7cb))
- update SDK settings
  ([586bfde](https://github.com/Nimbleway/nimble-cli/commit/586bfde274fd1dd12b59d3c236d2f458fd409594))

## 0.11.0 (2026-05-06)

Full Changelog:
[v0.10.0...v0.11.0](https://github.com/Nimbleway/nimble-cli/compare/v0.10.0...v0.11.0)

### Features

- allow `-` as value representing stdin to binary-only file parameters in CLIs
  ([948c073](https://github.com/Nimbleway/nimble-cli/commit/948c0730a137295bbd571e68ab9c5b4bb5ae067d))
- **api:** add client-source/FileInput params, base URL validation, remove npm
  packaging
  ([d87a6bb](https://github.com/Nimbleway/nimble-cli/commit/d87a6bb87e56dc3d4b4f75ff1b5b4bd20a636cdf))
- **api:** api update
  ([7a8ca4f](https://github.com/Nimbleway/nimble-cli/commit/7a8ca4fa1b0cca2daa6b7b6c2d7337dac5bcf450))
- **api:** api update
  ([5217aed](https://github.com/Nimbleway/nimble-cli/commit/5217aed5e2991d9b5cba57c4dfe9cf245049a778))
- **api:** api update
  ([e4ea036](https://github.com/Nimbleway/nimble-cli/commit/e4ea036448daf768bd9742b5d148273f01632779))
- **api:** api update
  ([7e68a7e](https://github.com/Nimbleway/nimble-cli/commit/7e68a7e359d5f58e673a9980a4f78200dc4c9422))
- better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url`
  ([5a9fe42](https://github.com/Nimbleway/nimble-cli/commit/5a9fe42ec4646123e94c8ff68c4a4209324f883d))
- binary-only parameters become CLI flags that take filenames only
  ([3325ed0](https://github.com/Nimbleway/nimble-cli/commit/3325ed037dd83b3d936f778e148a846e9b673e29))
- **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings
  ([a5da1f1](https://github.com/Nimbleway/nimble-cli/commit/a5da1f1f73c9100c788cf57ffa500f369d9b5fe3))
- **cli:** alias parameters in data with `x-stainless-cli-data-alias`
  ([13abd47](https://github.com/Nimbleway/nimble-cli/commit/13abd47a9444ff51f88bcb2d33e15ad37498bda5))
- **cli:** send filename and content type when reading input from files
  ([de52976](https://github.com/Nimbleway/nimble-cli/commit/de52976d505b7f57e12b2fbc457b1bab192f862a))
- support passing path and query params over stdin
  ([c09ce9c](https://github.com/Nimbleway/nimble-cli/commit/c09ce9cbbfda60f54791fbec42a697e0dd6b1fbb))

### Bug Fixes

- **cli:** correctly load zsh autocompletion
  ([0639155](https://github.com/Nimbleway/nimble-cli/commit/0639155aa512ef0aff69671517b9370722ca0a84))
- **cli:** fix incompatible Go types for flag generated as array of maps
  ([586c83c](https://github.com/Nimbleway/nimble-cli/commit/586c83cd5463815c163c18b867b9848c468ef6bd))
- fall back to main branch if linking fails in CI
  ([931cd09](https://github.com/Nimbleway/nimble-cli/commit/931cd096320383928a7b2ea139caa19a569bb605))
- fix for failing to drop invalid module replace in link script
  ([b11366c](https://github.com/Nimbleway/nimble-cli/commit/b11366c187001aa38629f54c8a56a5bb1b4c6437))
- fix quoting typo
  ([f74d3aa](https://github.com/Nimbleway/nimble-cli/commit/f74d3aa0520c1f45874317dadb90714b9925ed55))
- flags for nullable body scalar fields are strictly typed
  ([1240867](https://github.com/Nimbleway/nimble-cli/commit/1240867c4feef6e272a7700ee4d66b7f3f529bf1))
- handle empty data set using `--format explore`
  ([8b27aed](https://github.com/Nimbleway/nimble-cli/commit/8b27aed2080198a36026ec6a3b9ef95958db1aea))
- use `RawJSON` when iterating items with `--format explore` in the CLI
  ([ee0ee93](https://github.com/Nimbleway/nimble-cli/commit/ee0ee93a20b646d87845b559f38e04c7eafc01a3))

### Chores

- add documentation for ./scripts/link
  ([06d8b9b](https://github.com/Nimbleway/nimble-cli/commit/06d8b9b607a0baa31efa3e7767d6333e4ad88033))
- **ci:** support manually triggering release workflow
  ([8bf74b9](https://github.com/Nimbleway/nimble-cli/commit/8bf74b90e6f29c0d3811a79d6b45c7b8cd956ea7))
- **cli:** additional test cases for `ShowJSONIterator`
  ([454e8a5](https://github.com/Nimbleway/nimble-cli/commit/454e8a52c5b6afc1adf888c02a6ac9e31da34ed3))
- **cli:** fall back to JSON when using default "explore" with non-TTY
  ([514ace9](https://github.com/Nimbleway/nimble-cli/commit/514ace9b60c92a48fd1aafbebe3544a63a1c882e))
- **cli:** let `--format raw` be used in conjunction with `--transform`
  ([a19a3ce](https://github.com/Nimbleway/nimble-cli/commit/a19a3ce63db8f2623fa77ff5f4624523f126e62b))
- **cli:** switch long lists of positional args over to param structs
  ([9d5c596](https://github.com/Nimbleway/nimble-cli/commit/9d5c596745a7167dac6d819dd825fcbe1ea72cf3))
- **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many
  positionals
  ([86b29c5](https://github.com/Nimbleway/nimble-cli/commit/86b29c56402d82c2e682657953c828160bba2fc8))
- **internal:** more robust bootstrap script
  ([468f6b4](https://github.com/Nimbleway/nimble-cli/commit/468f6b41868b5255eae9f71d90244c10a8c9cf0c))
- mark all CLI-related tests in Go with `t.Parallel()`
  ([6f6aec0](https://github.com/Nimbleway/nimble-cli/commit/6f6aec0fa8bc971e2544df02ed144458c12e1b80))
- modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary
  ([e1792f8](https://github.com/Nimbleway/nimble-cli/commit/e1792f8342d4572ab972343927444ecc8d485a69))
- switch some CLI Go tests from `os.Chdir` to `t.Chdir`
  ([1937191](https://github.com/Nimbleway/nimble-cli/commit/1937191a2b506e804f707098ae075e8e802f346f))

### Documentation

- improve examples
  ([2809339](https://github.com/Nimbleway/nimble-cli/commit/28093397213052e4cc1c6ec5323fec0bcbb619f9))

## 0.10.0 (2026-03-30)

Full Changelog:
[v0.9.0...v0.10.0](https://github.com/Nimbleway/nimble-cli/compare/v0.9.0...v0.10.0)

### Features

- **api:** api update
  ([3853f10](https://github.com/Nimbleway/nimble-cli/commit/3853f10c43158cf7bbcc1232ac04e7441b19b438))
- **api:** manual updates
  ([3451260](https://github.com/Nimbleway/nimble-cli/commit/34512601014966c7f5cbef8935f58e597e84138f))
- **api:** manual updates
  ([a60bd86](https://github.com/Nimbleway/nimble-cli/commit/a60bd86a2a818b4472388c3d0820638ee47f2ae1))
- **api:** manual updates
  ([c4ef450](https://github.com/Nimbleway/nimble-cli/commit/c4ef45062cc3cd12a05f979650cf03e552991019))
- **api:** rename agent resource -&gt; agents
  ([a2abcba](https://github.com/Nimbleway/nimble-cli/commit/a2abcbaf793632ca040601472d440c79872e6977))

## 0.9.0 (2026-03-29)

Full Changelog:
[v0.8.0...v0.9.0](https://github.com/Nimbleway/nimble-cli/compare/v0.8.0...v0.9.0)

### Features

- **api:** api update
  ([ec48bc8](https://github.com/Nimbleway/nimble-cli/commit/ec48bc8f077f60fa86f08bdae82ff48ccada4c08))
- **api:** api update
  ([38f3d80](https://github.com/Nimbleway/nimble-cli/commit/38f3d80e0d68ec5dfadb7194c15788544e626386))
- **api:** manual updates
  ([20e1bcf](https://github.com/Nimbleway/nimble-cli/commit/20e1bcfa447b3333ce234a702d26efa2b79817bb))
- set CLI flag constant values automatically where `x-stainless-const` is set
  ([7a9161c](https://github.com/Nimbleway/nimble-cli/commit/7a9161ccffc7562dafa0fbb225a20733f1e69ef1))

### Bug Fixes

- fix for off-by-one error in pagination logic
  ([0bd270c](https://github.com/Nimbleway/nimble-cli/commit/0bd270cd50a6a3c9b6e920d737de839c45094e15))

### Chores

- **ci:** skip lint on metadata-only changes
  ([c31626b](https://github.com/Nimbleway/nimble-cli/commit/c31626bb47e30b37ffd433cb0c6960a78a85008f))
- **internal:** codegen related update
  ([6193015](https://github.com/Nimbleway/nimble-cli/commit/6193015a604fb09368ddccebb793582b88016866))
- omit full usage information when missing required CLI parameters
  ([3516aa9](https://github.com/Nimbleway/nimble-cli/commit/3516aa96b9314254cbf70d6ae1283ea15889c0af))

## 0.8.0 (2026-03-24)

Full Changelog:
[v0.7.0...v0.8.0](https://github.com/Nimbleway/nimble-cli/compare/v0.7.0...v0.8.0)

### Features

- add default description for enum CLI flags without an explicit description
  ([c175105](https://github.com/Nimbleway/nimble-cli/commit/c1751051f538856b25ca4df9c26912af89300c23))
- **api:** manual updates
  ([1047e40](https://github.com/Nimbleway/nimble-cli/commit/1047e4046f3ba2f0145a643f2ff1498e06ba0db4))

### Bug Fixes

- cli no longer hangs when stdin is attached to a pipe with empty input
  ([c887be7](https://github.com/Nimbleway/nimble-cli/commit/c887be730c5d184084df3e2a43c1f8541a8fdd5f))
- improve linking behavior when developing on a branch not in the Go SDK
  ([4ba2d72](https://github.com/Nimbleway/nimble-cli/commit/4ba2d72445fa7076186f59495aec56c6723512f9))

### Chores

- **internal:** update gitignore
  ([ff0db9d](https://github.com/Nimbleway/nimble-cli/commit/ff0db9d15a6c9a27effc45ff1420c77351d6f0ee))

## 0.7.0 (2026-03-18)

Full Changelog:
[v0.6.0...v0.7.0](https://github.com/Nimbleway/nimble-cli/compare/v0.6.0...v0.7.0)

### Features

- **api:** manual updates
  ([95cf9c6](https://github.com/Nimbleway/nimble-cli/commit/95cf9c682842315af7be18d4dcba8cd0d8a50fb9))

### Bug Fixes

- avoid reading from stdin unless request body is form encoded or json
  ([aa9fd7f](https://github.com/Nimbleway/nimble-cli/commit/aa9fd7fcadf028e07a6456855ef591998ec72e78))
- better support passing client args in any position
  ([c10c7c9](https://github.com/Nimbleway/nimble-cli/commit/c10c7c923c776a5b056831bcc8d90cbcd8d5f072))
- fix for test cases with newlines in YAML and better error reporting
  ([615aa31](https://github.com/Nimbleway/nimble-cli/commit/615aa316e34d33cb6edfe29b2587adde3f91e8c3))
- improved workflow for developing on branches
  ([e0ff301](https://github.com/Nimbleway/nimble-cli/commit/e0ff301bbf3f55c7436a8f6df076756c7cdd1a04))
- no longer require an API key when building on production repos
  ([5d78761](https://github.com/Nimbleway/nimble-cli/commit/5d78761a8c9375add3a13c707c2e313a8ac14e80))
- only set client options when the corresponding CLI flag or env var is
  explicitly set
  ([8e2d940](https://github.com/Nimbleway/nimble-cli/commit/8e2d94006048dc6fed9884840a5bcb7812fd9a4f))

### Chores

- **internal:** codegen related update
  ([433e269](https://github.com/Nimbleway/nimble-cli/commit/433e269d3b19f915785ec4b773301605c14a95b1))
- **internal:** tweak CI branches
  ([a7ee466](https://github.com/Nimbleway/nimble-cli/commit/a7ee466453f85944128c389e92f737b06db839cb))

## 0.6.0 (2026-03-10)

Full Changelog:
[v0.5.0...v0.6.0](https://github.com/Nimbleway/nimble-cli/compare/v0.5.0...v0.6.0)

### Features

- add `--max-items` flag for paginated/streaming endpoints
  ([f227d2e](https://github.com/Nimbleway/nimble-cli/commit/f227d2e67b41b5afe3f4df4596a58c69975931f3))
- add support for file downloads from binary response endpoints
  ([d52040c](https://github.com/Nimbleway/nimble-cli/commit/d52040cbb309d4d8e9ef5ac517f060677f28d513))
- **api:** api update
  ([2939c8e](https://github.com/Nimbleway/nimble-cli/commit/2939c8e1fac755d5439f5c5cddb2b84411cbe8c6))
- improved documentation and flags for client options
  ([9da30d3](https://github.com/Nimbleway/nimble-cli/commit/9da30d30bbded8a192f8652e0bd67f2df8752499))
- support passing required body params through pipes
  ([a7b02a0](https://github.com/Nimbleway/nimble-cli/commit/a7b02a011d1d9b6827e10a62d2b53e6f3eae55b4))

### Bug Fixes

- avoid printing usage errors twice
  ([4030bc9](https://github.com/Nimbleway/nimble-cli/commit/4030bc9f4cdb7ce23a6ef12381cd640657076ecf))
- fix for encoding arrays with `any` type items
  ([1561b57](https://github.com/Nimbleway/nimble-cli/commit/1561b5794ebffe66bcda053bca842359bdc6fa6c))
- more gracefully handle empty stdin input
  ([b9e421a](https://github.com/Nimbleway/nimble-cli/commit/b9e421aab44a8deae8f120e5fd71e2c6e8ae0f1e))

### Chores

- **ci:** skip uploading artifacts on stainless-internal branches
  ([91e8ea8](https://github.com/Nimbleway/nimble-cli/commit/91e8ea89a329bd6cca606b4da47e0bb3ad06792b))
- **internal:** codegen related update
  ([b5e5738](https://github.com/Nimbleway/nimble-cli/commit/b5e57383eaeb031077d594c785a89a254a44c22f))
- **internal:** codegen related update
  ([15720ce](https://github.com/Nimbleway/nimble-cli/commit/15720cebc33274fd35cbfd825b4103d3dbd004e7))
- **internal:** codegen related update
  ([26fecfd](https://github.com/Nimbleway/nimble-cli/commit/26fecfd207cb7507cdfa8a02936c48005aec3e51))
- zip READMEs as part of build artifact
  ([08aef64](https://github.com/Nimbleway/nimble-cli/commit/08aef645e37003b01b45d4b3a3fc714ce790f8ee))

## 0.5.0 (2026-02-25)

Full Changelog:
[v0.4.3...v0.5.0](https://github.com/Nimbleway/nimble-cli/compare/v0.4.3...v0.5.0)

### Features

- **api:** api update
  ([21b34a8](https://github.com/Nimbleway/nimble-cli/commit/21b34a84903730beb75e47592eef97104464cc8b))

### Bug Fixes

- pin formatting for headers to always use repeat/dot formats
  ([7217885](https://github.com/Nimbleway/nimble-cli/commit/7217885e1862c48dc38988793c079753e8a8791b))

## 0.4.3 (2026-02-22)

Full Changelog:
[v0.4.2...v0.4.3](https://github.com/Nimbleway/nimble-cli/compare/v0.4.2...v0.4.3)

## 0.4.2 (2026-02-22)

Full Changelog:
[v0.4.1...v0.4.2](https://github.com/Nimbleway/nimble-cli/compare/v0.4.1...v0.4.2)

### Bug Fixes

- reset package version to 0.0.0 and update release-please config for versioning
  ([eac05be](https://github.com/Nimbleway/nimble-cli/commit/eac05bede8ae639c96c4a4a72087554114fa141f))

## 0.4.1 (2026-02-22)

Full Changelog:
[v0.4.0...v0.4.1](https://github.com/Nimbleway/nimble-cli/compare/v0.4.0...v0.4.1)

### Chores

- bump npm version
  ([6053477](https://github.com/Nimbleway/nimble-cli/commit/60534777966ca361466c7ba75f7c75958c26db50))
- remove custom code
  ([ad01a20](https://github.com/Nimbleway/nimble-cli/commit/ad01a20f59bf892af667e60392c4970db6fb3b8b))

## 0.4.0 (2026-02-22)

Full Changelog:
[v0.3.0...v0.4.0](https://github.com/Nimbleway/nimble-cli/compare/v0.3.0...v0.4.0)

### Features

- add readme documentation for passing files as arguments
  ([316cc78](https://github.com/Nimbleway/nimble-cli/commit/316cc783d9eda456d9048fa4ad55a645e08f5a62))
- add workflow_dispatch support for npm release
  ([a495f66](https://github.com/Nimbleway/nimble-cli/commit/a495f66d7b03bfcc7fafcc95baf1db34f6f69305))
- **api:** align `extract_async`
  ([25b589a](https://github.com/Nimbleway/nimble-cli/commit/25b589ab5885dc008dea13a05770f6b25fc11046))
- **api:** Align new endpoints
  ([c9c8c64](https://github.com/Nimbleway/nimble-cli/commit/c9c8c647fe5252a488e0ed3fa3cc82b47fb6fd3e))
- **api:** api update
  ([ee584dc](https://github.com/Nimbleway/nimble-cli/commit/ee584dc3c918ea04b7cd84a3e42f2969c17e59d1))
- **api:** api update
  ([edb22ed](https://github.com/Nimbleway/nimble-cli/commit/edb22ed21d75d74b0c9bf6c37e8d22a1f8b66647))
- **api:** manual test
  ([7b32d10](https://github.com/Nimbleway/nimble-cli/commit/7b32d106aa00b314d42ad4bea7e7daff15ae1be9))
- **api:** manual updates
  ([9aa65d0](https://github.com/Nimbleway/nimble-cli/commit/9aa65d08eaf75e8f26d862757cdcd8931d42fdf0))
- **api:** manual updates
  ([33ac975](https://github.com/Nimbleway/nimble-cli/commit/33ac975e4be57ab83f29482de4ccb40059b9eb66))
- **api:** manual updates
  ([8d45cb2](https://github.com/Nimbleway/nimble-cli/commit/8d45cb20afa067da440e698dc99c2900eed0b296))
- **api:** manual updates
  ([1d24021](https://github.com/Nimbleway/nimble-cli/commit/1d2402183e38bb34445ce45c5703c832b0c1ee28))
- **api:** Move /agent to /agents/run
  ([1670467](https://github.com/Nimbleway/nimble-cli/commit/1670467ba4ccf28101471a8e680326012328a322))
- **api:** re-add extract
  ([a87a614](https://github.com/Nimbleway/nimble-cli/commit/a87a614ea89599d008f28d36f0ededf1cb6e6ffe))

### Bug Fixes

- fix for file uploads to octet stream and form encoding endpoints
  ([64deca6](https://github.com/Nimbleway/nimble-cli/commit/64deca67d138628b8d3ed612aca3fcbb09a56ff5))
- fix for resources with the same name as the client
  ([b10dd5a](https://github.com/Nimbleway/nimble-cli/commit/b10dd5a32684593aaa397f66821a9c4981a89ebd))

### Chores

- add build step to ci
  ([57c83d6](https://github.com/Nimbleway/nimble-cli/commit/57c83d6e84bd97e4b3b802db12b0b31b0c7a4fa9))
- **api:** Revert to gateway.webit.live
  ([9b9e640](https://github.com/Nimbleway/nimble-cli/commit/9b9e64085521e2afc45c41872d47a1125f858046))
- configure new SDK language
  ([e3618d6](https://github.com/Nimbleway/nimble-cli/commit/e3618d63143129fe4cbef2f6171b165702ad4089))
- configure new SDK language
  ([183044a](https://github.com/Nimbleway/nimble-cli/commit/183044af0d6429ee1dbdcc8d4b359ce4dbb3730b))
- **internal:** codegen related update
  ([b6fdf17](https://github.com/Nimbleway/nimble-cli/commit/b6fdf175780b024401bc87d6346044992998a3ea))
- **internal:** remove mock server code
  ([9960d8d](https://github.com/Nimbleway/nimble-cli/commit/9960d8d4c82416733b15e815976c70ba11a3963a))
- update documentation in readme
  ([e23a594](https://github.com/Nimbleway/nimble-cli/commit/e23a59419198f8e33753b80e584675ba1fe27354))
- update mock server docs
  ([19b48bc](https://github.com/Nimbleway/nimble-cli/commit/19b48bc7e9a38cedb3e1c47a17e77eb783c5b7f0))
- Update OpenAPI
  ([b2e1bdd](https://github.com/Nimbleway/nimble-cli/commit/b2e1bddfa9bfdd27303deba588a4ef17993df0d7))
- update SDK settings
  ([7a5a3bf](https://github.com/Nimbleway/nimble-cli/commit/7a5a3bf7e68af6c9bb3b29be520f192781a08929))

## 0.2.0 (2026-02-22)

Full Changelog:
[v0.1.0...v0.2.0](https://github.com/Nimbleway/nimble-cli/compare/v0.1.0...v0.2.0)

### Features

- add workflow_dispatch support for npm release
  ([a495f66](https://github.com/Nimbleway/nimble-cli/commit/a495f66d7b03bfcc7fafcc95baf1db34f6f69305))
- **api:** align `extract_async`
  ([25b589a](https://github.com/Nimbleway/nimble-cli/commit/25b589ab5885dc008dea13a05770f6b25fc11046))
- **api:** Align new endpoints
  ([c9c8c64](https://github.com/Nimbleway/nimble-cli/commit/c9c8c647fe5252a488e0ed3fa3cc82b47fb6fd3e))
- **api:** api update
  ([ee584dc](https://github.com/Nimbleway/nimble-cli/commit/ee584dc3c918ea04b7cd84a3e42f2969c17e59d1))
- **api:** manual test
  ([7b32d10](https://github.com/Nimbleway/nimble-cli/commit/7b32d106aa00b314d42ad4bea7e7daff15ae1be9))
- **api:** manual updates
  ([9aa65d0](https://github.com/Nimbleway/nimble-cli/commit/9aa65d08eaf75e8f26d862757cdcd8931d42fdf0))
- **api:** manual updates
  ([33ac975](https://github.com/Nimbleway/nimble-cli/commit/33ac975e4be57ab83f29482de4ccb40059b9eb66))
- **api:** Move /agent to /agents/run
  ([1670467](https://github.com/Nimbleway/nimble-cli/commit/1670467ba4ccf28101471a8e680326012328a322))
- **api:** re-add extract
  ([a87a614](https://github.com/Nimbleway/nimble-cli/commit/a87a614ea89599d008f28d36f0ededf1cb6e6ffe))

### Chores

- **internal:** codegen related update
  ([b6fdf17](https://github.com/Nimbleway/nimble-cli/commit/b6fdf175780b024401bc87d6346044992998a3ea))
- **internal:** remove mock server code
  ([9960d8d](https://github.com/Nimbleway/nimble-cli/commit/9960d8d4c82416733b15e815976c70ba11a3963a))
- update mock server docs
  ([19b48bc](https://github.com/Nimbleway/nimble-cli/commit/19b48bc7e9a38cedb3e1c47a17e77eb783c5b7f0))

## 0.1.0 (2026-02-15)

Full Changelog:
[v0.0.1...v0.1.0](https://github.com/Nimbleway/nimble-cli/compare/v0.0.1...v0.1.0)

### Features

- add readme documentation for passing files as arguments
  ([316cc78](https://github.com/Nimbleway/nimble-cli/commit/316cc783d9eda456d9048fa4ad55a645e08f5a62))
- **api:** api update
  ([edb22ed](https://github.com/Nimbleway/nimble-cli/commit/edb22ed21d75d74b0c9bf6c37e8d22a1f8b66647))
- **api:** manual updates
  ([093bcf3](https://github.com/Nimbleway/nimble-cli/commit/093bcf372a42904eeabffbabb3c56dd94b5dcfeb))
- **api:** manual updates
  ([1d24021](https://github.com/Nimbleway/nimble-cli/commit/1d2402183e38bb34445ce45c5703c832b0c1ee28))

### Bug Fixes

- fix for file uploads to octet stream and form encoding endpoints
  ([64deca6](https://github.com/Nimbleway/nimble-cli/commit/64deca67d138628b8d3ed612aca3fcbb09a56ff5))
- fix for resources with the same name as the client
  ([b10dd5a](https://github.com/Nimbleway/nimble-cli/commit/b10dd5a32684593aaa397f66821a9c4981a89ebd))

### Chores

- add build step to ci
  ([57c83d6](https://github.com/Nimbleway/nimble-cli/commit/57c83d6e84bd97e4b3b802db12b0b31b0c7a4fa9))
- **api:** Revert to gateway.webit.live
  ([926a26a](https://github.com/Nimbleway/nimble-cli/commit/926a26aa1cc10a981dd10ecd76f30b06efa619a8))
- configure new SDK language
  ([d099d33](https://github.com/Nimbleway/nimble-cli/commit/d099d33824b7a42dd87757181d9322c8762cbdf7))
- configure new SDK language
  ([183044a](https://github.com/Nimbleway/nimble-cli/commit/183044af0d6429ee1dbdcc8d4b359ce4dbb3730b))
- update documentation in readme
  ([e23a594](https://github.com/Nimbleway/nimble-cli/commit/e23a59419198f8e33753b80e584675ba1fe27354))
- Update OpenAPI
  ([b2e1bdd](https://github.com/Nimbleway/nimble-cli/commit/b2e1bddfa9bfdd27303deba588a4ef17993df0d7))
- update SDK settings
  ([7a5a3bf](https://github.com/Nimbleway/nimble-cli/commit/7a5a3bf7e68af6c9bb3b29be520f192781a08929))
