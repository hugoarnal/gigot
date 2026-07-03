package main

import (
	"flag"
	"fmt"
	"gigot/commands"
	"os"
)

func printUsage(cmds []commands.Commands) {
	fmt.Println("Usage: gigot [subcommand]")
	fmt.Println("")
	fmt.Println("Subcommands:")
	for _, cmd := range cmds {
		fmt.Printf("    %s\n", cmd.Name)
	}
}

func main() {
	cmds := []commands.Commands{
		{FlagSet: commands.InitFlagSet(), Name: "init", Run: commands.Init},
		{FlagSet: commands.SwitchFlagSet(), Name: "switch", Run: commands.Switch},
		{FlagSet: commands.AddFlagSet(), Name: "add", Run: commands.Add},
		{FlagSet: commands.RemoveFlagSet(), Name: "remove", Run: commands.Remove},
		{FlagSet: commands.GetEnabledFlagSet(), Name: "get-enabled", Run: commands.GetEnabled},
		{FlagSet: commands.ListFlagSet(), Name: "list", Run: commands.List},
		{FlagSet: commands.TempFlagSet(), Name: "temp", Run: commands.Temp},
		{FlagSet: commands.GetFlagSet(), Name: "get", Run: commands.Get},
	}

	help := flag.Bool("help", false, "Shows the help menu")

	flag.Parse()

	if *help {
		printUsage(cmds)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		printUsage(cmds)
		os.Exit(1)
	}

	for _, cmd := range cmds {
		if os.Args[1] == cmd.Name {
			cmd.Run(cmd.FlagSet)
			return
		}
	}

	printUsage(cmds)
	os.Exit(1)
}
