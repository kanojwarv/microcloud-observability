package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/doctor"
)

func executeDoctor() error {

	report := doctor.Run(".")

	fmt.Println(
		"Repository Health",
	)

	fmt.Println(
		"=================",
	)

	fmt.Println()

	fmt.Printf(
		"Overall Status: %s\n\n",
		report.Status,
	)

	for _, check := range report.Checks {

		fmt.Printf(
			"[%s] %s\n",
			check.Status,
			check.Name,
		)

		fmt.Printf(
			"    %s\n\n",
			check.Message,
		)
	}

	return nil
}
