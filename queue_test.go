//@author Devansh gupta

package dsal

import (
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	x := NewBasicQueue(10)
	for i := 10; i < 21; i++ {
		err := x.Enqueue(i)

		if err != nil {
			t.Error(err)

			return
		}
	}

}
func TestQueueDequeue(t *testing.T) {
	x := NewBasicQueue(10)
	for i := 10; i < 20; i++ {
		err := x.Enqueue(i)
		if err != nil {
			t.Error(err)
		}
	}
	for !x.Empty() {

		v, err := x.Dequeue()
		t.Log(v, "\n")
		if err != nil {
			t.Error(err)
		}
	}
}
