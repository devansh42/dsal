//@author Devansh Gupta
package dsal

import "errors"

//StackList, is the stack implemented by single linked list
type StackList struct {
	slist SingleList
}

func (s *StackList) Push(v interface{}) {
	s.slist.Push(v)
}
func (s *StackList) Pop() (interface{}, error) {
	v := s.slist.Pop()
	if v == nil {
		return v, errors.New("Stack underflow")
	}
	return v, nil
}
func (s StackList) Length() int {
	return s.slist.Length()
}