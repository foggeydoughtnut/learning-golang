package main

import (
	binarytree "binaryTreeOpertions/binaryTree"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarytree.BinarySearch(list, 2)
	fmt.Printf("2 is in the list : %t\n\n", result)

	tree := &binarytree.Tree{}
	valuesToInsert := []int{1, 10, 2, -5, 5, 6, 8, 11, 9}
	for _, val := range valuesToInsert {
		tree.Insert(val)
	}

	fmt.Println("In order traversal")
	tree.InOrderTraversal()
	fmt.Print("\n\n")

	numOfNodes := tree.Size()
	fmt.Printf("Number of nodes in the tree : %d\n\n", numOfNodes)

	searchResult := tree.Search(6)
	fmt.Printf("6 is in the tree : %t\n", searchResult)

	fmt.Println("Deleting 6...")
	tree.Delete(6)
	fmt.Printf("Result from search for 6 is : %t\n", tree.Search(6))
}
