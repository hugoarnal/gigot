package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func TempFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("temp", flag.ExitOnError)

	return cmd
}

func Temp(cmd *flag.FlagSet) {
	if err := cmd.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	config.GigotShellWarning()

	filename := config.GetGitConfigFilename()

	parsedConfig, err := config.ParseGitConfigFile(filename)

	if err != nil {
		panic(err)
	}

	if cmd.NArg() == 0 {
		SwitchList(parsedConfig)
	} else if cmd.NArg() == 1 {
		selectedConfig := cmd.Arg(0)

		_, err := config.FindGitConfigByName(parsedConfig, selectedConfig)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Switching for this shell session to \"%s\"\n", selectedConfig)
	} else {
		fmt.Println("Incorrect amount of arguments")
		os.Exit(1)
	}
}
