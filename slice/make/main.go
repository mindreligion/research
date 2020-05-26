package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10, 15)
	fmt.Printf("Length %d cap %d", len(s), cap(s))
	sn := make([]int, 40)

	s[5] = 7
	copy(s, s[1:])

	s[5] = 8
	fmt.Println(sn, s)
	/*
		var b [10]int
		s := b[0:0]
		for i := 0; i< 20; i++ {
			s = extendRithgOne(s, i)
			fmt.Println(s)
		}
	*/
}
