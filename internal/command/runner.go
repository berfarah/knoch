package command

import (
	"os"

	"github.com/berfarah/knoch/internal/utils"
)

type Runner struct {
	Runtime  *Runtime
	Commands map[string]*Command
}

func NewRunner() *Runner {
	return &Runner{
		Runtime:  NewRuntime(os.Args),
		Commands: make(map[string]*Command),
	}
}

func (r *Runner) Execute() {
	err := r.Runtime.LoadConfig()
	r.failWithoutConfig(err)

	command := r.Command(r.Runtime.Command)
	if command != nil {
		command.Call(r.Runtime)
	} else {
		r.Command("help").Call(r.Runtime)
	}
}

func (r *Runner) failWithoutConfig(err error) {
	if err != nil && r.Runtime.Command != "init" {
		init := r.Command("init")
		utils.Exit(init.UsageText())
	}
}

func (r *Runner) Register(c *Command) {
	r.Commands[c.Name] = c
}

func (r Runner) Command(name string) *Command {
	return r.Commands[name]
}
