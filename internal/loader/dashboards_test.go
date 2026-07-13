package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDashboards(
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
	)

	dashboards, err := LoadDashboards(
		path,
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if len(dashboards) == 0 {
		t.Fatal(
			"expected at least one dashboard",
		)
	}

	found := false

	for _, dashboard := range dashboards {

		if dashboard.Name == "executive" {

			found = true
			break
		}
	}

	if !found {
		t.Fatal(
			"expected executive dashboard",
		)
	}
}

func TestLoadDashboardsMissingDirectory(
	t *testing.T,
) {

	_, err := LoadDashboards(
		"does-not-exist",
	)

	if err == nil {
		t.Fatal(
			"expected error",
		)
	}
}
