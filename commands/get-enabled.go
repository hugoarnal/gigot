// TODO: this command could probably be merged with gigot get
// something like gigot get --current
package commands

import (
	"gigot/internal/config"

	"github.com/spf13/cobra"
)

func GetEnabledCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-enabled",
		Short: "Get associated information of currently enabled config",
		RunE:  GetEnabled,
	}
	cmd.Flags().Bool("name", false, "Get only name of the gitconfig")
	cmd.Flags().Bool("path", false, "Get only path of the gitconfig")
	cmd.Flags().Bool("enabled", false, "Get only enabled status of the gitconfig")
	return cmd
}

func GetEnabled(cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetBool("name")
	path, _ := cmd.Flags().GetBool("path")
	enabled, _ := cmd.Flags().GetBool("enabled")

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		return err
	}

	selectedConfig := config.GitConfigConfig{
		Name: "", Path: "", Enabled: false,
	}

	for i := range parsedConfig {
		if parsedConfig[i].Enabled {
			selectedConfig = parsedConfig[i]
			break
		}
	}

	GetPrint(selectedConfig, name, path, enabled)

	return nil
}
