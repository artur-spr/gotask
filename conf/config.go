package conf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var Global string

func init() {
	files := make([]string, 0)
	files = append(files, "~/.config/gotask/config.yaml")
	files = append(files, "~/.local/share/gotask/tasks.json")

	if err := createAll(files); err != nil {
		log.Fatal(err)
	}
}

// Creates all directoryes by the way to the file
// and open file in created directory without rewriting it.
func createAll(paths []string) error {
	// Owner: read/write/execute; group: read/execute; others: read/execute.
	var perm os.FileMode = 0755

	for _, path := range paths {
		path, err := homedir.Expand(path)
		if err != nil {
			return err
		}

		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, perm); err != nil {
			return err
		}

		file, err := os.OpenFile(path, os.O_CREATE, perm)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
