package main

func doMigrate(arg2, arg3 string) error {
	dsn := getDSN()

	switch arg2 {
	case "up":
		if err := cel.MigrateUp(dsn); err != nil {
			return err
		}
	case "down":
		if arg3 == "all" {
			if err := cel.MigrateDownAll(dsn); err != nil {
				return err
			}
		} else {
			if err := cel.MigrateSteps(-1, dsn); err != nil {
				return err
			}
		}
	case "reset":
		if err := cel.MigrateDownAll(dsn); err != nil {
			return err
		}
		if err := cel.MigrateUp(dsn); err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
