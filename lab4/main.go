package main

import (
	"github.com/redcuckoo/go-labs-SE/lab4/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
