package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchValue())
}

func searchValue() (int8, bool) {
	fmt.Print("Enter a number: ")
	var input int64
	fmt.Scanf("%d", &input)

	x := input % 2
	fmt.Println(x)
	even := x == 0
	var boolToInt int8
	if even {
		boolToInt = 1
	}
	return boolToInt, even
}
