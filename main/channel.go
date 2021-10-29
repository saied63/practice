package main

import (
	"fmt"
	"time"
)

func init() {
	//makeAChannell()
}

func makeAChannell() {
	c := make(chan string)
	go goroutineChan(c)
	result := <-c
	fmt.Printf("end reading from channel : %s ", result)
}

func goroutineChan(doneChannel chan string) {
	time.Sleep(3 * time.Second)
	doneChannel <- "end of chanel"
}
