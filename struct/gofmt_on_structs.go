package main

import "fmt"

type test struct {
	name   string
	age    int
	slogan string
}

func main() {

	t := test{
		name:   "John",
		age:    13,
		slogan: "For the evel",
	}
	fmt.Println(t)
}
