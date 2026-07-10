package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/validate"
)

// Execute runs the CLI command.
func Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: mco <version|validate>")
	}

	switch args[0] {

	case "version":
		fmt.Println("MicroCloud Observability")
		fmt.Println("Version: dev")

		return nil

	case "validate":
		engine := validate.New()

		engine.Register(
			validate.NewRepositoryValidator(),
		)

		results := engine.Run(
			context.Background(),
		)

		fmt.Println("MicroCloud Observability")
		fmt.Println("========================")

		for _, result := range results {
			fmt.Printf(
				"%s: %t\n",
				result.Name,
				result.Passed,
			)
		}

		return nil

	default:
		return fmt.Errorf(
			"unknown command: %s",
			args[0],
		)
	}
}
