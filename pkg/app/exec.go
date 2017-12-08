package app

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func Exec(s []string) {

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	cmd := exec.Command("terraform", strings.Join(s, " "))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	go cmd.Run()

	for {
		select {
		case s := <-signalCh:
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Received interrupt...\n")
				cmd.Process.Kill()
				os.Exit(1)
			case syscall.SIGHUP:
				fmt.Println("Received HUP...\n")
				//todo reload
				os.Exit(1)
			}
		}
	}

}
