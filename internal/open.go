package internal

import (
	"os/exec"
	"syscall"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runOpen,

		Usage: "open",
		Name:  "open",
		Long:  "Open the selected project in $EDITOR",
	})
}

const defaultEditor = "vim"

func runOpen(c *command.Command, r *command.Runtime) {
	var editor string

	if len(r.Args) < 1 {
		utils.Exit("No project provided")
	}

	proj, ok := project.Fetch(r.Args[0])
	if !ok {
		utils.Exit("No project by that name exists")
	}

	editor, _ = syscall.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}

	binary, err := exec.LookPath(editor)
	if err != nil {
		utils.Exit("$EDITOR not found in path")
	}

	syscall.Exec(binary, []string{binary, proj.Path()}, syscall.Environ())
}
