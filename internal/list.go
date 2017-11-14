package internal

import (
	"fmt"
	"sort"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runList,

		Usage: "list",
		Name:  "list",
		Long:  "List tracked repositories",

		Aliases: []string{"ls"},
	})
}

const listWorkerCount = 4

type listGitDetail struct {
	Dir          string
	Branch       string
	LatestCommit string
}

type listMaxLengths struct {
	Dir    int
	Branch int
}

func (l *listMaxLengths) Add(dir, branch string) {
	if len(dir) > l.Dir {
		l.Dir = len(dir)
	}
	if len(branch) > l.Branch {
		l.Branch = len(branch)
	}
}

type listTable struct {
	maxLengths listMaxLengths
	order      []string
	results    map[string]listGitDetail
}

func newListTable() listTable {
	return listTable{
		maxLengths: listMaxLengths{},
		order:      []string{},
		results:    make(map[string]listGitDetail),
	}
}

func (l *listTable) RecordOrder(s string) {
	l.order = append(l.order, s)
}

func (l *listTable) Sort() {
	sort.Strings(l.order)
}

func (l *listTable) Print() {
	for _, dir := range l.order {
		details := l.results[dir]
		format := fmt.Sprintf("%%-%ds\t%%%ds\t%%s", l.maxLengths.Dir, l.maxLengths.Branch)
		str := fmt.Sprintf(format, dir, details.Branch, details.LatestCommit)
		utils.Println(str)
	}
}

func (l *listTable) RecordResult(d listGitDetail) {
	l.results[d.Dir] = d
	l.maxLengths.Add(d.Dir, d.Branch)
}

func runList(c *command.Command, r *command.Runtime) {
	count := len(r.Config.Projects)

	table := newListTable()
	done := make(chan listGitDetail, count)
	projectDirs := make(chan config.Project, count)

	for w := 0; w < listWorkerCount; w++ {
		go listWorker(projectDirs, done)
	}

	for _, project := range r.Config.Projects {
		table.RecordOrder(project.Dir) // the table should loop over projects to create this, but that would add another loop :P
		projectDirs <- project
	}

	table.Sort()
	for i := 0; i < count; i++ {
		table.RecordResult(<-done)
	}
	table.Print()
}

func listWorker(projectDirs <-chan config.Project, done chan<- listGitDetail) {
	for project := range projectDirs {
		g := git.New().InDir(project.Path())

		branch, err := g.Branch()
		if err != nil {
			utils.Errorln(err)
		}
		lastCommit, err := g.LastCommit()
		if err != nil {
			utils.Errorln(err)
		}

		done <- listGitDetail{
			Dir:          project.Dir,
			Branch:       branch,
			LatestCommit: lastCommit,
		}
	}
}
