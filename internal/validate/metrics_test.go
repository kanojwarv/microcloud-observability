package validate

import (
	"context"
	"testing"
)

func TestMetricsValidator(t *testing.T) {
	validator := MetricsValidator{
		Path: "../../metrics/compute.yaml",
	}

	result := validator.Validate(context.Background())

	if !result.Passed {
		t.Fatal("expected metrics validation to pass")
	}
}
