package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func (root *Node)PreOrderTraversal(){
	if root !=nil{
		fmt.Printf("%d ", root.data)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}
	return
}

func main(){
	tree := Node{1, &Node{2, &Node{4, nil, nil}, &Node{5, nil, nil}}, &Node{3, &Node{6, nil, nil}, &Node{7, nil, nil}}}

	fmt.Printf("Pre Order Traversal of the given tree is: ")
	tree.PreOrderTraversal()
}