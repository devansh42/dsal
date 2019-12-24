package dsal

import (
	"math/rand"
	"testing"
	"time"
)

type ft func(x, y interface{}) bool

func TestFunc(t *testing.T) {
	f := func(x, y interface{}) bool {
		return true
	}
	tg(f, 12, 13)
}

func tg(fn ft, x, y int) {

}

type element struct {
	value string
	score float
}

func (e element) Score() float {
	return e.score
}

func TestSkiplist(t *testing.T) {
	x := NewSkipList()
	x.Insert(element{"hello", 1})
	x.Insert(element{"world", 6})
	x.Insert(element{"fuck", 8})
	x.Insert(element{"you", 9})
	x.Insert(element{"hate", 2})
	x.Insert(element{"me", 10})
	x.Insert(element{"wrong", 89})
	t.Log("Insertion completed")
	t.Log("Found 9 : ", x.Contains(element{"you", 9}))

}

func TestRand(t *testing.T) {
	z := time.Now()
	rand.Seed(z.UnixNano())
	t.Log(rand.Intn(1))
}
