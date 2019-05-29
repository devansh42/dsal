//@author Devansh Gupta
//BasicStack is the Array Stack Implementation
package dsal

import "errors"

//NewBasicStack, is the constructer for basic Stack
func NewBasicStack(size int) *BasicStack {
	b := new(BasicStack)
	b.array = make([]interface{}, size)
	return b
}

//BasicStack, is the array implementation in stack data structure
type BasicStack struct {
	array []interface{}
	i     int
}

func (s *BasicStack) Push(v interface{}) error {
	if s.i >= len(s.array) {
		return errors.New("Stackoverflow")
	}
	s.array[s.i] = v
	s.i++
	return nil
}
func (s *BasicStack) Pop() (interface{}, error) {
	if s.i <= 0 {
		return nil, errors.New("Stackunderflow")
	}
	v := s.array[s.i]
	s.i--
	return v, nil
}
func (s *BasicStack) Length() int {
	return s.i + 1
}

//google-site-verification=G5uuV911Fd454yXIF4iAZsIzrrEfqPrl1cgEkAWDnl4
