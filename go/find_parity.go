package main

import (
	"fmt"
	"strconv"
)

func FindParity(n int) bool {
	parity := false

	for n != 0 {
		if n & 1 != 0{
			parity = !parity
		}
		n = n >> 1
	}
	return parity
}

func main(){
	n := 20
	fmt.Printf("Binary of %d is: %s.\n", n, strconv.FormatInt(int64(n), 2))
	if FindParity(n){
		fmt.Printf("Parity of the %d is Odd.\n", n)
	} else {
		fmt.Printf("Parity of the %d is Even.\n", n)
	}

}