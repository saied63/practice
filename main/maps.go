package main

import (
	"fmt"
)

func init() {
	//makeAMap()
	//initDuringDecleation()
	//zeroValueMap()
	//noExistingEleman()
	//deleteFromMap()
	//mapOfStrucs()
}

func makeAMap() {
	employees := make(map[string]int)
	employees["saied"] = 10
	employees["ali"] = 18
	fmt.Println(employees)
	fmt.Println(employees["ali"])
	for key, value := range employees {
		fmt.Printf("  %s  :  %d  ", key, value)
	}
}

func initDuringDecleation() {
	myMap := map[string]int{
		"saied": 18,
		"ali":   19,
	}
	myMap["soghra"] = 20
	for key, value := range myMap {
		fmt.Printf(" myMap key %s and value %d ", key, value)
	}
}

func zeroValueMap() {
	myMap := map[string]int{}
	//var myMap map[string]int
	myMap["saied"] = 2
	fmt.Println(myMap)

}

func noExistingEleman() {
	myMap := map[string]int{}
	myMap["saied"] = 12
	value, ok := myMap["ali"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("eleman dose not exist")
	}
}

func deleteFromMap() {
	myMap := map[string]int{}
	myMap["saied"] = 20
	delete(myMap, "ali")
	fmt.Println(myMap)
}

type employee struct {
	salary  int
	age     int
	country string
}

func mapOfStrucs() {
	myMap := map[string]employee{}
	emp1 := employee{
		salary:  1000,
		age:     20,
		country: "iran",
	}
	emp2 := employee{
		salary:  15000,
		age:     31,
		country: "india",
	}
	myMap["saied"] = emp1
	myMap["ali"] = emp2
	for key, value := range myMap {
		fmt.Println(key)
		fmt.Println((value.age))
	}
	fmt.Println(len(myMap))

}
