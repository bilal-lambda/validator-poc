package rules

import (
	"fmt"

	"github.com/bilal-lambda/validator-poc/models"
	"gopkg.in/yaml.v3"
)

func Validate(source string) (bool, error) {
	var yml models.SampleYamlStruct
	err := yaml.Unmarshal([]byte(source), &yml)
	if err != nil {
		return false, fmt.Errorf("incorrect yaml - %s", err)
	}
	
	if yml.Runson != "linux" && yml.Runson != "macos" && yml.Runson != "win" {
		return false, fmt.Errorf("runson can only be linux, macos or win")
	}

	if yml.Version == "" {
		return false, fmt.Errorf("missing version")
	}

	if yml.Version != "0.1" && yml.Version != "0.2" {
		return false, fmt.Errorf("incorrect version")
	}

	return true, nil
}
