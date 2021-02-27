package main

import (
	"fmt"
	"strconv"
)

func CheckNumberPowerOfTwo(n int) int {
	return n & (n-1)
}

func main(){
	var n = 16
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	flag := CheckNumberPowerOfTwo(n)
	if flag == 0{
		fmt.Printf("Given %d number is the power of 2.\n", n)
	} else {
		fmt.Printf("Given %d number is not the power of 2.\n", n)
	}
}