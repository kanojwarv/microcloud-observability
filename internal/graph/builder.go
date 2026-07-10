package graph

import (
	"os"
	"path/filepath"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
)

func Build() (*Graph, error) {

	g := New()

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dashboardsDir := filepath.Join(
		cwd,
		"..",
		"..",
		"dashboards",
	)

	dashboards, err := loader.LoadDashboards(
		dashboardsDir,
	)

	if err != nil {
		return nil, err
	}

	for _, dashboard := range dashboards {

		for _, section := range dashboard.Dashboard.Sections {

			for _, panel := range section.Panels {

				g.AddDependency(
					dashboard.Dashboard.Name,
					panel.Metric,
				)
			}
		}
	}

	return g, nil
}
