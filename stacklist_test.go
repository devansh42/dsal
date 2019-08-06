package dsal

import "testing"

func TestStackList(t *testing.T) {
	s := NewStackList()
	for x := 0; x < 5; x++ {
		s.Push(x)
	}
	for s.Length() > 0 {
		t.Log(s.Pop())
	}
}
