package graph

import (
	"github.com/kanojwarv/microcloud-observability/internal/loader"
)

func Build() (*Graph, error) {

	g := New()

	dashboard, err := loader.LoadDashboardFile(
		"dashboards/executive.dashboard.yaml",
	)

	if err != nil {
		return nil, err
	}

	for _, section := range dashboard.Dashboard.Sections {

		for _, panel := range section.Panels {

			g.AddDependency(
				dashboard.Dashboard.Name,
				panel.Metric,
			)
		}
	}

	return g, nil
}
