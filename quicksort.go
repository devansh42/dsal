//@author Devansh Gupta
//This is the Quicksort Implementation in Golang
//This select a random index as a pivot
package dsal

import (
	"fmt"
	"math/rand"
)

//QuickSorted, sorts given index and returns sorted array
func QuickSorted(ar []int, increasing bool) []int {
	x := ar
	quicksort(x, 0, len(ar), increasing)
	return x
}

//QuickSort, sorts given array in place
func QuickSort(ar []int, increasing bool) {
	quicksort(ar, 0, len(ar), increasing)

}

//quicksort, is the internal method for sorting
func quicksort(ar []int, p, q int, increasing bool) {
	if p < q {
		g := rand.Intn(q) //Using a random no. as pivot
		for ; g < p; g = rand.Intn(q) {
		} //g is the random index in [p,q)
		swap(ar, g, q-1) //Swaping pivot from last element in array
		x := partition(ar, p, q, increasing)
		quicksort(ar, p, x, increasing)
		quicksort(ar, x+1, q, increasing)

	}
}

//swap, swaps element at given index
func swap(ar []int, x, y int) {
	t := ar[x]
	ar[x] = ar[y]
	ar[y] = t
}

func partition(ar []int, p, q int, increasing bool) int {
	i := p - 1
	var pivot = ar[q-1] //Last Element as pivot
	fmt.Println(pivot)
	for j := p; j < q-1; j++ {
		if ar[j] <= pivot {
			i++
			swap(ar, i, j)
		}
	}
	swap(ar, i+1, q-1) //Swap last element with first element which is greater than pivot

	return 1 + i
}
