package main

import (
	"flag"
	"fmt"
	"os"
)
1
func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)

	switch os.Args[1] {
	case "create":
		if err := createCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing 'create' command: %s\n", err)
			os.Exit(1)
		}
		// TODO: Implement logic for creating a container
		fmt.Println("Creating a new container...")

	case "run":
		if err := runCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing 'run' command: %s\n", err)
			os.Exit(1)
		}
		// TODO: Implement logic for running a container
		fmt.Println("Running a container...")

	default:
		fmt.Fprintf(os.Stderr, "Unsupported command '%s'\n", os.Args[1])
		os.Exit(1)
	}
}
