package main

import "fmt"

func main() {
	// Same text using:
	// hexadecimal bytes
	a := "\xe6\x97\xa5"
	// input text
	a2 := "日"
	// octal bytes
	octalBytes := "\346\227\245"

	b := '\xe6'
	b2 := '日'
	b3 := int32(26085)
	const shota = iota
	fmt.Println(a, string(a2[0]), string(b), b2, string(b3), octalBytes, shota)
	fmt.Printf("%T\n", shota)
	if a == a2 && a == string(b2) && octalBytes == a {
		fmt.Println("equal")
	}
}
