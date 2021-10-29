package main

import (
	"errors"
	"fmt"
)

/*
this is the error package and it's New func that we can use to make an custom error
  package errors

  // New returns an error that formats as the given text.
  func New(text string) error {
      return &errorString{text}
  }

  // errorString is a trivial implementation of error.
  type errorString struct {
      s string
  }

  func (e *errorString) Error() string {
      return e.s
  }
*/

func init() {
	//callDevide()
	//callDevide2()
	//checkThisAccount("ali", "ali")
}

// usign error package New func to make a custom error
func devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("devide by zero is denied")
	}
	return (a / b), nil
}

func callDevide() {
	value, err := devide(4, 0)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	fmt.Println(" value : ", value)
}

// using PrintF to return more details because New methode of errors package just get text

func devideWithMoreInfo(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("b parameter (%d) can't be zero", b)
	}
	return a / b, nil
}

func callDevide2() {
	value, err := devideWithMoreInfo(5, 0)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	fmt.Println("result of devide is ", value)
}

// using struct to make custom error
// in this case we use a struct which implement error interface (by implementin it's methode name Error)
// the Error methode will return string contain error text
// then in func that we want to return error we will return a adress of instance of that struct as error interface
//now we can get error format from this function like blowe and also we can compare error using assert type of underlying layer of error (wrongPass ins this case)
type wrongPass struct {
	username string
	err      string
}

func (wup *wrongPass) Error() string {
	return fmt.Sprintf("wrong password for user %s : %s ", wup.username, wup.err)
}

func checkUserAndPass(user string, pass string) error {
	//for example
	if user == pass {
		return &wrongPass{user, "wrong username and password"}
	}
	return nil
}

func checkThisAccount(user string, pass string) {
	err := checkUserAndPass(user, pass)
	if err != nil {
		if err, ok := err.(*wrongPass); ok {
			fmt.Println(err.Error())
		}
	}
}

// another example :
type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func test() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
	}
	fmt.Println("area of rect", area)
}
