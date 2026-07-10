package main

import (
	"fmt"
	"github.com/kanojwarv/microcloud-observability/internal/cli"
	"os"
)

func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
