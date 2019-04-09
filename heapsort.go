//@author Devansh Gupta
//This is the heap sort algorithm implementation in golang
package dsal

//HeapSort, This sorts given slice using heap sort algorithm
//ar=> This is the slice to be sorted
//increasing=> This specify if we want to sort in increasing or decreasing order
//Incase of increasing order we build a max heap else min heap
func HeapSort(ar []int, increasing bool) []int {
	x := NewHeap(ar, increasing)
	x.sort()
	return x.ar
}
