//@author Devansh Gupta
package dsal

import (
	"testing"
)

func TestBasicListPush(t *testing.T) {
	b := NewBasicList()
	for x := 100; x < 106; x++ {
		b.Push(x)
	}
	printList(b, t)
	//b.Remove(103)
	t.Log(b.Find(104))
	t.Log("Printing")
	printList(b, t)
	b.Append(30, 8989)
	b.Prepend(289, 8989978)
	t.Log("Printing")
	printList(b, t)
	b.Remove(1003)
	t.Log("Printing")

	printList(b, t)
}

func printList(b *BasicList, t *testing.T) {
	for x := b.Head(); x != b.Sentinel(); x = *x.next {
		t.Log(x.key)
	}
}
