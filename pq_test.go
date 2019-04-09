//@author Devansh Gupta
package dsal

import "testing"

func TestPq(t *testing.T) {
	var ar = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	x := NewPriorityQueue(ar, true)
	t.Log(x.GetHeap())
	t.Log(x.GetExtreme())
	t.Log(x.ExtractExtreme())
	t.Log(x.GetHeap())
	t.Log(x.IncreaseKey(2, 89))
	t.Log(x.GetHeap())

}
