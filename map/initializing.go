package main

import "fmt"

func main() {
	ma := make(map[string]int)
	ma["peter"] = 17
	fmt.Printf("%T", ma)
}
