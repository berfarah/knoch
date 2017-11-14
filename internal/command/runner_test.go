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
	assert.Equal(t, cmd, Runner.Command("foo"))
}

func TestRegisterHelp(t *testing.T) {
	Runner.RegisterHelp(cmd)
	assert.Equal(t, cmd, Runner.HelpCommand())
}

func TestRegisterDefault(t *testing.T) {
	Runner.RegisterDefault(cmd)
	assert.Equal(t, cmd, Runner.DefaultCommand())
}
