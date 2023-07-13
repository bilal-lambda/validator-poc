package rules

import (
	"fmt"
	"strings"

	"github.com/bilal-lambda/validator-poc/models"
	"gopkg.in/yaml.v3"
)

func Validate(source string) error {
	var yml models.SampleYamlStruct
	decoder := yaml.NewDecoder(strings.NewReader(source))
	decoder.KnownFields(true)
	err := decoder.Decode(&yml)
	if err != nil {
		return fmt.Errorf("invalid yaml - %s", err)
	}
	
	if yml.Runson != "linux" && yml.Runson != "macos" && yml.Runson != "win" {
		return fmt.Errorf("runson can only be linux, macos or win")
	}

	if yml.Version == "" {
		return fmt.Errorf("missing version")
	}

	if yml.Version != "0.1" && yml.Version != "0.2" {
		return fmt.Errorf("incorrect version, use 0.1 or 0.2")
	}

	return nil
}
