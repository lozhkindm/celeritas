package main

import (
	"fmt"
	"time"
)

func doSession() error {
	dbType := cel.DB.DataType
	if dbType == "mariadb" {
		dbType = "mysql"
	}
	if dbType == "postgresql" {
		dbType = "postgres"
	}

	filename := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())
	upFile := fmt.Sprintf("%s/migrations/%s.%s.up.sql", cel.RootPath, filename, dbType)
	downFile := fmt.Sprintf("%s/migrations/%s.%s.down.sql", cel.RootPath, filename, dbType)
	upTmpl := fmt.Sprintf("templates/migrations/session/%s.up.sql", dbType)
	downTmpl := fmt.Sprintf("templates/migrations/session/%s.down.sql", dbType)
	if err := copyFileFromTemplate(upTmpl, upFile); err != nil {
		return err
	}
	if err := copyFileFromTemplate(downTmpl, downFile); err != nil {
		return err
	}

	if err := doMigrate("up", ""); err != nil {
		return err
	}

	return nil
}
