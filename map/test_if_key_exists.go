package main

import "fmt"

func main() {
	m := map[string]string{
		"blow-fish": "carp",
		"empty-dot": "swift",
	}
	v, ok := m["blow-fish"]
	fmt.Println(v, ok)
	a := [...]int{
		1,
		2,
		3,
	}
	delete(m, "blow-fish")
	fmt.Println(a, m)
}
