package main

import (
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(tmplPath, targetPath string) error {
	if fileExists(targetPath) {
		return errors.New(fmt.Sprintf("%s already exists", targetPath))
	}

	data, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return err
	}
	if err := copyDataToFile(data, targetPath); err != nil {
		return err
	}
	return nil
}

func copyDataToFile(data []byte, to string) error {
	if err := ioutil.WriteFile(to, data, 0644); err != nil {
		return err
	}
	return nil
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
