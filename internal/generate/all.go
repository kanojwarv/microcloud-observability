package generate

import (
	"os"
	"path/filepath"
)

func All(
	repositoryRoot string,
) error {

	outputDir := "artifacts"

	err := os.MkdirAll(
		outputDir,
		0755,
	)

	if err != nil {
		return err
	}

	graphJSON, err := GraphJSON(
		repositoryRoot,
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(
			outputDir,
			"graph.json",
		),
		graphJSON,
		0644,
	)

	if err != nil {
		return err
	}

	graphHTML, err := GraphHTML(
		repositoryRoot,
	)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		filepath.Join(
			outputDir,
			"graph.html",
		),
		graphHTML,
		0644,
	)

	if err != nil {
		return err
	}

	return nil
}
