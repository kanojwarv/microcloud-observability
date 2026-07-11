package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/server"
)

func executeServe() error {

	fmt.Println(
		"MicroCloud Observability",
	)

	fmt.Println()

	fmt.Println(
		"Serving on:",
	)

	fmt.Println(
		"    http://localhost:8080",
	)

	return server.Start(
		":8080",
	)
}
