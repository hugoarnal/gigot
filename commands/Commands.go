package commands

import "flag"

type Commands struct {
	FlagSet *flag.FlagSet
	Name    string
	Run     func(*flag.FlagSet)
}
