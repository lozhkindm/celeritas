package celeritas

import (
	"fmt"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/pop"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (c *Celeritas) popConnect() (*pop.Connection, error) {
	tx, err := pop.Connect("development")
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Celeritas) CreatePopMigration(up, down []byte, name, ext string) error {
	if err := pop.MigrationCreate(path.Join(c.RootPath, "migrations"), name, ext, up, down); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) RunPopMigrations(tx *pop.Connection) error {
	migrator, err := pop.NewFileMigrator(path.Join(c.RootPath, "migrations"), tx)
	if err != nil {
		return err
	}
	if err := migrator.Up(); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) RollbackPopMigrations(tx *pop.Connection, steps ...int) error {
	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}
	migrator, err := pop.NewFileMigrator(path.Join(c.RootPath, "migrations"), tx)
	if err != nil {
		return err
	}
	if err := migrator.Down(step); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) MigrateUp(dsn string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s/migrations", c.RootPath), dsn)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) MigrateDownAll(dsn string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s/migrations", c.RootPath), dsn)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Down(); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) MigrateSteps(n int, dsn string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s/migrations", c.RootPath), dsn)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Steps(n); err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) MigrateForce(dsn string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s/migrations", c.RootPath), dsn)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = m.Close()
	}()

	if err := m.Force(-1); err != nil {
		return err
	}
	return nil
}
