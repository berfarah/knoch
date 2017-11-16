package config

import (
	"github.com/berfarah/knoch/internal/config/finder"
	"github.com/berfarah/knoch/internal/config/project"
)

const defaultWorkers = 4

var Instance = Config{
	File:     finder.File(),
	WorkDir:  finder.WorkDir,
	Registry: project.Tracker,
	Projects: []project.Project{},
	General: generalSettings{
		MaxWorkers: defaultWorkers,
	},
}

type generalSettings struct {
	MaxWorkers int `toml:"parallel_workers"`
}

type Config struct {
	File    string `toml:"-"`
	WorkDir string `toml:"-"`

	Registry project.Registry `toml:"-"`

	General  generalSettings   `toml:"general"`
	Projects []project.Project `toml:"project"`
}

func (c Config) Workers() int {
	if c.General.MaxWorkers < 1 {
		return 1
	}

	if len(c.Projects) < c.General.MaxWorkers {
		return len(c.Projects)
	}

	return c.General.MaxWorkers
}
