package cmd

import (
	"fmt"

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
					if err = wrapper.PreExecCmd(object().Authenticated, object().Quiet); err != nil {
						fmt.Println(err.Error())
					}
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
