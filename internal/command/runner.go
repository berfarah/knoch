package command

import (
	"os"

	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/utils"
)

type Runner struct {
	Args     *Args
	Commands map[string]*Command
}

func NewRunner() *Runner {
	return &Runner{
		Args:     NewArgs(os.Args[1:]),
		Commands: make(map[string]*Command),
	}
}

func (r *Runner) Execute() {
	cfg, err := config.New()
	r.failWithoutConfig(err)

	command := r.Command(r.Args.Command)
	if command == nil {
		r.Command("help").Call(cfg, r.Args)
		return
	}
	command.Call(cfg, r.Args)
}

func (r *Runner) failWithoutConfig(err error) {
	if err != nil && r.Args.Command != "init" {
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
