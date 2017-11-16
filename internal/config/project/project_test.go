package project_test

import (
	"testing"

	project "github.com/berfarah/knoch/internal/config/project"
	"github.com/stretchr/testify/assert"
)

var p = project.Project{
	Repo: "github.com/berfarah/dotfiles",
	Dir:  "berfarah/dotfiles",
}

func TestAddProject(t *testing.T) {
	assert := assert.New(t)
	r := project.Registry{}
	r.Add(p)

	_, ok := r[p.Dir]

	assert.True(ok, "should add the project")
}

func TestRemoveProject(t *testing.T) {
	assert := assert.New(t)
	r := project.Registry{}
	r.Add(p)
	r.Remove(p)

	_, ok := r[p.Dir]

	assert.False(ok, "should remove the project")
}
