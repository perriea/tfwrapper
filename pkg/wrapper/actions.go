package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/pkg/aws"
)

// Action application Terraform
func Action(action string, args []string) {
	data = append([]string{action}, args...)
	execution(data)
}

// ActionAuth application terraform
func ActionAuth(action string, args []string, quiet bool) {
	err, configuration := readConfig()
	if err {
		if !quiet {
			fmt.Printf("\033[1;31mAccount: \033[1;0m%s\n", configuration.Aws.General.Account)
			fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", configuration.Aws.General.Region)
			fmt.Printf("\033[1;34mProfile: \033[1;0m%s\n", configuration.Aws.Credentials.Profile)
			fmt.Printf("\033[1;33mClient: \033[1;0m%s\n", configuration.Terraform.Vars.ClientName)
			fmt.Print("--------------------------------------\n\n")
		}

		fmt.Println(!existVarsConfig())
		if !existVarsConfig() {
			auth.Run(&configuration.Aws.Credentials.Profile, configuration.Aws.Credentials.Role)
			FatalError(writeVarsConfig())
		}

		data = append([]string{action}, args...)
		FatalError(execution(data))
	} else {
		fmt.Println("\033[1;31mError: No configuration files found!\nApply requires configuration to be present.")
	}
}
