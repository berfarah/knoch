package internal

import (
	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/utils"
)

func init() {
	Runner.RegisterHelp(&command.Command{
		Run:  runHelp,
		Name: "help",
	})
}

func runHelp(c *command.Command, r *command.Runtime) {
	for _, command := range Runner.SortedCommands() {
		if command.Usage != "" {
			utils.Println("")
			utils.Println(command.Usage)
			utils.Println("  ", command.Long)
		}
	}
}
