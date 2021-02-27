package main

import (
	"fmt"
	"strconv"
)

func CheckKthBitSet(n, k int) int {
	return n & (1 << (k-1))
}

func main(){
	var n = 20
	var k = 3
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	flag := CheckKthBitSet(n, k)
	if flag !=0 {
		fmt.Printf("%d th bit of %d is aleady set.\n", k, n)
	} else {
		fmt.Printf("%d th bit of %d is not set.\n", k, n)
	}
}