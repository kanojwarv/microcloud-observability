package doctor

import "testing"

func TestCheck(
	t *testing.T,
) {

	report := Run(
		"../..",
	)

	if report.Status != "WARN" {

		t.Fatalf(
			"expected WARN, got %s",
			report.Status,
		)
	}

	found := false

	for _, check := range report.Checks {

		if check.Name == "alerts" {

			found = true

			if check.Status != "WARN" {

				t.Fatalf(
					"expected alerts to be WARN",
				)
			}
		}
	}

	if !found {

		t.Fatal(
			"alerts check missing",
		)
	}
}
