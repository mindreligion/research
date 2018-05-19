package main

import (
	"fmt"
)

func substractOneElem(sp *[]int) {
	s := *sp
	*sp = s[0 : len(s)-1]
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[5:10]
	fmt.Println("Before: ", s, a)
	substractOneElem(&s)
	fmt.Println("After: ", s, a)
}
