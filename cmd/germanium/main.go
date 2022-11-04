package main

import (
	"fmt"
	"os"

	"github.com/Hiroya-W/germanium/cli"
)

var exit = os.Exit

func main() {
	if err := cli.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exit(1)
	}
}
