package main

import (
	"errors"
	"os"

	"github.com/lozhkindm/celeritas"

	"github.com/fatih/color"
)

const version = "1.0.7"

var cel celeritas.Celeritas

func main() {
	var message string

	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	if err := setup(); err != nil {
		exitGracefully(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version: %s", version)
	case "migrate":
		if arg2 == "" {
			arg2 = "up"
		}
		if err := doMigrate(arg2, arg3); err != nil {
			exitGracefully(err)
		}
		message = "Migrations complete!"
	case "make":
		if arg2 == "" {
			exitGracefully(errors.New("make requires a subcommand: (migration|model|handler)"))
		}
		if err := doMake(arg2, arg3); err != nil {
			exitGracefully(err)
		}
	default:
		showHelp()
	}

	exitGracefully(nil, message)
}

func validateInput() (string, string, string, error) {
	var (
		arg1, arg2, arg3 string
		err              error
	)

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		showHelp()
		err = errors.New("command required")
	}

	return arg1, arg2, arg3, err
}

func exitGracefully(err error, msg ...string) {
	var m string

	if len(msg) > 0 {
		m = msg[0]
	}

	if err != nil {
		color.Red("Error: %v\n", err)
	}

	if len(m) > 0 {
		color.Yellow(m)
	} else {
		color.Green("Finished")
	}

	os.Exit(0)
}
