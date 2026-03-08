package main

import "fmt"

type Node struct { // Nó para si mesmo na struct
	value int
	left * Node
	right * Node
}

func insert() *Node{

	var value int

	var valueLeft int

	var valueRight int

	var example Node
	
	fmt.Println("Insira o valor")
	fmt.Scanln(&value)

	example.value = value

	fmt.Println("Insira o valor do no esquerdo:")
	fmt.Scanln(&valueLeft)

	example.left = &Node{value:valueLeft}

	fmt.Println("Insira o valor do no direito:")
	fmt.Scanln(&valueRight)

	example.right = &Node{value:valueRight}

	
	return &example


}

func print(node *Node, n int){

	if(node == nil){
		return
	}

	fmt.Println("Árvore: ",n,"Valor: ",node.value," \n")

	if(node.left != nil){
		fmt.Println("Nó esquerdo: \n")
	}
		
	print(node.left,n)

	if(node.right != nil){
		fmt.Println("Nó direito: \n")
	}
		
	print(node.right,n)

}

func main(){

	var numberNodes int

	slice := []*Node{}

	fmt.Println("Insira quantos nós terá essa a árvore binária: \n")
	fmt.Scanln(&numberNodes)

	for i:=0; i < numberNodes; i++ {

		node:= insert()

		slice = append(slice,node)

		print(slice[i],i)

	}
	


}