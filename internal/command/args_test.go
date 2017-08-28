package command_test

import (
	"testing"

	"github.com/berfarah/knoch/internal/command"
	"github.com/stretchr/testify/assert"
)

func TestNewRuntime(t *testing.T) {
	assert := assert.New(t)

	noArgs := []string{"knoch"}

	cases := []struct {
		In  []string
		Out *command.Runtime
	}{
		{
			noArgs,
			&command.Runtime{Executable: "knoch", Command: "help", Args: []string{}},
		},
		{
			[]string{"knoch", "help"},
			&command.Runtime{Executable: "knoch", Command: "help", Args: []string{}},
		},
		{
			[]string{"knoch", "--help"},
			&command.Runtime{Executable: "knoch", Command: "help", Args: []string{}},
		},
		{
			[]string{"knoch", "foo", "--help", "baz"},
			&command.Runtime{Executable: "knoch", Command: "foo", Args: []string{"--help", "baz"}},
		},
		{
			[]string{"knoch", "foo", "--version", "baz"},
			&command.Runtime{Executable: "knoch", Command: "version", Args: []string{}},
		},
		{
			[]string{"knoch", "foo", "bar", "baz"},
			&command.Runtime{Executable: "knoch", Command: "foo", Args: []string{"bar", "baz"}},
		},
	}

	for _, c := range cases {
		args := command.NewRuntime(c.In)
		assert.Equal(
			args,
			c.Out,
		)
	}
}
