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

func InsertNodeAtIthIndex(head *Node, index, data int) *Node{
	if head == nil{
		head = NewNode(data, nil)
		return head
	}
	if index == 0{
		newNode := NewNode(data, nil)
		newNode.next = head
		head = newNode
		return head
	}
	i := 0
	temp := head
	preNode := temp
	for temp != nil {
		if i == index{
			newNode := NewNode(data, nil)
			preNode.next = newNode
			newNode.next = temp
			break
		}
		i++
		preNode = temp
		temp = temp.next
	}
	return head
}

func main(){
	head := NewNode(30, NewNode(10, NewNode(40, NewNode(40, nil))))
	fmt.Printf("Input Linked list is: ")
	TraverseLinkedList(head)

	index := 3
	head = InsertNodeAtIthIndex(head, index, 35)
	fmt.Printf("Inserting new node at %dth index, Linked List is: ", index)
	TraverseLinkedList(head)
}