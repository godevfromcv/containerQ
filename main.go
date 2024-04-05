package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
)

func main() {
	// Ваш код тут
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &unix.SysProcAttr{
		Cloneflags: unix.CLONE_NEWNS | unix.CLONE_NEWUTS | unix.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run the command: %v\n", err)
		os.Exit(1)
	}
}
