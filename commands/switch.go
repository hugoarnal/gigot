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
	disable := cmd.Bool("disable", false, "Disable the currently enabled gitconfig")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	if !config.IsGigotShellSet() {
		fmt.Println("WARNING: the \"GIGOT_SHELL\" environment variable isn't set")
		fmt.Println("That means the shell session probably wasn't initialized using \"gigot init\"")
		fmt.Println("Please look at the documentation to setup your shell environment properly")
		fmt.Println("Proceeding with the switch command")
		fmt.Println()
	}

	filename := config.GetGitConfigFilename()

	parsedConfig, err := config.ParseGitConfigFile(filename)

	if err != nil {
		panic(err)
	}

	if *disable {
		for i := range parsedConfig {
			parsedConfig[i].Enabled = false
		}

		if err := config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig); err != nil {
			panic(err)
		}

		fmt.Println("Disabled all enabled configurations")
		return
	}

	if cmd.NArg() == 0 {
		// Listing all of them for now
		// TODO: switch to TUI
		fmt.Println("You must specify the gitconfig to switch to")
		fmt.Println()
		fmt.Println("The currently available configs are:")

		for _, c := range parsedConfig {
			fmt.Printf("    %s", c.Name)
			if c.Enabled {
				fmt.Printf(" (enabled)")
			}
			fmt.Printf("\n")
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
		if !parsedConfig[configIndex].Enabled {
			particle = "off"
		}

		fmt.Printf("Toggled %s \"%s\" configuration\n", particle, selectedConfig)
	} else {
		fmt.Println("Incorrect amount of arguments")
		os.Exit(1)
	}
}
