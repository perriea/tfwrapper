package wrapper

import (
	"os"
	"os/exec"

	"github.com/perriea/tfversion/terraform"
)

func switchVersion() error {
	if yamlProvider.Terraform != "" {
		err = terraform.Install(yamlProvider.Terraform, false)
		return err
	}

	err = terraform.Install(terraformDefaultVersion, false)
	return err
}

// execution : Execute terraform command
func execCmd(args []string) {
	var (
		cmd *exec.Cmd
	)

	// Prepare command with arguments
	cmd = exec.Command(binary, args...)

	// redirect stdout/err/in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute command
	cmd.Run()
}
