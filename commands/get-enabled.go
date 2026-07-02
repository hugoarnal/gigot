package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func GetEnabledFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("get-enabled", flag.ExitOnError)

	return cmd
}

func GetEnabled(cmd *flag.FlagSet) {
	name := cmd.Bool("name", false, "Get the currently enabled name")
	path := cmd.Bool("path", false, "Get the currently enabled path")

	cmd.Parse(os.Args[2:])

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		panic(err)
	}

	selectedConfig := config.GitConfigConfig{
		Name: "", Path: "", Enabled: false,
	}

	for i := range parsedConfig {
		if parsedConfig[i].Enabled == true {
			selectedConfig = parsedConfig[i]
			break
		}
	}

	if *path == true {
		fmt.Println(selectedConfig.Path)
		return
	}

	if *name == true {
		fmt.Println(selectedConfig.Name)
		return
	}

	fmt.Println(selectedConfig.Name, selectedConfig.Path)
}
