package wrapper

import (
	"os"
	"os/exec"

	"github.com/perriea/tfversion/terraform"
)

// execution : Execute terraform command
func execution(args []string) {

	version, err := readConfigHCL()
	if err != nil {
		err = terraform.Install(terraformDefaultVersion, false)
		FatalError(err)
	} else {
		err = terraform.Install(version, false)
		FatalError(err)
	}

	// Prepare command with arguments
	cmd := exec.Command(binary, args...)

	// redirect stdout/err/in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute command
	cmd.Run()
}
