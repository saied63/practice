package main

import "fmt"

func init() {
	//useMultipleInterface()
	//testZeroINterfaceValue()
}

// iChange and iPrint is decleared inside interace multipleInterface
type iEmbededInterface interface {
	iChange
	iPrint
}

type cityInfo struct {
	name string
	no   int
}

func (cs cityInfo) Print() {
	fmt.Printf("cs.name  is %s and cs.streetNo is %d  \n", cs.name, cs.no)
}

func (cs *cityInfo) Change(streetno int) {
	cs.no = streetno
	fmt.Printf("cs.name  is %s and cs.streetNo is %d  \n", cs.name, cs.no)
}

func useMultipleInterface() {
	var embededInterface iEmbededInterface
	cityInfor1 := cityInfo{
		name: "tehran",
		no:   12,
	}
	embededInterface = &cityInfor1
	embededInterface.Change(15)
	embededInterface.Print()
}

// zero or default value of interface in nil
func testZeroINterfaceValue() {

	var embededInterfaceInstance iEmbededInterface
	fmt.Println(embededInterfaceInstance)
}
