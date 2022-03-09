package celeritas

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
)

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
