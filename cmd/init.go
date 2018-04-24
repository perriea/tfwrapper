package cmd

import (
	"os"
)

var (
	c    map[string]CommandFactory
	args []string
	err  error
)

// CommandFactory mapping
type CommandFactory func() Command

// Command Terraform
type Command struct {
	Authenticated bool
	Terraform     bool
	Description   string
	Quiet         bool
}

func init() {
	// stdin terminal
	args = os.Args[1:]

	// CommandFactory mapping
	c = map[string]CommandFactory{
		"apply": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"console": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"destroy": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"env": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     false,
				Quiet:         true,
			}
		},
		"fmt": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     false,
				Quiet:         true,
			}
		},
		"get": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     false,
				Quiet:         true,
			}
		},
		"graph": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     false,
				Quiet:         true,
			}
		},
		"import": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"init": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"output": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     false,
				Quiet:         false,
			}
		},
		"plan": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"push": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     false,
				Quiet:         false,
			}
		},
		"refresh": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"show": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"taint": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"untaint": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"validate": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     true,
				Quiet:         true,
			}
		},
		"version": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     false,
				Quiet:         true,
			}
		},

		// DEBUG
		"debug": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"force-unlock": func() Command {
			return Command{
				Authenticated: true,
				Terraform:     true,
				Quiet:         false,
			}
		},
		"state": func() Command {
			return Command{
				Authenticated: false,
				Terraform:     true,
				Quiet:         true,
			}
		},
	}
}
