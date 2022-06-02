package main

func doMigrate(arg2, arg3 string) error {
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

	switch arg2 {
	case "up":
		if err := cel.MigrateUp(tx); err != nil {
			return err
		}
	case "down":
		var steps int
		if arg3 == "all" {
			steps = -1
		} else {
			steps = 1
		}
		if err := cel.MigrateDown(tx, steps); err != nil {
			return err
		}
	case "reset":
		if err := cel.MigrateReset(tx); err != nil {
			return err
		}
	default:
		showHelp()
	}

	return nil
}
