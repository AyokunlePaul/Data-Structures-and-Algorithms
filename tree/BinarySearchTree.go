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

func (bst *BinarySearchTree) Insert(value int) {
	bst.Root.insert(value)
}

func (bst *BinarySearchTree) Get(value int) *Node {
	if bst.Root != nil {
		return bst.Root.get(value)
	}
	return nil
}

func (bst *BinarySearchTree) TraverseInOrder() {
	if bst.Root != nil {
		bst.Root.traverseInOrder()
	}
}

func (bst *BinarySearchTree) TraversPreOrder() {
	if bst.Root != nil {
		bst.Root.traversePreOrder()
	}
}

func (bst *BinarySearchTree) TraversePostOrder() {
	if bst.Root != nil {
		bst.Root.traversePostOrder()
	}
}

func (bst *BinarySearchTree) String() string {
	builder := strings.Builder{}
	return builder.String()
}
