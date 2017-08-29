package git

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/berfarah/knoch/internal/utils"
)

func (git Git) RepoFromDir(dir string) (string, error) {
	remotes := make(map[string]string)
	// output:
	// origin\tgit@github.com:berfarah/knoch.git\t(fetch)
	out, err := git.Output("remote", "-v")
	if err != nil {
		return "", err
	}

	for _, entry := range out {
		all := strings.Fields(entry)
		if all[2] == "(push)" {
			continue
		}
		remotes[all[0]] = all[1]
	}

	if repo, ok := remotes["origin"]; ok {
		return repo, nil
	}

	if len(remotes) == 0 {
		utils.Exit("Repository must have a remote set up")
	}

	var repoText bytes.Buffer
	repoText.WriteString("Which repo is the origin?\n")
	for _, key := range remotes {
		line := fmt.Sprintf("  [%s] %s\n", key, remotes[key])
		repoText.WriteString(line)
	}
	fmt.Println(repoText)

	reader := bufio.NewReader(os.Stdin)
	selection, _, _ := reader.ReadLine()
	return remotes[strings.TrimSpace(string(selection))], nil
}

func (git Git) RepoFromString(arg string) string {
	if strings.HasSuffix(git.Path, "hub") {
		out, err := git.Output("--noop", "clone", arg)
		if err != nil {
			return arg
		}

		// example output:
		// git clone git@github.com:berfarah/knoch.git
		str := strings.Fields(out[0])
		return str[2]
	}

	return arg
}

var repositoryRegex = regexp.MustCompile("(?:[/:])((?:[^/]+/)?[^/]+?)(?:.git)?$")

func (git Git) DirFromRepo(repository string, args []string) string {
	if len(args) == 1 {
		matches := repositoryRegex.FindStringSubmatch(repository)
		if len(matches) > 1 {
			return matches[1]
		}
	}

	return args[1]
}
