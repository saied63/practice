package main

import (
	"fmt"
	"time"
)

func init() {
	//unbuffered()
}

func unbuffered() {
	ch := make(chan int, 1)
	go processbuffered(ch)
	//time.Sleep(3 * time.Second)
	//message := <-ch
	//fmt.Println("messsage from chanel ", message)
	fmt.Println("messsage from chanel ")
}

func processbuffered(ch chan int) {
	ch <- 1
	for i := 0; i < 1000; i++ {
		fmt.Println("in processing", i)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("finish process unbuffere")

}
