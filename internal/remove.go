package internal

import (
	"os"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runRemove,

		Usage: "remove <REPOSITORY/DIRECTORY>",
		Name:  "remove",
		Long:  "Remove and stop tracking repository",
	})
}

func runRemove(c *command.Command, r *command.Runtime) {
	if len(r.Args) == 0 {
		utils.Exit(c.UsageText())
	}

	var repo string
	var err error

	if utils.IsDir(r.Args[0]) {
		repo, err = git.RepoFromDir(r.Args[0])
		utils.Check(err, "")

	} else {
		repo = git.RepoFromString(r.Args[0])
	}

	project, found := r.Config.Projects[repo]
	if !found {
		utils.Exit("Not tracking " + r.Args[0] + ", did nothing")
	}

	dir := project.Path()

	r.Config.Projects.Remove(project)
	r.Config.Write()

	err = os.RemoveAll(dir)
	utils.Check(err, "")
}
