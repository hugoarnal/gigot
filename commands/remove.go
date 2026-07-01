package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func RemoveFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("remove", flag.ExitOnError)

	return cmd
}

func Remove(cmd *flag.FlagSet) {
	name := cmd.String("name", "", "Name of the associated gitconfig")

	cmd.Parse(os.Args[2:])

	if *name == "" {
		fmt.Println("Incorrect usage, please use the \"--help\" flag as reference")
		os.Exit(1)
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		panic(err)
	}

	index, err := config.FindGitConfigByName(parsedConfig, *name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parsedConfig = append(parsedConfig[:index], parsedConfig[index+1:]...)

	err = config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Removed \"%s\" from the configuration\n", *name)
}
