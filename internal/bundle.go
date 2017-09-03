package internal

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runBundle,

		Usage: "bundle",
		Name:  "bundle",
		Long:  "Download all repositories",
	})
}

func runBundle(c *command.Command, r *command.Runtime) {
	count := len(r.Config.Projects)
	done := make(chan bool, 1)
	projects := make(chan config.Project, count)
	results := make(chan status, count)
	progress := newProgress(count)

	go progress.Track(results, done)
	for w := 0; w < r.Config.Workers; w++ {
		go worker(w, projects, results)
	}
	for _, project := range r.Config.Projects {
		projects <- project
	}
	close(projects)

	<-done
}

func worker(id int, projects <-chan config.Project, results chan<- status) {
	for project := range projects {
		if utils.IsDir(project.Dir) {
			err := git.Sync(project.Dir)
			results <- status{Repo: project.Repo, Sync: true, Error: err}
		} else {
			err := git.New().WithArgs("clone", "--quiet", project.Repo, project.Dir).Run()
			results <- status{Repo: project.Repo, Download: true, Error: err}
		}
	}
}

type status struct {
	Sync     bool
	Download bool
	Repo     string
	Error    error
}

func (s status) Success() bool {
	return s.Error == nil
}

type progress struct {
	current  int
	sync     int
	download int
	failed   int
	total    int
}

func newProgress(total int) *progress {
	return &progress{
		current:  0,
		sync:     0,
		download: 0,
		failed:   0,
		total:    total,
	}
}

func (p *progress) Track(results <-chan status, done chan<- bool) {
	p.report()
	errors := []status{}

	for s := 0; s < p.total; s++ {
		status := <-results
		p.current++

		if status.Error != nil {
			p.failed++
			errors = append(errors, status)
		} else {
			if status.Download {
				p.download++
			}

			if status.Sync {
				p.sync++
			}
		}

		p.report()
	}

	utils.Println("")

	if len(errors) > 0 {
		utils.Errorln("\nErrors:")
	}

	for _, status := range errors {
		if serr, ok := status.Error.(*exec.ExitError); ok {
			stderr := strings.Split(string(serr.Stderr), "\n")
			if len(stderr) < 1 {
				utils.Errorln(status.Repo, serr)
				continue
			}
			for _, line := range stderr {
				if strings.HasPrefix(line, "fatal:") {
					utils.Errorln(status.Repo, "-", line)
				}
			}
		} else {
			utils.Errorln(status.Repo, status.Error)
		}
	}

	done <- true
}

func (p progress) report() {
	text := fmt.Sprintf(
		"\r[%v/%v] %v clone %v sync %v error",
		p.current,
		p.total,
		p.download,
		p.sync,
		p.failed,
	)
	utils.Print(text)
}
