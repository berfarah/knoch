package config_test

import (
	"testing"

	"github.com/berfarah/knoch/internal/config"
	"github.com/stretchr/testify/assert"
)

var p = config.Project{
	Repo: "github.com/berfarah/dotfiles",
	Dir:  "berfarah/dotfiles",
}

func TestAddProject(t *testing.T) {
	assert := assert.New(t)
	set := config.Projects{}
	set.Add(p)

	_, ok := set[p.Repo]

	assert.True(ok, "should add the project")
}

func TestRemoveProject(t *testing.T) {
	assert := assert.New(t)
	set := config.Projects{}
	set.Add(p)
	set.Remove(p)

	_, ok := set[p.Repo]

	assert.False(ok, "should remove the project")
}
