package main

import "fmt"

func main() {
	newnode1 := &node{value: 1}
	newnode2 := &node{value: 10}
	newnode3 := &node{value: 20}

	var bt BinaryTree
	bt.insert(newnode1)
	bt.insert(newnode2)
	bt.insert(newnode3)
	resp := bt.lookup(1)
	fmt.Println(resp)
}

type node struct {
	left  *node
	right *node
	value int
}

type BinaryTree struct {
	root *node
}

func (bt *BinaryTree) lookup(val int) bool {
	if bt.root == nil {
		return false
	}

	currentNode := bt.root
	if currentNode.value == val {
		return true
	}

	for {
		if val > currentNode.value {
			//right

			if currentNode.right == nil {
				return false
			}
			if currentNode.right.value == val {
				return true
			}

			currentNode = currentNode.right

		} else {
			//left
			if currentNode.left == nil {
				return false
			}

			if currentNode.left.value == val {
				return true
			}

			currentNode = currentNode.left

		}
	}
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
