package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func genHTML(filename string, md []byte) string {
	p := modifyFileLinks(filename, parser.NewWithExtensions(joinExtensions()).Parse(md))
	r := html.NewRenderer(html.RendererOptions{Flags: joinFlags()})

	return string(markdown.Render(p, r))
}
