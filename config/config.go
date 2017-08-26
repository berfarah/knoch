package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

const Filename = ".wsrc"

type Config struct {
	File     string   `json:"-"`
	Projects Projects `json:"projects"`
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
	err = json.Unmarshal(b, c)
	return err
}

func (c *Config) Write() error {
	f, err := os.Create(c.File)
	defer f.Close()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(*c, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(b)

	return err
}
