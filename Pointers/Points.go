package main

import "fmt"

type P int

func add2(a int) func() int {
	return func() int {
		return a + 2
	}
}
func addRef(a *int) func() int {
	return func() int {
		return *a + 2
	}
}
func add2Ref2(a *int) func() int {
	var b = *a
	return func() int {
		return b + 2
	}

}

func main() {

	n := new(P)
	*n = 1
	fmt.Println(&*n == n)
	a2 := add2(1)
	fmt.Println(a2())

	var a = 3
	a2r := addRef(&a)
	fmt.Println(a2r())

	a = 5
	fmt.Println(a2r())

	a2r2 := add2Ref2(&a)
	fmt.Println(a2r2())

	s2 := [...]int{1, 2, 3}
	for _, v := range &s2 {
		fmt.Print(v)
	}

}
