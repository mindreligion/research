package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		s := []string{"a", "b", "c", "d"}
		fmt.Println(s[r.Intn(len(s))])
	}
}
