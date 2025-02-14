package main

import (
	"github.com/jforde/tflint-ruleset-hackathon/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "template",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewStandardModuleStructureRule(),
			},
		},
	})
}
