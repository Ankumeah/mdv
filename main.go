package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	file := handleArgs()
	contents := readFile(file, true)

	parser := parser.NewWithExtensions(joinExtensions()).Parse(contents)
	doc := modifyFileLinks(file, parser)
	renderer := html.NewRenderer(html.RendererOptions{Flags: joinFlags()})

	html := inject(string(markdown.Render(doc, renderer)))

	showWebview(file, html)
}
