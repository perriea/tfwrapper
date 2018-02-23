package wrapper

import (
	"os"
	"os/exec"
	"strings"
)

// func Lookup() {}

// execution : Execute terraform command
func execution(args []string) error {

	// err = terraform.Install("0.10.8")
	// FatalError(err)

	// Prepare command with arguments
	cmd := exec.Command(binary, strings.Join(args, " "))

	// redirect stdout/err/in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Execute command
	err = cmd.Run()
	if Error(err) {
		return err
	}

	return nil
}
