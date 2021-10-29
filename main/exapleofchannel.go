package main

import (
	"fmt"
)

type myInt int

func init() {
	//callQuroutun()
}

func callQuroutun() {
	var a myInt
	ch := make(chan myInt)
	go a.sumSquarandCube(2, ch)
	result := <-ch
	fmt.Println(result)
}

func (sum myInt) sumSquarandCube(a int, done chan myInt) {
	var squar int
	var cube int
	for i := 1; i <= a; i++ {
		squar += i * i
		cube += i * i * i
	}
	sum = myInt(squar + cube)
	done <- sum
}
