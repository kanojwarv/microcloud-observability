package validate

import (
	"context"
	"testing"
)

func TestVariablesValidator(
	t *testing.T,
) {

	validator := VariablesValidator{
		Path: "../../variables/variables.yaml",
	}

	result := validator.Validate(
		context.Background(),
	)

	if !result.Passed {
		t.Fatal(
			"expected variables validation to pass",
		)
	}

	if len(result.Checks) == 0 {
		t.Fatal(
			"expected validation checks",
		)
	}
}
