# Changelog

## 0.7.0 (2026-03-18)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/Nimbleway/nimble-cli/compare/v0.6.0...v0.7.0)

### Features

* **api:** manual updates ([95cf9c6](https://github.com/Nimbleway/nimble-cli/commit/95cf9c682842315af7be18d4dcba8cd0d8a50fb9))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([aa9fd7f](https://github.com/Nimbleway/nimble-cli/commit/aa9fd7fcadf028e07a6456855ef591998ec72e78))
* better support passing client args in any position ([c10c7c9](https://github.com/Nimbleway/nimble-cli/commit/c10c7c923c776a5b056831bcc8d90cbcd8d5f072))
* fix for test cases with newlines in YAML and better error reporting ([615aa31](https://github.com/Nimbleway/nimble-cli/commit/615aa316e34d33cb6edfe29b2587adde3f91e8c3))
* improved workflow for developing on branches ([e0ff301](https://github.com/Nimbleway/nimble-cli/commit/e0ff301bbf3f55c7436a8f6df076756c7cdd1a04))
* no longer require an API key when building on production repos ([5d78761](https://github.com/Nimbleway/nimble-cli/commit/5d78761a8c9375add3a13c707c2e313a8ac14e80))
* only set client options when the corresponding CLI flag or env var is explicitly set ([8e2d940](https://github.com/Nimbleway/nimble-cli/commit/8e2d94006048dc6fed9884840a5bcb7812fd9a4f))


### Chores

* **internal:** codegen related update ([433e269](https://github.com/Nimbleway/nimble-cli/commit/433e269d3b19f915785ec4b773301605c14a95b1))
* **internal:** tweak CI branches ([a7ee466](https://github.com/Nimbleway/nimble-cli/commit/a7ee466453f85944128c389e92f737b06db839cb))

## 0.6.0 (2026-03-10)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/Nimbleway/nimble-cli/compare/v0.5.0...v0.6.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([f227d2e](https://github.com/Nimbleway/nimble-cli/commit/f227d2e67b41b5afe3f4df4596a58c69975931f3))
* add support for file downloads from binary response endpoints ([d52040c](https://github.com/Nimbleway/nimble-cli/commit/d52040cbb309d4d8e9ef5ac517f060677f28d513))
* **api:** api update ([2939c8e](https://github.com/Nimbleway/nimble-cli/commit/2939c8e1fac755d5439f5c5cddb2b84411cbe8c6))
* improved documentation and flags for client options ([9da30d3](https://github.com/Nimbleway/nimble-cli/commit/9da30d30bbded8a192f8652e0bd67f2df8752499))
* support passing required body params through pipes ([a7b02a0](https://github.com/Nimbleway/nimble-cli/commit/a7b02a011d1d9b6827e10a62d2b53e6f3eae55b4))


### Bug Fixes

* avoid printing usage errors twice ([4030bc9](https://github.com/Nimbleway/nimble-cli/commit/4030bc9f4cdb7ce23a6ef12381cd640657076ecf))
* fix for encoding arrays with `any` type items ([1561b57](https://github.com/Nimbleway/nimble-cli/commit/1561b5794ebffe66bcda053bca842359bdc6fa6c))
* more gracefully handle empty stdin input ([b9e421a](https://github.com/Nimbleway/nimble-cli/commit/b9e421aab44a8deae8f120e5fd71e2c6e8ae0f1e))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([91e8ea8](https://github.com/Nimbleway/nimble-cli/commit/91e8ea89a329bd6cca606b4da47e0bb3ad06792b))
* **internal:** codegen related update ([b5e5738](https://github.com/Nimbleway/nimble-cli/commit/b5e57383eaeb031077d594c785a89a254a44c22f))
* **internal:** codegen related update ([15720ce](https://github.com/Nimbleway/nimble-cli/commit/15720cebc33274fd35cbfd825b4103d3dbd004e7))
* **internal:** codegen related update ([26fecfd](https://github.com/Nimbleway/nimble-cli/commit/26fecfd207cb7507cdfa8a02936c48005aec3e51))
* zip READMEs as part of build artifact ([08aef64](https://github.com/Nimbleway/nimble-cli/commit/08aef645e37003b01b45d4b3a3fc714ce790f8ee))

## 0.5.0 (2026-02-25)

Full Changelog: [v0.4.3...v0.5.0](https://github.com/Nimbleway/nimble-cli/compare/v0.4.3...v0.5.0)

### Features

* **api:** api update ([21b34a8](https://github.com/Nimbleway/nimble-cli/commit/21b34a84903730beb75e47592eef97104464cc8b))


### Bug Fixes

* pin formatting for headers to always use repeat/dot formats ([7217885](https://github.com/Nimbleway/nimble-cli/commit/7217885e1862c48dc38988793c079753e8a8791b))

## 0.4.3 (2026-02-22)

Full Changelog: [v0.4.2...v0.4.3](https://github.com/Nimbleway/nimble-cli/compare/v0.4.2...v0.4.3)

## 0.4.2 (2026-02-22)

Full Changelog: [v0.4.1...v0.4.2](https://github.com/Nimbleway/nimble-cli/compare/v0.4.1...v0.4.2)

### Bug Fixes

* reset package version to 0.0.0 and update release-please config for versioning ([eac05be](https://github.com/Nimbleway/nimble-cli/commit/eac05bede8ae639c96c4a4a72087554114fa141f))

## 0.4.1 (2026-02-22)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/Nimbleway/nimble-cli/compare/v0.4.0...v0.4.1)

### Chores

* bump npm version ([6053477](https://github.com/Nimbleway/nimble-cli/commit/60534777966ca361466c7ba75f7c75958c26db50))
* remove custom code ([ad01a20](https://github.com/Nimbleway/nimble-cli/commit/ad01a20f59bf892af667e60392c4970db6fb3b8b))

## 0.4.0 (2026-02-22)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/Nimbleway/nimble-cli/compare/v0.3.0...v0.4.0)

### Features

* add readme documentation for passing files as arguments ([316cc78](https://github.com/Nimbleway/nimble-cli/commit/316cc783d9eda456d9048fa4ad55a645e08f5a62))
* add workflow_dispatch support for npm release ([a495f66](https://github.com/Nimbleway/nimble-cli/commit/a495f66d7b03bfcc7fafcc95baf1db34f6f69305))
* **api:** align `extract_async` ([25b589a](https://github.com/Nimbleway/nimble-cli/commit/25b589ab5885dc008dea13a05770f6b25fc11046))
* **api:** Align new endpoints ([c9c8c64](https://github.com/Nimbleway/nimble-cli/commit/c9c8c647fe5252a488e0ed3fa3cc82b47fb6fd3e))
* **api:** api update ([ee584dc](https://github.com/Nimbleway/nimble-cli/commit/ee584dc3c918ea04b7cd84a3e42f2969c17e59d1))
* **api:** api update ([edb22ed](https://github.com/Nimbleway/nimble-cli/commit/edb22ed21d75d74b0c9bf6c37e8d22a1f8b66647))
* **api:** manual test ([7b32d10](https://github.com/Nimbleway/nimble-cli/commit/7b32d106aa00b314d42ad4bea7e7daff15ae1be9))
* **api:** manual updates ([9aa65d0](https://github.com/Nimbleway/nimble-cli/commit/9aa65d08eaf75e8f26d862757cdcd8931d42fdf0))
* **api:** manual updates ([33ac975](https://github.com/Nimbleway/nimble-cli/commit/33ac975e4be57ab83f29482de4ccb40059b9eb66))
* **api:** manual updates ([8d45cb2](https://github.com/Nimbleway/nimble-cli/commit/8d45cb20afa067da440e698dc99c2900eed0b296))
* **api:** manual updates ([1d24021](https://github.com/Nimbleway/nimble-cli/commit/1d2402183e38bb34445ce45c5703c832b0c1ee28))
* **api:** Move /agent to /agents/run ([1670467](https://github.com/Nimbleway/nimble-cli/commit/1670467ba4ccf28101471a8e680326012328a322))
* **api:** re-add extract ([a87a614](https://github.com/Nimbleway/nimble-cli/commit/a87a614ea89599d008f28d36f0ededf1cb6e6ffe))


### Bug Fixes

* fix for file uploads to octet stream and form encoding endpoints ([64deca6](https://github.com/Nimbleway/nimble-cli/commit/64deca67d138628b8d3ed612aca3fcbb09a56ff5))
* fix for resources with the same name as the client ([b10dd5a](https://github.com/Nimbleway/nimble-cli/commit/b10dd5a32684593aaa397f66821a9c4981a89ebd))


### Chores

* add build step to ci ([57c83d6](https://github.com/Nimbleway/nimble-cli/commit/57c83d6e84bd97e4b3b802db12b0b31b0c7a4fa9))
* **api:** Revert to gateway.webit.live ([9b9e640](https://github.com/Nimbleway/nimble-cli/commit/9b9e64085521e2afc45c41872d47a1125f858046))
* configure new SDK language ([e3618d6](https://github.com/Nimbleway/nimble-cli/commit/e3618d63143129fe4cbef2f6171b165702ad4089))
* configure new SDK language ([183044a](https://github.com/Nimbleway/nimble-cli/commit/183044af0d6429ee1dbdcc8d4b359ce4dbb3730b))
* **internal:** codegen related update ([b6fdf17](https://github.com/Nimbleway/nimble-cli/commit/b6fdf175780b024401bc87d6346044992998a3ea))
* **internal:** remove mock server code ([9960d8d](https://github.com/Nimbleway/nimble-cli/commit/9960d8d4c82416733b15e815976c70ba11a3963a))
* update documentation in readme ([e23a594](https://github.com/Nimbleway/nimble-cli/commit/e23a59419198f8e33753b80e584675ba1fe27354))
* update mock server docs ([19b48bc](https://github.com/Nimbleway/nimble-cli/commit/19b48bc7e9a38cedb3e1c47a17e77eb783c5b7f0))
* Update OpenAPI ([b2e1bdd](https://github.com/Nimbleway/nimble-cli/commit/b2e1bddfa9bfdd27303deba588a4ef17993df0d7))
* update SDK settings ([7a5a3bf](https://github.com/Nimbleway/nimble-cli/commit/7a5a3bf7e68af6c9bb3b29be520f192781a08929))

## 0.2.0 (2026-02-22)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Nimbleway/nimble-cli/compare/v0.1.0...v0.2.0)

### Features

* add workflow_dispatch support for npm release ([a495f66](https://github.com/Nimbleway/nimble-cli/commit/a495f66d7b03bfcc7fafcc95baf1db34f6f69305))
* **api:** align `extract_async` ([25b589a](https://github.com/Nimbleway/nimble-cli/commit/25b589ab5885dc008dea13a05770f6b25fc11046))
* **api:** Align new endpoints ([c9c8c64](https://github.com/Nimbleway/nimble-cli/commit/c9c8c647fe5252a488e0ed3fa3cc82b47fb6fd3e))
* **api:** api update ([ee584dc](https://github.com/Nimbleway/nimble-cli/commit/ee584dc3c918ea04b7cd84a3e42f2969c17e59d1))
* **api:** manual test ([7b32d10](https://github.com/Nimbleway/nimble-cli/commit/7b32d106aa00b314d42ad4bea7e7daff15ae1be9))
* **api:** manual updates ([9aa65d0](https://github.com/Nimbleway/nimble-cli/commit/9aa65d08eaf75e8f26d862757cdcd8931d42fdf0))
* **api:** manual updates ([33ac975](https://github.com/Nimbleway/nimble-cli/commit/33ac975e4be57ab83f29482de4ccb40059b9eb66))
* **api:** Move /agent to /agents/run ([1670467](https://github.com/Nimbleway/nimble-cli/commit/1670467ba4ccf28101471a8e680326012328a322))
* **api:** re-add extract ([a87a614](https://github.com/Nimbleway/nimble-cli/commit/a87a614ea89599d008f28d36f0ededf1cb6e6ffe))


### Chores

* **internal:** codegen related update ([b6fdf17](https://github.com/Nimbleway/nimble-cli/commit/b6fdf175780b024401bc87d6346044992998a3ea))
* **internal:** remove mock server code ([9960d8d](https://github.com/Nimbleway/nimble-cli/commit/9960d8d4c82416733b15e815976c70ba11a3963a))
* update mock server docs ([19b48bc](https://github.com/Nimbleway/nimble-cli/commit/19b48bc7e9a38cedb3e1c47a17e77eb783c5b7f0))

## 0.1.0 (2026-02-15)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/Nimbleway/nimble-cli/compare/v0.0.1...v0.1.0)

### Features

* add readme documentation for passing files as arguments ([316cc78](https://github.com/Nimbleway/nimble-cli/commit/316cc783d9eda456d9048fa4ad55a645e08f5a62))
* **api:** api update ([edb22ed](https://github.com/Nimbleway/nimble-cli/commit/edb22ed21d75d74b0c9bf6c37e8d22a1f8b66647))
* **api:** manual updates ([093bcf3](https://github.com/Nimbleway/nimble-cli/commit/093bcf372a42904eeabffbabb3c56dd94b5dcfeb))
* **api:** manual updates ([1d24021](https://github.com/Nimbleway/nimble-cli/commit/1d2402183e38bb34445ce45c5703c832b0c1ee28))


### Bug Fixes

* fix for file uploads to octet stream and form encoding endpoints ([64deca6](https://github.com/Nimbleway/nimble-cli/commit/64deca67d138628b8d3ed612aca3fcbb09a56ff5))
* fix for resources with the same name as the client ([b10dd5a](https://github.com/Nimbleway/nimble-cli/commit/b10dd5a32684593aaa397f66821a9c4981a89ebd))


### Chores

* add build step to ci ([57c83d6](https://github.com/Nimbleway/nimble-cli/commit/57c83d6e84bd97e4b3b802db12b0b31b0c7a4fa9))
* **api:** Revert to gateway.webit.live ([926a26a](https://github.com/Nimbleway/nimble-cli/commit/926a26aa1cc10a981dd10ecd76f30b06efa619a8))
* configure new SDK language ([d099d33](https://github.com/Nimbleway/nimble-cli/commit/d099d33824b7a42dd87757181d9322c8762cbdf7))
* configure new SDK language ([183044a](https://github.com/Nimbleway/nimble-cli/commit/183044af0d6429ee1dbdcc8d4b359ce4dbb3730b))
* update documentation in readme ([e23a594](https://github.com/Nimbleway/nimble-cli/commit/e23a59419198f8e33753b80e584675ba1fe27354))
* Update OpenAPI ([b2e1bdd](https://github.com/Nimbleway/nimble-cli/commit/b2e1bddfa9bfdd27303deba588a4ef17993df0d7))
* update SDK settings ([7a5a3bf](https://github.com/Nimbleway/nimble-cli/commit/7a5a3bf7e68af6c9bb3b29be520f192781a08929))
