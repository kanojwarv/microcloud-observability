package doctor

func overallStatus(
	checks []Check,
) string {

	for _, check := range checks {

		if check.Status == "ERROR" {

			return "ERROR"
		}
	}

	for _, check := range checks {

		if check.Status == "WARN" {

			return "WARN"
		}
	}

	return "OK"
}
