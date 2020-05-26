package main

import "fmt"

type sum func(int, int) int

func (fn sum) doSum(a, b int) int {
	return fn(a, b)
}
func main() {
	var fn sum = func(a int, b int) int {
		return a + b
	}
	fmt.Print(fn.doSum(3, 5))
}
