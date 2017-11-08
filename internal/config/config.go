package config

import (
	"path"
)

const Filename = ".knoch"
const defaultWorkers = 4

type Config struct {
	Filename  string   `toml:"-"`
	Directory string   `toml:"-"`
	Projects  Projects `toml:"projects"`
	Workers   int      `toml:"parallel_workers"`

	encoded encodableConfig
}

func New() (*Config, error) {
	c := Config{
		Filename:  Filename,
		Directory: ".",
		Projects:  Projects{},
		Workers:   defaultWorkers,

		encoded: encodableConfig{},
	}
	err := c.Read()
	return &c, err
}

func (c *Config) File() string {
	return path.Join(c.Directory, c.Filename)
}
