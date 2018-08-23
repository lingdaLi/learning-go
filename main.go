package main

import (
	"fmt"
	"strings"
)

func main() {
	// string manipulate
	stringManipulate()
}

func stringManipulate() {
	blockPrint("***string mainipulate start***")

	str1 := "An implicitly typed string"

	// Print str1 value and type
	fmt.Printf("%v:%T\n", str1, str1)

	// Transfer str1 to upper case
	fmt.Println(strings.ToUpper(str1))

	// Make str1 as a title
	fmt.Println(strings.Title(str1))

	// Compare two string
	str2 := strings.ToUpper(str1)
	fmt.Println("Equals?", (str1 == str2))
	fmt.Println("Equals case insensitive?", strings.EqualFold(str1, str2))

	// Check whether a string contains a value
	fmt.Println("Contains?", strings.Contains(str1, "typed"))

	blockPrint("***string manipulate finish***")
}

func blockPrint(str string) {
	fmt.Println()
	fmt.Println(str)
	fmt.Println()
}
