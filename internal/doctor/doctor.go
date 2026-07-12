package doctor

func Run(
	repositoryRoot string,
) Report {

	checks := []Check{

		CheckDashboards(
			repositoryRoot,
		),

		CheckMetrics(
			repositoryRoot,
		),

		CheckRecordingRules(
			repositoryRoot,
		),

		CheckVariables(
			repositoryRoot,
		),

		CheckAlerts(
			repositoryRoot,
		),
	}

	return Report{
		Status: overallStatus(
			checks,
		),
		Checks: checks,
	}
}
