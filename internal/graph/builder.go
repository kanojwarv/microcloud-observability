package graph

import (
	"github.com/kanojwarv/microcloud-observability/internal/loader"
)

func Build() (*Graph, error) {

	g := New()

	dashboards, err := loader.LoadDashboards(
		"dashboards",
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
