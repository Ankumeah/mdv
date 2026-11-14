package main

import (
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var parserExtensions = []parser.Extensions{
	parser.Tables,
	parser.FencedCode,
	parser.Strikethrough,
	parser.SuperSubscript,
}

var htmlFlags = []html.Flags{
	html.SkipHTML,
	html.CompletePage,
}
