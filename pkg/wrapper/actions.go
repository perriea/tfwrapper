package wrapper

import (
	"fmt"
	"os"

	"github.com/perriea/tfwrapper/pkg/aws"
)

// Action application Terraform
func Action(args []string) {
	execution(args)
}

// ActionAuth application terraform
func ActionAuth(args []string, quiet bool) {
	var (
		configuration YAMLConfig
	)

	// read YAML config
	configuration, err = readYAMLConfig()
	if err != nil {
		fmt.Println("\033[1;31mError: No configuration files found!\nApply requires configuration to be present.")
		os.Exit(1)
	}

	// if the action must be silent
	if !quiet {
		fmt.Printf("\033[1;31mAccount: \033[1;0m%s\n", configuration.AWS.General.Account)
		fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", configuration.AWS.General.Region)
		fmt.Printf("\033[1;34mProfile: \033[1;0m%s (%s)\n", configuration.AWS.Credentials.Profile, configuration.AWS.Credentials.Role)
		fmt.Printf("\033[1;35mEnv: \033[1;0m%s\n", configuration.AWS.General.Env)
		fmt.Print("--------------------------------------\n\n")
	}

	// Check file config AWS
	if !existVarsConfig() {
		authAWS.Run(&configuration.AWS.Credentials.Profile, configuration.AWS.Credentials.Role, durationSess)
		if err = writeVarsConfig(); err != nil {
			panic(err)
		}
	}

	execution(args)
}
