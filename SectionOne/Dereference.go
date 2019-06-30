package main

import "fmt"

func main() {

	var PI *int
	I:=3
	PI=&I
	increment(PI)
	fmt.Println(I)
}

func increment(pi *int){
	*pi++
	fmt.Println(*pi)
	*pi++
}
