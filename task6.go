package main

import "fmt"

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) //
	fmt.Println(nextOdd()) //
	fmt.Println(nextOdd()) //

}

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		if i == 0 {
			i++
		}else {
			i += 2
		}
		ret = i
		return
	}
}