package wrapper

import (
	"fmt"

	"github.com/perriea/tfwrapper/pkg/aws"
)

func preExecAWS() {
	if !validConfigAuth() {
		// Auth AWS STS
		authAWS.Run(&yamlProvider.Provider.Credentials.Profile, yamlProvider.Provider.Credentials.Role, durationSess)
		if err = writeAuthConfig("aws"); err != nil {
			panic(err)
		}
	}
}

func preExecGCP() {
	if !validConfigAuth() {
		if err = writeAuthConfig("gcp"); err != nil {
			panic(err)
		}
	}
}

func lookupProvider(quiet bool) {
	switch yamlProvider.Cloud {
	case "aws":
		if !quiet {
			fmt.Printf("\033[1;31mCloud: \033[1;0m%s\n", yamlProvider.Cloud)
			fmt.Printf("\033[1;31mAccount: \033[1;0m%s\n", yamlProvider.Provider.General.Account)
			fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", yamlProvider.Provider.General.Region)
			fmt.Printf("\033[1;34mProfile: \033[1;0m%s\n", yamlProvider.Provider.Credentials.Profile)
			fmt.Printf("\033[1;35mEnv: \033[1;0m%s\n", yamlProvider.Provider.General.Env)
			fmt.Print("--------------------------------------\n\n")
		}

		preExecAWS()
	case "gcp":
		if !quiet {
			fmt.Printf("\033[1;31mCloud: \033[1;0m%s\n", yamlProvider.Cloud)
			fmt.Printf("\033[1;31mProject: \033[1;0m%s\n", yamlProvider.Provider.General.Project)
			fmt.Printf("\033[1;32mRegion: \033[1;0m%s\n", yamlProvider.Provider.General.Region)
			fmt.Printf("\033[1;34mProfile: \033[1;0m%s\n", yamlProvider.Provider.Credentials.Profile)
			fmt.Printf("\033[1;35mEnv: \033[1;0m%s\n", yamlProvider.Provider.General.Env)
			fmt.Print("--------------------------------------\n\n")
		}

		preExecGCP()
	}
}
