//@author Devansh Gupta
package dsal

import "testing"

func TestQuickSort(t *testing.T) {
	var x = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	Quicksort(x, true)
	t.Log(x)
}
