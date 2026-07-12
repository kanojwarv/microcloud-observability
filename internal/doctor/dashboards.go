package doctor

import (
	"path/filepath"
)

func CheckDashboards(
	repositoryRoot string,
) Check {

	files, err := filepath.Glob(
		filepath.Join(
			repositoryRoot,
			"dashboards",
			"*.yaml",
		),
	)

	if err != nil {

		return Check{
			Name:    "dashboards",
			Status:  "WARN",
			Message: err.Error(),
		}
	}

	if len(
		files,
	) == 0 {

		return Check{
			Name:    "dashboards",
			Status:  "WARN",
			Message: "No dashboards found.",
		}
	}

	return Check{
		Name:    "dashboards",
		Status:  "OK",
		Message: "Dashboards loaded successfully.",
	}
}
