package internal

import (
	"fmt"
	"sort"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

var (
	flagListSimple bool
)

func init() {
	Runner.Register(&command.Command{
		Run: runList,

		Usage: "list",
		Name:  "list",
		Long:  "List tracked repositories",

		Aliases: []string{"ls"},
	})
	cmd := Runner.Command("list")
	cmd.Flag.BoolVar(&flagListSimple, "name-only", false, "list only repo names")
}

const listWorkerCount = 4

type listGitDetail struct {
	Dir          string
	Branch       string
	LatestCommit string
	Error        error
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

		if details.Error != nil {
			format := fmt.Sprintf("%%-%ds\t%%s", l.maxLengths.Dir)
			str := fmt.Sprintf(format, dir, details.Error)
			utils.Println(str)
			continue
		}

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
	if flagListSimple {
		runSimpleList(c, r)
	} else {
		runFullList(c, r)
	}
}

func runSimpleList(c *command.Command, r *command.Runtime) {
	for _, p := range project.Tracker.Sorted() {
		utils.Println(p.Dir)
	}
}

func runFullList(c *command.Command, r *command.Runtime) {
	count := len(project.Tracker)

	table := newListTable()
	done := make(chan listGitDetail, count)
	projDirs := make(chan project.Project, count)

	for w := 0; w < listWorkerCount; w++ {
		go listWorker(projDirs, done)
	}

	for _, proj := range project.Tracker {
		table.RecordOrder(proj.Dir) // the table should loop over projs to create this, but that would add another loop :P
		projDirs <- proj
	}

	table.Sort()
	for i := 0; i < count; i++ {
		table.RecordResult(<-done)
	}
	table.Print()
}

func listWorker(projDirs <-chan project.Project, done chan<- listGitDetail) {
	for proj := range projDirs {
		g := git.New().InDir(proj.Path())

		branch, err := g.Branch()
		lastCommit, err := g.LastCommit()

		done <- listGitDetail{
			Dir:          proj.Dir,
			Branch:       branch,
			LatestCommit: lastCommit,
			Error:        err,
		}
	}
}
