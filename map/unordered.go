package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9
	m["10"] = 10
	for i, v := range m {
		fmt.Println(i, v)
	}
}
