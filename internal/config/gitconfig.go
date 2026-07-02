package config

import (
	"fmt"
	utils "gigot/internal"
	"os"

	"github.com/goccy/go-yaml"
)

// Terribly named
// Essentially, how the configuration keeps metadata about gitconfig files
type GitConfigConfig struct {
	Name    string
	Path    string
	Enabled bool
}

func GetGitConfigFilename() string {
	return GetPath() + "/gitconfig.yml"
}

func CheckGitConfigFile() bool {
	return utils.CheckPath(GetGitConfigFilename())
}

func CreateGitConfigFile() error {
	if !CheckPath() {
		if err := CreatePath(); err != nil {
			return err
		}
	}

	f, err := os.Create(GetGitConfigFilename())

	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err = f.Chmod(0644)

	if err != nil {
		return err
	}

	return nil
}

func ParseGitConfigFile(filename string) ([]GitConfigConfig, error) {
	if !CheckGitConfigFile() {
		if err := CreateGitConfigFile(); err != nil {
			return nil, err
		}
	}

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

func FindGitConfigByName(config []GitConfigConfig, configName string) (int, error) {
	configIndex := -1

	// Finds the config in the array
	for i := range config {
		if config[i].Name == configName {
			configIndex = i
			break
		}
	}

	if configIndex == -1 {
		return configIndex, fmt.Errorf("couldn't find \"%s\" in the configuration", configName)
	}

	return configIndex, nil
}

func SwitchGitConfigConfig(config *[]GitConfigConfig, configName string) error {
	configIndex, err := FindGitConfigByName(*config, configName)

	if err != nil {
		return err
	}

	// Toggles the enabled value from the given config
	// Removes enabled from all other configs
	for i := range *config {
		if i == configIndex {
			(*config)[i].Enabled = !(*config)[i].Enabled
		} else {
			(*config)[i].Enabled = false
		}
	}

	return nil
}
