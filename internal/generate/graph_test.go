package generate

import (
	"testing"
)

func TestGraphJSON(
	t *testing.T,
) {

	data, err := GraphJSON()

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if len(data) == 0 {

		t.Fatal(
			"expected JSON output",
		)
	}
}
