package main

import (
	"github.com/redcuckoo/go-labs-SE/lab3/cli"
	"github.com/redcuckoo/go-labs-SE/lab3/service"
	"os"
)

func main() {
	service.Main()

	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
