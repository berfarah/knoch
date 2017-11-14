package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type generalSettings struct {
	MaxWorkers int `toml:"parallel_workers"`
}

type encodableConfig struct {
	General  generalSettings `toml:"general"`
	Projects []Project       `toml:"project"`
}

func (c *Config) decode() {
	c.MaxWorkers = c.encoded.General.MaxWorkers
	for _, p := range c.encoded.Projects {
		c.Projects.Add(p)
	}
}

func (c *Config) refreshEncoded() {
	c.encoded = encodableConfig{
		General: generalSettings{
			MaxWorkers: c.MaxWorkers,
		},
		Projects: make([]Project, 0, len(c.Projects)),
	}

	for _, p := range c.Projects {
		c.encoded.Projects = append(c.encoded.Projects, p)
	}
}

func (c *Config) Read() error {
	b, err := ioutil.ReadFile(c.File())
	if err != nil {
		return err
	}

	err = toml.Unmarshal(b, &c.encoded)
	c.decode()
	return err
}

func (c *Config) Write() error {
	f, err := os.Create(c.File())
	defer f.Close()
	if err != nil {
		return err
	}

	c.refreshEncoded()
	err = toml.NewEncoder(f).Encode(c.encoded)
	if err != nil {
		return err
	}

	return err
}
