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
	if s.i < len(s.array)-1 {
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
