package internal

import (
	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.Register(&command.Command{
		Run:  runHelp,
		Name: "help",
	})
}

func runHelp(c *command.Command, r *command.Runtime) {
	for _, command := range Runner.Commands {
		if command.Usage != "" {
			utils.Println("")
			utils.Println(command.Usage)
			utils.Println("  ", command.Long)
		}
	}
}
