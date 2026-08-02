package commands

import (
	"fmt"
	"gigot/internal/config"

	"github.com/spf13/cobra"
)

func GetPrint(selectedConfig config.GitConfigConfig, name bool, path bool, enabled bool) {
	if name {
		fmt.Println(selectedConfig.Name)
		return
	}

	if path {
		fmt.Println(selectedConfig.Path)
		return
	}

	if enabled {
		fmt.Println(selectedConfig.Enabled)
		return
	}

	fmt.Println(selectedConfig.Name, selectedConfig.Path, selectedConfig.Enabled)
}

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get associated information of the given config",
		RunE:  Get,
	}
	cmd.Flags().Bool("name", false, "Get only name of the gitconfig")
	cmd.Flags().Bool("path", false, "Get only path of the gitconfig")
	cmd.Flags().Bool("enabled", false, "Get only enabled status of the gitconfig")
	return cmd
}

func Get(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetBool("name")
	path, _ := cmd.Flags().GetBool("path")
	enabled, _ := cmd.Flags().GetBool("enabled")

	if len(args) != 1 {
		return fmt.Errorf("invalid amount of arguments")
	}

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		return err
	}

	configIndex, err := config.FindGitConfigByName(parsedConfig, args[0])

	if err != nil {
		return err
	}

	selectedConfig := parsedConfig[configIndex]

	GetPrint(selectedConfig, name, path, enabled)

	return nil
}
