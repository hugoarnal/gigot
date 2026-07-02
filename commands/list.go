package commands

import (
	"flag"
	"fmt"
	"gigot/internal/config"
	"os"
)

func ListFlagSet() *flag.FlagSet {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)

	return cmd
}

func List(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	parsedConfig, err := config.ParseGitConfigFile(config.GetGitConfigFilename())

	if err != nil {
		panic(err)
	}

	for _, c := range parsedConfig {
		fmt.Printf("%s | %s | %t\n", c.Name, c.Path, c.Enabled)
	}
}
