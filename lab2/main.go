package main

import (
	"github.com/redcuckoo/go-labs-SE/lab2/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
