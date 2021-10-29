package main

import "fmt"

func init() {

	//callBiDirectionalChannel()
	// call uni directional chanel and try read from that and get error
	//callUniDirectionalChan()
	usageOfUnidirectionalChan()
}

func callBiDirectionalChannel() {
	ch := make(chan bool)
	go declearBiDirectionalChanel(ch)
	retult := <-ch
	fmt.Println(retult)
}

func declearBiDirectionalChanel(result chan bool) {
	result <- true
}

// above code we write a bidirectionalchanell that can resive and write to it
// but blow we whant whrite a chanel that only can write to it

func callUniDirectionalChan() {
	ch := make(chan<- bool)
	uniDirectionalChannel(ch)
	// blow line will cause error because ch is send only chanel and cand read from a send only chanel
	//result := <-ch
}

func uniDirectionalChannel(result chan<- bool) {
	result <- true
}

// so what is the usage of unidirectinalchanel if we can write to a sent only channel but cant read from it
// or we can read from unidirectionalchannel and cant write to it
// this is where this is where conversion chanel show itself : we can declear a bidirectional chanel and sent it as
// unidirectional package in this way we have a bidirectional outside  and unidirectional inside local func
//
func usageOfUnidirectionalChan() {
	ch := make(chan int)
	go sendOnlyChanel(ch)
	// ch is bidirectional here
	fmt.Println(<-ch)
}

func sendOnlyChanel(ch chan<- int) {
	// ch is unidirectinal here
	ch <- 2
}
