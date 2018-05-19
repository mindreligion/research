package main

import (
	"fmt"
	"flag"
	"os"
)

func main () {
	fmt.Println(os.Args)
	for i, v := range os.Args {
		fmt.Printf("Index: %d, value: %s\n", i, v)
	}
	return
	var word2 string
	wordPtr := flag.String("word", "starting_word", "a string parameter")
	flag.StringVar(&word2,"word2", "starting_word2", "a string parameter")
	intPtr := flag.Int("num", 10, "a number param")
	flag.Parse()
	fmt.Printf("Word is: %s\n", *wordPtr)
	fmt.Printf("Word2 is: %s\n", word2)
	fmt.Printf("Num is: %d\n", *intPtr)
	fmt.Printf("Tail is: %v\n", flag.Args())
}
