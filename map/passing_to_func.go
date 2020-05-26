package main

import "fmt"

func main() {
	m := map[string]string{
		"a": "abraham",
		"b": "boleward",
	}
	fmt.Println(m)
	// here map passed by value, but changes effects initial map
	// so maps are always being passed by ref
	removeVal(m, "a")
	fmt.Println(m)
}

func removeVal(m map[string]string, k string) {
	delete(m, k)
}
