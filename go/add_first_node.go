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

func AddFirstNode(head *Node, data int) *Node{

	if head == nil{
		head = NewNode(data, nil)
		return head
	}
	newNode := NewNode(data, nil)
	newNode.next = head
	return newNode
}
func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	head = AddFirstNode(head, 5)
	fmt.Printf("After adding first node, linked list is: ")
	TraverseLinkedList(head)
}
