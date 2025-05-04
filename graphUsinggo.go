package main

import "fmt"

type Graph struct {
	adjacentList  map[int][]int
	numberofNodes int
}

func main() {
	var graph Graph

	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)

	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 3)

	graph.PrintGraph()

}

func (list *Graph) addVertex(node int) {
	if list.adjacentList == nil {
		list.adjacentList = make(map[int][]int)
	}
	if _, exist := list.adjacentList[node]; !exist {
		list.adjacentList[node] = []int{}
		list.numberofNodes++
	}

}

func (list *Graph) addEdge(node1, node2 int) {
	list.adjacentList[node1] = append(list.adjacentList[node1], node2)
	list.adjacentList[node2] = append(list.adjacentList[node2], node1)
}

// PrintGraph displays the graph structure
func (g *Graph) PrintGraph() {
	for node, neighbors := range g.adjacentList {
		fmt.Printf("%d: %v\n", node, neighbors)
	}
}
