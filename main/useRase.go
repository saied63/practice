package main

import (
	"fmt"
	"sync"
)

func init() {
	//checkConfilict()
}

func checkConfilict() {
	var wg sync.WaitGroup
	var myInt int
	for i := 0; i < 1000; i++ {
		go makeConfilict(&myInt, &wg)
	}
}

func makeConfilict(myInt *int, wg *sync.WaitGroup) {
	fmt.Println("printing")
	i := *myInt
	fmt.Println("myInt is ", i)
	*myInt++
}

// run programe by go run -race main to find confilict of data usage in cocurrent programs
