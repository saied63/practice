package main

import (
	"fmt"
	"time"
)

func init() {
	//testGoRoutine()
}

func testGoRoutine() {
	go goroutineCaller()
}

func goroutineCaller() {

	go goRoutine1()
	go goRoutine2()
	time.Sleep(2 * time.Second)
	fmt.Println("hello from caller")
}

func goRoutine1() {

	for i := 0; i < 100; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("hello from goRoutine 1")
	}

}

func goRoutine2() {
	for i := 0; i < 100; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("hello from goRoutine 2")
	}
}
