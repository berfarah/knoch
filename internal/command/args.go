package command

type Args struct {
	Command string
	Params  []string
}

func NewArgs(args []string) *Args {
	a := &Args{}
	parseArgs(args, a)
	return a
}

const (
	helpFlag    = "--help"
	versionFlag = "--version"
)

func parseArgs(args []string, a *Args) {
	a.Params = []string{}

	if len(args) == 0 {
		a.Command = "help"
		return
	}

	for _, arg := range args {
		switch arg {
		case helpFlag:
			a.Command = "help"
			return
		case versionFlag:
			a.Command = "version"
			return
		}
	}

	a.Command = args[0]
	a.Params = args[1:]
}
