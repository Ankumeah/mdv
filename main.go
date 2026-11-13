package main

import (
	"fmt"
)

func main() {
	// No Im not a lisp user but I like chaos
	fmt.Println(
		string(
			genHTML(
				readFile(
					handleArgs(),
				),
			),
		),
	)
}
