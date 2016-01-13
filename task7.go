package main

import "fmt"

func factorial(x int64) int64 {

	if x == 0 {
//		fmt.Println(x==0)
//		fmt.Println(x)
		return 1
	}

	return x * factorial(x - 1)
}

func main(){
	fmt.Println(factorial(30))
	fmt.Println(2*3*4*5*6*7*8*9*10*11*12*13*14*15*16*17*18*19*20)
}