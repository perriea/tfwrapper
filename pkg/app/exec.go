package app

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Exec(s []string) {
	cmd := exec.Command("terraform", strings.Join(s, " "))
	// fmt.Println(cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
