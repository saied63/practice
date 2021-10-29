package main

/*
errore in go comes from blow interace
type error interface {
    Error() string
}
*/

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func init() {
	//readfile()
	//getIpAddress()
	testErrorComparing()
}

// in build int error handlingg this interface will return a string which describe an error

func readfile() {
	file, err := os.Open("./txt.txt")
	if err != nil {
		pathError, ok := err.(*os.PathError) // using type assert to find if error have a underlying layer of ok.pathError or not(is error == ok.PathError or not)
		if ok {
			fmt.Println("ok : ", ok)
			fmt.Printf(" \n\n pathError.Op :  %s \n pathError.Path : %s \n ", pathError.Op, pathError.Path)
		}
		//fmt.Println("error ", err)
		return
	}
	fmt.Printf("file.name : %s ", file.Name())
}

// there is alot of error built in struct which implemets error interface , the open command above has a error struct type name PathError
// that is like this
/*
type PathError struct {
    Op   string
    Path string
    Err  error
}
*/
// the above struct  emplement error interface like this
//func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

// another example :
// in this example we try to get underlying layer of dnsError that is contained a timeout() bool , a temporary() bool and Error()

func getIpAddress() {
	ipAddress, err := net.LookupHost("www.google.com")
	if err != nil {
		errorType, ok := err.(*net.DNSError)
		if ok {
			if errorType.Timeout() {
				println("Timeout")
				return
			}
			if errorType.Temporary() {
				fmt.Printf("Temporary error")
				return
			}
			fmt.Println("error ", err)
		}

	}
	for index, value := range ipAddress {
		fmt.Println(index, value)
	}

}

// another way of getting more detailes from error is comparsion errors with defined error like filepath.ErrBadPattern

func testErrorComparing() {
	matched, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("bad pattern error : ", err)
			return
		}
		fmt.Println("error : ", err.Error())
	}
	for _, value := range matched {
		fmt.Println(value)
	}
}
