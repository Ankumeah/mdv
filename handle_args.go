package main

import (
	"fmt"
	"os"
)

const helpText = `
mdv: No BS, minimal markdown viewer

Usage: mdv [FILE]`

func handleArgs() string {
	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "Nice potato there! (argc < 1)")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Println("No file provided")
		fmt.Println(helpText)
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "Expected a single input")
		os.Exit(1)
	} else if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(helpText)
		os.Exit(0)
	}

	return os.Args[1]
}
