package tree

import "fmt"

type Node struct {
	value int
	Right *Node
	Left  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		Right: nil,
		Left:  nil,
	}
}

func (node *Node) insert(value int) {
	if value > node.value {
		if node.Right == nil {
			node.Right = NewNode(value)
		} else {
			node.Right.insert(value)
		}
	} else if value < node.value {
		if node.Left == nil {
			node.Left = NewNode(value)
		} else {
			node.Left.insert(value)
		}
	}
}

func (node *Node) traverseInOrder() {
	if node.Left != nil {
		node.Left.traverseInOrder()
	}
	fmt.Printf("Data = %d, ", node.value)
	if node.Right != nil {
		node.Right.traverseInOrder()
	}
}
