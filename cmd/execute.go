package cmd

import (
	"github.com/perriea/tfwrapper/pkg/wrapper"
)

// Run command
func Run() {
	wrapper.DisplayVersion()
	// check len args
	if len(args) > 0 {
		// verify first arg
		for cli, object := range c {
			// if command exist
			if cli == args[0] {
				if object().Authenticated {
					wrapper.ActionAuth(args, object().Quiet)
				} else {
					wrapper.Action(args)
				}
			}
		}
		// no command
	} else {
		// launch helper
		wrapper.Help()
	}
}
