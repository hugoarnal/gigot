package utils

import (
	"errors"
	"os"
)

func CheckPath(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}
