package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		if arg3 == "" {
			return errors.New("you must give the migration a name")
		}

		dbType := cel.DB.DataType
		filename := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := fmt.Sprintf("%s/migrations/%s.%s.up.sql", cel.RootPath, filename, dbType)
		downFile := fmt.Sprintf("%s/migrations/%s.%s.down.sql", cel.RootPath, filename, dbType)
		upTmpl := fmt.Sprintf("templates/migrations/migration.%s.up.sql", dbType)
		downTmpl := fmt.Sprintf("templates/migrations/migration.%s.down.sql", dbType)
		if err := copyFileFromTemplate(upTmpl, upFile); err != nil {
			return err
		}
		if err := copyFileFromTemplate(downTmpl, downFile); err != nil {
			return err
		}
	case "auth":
		if err := doAuth(); err != nil {
			return err
		}
	case "handler":
		if arg3 == "" {
			return errors.New("you must give the handler a name")
		}

		filename := fmt.Sprintf("%s/handlers/%s.go", cel.RootPath, strings.ToLower(arg3))
		if fileExists(filename) {
			return errors.New(fmt.Sprintf("%s already exists", filename))
		}

		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			return err
		}

		contents := strings.ReplaceAll(string(data), "$HANDLER_NAME$", strcase.ToCamel(arg3))
		if err := ioutil.WriteFile(filename, []byte(contents), 0644); err != nil {
			return err
		}
	}
	return nil
}
