package command

type Runtime struct {
	Executable string
	Command    string
	Args       []string
}

func NewRuntime(args []string) *Runtime {
	runtime := &Runtime{
		Executable: args[0],
	}
	parseArgs(args[1:], runtime)

	return runtime
}

const (
	helpFlag    = "--help"
	versionFlag = "--version"
)

func parseArgs(args []string, r *Runtime) {
	r.Args = []string{}

	if len(args) == 0 {
		return
	}

	if args[0] == helpFlag {
		r.Command = "help"
		return
	}

	for _, arg := range args {
		if arg == versionFlag {
			r.Command = "version"
			return
		}
	}

	r.Command = args[0]
	r.Args = args[1:]
}
