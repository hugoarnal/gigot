package config

import (
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
