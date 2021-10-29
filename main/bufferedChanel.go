package main

import (
	"fmt"
	"time"
)

func init() {

	//normalChannel()
	//bufferedChannel()
	//timer()

}

func timer() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(" unlimite loop")
	}
}

func normalChannel() {
	ch := make(chan int)
	go goNormalChanel(ch)

	time.Sleep(5 * time.Second)

	for value := range ch {
		fmt.Println(value)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("chanel closed")
}

func goNormalChanel(ch chan int) {
	var counter int
	for i := 0; i < 10; i++ {
		counter++
		ch <- counter
		//time.Sleep(1 * time.Second)
	}
	close(ch)
}

func bufferedChannel() {
	buffChan := make(chan int, 10)
	go useBufferedChannel(buffChan)
	for value := range buffChan {
		fmt.Println(value)

	}
	time.Sleep(5 * time.Second)
	for value := range buffChan {
		fmt.Println(value)

	}
	fmt.Println("end of buffered chanell")
}

func useBufferedChannel(buffChan chan int) {
	for i := 0; i < 10; i++ {
		buffChan <- i
		time.Sleep(1 * time.Second)
	}
	close(buffChan)
}
