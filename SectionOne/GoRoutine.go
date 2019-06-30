package main

import "fmt"

func main() {

	quitSignal:=make(chan bool)
    go SayHelloFromGoRouting(quitSignal)
	fmt.Println("Waiting for signal")
	<-quitSignal

}
func SayHelloFromGoRouting(qs chan bool){

	fmt.Println("Hello from new go routing")
	qs<-true
}
