package main

import (
	"os/exec"
	"log"
	"os"
	"strings"
	"path/filepath"
	"syscall"
)

func main() {
	baseName := filepath.Base(os.Args[0])
	args := os.Args[1:]

	if ! strings.HasPrefix(baseName, "wsl") {
		log.Println("Basename does not have prefix: ", baseName)
		log.Fatal("Rename this binary to command name with prefix `wsl`. For example, rename to `wslgit` to run git command on WSL.")
	}

	commandName := strings.TrimSuffix(strings.TrimPrefix(baseName, "wsl"), ".exe")
	commandLine := append([]string{commandName}, args...)

	cmd := exec.Command("wsl", commandLine...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		status := 1
		if exitError, ok := err.(*exec.ExitError); ok {
			if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
				status = waitStatus.ExitStatus()
			}
		}
		os.Exit(status)
	}
}
