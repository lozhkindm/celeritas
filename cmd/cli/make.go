package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gertd/go-pluralize"
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
		if err := copyDataToFile([]byte(contents), filename); err != nil {
			return err
		}
	case "model":
		if arg3 == "" {
			return errors.New("you must give the model a name")
		}

		pl := pluralize.NewClient()
		modelName, tableName := strings.ToLower(arg3), strings.ToLower(arg3)

		if pl.IsPlural(arg3) {
			modelName = pl.Singular(modelName)
		} else {
			tableName = pl.Plural(tableName)
		}

		filename := fmt.Sprintf("%s/data/%s.go", cel.RootPath, modelName)
		if fileExists(filename) {
			return errors.New(fmt.Sprintf("%s already exists", filename))
		}

		data, err := templateFS.ReadFile("templates/data/model.go.txt")
		if err != nil {
			return err
		}

		contents := strings.ReplaceAll(string(data), "$MODEL_NAME$", strcase.ToCamel(modelName))
		contents = strings.ReplaceAll(contents, "$TABLE_NAME$", tableName)
		if err := copyDataToFile([]byte(contents), filename); err != nil {
			return err
		}
	case "session":
		if err := doSession(); err != nil {
			return err
		}
	}
	return nil
}
