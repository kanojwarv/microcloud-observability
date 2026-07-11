package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kanojwarv/microcloud-observability/internal/generate"
)

func executeGenerate(
	target string,
) error {

	switch target {

	case "graph":

		data, err := generate.GraphJSON(
			".",
		)

		if err != nil {
			return err
		}

		outputDir := "artifacts"

		err = os.MkdirAll(
			outputDir,
			0755,
		)

		if err != nil {
			return err
		}

		outputPath := filepath.Join(
			outputDir,
			"graph.json",
		)

		err = os.WriteFile(
			outputPath,
			data,
			0644,
		)

		if err != nil {
			return err
		}

		fmt.Println(
			"Generated:",
		)

		fmt.Printf(
			"    %s\n",
			outputPath,
		)

		return nil

	case "graph-html":

		data, err := generate.GraphHTML(
			".",
		)

		if err != nil {
			return err
		}

		outputPath := filepath.Join(
			"artifacts",
			"graph.html",
		)

		err = os.WriteFile(
			outputPath,
			data,
			0644,
		)

		if err != nil {
			return err
		}

		fmt.Println(
			"Generated:",
		)

		fmt.Printf(
			"    %s\n",
			outputPath,
		)

		return nil
	}

	return fmt.Errorf(
		"unknown generation target: %s",
		target,
	)
}
