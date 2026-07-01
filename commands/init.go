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
