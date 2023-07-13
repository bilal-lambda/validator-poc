package main

import (
	"fmt"
	"syscall/js"

	"gopkg.in/yaml.v3"
)

type SampleYamlStruct struct {
	Version string `yaml:"version"`
	Runson string `yaml:"runson"`
	CodeDirectory string `yaml:"codeDirectory"`
	TestSuites []string `yaml:"testSuites"`
}

func validateYaml(_this js.Value, args []js.Value) interface{} {
	ok := js.ValueOf("ok")

	if len(args) == 0 {
		return ok
	}

	source := args[0].String()

	var yml SampleYamlStruct
	err := yaml.Unmarshal([]byte(source), &yml)
	if err != nil {
		return js.ValueOf(fmt.Sprintf("invalid yaml - %s", err))
	}

	if yml.Runson != "linux" && yml.Runson != "macos" && yml.Runson != "win" {
		return js.ValueOf("runson can only be linux, macos or win")
	}

  if yml.Version == "" {
    		return js.ValueOf("missing version")
  }

	if yml.Version != "0.1" && yml.Version != "0.2" {
		return js.ValueOf("incorrect version")
	}

	return ok
}

func main() {
	js.Global().Set("validateYaml", js.FuncOf(validateYaml))
	select {}
}
