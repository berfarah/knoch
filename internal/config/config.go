package config

import (
	"github.com/berfarah/knoch/internal/config/finder"
	"github.com/berfarah/knoch/internal/config/project"
)

const defaultWorkers = 4

var general = GeneralSettings{
	MaxWorkers: defaultWorkers,
}

var instance = Config{
	File:     finder.File(),
	WorkDir:  finder.WorkDir,
	Registry: project.Tracker,
	Projects: []project.Project{},
	General:  &general,
}

type GeneralSettings struct {
	MaxWorkers int `toml:"parallel_workers"`
}

type Config struct {
	File    string `toml:"-"`
	WorkDir string `toml:"-"`

	Registry project.Registry `toml:"-"`

	General  *GeneralSettings  `toml:"general"`
	Projects []project.Project `toml:"project"`
}

func Workers() int {
	if general.MaxWorkers < 1 {
		return 1
	}

	if len(project.Tracker) < general.MaxWorkers {
		return len(project.Tracker)
	}

	return general.MaxWorkers
}
