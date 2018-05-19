package main

import "fmt"

func main() {
	var i int = 7
	var in int = 9
	pi := &i
	fmt.Printf("before reasigning:\n\t%v\n", pi)
	pi = &in
	fmt.Printf("after reasigning:\n\t%v\n", pi)
	*pi = 11
	fmt.Printf("final i = %v in = %v", i, in)
}
