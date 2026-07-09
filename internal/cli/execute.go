package cli

import (
	"errors"
	"fmt"
	"github.com/vishalkanojwar/microcloud-observability/pkg/validate"
)

func Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("usage: mco <version|validate>")
	}
	switch args[0] {
	case "version":
		fmt.Println("MicroCloud Observability")
		fmt.Println("Version: dev")
		return nil
	case "validate":
		return validate.Run(".")
	default:
		return fmt.Errorf("unknown command: %s", args[0])
	}
}
