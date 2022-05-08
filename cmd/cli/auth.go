package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func doAuth() error {
	// create auth migrations
	dbType := cel.DB.DataType
	if dbType == "mariadb" {
		dbType = "mysql"
	}
	if dbType == "postgresql" {
		dbType = "postgres"
	}

	filename := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upTmpl := fmt.Sprintf("templates/migrations/auth/%s.up.sql", dbType)
	downTmpl := fmt.Sprintf("templates/migrations/auth/%s.down.sql", dbType)
	upFile := fmt.Sprintf("%s/migrations/%s.%s.up.sql", cel.RootPath, filename, dbType)
	downFile := fmt.Sprintf("%s/migrations/%s.%s.down.sql", cel.RootPath, filename, dbType)
	if err := copyFileFromTemplate(upTmpl, upFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(downTmpl, downFile); err != nil {
		return err
	}

	// run auth migrations
	if err := doMigrate("up", ""); err != nil {
		return err
	}

	userTmpl := "templates/data/user.go.txt"
	tokenTmpl := "templates/data/token.go.txt"
	rememberTmpl := "templates/data/remember_token.go.txt"
	userFile := fmt.Sprintf("%s/data/user.go", cel.RootPath)
	tokenFile := fmt.Sprintf("%s/data/token.go", cel.RootPath)
	rememberFile := fmt.Sprintf("%s/data/remember_token.go", cel.RootPath)
	if err := copyFileFromTemplate(userTmpl, userFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(tokenTmpl, tokenFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(rememberTmpl, rememberFile); err != nil {
		return err
	}

	// create auth middlewares
	authMwTmpl := "templates/middlewares/auth.go.txt"
	tokenMwTmpl := "templates/middlewares/auth-token.go.txt"
	rememberMwTmpl := "templates/middlewares/remember.go.txt"
	authMwFile := fmt.Sprintf("%s/middlewares/auth.go", cel.RootPath)
	tokenMwFile := fmt.Sprintf("%s/middlewares/auth-token.go", cel.RootPath)
	rememberMwFile := fmt.Sprintf("%s/middlewares/remember.go", cel.RootPath)
	if err := copyFileFromTemplate(authMwTmpl, authMwFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(tokenMwTmpl, tokenMwFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(rememberMwTmpl, rememberMwFile); err != nil {
		return err
	}

	authHsTmpl := "templates/handlers/auth-handlers.go.txt"
	authHsFile := fmt.Sprintf("%s/handlers/auth-handlers.go", cel.RootPath)
	if err := copyFileFromTemplate(authHsTmpl, authHsFile); err != nil {
		return err
	}

	resetMlHtmlTmpl := "templates/mails/reset-password.html.txt"
	resetMlPlainTmpl := "templates/mails/reset-password.plain.txt"
	resetMlHtmlFile := fmt.Sprintf("%s/mails/reset-password.html.tmpl", cel.RootPath)
	resetMlPlainFile := fmt.Sprintf("%s/mails/reset-password.plain.tmpl", cel.RootPath)
	if err := copyFileFromTemplate(resetMlHtmlTmpl, resetMlHtmlFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(resetMlPlainTmpl, resetMlPlainFile); err != nil {
		return err
	}

	forgotVwTmpl := "templates/views/forgot.jet.txt"
	loginVwTmpl := "templates/views/login.jet.txt"
	resetVwTmpl := "templates/views/reset-password.jet.txt"
	forgotVwFile := fmt.Sprintf("%s/views/forgot.jet", cel.RootPath)
	loginVwFile := fmt.Sprintf("%s/views/login.jet", cel.RootPath)
	resetVwFile := fmt.Sprintf("%s/views/reset-password.jet", cel.RootPath)
	if err := copyFileFromTemplate(forgotVwTmpl, forgotVwFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(loginVwTmpl, loginVwFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(resetVwTmpl, resetVwFile); err != nil {
		return err
	}

	color.Yellow("  - users, tokens, remember_tokens migrations created and executed")
	color.Yellow("  - user and token models created")
	color.Yellow("  - auth and auth-token middlewares created")
	color.Yellow("")
	color.Yellow("Don't forget:")
	color.Yellow("1. Add user and token models in data/models.go")
	color.Yellow("2. Add appropriate middlewares to your routes")

	return nil
}
