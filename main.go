package main

import (
	"fmt"
	"github.com/AyokunlePaul/Data-Structures-and-Algorithms/tree"
)

func main() {
	myTree := tree.NewBST(25)
	myTree.Insert(20)
	myTree.Insert(15)
	myTree.Insert(27)
	myTree.Insert(30)
	myTree.Insert(29)
	myTree.Insert(26)
	myTree.Insert(22)
	myTree.Insert(32)

	myTree.TraverseInOrder()
	fmt.Println()
	myTree.TraversPreOrder()
	fmt.Println()
	myTree.TraversePostOrder()
}
