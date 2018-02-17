package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/pkg/app"
	"github.com/perriea/tfwrapper/pkg/aws"
)

var (
	data    []string
	profile *string
	err     error
)

// Action : application terraform
func Action(action string, args []string) {
	data = append([]string{action}, args...)
	app.Exec(data)
}

// ActionAuth : application terraform
func ActionAuth(action string, args []string) {
	err, configuration := readConfig()
	if err {
		fmt.Printf("Account: %s\n", configuration.Aws.General.Account)
		fmt.Printf("Region: %s\n", configuration.Aws.General.Region)
		fmt.Printf("Profile: %s\n", configuration.Aws.Credentials.Profile)
		fmt.Printf("Client: %s\n\n", configuration.Terraform.Vars.ClientName)

		if !existConfig() {
			auth.Run(&configuration.Aws.Credentials.Profile, configuration.Aws.Credentials.Role)
			writeConfig()
		}

		data = append([]string{action}, args...)
		app.Exec(data)
	} else {
		fmt.Println("Config Folder not found !")
	}
}
