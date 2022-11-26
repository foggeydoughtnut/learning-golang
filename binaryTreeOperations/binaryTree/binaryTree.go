package binarytree

import (
	"fmt"
)

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

type Tree struct {
	root *Node
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

func (tree *Tree) InOrderTraversal() {
	tree.InOrderTraversalRec(tree.root)
}

func (tree *Tree) InOrderTraversalRec(root *Node) {
	if root == nil {
		return
	}
	tree.InOrderTraversalRec(root.left)
	fmt.Printf("%d ", root.data)
	tree.InOrderTraversalRec(root.right)
}

func (tree *Tree) Size() int {
	return tree.sizeRec(tree.root)
}

func (tree *Tree) sizeRec(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + tree.sizeRec(root.left) + tree.sizeRec(root.right)

}

func (tree *Tree) Search(value int) bool {
	return tree.searchRec(tree.root, value)
}

func (tree *Tree) searchRec(root *Node, valueToLookFor int) bool {
	if root == nil {
		return false
	}
	if root.data == valueToLookFor {
		return true
	}
	if root.data < valueToLookFor {
		return tree.searchRec(root.right, valueToLookFor)
	} else {
		return tree.searchRec(root.left, valueToLookFor)
	}
}

// Inserts node into the tree, if successful return true
func (tree *Tree) Insert(value int) {
	tree.insertRec(tree.root, value)
}

func (tree *Tree) insertRec(node *Node, value int) *Node {
	if tree.root == nil {
		tree.root = &Node{data: value}
		return tree.root
	}
	if node == nil {
		return &Node{data: value}
	}
	if value <= node.data {
		node.left = tree.insertRec(node.left, value)
	}
	if value > node.data {
		node.right = tree.insertRec(node.right, value)
	}
	return node
}

func (tree *Tree) getLeftMostNode(root *Node) *Node {
	currentNode := root
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (tree *Tree) Delete(value int) *Node {
	return tree.deleteRec(tree.root, value)
}

func (tree *Tree) deleteRec(root *Node, value int) *Node {
	if root == nil {
		return root
	}
	if value < root.data {
		root.left = tree.deleteRec(root.left, value)
	} else if value > root.data {
		root.right = tree.deleteRec(root.right, value)
	} else {
		if root.right == nil {
			temp := root.left
			root = nil
			return temp
		} else if root.left == nil {
			temp := root.right
			root = nil
			return temp
		}

		temp := tree.getLeftMostNode(root.right)
		root.data = temp.data
		root.right = tree.deleteRec(root.right, temp.data)
	}
	return root
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (tree *Tree) Height() int {
	return tree.getHeight(tree.root)
}

func (tree *Tree) getHeight(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(tree.getHeight(root.left), tree.getHeight(root.right))
}
