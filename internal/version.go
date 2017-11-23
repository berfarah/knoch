package internal

import (
	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/utils"
)

// VERSION of knoch
const VERSION = "0.3.3"

func init() {
	Runner.Register(&command.Command{
		Run:  runVersion,
		Name: "version",
	})
}

func runVersion(c *command.Command, r *command.Runtime) {
	utils.Println("knoch version " + VERSION)
}
