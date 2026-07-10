package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/graph"
)

func executeImpact(
	artifact string,
) error {

	g, err := graph.Build()

	if err != nil {
		return err
	}

	dependents := g.DependentsOf(
		artifact,
	)

	fmt.Println("Impact Analysis")
	fmt.Println("================")
	fmt.Println()

	fmt.Printf(
		"Artifact: %s\n\n",
		artifact,
	)

	if len(dependents) == 0 {

		fmt.Println(
			"No dependents found.",
		)

		return nil
	}

	fmt.Println(
		"Referenced by:",
	)

	for _, dependent := range dependents {

		fmt.Printf(
			"  - %s\n",
			dependent,
		)
	}

	return nil
}
