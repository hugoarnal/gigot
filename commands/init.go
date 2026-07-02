package commands

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

type initShell struct {
	name    string
	content string
}

func InitFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("init", flag.ExitOnError)

	return cmd
}

//go:embed shell/gigot.bash
var shellBash string

// gigot.zsh is a copy of gigot.bash
// bash and zsh have quite the same syntax which makes sense
// I tried doing a symlink but it is strictly forbidden by the embed module:
// https://github.com/golang/go/issues/44507
//
//go:embed shell/gigot.zsh
var shellZsh string

func Init(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	if cmd.NArg() != 1 {
		fmt.Println("No argument given")
		os.Exit(1)
	}

	shells := []initShell{
		{
			name:    "bash",
			content: shellBash,
		},
		{
			name:    "zsh",
			content: shellZsh,
		},
	}

	for _, shell := range shells {
		if cmd.Arg(0) == shell.name {
			fmt.Print(shell.content)
			return
		}
	}

	fmt.Println("Unknown shell")
	os.Exit(1)
}
