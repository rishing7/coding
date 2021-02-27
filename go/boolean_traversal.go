package main

import "fmt"

func main(){
	arr := []int{10, 20, 30, 60, 40, 50}
	boolArr := []bool{true, false, true, false, true, false}
	fmt.Println("Input Array is: ", arr)
	fmt.Println("Input Boolean Array is: ", boolArr)

	visitedArray := []struct{
		i int
		b bool
	}{
		{10, true},
		{20, false},
		{30, true},
		{60, false},
		{40, true},
		{50, false},
	}
	fmt.Println("Boolean array using struct: ", visitedArray)
}