package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/pkg/aws"
)

// func Lookup() {

// }

// Action application Terraform
func Action(args []string) {
	execution(args)
}

// ActionAuth application terraform
func ActionAuth(args []string, quiet bool) {
	configuration, err := readConfigYAML()
	if err != nil {
		fmt.Println("\033[1;31mError: No configuration files found!\nApply requires configuration to be present.")
	}

	// if the action must be silent
	if !quiet && configuration.Aws.General.Account != "" {
		fmt.Printf("\033[1;31mAccount: \033[1;0m%s\n", configuration.Aws.General.Account)
		fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", configuration.Aws.General.Region)
		fmt.Printf("\033[1;34mProfile: \033[1;0m%s\n", configuration.Aws.Credentials.Profile)
		fmt.Printf("\033[1;33mClient: \033[1;0m%s\n", configuration.Terraform.Vars.ClientName)
		fmt.Print("--------------------------------------\n\n")
	}

	// Check file config AWS
	if !existVarsConfig() {
		auth.Run(&configuration.Aws.Credentials.Profile, configuration.Aws.Credentials.Role, durationSess)
		if err = writeVarsConfig(); err != nil {
			panic(err)
		}
	}

	execution(args)
}
