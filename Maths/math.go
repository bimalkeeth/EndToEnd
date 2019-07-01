package main

import (
	"fmt"
	"math"
)

func main() {
	Example()
}

func Example() {
	i := 25
	result := math.Sqrt(float64(i))
	fmt.Println(result)

	result = math.Ceil(9.5)
	fmt.Println(result)

	result = math.Floor(9.5)
	fmt.Println(result)
}
