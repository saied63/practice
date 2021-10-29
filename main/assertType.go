package main

/*
access to underlying layers of "interface" is called type asserting in go lang
*/

import (
	"fmt"
)

func init() {
	useAssert()
}

type myInterface interface {
	myFunc()
}

type myStruct struct {
	name    string
	address string
}

func (ms myStruct) myFunc() {

}

func useAssert() {
	ms := myStruct{
		name:    "saied",
		address: "amiriye",
	}
	var inf myInterface
	inf = ms
	value, ok := inf.(myStruct)
	fmt.Printf("%t   %s  ", ok, value)

}
