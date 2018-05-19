package main

import "fmt"

func main() {
	s := []string{"lalaloop", "halahoop"}
	fmt.Println(s)
	change(s)
	fmt.Println(s)
	fmt.Println("But:")
	// Slice i reference type even if it used as underlying type for other named type
	fmt.Println(s)
	st := t(s)
	changeT(st)
	fmt.Println(s)
}

type t []string
func change(s []string) {
	s[1] = "balaboop"
}
func changeT(s t) {
	s[1] = "mamamoop"
}
