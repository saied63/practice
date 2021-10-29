package main

// if the type of reciiver in implemention of interface declear as value type both value type and pointer type(sent by &a)
// will mean value type and methode can reach the value eighter eith sending value or sending adress of value
// in this way methode work lile normal value type  reciver hence eighter send value or &value will result the same and therir
// change inside function is not visible outside it (the point is if we inplement extention methode with reciver type value
// we the compiler can acceess the value either by sending value or wit it's address)

// but when we implement interface methode with pointer type it can be only reach with pointer reciver not a normal value type
// because in this way compiler can not access type of reciver without &
// in this way th change inside compiler will be visible outside it and will change refrence variable
// but in way one said above compiler will get value of reciver and dont care of refrence value and change will not
// visible otside function either by sending value type of pointer adress

import "fmt"

func init() {
	//checkInterfaceWithPointer()
}

type iDescriber interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
	p.age++
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
	a.state = "iran"
}

func checkInterfaceWithPointer() {
	var d1 iDescriber
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	fmt.Printf("%s is %d years old\n", p1.name, p1.age)
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()
	fmt.Printf("%s is %d years old\n", p2.name, p2.age)
	var d2 iDescriber
	a := Address{"Washington", "USA"}

	/* compilation error if the following line is
	   uncommented
	   cannot use a (type Address) as type iDescriber
	   in assignment: Address does not implement
	   iDescriber (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a //This works since iDescriber interface
	//is implemented by Address pointer in line 22
	d2.Describe()
	fmt.Printf("State %s Country %s", a.state, a.country)
}
