package config

import (
	"os"
	"path"
	"path/filepath"
)

const Filename = ".knoch"
const defaultDir = "."
const defaultWorkers = 4

var workDir string

type Config struct {
	Filename string `toml:"-"`
	WorkDir  string `toml:"-"`

	Projects   Projects `toml:"projects"`
	MaxWorkers int      `toml:"parallel_workers"`

	encoded encodableConfig
}

func (c Config) Workers() int {
	if c.MaxWorkers < 1 {
		return 1
	}

	if len(c.Projects) < c.MaxWorkers {
		return len(c.Projects)
	}

	return c.MaxWorkers
}

func New() (*Config, error) {
	c := Config{
		Filename:   Filename,
		Projects:   Projects{},
		MaxWorkers: defaultWorkers,

		encoded: encodableConfig{},
	}
	c.findConfig()
	workDir = c.WorkDir
	err := c.Read()

	return &c, err
}

func (c *Config) findConfig() {
	var err error

	c.WorkDir, err = os.Getwd()

	if err != nil {
		c.WorkDir = defaultDir
		return
	}

	for {
		if doesFileExist(c.File()) {
			return
		}

		if isHome(c.WorkDir) || isRoot(c.WorkDir) {
			c.WorkDir = defaultDir
			return
		}

		c.WorkDir = filepath.Dir(c.WorkDir)
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

func (c *Config) File() string {
	return path.Join(c.WorkDir, c.Filename)
}
