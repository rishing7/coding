package main

import "fmt"

func FindMax(x, y int){
	fmt.Printf("Maximum element in %d, and %d is: %d\n", x, y, x - ((x - y) & ((x - y) >> 31)))
}

func FindMin(x, y int) {
	fmt.Printf("Minimum element in %d, and %d is: %d\n", x, y, y + ((x - y) & ((x - y) >> 31)))
}

func main(){
	FindMax(12, 15)
	FindMin(13, 17)
	FindMax(1, 0)
}
