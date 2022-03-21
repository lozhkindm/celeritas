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
	upFile := fmt.Sprintf("%s/migrations/%s.%s.up.sql", cel.RootPath, filename, dbType)
	downFile := fmt.Sprintf("%s/migrations/%s.%s.down.sql", cel.RootPath, filename, dbType)
	upTmpl := fmt.Sprintf("templates/migrations/auth/%s.up.sql", dbType)
	downTmpl := fmt.Sprintf("templates/migrations/auth/%s.down.sql", dbType)
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
	userFile := fmt.Sprintf("%s/data/user.go", cel.RootPath)
	tokenFile := fmt.Sprintf("%s/data/token.go", cel.RootPath)
	if err := copyFileFromTemplate(userTmpl, userFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(tokenTmpl, tokenFile); err != nil {
		return err
	}

	// create auth middlewares
	authMwTmpl := "templates/middlewares/auth.go.txt"
	tokenMwTmpl := "templates/middlewares/auth-token.go.txt"
	authMwFile := fmt.Sprintf("%s/middlewares/auth.go", cel.RootPath)
	tokenMwFile := fmt.Sprintf("%s/middlewares/auth-token.go", cel.RootPath)
	if err := copyFileFromTemplate(authMwTmpl, authMwFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(tokenMwTmpl, tokenMwFile); err != nil {
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
