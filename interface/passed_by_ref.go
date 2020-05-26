package main

import "fmt"

type I interface {
}

// Don't understand why i - reference type. What conditions should be met to pass i by ref implicitly.
func main() {
	var i I
	var a int = 17
	i = 0
	change(i)
	i = a
}

func change(i I) {
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
	subChange(&i)
	fmt.Printf("%v\n", i)
}

func subChange(i *I) {
	fmt.Printf("%v\n", *i)
	*i = 7
	fmt.Printf("%v\n", *i)
}
