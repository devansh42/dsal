//@author Devansh Gupta
package dsal

import "testing"

func TestQueueList(t *testing.T) {
	l := NewQueueList()
	for x := 70; x < 110; x++ {
		l.Enqueue(x)
	}
	printQueuelist(*l, t)
}

func printQueuelist(l QueueList, t *testing.T) {
	for l.Length() != 0 {
		t.Log(l.Dequeue())
	}
}
