# Changelog

## 0.8.1 (2026-06-01)

Full Changelog: [v0.8.0...v0.8.1](https://github.com/stiggio/stigg-cli/compare/v0.8.0...v0.8.1)

### Chores

* remove custom code ([e1bf90c](https://github.com/stiggio/stigg-cli/commit/e1bf90c39897dc670d495aa2a1c98379868ee1d1))

## 0.8.0 (2026-05-26)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/stiggio/stigg-cli/compare/v0.7.1...v0.8.0)

### Features

* add default description for enum CLI flags without an explicit description ([e073679](https://github.com/stiggio/stigg-cli/commit/e0736799f52f2bda2707065bb18de993c0378ea2))
* allow `-` as value representing stdin to binary-only file parameters in CLIs ([2e80a9d](https://github.com/stiggio/stigg-cli/commit/2e80a9d67551395be133547f636033bfa8cc864c))
* **api:** api update ([3edca52](https://github.com/stiggio/stigg-cli/commit/3edca52125babd1041082aca4dcd5fafdddffb12))
* **api:** api update ([8c4a5d6](https://github.com/stiggio/stigg-cli/commit/8c4a5d68fd4821eb9f707d3cf13c492e746809b2))
* **api:** api update ([2588549](https://github.com/stiggio/stigg-cli/commit/2588549292cbe883d67a630c883a755aafa023a5))
* **api:** api update ([c608f59](https://github.com/stiggio/stigg-cli/commit/c608f599febf2a401ea7feb52dd793200d828bdc))
* **api:** api update ([307016d](https://github.com/stiggio/stigg-cli/commit/307016d6d5f1b8addbd0177b839e8c1fd2daa97e))
* **api:** api update ([0e59b7c](https://github.com/stiggio/stigg-cli/commit/0e59b7cef4d02f8f94477fdf57273437de026011))
* **api:** api update ([4540fac](https://github.com/stiggio/stigg-cli/commit/4540facb99f4d976731744eeef4f9303bd2c90b6))
* **api:** updated stainless config with new endpoint ([860152d](https://github.com/stiggio/stigg-cli/commit/860152d4c9b576466d1ebc4c1aaabec350a37933))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([b9b06a4](https://github.com/stiggio/stigg-cli/commit/b9b06a40dac741cbb9955c3e6baac6360970d23f))
* binary-only parameters become CLI flags that take filenames only ([883d0f8](https://github.com/stiggio/stigg-cli/commit/883d0f8c73dcb868af0e68c71adbd4459219e337))
* set CLI flag constant values automatically where `x-stainless-const` is set ([bdcdbae](https://github.com/stiggio/stigg-cli/commit/bdcdbae10b4743038f82c66a5285593ee2f9e2a9))
* **STIGG-6815:** auto-generate openapi schema stainless ([a41072d](https://github.com/stiggio/stigg-cli/commit/a41072da21d8c7467139a155110d40957535c9ea))
* **STIGG-7501:** split cluster-fargate-vpc and doggo-persistent-cach… ([6596885](https://github.com/stiggio/stigg-cli/commit/65968857271608efddf410a228ac8fee9d204f97))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([77c71ae](https://github.com/stiggio/stigg-cli/commit/77c71ae1e41f42226499638bda15c14a2f8c5476))
* cli no longer hangs when stdin is attached to a pipe with empty input ([c2da831](https://github.com/stiggio/stigg-cli/commit/c2da8318e7d73685cbf7fcdd17f56f6b2971e1c6))
* **cli:** fix incompatible Go types for flag generated as array of maps ([0c04026](https://github.com/stiggio/stigg-cli/commit/0c040269594b4304550c75e72d08f1f69ecd7ab2))
* fall back to main branch if linking fails in CI ([80319f4](https://github.com/stiggio/stigg-cli/commit/80319f465823c832a99493a74a49437a0fd19d7f))
* fix for failing to drop invalid module replace in link script ([1c4fa49](https://github.com/stiggio/stigg-cli/commit/1c4fa4953d1f4c3a59f0dae67251b419bca015ad))
* fix for off-by-one error in pagination logic ([ce9aded](https://github.com/stiggio/stigg-cli/commit/ce9aded4af3caf3d3c5435803ad965acffb70929))
* fix quoting typo ([54a361f](https://github.com/stiggio/stigg-cli/commit/54a361fdb635f938f2cdb456feec6fa8318b547b))
* handle empty data set using `--format explore` ([2737d6b](https://github.com/stiggio/stigg-cli/commit/2737d6bb5eaf4b359277ea9fbe9c0367b10ed2be))
* improve linking behavior when developing on a branch not in the Go SDK ([e3a4570](https://github.com/stiggio/stigg-cli/commit/e3a457095cbf85d3cd021041875fe9aff8ae248c))
* **STIGG-7500:** stripe account not deauthorized when connected to multiple environments ([ccf98a6](https://github.com/stiggio/stigg-cli/commit/ccf98a6b0da218407620f2d55b247f8a46d00522))
* **STIGG-7502:** sync paid credit grants to DynamoDB cache after acti… ([85640d4](https://github.com/stiggio/stigg-cli/commit/85640d48935a64ccc136e29c893e74ee11464138))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([8984eaa](https://github.com/stiggio/stigg-cli/commit/8984eaace04ea0a2965a94d17627370578a1c004))


### Reverts

* **STIGG-7500:** stripe account not deauthorized when connected to multiple environments ([cdb8d29](https://github.com/stiggio/stigg-cli/commit/cdb8d29056f38ef531ef187d8eaed6eb069f0bca))


### Chores

* **ci:** skip lint on metadata-only changes ([bf5cb08](https://github.com/stiggio/stigg-cli/commit/bf5cb08995b03c7831af19d9cbd65781e9eb995f))
* **cli:** additional test cases for `ShowJSONIterator` ([e3099ee](https://github.com/stiggio/stigg-cli/commit/e3099ee4b0ce6e15f167d86dbc8ee3cdf70a287b))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([9a7e062](https://github.com/stiggio/stigg-cli/commit/9a7e062fd97dac93c22b9304adac201dc781ee95))
* **internal:** codegen related update ([99c5b6c](https://github.com/stiggio/stigg-cli/commit/99c5b6cfac2974ea2f7cf6938ed8f3ab0adc4be8))
* **internal:** codegen related update ([504ba0e](https://github.com/stiggio/stigg-cli/commit/504ba0e75676907ba2b854e233348482cb75ca4c))
* **internal:** codegen related update ([2f6bf32](https://github.com/stiggio/stigg-cli/commit/2f6bf321ce4eb98082e404d6184f5c5fe83038ca))
* **internal:** codegen related update ([12634e0](https://github.com/stiggio/stigg-cli/commit/12634e04985e2c480f67c197942e37e8b3a261b3))
* **internal:** update gitignore ([25a968a](https://github.com/stiggio/stigg-cli/commit/25a968ac204c23ec2f75f45eca3028950ac8253a))
* mark all CLI-related tests in Go with `t.Parallel()` ([751fb4f](https://github.com/stiggio/stigg-cli/commit/751fb4f8143fa21551d10c64665003da31eac159))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([21aa260](https://github.com/stiggio/stigg-cli/commit/21aa260da5925f765c19ee93ed8cdca4631fdef0))
* omit full usage information when missing required CLI parameters ([ac35bd1](https://github.com/stiggio/stigg-cli/commit/ac35bd1e2b34dadf1b80e870b571a9f6398351a1))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([beccd9a](https://github.com/stiggio/stigg-cli/commit/beccd9a38d07fe2b4c253576298dabbb39958edc))

## 0.7.1 (2026-03-17)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/stiggio/stigg-cli/compare/v0.7.0...v0.7.1)

### Bug Fixes

* better support passing client args in any position ([bf63688](https://github.com/stiggio/stigg-cli/commit/bf6368819c4f50c72d7f5fbe48533efdff3a7bc7))
* improved workflow for developing on branches ([fe4d128](https://github.com/stiggio/stigg-cli/commit/fe4d1286597dc5e967ac1b125dcb76bace5f8ce8))
* no longer require an API key when building on production repos ([2622e41](https://github.com/stiggio/stigg-cli/commit/2622e41a233b98e7fd9401e4b06768f8af9b769f))


### Chores

* **internal:** tweak CI branches ([8f8ba0d](https://github.com/stiggio/stigg-cli/commit/8f8ba0d1b19d01ec1ace9b7bef1b8329b37edcf4))

## 0.7.0 (2026-03-16)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/stiggio/stigg-cli/compare/v0.6.0...v0.7.0)

### Features

* **api:** api update ([7aeca96](https://github.com/stiggio/stigg-cli/commit/7aeca96df203df0deda9ee87e9ddd5e1996b16bc))

## 0.6.0 (2026-03-16)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/stiggio/stigg-cli/compare/v0.5.0...v0.6.0)

### Features

* **api:** api update ([bf47346](https://github.com/stiggio/stigg-cli/commit/bf47346e810ec845228ca557337d67ae88aa740f))

## 0.5.0 (2026-03-16)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/stiggio/stigg-cli/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([038e00b](https://github.com/stiggio/stigg-cli/commit/038e00b19c4405f14a46ccdbf94220aa5b1b870d))
* **api:** api update ([879938d](https://github.com/stiggio/stigg-cli/commit/879938db058215a83dc2ec3bcab85ff9ec462386))

## 0.4.0 (2026-03-15)

Full Changelog: [v0.3.1...v0.4.0](https://github.com/stiggio/stigg-cli/compare/v0.3.1...v0.4.0)

### Features

* **api:** api update ([dad9e83](https://github.com/stiggio/stigg-cli/commit/dad9e8391e02d062d179a2f45141b53e0347b826))

## 0.3.1 (2026-03-14)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/stiggio/stigg-cli/compare/v0.3.0...v0.3.1)

### Bug Fixes

* only set client options when the corresponding CLI flag or env var is explicitly set ([aff5757](https://github.com/stiggio/stigg-cli/commit/aff5757bbd00997b170ffc068aa69ad38c14d938))

## 0.3.0 (2026-03-12)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/stiggio/stigg-cli/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([a8de03d](https://github.com/stiggio/stigg-cli/commit/a8de03d1ed6f0e2c73c90ffd8cdfab30c7f945bb))

## 0.2.0 (2026-03-12)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/stiggio/stigg-cli/compare/v0.1.1...v0.2.0)

### Features

* **api:** api update ([b3e9558](https://github.com/stiggio/stigg-cli/commit/b3e95584b8632d7198f5e48f326a053352902f3d))

## 0.1.1 (2026-03-12)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/stiggio/stigg-cli/compare/v0.1.0...v0.1.1)

### Bug Fixes

* fix for test cases with newlines in YAML and better error reporting ([44c8bc2](https://github.com/stiggio/stigg-cli/commit/44c8bc2b18181c9295920a3f9f45c9f27cbe8723))

## 0.1.0 (2026-03-10)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/stiggio/stigg-cli/compare/v0.0.1...v0.1.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([d1b2890](https://github.com/stiggio/stigg-cli/commit/d1b28908466b7ea8cfa9246ef1f45c356ef693be))
* add support for file downloads from binary response endpoints ([31bdd04](https://github.com/stiggio/stigg-cli/commit/31bdd04dbf0e39c97aec6aea6aa9ccb62c51fbd1))
* **api:** added credits resources ([14c9d47](https://github.com/stiggio/stigg-cli/commit/14c9d475f33faca3f977542a41574e78882d53a4))
* **api:** added homebrew publishing of stigg cli ([84fefa4](https://github.com/stiggio/stigg-cli/commit/84fefa46018b42459f6597bfefb81b058b2501ca))
* **api:** api update ([d1049e6](https://github.com/stiggio/stigg-cli/commit/d1049e6742951601d8037e02989ddc294e8e91f3))
* **api:** api update ([0784a04](https://github.com/stiggio/stigg-cli/commit/0784a044d57a879ae99b98953fed1569a19d97c8))
* **api:** api update ([df5191b](https://github.com/stiggio/stigg-cli/commit/df5191b27924240935ae51a4e4e367343b700f49))
* **api:** api update ([ad0ebbc](https://github.com/stiggio/stigg-cli/commit/ad0ebbcc7fe3eb716194de670722edb745da9be5))
* **api:** api update ([899fe7a](https://github.com/stiggio/stigg-cli/commit/899fe7a747e567e235b7424fc4312324d7b7ffa0))
* **api:** api update ([dfcedb9](https://github.com/stiggio/stigg-cli/commit/dfcedb9df28a1366ea26db6dbedab0044d0dd988))
* **api:** api update ([3b4d8d6](https://github.com/stiggio/stigg-cli/commit/3b4d8d67f65226c40c9562ecc87aa6b9ef63591e))
* **api:** api update ([0843656](https://github.com/stiggio/stigg-cli/commit/08436565a849f9801705f07e7eefd95dedd39e0f))
* **api:** api update ([47458a0](https://github.com/stiggio/stigg-cli/commit/47458a0f98187fe32a7f1f309ec534575d6a48ab))
* **api:** api update ([62a9734](https://github.com/stiggio/stigg-cli/commit/62a97342e541550f92174376d82f13a36ea4b994))
* **api:** api update ([b08059d](https://github.com/stiggio/stigg-cli/commit/b08059daadfd6d7120ea2ec3a56e67a70e37c847))
* **api:** api update ([44388e5](https://github.com/stiggio/stigg-cli/commit/44388e5f477b377cd35f7138a9834d7e0c41c781))
* **api:** api update ([5121197](https://github.com/stiggio/stigg-cli/commit/5121197c4b9f010ca46cbd2e8cb9b4d7d8ada4f0))
* **api:** api update ([72d25a4](https://github.com/stiggio/stigg-cli/commit/72d25a42e70b74fa49f2c511d60b1bae62d33eab))
* **api:** api update ([8b73810](https://github.com/stiggio/stigg-cli/commit/8b73810e20d38dcda76742a51da21f1f557fe24b))
* **api:** api update ([57a3791](https://github.com/stiggio/stigg-cli/commit/57a37919139fed6e2b4e7d2f7bb414f66bc74611))
* **api:** api update ([7f4fe05](https://github.com/stiggio/stigg-cli/commit/7f4fe05f66d7508fb5742538936c437041b02d85))
* **api:** stainless config updates ([d2d6213](https://github.com/stiggio/stigg-cli/commit/d2d62135ccfb7884fb5c2ea597c096411a6865fb))
* **api:** update endpoints and models ([f642ba1](https://github.com/stiggio/stigg-cli/commit/f642ba1af8df5a59b610a2510c5186c6217ebda0))
* improved documentation and flags for client options ([60c9c95](https://github.com/stiggio/stigg-cli/commit/60c9c95d849ac15521a1ae448417c54427d5325c))
* support passing required body params through pipes ([8ee3f0c](https://github.com/stiggio/stigg-cli/commit/8ee3f0cbd5c1f660680eed75e3519fcf97a4c327))


### Bug Fixes

* avoid printing usage errors twice ([3ac2fd9](https://github.com/stiggio/stigg-cli/commit/3ac2fd97d6c7d39bd5b3bd04be2eaa30dad7b52c))
* fix for encoding arrays with `any` type items ([b4425f8](https://github.com/stiggio/stigg-cli/commit/b4425f85e8adcff2ab803132147a21358ca45c33))
* more gracefully handle empty stdin input ([0c3cbd3](https://github.com/stiggio/stigg-cli/commit/0c3cbd33f38091d89e7cf38a764f059e41d7c31a))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([794af05](https://github.com/stiggio/stigg-cli/commit/794af059a441bae63037bb0ea6c1f89445af3e3a))
* configure new SDK language ([42c5c5a](https://github.com/stiggio/stigg-cli/commit/42c5c5ab2e3234fe1e9675ff9069383fbefaacaa))
* configure new SDK language ([ebe422f](https://github.com/stiggio/stigg-cli/commit/ebe422ffeb8a3bd2489c15722349790080f92865))
* **internal:** codegen related update ([0299346](https://github.com/stiggio/stigg-cli/commit/02993468cc0734e324bee52ccec35b5ae0945c0c))
* **internal:** codegen related update ([6beca1e](https://github.com/stiggio/stigg-cli/commit/6beca1e3111669ae8648244047555a7f960ade06))
* update SDK settings ([5866d69](https://github.com/stiggio/stigg-cli/commit/5866d697537a8cf257cb143050c2c198cd8a95fb))
