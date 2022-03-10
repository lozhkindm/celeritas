package main

import (
	"embed"
	"io/ioutil"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(tmplPath, targetPath string) error {
	// TODO: check if file already exists
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
