package main

import "fmt"

type A struct {
	B
	C
}

type B struct {
	a int
}

type C struct {
	b int
}

// You can/must reference A.B.a as A.a
// It's seems like convenient way to use composition
func main() {
	a := A{}
	a.a = 2
	a.b = 3
	fmt.Printf("%v", a)
}
