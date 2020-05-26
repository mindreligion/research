package main

import "fmt"

func main() {
	s := []int{}
	c := 0
	n := 1
	maxL := 1000000
	for i := 0; i < maxL; i++ {
		a := []int{}
		j := i
		for j = i; j < i+n && j < maxL; j++ {
			a = append(a, j)
		}
		s = append(s, a...)
		if c != cap(s) {
			fmt.Println(i+1, ":", cap(s)-c, n)
			c = cap(s)
		}
		n++
		i = j - 1
	}
	fmt.Println(len(s))
}
