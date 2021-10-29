package main

import (
	"log"
	"net/http"
	"sync"
)

func init() {
	//startWorkers()
}

func workerGetHttp(in <-chan string, wg *sync.WaitGroup) {
	for {
		url := <-in
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("ststus code of %s is %d ", url, res.StatusCode)
		}
		wg.Done()
	}
}

func startWorkers() {
	var wg sync.WaitGroup
	ch := make(chan string, 1024)
	countOfWorker := 100
	for i := 0; i < countOfWorker; i++ {
		go workerGetHttp(ch, &wg)
	}

	urls := []string{"https://www.google.com", "https://www.bing.com", "https://www.farsnews.ir"}
	for i := 0; i < 100; i++ {
		for i := 0; i < len(urls); i++ {
			wg.Add(1)
			ch <- urls[i]
		}
	}

	wg.Wait()
}
