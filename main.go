package main

func main() {
	file := handleArgs()
	html := inject(genHTML(file, readFile(file)))

	showWebview(file, html)
}
