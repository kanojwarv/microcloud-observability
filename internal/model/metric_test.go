package model

import "testing"

func TestMetricDefinition(t *testing.T) {
	metric := Metric{
		Metric:        "lxd_cpu_seconds_total",
		Object:        ObjectVM,
		Category:      CategoryPerformance,
		Unit:          "cpu_seconds",
		RecordingRule: "vm_cpu_usage_seconds_rate",
		Dashboards: []string{
			"executive",
			"vm",
		},
		Alerts: []string{
			"high_vm_cpu",
		},
	}

	if metric.Object != ObjectVM {
		t.Fatalf("expected object %q, got %q", ObjectVM, metric.Object)
	}

	if metric.Category != CategoryPerformance {
		t.Fatalf(
			"expected category %q, got %q",
			CategoryPerformance,
			metric.Category,
		)
	}
}
