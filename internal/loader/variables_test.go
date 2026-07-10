package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadVariablesFile(
	t *testing.T,
) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(
			"failed to get cwd: %v",
			err,
		)
	}

	path := filepath.Join(
		cwd,
		"..",
		"..",
		"variables",
		"variables.yaml",
	)

	variables, err := LoadVariablesFile(path)
	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if len(variables.Variables) == 0 {
		t.Fatal(
			"expected variables",
		)
	}
}

func TestLoadVariablesFileMissing(
	t *testing.T,
) {
	_, err := LoadVariablesFile(
		"does-not-exist.yaml",
	)

	if err == nil {
		t.Fatal(
			"expected error",
		)
	}
}
