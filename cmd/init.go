package cmd

import (
	"os"
)

type CommandFactory func() Command

// Command Terraform
type Command struct {
	Description   string
	Authenticated bool
	InitPullState bool
	Quiet         bool
}

var (
	c    map[string]CommandFactory
	args []string
	err  error
)

func init() {
	c = map[string]CommandFactory{
		"apply": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"console": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"destroy": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"env": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: false,
				Quiet:         true,
			}
		},
		"fmt": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: false,
				Quiet:         true,
			}
		},
		"get": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: false,
				Quiet:         true,
			}
		},
		"graph": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: false,
				Quiet:         true,
			}
		},
		"import": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"init": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: true,
				Quiet:         true,
			}
		},
		"output": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: false,
				Quiet:         false,
			}
		},
		"plan": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"push": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: false,
				Quiet:         false,
			}
		},
		"refresh": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"show": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"taint": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"untaint": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"validate": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: true,
				Quiet:         true,
			}
		},
		"version": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: false,
				Quiet:         true,
			}
		},

		// DEBUG
		"debug": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"force-unlock": func() Command {
			return Command{
				Authenticated: true,
				InitPullState: true,
				Quiet:         false,
			}
		},
		"state": func() Command {
			return Command{
				Authenticated: false,
				InitPullState: true,
				Quiet:         true,
			}
		},
	}

	args = os.Args[1:]
}
