package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

func doNew(appName string) error {
	appName = strings.ToLower(appName)
	if strings.Contains(appName, "/") {
		parts := strings.Split(appName, "/")
		appName = parts[len(parts)-1]
	}

	color.Green("Cloning repository...")

	_, err := git.PlainClone(fmt.Sprintf("./%s", appName), false, &git.CloneOptions{
		URL:      "https://github.com/lozhkindm/celeritas-skeleton.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		return err
	}

	if err := os.RemoveAll(fmt.Sprintf("./%s/.git", appName)); err != nil {
		return err
	}

	color.Green("Creating .env file...")

	contents, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		return err
	}
	env := string(contents)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${KEY}", cel.RandStr(32))
	if err := copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName)); err != nil {
		return err
	}

	return nil
}
