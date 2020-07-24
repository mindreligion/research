package main

import "fmt"

func main() {
	fmt.Println("what is your favorite color?")
	var color int
	fmt.Scan(&color)
	fmt.Println("your favorite color is", color)

	//fmt.Println("my favorite color is", color)
}
