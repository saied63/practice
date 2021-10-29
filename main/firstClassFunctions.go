package main

import (
	"errors"
	"fmt"
	"strconv"
)

func AnymousFunc(s string) {
	// this is anymous function
	fn := func() error {
		if len(s) > 0 {
			fmt.Println(s)
			return nil
		}
		return errors.New("entry is null")
	}

	err := fn()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func AnymousFuncWithoutAssyncToVariable(s string) {
	// another type or error execute by ()
	func() {
		if len(s) <= 0 {
			fmt.Println("error : s is null")
		}
	}()
}

func AnymousFuncWithParameter() {
	// we pass params of anymous func inside ()
	var c int = 1
	func(a int, b int) {
		fmt.Println(a / b / c)
	}(6, 3)
}

// user define anymous function type
// by using this type of func signiture we can define any number of body for it . it is like a polymorfism in interface
type add func(a int, b int) (int, error)

func UseUserDefineFunctionType() {
	var fn add = func(a int, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("secound param can't be null")
		}
		return a / b, nil
	}
	result, err := fn(6, 3)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return
	}
	fmt.Println("result : ", result)
}

func UseUserDefineFunctionType_Sum() {
	var fn add = func(a int, b int) (int, error) {
		if a == 0 && b == 0 {
			return 0, errors.New("both integer is zero so result is zero")
		}
		return a + b, nil
	}

	result, err := fn(2, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

////////////////////////////////////////////

// High order function
func UseHighOrderFunc() {
	//HighOrderFanc(ParamFunc2)
	// we can also send anymous func to higher order function
	HighOrderFanc(func(s string) error {
		fmt.Println("using anymous func in higher - order function  ", s)
		return nil
	})
}

func ParamFunc(s string) error {
	fmt.Println(s)
	return nil
}

func ParamFunc2(s string) error {
	if len(s) <= 0 {
		return errors.New("s is empty")
	}
	return nil
}

func HighOrderFanc(fn func(s string) error) {
	err := fn("saied")
	if err != nil {
		fmt.Println(err)
	}
}

// there is two kind of higer-order function one of them is whose accept a function as param like above
// and second one are whose return a function from other function

func ReturnHigherOrderFunction(r string) func() (int, error) {

	fn := func() (int, error) {
		val, err := strconv.Atoi(r)
		if err != nil {
			return -1000, err
		}
		return val, nil
	}
	return fn
}
