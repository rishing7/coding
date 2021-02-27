package main

import (
	"fmt"
	"strconv"
)

func SwapTwoBits(n, x, y int) int{
	if ((n & (1 << x)) >> x) ^ ((n & (1 << y)) >> y) != 0{
		n ^= 1 << x
		n ^= 1 << y
	}
	return n
}

func main(){
	n, x, y := 31, 2, 6
	fmt.Printf("Binary representation of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	res := SwapTwoBits(n, x, y)
	fmt.Printf("After swapping %d th and %d th bits, binary representation of result is: %s\n", x, y, strconv.FormatInt(int64(res), 2))
}
