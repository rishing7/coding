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

func TraverseLinkedList(head *Node){
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func UpdateLastNodeValue(head *Node, data int) *Node{

	if head == nil{
		return head
	}
	temp := head
	for temp.next != nil{
		temp = temp.next
	}
	temp.value = data
	return head
}
func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	head = UpdateLastNodeValue(head, 41)
	fmt.Printf("After updating last node value, linked list is: ")
	TraverseLinkedList(head)
}

