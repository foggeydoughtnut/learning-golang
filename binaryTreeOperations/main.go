package main

import (
	binarytree "binaryTreeOpertions/binaryTree"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarytree.BinarySearch(list, 2)
	fmt.Printf("2 is in the tree : %t\n\n", result)

	root := binarytree.CreateBinaryTree()
	fmt.Println("In Order traversal")
	binarytree.InOrderTraversal(&root)
	fmt.Print("\n\n")

	numOfNodes := binarytree.CountNumberOfNodes(&root)
	fmt.Printf("Number of nodes in the tree : %d\n", numOfNodes)
}
