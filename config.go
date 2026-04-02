package main

import (
	"log"
	"os"
	// "github.com/spf13/viper"
)

func init() {
	log.Println("Hello world!")
}

func createDir(path string) error {
	// Owner: read/write/execute; group: read/execute; others: read/execute.
	var perm os.FileMode = 0755

	f := func() error {
		if err := os.Mkdir(path, perm); err != nil {
			return err
		}
		return nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return f()
	} else if !info.IsDir() {
		return f()
	}
	return nil
}

func createFile(path string) error {
	// Owner: read/write/execute; group: read/execute; others: read/execute.
	var perm os.FileMode = 0755

	file, err := os.OpenFile(path, os.O_CREATE, perm)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
