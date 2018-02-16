package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/pkg/app"
	"github.com/perriea/tfwrapper/pkg/aws"
)

var (
	profile *string
	err     error
)

// Apply : application terraform
func Apply(args []string) {
	err, configuration := ReadConfig()
	if err {
		fmt.Printf("Account: %s\n", configuration.Aws.General.Account)
		fmt.Printf("Region: %s\n", configuration.Aws.General.Region)
		fmt.Printf("Profile: %s\n", configuration.Aws.Credentials.Profile)
		fmt.Printf("Client: %s\n\n", configuration.Terraform.Vars.ClientName)

		auth.Run(&configuration.Aws.Credentials.Profile, configuration.Aws.Credentials.Role)
		app.Exec("terraform", args)
	} else {
		fmt.Println("Config Folder not found !")
	}
}

// Init : application terraform
func Init(args []string) {
	err, configuration := ReadConfig()
	if err {
		fmt.Printf("Account: %s\n", configuration.Aws.General.Account)
		fmt.Printf("Region: %s\n", configuration.Aws.General.Region)
		fmt.Printf("Profile: %s\n", configuration.Aws.Credentials.Profile)
		fmt.Printf("Client: %s\n\n", configuration.Terraform.Vars.ClientName)

		auth.Run(&configuration.Aws.Credentials.Profile, configuration.Aws.Credentials.Role)
		app.Exec("terraform", args)
	} else {
		fmt.Println("Config Folder not found !")
	}
}
