package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/doctor"
)

func executeDoctor() error {

	health := doctor.Check()

	fmt.Println(
		"Repository Health",
	)

	fmt.Println(
		"=================",
	)

	fmt.Println()

	fmt.Println(
		health.Status,
	)

	fmt.Println()

	fmt.Println(
		health.Message,
	)

	return nil
}
