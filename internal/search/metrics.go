package search

import (
	"strings"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
)

func findMetrics(
	repositoryRoot string,
	query string,
) []Result {

	results := []Result{}

	metrics, err := loader.LoadMetrics(
		repositoryRoot,
	)

	if err != nil {
		return results
	}

	query = strings.ToLower(
		query,
	)

	for _, metric := range metrics {

		query = strings.ToLower(
			query,
		)

		if strings.Contains(
			strings.ToLower(
				metric.Metric,
			),
			query,
		) {

			results = append(
				results,
				Result{
					Kind: "metric",
					Name: metric.Metric,
					Path: metric.Source,
				},
			)
		}
	}

	return results

}
