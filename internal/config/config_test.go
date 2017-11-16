package config_test

import (
	"io/ioutil"
	"testing"

	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)

	p := project.Project{
		Repo: "github.com/berfarah/dotfiles",
		Dir:  "berfarah/dotfiles",
	}
	registry := project.Registry{}
	registry.Add(p)

	cfg := config.Config{
		File:     "../testdata/.knoch.read",
		Projects: []project.Project{},
		Registry: registry,
	}

	err := cfg.Read()
	assert.Nil(err)

	assert.Equal(8, cfg.General.MaxWorkers)
	assert.Equal(registry.Sorted(), cfg.Projects)
}

func TestWrite(t *testing.T) {
	assert := assert.New(t)

	p := project.Project{Repo: "foo", Dir: "bar"}
	expected := `[general]
  parallel_workers = 0

[[project]]
  dir = "bar"
  repo = "foo"
`

	registry := project.Registry{}
	registry.Add(p)
	cfg := config.Config{
		General:  &config.GeneralSettings{0},
		File:     "../testdata/.knoch.write",
		Projects: []project.Project{},
		Registry: registry,
	}

	err := cfg.Write()
	b, err := ioutil.ReadFile("../testdata/.knoch.write")
	assert.Nil(err)
	assert.Equal(expected, string(b))
}
