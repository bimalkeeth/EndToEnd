package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	alex := User{}
	fmt.Println(alex)
	alexP := &alex
	fmt.Println(alexP)
	var Worker = struct {
		User
		salary float64
	}{salary: 2344.66, User: User{age: 10, name: "bimal"}}
	fmt.Println(Worker)
}
