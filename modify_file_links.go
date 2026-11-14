package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown/ast"
)

const absPathErrorMessage = "Got an error while getting the full path of a file. Something has gone seriously wrong. Im non relegious but may god be with you: %v\n"

func isLocalFile(uri []byte) bool {
	return !bytes.HasPrefix(uri, []byte("https://")) &&
		!bytes.HasPrefix(uri, []byte("http://"))
}

func getLocalFileData(mdFile string, uri []byte) []byte {
	mdDir, err := filepath.Abs(filepath.Dir(mdFile))
	if err != nil {
		fmt.Printf(absPathErrorMessage, err)
		os.Exit(1)
	}

	path := string(uri)
	if !filepath.IsAbs(path) {
		path = filepath.Join(mdDir, path)
	}

	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Printf(absPathErrorMessage, err)
		os.Exit(1)
	}

	data := readFile(path)

	dataType := http.DetectContentType(data)
	encodedData := base64.StdEncoding.EncodeToString(data)

	return []byte("data:" + dataType + ";base64," + encodedData)
}

func modifyFileLinks(mdFile string, root ast.Node) ast.Node {
	ast.WalkFunc(
		root,
		func(node ast.Node, entering bool) ast.WalkStatus {
			if link, ok := node.(*ast.Link); ok && entering && isLocalFile(link.Destination) {
				link.Destination = getLocalFileData(mdFile, link.Destination)
			}

			return ast.GoToNext
		},
	)

	return root
}
