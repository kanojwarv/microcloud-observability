package search

import (
	"strings"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
)

func findDashboards(
	repositoryRoot string,
	query string,
) []Result {

	results := []Result{}

	dashboards, err := loader.LoadDashboards(
		"dashboards",
	)

	if err != nil {
		return results
	}

	query = strings.ToLower(
		query,
	)

	for _, dashboard := range dashboards {

		if strings.Contains(
			strings.ToLower(
				dashboard.Title,
			),
			query,
		) {

			results = append(
				results,
				Result{
					Kind: "dashboard",
					Name: dashboard.Title,
					Path: dashboard.Source,
				},
			)
		}
	}

	return results
}
