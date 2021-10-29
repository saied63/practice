package main

import "fmt"

var allAddresses []newAddress = []newAddress{{streetName: "amiriye", streetNo: 5, pelak: 2},
	{"aghdasiye", 5, 9}, {"emamiye", 14, 5}}

var allcityAddresses []cityAddress = []cityAddress{{cityName: "mashhad", streetName: "amiriye", streetNo: 5, pelak: 2},
	{"tehran", "aghdasiye", 5, 9}, {"esfehan", "emamiye", 14, 5}}

func init() {
	//useInterface()
	//useInterfaceOnAnotherType()
	//useEmptyInterfaceFunc()
	//assert()
}

type Iaddress interface {
	FindStreetNoName() bool
	ChangeStreetNo(no int)
}

type newAddress struct {
	streetName string
	streetNo   int
	pelak      int
}

type cityAddress struct {
	cityName   string
	streetName string
	streetNo   int
	pelak      int
}

func (e newAddress) FindStreetNoName() bool {

	for _, a := range allAddresses {
		if a == e {
			fmt.Println("address fount  : ")
			fmt.Println(a)
			return true
		}
	}
	fmt.Println("address not fount ")
	return false
}

func (e newAddress) ChangeStreetNo(no int) {
	for i, a := range allAddresses {
		if a.streetName == e.streetName {
			fmt.Println("address fount and changed to : ")
			allAddresses[i].streetNo = no

		}
	}
}

func useInterface() {
	add := newAddress{"amiriye", 5, 2}
	add.FindStreetNoName()
	add2 := newAddress{"amiriye", 8, 6}
	add2.ChangeStreetNo(32)
	fmt.Println(allAddresses)

	var addressInterface Iaddress = add
	getInterface(addressInterface)
	//addressInterface.ChangeStreetNo(42)
	fmt.Println(allAddresses)
}

func (e cityAddress) FindStreetNoName() bool {

	for _, a := range allcityAddresses {
		if a == e {
			fmt.Println("address fount  : ")
			fmt.Println(a)
			return true
		}
	}
	fmt.Println("address not fount ")
	return false
}

func (e cityAddress) ChangeStreetNo(no int) {
	for i, a := range allcityAddresses {
		if a.streetName == e.streetName {
			fmt.Println("address fount and changed to : ")
			allcityAddresses[i].streetNo = no
			fmt.Println(allcityAddresses[i])
		}
	}
}

func useInterfaceOnAnotherType() {
	cadd1 := allcityAddresses[1]
	var cityInterface Iaddress = cadd1

	getInterface(cityInterface)
}

func getInterface(adressInterface Iaddress) {
	adressInterface.ChangeStreetNo(15)
}

func useEmptyInterfaceFunc() {
	geemptyInterface(2)
	geemptyInterface("3")
	add := newAddress{
		streetName: "molavi",
		streetNo:   50,
		pelak:      10,
	}
	geemptyInterface(add)
}

func geemptyInterface(i interface{}) {
	fmt.Printf("type is : %T and value is %v : ", i, i)
}

func assert() {
	getUnderlyingValueOfInterface("3")
	getUnderlyingValueOfInterface(2)
}

func getUnderlyingValueOfInterface(i interface{}) {
	value, ok := i.(int)
	if ok {
		fmt.Printf("type is ok value is : %v", value)
	} else {
		fmt.Printf("get wrong type from interface  value is default : %v", value)
	}
}
