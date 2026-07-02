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

	cmd.Usage = func() {
		shells := initShellList()

		fmt.Println("Available shells:")

		for _, shell := range shells {
			fmt.Printf("    %s\n", shell.name)
		}
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

func Init(cmd *flag.FlagSet) {
	if err := cmd.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	shells := initShellList()

	if cmd.NArg() != 1 {
		fmt.Println("No argument given")
		os.Exit(1)
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
