package doctor

import (
	"path/filepath"
)

func CheckVariables(
	repositoryRoot string,
) Check {

	files, err := filepath.Glob(
		filepath.Join(
			repositoryRoot,
			"variables",
			"*.yaml",
		),
	)

	if err != nil {

		return Check{
			Name:    "variables",
			Status:  "WARN",
			Message: err.Error(),
		}
	}

	if len(
		files,
	) == 0 {

		return Check{
			Name:    "variables",
			Status:  "WARN",
			Message: "No variables found.",
		}
	}

	return Check{
		Name:    "variables",
		Status:  "OK",
		Message: "Variables loaded successfully.",
	}
}
