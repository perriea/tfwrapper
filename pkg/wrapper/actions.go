package wrapper

import (
	"fmt"
)

// Action application Terraform
func Action(args []string) {
	if err = switchVersion(); err != nil {
		panic(err)
	}

	execCmd(args)
}

// ActionAuth application terraform
func ActionAuth(args []string, quiet bool) {
	// read YAML config
	yamlProvider, err = readYAMLConfig()
	if err != nil {
		fmt.Printf("\033[1;31m%s\033[1;0m\n", err.Error())
	} else {
		if err = switchVersion(); err != nil {
			panic(err)
		}

		lookupProvider(quiet)
		execCmd(args)
	}
}
