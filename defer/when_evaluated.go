package main

import "fmt"

func a() (i int) {
	i = 0
	defer func(i int) {
		i = b(i)
		i += 2
		fmt.Println(i)
	}(i)
	i += 4
	fmt.Printf(`before return, i=%v\n`, i)
	return
}
func b(i int) int {
	fmt.Println(`in b `, i)
	return i + 1
}
func main() {

	fmt.Println(`main result: `, a())
}
