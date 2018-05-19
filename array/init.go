package main

import (
	"fmt"
)

func main() {
	a := [...]int{1,2,3,4,5,6,7}
	b := [7]int{2:5, 5:2}
	fmt.Println("a: " , a)
	fmt.Println("b: " , b)
	b = a
	fmt.Println("after asignment 'a=b':")
	fmt.Println("a: " , a)
	fmt.Println("b: " , b)
}
