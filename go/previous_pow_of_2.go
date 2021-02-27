package main

import "fmt"

func PreviousPowOf2(n int) int{
	n = n | (n >> 1)
	n = n | (n >> 2)
	n = n | (n >> 4)
	n = n | (n >> 8)
	n = n | (n >> 16)
	return n - (n >> 1)
}

func main(){
	fmt.Printf("Previous round number, power of 2 for %d is %d\n", 20, PreviousPowOf2(20))
	fmt.Printf("Previous round number, power of 2 for %d is %d\n", 12, PreviousPowOf2(12))
	fmt.Printf("Previous round number, power of 2 for %d is %d\n", 131, PreviousPowOf2(131))
}
