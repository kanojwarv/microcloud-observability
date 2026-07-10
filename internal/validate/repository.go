package validate

import (
	"context"
	"os"
)

type RepositoryValidator struct {
	Root string
}

func NewRepositoryValidator() RepositoryValidator {
	return RepositoryValidator{
		Root: ".",
	}
}

func (r RepositoryValidator) Name() string {
	return "repository"
}

func (r RepositoryValidator) Validate(ctx context.Context) Result {
	required := []string{
		"docs",
		"metrics",
		"recording-rules",
		"variables",
	}

	result := Result{
		Name:   "Repository",
		Passed: true,
	}

	for _, path := range required {
		_, err := os.Stat(r.Root + "/" + path)

		ok := err == nil

		result.Checks = append(result.Checks, Check{
			Name:   path,
			Passed: ok,
		})

		if !ok {
			result.Passed = false
		}
	}

	return result
}
