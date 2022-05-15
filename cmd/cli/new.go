package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var appURL string

func doNew(appName string) error {
	appName = strings.ToLower(appName)
	appURL = appName
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

	color.Green("Creating Makefile...")

	makefileExt := "mac"
	if runtime.GOOS == "windows" {
		makefileExt = "windows"
	}

	src, err := os.Open(fmt.Sprintf("./%s/Makefile.%s", appName, makefileExt))
	if err != nil {
		return err
	}
	defer func() {
		_ = src.Close()
	}()
	dst, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
	defer func() {
		_ = dst.Close()
	}()
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	_ = os.Remove(fmt.Sprintf("./%s/Makefile.mac", appName))
	_ = os.Remove(fmt.Sprintf("./%s/Makefile.windows", appName))

	color.Green("Creating go.mod file...")

	_ = os.Remove(fmt.Sprintf("./%s/go.mod", appName))
	contents, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		return err
	}
	mod := string(contents)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)
	if err := copyDataToFile([]byte(mod), fmt.Sprintf("./%s/go.mod", appName)); err != nil {
		return err
	}

	color.Green("Updating source files...")

	if err := updateSource(); err != nil {
		return err
	}

	return nil
}
