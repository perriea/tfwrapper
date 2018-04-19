package wrapper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/perriea/tfversion/terraform"
)

func switchVersion() error {
	if yamlProvider.Provider.Version.Terraform != "" {
		if err = terraform.Install(yamlProvider.Provider.Version.Terraform, true); err != nil {
			return err
		}
	}

	return terraform.Install(terraformDefaultVersion, true)
}

// PreExecCmd application Terraform
func preExecCmd(authenticated bool, quiet bool) error {
	// read YAML config
	yamlProvider, err = readYAMLConfig()
	if err != nil {
		fmt.Printf("\033[1;31m%s\033[1;0m\n", err.Error())
		return err
	}

	if err = switchVersion(); err != nil {
		return err
	}

	if authenticated {
		lookupProvider(quiet)
	}

	return nil
}

// ExecCmd : Execute terraform command
func ExecCmd(args []string, authenticated bool, quiet bool) error {
	var cmd *exec.Cmd

	if err = preExecCmd(authenticated, quiet); err != nil {
		return err
	}

	// exec && redirect stdout/err/in
	cmd = exec.Command(binary, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute command
	if err = cmd.Run(); err != nil {
		return err
	}

	return cmd.Wait()
}
