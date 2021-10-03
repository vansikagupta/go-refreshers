package main

import (
	"fmt"
	"strings"
)

func main() {

	//stringBasics()
	packageString()

}

func stringBasics() {
	//String literals
	mystr1 := `Raw string format.
	Doesn't support escape character \n.
	Allows multi-line string literal.`
	fmt.Println(mystr1)

	mystr2 := "Interpreted string literal. Supports escape \ncharacters"
	fmt.Println(mystr2)

	//****************

	//Ranging through string
	ascii := "ou"
	// len returns number of bytes
	fmt.Println(len(ascii))
	for i := 0; i < len(ascii); i++ {
		fmt.Print(ascii[i], " ")
	}

	//character represented in more than one byte, rune
	fmt.Println("")
	nonascii := "õµ"
	// len not represents number of character
	fmt.Println(len(nonascii))
	for i := 0; i < len(nonascii); i++ {
		fmt.Print(nonascii[i], " ")
	}

	// range returns rune
	fmt.Println("")
	for i, r := range nonascii {
		fmt.Println(i, nonascii[i], r, string(r))
	}

	//****************

	//character comparision
	a := 'a'
	b := 'b'
	if a == b {
		fmt.Println("Comparision successful")
	} else {
		fmt.Println(b - a)
	}
}

func packageString() {
	splitString := "1, 2, 3, 4, 5"

	substringSlice := strings.Split(splitString, ", ")
	fmt.Println(substringSlice)

	newString := strings.Join(substringSlice, "|")
	fmt.Println(newString)
}
