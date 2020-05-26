package main

import "fmt"

func main() {
	a := "\063"
	//b := 'a'
	s := "a"
	fmt.Printf("%T %v %T %v", a, a, s[0], s[0])
}
