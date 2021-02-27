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

func DeleteKthIndexNode(head *Node, index int) *Node{
	if head == nil{
		return head
	}
	if index == 0{
		head = head.next
		return head
	}
	i := 1
	temp := head
	for temp != nil{
		if i == index{
			temp.next = temp.next.next
		}
		i++
		temp = temp.next
	}
	return head
}

func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)

	index := 5
	head = DeleteKthIndexNode(head, index)
	fmt.Printf("After Deletion of %dth index node, Linked List is: ", index)
	TraverseLinkedList(head)
}