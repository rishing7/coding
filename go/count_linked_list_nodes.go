package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func NewNode(value int, next *Node) *Node{
	var n Node
	n.value = value
	n.next = next
	return &n
}

func CountNodes(head *Node){

	fmt.Printf("Input Linked List is: ")
	count :=0
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
		count += 1
	}
	fmt.Printf("\nNumber of nodes in the linked list is: %d\n", count)
}

func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	CountNodes(head)
}
