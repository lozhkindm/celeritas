package main

import (
	"fmt"
	"time"
)

func doAuth() error {
	// create auth migrations
	dbType := cel.DB.DataType
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

	return nil
}
