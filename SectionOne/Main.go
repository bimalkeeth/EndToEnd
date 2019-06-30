package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	s:=make([]int,5)
	s[0],s[1],s[2],s[3],s[4]=1,2,3,4,5
	fmt.Println(s)
	s1:=s[2:5]
	fmt.Println(s1)
	fmt.Println(s1[:cap(s1)])
	s2:=make([]int,2)
	copy(s2,s[1:3])
	fmt.Println(s2)

}
