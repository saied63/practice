package main

import "fmt"

func init() {
	//testDefer()

}

// defer will be use to call a function or methode in another metihote to return a value or do somthing just before return from that function
// it is posiible we return from function without reach at the end '}' in these cases we use defer for  that things we wanna return when function done it's work
// in any state
// defere can take parameter and parameter value will be declear in exact point ther feeds not at the end : see bloww function

var X int

// in this case when we called XpluOne using defere the value of the X was 1 and we know defere will call just befor retrn from func that invoke it and then in the middle of func we add 1000
// unit to the x and fmt inside func will print 1002 and defer will return 1003
func testDefer() {
	X = 1
	defer XplusOne()

	X += 1000
	XplusOne()
	fmt.Printf("end of func %d \n", X)
}

func XplusOne() {
	X = X + 1
	fmt.Printf("after add on unit %d \n", X)
}
