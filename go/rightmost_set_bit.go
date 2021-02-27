package main

import (
	"fmt"
	"math"
	"strconv"
)

func FindRightMostSetBit(n int) int {
	if (n & 1) != 0{
		return 1
	}
	return int(math.Log2(float64(n & -n))) + 1
}

func main(){
	var n = 20
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	fmt.Printf("Position of the rightmost set bit of the %d is %d.\n", n, FindRightMostSetBit(n))
}