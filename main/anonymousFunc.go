package main

import (
	"fmt"
)

func init() {
	//callAnnonymousFunc()
}

func callAnnonymousFunc() {
	fmt.Println("start")

	messsage := func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	}
	messsage()
	//or

	func() {
		fmt.Println("anonymouse mode 2")
	}() // this mean call function

	fmt.Println("End")
}
