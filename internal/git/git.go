package git

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/berfarah/knoch/internal/utils"
)

var binary string

func init() {
	bin, err := exec.LookPath("hub")
	if err == nil {
		binary = bin
		return
	}
	bin, err = exec.LookPath("git")
	utils.Check(err, "Must have git in $PATH")
	binary = bin
}

func Exec(args ...string) error {
	return syscall.Exec(
		binary,
		append([]string{binary}, args...),
		os.Environ(),
	)
}

type Command struct {
	Path string
	Args []string
	Dir  string
}

func New() *Command {
	return &Command{
		Path: binary,
		Args: []string{},
		Dir:  "",
	}
}

func (c *Command) WithArgs(args ...string) *Command {
	c.Args = append(c.Args, args...)
	return c
}

func (c *Command) InDir(dir string) *Command {
	c.Dir = dir
	return c
}

func (c *Command) Cmd() *exec.Cmd {
	cmd := exec.Command(c.Path, c.Args...)
	cmd.Dir = c.Dir
	return cmd
}

func (c *Command) Output() ([]string, error) {
	b, err := c.Cmd().Output()
	return formatOutput(b), err
}

func (c *Command) Run() error {
	return c.Cmd().Run()
}

func (c *Command) Success() bool {
	return c.Run() == nil
}

func formatOutput(b []byte) (out []string) {
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			out = append(out, string(line))
		}
	}
	return out
}
