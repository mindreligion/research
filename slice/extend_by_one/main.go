package main

import (
	"fmt"
)

func extendRithgOne(s []int, e int) []int {
	n := len(s)
	s = s[:n+1]
	s[n] = e
	return s
}

func main() {
	var b [10]int
	s := b[0:0]
	for i := 0; i < 20; i++ {
		s = extendRithgOne(s, i)
		fmt.Println(s)
	}
}
