package main

import (
	"fmt"
	"time"
)

func init() {
	//processChanels()
}

func worker1(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- 1
	}

}
func worker2(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- 2
	}
}
func worker3(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- 3
	}
}

func processChanels() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)

	go worker1(ch1)
	go worker2(ch2)
	go worker3(ch3)
	for i := 0; i < 10; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("msg from ch1  : ", msg)
		case msg := <-ch2:
			fmt.Println("msg from ch2  : ", msg)
		case msg := <-ch3:
			fmt.Println("msg from ch3  : ", msg)
		}
	}

}
