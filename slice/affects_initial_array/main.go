package main

import (
	"fmt"
)

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[5:10]
	fmt.Println("Before: ", s, a)
	incrementEach(s)
	fmt.Println("After: ", s, a)
}

func incrementEach(s []int) {
	for i := range s {
		s[i]++
	}
}
