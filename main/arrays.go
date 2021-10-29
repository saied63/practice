package main

import (
	"fmt"
)

func init() {
	//fmt.Println("hello from init inside array and slice")
}

func testArray() {
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = i + 1
		fmt.Printf("array in index (%d) is (%d)", i, array[i])
	}
}

func testarrayInitialize() {
	array1 := [3]int{1, 2, 3}
	var array2 = [3]string{"a", "b"}
	fmt.Printf("third party of non defined array is %s \n", array2[2])

	for i := 0; i < len(array1); i++ {
		fmt.Printf("array 1 is ")
		fmt.Println(array1)
	}

	for i := 0; i < len(array2); i++ {
		fmt.Printf("array 2 is ")
		fmt.Println(array2)
	}

	array3 := []float32{2}
	array3[0] = 1
	fmt.Println(array3[0])
	fmt.Println(array3)
}

func testRangeOnIterator() {
	array := [6]float32{1.5, 2.5, 3.5, 4.5, 0.1, 1.02}
	for index, value := range array {
		fmt.Printf("value for index %d is %0.1f  : ", index, value)
	}
	// ins case we don't need index
	for _, v := range array {
		fmt.Printf("  value is %0.1f : ", v)
	}
}

func testMultidecimalIteratorWithRang() {
	array := [3][2]string{{"a1", "a2"}, {"b1", "b2"}, {"c1", "c2"}}
	for i1, v1 := range array {
		fmt.Printf("array decimal is %d and value is : ", i1)
		fmt.Println(v1)
		for i2, v2 := range array[i1] {
			fmt.Printf("in decimal %d and index %d value is : %s \n", i1, i2, v2)

		}
	}
}

// slices are a refrence from a pice of array
func testSclice() {
	array := []int{20, 30, 40, 50, 60}
	for i := range array {
		array[i]++
	}
	fmt.Println(array)
	slice := array[:]
	fmt.Println(slice)
	slice[5] = 70
	fmt.Println(slice)
}

func testSliceLengthAndCapacity() {
	array := []int{1, 2, 3, 4, 5, 6, 8}
	fmt.Printf("len : %d and cap : %d ", len(array), cap(array))
	slice1 := array[1:3]
	fmt.Printf("len : %d and cap : %d ", len(slice1), cap(slice1))
}

func makeSliceByMake() {
	//test := [5]int{1, 2, 3, 4, 5} // test is array because length is assigned
	//test = test[1:2]
	slice := make([]int, 2, 10)
	slice = slice[:cap(slice)] // resize slice until it's capacity
	fmt.Println(slice[5])
}

func testAppend() {
	slice := []int{0, 1, 2, 3, 4}
	slice = append(slice, 7)
	fmt.Println((slice))
}

func nilSlice() {
	var slice []string
	if slice == nil {
		fmt.Printf("slice is nill len is : %d  capacity is %d", len(slice), cap(slice))
	}
	slice = append(slice, "saied", "ali")
	fmt.Println(slice)

}

/*
type of slice is like this :
type slice struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}
hence when we append slice the new one will be new slice because length and capacity is changed
*/
func testSliceRefrence() {
	slice := []int{2, 4, 6, 7}
	slice = slice[0:2]
	reciveRefrenceOfSlice(slice)
	fmt.Println(slice)
}

func reciveRefrenceOfSlice(sc []int) {
	//sc = append(sc, 8)
	sc[0] -= 2
	fmt.Println(sc)
}

//memory managment
/* slice is a refrence on top of array so if we call a function with an array and
then pick a slice from it , if we return that slice the array in local function will be
remain in memory and garbag collector can't remove it so to solve this issue we can make a
new slice and the copy target slice to it then return new slice so that previous slice and
depended array can be remove*/
func aFuncWithLocalArray() []int {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:2]
	// if we just return slice the array will be remain in memory
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func useCoppySlice() {
	slice := aFuncWithLocalArray()
	fmt.Println(slice)
}
