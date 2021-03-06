package main

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup(arg1 string) error {
	if arg1 == "new" || arg1 == "version" || arg1 == "help" {
		return nil
	}
	if err := godotenv.Load(); err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	cel.RootPath = wd
	cel.DB.DataType = os.Getenv("DATABASE_TYPE")

	return nil
}

func showHelp() {
	color.Yellow(`Available commands:

    help                           - show the help commands
	up                             - take the server out of maintenance mode
	down                           - put the server into maintenance mode
    version                        - print application version
    migrate                        - runs all up migrations that have not been run previously
    migrate down                   - reverses the most recent migration
    migrate reset                  - runs all down migrations in reverse order, and then all up migrations
    make migration <name> <format> - creates two new up and down migrations in the migrations folder (format=pop,sql)
    make auth                      - prepares auth functionality
    make handler <name>            - creates a stub handler in the handlers folder
    make model <name>              - creates a new model in the data folder
    make session                   - creates a table in the database as a session store
    make mail <name>               - create two html and plain mail templates

	`)
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return nil
	}
	match, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}
	if !match {
		return nil
	}
	if match {
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		contents := strings.Replace(string(file), "myapp", appURL, -1)
		if err := os.WriteFile(path, []byte(contents), 0); err != nil {
			return err
		}
	}
	return nil
}

func updateSource() error {
	return filepath.Walk(".", updateSourceFiles)
}

func checkDB() error {
	if cel.DB.DataType == "" {
		return errors.New("no database connection provided")
	}
	if !fileExists(path.Join(cel.RootPath, "config", "database.yml")) {
		return errors.New("database.yml does not exist in a config folder")
	}
	return nil
}
