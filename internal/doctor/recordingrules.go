package doctor

import (
	"path/filepath"
)

func CheckRecordingRules(
	repositoryRoot string,
) Check {

	files, err := filepath.Glob(
		filepath.Join(
			repositoryRoot,
			"recordingrules",
			"*.yaml",
		),
	)

	if err != nil {

		return Check{
			Name:    "recordingrules",
			Status:  "WARN",
			Message: err.Error(),
		}
	}

	if len(
		files,
	) == 0 {

		return Check{
			Name:    "recordingrules",
			Status:  "WARN",
			Message: "No recordingrules found.",
		}
	}

	return Check{
		Name:    "recordingrules",
		Status:  "OK",
		Message: "RecordingRules loaded successfully.",
	}
}
