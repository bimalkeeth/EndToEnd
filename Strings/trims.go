package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "\t Hello world"
	trims := strings.TrimSpace(greeting)
	fmt.Printf("%d %s\n", len(trims), trims)
}
