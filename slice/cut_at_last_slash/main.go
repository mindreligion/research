package main

import (
	"bytes"
	"fmt"
)

func truncateAtFinalSlash(sp *[]byte) {
	i := bytes.LastIndex(*sp, []byte{'\\'})
	if i >= 0 {
		*sp = (*sp)[:i]
	}
}

func main() {
	s := "http://the.net/very/good/catalog/item1"
	slice := []byte(s)
	fmt.Println("Before: ", s, string(slice))
	truncateAtFinalSlash(&slice)
	fmt.Println("After: ", s, string(slice))
}
