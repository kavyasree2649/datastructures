package main

import "fmt"

func main() {
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}

	var bt BinaryTree

	// Insert multiple values
	for _, val := range values {
		bt.insert(&node{value: val})
	}

	list := bt.breathFirstSearch()
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

func (bt *BinaryTree) breathFirstSearch() []int {

	currentNode := bt.root
	var list []int
	var queue []*node

	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		list = append(list, currentNode.value)

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}

	return list

}
