package validate

import (
	"context"
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
	"github.com/kanojwarv/microcloud-observability/internal/model"
)

type MetricsValidator struct {
	Path string
}

func NewMetricsValidator() MetricsValidator {
	return MetricsValidator{
		Path: "metrics/compute.yaml",
	}
}

func (m MetricsValidator) Name() string {
	return "metrics"
}

func (m MetricsValidator) Validate(ctx context.Context) Result {
	result := Result{
		Name:   "Metrics",
		Passed: true,
	}

	metricsFile, err := loader.LoadMetricsFile(m.Path)
	if err != nil {
		result.Passed = false

		result.Checks = append(result.Checks, Check{
			Name:    "load metrics",
			Passed:  false,
			Message: err.Error(),
		})

		return result
	}

	for _, metric := range metricsFile.Metrics {

		m.validateMetricName(metric, &result)
		m.validateObject(metric, &result)
		m.validateCategory(metric, &result)
		m.validateRecordingRule(metric, &result)
	}

	return result
}

func (m MetricsValidator) validateMetricName(
	metric model.Metric,
	result *Result,
) {
	ok := metric.Metric != ""

	result.Checks = append(result.Checks, Check{
		Name:    fmt.Sprintf("%s metric name", metric.Metric),
		Passed:  ok,
		Message: "metric name must not be empty",
	})

	if !ok {
		result.Passed = false
	}
}

func (m MetricsValidator) validateObject(
	metric model.Metric,
	result *Result,
) {
	valid := map[model.Object]bool{
		model.ObjectVM:      true,
		model.ObjectHost:    true,
		model.ObjectCluster: true,
		model.ObjectPool:    true,
		model.ObjectOSD:     true,
	}

	ok := valid[metric.Object]

	result.Checks = append(result.Checks, Check{
		Name:    fmt.Sprintf("%s object", metric.Metric),
		Passed:  ok,
		Message: "invalid object",
	})

	if !ok {
		result.Passed = false
	}
}

func (m MetricsValidator) validateCategory(
	metric model.Metric,
	result *Result,
) {
	valid := map[model.Category]bool{
		model.CategoryPerformance: true,
		model.CategoryStorage:     true,
		model.CategoryNetwork:     true,
		model.CategoryOperations:  true,
	}

	ok := valid[metric.Category]

	result.Checks = append(result.Checks, Check{
		Name:    fmt.Sprintf("%s category", metric.Metric),
		Passed:  ok,
		Message: "invalid category",
	})

	if !ok {
		result.Passed = false
	}
}

func (m MetricsValidator) validateRecordingRule(
	metric model.Metric,
	result *Result,
) {
	ok := metric.RecordingRule != ""

	result.Checks = append(result.Checks, Check{
		Name:    fmt.Sprintf("%s recording rule", metric.Metric),
		Passed:  ok,
		Message: "recording rule must not be empty",
	})

	if !ok {
		result.Passed = false
	}
}
