package main

import (
	"fmt"
	"syscall/js"

	"github.com/bilal-lambda/validator-poc/rules"
)

func validateYaml(_this js.Value, args []js.Value) interface{} {
	ok := js.ValueOf("ok")

	if len(args) == 0 {
		return ok
	}

	source := args[0].String()
	if _, err := rules.Validate(source); err != nil {
		return js.ValueOf(err.Error())
	}

	return ok
}

func main() {
	js.Global().Set("validateYaml", js.FuncOf(validateYaml))
	select {}
}
