package doctor

import (
	"path/filepath"
)

func CheckAlerts(
	repositoryRoot string,
) Check {

	files, err := filepath.Glob(
		filepath.Join(
			repositoryRoot,
			"alerts",
			"*.yaml",
		),
	)

	if err != nil {

		return Check{
			Name:    "alerts",
			Status:  "WARN",
			Message: err.Error(),
		}
	}

	if len(
		files,
	) == 0 {

		return Check{
			Name:    "alerts",
			Status:  "WARN",
			Message: "No alerts found.",
		}
	}

	return Check{
		Name:    "alerts",
		Status:  "OK",
		Message: "Alerts loaded successfully.",
	}
}
