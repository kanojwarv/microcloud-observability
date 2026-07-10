package cli

import (
	"context"
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/validate"
)

func Execute(args []string) error {
	if len(args) == 0 {
		printHelp()
		return nil
	}

	switch args[0] {

	case "help", "--help", "-h":
		printHelp()
		return nil

	case "version":
		fmt.Println("MicroCloud Observability")
		fmt.Println("Version: dev")
		return nil

	case "validate":
		engine := validate.New()

		engine.Register(
			validate.NewRepositoryValidator(),
		)

		report := engine.Run(
			context.Background(),
		)

		fmt.Println("MicroCloud Observability")
		fmt.Println("========================")

		for _, result := range report.Results {
			fmt.Printf(
				"%s: %t\n",
				result.Name,
				result.Passed,
			)
		}

		return nil

	default:
		printHelp()
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func printHelp() {
	fmt.Println(`MicroCloud Observability

    Usage:

    mco <command>

    Commands:

        validate     Validate the repository
        version      Show version information
        help         Show this help message

    Examples:

        mco validate
        mco version
        mco help
    `)
}
