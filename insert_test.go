package dsal

import "testing"

func TestInsertionSort(t *testing.T) {
	var x = []int64{2, 6, 3, 4, 10, 7}
	InsertionSort(x)
	t.Logf("%+v", x)
}
