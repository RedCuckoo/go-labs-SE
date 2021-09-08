package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/redcuckoo/go-labs-SE/lab2/config"
	"log"
)

func Run(args []string) bool{
	defer func() {
		if rvr := recover(); rvr != nil {
			log.Fatalf("app paniced with recover: %v", rvr)
		}
	}()

	configure, err := config.Configure()
	if err != nil {
		log.Fatal(err)
		return false
	}

	app := kingpin.New("lab2", "")

	migrateCmd := app.Command("migrate", "migrate command")
	migrateUpCmd := migrateCmd.Command("up", "migrate database up")
	migrateDownCmd := migrateCmd.Command("down", "migrate database down")

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.Fatalf("failed to parse arguments, err: %v", err)
		return false
	}

	switch cmd {
	case migrateUpCmd.FullCommand():
		err = MigrateUp(configure)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(configure)
	default:
		log.Fatalf("unknown command %v", cmd)
		return false
	}

	if err != nil {
		log.Fatalf("failed to exec cmd, err: %v", err)
		return false
	}

	return true
}
