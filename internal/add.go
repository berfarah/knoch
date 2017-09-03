package internal

import (
	"os"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config"
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
	if len(r.Args) == 0 {
		utils.Exit(c.UsageText())
	}

	if addDirProject(r) {
		return
	}

	addRepoProject(r)
}

func addRepoProject(r *command.Runtime) {
	repository := git.RepoFromString(r.Args[0])
	directory := git.DirFromRepo(repository, r.Args)

	r.Config.Projects.Add(config.Project{
		Repo: repository,
		Dir:  directory,
	})

	if utils.IsDir(directory) {
		return
	}

	err := os.MkdirAll(directory, 0755)
	utils.Check(err, "")

	err = r.Config.Write()
	utils.Check(err, "")

	err = git.Exec("clone", repository, directory)
	utils.Check(err, "")
}

func addDirProject(r *command.Runtime) bool {
	dir := r.Args[0]

	if !utils.IsDir(dir) {
		return false
	}

	if !git.IsRepo(dir) {
		return false
	}

	repo, err := git.RepoFromDir(dir)
	if err != nil {
		return false
	}

	added := r.Config.Projects.Add(config.Project{
		Repo: repo,
		Dir:  dir,
	})
	if !added {
		return true
	}

	err = r.Config.Write()
	utils.Check(err, "")

	return true
}
