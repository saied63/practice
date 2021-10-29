package main

import "fmt"

func init() {
	//useMethode()
}

type student struct {
	name   string
	family string
	age    int
}

func (e student) extentionMethode() {
	fmt.Println(e.name)
	fmt.Println(e.family)
	fmt.Println(e.age)
}

func useMethode() {
	std1 := student{
		name:   "saied",
		family: "rajab",
		age:    28,
	}
	std1.extentionMethode()
	std1 = std1.addPlayerAge()
	fmt.Println(std1)
	//(&std1).pointerTypeAddAge() go allow us to remove & and just call like this:
	(std1).pointerTypeAddAge()
	fmt.Println(std1)
}

func (e student) addPlayerAge() student {
	e.age++
	return e
}

func (e *student) pointerTypeAddAge() {
	e.age++
}
