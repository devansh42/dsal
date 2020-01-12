package dsal

import (
	"fmt"
	"testing"
)

func TestStackPush(t *testing.T) {

	b := NewBasicStack(1 << 10) //1024 elements
	for x := 0; x < 50; x++ {
		err := b.Push(x)
		if err != nil {
			t.Error(err.Error())
		}
	}
}

func TestStackPop(t *testing.T) {
	b := NewBasicStack(1 << 10) //1024 elements
	for x := 78; x < 1000; x++ {
		err := b.Push(x)
		if err != nil {
			t.Error(err.Error())
		}
	}
	for x := 78; x < 100; x++ {
		v, err := b.Pop()
		if err != nil {
			t.Error(err.Error())
			return
		}
		fmt.Println(v)
	}

}
func TestPushAndPop(t *testing.T) {
	x := NewBasicStack(5)
	x.Push(555)
	v, _ := x.Pop()
	t.Log(v)
}
