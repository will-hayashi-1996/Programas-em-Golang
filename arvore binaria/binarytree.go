package main

import "fmt"

type Node struct { // Nó para si mesmo na struct
	value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {

	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {

		root.left = insert(root.left, value)

	} else {
		root.right = insert(root.right, value)
	}

	return root

}

func inOrder(node *Node) {

	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.value)
	inOrder(node.right)

}

func main() {

	var numberNodes int

	var root *Node

	var value int

	fmt.Println("Insira quantos nós terá essa a árvore binária: ")
	fmt.Scanln(&numberNodes)

	for i := 0; i < numberNodes; i++ {

		fmt.Println("Insira o valor: ")
		fmt.Scanln(&value)
		root = insert(root, value)
	}

	inOrder(root)

}
