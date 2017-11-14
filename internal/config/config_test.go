package config_test

import (
	"io/ioutil"
	"testing"

	"github.com/berfarah/knoch/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)

	p := config.Project{
		Repo: "github.com/berfarah/dotfiles",
		Dir:  "berfarah/dotfiles",
	}
	set := config.Projects{}
	set.Add(p)

	cfg := config.Config{
		WorkDir:  "../testdata",
		Filename: ".knoch",
		Projects: config.Projects{},
	}

	err := cfg.Read()
	assert.Nil(err)
	assert.Equal(set, cfg.Projects)
}

func TestWrite(t *testing.T) {
	assert := assert.New(t)

	p := config.Project{Repo: "foo", Dir: "bar"}
	expected := `[general]
  parallel_workers = 0

[[project]]
  repo = "foo"
  dir = "bar"
`

	cfg := config.Config{
		WorkDir:  "../testdata",
		Filename: ".knoch.test",
		Projects: config.Projects{},
	}
	cfg.Projects.Add(p)

	err := cfg.Write()
	b, err := ioutil.ReadFile("../testdata/.knoch.test")
	assert.Nil(err)
	assert.Equal(expected, string(b))
}
