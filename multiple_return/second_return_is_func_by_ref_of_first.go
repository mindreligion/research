package main

import "fmt"

func main() {
	fmt.Println(multiReturn())
}

func transformFirst(msg *string) bool {
	*msg = fmt.Sprintf("%s: appended", *msg)
	return true
}

func multiReturn() (string, bool) {
	msg := `initial`
	return msg, transformFirst(&msg)
}
