package config_test

import (
	"testing"

	"github.com/berfarah/knoch/config"
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

var expectedJSON = []byte(`[{"repo":"github.com/berfarah/dotfiles","dir":"berfarah/dotfiles"}]`)

func TestMarshalProjects(t *testing.T) {
	assert := assert.New(t)
	set := config.Projects{}
	set.Add(p)

	b, err := set.MarshalJSON()

	assert.Nil(err)
	assert.Equal(expectedJSON, b)
}

func TestUnmarshalProjects(t *testing.T) {
	assert := assert.New(t)
	set := config.Projects{}

	err := set.UnmarshalJSON(expectedJSON)
	_, ok := set[p.Repo]

	assert.Nil(err, "should not fail")
	assert.True(ok, "should have the project")
}
