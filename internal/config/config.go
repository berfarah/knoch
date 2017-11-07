package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const Filename = ".knoch"
const defaultWorkers = 4

type Config struct {
	Filename  string   `yaml:"-"`
	Directory string   `yaml:"-"`
	Projects  Projects `yaml:"projects"`
	Workers   int      `yaml:"parallel_workers"`
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
	err = yaml.Unmarshal(b, ec)
	c.Decode(ec)
	return err
}

func (c *Config) Write() error {
	f, err := os.Create(c.File())
	defer f.Close()
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(c.Encode())
	if err != nil {
		return err
	}

	_, err = f.Write(b)

	return err
}
