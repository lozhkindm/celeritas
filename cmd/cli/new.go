package main

import "strings"

func doNew(appName string) error {
	appName = strings.ToLower(appName)
	if strings.Contains(appName, "/") {
		parts := strings.Split(appName, "/")
		appName = parts[len(parts)-1]
	}

	return nil
}
