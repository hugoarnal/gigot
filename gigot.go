package main

import (
	"fmt"
	"gigot/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid amount of arguments")
		os.Exit(1)
	}

	cmds := []commands.Commands{
		{FlagSet: commands.InitFlagSet(), Name: "init", Run: commands.Init},
	}

	for _, cmd := range cmds {
		if os.Args[1] == cmd.Name {
			cmd.Run(cmd.FlagSet)
			return
		}
	}

	fmt.Println("Incorrect subcommand")
	os.Exit(1)
}
