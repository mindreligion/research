package main

import (
	"fmt"
	"reflect"
)

func vArgs(s *[]int, elements ...int) {
	fmt.Printf("Slice pointer passed to func - \n\t- as itIs: %#p\n\t- pointer: %#p\n\t- dereference:%#p\n\n", s, &s, *s)
	fmt.Printf("Type of argument s:\n\t%v\n", reflect.TypeOf(s))

	fmt.Println("Reflection arg s type: ", reflect.TypeOf(s))
	fmt.Println("Reflection arg s value: ", reflect.ValueOf(s))
	fmt.Println("Reflection arg &s type: ", reflect.TypeOf(&s))
	fmt.Println("Reflection arg &s value: ", reflect.ValueOf(&s))
	var ns []int
	fmt.Printf("New slice defined - \n\t- as itIs: %#p\n\t- pointer: %#p\n\n", ns, &ns)

	for _, element := range elements {
		ns = (*s)[:len(*s)+1]
		fmt.Printf("New slice assigned new value - \n\t- as itIs: %#p\n\t- pointer: %#p\n\n", ns, &ns)
		fmt.Println("Reflection arg ns type: ", reflect.TypeOf(ns))
		fmt.Println("Reflection arg ns value: ", reflect.ValueOf(ns))
		fmt.Println("Reflection arg &ns type: ", reflect.TypeOf(&ns))
		fmt.Println("Reflection arg &ns value: ", reflect.ValueOf(&ns))
		fmt.Printf("Slice arg pointer before New slice assigning - \n\t- as itIs: %#p\n\t- pointer: %#p\n\t- dereference:%#p\n\n", s, &s, *s)
		s = &ns
		fmt.Printf("Slice arg pointer after New slice assigning -\n\t- as itIs: %#p\n\t- pointer: %#p\n\t- dereference:%#p\n\n", s, &s, *s)
		fmt.Println("Reflection arg after s type: ", reflect.TypeOf(s))
		fmt.Println("Reflection arg after s value: ", reflect.ValueOf(s))
		fmt.Println("Reflection arg after &s type: ", reflect.TypeOf(&s))
		fmt.Println("Reflection arg after &s value: ", reflect.ValueOf(&s))
		(*s)[len(*s)-1] = element
	}
}

var (
	s []int
)

func main() {
	var a [10]int
	s = a[:0]
	fmt.Printf("Slice of array - \n\t- as itIs: %#p\n\t- pointer: %#p\n\n", s, &s)
	vArgs(&s, 1, 2, 3, 4)
	// expect [1 2 3 4]
	fmt.Println(s)
	// got []  -initial value
	fmt.Println(s[:cap(s)])
	fmt.Printf("Final slice - \n\t- as itIs: %#p\n\t- pointer: %#p\n\n", s, &s)
}
