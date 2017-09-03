package utils

import "os"

func IsDir(dir string) bool {
	f, err := os.Stat(dir)
	return err == nil && f.IsDir()
}
