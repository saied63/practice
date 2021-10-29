package main

import (
	"fmt"
)

func init() {
	//testPointer()
	//zerovalueofPointerIsNil()
	//makePointerWithNewFunc()
	//returnPointerFromFunc()
}

func testPointer() {
	a := 1
	var b *int = &a
	getPointer(b)
	fmt.Println(a)
}

func getPointer(p *int) {
	*p++
	fmt.Println(p)
}

func zerovalueofPointerIsNil() {
	var a *int
	if a == nil {
		b := 2
		a = &b
	}
	fmt.Println(a)
}

func makePointerWithNewFunc() {

	size := new(int)
	*size = 2
	fmt.Println(*size)
	*size += 2
	fmt.Println(*size)
}

func returnPointerFromFunc() {
	p := useReturnPointer()
	fmt.Println(*p)
}

func useReturnPointer() *int {
	a := 11
	return &a
}
