package main

import "fmt"

func main(){
	fmt.Printf("Printing Lowercase characters from Uppercase characters using OR operator: ")
	for ch:='A'; ch<='Z'; ch++ {
		fmt.Printf(string(ch | ' '))
	}
	fmt.Println()
	fmt.Printf("Printing Uppercase characters from Lowercase characters using AND operator: ")
	for ch:='a'; ch<='z'; ch++ {
		fmt.Printf(string(ch & '_'))
	}
}