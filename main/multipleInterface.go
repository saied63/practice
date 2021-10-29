package main

import "fmt"

func init() {
	//testMultipleInterface()
}

type iPrint interface {
	Print()
}
type iChange interface {
	Change(streetno int)
}

type citystruct struct {
	name     string
	streetNo int
}

func (cs citystruct) Print() {
	fmt.Printf("cs.name  is %s and cs.streetNo is %d  \n", cs.name, cs.streetNo)
}

func (cs *citystruct) Change(streetno int) {
	cs.streetNo = streetno
	fmt.Printf("cs.name  is %s and cs.streetNo is %d  \n", cs.name, cs.streetNo)
}

func testMultipleInterface() {
	var printInterface iPrint
	cs1 := citystruct{
		"mashhad",
		16,
	}
	printInterface = cs1
	printInterface.Print()
	var changeInterface iChange = &cs1
	changeInterface.Change(15)
	fmt.Printf("cs.name  is %s and cs.streetNo is %d  \n", cs1.name, cs1.streetNo)
}
