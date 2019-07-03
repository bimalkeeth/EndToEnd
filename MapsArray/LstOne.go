package main

import "fmt"

func main() {
	intSlice := []int{1, 5, 5, 5, 5, 7, 8, 6, 6, 6}
	fmt.Println(intSlice)
	unuqueSlice := unique(intSlice)
	fmt.Println(unuqueSlice)

}

func unique(ints []int) []int {
	keys := make(map[int]bool)
	uniqueElements := []int{}
	for _, entry := range ints {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements

}
