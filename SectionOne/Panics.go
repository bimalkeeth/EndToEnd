package main

import "fmt"

func main() {

	fmt.Println("Hello")
	testPanic()
	fmt.Println("World")

}

func testPanic(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("recover panic")
		}
	}()
	panic("A panic happened")
}