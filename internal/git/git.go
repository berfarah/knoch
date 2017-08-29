package git

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/berfarah/knoch/internal/utils"
)

type Git struct {
	Path string
}

func New() Git {
	binary, err := exec.LookPath("hub")
	if err == nil {
		return Git{Path: binary}
	}

	binary, err = exec.LookPath("git")
	utils.Check(err, "Must have git in $PATH")

	return Git{Path: binary}
}

func (g Git) Exec(args ...string) error {
	some := append([]string{g.Path}, args...)
	return syscall.Exec(g.Path, some, os.Environ())
}

func (g Git) Output(args ...string) (out []string, err error) {
	b, err := exec.Command(g.Path, args...).CombinedOutput()
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			out = append(out, string(line))
		}
	}

	return out, err
}

func (g Git) Success(args ...string) bool {
	err := exec.Command(g.Path, args...).Run()
	return err == nil
}
