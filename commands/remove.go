package commands

import (
	"fmt"
	"gigot/internal/config"

	"github.com/spf13/cobra"
)

func RemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a gitconfig from the configuration file",
		RunE:  Remove,
	}
	cmd.Flags().String("name", "", "Name of the associated gitconfig")
	return cmd
}

func Remove(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")

	if name == "" {
		return fmt.Errorf("incorrect usage, missing name")
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		return err
	}

	index, err := config.FindGitConfigByName(parsedConfig, name)

	if err != nil {
		return err
	}

	parsedConfig = append(parsedConfig[:index], parsedConfig[index+1:]...)

	err = config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig)

	if err != nil {
		return err
	}

	fmt.Printf("Removed \"%s\" from the configuration\n", name)
	return nil
}
