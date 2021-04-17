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
	fmt.Printf("%d->", node.value)
	if node.Right != nil {
		node.Right.traverseInOrder()
	}
}

func (node *Node) traversePreOrder() {
	fmt.Printf("%d->", node.value)
	if node.Left != nil {
		node.Left.traversePreOrder()
	}
	if node.Right != nil {
		node.Right.traversePreOrder()
	}
}

func (node *Node) traversePostOrder() {
	if node.Left != nil {
		node.Left.traversePostOrder()
	}
	if node.Right != nil {
		node.Right.traversePostOrder()
	}
	fmt.Printf("%d->", node.value)
}

func (node *Node) get(value int) *Node {
	if node.value == value {
		return node
	} else if value < node.value {
		if node.Left != nil {
			node.Left.get(value)
		}
	} else {
		if node.Right != nil {
			node.Right.get(value)
		}
	}
	return nil
}

func (node *Node) String() string {
	return string(rune(node.value))
}
