package commands

import (
	"fmt"
	"gigot/internal/config"

	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available gitconfigs",
		RunE:  List,
	}
	return cmd
}

func List(cmd *cobra.Command, args []string) error {
	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		return err
	}

	for _, c := range parsedConfig {
		fmt.Printf("%s | %s | %t\n", c.Name, c.Path, c.Enabled)
	}

	return nil
}
