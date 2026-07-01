package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

// Terribly named
// Essentially, how the configuration keeps metadata about gitconfig files
type GitConfigConfig struct {
	Name     string
	Path     string
	Selected bool
}

func GetGitConfigFilename() string {
	return GetPath() + "/gitconfig.yml"
}

func ParseGitConfigFile(filename string) ([]GitConfigConfig, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var v []GitConfigConfig

	err = yaml.Unmarshal([]byte(data), &v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func WriteGitConfigFile(filename string, config []GitConfigConfig) error {
	bytes, err := yaml.Marshal(config)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, bytes, 0644)

	return err
}
