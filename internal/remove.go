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
	})
}

func runRemove(c *command.Command, r *command.Runtime) {
	var (
		proj project.Project
		err  error
	)

	if len(r.Args) == 0 {
		utils.Exit(c.UsageText())
	}

	proj, err = project.FromDir(r.Args[0])
	utils.Check(err, "")

	if !project.Remove(proj) {
		utils.Exit("Not tracking " + r.Args[0] + ", did nothing")
	}

	err = config.Write()

	err = os.RemoveAll(proj.Path())
	utils.Check(err, "")
}
