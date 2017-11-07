package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

const Filename = ".knoch"
const defaultWorkers = 4

type Config struct {
	Filename  string   `toml:"-"`
	Directory string   `toml:"-"`
	Projects  Projects `toml:"projects"`
	Workers   int      `toml:"parallel_workers"`
}

func New() (*Config, error) {
	c := Config{
		Filename:  Filename,
		Directory: ".",
		Projects:  Projects{},
		Workers:   defaultWorkers,
	}
	err := c.Read()
	return &c, err
}

func (c *Config) File() string {
	return path.Join(c.Directory, c.Filename)
}

func (c *Config) Read() error {
	ec := &EncodableConfig{}
	b, err := ioutil.ReadFile(c.File())
	if err != nil {
		return err
	}
	err = toml.Unmarshal(b, ec)
	c.Decode(ec)
	return err
}

func (c *Config) Write() error {
	f, err := os.Create(c.File())
	defer f.Close()
	if err != nil {
		return err
	}

	err = toml.NewEncoder(f).Encode(c.Encode())
	if err != nil {
		return err
	}

	return err
}
