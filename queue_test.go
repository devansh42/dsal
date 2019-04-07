//@author Devansh gupta

package dsal

import (
	"testing"
)

func TestQueue(t *testing.T) {
	x := NewBasicQueue(10)
	for i := 10; i < 20; i++ {
		err := x.Enqueue(i)

		if err != nil {
			t.Error(err)

			return
		}
	}

	t.Log(x.array)
	t.Log(x.i)
	t.Log(x.j)
	t.Log(x.Dequeue())
	t.Log(x.array)

}
