package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func GetPrint(selectedConfig config.GitConfigConfig, name bool, path bool) {
	if name {
		fmt.Println(selectedConfig.Name)
		return
	}

	if path {
		fmt.Println(selectedConfig.Path)
		return
	}

	fmt.Println(selectedConfig.Name, selectedConfig.Path)
}

func GetFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("get", flag.ExitOnError)

	return cmd
}

func Get(cmd *flag.FlagSet) {
	name := cmd.Bool("name", false, "Get by name")
	path := cmd.Bool("path", false, "Get by path")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	if cmd.NArg() != 1 {
		fmt.Println("Invalid amount of arguments")
		os.Exit(1)
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		panic(err)
	}

	configIndex, err := config.FindGitConfigByName(parsedConfig, cmd.Arg(0))

	if err != nil {
		panic(err)
	}

	selectedConfig := parsedConfig[configIndex]

	GetPrint(selectedConfig, *name, *path)
}
