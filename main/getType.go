package main

import "fmt"

func init() {
	//getType(2)
}

func getType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case float32:
		fmt.Println("type is float32")
	case student:
		fmt.Println("type is student struct")
	case newAddress:
		fmt.Println("type is newAddress struct")

	}
}
