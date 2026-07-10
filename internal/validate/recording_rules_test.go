package validate

import (
	"context"
	"testing"
)

func TestRecordingRulesValidator(t *testing.T) {
	validator := RecordingRulesValidator{
		Path: "../../recording-rules/compute.rules.yaml",
	}

	result := validator.Validate(context.Background())

	if !result.Passed {
		for _, check := range result.Checks {
			t.Logf(
				"%s passed=%v message=%s",
				check.Name,
				check.Passed,
				check.Message,
			)
		}

		t.Fatal("expected recording rules validation to pass")
	}

	if len(result.Checks) == 0 {
		t.Fatal("expected validation checks")
	}
}
