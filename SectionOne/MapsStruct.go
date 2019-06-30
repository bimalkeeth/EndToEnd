package main

import "fmt"

func main() {

	 ma:=make(map[string]int)
	 ma["first"]=1
	 ma["second"]=2
	 fmt.Println(ma["first"])
	 if _,ok :=ma["second"];ok{

	 	fmt.Println("Value found")
	 	delete(ma,"second")
	 }
}
