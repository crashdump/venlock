package main

import (
	"fmt"

	"github.com/avelino/slugify"
)

// Dummy application

func main() {
	text := "Example slugify"
	fmt.Printf(text + ": " + slugify.Slugify(text))
}
