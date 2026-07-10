package loader

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func TestLoadMetricsFile(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}

	path := filepath.Join(
		cwd,
		"..",
		"..",
		"metrics",
		"compute.yaml",
	)

	metrics, err := LoadMetricsFile(path)
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

	if first.Object != model.ObjectVM {
		t.Fatalf(
			"expected object %q, got %q",
			model.ObjectVM,
			first.Object,
		)
	}

	if first.RecordingRule == "" {
		t.Fatal("expected recording rule")
	}
}

func TestLoadMetricsFileMissing(t *testing.T) {
	_, err := LoadMetricsFile("does-not-exist.yml")

	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
