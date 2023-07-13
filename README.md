# YAML validator

This is a proof-of-concept YAML validation tool which can be used in the CLI as golang code and can be used on FE as wasm code.

# Motivation

During onboarding process of Hyperexecute, users may make mistakes in the yaml. The job will fail after the code is uploaded, VMs are assigned and tests are ready to execute on the VM. Where applicable, this feedback can be shortened by providing a validation before code is uploaded.

Recent examples:
- language `version` for the `runtime` provided with quotes which fails on the VM
- `token` used in the `sourcePayload` section instead of `accessToken`

We already have certain rules baked in the CLI which can be updated regularly. It might be helpful to have a user interface to generate and validate the yaml via FE. A lot of these rules which are not dependent on the file system can be reused. Moreover, it seems possible to refactor the CLI so that the validations are extracted out and can be targetted for FE too. Having a same source for CLI validation and frontend reduces the chances of out-of-sync validation where CLI and FE might differ on the correctness of provided yaml.

A sample yaml which this tool validates

```
version: "0.1"
runson: win
codeDirectory: code
testSuites:
  - npm run test
```
