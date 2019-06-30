package main

import "fmt"

type TestInterface interface {
	SayHello()
	Say(s string)
}

type testCompleteImpl struct{}
func(p testCompleteImpl)SayHello(){
	fmt.Println("Hi Hello")
}
func(p testCompleteImpl)Say(s string){
	fmt.Println(s)
}


func main() {
   var ti TestInterface
  ti=testCompleteImpl{}
  ti.SayHello()
}
