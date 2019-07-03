package main

import (
	"fmt"
	"github.com/bradfitz/slice"
)

type NameAge struct {
	Name string
	Age  int
}

func main() {
	nameAges := map[string]int{
		"Michael": 30,
		"John":    25,
		"Jessica": 26,
		"Ali":     18,
	}
	var nameAgeSlice []NameAge = nil
	for key, val := range nameAges {
		nameAgeSlice = append(nameAgeSlice, NameAge{Age: val, Name: key})
	}

	slice.SortInterface(nameAgeSlice[:], func(i, j int) bool {
		return nameAgeSlice[i].Age < nameAgeSlice[j].Age
	})

	fmt.Println(nameAgeSlice)
	fmt.Println(nameAgeSlice)
}
