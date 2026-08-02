package commands

import (
	"fmt"
	"gigot/internal/config"
	"os"

	"github.com/spf13/cobra"
)

func SwitchList(parsedConfig []config.GitConfigConfig) {
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
}

func SwitchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch",
		Short: "Switch to another gitconfig",
		RunE:  Switch,
	}
	cmd.Flags().Bool("disable", false, "Disable the currently enabled gitconfig")
	return cmd
}

func Switch(cmd *cobra.Command, args []string) error {
	disable, _ := cmd.Flags().GetBool("disable")

	config.GigotShellWarning()

	filename := config.GetGitConfigFilename()

	parsedConfig, err := config.ParseGitConfigFile(filename)

	if err != nil {
		return err
	}

	if disable {
		for i := range parsedConfig {
			parsedConfig[i].Enabled = false
		}

		if err := config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig); err != nil {
			return err
		}

		fmt.Println("Disabled all enabled configurations")
		return nil
	}

	switch len(args) {
	case 0:
		SwitchList(parsedConfig)
	case 1:
		selectedConfig := args[0]

		if err := config.SwitchGitConfigConfig(&parsedConfig, selectedConfig); err != nil {
			return err
		}

		if err := config.WriteGitConfigFile(filename, parsedConfig); err != nil {
			return err
		}

		configIndex, err := config.FindGitConfigByName(parsedConfig, selectedConfig)

		if err != nil {
			return err
		}

		particle := "on"
		if !parsedConfig[configIndex].Enabled {
			particle = "off"
		}

		fmt.Printf("Toggled %s \"%s\" configuration\n", particle, selectedConfig)
	default:
		return fmt.Errorf("incorrect amount of arguments")
	}

	return nil
}
