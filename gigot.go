package main

import (
	"gigot/commands"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "gigot",
		Short: "An easy gitconfig switcher",
	}

	cmds := []*cobra.Command{
		commands.InitCommand(),
		commands.AddCommand(),
		commands.RemoveCommand(),
		commands.ListCommand(),
		commands.GetCommand(),
		commands.GetEnabledCommand(),
		commands.SwitchCommand(),
		commands.TempCommand(),
	}

	// Disable autocomplete
	// Doesn't really make sense for now
	root.CompletionOptions.HiddenDefaultCmd = true
	// Silence usage whenever an error occurs
	root.SilenceUsage = true

	for _, cmd := range cmds {
		root.AddCommand(cmd)
	}

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
