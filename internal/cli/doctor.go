package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/graph"
)

func executeDoctor() error {

	g, err := graph.Build()

	if err != nil {
		return err
	}

	fmt.Println("Repository Health")
	fmt.Println("=================")
	fmt.Println()

	found := false

	dependencies := g.DependenciesOf(
		"executive",
	)

	if len(dependencies) == 0 {

		found = true

		fmt.Println(
			"WARNING",
		)

		fmt.Println(
			"dashboard has no dependencies",
		)
	}

	if !found {

		fmt.Println(
			"OK",
		)

		fmt.Println()

		fmt.Println(
			"All dashboard dependencies are used.",
		)
	}

	return nil
}
