package main

import (
	"fmt"
)

func doSession() error {
	if err := checkDB(); err != nil {
		return err
	}
	tx, err := cel.PopConnect()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Close()
	}()

	// create auth migrations
	dbType := cel.DB.DataType
	if dbType == "mariadb" {
		dbType = "mysql"
	}
	if dbType == "postgresql" {
		dbType = "postgres"
	}

	upTmpl := fmt.Sprintf("templates/migrations/session/%s.up.sql", dbType)
	up, err := templateFS.ReadFile(upTmpl)
	if err != nil {
		return err
	}
	downTmpl := fmt.Sprintf("templates/migrations/session/%s.down.sql", dbType)
	down, err := templateFS.ReadFile(downTmpl)
	if err != nil {
		return err
	}
	if err := cel.CreatePopMigration(up, down, "sessions", "sql"); err != nil {
		return err
	}
	if err := cel.MigrateUp(tx); err != nil {
		return err
	}

	return nil
}
