# validator-poc
Proof-of-Concept for yaml validation. Can be used for validating hyperexecute yaml.

# Description

This is a YAML validation tool which can be used in the CLI as golang code and can be used on FE as wasm code.

# Motivation

During onboarding process of Hyperexecute, users may make mistakes in the yaml which fails after the code is uploaded, VMs are assigned and tests are ready to execute on the VM. Where applicable, this feedback can be shortened by providing a validation before code is uploaded. Having a same source for CLI validation and frontend reduces the chances of out-of-sync validation where CLI and FE might differ on the correctness of provided yaml.

A sample yaml which this tool validates

```
version: "0.1"
runson: win
codeDirectory: code
testSuites:
  - npm run test
```
