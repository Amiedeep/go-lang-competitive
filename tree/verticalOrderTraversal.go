package main

import (
	"fmt"
)

//Node is the binary tree node
type Node struct {
	leftChild  *Node
	rightChild *Node
	data       int
}

var (
	hm = map[int][]int{}

	min, max int
)

func verticalOrderTraversal(node *Node) {
	for i := min; i <= max; i++ {
		for j := 0; j < len(hm[i]); j++ {
			fmt.Println(hm[i][j])
		}
	}
}

func initializeHashMap(node *Node, v int) {
	if node == nil {
		return
	}
	if v < min {
		min = v
	}

	if v > max {
		max = v
	}

	//if value of map is of type []int{}, then append works. However, if it is []int then it gives `assignment to entry in nil map` error
	// if hm[v] == nil {
	// 	fmt.Println(hm[v])
	// }

	hm[v] = append(hm[v], node.data)

	initializeHashMap(node.leftChild, v-1)
	initializeHashMap(node.rightChild, v+1)
}

func main() {

	thirdNode := &Node{leftChild: nil, rightChild: nil, data: 4}
	secondNode := &Node{leftChild: thirdNode, rightChild: nil, data: 3}
	firstNode := &Node{leftChild: nil, rightChild: secondNode, data: 2}

	initializeHashMap(firstNode, 0)
	verticalOrderTraversal(firstNode)

}
