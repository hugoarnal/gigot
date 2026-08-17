package commands

import (
	"fmt"
	"gigot/internal/config"
	"os"

	"github.com/spf13/cobra"
)

func TempCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "temp",
		Short: "Temporarily (only for this shell session) switch to another gitconfig",
		RunE:  Temp,
	}
	cmd.Flags().Bool("disable", false, "Disable the currently enabled gitconfig")
	return cmd
}

func Temp(cmd *cobra.Command, args []string) error {
	disable, _ := cmd.Flags().GetBool("disable")

	config.GigotShellWarning()

	filename := config.GetGitConfigFilename()

	parsedConfig, err := config.ParseGitConfigFile(filename)

	if err != nil {
		return err
	}

	if disable {
		fmt.Println("Disabling the currently used config")
		return nil
	}

	switch len(args) {
	case 0:
		err := SwitchList(parsedConfig, filename)

		if err != nil {
			fmt.Println("")
			os.Exit(1)
		}
	case 1:
		selectedConfig := args[0]

		_, err := config.FindGitConfigByName(parsedConfig, selectedConfig)

		if err != nil {
			return err
		}

		fmt.Printf("Switching for this shell session to \"%s\"\n", selectedConfig)
	default:
		return fmt.Errorf("incorrect amount of arguments")
	}

	return nil
}
