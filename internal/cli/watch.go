package cli

import (
	"fmt"

	"github.com/kanojwarv/microcloud-observability/internal/watch"
)

func executeWatch() error {

	fmt.Println(
		"Watching repository...",
	)

	fmt.Println()

	fmt.Println(
		"Regenerating artifacts every 5 seconds.",
	)

	watch.Start(".")

	return nil
}
