package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/pkg/aws"
)

func preExecAWS(quiet bool) {
	if !quiet {
		fmt.Printf("\033[1;31mCloud: \033[1;0m%s\n", yamlProvider.Cloud)
		fmt.Printf("\033[1;31mAccount: \033[1;0m%s\n", yamlProvider.Terraform.General.Account)
		fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", yamlProvider.Terraform.General.Region)
		fmt.Printf("\033[1;34mProfile: \033[1;0m%s\n", yamlProvider.Terraform.Credentials.Profile)
		fmt.Printf("\033[1;35mEnv: \033[1;0m%s\n", yamlProvider.Terraform.General.Env)
		fmt.Print("--------------------------------------\n\n")
	}

	if !validConfigAuth() {
		// Auth AWS STS
		authAWS.Run(&yamlProvider.Terraform.Credentials.Profile, yamlProvider.Terraform.Credentials.Role, durationSess)
		FatalError(writeAuthConfig("aws"))
	}
}

func preExecGCP(quiet bool) {
	if !quiet {
		fmt.Printf("\033[1;31mCloud: \033[1;0m%s\n", yamlProvider.Cloud)
		fmt.Printf("\033[1;31mProject: \033[1;0m%s\n", yamlProvider.Terraform.General.Project)
		fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", yamlProvider.Terraform.General.Region)
		fmt.Printf("\033[1;34mProfile: \033[1;0m%s\n", yamlProvider.Terraform.Credentials.Profile)
		fmt.Printf("\033[1;35mEnv: \033[1;0m%s\n", yamlProvider.Terraform.General.Env)
		fmt.Print("--------------------------------------\n\n")
	}

	if !validConfigAuth() {
		FatalError(writeAuthConfig("gcp"))
	}
}

func lookupProvider(quiet bool) {
	switch yamlProvider.Cloud {
	case "aws":
		preExecAWS(quiet)
	case "gcp":
		preExecGCP(quiet)
	}
}
