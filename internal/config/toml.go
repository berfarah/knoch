package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/berfarah/knoch/internal/config/project"
)

func Read() error {
	return Instance.Read()
}

func (c *Config) Read() error {
	b, err := ioutil.ReadFile(c.File)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(b, &c)
	for _, p := range c.Projects {
		project.Register(p)
	}
	return err
}

func Write() error {
	return Instance.Write()
}

func (c *Config) Write() error {
	c.Projects = c.Registry.Sorted()

	f, err := os.Create(c.File)
	defer f.Close()
	if err != nil {
		return err
	}

	err = toml.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}

	return err
}
