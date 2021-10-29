package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func anotherKindOfFor() {
	i := 0
	for ; i < 8; i++ {
		fmt.Println("another for" + strconv.FormatInt(int64(i), 10))
	}
	fmt.Println(i)

}

func forStatementWithLable() {
outer:
	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}

		for j := 0; j < 10; j++ {
			fmt.Println("j" + strconv.FormatInt(int64(j), 10))
			if j > 5 && i > 7 {
				break outer
			}
		}

		fmt.Println("i" + strconv.FormatInt(int64(i), 10))
	}
}

func anotherKingFor2() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func anotherKindOfFo3() {
	i := 0
	for {
		if i > 4 {
			break
		}
		i++
	}
	fmt.Println(i)
}

func testSwitch(i int) {

	switch {
	case i < 10:
		if i < 0 {
			fmt.Println("less than zero")
			break
		}
		fmt.Println("less than 10")
	case i >= 10 && i < 20:
		fmt.Println("i is between 10 & 20")
	case i >= 20 && i < 30:
		fmt.Println("i is between 20 than 30")
	default:
		fmt.Println("i is more than 30")

	}
}

func testmultiplecase(i int) {
	switch i {
	case 1, 2, 3, 4:
		fmt.Println("i is less than 5")
	default:
		fmt.Println("i is 5 or more")
	}
}

func testAnotherSwitchCase() {
	switch char := 'c'; char {
	case 'a', 'f', 'h':
		fmt.Println("char is included")
	default:
		fmt.Println("char is not included")
	}
}

func switchInsideFor() int {
	result := 0
breakelable:
	for i := 0; i < 10; i++ {
		switch {
		case i > 4 && i < 6:
			result = 5
			break breakelable
		}
		fmt.Println(i)
	}
	return result
}

func randomNumber() {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(51)
	fmt.Printf("random number is %d", rnd)
	//fmt.Println(strconv.FormatInt(rnd, 10))
}
