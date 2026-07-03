package commands

import (
	"flag"
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

	if err := cmd.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		panic(err)
	}

	selectedConfig := config.GitConfigConfig{
		Name: "", Path: "", Enabled: false,
	}

	for i := range parsedConfig {
		if parsedConfig[i].Enabled {
			selectedConfig = parsedConfig[i]
			break
		}
	}

	GetPrint(selectedConfig, *name, *path)
}
