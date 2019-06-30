package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/Users/bimalkeeth/Downloads/3.png")
	text, _ := client.HOCRText()
	fmt.Println(text)
	// Hello, World!
}