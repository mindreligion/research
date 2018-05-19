package main

import "fmt"

type myString struct {
	s string
}

func main() {
	s := myString{`string1`}
	ps := &s
	cs := s
	fmt.Printf(`%T%v %T%v %T%v`, s, s, ps, ps, cs, cs)
	s.s = `string_changed`
	fmt.Printf(`%T%v %T%v %T%v`, s, s, ps, ps.s, cs, cs)
}
