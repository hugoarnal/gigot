package commands

import (
	"fmt"
	"gigot/internal/config"

	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Adds a new gitconfig to the configuration file",
		RunE:  Add,
	}
	cmd.Flags().String("name", "", "Name of the associated gitconfig")
	cmd.Flags().String("path", "", "Path of the associated gitconfig")
	return cmd
}

func Add(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	path, _ := cmd.Flags().GetString("path")

	if name == "" || path == "" {
		return fmt.Errorf("incorrect usage, missing name or path")
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		return err
	}

	parsedConfig = append(parsedConfig, config.GitConfigConfig{Name: name, Path: path, Enabled: false})

	err = config.WriteGitConfigFile(config.GetGitConfigFilename(), parsedConfig)

	if err != nil {
		return err
	}

	fmt.Printf("Added \"%s\" to the configuration\n", name)

	return nil
}
