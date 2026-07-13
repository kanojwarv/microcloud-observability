package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/inspect"
)

func executeInspect() error {

	report := inspect.Run(
		".",
	)

	for _, check := range report.Checks {

		status := "OK"

		if !check.Passed {

			status = "WARN"
		}

		fmt.Printf(
			"[%s] %s\n",
			status,
			check.Name,
		)

		if check.Message != "" {

			fmt.Printf(
				"    %s\n",
				check.Message,
			)
		}
	}

	return nil
}
