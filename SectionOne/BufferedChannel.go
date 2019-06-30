package main

import (
	"fmt"
	"time"
)

func main() {

	select{
	  case v1:=<-waitAndSend(3,1):
	       fmt.Println(v1)
	  case v2:=<-waitAndSend(5,1):
	  	   fmt.Println(v2)
	default:
		fmt.Println("all channel are slow")
	}

}

func waitAndSend(v,i int ) chan int{

	ret :=make(chan int)

	go func(){
		time.Sleep(time.Duration(i) * time.Second)
		ret<-v

	}()

   return ret
}