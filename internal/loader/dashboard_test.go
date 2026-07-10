package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDashboardFile(
	t *testing.T,
) {

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(
			"failed to get cwd: %v",
			err,
		)
	}

	path := filepath.Join(
		cwd,
		"..",
		"..",
		"dashboards",
		"executive.dashboard.yaml",
	)

	dashboard, err := LoadDashboardFile(
		path,
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if dashboard.Dashboard.Name != "executive" {
		t.Fatalf(
			"expected dashboard %q, got %q",
			"executive",
			dashboard.Dashboard.Name,
		)
	}

	if len(
		dashboard.Dashboard.Sections,
	) == 0 {

		t.Fatal(
			"expected sections",
		)
	}
}

func TestLoadDashboardFileMissing(
	t *testing.T,
) {

	_, err := LoadDashboardFile(
		"does-not-exist.yaml",
	)

	if err == nil {
		t.Fatal(
			"expected error",
		)
	}
}
