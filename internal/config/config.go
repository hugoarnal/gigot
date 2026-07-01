package config

import (
	"github.com/adrg/xdg"
)

func GetPath() string {
	return xdg.ConfigHome + "/gigot"
}
