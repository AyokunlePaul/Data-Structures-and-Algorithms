package tree

import "strings"

type BinarySearchTree struct {
	Root *Node
}

func NewBST(value int) *BinarySearchTree {
	return &BinarySearchTree{
		Root: NewNode(value),
	}
}

func (bst *BinarySearchTree) insert(value int) {
	bst.Root.insert(value)
}

func (bst *BinarySearchTree) traverseInOrder() {
	if bst.Root != nil {
		bst.Root.traverseInOrder()
	}
}

func (bst *BinarySearchTree) String() string {
	builder := strings.Builder{}
	return builder.String()
}
