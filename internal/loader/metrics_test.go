package loader

import (
	"testing"
)

func TestLoadMetricsFile(t *testing.T) {
	metrics, err := LoadMetricsFile("../../metrics/compute.yml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(metrics.Metrics) != 3 {
		t.Fatalf(
			"expected 3 metrics, got %d",
			len(metrics.Metrics),
		)
	}

	first := metrics.Metrics[0]

	if first.Metric != "lxd_cpu_seconds_total" {
		t.Fatalf(
			"expected metric %q, got %q",
			"lxd_cpu_seconds_total",
			first.Metric,
		)
	}

	if first.Object != "vm" {
		t.Fatalf(
			"expected object %q, got %q",
			"vm",
			first.Object,
		)
	}
}

func TestLoadMetricsFileMissing(t *testing.T) {
	_, err := LoadMetricsFile("does-not-exist.yml")

	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
