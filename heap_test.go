//@author Devansh Gupta
package dsal

import "testing"

func TestHeap(t *testing.T) {
	var x = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	h := NewHeap(x, true)
	t.Log(h.GetHeap())
	h.sort()
	t.Log(h.GetHeap())
}
