package dsal

import (
	"log"
	"testing"
)

func TestGarbage(t *testing.T) {
	x := make([]int8, 0, 5)
	t.Log(x[3])
}

func TestSimpleBitarray(t *testing.T) {
	x := NewBitArray(30)
	var a = []uint8{1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1}
	for _, v := range a {
		x.Append(v == 1)

	}
	var i int64
	var b []uint8
	for i = 0; i < x.size; i++ {
		if x.IsSet(i) {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}
	t.Log(b)
	t.Log(a)
}

func TestAndOrXor(t *testing.T) {
	x, y := NewBitArray(30), NewBitArray(30)
	var a = []uint8{1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1}
	var b = []uint8{0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1}
	for i := 0; i < len(b); i++ {
		x.Append(a[i] == 1)
		y.Append(b[i] == 1)
	}
	and := x.AND(*y)
	or := x.OR(*y)
	xor := x.XOR(*y)
	var oand, oor, oxor []uint8
	var i int64
	for i = 0; i < x.size; i++ {
		oand = append(oand, ch(and.IsSet(i)))
		oor = append(oor, ch(or.IsSet(i)))
		oxor = append(oxor, ch(xor.IsSet(i)))

	}
	log.Print(a)
	log.Print(b)
	log.Print(oand)
	log.Print(oor)
	log.Print(oxor)

}

func TestSetUnset(t *testing.T) {
	var a = []uint8{1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1}
	x := NewBitArray(30)
	for _, v := range a {
		x.Append(v == 1)
	}
	t.Log(toArr(*x))
	x.Set(2) //setting bit '2'
	t.Log(toArr(*x))
	x.Unset(4)
	t.Log(toArr(*x))
}

func toArr(x BitArray) []uint8 {
	var i int64
	var ar []uint8
	for i = 0; i < x.size; i++ {
		ar = append(ar, ch(x.IsSet(i)))
	}
	return ar
}

func ch(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
