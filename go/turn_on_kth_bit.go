package main

import (
	"fmt"
	"strconv"
)

func TurnOnKthBit(n, k int) int {
	return n | (1 << (k-1))
}

func main(){
	var n = 20
	var k = 4
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	newNumber := TurnOnKthBit(20, 4)
	fmt.Printf("After turning on %d th bit of %d is: %d.\n", k, n, newNumber)
	fmt.Printf("Binary of %d is: %s.\n", newNumber, strconv.FormatInt(int64(newNumber), 2))
}