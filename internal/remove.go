package internal

import (
	"os"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runRemove,

		Usage: "remove <DIRECTORY>",
		Name:  "remove",
		Long:  "Remove and stop tracking repository",

		Aliases: []string{"rm"},
	})
}

func runRemove(c *command.Command, r *command.Runtime) {
	var (
		proj project.Project
		err  error
		dir  string
	)

	if len(r.Args) == 0 {
		utils.Exit(c.UsageText())
	}

	dir = r.Args[0]

	if utils.IsDir(dir) {
		proj, err = project.FromDir(dir)
		utils.Check(err, "")
	} else {
		proj, _ = project.Fetch(dir)
	}

	if !project.Remove(proj) {
		utils.Exit("Not tracking " + dir + ", did nothing")
	}

	err = config.Write()

	err = os.RemoveAll(proj.Path())
	utils.Check(err, "")
}
