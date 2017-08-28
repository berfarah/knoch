package command_test

import (
	"testing"

	"github.com/berfarah/knoch/internal/command"
	"github.com/stretchr/testify/assert"
)

func TestNewArgs(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		In  []string
		Out *command.Args
	}{
		{
			[]string{},
			&command.Args{Command: "help", Params: []string{}},
		},
		{
			[]string{"help"},
			&command.Args{Command: "help", Params: []string{}},
		},
		{
			[]string{"foo", "--help", "baz"},
			&command.Args{Command: "help", Params: []string{}},
		},
		{
			[]string{"foo", "--version", "baz"},
			&command.Args{Command: "version", Params: []string{}},
		},
		{
			[]string{"foo", "bar", "baz"},
			&command.Args{Command: "foo", Params: []string{"bar", "baz"}},
		},
	}

	for _, c := range cases {
		args := command.NewArgs(c.In)
		assert.Equal(
			args,
			c.Out,
		)
	}
}
