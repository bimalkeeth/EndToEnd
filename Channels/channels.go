package main

import (
	"fmt"
	"sync"
)

type T int

func IsClosed(c chan T) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

type ClosableChan struct {
	ch       chan T
	onece    sync.Once
	isClosed bool
}

func (cc *ClosableChan) Close() {
	cc.onece.Do(func() {
		close(cc.ch)
		cc.isClosed = true
	})
}
func (cc ClosableChan) IsClosed() bool {
	return cc.isClosed
}

func ping1(ch chan string) {
	ch <- "ping-1 successful"
}
func ping2(ch chan string) {
	ch <- "ping-2 successful"
}

func main() {
	var c = make(chan T, 1)
	fmt.Println(IsClosed(c))
	fmt.Println(cap(c))
	fmt.Println(len(c))
	c <- 10
	fmt.Println(cap(c))
	fmt.Println(len(c))

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)

}
