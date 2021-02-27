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
func AddAfterKthNode(head *Node, k , data int) *Node{
	// Insert after Kth node(K is not in the linked list).

	if head == nil{
		return head
	}
	temp := head
	for temp != nil{
		if temp.value == k{
			newNode := NewNode(data, nil)
			newNode.next = temp.next
			temp.next = newNode
			break
		}
		temp = temp.next
	}
	return head
}

func DeleteAfterKthNode(head *Node, k int) *Node{
	// Delete after Kth node.

	if head == nil{
		return head
	}
	temp := head
	for temp != nil{
		if temp.value == k{
			temp.next = temp.next.next
		}
		temp = temp.next
	}
	return head
}

func UpdateAfterKthNode(head *Node, k, data int) *Node{
	// Update after Kth node(K is not in the linked list).

	if head == nil{
		return head
	}
	temp := head
	for temp != nil{
		if temp.value == k{
			temp.next.value = data
		}
		temp = temp.next
	}
	return head
}


func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)
	head = AddAfterKthNode(head, 50, 15)
	fmt.Printf("Adding node after %dth value node, Linked List is: ", 50)
	TraverseLinkedList(head)
}
