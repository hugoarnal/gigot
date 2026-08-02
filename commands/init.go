package commands

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

type initShell struct {
	name    string
	content string
}

func InitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize your shell session",
		RunE:  Init,
	}
	return cmd
}

//go:embed shell/gigot.bash
var shellBash string

// It is intentional for the zsh config to depend on the bash one
// bash and zsh have quite the same syntax which makes sense
// I tried doing a symlink but it is strictly forbidden by the embed module:
// https://github.com/golang/go/issues/44507
//
//go:embed shell/gigot.bash
var shellZsh string

func initShellList() []initShell {
	return []initShell{
		{
			name:    "bash",
			content: shellBash,
		},
		{
			name:    "zsh",
			content: shellZsh,
		},
	}
}

func Init(cmd *cobra.Command, args []string) error {
	shells := initShellList()

	if len(args) != 1 {
		return fmt.Errorf("invalid amount of arguments")
	}

	for _, shell := range shells {
		if args[0] == shell.name {
			fmt.Print(shell.content)
			return nil
		}
	}

	return fmt.Errorf("unknown shell")
}
