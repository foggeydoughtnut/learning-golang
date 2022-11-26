package binarytree

import "fmt"

// Returns true if the item being searched for is in the tree, else false
func BinarySearch(list []int, item int) bool {
	median := len(list) / 2
	if len(list) == 0 {
		return false
	}
	if list[median] == item {
		return true
	}
	if item < list[median] {
		return BinarySearch(list[:median], item)
	} else {
		return BinarySearch(list[median:], item)
	}
}

// Node structure
type Node struct {
	data  int
	left  *Node
	right *Node
}

func CreateBinaryTree() Node {
	root := Node{5, nil, nil}
	root.left = &Node{3, nil, nil}
	root.right = &Node{8, nil, nil}
	root.left.left = &Node{2, nil, nil}
	root.left.right = &Node{4, nil, nil}
	root.right.left = &Node{6, nil, nil}
	root.right.right = &Node{10, nil, nil}
	return root
}

func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Printf("%d ", root.data)
	InOrderTraversal(root.right)
}

func CountNumberOfNodes(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + CountNumberOfNodes(root.left) + CountNumberOfNodes(root.right)
}
