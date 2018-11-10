package main

import "fmt"

func sum(numbers ...float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(110, 45, 32))
}
