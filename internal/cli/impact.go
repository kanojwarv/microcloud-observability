package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/impact"
)

func executeImpact(
	artifact string,
) error {

	dependents, err := impact.Analyze(
		".",
		artifact,
	)

	if err != nil {
		return err
	}

	fmt.Println(
		"Impact Analysis",
	)

	fmt.Println(
		"================",
	)

	fmt.Println()

	fmt.Printf(
		"Artifact: %s\n\n",
		artifact,
	)

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
