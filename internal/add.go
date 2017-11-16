package internal

import (
	"os"
	"path/filepath"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runAdd,

		Usage: "add <REPOSITORY> [<DIRECTORY>]",
		Name:  "add",
		Long:  "Clone and track a repository locally",
	})
}

func runAdd(c *command.Command, r *command.Runtime) {
	var (
		proj project.Project
		err  error
	)

	if len(r.Args) == 0 {
		utils.Exit(c.UsageText())
	}

	proj, err = project.FromDir(r.Args[0])
	if err != nil {
		proj, err = project.FromRepo(r.Args[0])

		if len(r.Args) > 1 {
			proj.Dir = r.Args[1]
		}
	}

	utils.Check(err, "")

	project.Register(proj)
	err = config.Write()
	utils.Check(err, "")

	if !proj.Exists() {
		cloneProject(proj)
	}
}

func cloneProject(proj project.Project) {
	var err error

	parentDir := filepath.Dir(proj.Path())
	if !utils.IsDir(parentDir) {
		err = os.MkdirAll(parentDir, 0755)
		utils.Check(err, "")
	}

	err = git.Exec("clone", proj.Repo, proj.Path())
	utils.Check(err, "")
}
