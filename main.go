package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func run() {
	cmd := exec.Command("/bin/sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}
