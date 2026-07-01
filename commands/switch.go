package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func SwitchFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("switch", flag.ExitOnError)

	return cmd
}

func Switch(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	filename := config.GetGitConfigFilename()

	parsedConfig, err := config.ParseGitConfigFile(filename)

	if err != nil {
		// TODO: create the configuration file here instead
		panic(err)
	}

	if cmd.NArg() == 0 {
		// Listing all of them for now
		// TODO: switch to TUI
		fmt.Println("You must specify the gitconfig to switch to")
		fmt.Println()
		fmt.Println("The currently available configs are:")

		for _, c := range parsedConfig {
			fmt.Printf("    %s\n", c.Name)
		}

		os.Exit(1)
	} else if cmd.NArg() == 1 {
		selectedConfig := cmd.Arg(0)

		if err := config.SwitchGitConfigConfig(&parsedConfig, selectedConfig); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := config.WriteGitConfigFile(filename, parsedConfig); err != nil {
			panic(err)
		}

		configIndex, err := config.FindGitConfigByName(parsedConfig, selectedConfig)

		if err != nil {
			panic(err)
		}

		particle := "on"
		if parsedConfig[configIndex].Selected == false {
			particle = "off"
		}

		fmt.Printf("Toggled %s \"%s\" configuration\n", particle, selectedConfig)
	} else {
		fmt.Println("Incorrect amount of arguments")
		os.Exit(1)
	}
}
