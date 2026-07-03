package config

import (
	"fmt"
	utils "gigot/internal"
	"os"

	"github.com/adrg/xdg"
)

func GetPath() string {
	return xdg.ConfigHome + "/gigot"
}

func CheckPath() bool {
	return utils.CheckPath(GetPath())
}

func CreatePath() error {
	err := os.MkdirAll(GetPath(), 0700)

	return err
}

func IsGigotShellSet() bool {
	_, present := os.LookupEnv("GIGOT_SHELL")

	return present
}

func GigotShellWarning() {
	if !IsGigotShellSet() {
		fmt.Println("WARNING: the \"GIGOT_SHELL\" environment variable isn't set")
		fmt.Println("That means the shell session probably wasn't initialized using \"gigot init\"")
		fmt.Println("Please look at the documentation to setup your shell environment properly")
		fmt.Println("Proceeding with the switch command")
		fmt.Println()
	}
}
