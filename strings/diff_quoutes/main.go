package main

import (
	"fmt"
)

func main() {
	a := `the
	new line\n
	text`
	b := "the" +
		"new \vline" +
		"text"
	fmt.Println(a, b)
}
