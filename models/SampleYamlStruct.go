package models

import (
	_ "gopkg.in/yaml.v3"
)

type SampleYamlStruct struct {
	Version string `yaml:"version"`
	Runson string `yaml:"runson"`
	CodeDirectory string `yaml:"codeDirectory"`
	TestSuites []string `yaml:"testSuites"`
}
