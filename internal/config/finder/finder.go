package finder

import (
	"os"
	"path"
	"path/filepath"
)

const Filename = ".knoch"
const defaultDir = "."

var WorkDir string

func File() string {
	return path.Join(WorkDir, Filename)
}

func init() {
	findConfig()
}

func findConfig() {
	var err error

	WorkDir, err = os.Getwd()

	if err != nil {
		WorkDir = defaultDir
		return
	}

	for {
		if doesFileExist(File()) {
			return
		}

		if isHome(WorkDir) || isRoot(WorkDir) {
			WorkDir = defaultDir
			return
		}

		WorkDir = filepath.Dir(WorkDir)
	}
}

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func isHome(path string) bool {
	return path == os.Getenv("HOME")
}

func isRoot(path string) bool {
	return path == "/"
}
