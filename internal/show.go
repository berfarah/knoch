package internal

import (
	"errors"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/utils"
)

var (
	flagPrintProjectPath bool
)

func init() {
	Runner.Register(&command.Command{
		Run: runShow,

		Usage: "show",
		Name:  "show",
		Long:  "Show full path of selected project",

		Aliases: []string{"s"},
	})
}

func runShow(c *command.Command, r *command.Runtime) {
	if len(r.Args) < 1 {
		utils.Exit("No project provided")
	}

	project, err := findProject(r, r.Args[0])
	if err != nil {
		utils.Exit("No project by that name exists")
	}

	utils.Println(project.Path())
}

func findProject(r *command.Runtime, dir string) (config.Project, error) {
	var p config.Project

	for _, project := range r.Config.Projects {
		if project.Dir == dir {
			return project, nil
		}
	}

	return p, errors.New("No project found")
}
