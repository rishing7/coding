package main

import (
	"fmt"
	"math"
)

func CheckPowerOf4(n int){
	res := math.Log(float64(n)) / math.Log(float64(4))
	if res == math.Floor(res) {
		fmt.Printf("%d is the power of 4.\n", n)
	} else {
		fmt.Printf("%d is not the power of 4.\n", n)
	}
}

func main(){
	CheckPowerOf4(13)
	CheckPowerOf4(16)
	CheckPowerOf4(0)
}
