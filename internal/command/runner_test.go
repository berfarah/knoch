package command_test

import (
	"testing"

	"github.com/berfarah/knoch/internal/command"
	"github.com/stretchr/testify/assert"
)

var (
	Runner = command.NewRunner()
	cmd    = &command.Command{Name: "foo"}
)

func TestRegister(t *testing.T) {
	Runner.Register(cmd)
	assert.Equal(t, cmd, Runner.Commands["foo"])
}

func TestCommand(t *testing.T) {
	assert.Equal(t, cmd, Runner.Command("foo"))
}
