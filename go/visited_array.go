package main

import (
	"fmt"
)


func FindDuplicateElement(arr []int)  {
	var Visited []VisitedArray
	for i:=0; i<len(arr); i++{
		if Visited[arr[i]] == true{

		}
		Visited[arr[i]] = true
	}
}

func main(){
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)

	for i:=0; i<size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}
}