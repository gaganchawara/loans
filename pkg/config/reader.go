package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

const (
	FilePath    = "%s/config/%s" // The format for the configuration file path
	EnvFilename = "%s.toml"      // The format for the environment-specific configuration file name
)

// LoadConfig loads configuration values from the environment-specific files
func LoadConfig(basePath string, env string, cfg interface{}) error {
	if err := LoadConfigFromFile(FilePath, basePath, EnvFilename, cfg, env); err != nil {
		return err
	}

	return nil
}

// LoadConfigFromFile loads configuration values from a TOML file located at the given path
// and populates the provided 'configStruct' with these values.
func LoadConfigFromFile(
	filePath string,
	basePath string,
	filename string,
	configStruct interface{},
	env string) error {
	path := fmt.Sprintf(filePath, basePath, fmt.Sprintf(filename, env))

	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	if _, err := toml.DecodeFile(path, configStruct); err != nil {
		return err
	}

	return nil
}
