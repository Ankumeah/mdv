package main

import "fmt"

func main() {
	file := handleArgs()
	html := inject(genHTML(file, readFile(file)))

	fmt.Println(html)

	showWebview(file, html)
}
