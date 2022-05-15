package main

import (
	"fmt"
	"os"
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

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	cel.RootPath = path
	cel.DB.DataType = os.Getenv("DATABASE_TYPE")

	return nil
}

func getDSN() string {
	var dsn string

	dbType := cel.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf(
				"postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf(
				"postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		}
	} else {
		dsn = fmt.Sprintf("mysql://%s", cel.BuildDSN())
	}

	return dsn
}

func showHelp() {
	color.Yellow(`Available commands:

    help                  - show the help commands
    version               - print application version
    migrate               - runs all up migrations that have not been run previously
    migrate down          - reverses the most recent migration
    migrate reset         - runs all down migrations in reverse order, and then all up migrations
    make migration <name> - creates two new up and down migrations in the migrations folder
    make auth             - prepares auth functionality
    make handler <name>   - creates a stub handler in the handlers folder
    make model <name>     - creates a new model in the data folder
    make session          - creates a table in the database as a session store
    make mail <name>      - create two html and plain mail templates

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
