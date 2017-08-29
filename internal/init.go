package internal

import (
	"os"

	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run: runInit,

		Usage: "init [<DIRECTORY>]",
		Name:  "init",
		Long:  "Creates a configuration file in the local directory",
	})
}

func runInit(c *command.Command, r *command.Runtime) {
	var directory string
	if len(r.Args) > 0 {
		directory = r.Args[0]
		err := os.MkdirAll(directory, 0755)
		utils.Check(err, "")
		r.Config.Directory = directory
	}

	r.Config.Write()
	utils.Println("Wrote configuration to " + r.Config.File())
}
