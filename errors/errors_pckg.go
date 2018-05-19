package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(func() string {
		err := errors.New(`Hello I'm a error occurred here`)
		return err.Error()
	}())
}
