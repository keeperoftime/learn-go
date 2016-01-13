package main

import "fmt"

func main() {
	fmt.Println(add(-100, -2, -6, -8, -9, -90, -87, -66))
}

func add(args...int) int {
	//	xs := []int{1, 2, 3, 4, 57, 878}
	var m int = args[0]
	for i := 0; i < len(args); i++ {
		if m < args[i] {
			m = args[i]
		}

	}
	return m
}


