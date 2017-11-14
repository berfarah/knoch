package command

import (
	"fmt"

	flag "github.com/ogier/pflag"

	"github.com/berfarah/knoch/internal/utils"
)

const binaryName = "knoch"

type Command struct {
	Run  func(c *Command, r *Runtime)
	Flag flag.FlagSet

	Name  string
	Usage string
	Long  string

	Aliases []string
}

func (c *Command) Call(r *Runtime) (err error) {
	if c.Run != nil {
		c.Run(c, r)
		return nil
	}

	return nil
}

func (c *Command) parseArgs(r *Runtime) (err error) {
	c.Flag.SetInterspersed(true)
	c.Flag.Init(c.Name, flag.ContinueOnError)
	c.Flag.Usage = func() {
		utils.Errorln("")
		utils.Errorln(c.UsageText())
	}

	return nil
}

func (c *Command) UsageText() string {
	return fmt.Sprintf("Usage: %s %s", binaryName, c.Usage)
}
