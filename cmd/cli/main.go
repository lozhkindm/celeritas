package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/lozhkindm/celeritas"
)

const version = "1.0.0"

var cel celeritas.Celeritas

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}

	switch arg1 {
	case "help":
		showHelp()
	case "version":
		color.Yellow("Application version: %s", version)
	default:
		fmt.Println(arg2, arg3)
	}
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
		color.Red("Error: command required")
		showHelp()
		err = errors.New("command required")
	}

	return arg1, arg2, arg3, err
}

func showHelp() {
	color.Yellow(`Available commands:
	help         - show the help commands
	version      - print application version
	`)
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