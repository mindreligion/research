package main

import (
	"fmt"
)

// Append appends the elements to the slice.
// Efficient version.
func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := len(slice) + len(elements)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	fmt.Println(slice, len(slice), cap(slice))
	slice = slice[:total]
	fmt.Println(slice, len(slice), cap(slice))
	copy(slice[n:], elements)
	return slice
}

func main() {
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := []int{55, 66, 77}
	fmt.Println(slice1)
	slice1 = Append(slice1, slice2...) // The '...' is essential!
	fmt.Println(slice1)
}
