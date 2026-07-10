package validate

import (
	"context"

	"github.com/kanojwarv/microcloud-observability/internal/loader"
	"github.com/kanojwarv/microcloud-observability/internal/model"
)

type DashboardValidator struct {
	Path string
}

func NewDashboardValidator() DashboardValidator {
	return DashboardValidator{
		Path: "dashboards/executive.dashboard.yaml",
	}
}

func (d DashboardValidator) Name() string {
	return "dashboards"
}

func (d DashboardValidator) Validate(
	ctx context.Context,
) Result {

	result := Result{
		Name:   "Dashboards",
		Passed: true,
	}

	dashboardFile, err := loader.LoadDashboardFile(
		d.Path,
	)

	if err != nil {

		result.Passed = false

		result.Checks = append(
			result.Checks,
			Check{
				Name:    "load dashboard",
				Passed:  false,
				Message: err.Error(),
			},
		)

		return result
	}

	d.validateDashboard(
		dashboardFile.Dashboard,
		&result,
	)

	return result
}

func (d DashboardValidator) validateDashboard(
	dashboard model.Dashboard,
	result *Result,
) {

	d.validateDashboardName(
		dashboard.Name,
		result,
	)

	d.validateDashboardTitle(
		dashboard.Title,
		result,
	)

	for _, section := range dashboard.Sections {

		for _, panel := range section.Panels {

			d.validatePanelTitle(
				panel.Title,
				result,
			)

			d.validatePanelType(
				panel.Type,
				result,
			)
		}
	}
}

func (d DashboardValidator) validateDashboardName(
	name string,
	result *Result,
) {
	ok := name != ""

	result.Checks = append(
		result.Checks,
		Check{
			Name:    "dashboard name",
			Passed:  ok,
			Message: "dashboard name required",
		},
	)

	if !ok {
		result.Passed = false
	}
}

func (d DashboardValidator) validateDashboardTitle(
	title string,
	result *Result,
) {
	ok := title != ""

	result.Checks = append(
		result.Checks,
		Check{
			Name:    "dashboard title",
			Passed:  ok,
			Message: "dashboard title required",
		},
	)

	if !ok {
		result.Passed = false
	}
}

func (d DashboardValidator) validatePanelTitle(
	title string,
	result *Result,
) {
	ok := title != ""

	result.Checks = append(
		result.Checks,
		Check{
			Name:    "panel title",
			Passed:  ok,
			Message: "panel title required",
		},
	)

	if !ok {
		result.Passed = false
	}
}

func (d DashboardValidator) validatePanelType(
	panelType string,
	result *Result,
) {
	ok := panelType != ""

	result.Checks = append(
		result.Checks,
		Check{
			Name:    "panel type",
			Passed:  ok,
			Message: "panel type required",
		},
	)

	if !ok {
		result.Passed = false
	}
}
