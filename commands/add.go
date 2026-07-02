package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func AddFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("add", flag.ExitOnError)

	return cmd
}

func Add(cmd *flag.FlagSet) {
	name := cmd.String("name", "", "Name of the associated gitconfig")
	path := cmd.String("path", "", "Path to the associated gitconfig")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	if *name == "" || *path == "" {
		fmt.Println("Incorrect usage, please use the \"--help\" flag as reference")
		os.Exit(1)
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		panic(err)
	}

	parsedConfig = append(parsedConfig, config.GitConfigConfig{Name: *name, Path: *path, Enabled: false})

	err = config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Added \"%s\" to the configuration\n", *name)
}
