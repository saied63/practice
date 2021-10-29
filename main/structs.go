package main

import "fmt"

func init() {
	//makesomeInstanceOfStruct()
	//makeAnonymousStruct()
	//pointerToStruct()
	//anonymouseFieldInsideStruct()
	//nestedStruct()
	//testPromotedStructField()
}

// name struct
type students struct {
	name, family string
	age          int
}

func makesomeInstanceOfStruct() {
	std1 := students{"saied", "rajab", 2}
	std2 := students{
		name:   "ali",
		family: "rajab",
		age:    30,
	}
	studentMap := make(map[int]students)
	studentMap[0] = std1
	studentMap[1] = std2
	for _, student := range studentMap {
		fmt.Println(student)
	}
}

func makeAnonymousStruct() {
	book := struct {
		name   string
		family string
		age    int
	}{
		name:   "name",
		family: "family",
		age:    30,
	}
	fmt.Println(book)
}

func pointerToStruct() {

	std1 := &students{
		name:   "saied",
		family: "rajab",
		age:    33,
	}
	fmt.Println(std1.age)
}

type anonymousField struct {
	string
	int
}

type account struct {
	accountnumber int
	userandpass   anonymousField
}

func anonymouseFieldInsideStruct() {
	anonymous := anonymousField{
		"saied",
		23,
	}
	fmt.Println(anonymous.string)
	println(anonymous.int)
}

func nestedStruct() {
	nestedSample := account{
		accountnumber: 10,
		userandpass: anonymousField{
			"saied63",
			666666,
		},
	}
	fmt.Println(nestedSample.accountnumber)
	fmt.Println(nestedSample.userandpass.string)
}

type city struct {
	city   string
	street string
}

type country struct {
	name  string
	state string
	city
}

func testPromotedStructField() {
	country1 := country{
		name:  "iran",
		state: "asia",
		city: city{
			city:   "mashhad",
			street: "amiriye",
		},
	}
	fmt.Println(country1.city)
	fmt.Println(country1.city.city)
	fmt.Println(country1.street)
}
