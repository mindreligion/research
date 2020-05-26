package main

import (
	"fmt"
	"reflect"
)

type IA interface {
	a()
}

type IB interface {
	b()
}

type X int
type Y int
type Z int

func (i X) a() {
	fmt.Println(`a() call`)
}

func (i X) b() {
	fmt.Println(`b() call`)
}
func (i Z) a() {
	fmt.Println(`a() call`)
}

func (i Z) b() {
	fmt.Println(`b() call`)
}

func (i Y) a() {
	fmt.Println(`a() call`)
}

func p(s string) {
	fmt.Println(s, ":")
}

func p2(s string) {

}
func reflectStuff(p string, i interface{}) {
	indent := "\t\t"
	if p != "" {
		fmt.Println(p, `:`)
	}
	//fmt.Printf("\tPrintf type is: %T\n", i)
	fmt.Println(indent, "Reflection type is:", reflect.TypeOf(i))
	//fmt.Printf("\tPrintf value is: %v\n", i)
	fmt.Println(indent, "Reflection value is:", reflect.ValueOf(i).Interface())
}

func mockCall(i interface{}) {
	vx, okx := i.(X)
	okxs := "FAIL"
	if okx {
		okxs = "OK"
	}
	vy, oky := i.(Y)
	okys := "FAIL"
	if oky {
		okys = "OK"
	}
	vz, okz := i.(Z)
	okzs := "FAIL"
	if okz {
		okzs = "OK"
	}
	via, okia := i.(IA)
	okias := "FAIL"
	if okia {
		okias = "OK"
	}
	vib, okib := i.(IB)
	okibs := "FAIL"
	if okib {
		okibs = "OK"
	}
	fmt.Println("\n\ti.(X) value: ", vx, okxs)
	reflectStuff("", vx)
	fmt.Println("\n\ti.(Y) value: ", vy, okys)
	reflectStuff("", vy)
	fmt.Println("\n\ti.(Z) value: ", vz, okzs)
	reflectStuff("", vz)
	fmt.Println("\n\ti.(IA) value: ", via, okias)
	reflectStuff("", via)
	fmt.Println("\n\ti.(IB) value: ", vib, okibs)
	reflectStuff("", vib)
}
func main() {
	var x X = 7
	p(`mockCall(X)`)
	mockCall(x)

	var x_ia IA = X(7)
	p(`mockCall(X->IA)`)
	mockCall(x_ia)

	var x_ib IB = x_ia.(X)
	p(`mockCall(X->IB)`)
	mockCall(x_ib)

	var y Y = 7
	p(`mockCall(Y)`)
	mockCall(y)

	var y_ia IA = Y(7)
	p(`mockCall(Y->IA)`)
	mockCall(y_ia)

	var z Z = 7
	p(`mockCall(Z)`)
	mockCall(z)

	var z_ia IA = Z(7)
	p(`mockCall(Z->IA)`)
	mockCall(z_ia)

	var z_ib IB = Z(7)
	p(`mockCall(Z->IB)`)
	mockCall(z_ib)
}
