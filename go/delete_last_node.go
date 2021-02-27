package main

import (
	"fmt"
)

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

func TraverseLinkedList(head *Node){
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func DeleteLastNode(head *Node) *Node{

	if head == nil{
		return head
	}
	temp := head
	for temp.next.next != nil{
		temp = temp.next
	}
	temp.next = nil
	return head
}
func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	head = DeleteLastNode(head)
	fmt.Printf("After deleting last node of the linked list: ")
	TraverseLinkedList(head)
}
