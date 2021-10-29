package main

import (
	"fmt"
	"sync"
)

func init() {
	//useMutex()
}

type cunter struct {
	sync.Mutex
	value int
}

func useMutex() {
	var wg sync.WaitGroup
	counter := cunter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}
	wg.Wait()
	fmt.Println("value", counter.value)
}

func increment(counter *cunter, wg *sync.WaitGroup) {
	counter.Lock()
	i := counter.value
	counter.value = i + 1
	//counter.value++
	wg.Done()
	counter.Unlock()
}
