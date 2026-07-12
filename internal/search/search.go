package search

func Find(
	repositoryRoot string,
	query string,
) (Response, error) {

	response := Response{
		Query: query,
	}

	results := []Result{}

	results = append(
		results,
		findMetrics(
			repositoryRoot,
			query,
		)...,
	)

	results = append(
		results,
		findDashboards(
			repositoryRoot,
			query,
		)...,
	)

	response.Results = results

	return response, nil
}
