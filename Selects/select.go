package main

import (
	"fmt"
	"sync"
)

func hello(w *sync.WaitGroup) {
	fmt.Println("Hello")
	w.Done()
}

func main() {
	var w sync.WaitGroup
	w.Add(1)
	go hello(&w)
	w.Wait()
}
