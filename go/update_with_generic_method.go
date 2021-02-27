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

func UpdateKthIndexNode(head *Node, index , data int) *Node{
	if head == nil{
		return head
	}

	i := 0
	temp := head
	for temp != nil{
		if i == index{
			temp.value = data
			break
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

	//index := 0
	//head = UpdateKthIndexNode(head, index, 15)
	//fmt.Printf("Update %dth index node, Linked List is: ", index)
	//TraverseLinkedList(head)

	//index := 3
	//head = UpdateKthIndexNode(head, index, 15)
	//fmt.Printf("Update %dth index node, Linked List is: ", index)
	//TraverseLinkedList(head)

	//index := 2
	//head = UpdateKthIndexNode(head, index, 15)
	//fmt.Printf("Update %dth index node, Linked List is: ", index)
	//TraverseLinkedList(head)

	index := 10
	head = UpdateKthIndexNode(head, index, 15)
	fmt.Printf("Update %dth index node, Linked List is: ", index)
	TraverseLinkedList(head)
}