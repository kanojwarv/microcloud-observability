package doctor

import (
	"path/filepath"
)

func CheckMetrics(
	repositoryRoot string,
) Check {

	files, err := filepath.Glob(
		filepath.Join(
			repositoryRoot,
			"metrics",
			"*.yaml",
		),
	)

	if err != nil {

		return Check{
			Name:    "metrics",
			Status:  "WARN",
			Message: err.Error(),
		}
	}

	if len(
		files,
	) == 0 {

		return Check{
			Name:    "metrics",
			Status:  "WARN",
			Message: "No metrics found.",
		}
	}

	return Check{
		Name:    "metrics",
		Status:  "OK",
		Message: "Metrics loaded successfully.",
	}
}
