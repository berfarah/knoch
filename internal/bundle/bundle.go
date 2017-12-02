package bundle

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/config/project"
	"github.com/berfarah/knoch/internal/git"
	"github.com/berfarah/knoch/internal/utils"
)

const threads = 16
const timeout = 60 // seconds

type job struct {
	Index   int
	Project project.Project
}

type status struct {
	Index   int
	Project string
	New     bool
	Stdout  string
	Error   error
}

var wg sync.WaitGroup

func Run(c *command.Command, r *command.Runtime) {
	count := len(project.Tracker)
	jobs := make(chan job, count)
	statuses := make(chan status, threads*3)
	done := make(chan bool)

	defer func() {
		wg.Wait()
		close(statuses)
		<-done
	}()

	for index, project := range project.Tracker.Sorted() {
		statuses <- status{Index: index, Project: project.Dir}
		jobs <- job{index, project}
	}
	close(jobs)

	go print(count, statuses, done)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(i, jobs, statuses)
	}
}

func perform(job job, statuses chan<- status) {
	project := job.Project
	s := status{
		Index:   job.Index,
		Project: job.Project.Dir,
	}
	defer func() {
		statuses <- s
	}()
	var cmd exec.Cmd

	if project.Exists() {
		cmd = *git.New().WithArgs("fetch", "--progress").InDir(project.Path()).Cmd()
		s.Stdout = "Updating..."
		statuses <- s
	} else {
		cmd = *git.New().WithArgs("clone", "--progress", project.Repo, project.Path()).Cmd()
		s.Stdout = "Downloading..."
		statuses <- s
	}

	cmdReader, err := cmd.StderrPipe()
	if err != nil {
		s.Error = err
		return
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			s.Stdout = strings.Trim(scanner.Text(), "\r")
			statuses <- s
		}
	}()

	if err := cmd.Start(); err != nil {
		s.Error = err
		statuses <- s
		return
	}

	timer := time.AfterFunc(60*time.Second, func() {
		cmd.Process.Kill()
	})
	err = cmd.Wait()
	s.Error = err
	s.Stdout = "Done!"
	timer.Stop()
}

func worker(id int, jobs <-chan job, statuses chan<- status) {
	for job := range jobs {
		perform(job, statuses)
	}
	wg.Done()
}

func (s status) String() string {
	if s.Error != nil {
		return fmt.Sprintf("\x1B[0;31m%s\x1B[0m", s.Error.Error())
	}

	return s.Stdout
}

type statusTable struct {
	runOnce      bool
	maxDirLength int
	writer       bufio.Writer
	dirty        bool
	table        []status
}

func newStatusTable(count int) statusTable {
	return statusTable{
		table: make([]status, count, count),
	}
}

func (st *statusTable) updateMax(length int) {
	if length > st.maxDirLength {
		st.maxDirLength = length
	}
}

func (st *statusTable) Add(s status) {
	st.updateMax(len(s.Project))
	st.table[s.Index] = s
	st.dirty = true
}

func (st *statusTable) hasRun() bool {
	if !st.runOnce {
		st.runOnce = true
		return false
	}
	return true
}

func (st *statusTable) Print() {
	var output string
	if !st.dirty {
		return
	}

	// copy table so that our max dir length doesn't get outdated
	table := append([]status{}, st.table...)
	format := fmt.Sprintf("%%-%ds\t%%-30s\n", st.maxDirLength)

	st.dirty = false
	if st.hasRun() {
		output = strings.Repeat("\x1B[A\x1B[2K\r", len(table))
	}

	for _, s := range table {
		output += fmt.Sprintf(format, s.Project, s.String())
	}

	utils.Print(output)
}

func print(count int, statuses <-chan status, done chan<- bool) {
	table := newStatusTable(count)
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()

	go func() {
		table.Print()
		for _ = range ticker.C {
			table.Print()
		}
	}()

	for status := range statuses {
		table.Add(status)
	}

	table.Print()
	done <- true
}
