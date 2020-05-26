package main

import "fmt"

type A struct {
	i int
}

func (a A) action() {
	fmt.Println(a.i)
}

func (pa *A) paction() {
	fmt.Printf(`%T  %v`, pa, pa)
}
func main() {
	a := A{3}
	pa := &a
	pa.action()
	pa.paction()
}
