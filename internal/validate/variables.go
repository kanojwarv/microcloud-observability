package validate

import (
	"context"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
	"github.com/kanojwarv/microcloud-observability/internal/model"
)

type VariablesValidator struct {
	Path string
}

func NewVariablesValidator() VariablesValidator {
	return VariablesValidator{
		Path: "variables/variables.yaml",
	}
}

func (v VariablesValidator) Name() string {
	return "variables"
}

func (v VariablesValidator) Validate(
	ctx context.Context,
) Result {

	result := Result{
		Name:   "Variables",
		Passed: true,
	}

	variablesFile, err := loader.LoadVariablesFile(v.Path)
	if err != nil {

		result.Passed = false

		result.Checks = append(
			result.Checks,
			Check{
				Name:    "load variables",
				Passed:  false,
				Message: err.Error(),
			},
		)

		return result
	}

	seen := make(map[string]bool)

	for _, variable := range variablesFile.Variables {

		v.validateName(
			variable,
			&result,
		)

		v.validateLabel(
			variable,
			&result,
		)

		v.validateDuplicate(
			variable,
			seen,
			&result,
		)

		seen[variable.Name] = true
	}

	return result
}

func (v VariablesValidator) validateName(
	variable model.Variable,
	result *Result,
) {

	ok := variable.Name != ""

	result.Checks = append(
		result.Checks,
		Check{
			Name:    variable.Name,
			Passed:  ok,
			Message: "name must not be empty",
		},
	)

	if !ok {
		result.Passed = false
	}
}

func (v VariablesValidator) validateLabel(
	variable model.Variable,
	result *Result,
) {

	ok := variable.Label != ""

	result.Checks = append(
		result.Checks,
		Check{
			Name:    variable.Name + " label",
			Passed:  ok,
			Message: "label must not be empty",
		},
	)

	if !ok {
		result.Passed = false
	}
}

func (v VariablesValidator) validateDuplicate(
	variable model.Variable,
	seen map[string]bool,
	result *Result,
) {

	ok := !seen[variable.Name]

	result.Checks = append(
		result.Checks,
		Check{
			Name:    variable.Name + " duplicate",
			Passed:  ok,
			Message: "duplicate variable",
		},
	)

	if !ok {
		result.Passed = false
	}
}
