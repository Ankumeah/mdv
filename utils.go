package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func joinExtensions() parser.Extensions {
	var ext parser.Extensions
	for _, e := range parserExtensions {
		ext |= e
	}

	return ext
}

func joinFlags() html.Flags {
	var flag html.Flags
	for _, f := range htmlFlags {
		flag |= f
	}

	return flag
}

func readFile(file string, exitOnError bool) []byte {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		if exitOnError {
			os.Exit(1)
		}
		return []byte("Invalid file")
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		if exitOnError {
			os.Exit(1)
		}
		return []byte("Invalid file")
	}

	return content
}
