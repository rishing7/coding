package main

import "fmt"

func NextPowOf2(n int) int{
	k := 1
	for ;k < n; {
		k = k << 1
	}
	return k
}

func main(){
	fmt.Printf("Round of highest power of 2 for %d is %d\n", 20, NextPowOf2(20))
	fmt.Printf("Round of highest power of 2 for %d is %d\n", 16, NextPowOf2(16))
	fmt.Printf("Round of highest power of 2 for %d is %d\n", 131, NextPowOf2(131))
}
