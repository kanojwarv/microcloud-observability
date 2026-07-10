package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/graph"
)

func executeGraph() error {

	g, err := graph.Build()

	if err != nil {
		return err
	}

	fmt.Println("Dependency Graph")
	fmt.Println("================")

	dependencies := g.DependenciesOf(
		"executive",
	)

	for _, dependency := range dependencies {

		fmt.Printf(
			"executive -> %s\n",
			dependency,
		)
	}

	return nil
}
