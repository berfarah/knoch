package config

import (
	"path"
	"path/filepath"
)

type Project struct {
	Repo string `toml:"repo"`
	Dir  string `toml:"dir"`
}

func (p Project) Clean() {
	var err error
	abs := p.Dir
	if !path.IsAbs(abs) {
		abs, err = filepath.Abs(abs)
		if err != nil {
			return
		}
	}

	p.Dir, err = filepath.Rel(workDir, abs)
	if err != nil {
		p.Dir = abs
	}
}

func (p Project) Path() string {
	return path.Join(workDir, p.Dir)
}

type Projects map[string]Project

func (set *Projects) Add(p Project) bool {
	p.Clean()
	_, found := (*set)[p.Dir]
	(*set)[p.Dir] = p
	return !found
}

func (set *Projects) Remove(p Project) {
	delete(*set, p.Dir)
}
