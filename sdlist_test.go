//@author Devansh Gupta
package dsal

import "testing"

func TestSingleList(t *testing.T) {
	b := NewSingleList()
	for x := 10; x < 30; x++ {
		b.Push(x)
	}
	printsList(b, t)
	t.Log(b.Head())
	t.Log(b.Tail())
	t.Log(b.Pop())
	printsList(b, t)
	t.Log(b.Find(89).key)
}
func printsList(b *SingleList, t *testing.T) {
	for x := b.null.next; x != b.null; x = x.next {
		t.Log(x.key)
	}
}
