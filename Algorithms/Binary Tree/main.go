package main

import (
	"fmt"
)

type Node struct {
	data  int
	right *Node
	left  *Node
}

func CreateNewNode(data int) (newNode *Node) {
	newNode = &Node{data, nil, nil}
	return
}
func main() {
	var root *Node
	root = CreateNewNode(229)
	n1 := CreateNewNode(12)
	n2 := CreateNewNode(29)
	root.right = n1
	root.left = n2
	n1n1r := CreateNewNode(200)
	n1.right = n1n1r
	fmt.Println(n1.right)
}

// next https://www.geeksforgeeks.org/binary-tree-set-2-properties/
