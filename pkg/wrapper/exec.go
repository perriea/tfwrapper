package wrapper

import (
	"os"
	"os/exec"
)

// func Lookup() {}

// execution : Execute terraform command
func execution(args []string) {

	// err = terraform.Install("0.10.8")
	// FatalError(err)

	// Prepare command with arguments
	cmd := exec.Command(binary, args...)

	// redirect stdout/err/in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute command
	cmd.Run()
}
