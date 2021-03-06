package main

import (
	"errors"
	"os"

	"github.com/lozhkindm/celeritas"

	"github.com/fatih/color"
)

const version = "1.0.0"

var cel celeritas.Celeritas

func main() {
	var message string

	arg1, arg2, arg3, arg4, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	if err := setup(arg1); err != nil {
		exitGracefully(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "up":
		if err := maintenanceMode(false); err != nil {
			exitGracefully(err)
		}
	case "down":
		if err := maintenanceMode(true); err != nil {
			exitGracefully(err)
		}
	case "new":
		if arg2 == "" {
			exitGracefully(errors.New("new required an application name"))
		}
		if err := doNew(arg2); err != nil {
			exitGracefully(err)
		}
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
			exitGracefully(errors.New("make requires a subcommand: (migration|auth|handler|model|session|mail)"))
		}
		if err := doMake(arg2, arg3, arg4); err != nil {
			exitGracefully(err)
		}
	default:
		showHelp()
	}

	exitGracefully(nil, message)
}

func validateInput() (string, string, string, string, error) {
	var (
		arg1, arg2, arg3, arg4 string
		err                    error
	)

	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}

		if len(os.Args) >= 5 {
			arg4 = os.Args[4]
		}
	} else {
		showHelp()
		err = errors.New("command required")
	}

	return arg1, arg2, arg3, arg4, err
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
