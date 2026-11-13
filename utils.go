package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gomarkdown/markdown"
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

func genHTML(md []byte) []byte {
	return markdown.Render(
		parser.NewWithExtensions(
			joinExtensions(),
		).Parse(md),
		html.NewRenderer(
			html.RendererOptions{
				Flags: joinFlags(),
			},
		),
	)
}

func readFile(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	return content
}
