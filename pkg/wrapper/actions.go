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
			fmt.Printf("Account: %s\n", configuration.Aws.General.Account)
			fmt.Printf("Region: %s\n", configuration.Aws.General.Region)
			fmt.Printf("Profile: %s\n", configuration.Aws.Credentials.Profile)
			fmt.Printf("Client: %s\n\n", configuration.Terraform.Vars.ClientName)
		}

		if !existVarsConfig() {
			auth.Run(&configuration.Aws.Credentials.Profile, configuration.Aws.Credentials.Role)
			FatalError(writeVarsConfig())
		}

		data = append([]string{action}, args...)
		FatalError(execution(data))
	} else {
		fmt.Println(ErrorNotFoundConfig)
	}
}

var ErrorNotFoundConfig = `Error: No configuration files found!
Apply requires configuration to be present.`
