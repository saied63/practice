package main

import (
	"fmt"
)

func init() {
	//callVariadicFunc()
	//callAnotherVariadicSlice()
}

func testVariadicFunc(a string, b ...int) {
	fmt.Println(a)
	for index, value := range b {
		fmt.Printf("index : %d , value %d", index, value)
	}
	b[0] = 62

}

func callVariadicFunc() {
	testVariadicFunc("saied", 1, 2, 3, 4, 5, 6, 7)
	// pass slise to variadic function
	var slice []int = []int{2, 3, 6, 5}
	testVariadicFunc("pass slice : ", slice...)
	fmt.Println(slice)
}

func callAnotherVariadicSlice() {
	var slice []string = []string{"saied", "ali", "moeen"}
	anotherVariiadicFunc(slice...)
	fmt.Println(slice)
}

func anotherVariiadicFunc(s ...string) {
	s[0] = "changed"
	s = append(s, "append to slice")
	fmt.Println(s)
}
