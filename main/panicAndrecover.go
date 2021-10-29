package main

import "fmt"

func init() {

	//testPanich(nil)
	//testPanic2()
	//testPanic3()
	//recoverPanic()
	//recoverPanic()
	recoverFromRuntimePanic()
}

func testPanich(user *string) {
	defer fmt.Println("defere")
	if user == nil {
		panic("user is null") // program will terminate here but befor that all defer function will execute
	}
	fmt.Println("end of func")
}

// Panics can also be caused by errors that happen during the runtime such as trying to access an index that is not present in a slice
func slicePanic() {
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}
func testPanic2() {
	slicePanic()
	fmt.Println("normally returned from main")
}

/*
When a function encounters a panic, its execution is stopped,
any deferred functions are executed and then the control returns to its caller.
This process continues until all the functions of the current goroutine have
returned at which point the program prints the panic message, followed by the stack trace and then terminates.
*/
func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func testPanic3() {
	defer fmt.Println("deferred call in testPanic3")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from testPanic3")
}

// recover panick using defer

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName2(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverPanic() {
	defer fmt.Println("deferred call in recoverPanic")
	firstName := "Elon"
	fullName2(&firstName, nil)
	fmt.Println("returned normally from recoverPanic")
}

// another panic
func recoverInvalidAccess() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func invalidSliceAccess() {
	defer recoverInvalidAccess()
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}

func recoverFromRuntimePanic() {
	invalidSliceAccess()
	fmt.Println("normally returned from main")
}
