package main

import (
	"fmt"
	"unicode/utf8"
)

func init() {
	//workingOnStrings()
	//getRightUtftforSpecialChard()
	//createStringFromSliceOfByte()
	//funchowToChangeACharInString()
}
func workingOnStrings() {
	name := "hello world"
	name = name[:5]
	fmt.Printf("%x", name[0])
	isStringRefrenceBase(name)
	fmt.Println(name)
}

func isStringRefrenceBase(s string) {
	s = s[5:]
}

func getRightUtftforSpecialChard() {
	name := "Señor"
	// normal print char for Señor
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i])
	}
	fmt.Print(" \n")
	// wrong format will correcr by range
	for _, c := range name {
		fmt.Printf("%c", c)
	}
	fmt.Print(" \n")
	// or we cat correct it with rune
	runes := []rune(name)
	// normal print char for Señor
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
}

func createStringFromSliceOfByte() {
	byts := []byte{67, 97, 102, 195, 169}
	s := string(byts)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

// strings is immutable and can't be change but we can use rune
func funchowToChangeACharInString() {
	name := "hello"
	//name[0] = 'p' error
	runeFormat := []rune(name)
	runeFormat[0] = 'p'
	fmt.Println(runeFormat)
	s := string(runeFormat)
	fmt.Println(s)
}
