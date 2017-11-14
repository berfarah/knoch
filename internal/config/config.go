package config

import (
	"path"
)

const Filename = ".knoch"
const defaultWorkers = 4

type Config struct {
	Filename   string   `toml:"-"`
	Directory  string   `toml:"-"`
	Projects   Projects `toml:"projects"`
	MaxWorkers int      `toml:"parallel_workers"`

	encoded encodableConfig
}

func (c Config) Workers() int {
	if len(c.Projects) < c.MaxWorkers {
		return len(c.Projects)
	}

	return c.MaxWorkers
}

func New() (*Config, error) {
	c := Config{
		Filename:  Filename,
		Directory: ".",
		Projects:  Projects{},

		encoded: encodableConfig{},
	}
	err := c.Read()
	if c.MaxWorkers == 0 {
		c.MaxWorkers = defaultWorkers
	}

	return &c, err
}

func (c *Config) File() string {
	return path.Join(c.Directory, c.Filename)
}
