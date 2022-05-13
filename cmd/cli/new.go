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
	return nil
}
