package main

import "fmt"

func main() {
	// if use exponent the result will be float64 - not int, event if base have no decimal point
	i := 10e10
	fmt.Printf("%T", i)
}
