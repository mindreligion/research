package main

import "fmt"

func main() {
	var a [10]int
	s := a[:3]
	fmt.Printf("%T\n%T\n", a, s)
	ps := &s
	fmt.Printf("%v\n%v\n", s, ps)
	newSlice := s[:5]
	fmt.Printf("%v\n%v\n", newSlice, (*ps))
}
