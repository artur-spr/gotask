package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	configPath       = "~/.config/gotask"
	configFileName   = "config"
	configFileFormat = "yaml"
	// configFullPath   = fmt.Sprintf("%s/%s.%s", configPath, configFileName, configFileFormat)
)

func ConfigInit() {
	if err := createDir(configPath); err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileFormat)
	viper.AddConfigPath(configPath)
}

func createDir(path string) error {
	var perm os.FileMode = 0755 // Owner: read/write/execute; group: read/execute; others: read/execute.

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
