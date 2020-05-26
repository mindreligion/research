package main

import "fmt"

func main() {
	initial := []string{"banana", "apple", "orange"}
	// in this case capacity restriction prevents initial slice override
	new := initial[:1:1]
	newAppended := append(new, "kiwi")
	fmt.Println(newAppended, initial)
}
