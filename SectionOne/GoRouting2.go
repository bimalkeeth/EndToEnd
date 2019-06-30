package main

import (
	"fmt"
	"time"
)

func main() {

	ic:=make(chan int)
	go PeriodicSend(ic)
	for i:=range ic{
		fmt.Println(i)
	}
	if v,ok:=<-ic;ok{

		fmt.Println(v)
	}
}

func PeriodicSend(ic chan int){
	i:=1
	for{
		i++
		ic <- i
		time.Sleep(1 *time.Second)
	}
	close(ic)
}