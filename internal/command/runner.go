package command

import (
	"os"
	"sort"

	"github.com/berfarah/knoch/internal/config"
)

type Runner struct {
	runtime  *Runtime
	Commands map[string]*Command
	Aliases  map[string]string

	defaultCommand string
	helpCommand    string
}

func NewRunner() *Runner {
	return &Runner{
		runtime:  NewRuntime(os.Args),
		Commands: make(map[string]*Command),
		Aliases:  make(map[string]string),
	}
}

func (r *Runner) SortedCommands() []Command {
	var keys []string
	for key := range r.Commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var commands []Command
	for _, key := range keys {
		commands = append(commands, *r.Command(key))
	}

	return commands
}

func (r *Runner) Execute() {
	err := config.Read()
	if err != nil && r.runtime.Command == "" {
		r.HelpCommand().Call(r.runtime)
		os.Exit(1)
	}

	if r.runtime.Command == "" {
		r.DefaultCommand().Call(r.runtime)
		return
	}

	command := r.Command(r.runtime.Command)
	if command == nil {
		r.HelpCommand().Call(r.runtime)
		os.Exit(1)
	}

	command.Call(r.runtime)
}

func (r *Runner) Register(c *Command) {
	r.Commands[c.Name] = c
	r.Alias(c.Name, c.Aliases...)
}

func (r *Runner) Alias(from string, to ...string) {
	for _, alias := range to {
		r.Aliases[alias] = from
	}
}

func (r *Runner) RegisterDefault(c *Command) {
	r.Register(c)
	r.defaultCommand = c.Name
}

func (r *Runner) RegisterHelp(c *Command) {
	r.Register(c)
	r.Commands["help"] = c
}

func (r Runner) Command(name string) *Command {
	command, ok := r.Commands[name]
	if !ok {
		command = r.Commands[r.Aliases[name]]
	}

	return command
}

func (r *Runner) DefaultCommand() *Command {
	return r.Commands[r.defaultCommand]
}

func (r *Runner) HelpCommand() *Command {
	return r.Commands["help"]
}
