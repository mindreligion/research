package main

import "fmt"

func main() {
	a, b := 4, 5
	fmt.Printf("sum(%v, %v) = %v\n", a, b, sum(a, b))
}

func sum(a, b int) int {
	return a + b
}
