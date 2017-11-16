package internal

import (
	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runShow,

		Usage: "show",
		Name:  "show",
		Long:  "Show full path of selected project",
	})
}

func runShow(c *command.Command, r *command.Runtime) {
	if len(r.Args) < 1 {
		utils.Exit("No project provided")
	}

	proj, ok := project.Fetch(r.Args[0])
	if !ok {
		utils.Exit("No project by that name exists")
	}

	utils.Println(proj.Path())
}
