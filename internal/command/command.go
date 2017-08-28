package command

import (
	"fmt"

	flag "github.com/ogier/pflag"

	"github.com/berfarah/knoch/internal/config"
	"github.com/berfarah/knoch/internal/utils"
)

type Command struct {
	Run  func(c *Command, cfg *config.Config, args *Args)
	Flag flag.FlagSet

	Name  string
	Usage string
	Long  string
}

func (c *Command) Call(cfg *config.Config, args *Args) (err error) {
	if c.Run != nil {
		c.Run(c, cfg, args)
		return nil
	}

	return nil
}

func (c *Command) parseArgs(args *Args) (err error) {
	c.Flag.SetInterspersed(true)
	c.Flag.Init(c.Name, flag.ContinueOnError)
	c.Flag.Usage = func() {
		utils.Errorln("")
		utils.Errorln(c.UsageText())
	}

	return nil
}

func (c *Command) UsageText() string {
	return fmt.Sprintf("Usage: ws %s", c.Usage)
}
