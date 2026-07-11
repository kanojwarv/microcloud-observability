package impact

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/graph"
)

func Analyze(
	repositoryRoot string,
	artifact string,
) ([]string, error) {

	g, err := graph.Build(
		repositoryRoot,
	)

	if err != nil {
		return nil, err
	}

	dependents := []string{}

	for parent, dependencies := range g.Dependencies() {

		for _, dependency := range dependencies {

			if dependency == artifact {

				dependents = append(
					dependents,
					parent,
				)
			}
		}
	}

	if len(dependents) == 0 {

		return nil, fmt.Errorf(
			"artifact not found: %s",
			artifact,
		)
	}

	return dependents, nil
}
