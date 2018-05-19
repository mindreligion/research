package main

import (
	"fmt"
)

type i interface {
	juggle(string)
}

type A string

type B string

func (A) juggle(s string) {}
func (B) juggle(t string) {}
func jobdone(x i) {
	a, ok := x.(A)
	fmt.Printf(`%v, %T, %v, %T;`, a, a, ok, ok)
	b, ok := x.(B)
	fmt.Printf(`%v, %T, %v, %T;`, b, b, ok, ok)
}
func main() {
	jobdone(A(`the a`))
	var b B = `the b`
	jobdone(b)

	println()
	type I interface {
	}
	var x I = 7
	fmt.Printf(`%T`, x)
	fmt.Printf(`%T`, x.(int))

}
