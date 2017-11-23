package project

import (
	"path"
	"path/filepath"
	"regexp"

	"github.com/berfarah/knoch/internal/config/finder"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

type Project struct {
	Dir  string `toml:"dir"`
	Repo string `toml:"repo"`
}

var repositoryRegex = regexp.MustCompile("(?:[/:])((?:[^/]+/)?[^/]+?)(?:.git)?$")

func FromRepo(repo string) (Project, error) {
	var dir string
	repo = git.RepoFromString(repo)

	matches := repositoryRegex.FindStringSubmatch(repo)
	if len(matches) > 1 {
		dir = matches[1]
	}

	return New(dir, repo), nil
}

func FromDir(dir string) (Project, error) {
	dir = cleanPath(dir)
	project := New(dir, "")

	repo, err := git.RepoFromDir(project.Path())
	project.Repo = repo

	return project, err
}

func New(dir, repo string) Project {
	return Project{Dir: dir, Repo: repo}
}

func cleanPath(dir string) string {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return dir
	}

	rel, err := filepath.Rel(finder.WorkDir, abs)
	if err != nil {
		return dir
	}

	return rel
}

func (p Project) Exists() bool {
	return utils.IsDir(p.Path()) && git.IsRepo(p.Path())
}

func (p Project) Path() string {
	return path.Join(finder.WorkDir, p.Dir)
}
