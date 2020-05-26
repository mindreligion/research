/*
	Example shows that you can create struct instance in different ways
	1,2) new(structName) and &structName - create pointer. Structure under it can be accessed without dereferencing
	3) using structName{field values...} creates structure variable not a pointer
		but after creating another copy of it change is source variable affects copy unlike usual variables act

*/
package main

import "fmt"

type myStruct struct {
	age int
}

func main() {
	// create struct using new keyword
	meNew := new(myStruct)
	meNew.age = 1
	fmt.Println("meNew before: ", meNew)
	setAgePointer(meNew, 11)
	fmt.Println("meNew after: ", meNew, " expcected: 11")

	// create struct using curly braces
	meCurlyBraces := myStruct{2}
	fmt.Println("meCurlyBraces before: ", meCurlyBraces)
	setAge(meCurlyBraces, 12)
	fmt.Println("meCurlyBraces after: ", meCurlyBraces, " expected: 12")

	// create struct using ampersand
	meAmpersand := &myStruct{3}
	fmt.Println("meAmpersand before: ", meAmpersand)
	setAgePointer(meAmpersand, 13)
	fmt.Println("meAmpersand after: ", meAmpersand, " expected: 13")

	// In depth with another instance.
	// For pointer.
	meNewCopy := meNew
	meNew.age = 100
	fmt.Println("\nmeNewCopy after changing meNew: ", meNewCopy)

	// For usual variable.
	meAmpersandCopy := meAmpersand
	meAmpersand.age = 100
	fmt.Println("\nmeAmpersandCopy after changing meAmpersand: ", meAmpersandCopy)

	// And what the types?
	fmt.Printf("meCurlyBraces type is %T\nmeNew type is %T", meCurlyBraces, meNew)

	// small research of same cases with variables
	v := 3
	pv := &v
	cv := v
	*pv = 33
	fmt.Println()
	fmt.Printf("v: %v\n*pv: %v\ncv: %v\n", v, *pv, cv)
	cv = 17
	fmt.Println()
	fmt.Printf("v: %v\n*pv: %v\ncv: %v\n", v, *pv, cv)
}

func setAge(obj myStruct, value int) {
	obj.age = value
}

func setAgePointer(pObj *myStruct, value int) {
	pObj.age = value
}
