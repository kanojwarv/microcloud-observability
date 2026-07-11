package cli

import (
	"context"
	"errors"
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

		engine.Register(
			validate.NewMetricsValidator(),
		)

		engine.Register(
			validate.NewRecordingRulesValidator(),
		)

		engine.Register(
			validate.NewVariablesValidator(),
		)

		engine.Register(
			validate.NewDashboardValidator(),
		)

		report := engine.Run(
			context.Background(),
		)

		fmt.Println("MicroCloud Observability")
		fmt.Println("========================")

		for _, result := range report.Results {

			status := "FAIL"

			if result.Passed {
				status = "PASS"
			}

			fmt.Printf(
				"%-15s %s\n",
				result.Name,
				status,
			)
		}

		return nil

	case "graph":
		return executeGraph()

	case "doctor":
		return executeDoctor()

	case "impact":

		if len(args) < 2 {
			return errors.New(
				"usage: mco impact <artifact>",
			)
		}

		return executeImpact(
			args[1],
		)
	case "generate":

		if len(args) < 2 {
			return errors.New(
				"usage: mco generate <graph>",
			)
		}

		return executeGenerate(
			args[1],
		)

	case "serve":
		return executeServe()

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
		impact       Show artifact impact
		generate.    Generate artifacts
	    graph        Show dependency graph
        doctor       Check repository health

		help         Show this help message

    Examples:

        mco validate
        mco version
		mco generate graph
		mco generate graph-html
        mco help

    `)
}
