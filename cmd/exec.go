package cmd

import (
	"github.com/perriea/tfwrapper/pkg/wrapper"
)

// Run command
func Run() {
	var command bool

	wrapper.DisplayVersion()

	if len(args) > 0 {
		// verify first argument
		for cli, object := range c {
			if cli == args[0] {
				if object().Terraform {
					wrapper.ExecCmd(args, object().Authenticated, object().Quiet)
				}
				// future external commands
				command = true
			}
		}

		// check if a command have been executed
		if !command {
			wrapper.Help()
		}
	} else {
		wrapper.Help()
	}
}
