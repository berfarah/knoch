package git

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/berfarah/knoch/internal/utils"
)

func IsRepo(dir string) bool {
	return New().InDir(dir).WithArgs("rev-parse", "--git-dir").Success()
}

// RepoFromDir gets upstream repo from directory
// git output:
// origin\tgit@github.com:berfarah/knoch.git\t(fetch)
func RepoFromDir(dir string) (string, error) {
	if !utils.IsDir(dir) {
		return "", fmt.Errorf("Not a directory")
	}

	if !IsRepo(dir) {
		return "", fmt.Errorf("Not a repository")
	}

	out, err := New().InDir(dir).WithArgs("remote", "-v").Output()
	if err != nil {
		return "", err
	}

	for _, entry := range out {
		all := strings.Fields(entry)
		if len(all) != 3 || all[2] != "(fetch)" {
			continue
		}
		return all[1], nil
	}

	return "", fmt.Errorf("No remotes found")
}

// RepoFromString gets a full repository from a string
// example output:
// git clone git@github.com:berfarah/knoch.git
func RepoFromString(s string) string {
	if !strings.HasSuffix(binary, "hub") {
		return s
	}

	out, err := New().WithArgs("--noop", "clone", s).Output()
	if err != nil {
		return s
	}

	str := strings.Fields(out[0])
	return str[2]
}

var repositoryRegex = regexp.MustCompile("(?:[/:])((?:[^/]+/)?[^/]+?)(?:.git)?$")

// DirFromRepo builds a directory from a repository name
func DirFromRepo(repository string, args []string) string {
	if len(args) == 1 {
		matches := repositoryRegex.FindStringSubmatch(repository)
		if len(matches) > 1 {
			return matches[1]
		}
	}

	return args[1]
}

func Sync(dir string) error {
	if !IsRepo(dir) {
		return fmt.Errorf("fatal: not a repository")
	}

	cmd := New().InDir(dir)
	if strings.HasSuffix(binary, "hub") {
		return cmd.WithArgs("sync").Run()
	}
	return cmd.WithArgs("fetch").Run()
}
