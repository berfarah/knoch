package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

const Filename = ".wsrc"

type Config struct {
	File     string   `yaml:"-"`
	Projects Projects `yaml:"projects"`
}

func New() (*Config, error) {
	p, _ := os.Getwd()

	c := Config{
		File: path.Join(p, Filename),
	}
	c.Projects = Projects{}
	err := c.Read()
	return &c, err
}

func (c *Config) Read() error {
	b, err := ioutil.ReadFile(c.File)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, c)
	return err
}

func (c *Config) Write() error {
	f, err := os.Create(c.File)
	defer f.Close()
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(*c)
	if err != nil {
		return err
	}

	_, err = f.Write(b)

	return err
}
