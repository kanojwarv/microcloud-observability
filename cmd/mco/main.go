package main

import (
	"fmt"
	"github.com/vishalkanojwar/microcloud-observability/internal/cli"
	"os"
)

func main() {
	if err := cli.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
