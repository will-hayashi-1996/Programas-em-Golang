package main

import "fmt"

func bubblesort(a []int, n int) {

	if n < 1 {
		return
	}

	var aux int = 0

	for i := 0; i < n; i++ {

		if a[i] > a[i+1] {
			aux = a[i]
			a[i] = a[i+1]
			a[i+1] = aux
		}

	}

	bubblesort(a, n-1)

}

func main() {

	numeros := []int{41, 14, 24, 2, 47, 16, 3, 7, 21, 34}

	bubblesort(numeros, len(numeros)-1)

	fmt.Print(numeros)

}
