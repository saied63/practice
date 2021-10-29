package main

import (
	"fmt"
	"sync"
	"time"
)

func callWorkGroup() {
	fmt.Println("start")
	message("Toplearn.com")
	fmt.Println("End")
}

func message(text string) {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printMessage(text, i, &wg)
	}
	wg.Wait()
}

func printMessage(text string, index int, wg *sync.WaitGroup) {
	if index == 2 {
		time.Sleep(2 * time.Second)
	}
	fmt.Println(text, index)
	wg.Done()
}
