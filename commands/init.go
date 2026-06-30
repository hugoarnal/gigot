package commands

import (
	"flag"
	"fmt"
	"os"
)

func InitFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("init", flag.ExitOnError)
}

func Init(cmd *flag.FlagSet) {
	cmd.Parse(os.Args[2:])

	fmt.Println("Hello world")
}
