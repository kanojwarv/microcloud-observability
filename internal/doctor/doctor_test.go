package doctor

import "testing"

func TestCheck(
	t *testing.T,
) {

	health := Check()

	if health.Status != "OK" {

		t.Fatalf(
			"expected OK, got %s",
			health.Status,
		)
	}
}
