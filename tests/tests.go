package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var res []string
	for i := 2; i <= 16000; i++ {
		res = append(res, fmt.Sprintf("%v,100000,300", i))
	}
	if err := ioutil.WriteFile("./lvl.csv", []byte(strings.Join(res, "\n")), 0777); err != nil {
		panic(err)
	}
}
