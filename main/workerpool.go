package main

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	//callGroupOfProcess()
}

func callGroupOfProcess() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("all process are done")
}

func process(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Printf("wg process number : %d \n ", i)
	wg.Done()
}
