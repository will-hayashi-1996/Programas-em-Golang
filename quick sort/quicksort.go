package main

import "fmt"

func quicksort(a *[]int, ini int, fim int) {

	if ini > fim {
		return
	}

	var partPos int = partition(a, ini, fim)

	quicksort(a, ini, partPos-1)

	quicksort(a, partPos+1, fim)

}

func partition(a *[]int, ini int, fim int) int {

	var pivo int = (*a)[fim]

	var i int = ini - 1

	var aux int = 0

	for j := ini; j < fim; j++ {

		aux = 0

		if (*a)[j] <= pivo {

			i++
			aux = (*a)[i]
			(*a)[i] = (*a)[j]
			(*a)[j] = aux
		}

	}

	aux = (*a)[fim]
	(*a)[fim] = (*a)[i+1]
	(*a)[i+1] = aux

	return i + 1
}

func main() {

	numeros := []int{41, 14, 24, 2, 47, 16, 12, 3, 7, 21, 34}

	quicksort(&numeros, 0, len(numeros)-1)

	fmt.Print(numeros)

}
