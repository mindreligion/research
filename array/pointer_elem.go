package main

import "fmt"

func main() {
	var a [3]*int
	b := [...]*int{new(int), new(int), new(int)}
	a = b

	fmt.Println(a, b)

	show(a,b)
	*b[1] = 7
	show(a,b)
}

func show(a, b [3]*int) {
	fmt.Println("a")
	for i, v := range a {
		fmt.Println(i, *v)
	}
	fmt.Println("b")
	for i, v := range b {
		fmt.Println(i, *v)
	}
}
