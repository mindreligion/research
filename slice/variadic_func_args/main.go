package main

import (
	"fmt"
)

func vArgs(elements ...int) {
	fmt.Println(elements)
}

func args(a, b int) {
	fmt.Println(a, b)
}

func main() {
	vArgs(1, 2)
	slice := []int{2, 3}
	// works onli with variable args function
	vArgs(slice...)
	// will not work
	// args(slice...) - func args needs two int params
}
