package dsal

import "errors"

//NewBasicStack, is the constructer for basic Stack
func NewBasicStack(size int) *BasicStack {
	b := new(BasicStack)
	b.array = make([]interface{}, size)
	b.i = -1
	return b
}

//BasicStack, is the array implementation of stack data structure
type BasicStack struct {
	array []interface{}
	//i points to current head
	i int
}

func (s *BasicStack) Push(v interface{}) (err error) {
	if s.i < len(s.array) {
		s.i++
		s.array[s.i] = v

	} else {
		err = errors.New("Stackoverflow")
	}
	return
}
func (s *BasicStack) Pop() (interface{}, error) {
	if s.i < 0 {
		return nil, errors.New("Stackunderflow")
	}
	v := s.array[s.i]
	s.i--
	return v, nil
}
func (s *BasicStack) Length() int {
	return s.i + 1
}

func (s *BasicStack) Begin() Iterator {

	return &basicStackIterator{beg: 0, end: uint(s.i), cur: 0, array: &s.array}
}

func (s *BasicStack) End() Iterator {
	return &basicStackIterator{beg: 0, end: uint(s.i), cur: uint(s.i), array: &s.array}

}

type basicStackIterator arrayIterableAttrs

func (b *basicStackIterator) Next() (v interface{}, err error) {
	ar := *b.array
	if b.HasNext() {
		v = ar[b.cur]
		b.cur++

	} else {
		err = errors.New("Iterator is out of bound")
	}
	return
}
func (b *basicStackIterator) Prev() (v interface{}, err error) {
	ar := *b.array
	if b.HasPrev() {
		v = ar[b.cur]
		b.cur--
	} else {
		err = errors.New("Iterator is out of bound")

	}
	return
}
func (b basicStackIterator) HasNext() bool {
	return b.cur <= b.end && b.inBounds()
}
func (b basicStackIterator) HasPrev() bool {
	return b.cur >= b.beg && b.inBounds()
}
func (b basicStackIterator) inBounds() bool {
	return b.cur <= b.end && b.cur >= b.beg
}
