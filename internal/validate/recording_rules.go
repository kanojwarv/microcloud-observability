package validate

import (
	"context"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
	"github.com/kanojwarv/microcloud-observability/internal/model"
)

type RecordingRulesValidator struct {
	Path string
}

func NewRecordingRulesValidator() RecordingRulesValidator {
	return RecordingRulesValidator{
		Path: "recording-rules/compute.rules.yaml",
	}
}

func (r RecordingRulesValidator) Name() string {
	return "recording-rules"
}

func (r RecordingRulesValidator) Validate(ctx context.Context) Result {
	result := Result{
		Name:   "RecordingRules",
		Passed: true,
	}

	rulesFile, err := loader.LoadRecordingRuleFile(r.Path)
	if err != nil {
		result.Passed = false

		result.Checks = append(result.Checks, Check{
			Name:    "load recording rules",
			Passed:  false,
			Message: err.Error(),
		})

		return result
	}

	for _, group := range rulesFile.Groups {
		r.validateGroup(group, &result)
	}

	return result
}

func (r RecordingRulesValidator) validateGroup(
	group model.RuleGroup,
	result *Result,
) {
	r.validateGroupName(group, result)
	r.validateInterval(group, result)

	for _, rule := range group.Rules {
		r.validateRule(rule, result)
	}
}

func (r RecordingRulesValidator) validateGroupName(
	group model.RuleGroup,
	result *Result,
) {
	ok := group.Name != ""

	result.Checks = append(result.Checks, Check{
		Name:    "group name",
		Passed:  ok,
		Message: "group name must not be empty",
	})

	if !ok {
		result.Passed = false
	}
}

func (r RecordingRulesValidator) validateInterval(
	group model.RuleGroup,
	result *Result,
) {
	ok := group.Interval != ""

	result.Checks = append(result.Checks, Check{
		Name:    group.Name + " interval",
		Passed:  ok,
		Message: "interval must not be empty",
	})

	if !ok {
		result.Passed = false
	}
}

func (r RecordingRulesValidator) validateRule(
	rule model.RecordingRule,
	result *Result,
) {
	ok := rule.Record != "" && rule.Expr != ""

	result.Checks = append(result.Checks, Check{
		Name:    rule.Record,
		Passed:  ok,
		Message: "record and expr must not be empty",
	})

	if !ok {
		result.Passed = false
	}
}
