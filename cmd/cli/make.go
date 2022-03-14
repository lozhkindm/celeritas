package main

import (
	"errors"
	"fmt"
	"time"
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
		if err := copyFileFromTemplate(upTmpl, upFile); err != nil {
			return err
		}

		downTmpl := fmt.Sprintf("templates/migrations/migration.%s.down.sql", dbType)
		if err := copyFileFromTemplate(downTmpl, downFile); err != nil {
			return err
		}
	case "auth":
		if err := doAuth(); err != nil {
			return err
		}
	}
	return nil
}
