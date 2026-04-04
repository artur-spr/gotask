package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func init() {
	configName, configType := "config", "yaml"
	configDir, err := homedir.Expand("~/.config/gotask")
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(configDir, configName) + "." + configType

	if err := mkdirAll(configDir); err != nil {
		log.Fatal(err)
	}

	if err := openFile(configPath); err != nil {
		log.Fatal(err)
	}
}

// Makes all dirs by [path] using os.MakeAll(path, 0755).
// Expands ~/ in path into /home/user/
func mkdirAll(path string) error {
	// Owner: read/write/execute; group: read/execute; others: read/execute.
	var perm os.FileMode = 0755
	path, err := homedir.Expand(path)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(path, perm); err != nil {
		return err
	}

	return nil
}

// Opens file [path] if does not exist using os.openFile(path, os.O_CREATE, 0755).
// Expands ~/ in path into /home/user/
func openFile(path string) error {
	// Owner: read/write/execute; group: read/execute; others: read/execute.
	var perm os.FileMode = 0755
	path, err := homedir.Expand(path)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
