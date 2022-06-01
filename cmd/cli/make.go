package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func doMake(arg2, arg3, arg4 string) error {
	switch arg2 {
	case "key":
		color.Yellow("Encryption key: %s", cel.RandStr(32))
	case "migration":
		if err := checkDB(); err != nil {
			return err
		}
		if arg3 == "" {
			return errors.New("you must give the migration a name")
		}
		var (
			up, down []byte
			ext      = "pop"
			err      error
		)
		if arg4 == "pop" || arg4 == "" {
			if up, err = templateFS.ReadFile("templates/migrations/migration.up.pop"); err != nil {
				return err
			}
			if down, err = templateFS.ReadFile("templates/migrations/migration.down.pop"); err != nil {
				return err
			}
		} else {
			ext = "sql"
		}
		if err := cel.CreatePopMigration(up, down, arg3, ext); err != nil {
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
	case "mail":
		if arg3 == "" {
			return errors.New("you must give the mail template a name")
		}

		htmlFile := fmt.Sprintf("%s/mails/%s.html.tmpl", cel.RootPath, strings.ToLower(arg3))
		plainFile := fmt.Sprintf("%s/mails/%s.plain.tmpl", cel.RootPath, strings.ToLower(arg3))
		htmlTmpl := "templates/mails/mail.html.txt"
		plainTmpl := "templates/mails/mail.plain.txt"
		if err := copyFileFromTemplate(htmlTmpl, htmlFile); err != nil {
			return err
		}
		if err := copyFileFromTemplate(plainTmpl, plainFile); err != nil {
			return err
		}

	}
	return nil
}
