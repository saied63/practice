package main

import (
	"fmt"
	"time"
)

// we can use extra parameter when recive from chanel that specify chanel is closed or open
func init() {
	//testCloseChannel()
}

func testCloseChannel() {
	ch := make(chan int)
	go reciver(ch)
	/* {
		value, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(value)
	}
	*/
	// we can use forrange and no need to check ok in if (range will do it for us till chanel closed)
	for value := range ch {
		fmt.Println(value)
	}
}

func reciver(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}

	close(ch)
}
