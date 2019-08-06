package dsal

import "testing"

func TestSet(t *testing.T) {
	x := make(Set)
	y := NewSet()
	x.Insert(1)
	x.Insert(3)
	x.Insert(5)
	x.Insert(2)
	x.Insert(100)
	y.Insert(1)
	y.Insert(100)
	y.Insert(67)

	l := t.Log

	u := x.Union(y)
	inter := x.Intersection(y)
	l(u)
	l(inter)
	l(x.IsMember(100))
	l(x.Remove(100))
	l(x.IsMember(100))
}
