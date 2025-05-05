package main

import (
	"fmt"
)

func main() {
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}

	var bt BinaryTree

	// Insert multiple values
	for _, val := range values {
		bt.insert(&node{value: val})
	}

	list := bt.DFSPreorderSearch()
	fmt.Println(list)
}

type node struct {
	left  *node
	right *node
	value int
}

type BinaryTree struct {
	root *node
}

func (bt *BinaryTree) DFSPreorderSearch() []int {
	return traversePreorder(bt.root, []int{})
}

func (bt *BinaryTree) insert(nn *node) {
	if bt.root == nil {
		bt.root = nn
		return
	}
	currentNode := bt.root
	for {

		if currentNode.value > nn.value {
			//Left

			if currentNode.left == nil {
				currentNode.left = nn
				return

			}
			currentNode = currentNode.left

		} else {

			//right
			if currentNode.right == nil {
				currentNode.right = nn
				return
			}
			currentNode = currentNode.right
		}
	}

}

func traversePreorder(node *node, list []int) []int {
	if node == nil {
		return nil
	}

	list = append(list, node.value)

	if node.left != nil {
		list = traversePreorder(node.left, list)
	}

	if node.right != nil {
		list = traversePreorder(node.right, list)
	}

	return list
}
