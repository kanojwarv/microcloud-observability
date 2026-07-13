package inspect

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/doctor"
	"github.com/kanojwarv/microcloud-observability/internal/graph"
	"github.com/kanojwarv/microcloud-observability/internal/impact"
	"github.com/kanojwarv/microcloud-observability/internal/search"
)

type Check struct {
	Name    string
	Passed  bool
	Message string
}

type Report struct {
	Checks []Check
}

func Run(
	repositoryRoot string,
) Report {

	report := Report{}

	_, err := graph.Build(
		repositoryRoot,
	)

	report.Checks = append(
		report.Checks,
		Check{
			Name:    "Graph",
			Passed:  err == nil,
			Message: errorMessage(err),
		},
	)

	_, err = search.Find(
		repositoryRoot,
		"cpu",
	)

	report.Checks = append(
		report.Checks,
		Check{
			Name:    "Search",
			Passed:  err == nil,
			Message: errorMessage(err),
		},
	)

	_, err = impact.Analyze(
		repositoryRoot,
		"vm_cpu_usage_seconds_rate",
	)

	report.Checks = append(
		report.Checks,
		Check{
			Name:    "Impact",
			Passed:  err == nil,
			Message: errorMessage(err),
		},
	)

	health := doctor.Run(
		repositoryRoot,
	)

	passed := health.Status != "ERROR"

	report.Checks = append(
		report.Checks,
		Check{
			Name:   "Doctor",
			Passed: passed,
			Message: fmt.Sprintf(
				"repository status: %s",
				health.Status,
			),
		},
	)

	return report
}

func errorMessage(
	err error,
) string {

	if err == nil {
		return ""
	}

	return err.Error()
}
